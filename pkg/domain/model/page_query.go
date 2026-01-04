package model

type PageQuery struct {
	Page int `json:"page" query:"page" form:"page" validate:"required,min=1"`
	Size int `json:"page_size" query:"page_size" form:"page_size" validate:"required,min=1,max=100"`
}

type PageResult[T any] struct {
	Total   int64 `json:"total"`
	Records []T   `json:"records"`
}
