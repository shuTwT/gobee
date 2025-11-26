package friendcirclerecord

import (
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

func ListFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

func ListFriendCircleRuleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

func CreateFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

func UpdateFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

func DeleteFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}
