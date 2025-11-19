package setting_handler

import (
	"gobee/internal/database"
	setting_service "gobee/internal/services/setting"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	settings, err := setting_service.GetAllSettings(ctx, client)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// 检查系统是否已初始化
	initialized, err := setting_service.IsSystemInitialized(ctx, client)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	settingsMap := make(map[string]string)
	for _, s := range settings {
		settingsMap[s.Key] = s.Value
	}

	return c.JSON(model.NewSuccess("success", fiber.Map{
		"settings":    settingsMap,
		"initialized": initialized,
	}))
}

func GetSettingsMap(c *fiber.Ctx) error {
	client := database.DB
	ctx := c.Context()

	// 获取所有系统设置
	settings, err := setting_service.GetAllSettings(ctx, client)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	settingsMap := make(map[string]string)
	for _, s := range settings {
		settingsMap[s.Key] = s.Value
	}

	return c.JSON(model.NewSuccess("success", settingsMap))
}

func SaveSettings(c *fiber.Ctx) error {
	client := database.DB
	ctx := c.Context()

	var req map[string]interface{}
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	for key, val := range req {
		var exist bool
		var err error
		exist, err = setting_service.ExistSettingByKey(ctx, client, key)
		if err != nil {
			log.Errorf("Error getting setting by key %s: %v", key, err)
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		if exist {
			err = setting_service.UpdateSettingByKey(ctx, client, key, val.(string))
		} else {
			err = setting_service.CreateSettingIfNotExist(ctx, client, key, val.(string))
		}
		if err != nil {
			log.Errorf("Error updating/creating setting by key %s: %v", key, err)
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}

	}

	return c.JSON(model.NewSuccess("success", nil))
}
