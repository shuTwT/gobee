package router

import (
	auth_handlers "gobee/internal/handlers/auth"
	paychannel_handler "gobee/internal/handlers/pay_channel"
	"gobee/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	console := router.Group("/console")
	{
		// 登录页面
		console.Get("/login", func(c *fiber.Ctx) error {
			return c.Render("pages/console/login", fiber.Map{
				"Title": "登录",
			}, "layouts/base")
		})
		// console.Use(middleware.ConsoleProtected)
		// 首页
		console.Get("/", func(c *fiber.Ctx) error {
			return c.Render("pages/console/index", fiber.Map{
				"Title": "Hello, World!",
			}, "layouts/admin", "layouts/base")
		})
		console.Get("/content_models", func(c *fiber.Ctx) error {
			return c.Render("pages/console/model-schema", fiber.Map{
				"Title": "内容模型",
			}, "layouts/admin", "layouts/base")
		})
		console.Get("/content", func(c *fiber.Ctx) error {
			return c.Render("pages/console/model-content", fiber.Map{
				"Title": "内容",
			}, "layouts/admin", "layouts/base")
		})
		// 相册管理
		console.Get("/album", func(c *fiber.Ctx) error {
			return c.Render("pages/console/album", fiber.Map{
				"Title": "相册管理",
			}, "layouts/admin", "layouts/base")
		})
		// 评论管理
		console.Get("/comment", func(c *fiber.Ctx) error {
			return c.Render("pages/console/comment", fiber.Map{
				"Title": "评论管理",
			}, "layouts/admin", "layouts/base")
		})
		// 文件管理
		console.Get("/file", func(c *fiber.Ctx) error {
			return c.Render("pages/console/file", fiber.Map{
				"Title": "文件管理",
			}, "layouts/admin", "layouts/base")
		})
		// 友链管理
		console.Get("/flink", func(c *fiber.Ctx) error {
			return c.Render("pages/console/flink", fiber.Map{
				"Title": "友链管理",
			}, "layouts/admin", "layouts/base")
		})
		// webhook
		console.Get("/webhook", func(c *fiber.Ctx) error {
			return c.Render("pages/console/webhook", fiber.Map{
				"Title": "Webhook",
			}, "layouts/admin", "layouts/base")
		})
		// 用户中心
		console.Get("/user-center", func(c *fiber.Ctx) error {
			return c.Render("pages/console/user-center", fiber.Map{
				"Title": "用户中心",
			}, "layouts/admin", "layouts/base")
		})
		// 用户管理
		console.Get("/users", func(c *fiber.Ctx) error {
			return c.Render("pages/console/users", fiber.Map{
				"Title": "用户管理",
			}, "layouts/admin", "layouts/base")
		})
	}

	api := router.Group("/api")
	{
		apiV1 := api.Group("/v1")
		{
			apiV1.Get("/users", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"message": "Hello, World!",
				})
			})
			payChannelApi := apiV1.Group("/pay-channel")
			{
				payChannelApi.Get("/list", paychannel_handler.ListPayChannel)
			}
		}
	}

	auth := router.Group("/auth")
	{
		auth.Post("/login/password", auth_handlers.LoginPassword)
	}

	router.Get("/initialize", func(c *fiber.Ctx) error {
		return c.Render("pages/initialize", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/base")
	})
}
