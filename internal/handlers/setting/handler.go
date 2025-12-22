package setting_handler

import (
	"encoding/json"
	"gobee/internal/database"
	setting_service "gobee/internal/services/setting"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type SettingHandler interface {
	GetSettings(c *fiber.Ctx) error
	GetJsonSettingsMap(c *fiber.Ctx) error
	SaveSettings(c *fiber.Ctx) error
}

type SettingHandlerImpl struct {
	settingService setting_service.SettingService
}

func NewSettingHandlerImpl(settingService setting_service.SettingService) *SettingHandlerImpl {
	return &SettingHandlerImpl{
		settingService: settingService,
	}
}

// @Summary 获取系统设置
// @Description 获取所有系统设置和系统初始化状态
// @Tags settings
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /api/v1/settings [get]
func (h *SettingHandlerImpl) GetSettings(c *fiber.Ctx) error {
	client := database.DB
	ctx := c.Context()

	// 获取所有系统设置
	settings, err := h.settingService.GetAllSettings(ctx, client)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// 检查系统是否已初始化
	initialized, err := h.settingService.IsSystemInitialized(ctx, client)
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

func (h *SettingHandlerImpl) GetJsonSettingsMap(c *fiber.Ctx) error {
	client := database.DB
	ctx := c.Context()

	key := c.Params("key")

	var exist bool
	var err error
	exist, err = h.settingService.ExistSettingByKey(ctx, client, key)
	if err != nil {
		log.Errorf("Error getting setting by key %s: %v", key, err)
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	if !exist {
		return c.JSON(model.NewSuccess("Setting not found", map[string]any{}))
	}

	// 获取所有系统设置
	setting, err := h.settingService.GetSettingByKey(ctx, client, key)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	var value any
	if err := json.Unmarshal([]byte(setting.Value), &value); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", value))
}

func (h *SettingHandlerImpl) SaveSettings(c *fiber.Ctx) error {
	client := database.DB
	ctx := c.Context()

	key := c.Params("key")

	var req map[string]interface{}

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	var exist bool
	var err error
	exist, err = h.settingService.ExistSettingByKey(ctx, client, key)
	if err != nil {
		log.Errorf("Error getting setting by key %s: %v", key, err)
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	value, err := json.Marshal(req)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	if exist {
		err = h.settingService.UpdateSettingByKey(ctx, client, key, string(value))
	} else {
		err = h.settingService.CreateSettingIfNotExist(ctx, client, key, string(value))
	}
	if err != nil {
		log.Errorf("Error updating/creating setting by key %s: %v", key, err)
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", nil))
}
