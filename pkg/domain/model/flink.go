package model

import "time"

type FlinkCreateReq struct {
	// 名称
	Name string `json:"name"`
	// 链接
	URL string `json:"url"`
	// logo
	AvatarURL string `json:"avatar_url"`
	// 简介
	Description string `json:"description"`
	CoverURL    string `json:"cover_url"`
	// 快照
	SnapshotURL string `json:"snapshot_url"`
	// 邮箱
	Email string `json:"email"`
	// 是否开启朋友圈
	EnableFriendCircle bool `json:"enable_friend_circle"`
	// 朋友圈解析规则
	FriendCircleRuleID *int `json:"friend_circle_rule_id"`
}

type FlinkUpdateReq struct {
	// ID of the ent.
	ID int `json:"id"`
	// 名称
	Name string `json:"name"`
	// 链接
	URL string `json:"url"`
	// logo
	AvatarURL string `json:"avatar_url"`
	CoverURL  string `json:"cover_url"`
	// 简介
	Description string `json:"description"`
	// 快照
	SnapshotURL string `json:"snapshot_url"`
	// 邮箱
	Email string `json:"email"`
	// 是否开启朋友圈
	EnableFriendCircle bool `json:"enable_friend_circle"`
	// 朋友圈解析规则
	FriendCircleRuleID int `json:"friend_circle_rule_id"`
}

type FlinkResp struct {
	// ID of the ent.
	ID int `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at"`
	// 名称
	Name string `json:"name"`
	// 链接
	URL string `json:"url"`
	// logo
	AvatarURL string `json:"avatar_url"`
	// 简介
	Description string `json:"description"`
	// 状态
	Status   int    `json:"status"`
	CoverURL string `json:"cover_url"`
	// 快照
	SnapshotUrl string `json:"snapshot_url"`
	// 邮箱
	Email string `json:"email"`
	// 是否开启朋友圈
	EnableFriendCircle bool `json:"enable_friend_circle"`
	// 朋友圈解析规则
	FriendCircleRuleID *int `json:"friend_circle_rule_id"`
}
