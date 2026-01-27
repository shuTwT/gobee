package theme

import (
	"log/slog"
	"strconv"

	"github.com/shuTwT/gobee/ent"
	"github.com/shuTwT/gobee/internal/services/infra/theme"
	"github.com/shuTwT/gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

type ThemeHandler interface {
	CreateTheme(c *fiber.Ctx) error
	ListThemePage(c *fiber.Ctx) error
	QueryTheme(c *fiber.Ctx) error
	DeleteTheme(c *fiber.Ctx) error
	EnableTheme(c *fiber.Ctx) error
	DisableTheme(c *fiber.Ctx) error
}

type ThemeHandlerImpl struct {
	themeService theme.ThemeService
}

func NewThemeHandlerImpl(themeService theme.ThemeService) *ThemeHandlerImpl {
	return &ThemeHandlerImpl{themeService: themeService}
}

func (h *ThemeHandlerImpl) CreateTheme(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("Failed to get uploaded file", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusBadRequest, "获取上传文件失败"))
	}

	themeEntity, err := h.themeService.CreateTheme(c.Context(), file)
	if err != nil {
		slog.Error("Failed to create theme", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	resp := h.buildThemeResp(themeEntity)
	slog.Info("Theme created successfully", "theme_id", themeEntity.ID, "theme_name", themeEntity.Name)
	return c.JSON(model.NewSuccess("主题创建成功", resp))
}

func (h *ThemeHandlerImpl) ListThemePage(c *fiber.Ctx) error {
	var pageQuery model.PageQuery
	if err := c.QueryParser(&pageQuery); err != nil {
		slog.Error("Failed to parse query parameters", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusBadRequest, "查询参数解析失败"))
	}

	count, themes, err := h.themeService.ListThemePage(c.Context(), pageQuery.Page, pageQuery.Size)
	if err != nil {
		slog.Error("Failed to list themes", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	themeResps := make([]*model.ThemeResp, 0, len(themes))
	for _, t := range themes {
		themeResps = append(themeResps, h.buildThemeResp(t))
	}

	pageResult := model.PageResult[*model.ThemeResp]{
		Total:   int64(count),
		Records: themeResps,
	}
	return c.JSON(model.NewSuccess("主题列表获取成功", pageResult))
}

func (h *ThemeHandlerImpl) QueryTheme(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		slog.Error("Invalid theme ID", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusBadRequest, "无效的主题ID"))
	}

	themeEntity, err := h.themeService.QueryTheme(c.Context(), id)
	if err != nil {
		slog.Error("Failed to query theme", "theme_id", id, "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	resp := h.buildThemeResp(themeEntity)
	return c.JSON(model.NewSuccess("主题查询成功", resp))
}

func (h *ThemeHandlerImpl) DeleteTheme(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		slog.Error("Invalid theme ID", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusBadRequest, "无效的主题ID"))
	}

	err = h.themeService.DeleteTheme(c.Context(), id)
	if err != nil {
		slog.Error("Failed to delete theme", "theme_id", id, "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	slog.Info("Theme deleted successfully", "theme_id", id)
	return c.JSON(model.NewSuccess("主题删除成功", nil))
}

func (h *ThemeHandlerImpl) EnableTheme(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		slog.Error("Invalid theme ID", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusBadRequest, "无效的主题ID"))
	}

	err = h.themeService.EnableTheme(c.Context(), id)
	if err != nil {
		slog.Error("Failed to enable theme", "theme_id", id, "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	slog.Info("Theme enabled successfully", "theme_id", id)
	return c.JSON(model.NewSuccess("主题启用成功", nil))
}

func (h *ThemeHandlerImpl) DisableTheme(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		slog.Error("Invalid theme ID", "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusBadRequest, "无效的主题ID"))
	}

	err = h.themeService.DisableTheme(c.Context(), id)
	if err != nil {
		slog.Error("Failed to disable theme", "theme_id", id, "error", err.Error())
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	slog.Info("Theme disabled successfully", "theme_id", id)
	return c.JSON(model.NewSuccess("主题禁用成功", nil))
}

func (h *ThemeHandlerImpl) buildThemeResp(t *ent.Theme) *model.ThemeResp {
	return &model.ThemeResp{
		ID:            t.ID,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
		Name:          t.Name,
		DisplayName:   t.DisplayName,
		Description:   t.Description,
		AuthorName:    t.AuthorName,
		AuthorEmail:   t.AuthorEmail,
		Logo:          t.Logo,
		Homepage:      t.Homepage,
		Repo:          t.Repo,
		Issue:         t.Issue,
		SettingName:   t.SettingName,
		ConfigMapName: t.ConfigMapName,
		Version:       t.Version,
		Require:       t.Require,
		License:       t.License,
		Path:          t.Path,
		Enabled:       t.Enabled,
	}
}
