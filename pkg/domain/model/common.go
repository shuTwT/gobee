package model

type HomeStatistic struct {
	PostCount    int `json:"postCount"`
	UserCount    int `json:"userCount"`
	CommentCount int `json:"commentCount"`
	VisitCount   int `json:"visitCount"`
}

// SiteStatistic 站点公开统计信息，仅统计已发布且可见的文章
type SiteStatistic struct {
	PostCount      int `json:"postCount"`
	CategoryCount  int `json:"categoryCount"`
	TagCount       int `json:"tagCount"`
	TotalWordCount int `json:"totalWordCount"`
}
