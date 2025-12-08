package common

import (
	"gobee/pkg/domain/model"

	common_service "gobee/internal/services/common"

	"github.com/gofiber/fiber/v2"
)

func GetHomeStatistics(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("统计信息获取成功", common_service.GetHomeStatistic(c.Context())))
}
