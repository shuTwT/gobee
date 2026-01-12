package model

type DocLibraryCreateReq struct {
	Name        string `json:"name" validate:"required"`
	Alias       string `json:"alias" validate:"required"`
	Description string `json:"description"`
	Source      string `json:"source" validate:"required,oneof=git openapi llms_txt website"`
	URL         string `json:"url" validate:"required"`
}

type DocLibraryUpdateReq struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Alias       string `json:"alias" validate:"required"`
	Description string `json:"description"`
	Source      string `json:"source" validate:"required,oneof=git openapi llms_txt website"`
	URL         string `json:"url" validate:"required"`
}

type DocLibraryResp struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Alias       string     `json:"alias"`
	Description string     `json:"description"`
	Source      string     `json:"source"`
	URL         string     `json:"url"`
	CreatedAt   *LocalTime `json:"created_at"`
	UpdatedAt   *LocalTime `json:"updated_at"`
}

type DocLibraryTreeResp struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name"`
	Alias       string               `json:"alias"`
	Description string               `json:"description"`
	Source      string               `json:"source"`
	URL         string               `json:"url"`
	CreatedAt   *LocalTime           `json:"created_at"`
	UpdatedAt   *LocalTime           `json:"updated_at"`
	Children    []DocLibraryTreeResp `json:"children,omitempty"`
}
