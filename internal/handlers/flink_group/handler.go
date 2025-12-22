package flinkgroup

import (
	"gobee/ent"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FlinkGroupHandler interface {
	ListFLinkGroup(c *fiber.Ctx) error
	CreateFlinkGroup(c *fiber.Ctx) error
	UpdateFlinkGroup(c *fiber.Ctx) error
	DeleteFLinkGroup(c *fiber.Ctx) error
}

type FlinkGroupHandlerImpl struct {
	client *ent.Client
}

func NewFlinkGroupHandlerImpl(client *ent.Client) *FlinkGroupHandlerImpl {
	return &FlinkGroupHandlerImpl{
		client: client,
	}
}

// @Summary 查询FlinkGroup
// @Description 查询FlinkGroup
// @Tags flink_groups
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]ent.FLinkGroup}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flink_groups [get]
func (h *FlinkGroupHandlerImpl) ListFLinkGroup(c *fiber.Ctx) error {
	flinkGroups, err := h.client.FLinkGroup.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flinkGroups))
}

// @Summary 创建FlinkGroup
// @Description 创建FlinkGroup
// @Tags flink_groups
// @Accept json
// @Produce json
// @Param flink_group_create_req body ent.FLinkGroup true "FlinkGroup创建请求体"
// @Success 200 {object} model.HttpSuccess{data=ent.FLinkGroup}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flink_groups [post]
func (h *FlinkGroupHandlerImpl) CreateFlinkGroup(c *fiber.Ctx) error {
	var createReq *ent.FLinkGroup
	if err := c.BodyParser(&createReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	flinkGroup, err := h.client.FLinkGroup.Create().SetName(createReq.Name).Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flinkGroup))
}

// @Summary 更新FlinkGroup
// @Description 更新FlinkGroup
// @Tags flink_groups
// @Accept json
// @Produce json
// @Param id path int true "FlinkGroup ID"
// @Param flink_group_update_req body ent.FLinkGroup true "FlinkGroup更新请求体"
// @Success 200 {object} model.HttpSuccess{data=ent.FLinkGroup}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flink_groups/{id} [put]
func (h *FlinkGroupHandlerImpl) UpdateFlinkGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid ID format"))
	}
	var updateReq *ent.FLinkGroup
	if err = c.BodyParser(&updateReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	flinkGroup, err := h.client.FLinkGroup.UpdateOneID(id).SetName(updateReq.Name).Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flinkGroup))
}

// @Summary 删除FlinkGroup
// @Description 删除FlinkGroup
// @Tags flink_groups
// @Accept json
// @Produce json
// @Param id path int true "FlinkGroup ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flink_groups/{id} [delete]
func (h *FlinkGroupHandlerImpl) DeleteFLinkGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	err = h.client.FLinkGroup.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
