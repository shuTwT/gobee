package model

import "time"

type FlinkCreateReq struct {
	// 名称
	Name string `json:"name"`
	// 链接
	URL string `json:"url"`
	// logo
	Logo string `json:"logo"`
	// 简介
	Description string `json:"description"`
	// 快照
	Snapshot string `json:"snapshot"`
	// 邮箱
	Email string `json:"email"`
}

type FlinkUpdateReq struct {
	// ID of the ent.
	ID int `json:"id"`
	// 名称
	Name string `json:"name"`
	// 链接
	URL string `json:"url"`
	// logo
	Logo string `json:"logo"`
	// 简介
	Description string `json:"description"`
	// 快照
	Snapshot string `json:"snapshot"`
	// 邮箱
	Email string `json:"email"`
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
	Logo string `json:"logo"`
	// 简介
	Description string `json:"description"`
	// 状态
	Status int `json:"status"`
	// 快照
	Snapshot string `json:"snapshot"`
	// 邮箱
	Email string `json:"email"`
}
