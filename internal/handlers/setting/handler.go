package setting_handler

import (
	"gobee/internal/database"
	"gobee/internal/model"
	"gobee/internal/services/setting"

	"github.com/gofiber/fiber/v2"
)

// @Summary 获取系统设置
// @Description 获取所有系统设置和系统初始化状态
// @Tags settings
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /api/v1/settings [get]
func GetSettings(c *fiber.Ctx) error {
	client := database.DB
	ctx := c.Context()

	// 获取所有系统设置
	settings, err := setting.GetAllSettings(ctx, client)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// 检查系统是否已初始化
	initialized, err := setting.IsSystemInitialized(ctx, client)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", fiber.Map{
		"settings":    settings,
		"initialized": initialized,
	}))
}
