package user_handlers

import (
	"gobee/ent"

	"github.com/gofiber/fiber/v2"
)

func FindUser(ctx *fiber.Ctx, client *ent.Client) ([]*ent.User, error) {
	user, err := client.User.Query().All(ctx.Context())
	if err != nil {
		return nil, err
	}

	return user, nil
}
