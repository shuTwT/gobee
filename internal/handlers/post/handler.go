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

func (h *PostHandlerImpl) ListPost(c *fiber.Ctx) error {
	client := database.DB
	posts, err := client.Post.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", posts))
}

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
