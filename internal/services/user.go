package services

type CreateUserRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
