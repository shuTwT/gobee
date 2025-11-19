package model

import "time"

// File represents the file domain model.
type File struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	URL       string    `json:"url"`
	Type      string    `json:"type"`
	Size      string    `json:"size"`
}

// FileCreateReq represents the request body for creating a file.

type FileCreateReq struct {
	Name string `json:"name" validate:"required"`
	Path string `json:"path"`
	URL  string `json:"url" validate:"required,url"`
	Type string `json:"type"`
	Size string `json:"size"`
}

// FileUpdateReq represents the request body for updating a file.

type FileUpdateReq struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

// FileResp represents the response body for a file.

type FileResp struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	URL       string    `json:"url"`
	Type      string    `json:"type"`
	Size      string    `json:"size"`
}