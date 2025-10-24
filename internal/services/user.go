package services

import (
	"context"
	"gobee/ent"
	"gobee/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
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
		// SetRole(user.RoleAdmin).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
