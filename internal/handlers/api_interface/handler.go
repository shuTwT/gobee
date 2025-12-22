package apiinterface

import (
	"gobee/ent"
	api_interface_service "gobee/internal/services/api_interface"

	"github.com/gofiber/fiber/v2"
)

type ApiInterfaceHandler interface {
	ListApiRoutesPage(c *fiber.Ctx) error
}

type ApiInterfaceHandlerImpl struct {
	client              *ent.Client
	apiInterfaceService api_interface_service.ApiInterfaceService
}

func NewApiInterfaceHandlerImpl(client *ent.Client, apiInterfaceService api_interface_service.ApiInterfaceService) *ApiInterfaceHandlerImpl {
	return &ApiInterfaceHandlerImpl{
		client:              client,
		apiInterfaceService: apiInterfaceService,
	}
}

func (h *ApiInterfaceHandlerImpl) ListApiRoutesPage(c *fiber.Ctx) error {
	return h.apiInterfaceService.ListApiRoutesPage(c)
}
