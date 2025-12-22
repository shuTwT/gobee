package payorder_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"gobee/ent"
	"gobee/ent/payorder"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
)

type PayOrderHandler interface {
	ListPayOrder(c *fiber.Ctx) error
	CreatePayOrder(c *fiber.Ctx) error
	UpdatePayOrder(c *fiber.Ctx) error
	QueryPayOrder(c *fiber.Ctx) error
	DeletePayOrder(c *fiber.Ctx) error
}

type PayOrderHandlerImpl struct {
	client *ent.Client
}

func NewPayOrderHandlerImpl(client *ent.Client) *PayOrderHandlerImpl {
	return &PayOrderHandlerImpl{client: client}
}

// @Summary 获取支付订单列表
// @Description 获取所有支付订单的列表
// @Tags payorders
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]ent.PayOrder}
// @Failure 500 {object} model.HttpError
// @Router /api/v1/payorders [get]
func (h *PayOrderHandlerImpl) ListPayOrder(c *fiber.Ctx) error {
	orders, err := h.client.PayOrder.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", orders))
}

// @Summary 创建支付订单
// @Description 创建一个新的支付订单
// @Tags payorders
// @Accept json
// @Produce json
// @Param payorder body ent.PayOrder true "支付订单信息"
// @Success 201 {object} ent.PayOrder
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/payorders [post]
func (h *PayOrderHandlerImpl) CreatePayOrder(c *fiber.Ctx) error {
	client := database.DB
	var order *model.PayOrderCreateReq
	if err := c.BodyParser(&order); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	newOrder, err := client.PayOrder.Create().
		SetChannelID(order.ChannelID).
		SetOrderID(order.OrderID).
		SetOutTradeNo(order.OutTradeNo).
		SetTotalFee(order.TotalFee).
		SetSubject(order.Subject).
		SetBody(order.Body).
		SetNotifyURL(order.NotifyURL).
		SetReturnURL(order.ReturnURL).
		SetExtra(order.Extra).
		SetPayURL(order.PayURL).
		SetState(order.State).
		SetErrorMsg(order.ErrorMsg).
		SetRaw(order.Raw).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newOrder))
}

// @Summary 更新支付订单
// @Description 更新指定支付订单的信息
// @Tags payorders
// @Accept json
// @Produce json
// @Param id path string true "支付订单ID"
// @Param payorder body ent.PayOrder true "支付订单信息"
// @Success 200 {object} ent.PayOrder
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/payorders/{id} [put]
func (h *PayOrderHandlerImpl) UpdatePayOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}

	var order *model.PayOrderUpdateReq
	if err = c.BodyParser(&order); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	updatedOrder, err := h.client.PayOrder.UpdateOneID(id).
		SetChannelID(order.ChannelID).
		SetOrderID(order.OrderID).
		SetOutTradeNo(order.OutTradeNo).
		SetTotalFee(order.TotalFee).
		SetSubject(order.Subject).
		SetBody(order.Body).
		SetNotifyURL(order.NotifyURL).
		SetReturnURL(order.ReturnURL).
		SetExtra(order.Extra).
		SetPayURL(order.PayURL).
		SetState(order.State).
		SetErrorMsg(order.ErrorMsg).
		SetRaw(order.Raw).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", updatedOrder))
}

// @Summary 查询支付订单
// @Description 查询指定支付订单的详细信息
// @Tags payorders
// @Accept json
// @Produce json
// @Param id path string true "支付订单ID"
// @Success 200 {object} ent.PayOrder
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/payorders/{id} [get]
func (h *PayOrderHandlerImpl) QueryPayOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	order, err := h.client.PayOrder.Query().
		Where(payorder.ID(id)).
		Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"PayOrder not found"))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", order))
}

// @Summary 删除支付订单
// @Description 删除指定支付订单
// @Tags payorders
// @Accept json
// @Produce json
// @Param id path string true "支付订单ID"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/payorders/{id} [delete]
func (h *PayOrderHandlerImpl) DeletePayOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	err = h.client.PayOrder.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusBadRequest,
				"PayOrder not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success",
		nil,
	))
}
