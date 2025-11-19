package model

import "time"

// Page represents the page domain model.
type Page struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Description *string   `json:"description,omitempty"`
}

// PageCreateReq represents the request body for creating a page.

type PageCreateReq struct {
	Title       string  `json:"title" validate:"required"`
	Content     string  `json:"content" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// PageUpdateReq represents the request body for updating a page.

type PageUpdateReq struct {
	Title       *string `json:"title,omitempty"`
	Content     *string `json:"content,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PageResp represents the response body for a page.

type PageResp struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Description *string   `json:"description,omitempty"`
}