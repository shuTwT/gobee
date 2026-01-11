package model

type NotificationPageQuery struct {
	Page  int  `json:"page" query:"page" form:"page" validate:"required,min=1"`
	Size  int  `json:"page_size" query:"page_size" form:"page_size" validate:"required,min=1,max=100"`
	IsRead *bool `json:"is_read" query:"is_read" form:"is_read"`
}

type NotificationBatchReadReq struct {
	IDs []int `json:"ids" validate:"required,min=1"`
}
