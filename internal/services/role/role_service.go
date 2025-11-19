package role

import (
	"gobee/ent"
	"gobee/ent/role"
	"gobee/internal/database"

	"github.com/gofiber/fiber/v2"
)

func QueryRole(c *fiber.Ctx, id int) (*ent.Role, error) {
	client := database.DB
	return client.Role.Query().Where(role.IDEQ(id)).First(c.Context())

}
