/**
 * 用户
 */
package model

import "time"

// User represents the user domain model.
type User struct {
	ID                  int       `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Email               string    `json:"email"`
	EmailVerified       bool      `json:"email_verified"`
	Name                string    `json:"name"`
	PhoneNumber         *string   `json:"phone_number,omitempty"`
	PhoneNumberVerified bool      `json:"phone_number_verified"`
	Password            string    `json:"-"` // Sensitive field, excluded from JSON output
	RoleID              *int      `json:"role_id,omitempty"`
	Role                *Role     `json:"role,omitempty"`
}

// UserCreateReq represents the request body for creating a user.

type UserCreateReq struct {
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number,omitempty"`
	RoleID      int    `json:"role_id,omitempty"`
}

// UserUpdateReq represents the request body for updating a user.

type UserUpdateReq struct {
	Name        string `json:"name,omitempty"`
	Password    string `json:"password,omitempty" validate:"min=8"`
	PhoneNumber string `json:"phone_number,omitempty"`
	RoleID      int    `json:"role_id,omitempty"`
}

// UserResp represents the response body for a user.

type UserResp struct {
	ID                  int       `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	Email               string    `json:"email"`
	EmailVerified       bool      `json:"email_verified"`
	Name                string    `json:"name"`
	PhoneNumber         *string   `json:"phone_number,omitempty"`
	PhoneNumberVerified bool      `json:"phone_number_verified"`
	RoleID              *int      `json:"role_id,omitempty"`
	Role                *Role     `json:"role,omitempty"`
}
