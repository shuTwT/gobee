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

func (h *FlinkGroupHandlerImpl) ListFLinkGroup(c *fiber.Ctx) error {
	flinkGroups, err := h.client.FLinkGroup.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flinkGroups))
}

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
