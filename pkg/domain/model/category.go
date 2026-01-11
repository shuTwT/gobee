package model

type CategoryCreateReq struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	SortOrder   int     `json:"sort_order" validate:"required"`
	Active      *bool   `json:"active,omitempty"`
}

type CategoryUpdateReq struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	SortOrder   *int    `json:"sort_order,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

type CategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	SortOrder   int    `json:"sort_order"`
	Active      bool   `json:"active"`
	PostCount   int    `json:"post_count"`
}
