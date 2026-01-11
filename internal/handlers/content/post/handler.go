package post

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gobee/ent"
	"gobee/ent/post"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
	"strconv"
	"time"

	post_service "gobee/internal/services/content/post"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type PostHandler interface {
	ListPost(c *fiber.Ctx) error
	ListPostPage(c *fiber.Ctx) error
	CreatePost(c *fiber.Ctx) error
	UpdatePostContent(c *fiber.Ctx) error
	UpdatePostSetting(c *fiber.Ctx) error
	PublishPost(c *fiber.Ctx) error
	UnpublishPost(c *fiber.Ctx) error
	QueryPost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
	GetSummaryForStream(c *fiber.Ctx) error
}

type PostHandlerImpl struct {
	postService post_service.PostService
}

func NewPostHandlerImpl(postService post_service.PostService) *PostHandlerImpl {
	return &PostHandlerImpl{
		postService: postService,
	}
}

// @Summary 查询所有文章
// @Description 查询所有文章
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]ent.Post}
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts [get]
func (h *PostHandlerImpl) ListPost(c *fiber.Ctx) error {
	posts, err := h.postService.QueryPostList(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", posts))
}

func (h *PostHandlerImpl) ListPostPage(c *fiber.Ctx) error {
	var req model.PostPageReq
	if err := c.QueryParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	posts, count, err := h.postService.QueryPostPage(c.Context(), req)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", model.PageResult[*ent.Post]{
		Total:   int64(count),
		Records: posts,
	}))
}

// @Summary 创建文章
// @Description 创建一篇新文章
// @Tags posts
// @Accept json
// @Produce json
// @Param post body model.PostCreateReq true "文章创建请求"
// @Success 200 {object} model.HttpSuccess{data=ent.Post}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts [post]
func (h *PostHandlerImpl) CreatePost(c *fiber.Ctx) error {
	var post model.PostCreateReq
	if err := c.BodyParser(&post); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newPost, err := h.postService.CreatePost(c.Context(), post.Title, post.Content)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newPost))
}

// @Summary 更新文章内容
// @Description 更新指定文章的内容
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Param post body model.PostUpdateReq true "文章更新请求"
// @Success 200 {object} model.HttpSuccess{data=ent.Post}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts/{id} [put]
func (h *PostHandlerImpl) UpdatePostContent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var post model.PostUpdateReq
	if err = c.BodyParser(&post); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newPost, err := h.postService.UpdatePostContent(c.Context(), id, post.Content)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newPost))
}

// @Summary 更新文章设置
// @Description 更新指定文章的设置
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Param post body model.PostUpdateReq true "文章更新请求"
// @Success 200 {object} model.HttpSuccess{data=ent.Post}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts/{id}/settings [put]
func (h *PostHandlerImpl) UpdatePostSetting(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var post model.PostUpdateReq
	if err = c.BodyParser(&post); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newPost, err := h.postService.UpdatePostSetting(c.Context(), id, post)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newPost))
}

// @Summary 发布文章
// @Description 发布指定文章
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.HttpSuccess{data=ent.Post}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts/{id}/publish [put]
func (h *PostHandlerImpl) PublishPost(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	newPost, err := client.Post.UpdateOneID(id).
		SetStatus("published").
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newPost))
}

// @Summary 取消发布文章
// @Description 取消发布指定文章
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.HttpSuccess{data=ent.Post}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts/{id}/unpublish [put]
func (h *PostHandlerImpl) UnpublishPost(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	newPost, err := client.Post.UpdateOneID(id).
		SetStatus("draft").
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newPost))
}

// @Summary 查询文章
// @Description 查询指定文章
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.HttpSuccess{data=ent.Post}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts/{id} [get]
func (h *PostHandlerImpl) QueryPost(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	post, err := client.Post.Query().WithCategories().WithTags().
		Where(post.ID(id)).
		First(c.Context())
	postResp := model.PostResp{
		ID:                    post.ID,
		Title:                 post.Title,
		Alias:                 post.Alias,
		Content:               post.Content,
		MdContent:             post.MdContent,
		HtmlContent:           post.HTMLContent,
		ContentType:           string(post.ContentType),
		Status:                string(post.Status),
		IsAutogenSummary:      post.IsAutogenSummary,
		IsVisible:             post.IsVisible,
		IsPinToTop:            post.IsPinToTop,
		IsAllowComment:        post.IsAllowComment,
		IsVisibleAfterComment: post.IsVisibleAfterComment,
		IsVisibleAfterPay:     post.IsVisibleAfterPay,
		Price:                 float32(post.Price) / 100,
		PublishedAt:           post.PublishedAt,
		ViewCount:             post.ViewCount,
		CommentCount:          post.CommentCount,
		Cover:                 &post.Cover,
		Keywords:              &post.Keywords,
		Copyright:             &post.Copyright,
		Author:                post.Author,
		Summary:               &post.Summary,
		Categories: func() []int {
			ids := make([]int, len(post.Edges.Categories))
			for i, cat := range post.Edges.Categories {
				ids[i] = cat.ID
			}
			return ids
		}(),
		Tags: func() []int {
			ids := make([]int, len(post.Edges.Tags))
			for i, tag := range post.Edges.Tags {
				ids[i] = tag.ID
			}
			return ids
		}(),
	}
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", postResp))
}

// @Summary 删除文章
// @Description 删除指定文章
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts/{id} [delete]
func (h *PostHandlerImpl) DeletePost(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	if err := client.Post.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}

// simulateAIProcessing 模拟AI处理过程
func simulateAIProcessing(targetText string, ch chan model.AIResponse) {
	defer close(ch)

	for i, text := range targetText {
		ch <- model.AIResponse{
			Content: string(text),
			Done:    i == len(targetText)-1,
		}
		time.Sleep(50 * time.Millisecond)
	}
}

// @Summary 获取文章摘要流
// @Description 获取指定文章的摘要流
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.HttpSuccess{data=model.AIResponse}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/posts/{id}/summary [get]
func (h *PostHandlerImpl) GetSummaryForStream(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	responseChan := make(chan model.AIResponse)

	go simulateAIProcessing("你好!我是AI助手,我正在处理你的请求,这是一个流式响应示例,马上就要完成了,处理完成!", responseChan)
	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for response := range responseChan {
			data, err := json.Marshal(response.Content)
			if err != nil {
				continue
			}

			fmt.Fprintf(w, "%s", string(data))

			err = w.Flush()

			if err != nil {
				break
			}

			if response.Done {
				break
			}
		}

	}))

	return nil
}
