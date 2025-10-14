package pay_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"gobee/ent"
	"gobee/ent/paychannel"
)

// ListPayChannel 获取支付渠道列表
func ListPayChannel(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	channels, err := client.PayChannel.Query().All(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": channels,
	})
}

// CreatePayChannel 创建支付渠道
func CreatePayChannel(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	var channel struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Type   string `json:"type"`
		Config string `json:"config"`
	}
	if err := c.BodyParser(&channel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newChannel, err := client.PayChannel.Create().
		SetName(channel.Name).
		SetCode(channel.Code).
		SetType(channel.Type).
		SetConfig(channel.Config).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": newChannel,
	})
}

// UpdatePayChannel 更新支付渠道
func UpdatePayChannel(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var channel struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Type   string `json:"type"`
		Config string `json:"config"`
	}
	if err := c.BodyParser(&channel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updatedChannel, err := client.PayChannel.UpdateOneID(id).
		SetName(channel.Name).
		SetCode(channel.Code).
		SetType(channel.Type).
		SetConfig(channel.Config).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": updatedChannel,
	})
}

// QueryPayChannel 查询支付渠道
func QueryPayChannel(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	channel, err := client.PayChannel.Query().
		Where(paychannel.ID(id)).
		Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "PayChannel not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": channel,
	})
}

// DeletePayChannel 删除支付渠道
func DeletePayChannel(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	err = client.PayChannel.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "PayChannel not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "PayChannel deleted successfully",
	})
}
