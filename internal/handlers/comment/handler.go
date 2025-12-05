package comment

import (
	commentService "gobee/internal/services/comment"
	"gobee/pkg/domain/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

var TWIKOO_EVENT = struct {
	GetConfig        string
	CommentGET       string
	GetCommentsCount string
	CommentSubmit    string
}{
	GetConfig:        "GET_CONFIG",
	CommentGET:       "COMMENT_GET",
	GetCommentsCount: "GET_COMMENTS_COUNT",
	CommentSubmit:    "COMMENT_SUBMIT",
}

func ListCommentPage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	resp, err := commentService.ListCommentPage(c.Context(), pageQuery)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("评论列表获取成功", resp))
}

type TwikooReqBody struct {
	AccessToken  *string   `json:"accessToken"`
	Event        string    `json:"event"`
	Url          *string   `json:"url"`
	EnvId        *string   `json:"envId"`
	IncludeReply *bool     `json:"includeReply"`
	Region       *string   `json:"region"`
	Urls         *[]string `json:"urls"`
	Comment      *string   `json:"comment"`
	Href         *string   `json:"href"`
	Link         *string   `json:"link"`
	Mail         *string   `json:"mail"`
	Nick         *string   `json:"nick"`
	UA           *string   `json:"ua"`
}

func HandleTwikoo(c *fiber.Ctx) error {

	var reqBody TwikooReqBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	// 配置
	if reqBody.Event == TWIKOO_EVENT.GetConfig {
		return c.JSON(fiber.Map{
			"code":   0,
			"config": fiber.Map{},
		})
	}
	// 获取评论
	if reqBody.Event == TWIKOO_EVENT.CommentGET {
		commmentList := []fiber.Map{}
		comments, err := commentService.ListComment(c.Context(), *reqBody.Url)
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		for _, comment := range comments {
			browser, os := commentService.ParseUserAgent(*comment.UserAgent)
			commmentList = append(commmentList, fiber.Map{
				"id":        comment.ID,
				"parentId":  comment.ParentID,
				"url":       comment.URL,
				"userAgent": comment.UserAgent,
				"comment":   comment.Content,
				"browser":   browser,
				"os":        os,
				"created":   time.Time(comment.CreatedAt).UnixMilli(),
				"updated":   time.Time(comment.UpdatedAt).UnixMilli(),
				"replies":   []fiber.Map{},
			})
		}
		return c.JSON(fiber.Map{
			"count": len(commmentList),
			"data":  commmentList,
			"more":  false,
		})
	}
	// 获取评论数量
	if reqBody.Event == TWIKOO_EVENT.GetCommentsCount {
		data := []fiber.Map{}
		count, err := commentService.CountComment(c.Context(), *reqBody.IncludeReply, *reqBody.Urls)
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		for _, url := range *reqBody.Urls {
			data = append(data, fiber.Map{
				"url":   url,
				"count": count,
			})
		}
		return c.JSON(fiber.Map{
			"data": data,
		})
	}

	// 提交评论
	if reqBody.Event == TWIKOO_EVENT.CommentSubmit {
		ipAddress := c.IP()
		id, err := commentService.CreateComment(c.Context(), *reqBody.Comment, *reqBody.Href, *reqBody.Link, *reqBody.Mail, *reqBody.Nick, *reqBody.UA, *reqBody.Url, ipAddress)
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		return c.JSON(fiber.Map{
			"id": id,
		})
	}
	return c.JSON(fiber.Map{})
}

func RecentComment(c *fiber.Ctx) error {
	comments, err := commentService.GetRecentComment(c.Context(), 10)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	commentList := []fiber.Map{}
	for _, entity := range comments {
		commentList = append(commentList, fiber.Map{
			"id":        entity.ID,
			"parentId":  entity.ParentID,
			"url":       entity.URL,
			"userAgent": entity.UserAgent,
			"comment":   entity.Content,
			"created":   time.Time(entity.CreatedAt).UnixMilli(),
			"updated":   time.Time(entity.UpdatedAt).UnixMilli(),
			"replies":   []fiber.Map{},
		})
	}
	return c.JSON(model.NewSuccess("最近评论获取成功", commentList))
}
