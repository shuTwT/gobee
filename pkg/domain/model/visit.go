package model

type VisitLogReq struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Path      string `json:"path"`
	OS        string `json:"os"`
	Browser   string `json:"browser"`
	Device    string `json:"device"`
}

type VisitLogPageQuery struct {
	Page  int    `json:"page" query:"page" form:"page" validate:"required,min=1"`
	Size  int    `json:"page_size" query:"page_size" form:"page_size" validate:"required,min=1,max=100"`
	IP    string `json:"ip" query:"ip" form:"ip"`
	Path  string `json:"path" query:"path" form:"path"`
}

type VisitLogBatchDeleteReq struct {
	IDs []int `json:"ids" validate:"required,min=1"`
}
