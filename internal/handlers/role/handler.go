package role

import (
	"gobee/ent"
	"gobee/ent/role"
	"gobee/internal/database"
	role_service "gobee/internal/services/role"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListRole(c *fiber.Ctx) error {
	client := database.DB
	roles, err := client.Role.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", roles))
}

func ListRolePage(c *fiber.Ctx) error {
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

func CreateRole(c *fiber.Ctx) error {
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

func UpdateRole(c *fiber.Ctx) error {
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

func QueryRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}
	role, err := role_service.QueryRole(c, id)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}
	return c.JSON(model.NewSuccess("success", role))
}

func DeleteRole(c *fiber.Ctx) error {
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
