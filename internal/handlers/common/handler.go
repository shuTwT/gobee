package common

import (
	"gobee/pkg/domain/model"

	common_service "gobee/internal/services/common"

	"github.com/gofiber/fiber/v2"
)

type CommonHandler interface {
	GetHomeStatistics(c *fiber.Ctx) error
}

type CommonHandlerImpl struct {
	commonService common_service.CommonService
}

func NewCommonHandlerImpl(commonService common_service.CommonService) *CommonHandlerImpl {
	return &CommonHandlerImpl{
		commonService: commonService,
	}
}

func (h *CommonHandlerImpl) GetHomeStatistics(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("统计信息获取成功", h.commonService.GetHomeStatistic(c.Context())))
}
