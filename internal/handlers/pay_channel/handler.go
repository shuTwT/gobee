package pay_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"gobee/ent"
	"gobee/ent/paychannel"
	"gobee/internal/database"
	"gobee/internal/model"
)

// @Summary 获取支付渠道列表
// @Description 获取所有支付渠道的列表
// @Tags paychannels
// @Accept json
// @Produce json
// @Success 200 {object} []ent.PayChannel
// @Failure 500 {object} model.HttpError
// @Router /api/v1/paychannels [get]
func ListPayChannel(c *fiber.Ctx) error {
	client := database.DB
	channels, err := client.PayChannel.Query().All(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", channels))
}

// @Summary 创建支付渠道
// @Description 创建一个新的支付渠道
// @Tags paychannels
// @Accept json
// @Produce json
// @Param paychannel body ent.PayChannel true "支付渠道信息"
// @Success 200 {object} ent.PayChannel
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/paychannels [post]
func CreatePayChannel(c *fiber.Ctx) error {
	client := database.DB
	var channel struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Type   string `json:"type"`
		Config string `json:"config"`
	}
	if err := c.BodyParser(&channel); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	newChannel, err := client.PayChannel.Create().
		SetName(channel.Name).
		SetCode(channel.Code).
		SetType(channel.Type).
		SetConfig(channel.Config).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newChannel))
}

// @Summary 更新支付渠道
// @Description 更新指定支付渠道的信息
// @Tags paychannels
// @Accept json
// @Produce json
// @Param id path string true "支付渠道ID"
// @Param paychannel body ent.PayChannel true "支付渠道信息"
// @Success 200 {object} ent.PayChannel
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/paychannels/{id} [put]
func UpdatePayChannel(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid ID format"))
	}

	var channel struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Type   string `json:"type"`
		Config string `json:"config"`
	}
	if err := c.BodyParser(&channel); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	updatedChannel, err := client.PayChannel.UpdateOneID(id).
		SetName(channel.Name).
		SetCode(channel.Code).
		SetType(channel.Type).
		SetConfig(channel.Config).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", updatedChannel))
}

// @Summary 查询支付渠道
// @Description 查询指定支付渠道的详细信息
// @Tags paychannels
// @Accept json
// @Produce json
// @Param id path string true "支付渠道ID"
// @Success 200 {object} ent.PayChannel
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/paychannels/{id} [get]
func QueryPayChannel(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	channel, err := client.PayChannel.Query().
		Where(paychannel.ID(id)).
		Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound, "ayChannel not found"))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", channel))
}

// @Summary 删除支付渠道
// @Description 删除指定支付渠道
// @Tags paychannels
// @Accept json
// @Produce json
// @Param id path string true "支付渠道ID"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/paychannels/{id} [delete]
func DeletePayChannel(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	err = client.PayChannel.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"PayChannel not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
