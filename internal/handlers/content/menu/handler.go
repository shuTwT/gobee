package menu

import (
	"strconv"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"

	menu_service "github.com/shuTwT/hoshikuzu/internal/services/content/menu"

	"github.com/gofiber/fiber/v2"
)

type MenuHandler struct {
	menuService menu_service.MenuService
}

func NewMenuHandler(menuService menu_service.MenuService) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
	}
}

// @Summary 查询菜单
// @Description 查询指定ID的菜单详情
// @Tags 后台管理接口/菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} model.HttpSuccess{data=ent.Menu}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/menu/query/{id} [get]
func (h *MenuHandler) QueryMenu(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	m, err := h.menuService.QueryMenu(c, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"Menu not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", m))
}

// @Summary 获取菜单列表
// @Description 获取所有菜单列表
// @Tags 后台管理接口/菜单
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]ent.Menu}
// @Failure 500 {object} model.HttpError
// @Router /api/v1/menu/list [get]
func (h *MenuHandler) QueryMenuList(c *fiber.Ctx) error {
	menus, err := h.menuService.QueryMenuList(c)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", menus))
}

// @Summary 分页查询菜单
// @Description 分页查询菜单列表
// @Tags 后台管理接口/菜单
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Success 200 {object} model.HttpSuccess{data=model.PageResult[ent.Menu]}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/menu/page [get]
func (h *MenuHandler) QueryMenuPage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	err := c.QueryParser(&pageQuery)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	count, menus, err := h.menuService.QueryMenuPage(c, pageQuery)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	pageResult := model.PageResult[*ent.Menu]{
		Total:   int64(count),
		Records: menus,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

// @Summary 创建菜单
// @Description 创建新的菜单
// @Tags 后台管理接口/菜单
// @Accept json
// @Produce json
// @Param request body model.MenuCreateReq true "菜单信息"
// @Success 200 {object} model.HttpSuccess{data=ent.Menu}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/menu/create [post]
func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var createReq model.MenuCreateReq
	if err := c.BodyParser(&createReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	m, err := h.menuService.CreateMenu(c.Context(), createReq)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", m))
}

// @Summary 更新菜单
// @Description 更新指定ID的菜单信息
// @Tags 后台管理接口/菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Param request body model.MenuUpdateReq true "菜单信息"
// @Success 200 {object} model.HttpSuccess{data=ent.Menu}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/menu/update/{id} [put]
func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	var updateReq model.MenuUpdateReq
	if err = c.BodyParser(&updateReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	updatedMenu, err := h.menuService.UpdateMenu(c, id, updateReq)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"Menu not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", updatedMenu))
}

// @Summary 删除菜单
// @Description 删除指定ID的菜单
// @Tags 后台管理接口/菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/menu/delete/{id} [delete]
func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	err = h.menuService.DeleteMenu(c, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"Menu not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", nil))
}
