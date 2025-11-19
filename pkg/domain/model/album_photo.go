package model

import "time"

// AlbumPhoto represents the album photo domain model.
type AlbumPhoto struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ImageURL  string    `json:"image_url"`
	ViewCount int       `json:"view_count"`
	AlbumID   int       `json:"album_id"`
}

// AlbumPhotoCreateReq represents the request body for creating an album photo.

type AlbumPhotoCreateReq struct {
	ImageURL string `json:"image_url" validate:"required,url"`
	AlbumID  int    `json:"album_id" validate:"required"`
}

// AlbumPhotoUpdateReq represents the request body for updating an album photo.

type AlbumPhotoUpdateReq struct {
	ImageURL string `json:"image_url,omitempty" validate:"omitempty,url"`
	AlbumID  int    `json:"album_id,omitempty"`
}

// AlbumPhotoResp represents the response body for an album photo.

type AlbumPhotoResp struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ImageURL  string    `json:"image_url"`
	ViewCount int       `json:"view_count"`
	AlbumID   int       `json:"album_id"`
}
