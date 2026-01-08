package model

type MemberLevelCreateReq struct {
	Name         string   `json:"name" validate:"required"`
	Description  string   `json:"description,omitempty"`
	Level        int      `json:"level" validate:"required"`
	MinPoints    int      `json:"min_points,omitempty"`
	DiscountRate int      `json:"discount_rate,omitempty"`
	Privileges   []string `json:"privileges,omitempty"`
	Icon         string   `json:"icon,omitempty"`
	SortOrder    int      `json:"sort_order,omitempty"`
}

type MemberLevelUpdateReq struct {
	Name         *string   `json:"name,omitempty"`
	Description  *string   `json:"description,omitempty"`
	Level        *int      `json:"level,omitempty"`
	MinPoints    *int      `json:"min_points,omitempty"`
	DiscountRate *int      `json:"discount_rate,omitempty"`
	Privileges   *[]string `json:"privileges,omitempty"`
	Icon         *string   `json:"icon,omitempty"`
	Active       *bool     `json:"active,omitempty"`
	SortOrder    *int      `json:"sort_order,omitempty"`
}

type MemberLevelResp struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description,omitempty"`
	Level        int      `json:"level"`
	MinPoints    int      `json:"min_points"`
	DiscountRate int      `json:"discount_rate"`
	Privileges   []string `json:"privileges,omitempty"`
	Icon         string   `json:"icon,omitempty"`
	Active       bool     `json:"active"`
	SortOrder    int      `json:"sort_order"`
}
