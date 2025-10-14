package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
)

// Protected 保护需要认证的路由
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte("your-secret-key"), // 注意：在生产环境中应该使用环境变量存储密钥
		SigningMethod:  "HS256",
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}

// jwtError 处理JWT错误
func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

// jwtSuccess 处理JWT验证成功
func jwtSuccess(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	
	// 将用户信息存储到上下文中
	c.Locals("userId", claims["id"])
	c.Locals("userEmail", claims["email"])
	c.Locals("userName", claims["name"])
	
	return c.Next()
}

// GetCurrentUser 获取当前登录用户信息
func GetCurrentUser(c *fiber.Ctx) fiber.Map {
	return fiber.Map{
		"id":    c.Locals("userId"),
		"email": c.Locals("userEmail"),
		"name":  c.Locals("userName"),
	}
}
