package model

import "time"

// ApiPerms represents the API permissions domain model.
type ApiPerms struct {
	ID             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	Path           string    `json:"path"`
	Method         string    `json:"method"`
	Desc           string    `json:"desc"`
	PermissionType string    `json:"permission_type"`
	Roles          []string  `json:"roles"`
	Status         string    `json:"status"`
}

// ApiPermsCreateReq represents the request body for creating an API permission.

type ApiPermsCreateReq struct {
	Name           string   `json:"name" validate:"required"`
	Path           string   `json:"path" validate:"required"`
	Method         string   `json:"method" validate:"required"`
	Desc           string   `json:"desc"`
	PermissionType string   `json:"permission_type" validate:"required,oneof=private public"`
	Roles          []string `json:"roles"`
	Status         string   `json:"status" validate:"required,oneof=active inactive"`
}

// ApiPermsUpdateReq represents the request body for updating an API permission.

type ApiPermsUpdateReq struct {
	Name           *string  `json:"name,omitempty"`
	Path           *string  `json:"path,omitempty"`
	Method         *string  `json:"method,omitempty"`
	Desc           *string  `json:"desc,omitempty"`
	PermissionType *string  `json:"permission_type,omitempty" validate:"omitempty,oneof=private public"`
	Roles          []string `json:"roles,omitempty"`
	Status         *string  `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}

// ApiPermsResp represents the response body for an API permission.

type ApiPermsResp struct {
	ID             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
	Path           string    `json:"path"`
	Method         string    `json:"method"`
	Desc           string    `json:"desc"`
	PermissionType string    `json:"permission_type"`
	Roles          []string  `json:"roles"`
	Status         string    `json:"status"`
}