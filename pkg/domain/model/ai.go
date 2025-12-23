package model

type AIResponse struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
