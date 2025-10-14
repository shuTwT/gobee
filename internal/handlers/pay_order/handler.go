package payorder_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"gobee/ent"
	"gobee/ent/payorder"
)

// ListPayOrder 获取支付订单列表
func ListPayOrder(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	orders, err := client.PayOrder.Query().All(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": orders,
	})
}

// CreatePayOrder 创建支付订单
func CreatePayOrder(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	var order struct {
		ChannelID  string `json:"channel_id"`
		OrderID    string `json:"order_id"`
		OutTradeNo string `json:"out_trade_no"`
		TotalFee   string `json:"total_fee"`
		Subject    string `json:"subject"`
		Body       string `json:"body"`
		NotifyURL  string `json:"notify_url"`
		ReturnURL  string `json:"return_url"`
		Extra      string `json:"extra"`
		PayURL     string `json:"pay_url"`
		State      string `json:"state"`
		ErrorMsg   string `json:"error_msg"`
		Raw        string `json:"raw"`
	}
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": newOrder,
	})
}

// UpdatePayOrder 更新支付订单
func UpdatePayOrder(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var order struct {
		ChannelID  string `json:"channel_id"`
		OrderID    string `json:"order_id"`
		OutTradeNo string `json:"out_trade_no"`
		TotalFee   string `json:"total_fee"`
		Subject    string `json:"subject"`
		Body       string `json:"body"`
		NotifyURL  string `json:"notify_url"`
		ReturnURL  string `json:"return_url"`
		Extra      string `json:"extra"`
		PayURL     string `json:"pay_url"`
		State      string `json:"state"`
		ErrorMsg   string `json:"error_msg"`
		Raw        string `json:"raw"`
	}
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updatedOrder, err := client.PayOrder.UpdateOneID(id).
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": updatedOrder,
	})
}

// QueryPayOrder 查询支付订单
func QueryPayOrder(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	order, err := client.PayOrder.Query().
		Where(payorder.ID(id)).
		Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "PayOrder not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": order,
	})
}

// DeletePayOrder 删除支付订单
func DeletePayOrder(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	err = client.PayOrder.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "PayOrder not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "PayOrder deleted successfully",
	})
}
