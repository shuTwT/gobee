package model

type CouponCreateReq struct {
	Name          string   `json:"name" validate:"required"`
	Code          string   `json:"code" validate:"required"`
	Description   *string  `json:"description,omitempty"`
	Type          int      `json:"type" validate:"required"`
	Value         int      `json:"value" validate:"required"`
	MinAmount     int      `json:"min_amount"`
	MaxDiscount   int      `json:"max_discount"`
	TotalCount    int      `json:"total_count" validate:"required"`
	PerUserLimit  int      `json:"per_user_limit"`
	StartTime     string   `json:"start_time" validate:"required"`
	EndTime       string   `json:"end_time" validate:"required"`
	Active        *bool    `json:"active,omitempty"`
	Image         *string  `json:"image,omitempty"`
	ProductIds    []int    `json:"product_ids,omitempty"`
	CategoryIds   []int    `json:"category_ids,omitempty"`
}

type CouponUpdateReq struct {
	Name         *string  `json:"name,omitempty"`
	Description  *string  `json:"description,omitempty"`
	Type         *int     `json:"type,omitempty"`
	Value        *int     `json:"value,omitempty"`
	MinAmount    *int     `json:"min_amount,omitempty"`
	MaxDiscount  *int     `json:"max_discount,omitempty"`
	TotalCount   *int     `json:"total_count,omitempty"`
	PerUserLimit *int     `json:"per_user_limit,omitempty"`
	StartTime    *string  `json:"start_time,omitempty"`
	EndTime      *string  `json:"end_time,omitempty"`
	Active       *bool    `json:"active,omitempty"`
	Image        *string  `json:"image,omitempty"`
	ProductIds   []int    `json:"product_ids,omitempty"`
	CategoryIds  []int    `json:"category_ids,omitempty"`
}

type CouponBatchUpdateReq struct {
	IDs    []int  `json:"ids" validate:"required"`
	Active *bool  `json:"active,omitempty"`
}

type CouponBatchDeleteReq struct {
	IDs []int `json:"ids" validate:"required"`
}

type CouponSearchReq struct {
	Keyword string `json:"keyword" query:"keyword" form:"keyword"`
	Type    *int   `json:"type" query:"type" form:"type"`
	Active  *bool  `json:"active" query:"active" form:"active"`
	Page    int    `json:"page" query:"page" form:"page" validate:"required,min=1"`
	Size    int    `json:"page_size" query:"page_size" form:"page_size" validate:"required,min=1,max=100"`
}
