package services

import (
	"context"
	"fmt"
	"gobee/ent"
	"gobee/ent/personalaccesstoken"
	"gobee/ent/user"
	"gobee/internal/database"
	"gobee/pkg/config"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	ListUser(c *fiber.Ctx) ([]*ent.User, error)
	ListUserPage(c *fiber.Ctx, pageQuery model.PageQuery) (int, []*ent.User, error)
	CreateUser(ctx context.Context, req model.UserCreateReq) (*ent.User, error)
	UpdateUser(ctx context.Context, id int, req model.UserUpdateReq) (*ent.User, error)
	GetUserCount(ctx context.Context) (int, error)
	DeleteUser(ctx context.Context, id int) error
	GetPersonalAccessTokenList(ctx context.Context, userId int) ([]*ent.PersonalAccessToken, error)
	GetPersonalAccessToken(ctx context.Context, userId int, id int) (*ent.PersonalAccessToken, error)
	CreatePersonalAccessToken(ctx context.Context, id int, req model.PersonalAccessTokenCreateReq) (*ent.PersonalAccessToken, error)
}

type UserServiceImpl struct {
	client *ent.Client
}

func NewUserServiceImpl(client *ent.Client) *UserServiceImpl {
	return &UserServiceImpl{client: client}
}

func (s *UserServiceImpl) ListUser(c *fiber.Ctx) ([]*ent.User, error) {
	client := s.client
	users, err := client.User.Query().All(c.Context())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) ListUserPage(c *fiber.Ctx, pageQuery model.PageQuery) (int, []*ent.User, error) {
	client := database.DB
	count, err := client.User.Query().Count(c.UserContext())

	if err != nil {
		c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
		return 0, nil, err
	}

	users, err := client.User.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
		return 0, nil, err
	}

	return count, users, err
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, req model.UserCreateReq) (*ent.User, error) {

	// 检查邮箱是否已存在
	exists, err := s.client.User.Query().
		Where(user.EmailEQ(req.Email)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("email %s already exists", req.Email)
	}

	// 如果提供了手机号，检查手机号是否已存在
	if req.PhoneNumber != "" {
		exists, err = s.client.User.Query().
			Where(user.PhoneNumberEQ(req.PhoneNumber)).
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("phone number %s already exists", req.PhoneNumber)
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 使用事务创建用户和钱包
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

	// 创建用户
	createdUser, err := tx.User.Create().
		SetName(req.Name).
		SetPassword(string(hashedPassword)).
		SetPhoneNumber(req.PhoneNumber).
		SetEmail(req.Email).
		SetRoleID(req.RoleID).
		Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	// 为用户创建钱包
	_, err = tx.Wallet.Create().
		SetUserID(createdUser.ID).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	// 为用户创建会员
	memberNo := fmt.Sprintf("M%06d", createdUser.ID)
	_, err = tx.Member.Create().
		SetUserID(createdUser.ID).
		SetMemberLevel(1).
		SetMemberNo(memberNo).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, id int, req model.UserUpdateReq) (*ent.User, error) {
	// 开始构建更新
	update := s.client.User.UpdateOneID(id)

	// 如果提供了新名称
	if req.Name != "" {
		update.SetName(req.Name)
	}

	// 如果提供了新手机号
	if req.PhoneNumber != "" {
		// 检查手机号是否已被其他用户使用
		var exists bool
		exists, err := s.client.User.Query().
			Where(
				user.And(
					user.PhoneNumberEQ(req.PhoneNumber),
					user.IDNEQ(id),
				),
			).
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("Phone number already exists")
		}
		update.SetPhoneNumber(req.PhoneNumber)
	}

	// 如果提供了新密码
	if req.Password != "" {
		var hashedPassword []byte
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("Failed to hash password")
		}
		update.SetPassword(string(hashedPassword))
	}

	// 执行更新
	updatedUser, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *UserServiceImpl) GetUserCount(ctx context.Context) (int, error) {
	client := database.DB
	count, err := client.User.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, id int) error {
	err := s.client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("User not found")
		}
		return err
	}

	return nil
}

func (s *UserServiceImpl) GetPersonalAccessTokenList(ctx context.Context, userId int) ([]*ent.PersonalAccessToken, error) {
	tokens, err := s.client.PersonalAccessToken.Query().Where(personalaccesstoken.UserID(userId)).All(ctx)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (s *UserServiceImpl) GetPersonalAccessToken(ctx context.Context, userId int, id int) (*ent.PersonalAccessToken, error) {
	token, err := s.client.PersonalAccessToken.Query().Where(personalaccesstoken.UserIDEQ(userId), personalaccesstoken.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *UserServiceImpl) CreatePersonalAccessToken(ctx context.Context, userId int, req model.PersonalAccessTokenCreateReq) (*ent.PersonalAccessToken, error) {
	// 查找用户
	u, _ := s.client.User.Query().
		Where(user.IDEQ(userId)).
		Only(ctx)

		// 生成JWT令牌
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
		"exp":   req.Expires, // 24小时过期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.GetString(config.AUTH_PAT_SECRET)

	// 使用密钥签名令牌
	t, err := token.SignedString([]byte(secret)) // 注意：在生产环境中应该使用环境变量存储密钥
	if err != nil {
		return nil, err
	}

	pat, err := s.client.PersonalAccessToken.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetExpires(req.Expires.Time()).
		SetToken(t).
		SetUserID(userId).Save(ctx)
	if err != nil {
		return nil, err
	}
	return pat, nil
}
