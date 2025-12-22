package post

import (
	"gobee/ent"
	"gobee/ent/post"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
	"strconv"

	post_service "gobee/internal/services/post"

	"github.com/gofiber/fiber/v2"
)

type PostHandler interface {
	ListPost(c *fiber.Ctx) error
	CreatePost(c *fiber.Ctx) error
	UpdatePostContent(c *fiber.Ctx) error
	UpdatePostSetting(c *fiber.Ctx) error
	PublishPost(c *fiber.Ctx) error
	UnpublishPost(c *fiber.Ctx) error
	QueryPost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
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
	client := database.DB
	posts, err := client.Post.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", posts))
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
	client := database.DB
	var post *ent.Post
	if err := c.BodyParser(&post); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newPost, err := client.Post.Create().
		SetTitle(post.Title).
		SetContent(post.Content).
		Save(c.Context())
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
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var post *model.PostUpdateReq
	if err = c.BodyParser(&post); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newPost, err := client.Post.UpdateOneID(id).
		SetContent(post.Content).
		Save(c.Context())
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
	var post *ent.Post
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
	post, err := client.Post.Query().
		Where(post.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", post))
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
