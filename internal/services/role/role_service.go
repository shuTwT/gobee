package role

import (
	"gobee/ent"
	"gobee/ent/role"
	"gobee/internal/database"

	"github.com/gofiber/fiber/v2"
)

type RoleService interface {
	QueryRole(c *fiber.Ctx, id int) (*ent.Role, error)
}

type RoleServiceImpl struct {
	client *ent.Client
}

func NewRoleServiceImpl(client *ent.Client) *RoleServiceImpl {
	return &RoleServiceImpl{client: client}
}

func (s *RoleServiceImpl) QueryRole(c *fiber.Ctx, id int) (*ent.Role, error) {
	client := database.DB
	return client.Role.Query().Where(role.IDEQ(id)).First(c.Context())

}
