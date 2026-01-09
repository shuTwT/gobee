package role

import (
	"gobee/ent"
	"gobee/ent/role"
	"gobee/internal/database"
	role_service "gobee/internal/services/system/role"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler interface {
	ListRole(c *fiber.Ctx) error
	ListRolePage(c *fiber.Ctx) error
	CreateRole(c *fiber.Ctx) error
	UpdateRole(c *fiber.Ctx) error
	QueryRole(c *fiber.Ctx) error
	DeleteRole(c *fiber.Ctx) error
}

type RoleHandlerImpl struct {
	roleService role_service.RoleService
}

func NewRoleHandlerImpl(roleService role_service.RoleService) *RoleHandlerImpl {
	return &RoleHandlerImpl{
		roleService: roleService,
	}
}

// @Summary 查询所有角色
// @Description 查询系统中所有角色的列表
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]ent.Role}
// @Failure 500 {object} model.HttpError
// @Router /api/v1/roles [get]
func (h *RoleHandlerImpl) ListRole(c *fiber.Ctx) error {
	client := database.DB
	roles, err := client.Role.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", roles))
}

// @Summary 查询角色分页列表
// @Description 查询系统中角色的分页列表
// @Tags roles
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Success 200 {object} model.HttpSuccess{data=model.PageResult[ent.Role]}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/roles/page [get]
func (h *RoleHandlerImpl) ListRolePage(c *fiber.Ctx) error {
	client := database.DB
	pageQuery := model.PageQuery{}
	err := c.QueryParser(&pageQuery)

	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	count, err := client.Role.Query().Count(c.UserContext())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	roles, err := client.Role.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}
	pageResult := model.PageResult[*ent.Role]{
		Total:   int64(count),
		Records: roles,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

// @Summary 创建角色
// @Description 创建一个新的角色
// @Tags roles
// @Accept json
// @Produce json
// @Param createReq body model.RoleCreateReq true "角色创建请求"
// @Success 200 {object} model.HttpSuccess{data=ent.Role}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/roles [post]
func (h *RoleHandlerImpl) CreateRole(c *fiber.Ctx) error {
	client := database.DB
	var roleData *model.RoleCreateReq
	if err := c.BodyParser(&roleData); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	// 检查角色代码是否已存在
	exists, err := client.Role.Query().
		Where(role.CodeEQ(roleData.Code)).
		Exist(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}
	if exists {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Role code already exists",
		))
	}

	role, err := client.Role.Create().SetName(roleData.Name).SetCode(roleData.Code).Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", role))
}

// @Summary 更新角色
// @Description 更新指定角色的信息
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param updateReq body model.RoleUpdateReq true "角色更新请求"
// @Success 200 {object} model.HttpSuccess{data=ent.Role}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/roles/{id} [put]
func (h *RoleHandlerImpl) UpdateRole(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	var roleData *model.RoleUpdateReq
	if err = c.BodyParser(&roleData); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	// 检查角色代码是否已存在
	exists, err := client.Role.Query().
		Where(role.CodeEQ(roleData.Code)).
		Exist(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}
	if exists {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Role code already exists",
		))
	}
	// 开始构建更新
	update := client.Role.UpdateOneID(id)

	// 如果提供了新名称
	if roleData.Name != "" {
		update.SetName(roleData.Name)
	}

	// 如果提供了新代码
	if roleData.Code != "" {
		update.SetCode(roleData.Code)
	}

	role, err := update.Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", role))
}

// @Summary 查询角色
// @Description 查询指定角色的详细信息
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} model.HttpSuccess{data=ent.Role}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/roles/{id} [get]
func (h *RoleHandlerImpl) QueryRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}
	role, err := h.roleService.QueryRole(c, id)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}
	return c.JSON(model.NewSuccess("success", role))
}

// @Summary 删除角色
// @Description 删除指定角色
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/roles/{id} [delete]
func (h *RoleHandlerImpl) DeleteRole(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}
	err = client.Role.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
