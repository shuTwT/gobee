package model

type CategoryCreateReq struct {
	Name        string  `json:"name" validate:"required"`
	Description  *string `json:"description,omitempty"`
	Slug         *string `json:"slug,omitempty"`
	SortOrder    int     `json:"sort_order" validate:"required"`
	Active       *bool   `json:"active,omitempty"`
}

type CategoryUpdateReq struct {
	Name        *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	Slug         *string `json:"slug,omitempty"`
	SortOrder    *int    `json:"sort_order,omitempty"`
	Active       *bool   `json:"active,omitempty"`
}
