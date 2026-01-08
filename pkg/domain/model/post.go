package model

import "time"

// PostCreateReq represents the request body for creating a post.
type PostCreateReq struct {
	Title            string  `json:"title" validate:"required"`
	Alias            *string `json:"alias,omitempty"`
	Content          string  `json:"content" validate:"required"`
	IsPublished      bool    `json:"is_published"`
	IsAutogenSummary bool    `json:"is_autogen_summary"`
	IsVisible        bool    `json:"is_visible"`
	IsTipToTop       bool    `json:"is_tip_to_top"`
	IsAllowComment   bool    `json:"is_allow_comment"`
	Cover            *string `json:"cover,omitempty"`
	Keywords         *string `json:"keywords,omitempty"`
	Copyright        *string `json:"copyright,omitempty"`
	Author           string  `json:"author" validate:"required"`
	Summary          *string `json:"summary,omitempty"`
}

// PostUpdateReq represents the request body for updating a post.
type PostUpdateReq struct {
	Title            string `json:"title,omitempty"`
	Alias            string `json:"alias,omitempty"`
	Content          string `json:"content,omitempty"`
	IsPublished      bool   `json:"is_published,omitempty"`
	IsAutogenSummary bool   `json:"is_autogen_summary,omitempty"`
	IsVisible        bool   `json:"is_visible,omitempty"`
	IsTipToTop       bool   `json:"is_tip_to_top,omitempty"`
	IsAllowComment   bool   `json:"is_allow_comment,omitempty"`
	Cover            string `json:"cover,omitempty"`
	Keywords         string `json:"keywords,omitempty"`
	Copyright        string `json:"copyright,omitempty"`
	Author           string `json:"author,omitempty"`
	Summary          string `json:"summary,omitempty"`
}

// PostResp represents the response body for a post.
type PostResp struct {
	ID               int        `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	Title            string     `json:"title"`
	Alias            *string    `json:"alias,omitempty"`
	Content          string     `json:"content"`
	IsPublished      bool       `json:"is_published"`
	IsAutogenSummary bool       `json:"is_autogen_summary"`
	IsVisible        bool       `json:"is_visible"`
	IsTipToTop       bool       `json:"is_tip_to_top"`
	IsAllowComment   bool       `json:"is_allow_comment"`
	PublishedAt      *time.Time `json:"published_at,omitempty"`
	ViewCount        int        `json:"view_count"`
	CommentCount     int        `json:"comment_count"`
	Cover            *string    `json:"cover,omitempty"`
	Keywords         *string    `json:"keywords,omitempty"`
	Copyright        *string    `json:"copyright,omitempty"`
	Author           string     `json:"author"`
	Summary          *string    `json:"summary,omitempty"`
}
