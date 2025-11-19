package services

import (
	"context"
	"gobee/ent"
	"gobee/internal/database"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
	RoleId   int
}

func ListUser(c *fiber.Ctx) ([]*ent.User, error) {
	client := database.DB
	users, err := client.User.Query().All(c.Context())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func ListUserPage(c *fiber.Ctx, pageQuery model.PageQuery) (int, []*ent.User, error) {
	client := database.DB
	count, err := client.User.Query().Count(c.UserContext())

	if err != nil {
		c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
		return 0, nil, err
	}

	users, err := client.User.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
		return 0, nil, err
	}

	return count, users, err
}

func CreateUser(req CreateUserRequest) (*ent.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	createdUser, err := database.DB.User.Create().
		SetName(req.Username).
		SetPassword(string(hashedPassword)).
		SetEmail(req.Email).
		SetRoleID(req.RoleId).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
