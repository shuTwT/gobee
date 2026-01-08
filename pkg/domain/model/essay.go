package model

type EssayCreateReq struct {
	Content string   `json:"content"`
	Draft   bool     `json:"draft"`
	Images  []string `json:"images"`
}

type EssayUpdateReq struct {
	ID      int      `json:"id"`
	Content string   `json:"content"`
	Draft   bool     `json:"draft"`
	Images  []string `json:"images,omitempty"`
}

type EssayResp struct {
	ID       int        `json:"id"`
	Content  string     `json:"content"`
	Draft    bool       `json:"draft"`
	Images   []string   `json:"images"`
	CreateAt *LocalTime `json:"create_at"`
}
