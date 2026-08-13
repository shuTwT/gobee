package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/refreshtoken"
	"github.com/shuTwT/hoshikuzu/ent/user"
	"github.com/shuTwT/hoshikuzu/pkg/config"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenTTL   = 24 * time.Hour
	refreshTokenTTL  = 7 * 24 * time.Hour
	refreshTokenSize = 32 // 字节数，生成 64 位十六进制字符串
)

type AuthService interface {
	Login(ctx context.Context, req *model.LoginRequest, userAgent, ip string) (*model.LoginResp, error)
	RefreshToken(ctx context.Context, req *model.RefreshTokenRequest, userAgent, ip string) (*model.RefreshTokenResp, error)
	RevokeRefreshToken(ctx context.Context, rawToken string) error
}

type AuthServiceImpl struct {
	client *ent.Client
}

func NewAuthServiceImpl(client *ent.Client) *AuthServiceImpl {
	return &AuthServiceImpl{client: client}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *model.LoginRequest, userAgent, ip string) (*model.LoginResp, error) {
	u, err := s.client.User.Query().
		Where(user.EmailEQ(req.Email)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("找不到该用户")
		}
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误")
	}
	role, err := u.QueryRole().Only(ctx)
	if err != nil {
		return nil, errors.New("用户角色不存在")
	}

	return s.issueTokens(ctx, u, role.Code, userAgent, ip)
}

// RefreshToken 使用有效 refresh token 轮换签发新的 access/refresh token。
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *model.RefreshTokenRequest, userAgent, ip string) (*model.RefreshTokenResp, error) {
	hash := hashRefreshToken(req.RefreshToken)
	rt, err := s.client.RefreshToken.Query().
		Where(
			refreshtoken.TokenHash(hash),
			refreshtoken.Revoked(false),
			refreshtoken.ExpiresAtGT(time.Now()),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("refresh token 无效或已过期")
		}
		return nil, err
	}

	u, err := s.client.User.Get(ctx, rt.UserID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	role, err := u.QueryRole().Only(ctx)
	if err != nil {
		return nil, errors.New("用户角色不存在")
	}

	accessToken, expires, err := s.generateAccessToken(u)
	if err != nil {
		return nil, err
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	// 吊销旧 refresh token
	if err := tx.RefreshToken.UpdateOne(rt).SetRevoked(true).Exec(ctx); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	// 签发新 refresh token
	raw, newHash, expiresAt, err := generateRefreshToken()
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if _, err := tx.RefreshToken.Create().
		SetTokenHash(newHash).
		SetUserID(u.ID).
		SetExpiresAt(expiresAt).
		SetUserAgent(userAgent).
		SetIP(ip).
		Save(ctx); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.RefreshTokenResp{
		AccessToken:  accessToken,
		RefreshToken: raw,
		Expires:      expires,
		Username:     u.Name,
		Roles:        []string{roleCodeForUI(role.Code)},
	}, nil
}

// RevokeRefreshToken 吊销指定 refresh token，用于登出。找不到也不报错。
func (s *AuthServiceImpl) RevokeRefreshToken(ctx context.Context, rawToken string) error {
	if rawToken == "" {
		return nil
	}
	hash := hashRefreshToken(rawToken)
	_, err := s.client.RefreshToken.Update().
		Where(refreshtoken.TokenHash(hash)).
		SetRevoked(true).
		Save(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	return nil
}

// issueTokens 生成 access token 与 refresh token 并持久化 refresh token。
func (s *AuthServiceImpl) issueTokens(ctx context.Context, u *ent.User, roleCode, userAgent, ip string) (*model.LoginResp, error) {
	accessToken, expires, err := s.generateAccessToken(u)
	if err != nil {
		return nil, err
	}

	raw, hash, expiresAt, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	if _, err := s.client.RefreshToken.Create().
		SetTokenHash(hash).
		SetUserID(u.ID).
		SetExpiresAt(expiresAt).
		SetUserAgent(userAgent).
		SetIP(ip).
		Save(ctx); err != nil {
		return nil, err
	}

	return &model.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: raw,
		Expires:      expires,
		Username:     u.Name,
		Roles:        []string{roleCodeForUI(roleCode)},
	}, nil
}

// generateAccessToken 生成登录态 JWT，返回 token 字符串与毫秒级过期时间戳。
func (s *AuthServiceImpl) generateAccessToken(u *ent.User) (string, int64, error) {
	expiresAt := time.Now().Add(accessTokenTTL)
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
		"exp":   expiresAt.Unix(), // JWT 规范：秒级时间戳
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.GetString(config.AUTH_TOKEN_SECRET)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return t, expiresAt.UnixMilli(), nil
}

// generateRefreshToken 生成随机 refresh token 明文、SHA256 哈希与过期时间。
func generateRefreshToken() (raw, hash string, expiresAt time.Time, err error) {
	b := make([]byte, refreshTokenSize)
	if _, err = rand.Read(b); err != nil {
		return "", "", time.Time{}, err
	}
	raw = hex.EncodeToString(b)
	sum := sha256.Sum256([]byte(raw))
	hash = hex.EncodeToString(sum[:])
	expiresAt = time.Now().Add(refreshTokenTTL)
	return raw, hash, expiresAt, nil
}

func hashRefreshToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

// roleCodeForUI 将后端角色码映射为前端使用的展示码。
func roleCodeForUI(roleCode string) string {
	// The UI historically calls the super administrator role "admin". Keep
	// that presentation label while no longer granting it to every user.
	if roleCode == "superAdmin" {
		return "admin"
	}
	return roleCode
}
