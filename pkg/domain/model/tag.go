package model

type TagCreateReq struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Color       *string `json:"color,omitempty"`
	SortOrder   *int    `json:"sort_order,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

type TagUpdateReq struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Color       *string `json:"color,omitempty"`
	SortOrder   *int    `json:"sort_order,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

type TagResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Color       string `json:"color"`
	SortOrder   int    `json:"sort_order"`
	Active      bool   `json:"active"`
	PostCount   int    `json:"post_count"`
}
