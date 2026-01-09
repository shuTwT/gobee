package model

import "time"

// PostCreateReq represents the request body for creating a post.
type PostCreateReq struct {
	Title                 string  `json:"title" validate:"required"`                                //文章标题
	Alias                 *string `json:"alias,omitempty"`                                          //文章别名
	Content               string  `json:"content" validate:"required"`                              //文章内容
	MdContent             *string `json:"md_content,omitempty"`                                     //md文章内容
	HtmlContent           *string `json:"html_content,omitempty"`                                   //html文章内容
	ContentType           *string `json:"content_type" validate:"required,enum=markdown,html"`      //内容类型
	Status                *string `json:"status" validate:"required,enum=draft,published,archived"` //状态
	IsAutogenSummary      bool    `json:"is_autogen_summary"`                                       //是否自动生成摘要
	IsVisible             bool    `json:"is_visible"`                                               //是否可见
	IsTipToTop            bool    `json:"is_tip_to_top"`                                            //是否置顶
	IsAllowComment        bool    `json:"is_allow_comment"`                                         //是否允许评论
	IsVisibleAfterComment bool    `json:"is_visible_after_comment"`                                 //是否评论后可见
	IsVisibleAfterPay     bool    `json:"is_visible_after_pay"`                                     //是否支付后可见
	Money                 int     `json:"money" validate:"required,min=0"`                          //支付金额
	Cover                 *string `json:"cover,omitempty"`                                          //文章封面
	Keywords              *string `json:"keywords,omitempty"`                                       //文章关键词
	Copyright             *string `json:"copyright,omitempty"`                                      //文章版权
	Author                string  `json:"author" validate:"required"`                               //作者
	Summary               *string `json:"summary,omitempty"`                                        //文章摘要
}

// PostUpdateReq represents the request body for updating a post.
type PostUpdateReq struct {
	Title            string  `json:"title,omitempty"`                                          //文章标题
	Alias            *string `json:"alias,omitempty"`                                          //文章别名
	Content          string  `json:"content"`                                                  //文章内容
	MdContent        *string `json:"md_content,omitempty"`                                     //md文章内容
	HtmlContent      *string `json:"html_content,omitempty"`                                   //html文章内容
	ContentType      *string `json:"content_type" validate:"required,enum=markdown,html"`      //内容类型
	Status           *string `json:"status" validate:"required,enum=draft,published,archived"` //状态
	IsAutogenSummary bool    `json:"is_autogen_summary,omitempty"`                             //是否自动生成摘要
	IsVisible        bool    `json:"is_visible,omitempty"`                                     //是否可见
	IsTipToTop       bool    `json:"is_tip_to_top,omitempty"`                                  //是否置顶
	IsAllowComment   bool    `json:"is_allow_comment,omitempty"`                               //是否允许评论
	Cover            string  `json:"cover,omitempty"`                                          //文章封面
	Keywords         string  `json:"keywords,omitempty"`                                       //文章关键词
	Copyright        string  `json:"copyright,omitempty"`                                      //文章版权
	Author           string  `json:"author,omitempty"`                                         //作者
	Summary          string  `json:"summary,omitempty"`                                        //文章摘要
}

// PostResp represents the response body for a post.
type PostResp struct {
	ID                    int        `json:"id"`                       //文章ID
	CreatedAt             time.Time  `json:"created_at"`               //创建时间
	Title                 string     `json:"title"`                    //文章标题
	Alias                 *string    `json:"alias,omitempty"`          //文章别名
	Content               string     `json:"content"`                  //文章内容
	MdContent             *string    `json:"md_content,omitempty"`     //md文章内容
	HtmlContent           *string    `json:"html_content,omitempty"`   //html文章内容
	ContentType           string     `json:"content_type"`             //内容类型
	Status                string     `json:"status"`                   //状态
	IsAutogenSummary      bool       `json:"is_autogen_summary"`       //是否自动生成摘要
	IsVisible             bool       `json:"is_visible"`               //是否可见
	IsTipToTop            bool       `json:"is_tip_to_top"`            //是否置顶
	IsAllowComment        bool       `json:"is_allow_comment"`         //是否允许评论
	IsVisibleAfterComment bool       `json:"is_visible_after_comment"` //是否评论后可见
	IsVisibleAfterPay     bool       `json:"is_visible_after_pay"`     //是否支付后可见
	Money                 int        `json:"money"`                    //支付金额
	PublishedAt           *time.Time `json:"published_at,omitempty"`   //发布时间
	ViewCount             int        `json:"view_count"`               //浏览次数
	CommentCount          int        `json:"comment_count"`            //评论次数
	Cover                 *string    `json:"cover,omitempty"`          //文章封面
	Keywords              *string    `json:"keywords,omitempty"`       //文章关键词
	Copyright             *string    `json:"copyright,omitempty"`      //文章版权
	Author                string     `json:"author"`                   //作者
	Summary               *string    `json:"summary,omitempty"`        //文章摘要
}
