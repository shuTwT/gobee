package model

import "time"

// Album represents the album domain model.
type Album struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Sort        int       `json:"sort"`
}

// AlbumCreateReq represents the request body for creating an album.

type AlbumCreateReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Sort        int    `json:"sort"`
}

// AlbumUpdateReq represents the request body for updating an album.

type AlbumUpdateReq struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Sort        int    `json:"sort,omitempty"`
}

// AlbumResp represents the response body for an album.

type AlbumResp struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Sort        int       `json:"sort"`
}
