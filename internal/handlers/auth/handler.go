package auth_handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"gobee/ent"
	"gobee/ent/user"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
)

// LoginRequest 登录请求结构
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// @Summary 用户登录
// @Description 验证用户凭据并返回JWT令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "登录请求"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 401 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /auth/login/password [post]
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(
			fiber.StatusBadRequest,
			"Invalid request body",
		))
	}

	// 获取数据库客户端
	client := database.DB

	// 查找用户
	u, err := client.User.Query().
		Where(user.EmailEQ(req.Email)).
		Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(
				fiber.StatusUnauthorized,
				"找不到该用户",
			))
		}
		return c.JSON(model.NewError(
			fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return c.JSON(model.NewError(
			fiber.StatusUnauthorized,
			"Invalid credentials",
		))
	}

	expires := time.Now().Add(time.Hour * 24).UnixMilli()
	// 生成JWT令牌
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
		"exp":   expires, // 24小时过期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名令牌
	t, err := token.SignedString([]byte("your-secret-key")) // 注意：在生产环境中应该使用环境变量存储密钥
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	// 返回令牌
	return c.JSON(model.NewSuccess("Login successful", fiber.Map{
		"accessToken": t,
		"expires":     expires,
		"username":    u.Name,
		"roles":       []string{"admin"},
	}))
}
