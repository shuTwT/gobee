package model

type HomeStatistic struct {
	PostCount    int `json:"postCount"`
	UserCount    int `json:"userCount"`
	CommentCount int `json:"commentCount"`
	VisitCount   int `json:"visitCount"`
}
