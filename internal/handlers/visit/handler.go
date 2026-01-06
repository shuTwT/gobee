package visit

import (
	"gobee/internal/services/visit"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

type VisitHandler interface {
	HandleVisitor(c *fiber.Ctx) error
}

type VisitHandlerImpl struct {
	visitService visit.VisitService
}

func NewVisitHandlerImpl(visitService visit.VisitService) VisitHandler {
	return &VisitHandlerImpl{visitService: visitService}
}

func (h *VisitHandlerImpl) HandleVisitor(c *fiber.Ctx) error {
	var req model.VisitLogReq
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return h.visitService.CreateVisitLog(c.Context(), req)
}
