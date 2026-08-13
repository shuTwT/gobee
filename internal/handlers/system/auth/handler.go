package auth_handler

import (
	"github.com/shuTwT/hoshikuzu/internal/services/system/auth"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// @Summary 用户登录
// @Description 验证用户凭据并返回JWT令牌
// @Tags 公开接口/认证
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "登录请求"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 401 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/auth/login/password [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req *model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(
			fiber.StatusBadRequest,
			"Invalid request body",
		))
	}

	loginResp, err := h.authService.Login(c.Context(), req, c.Get("User-Agent"), c.IP())
	if err != nil {
		return c.JSON(model.NewError(
			fiber.StatusUnauthorized,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("Login successful", loginResp))
}

// @Summary 刷新令牌
// @Description 使用refresh token换取新的access/refresh token
// @Tags 公开接口/认证
// @Accept json
// @Produce json
// @Param body body model.RefreshTokenRequest true "刷新请求"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 401 {object} model.HttpError
// @Router /api/auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req model.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid request body"))
	}
	if req.RefreshToken == "" {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "refreshToken is required"))
	}

	resp, err := h.authService.RefreshToken(c.Context(), &req, c.Get("User-Agent"), c.IP())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusUnauthorized, err.Error()))
	}
	return c.JSON(model.NewSuccess("Token refreshed", resp))
}

// @Summary 登出
// @Description 吊销当前refresh token
// @Tags 公开接口/认证
// @Accept json
// @Produce json
// @Param body body model.RefreshTokenRequest true "登出请求"
// @Success 200 {object} model.HttpSuccess
// @Router /api/auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var req model.RefreshTokenRequest
	// 登出时 refreshToken 可缺省，缺省时直接返回成功
	_ = c.BodyParser(&req)
	if req.RefreshToken != "" {
		_ = h.authService.RevokeRefreshToken(c.Context(), req.RefreshToken)
	}
	return c.JSON(model.NewSuccess("Logout successful", nil))
}
