package albumphoto

import (
	"gobee/ent"
	"gobee/ent/albumphoto"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlbumPhotoHandler interface {
	ListAlbumPhoto(c *fiber.Ctx) error
	ListAlbumPhotoPage(c *fiber.Ctx) error
	CreateAlbumPhoto(c *fiber.Ctx) error
	UpdateAlbumPhoto(c *fiber.Ctx) error
	QueryAlbumPhoto(c *fiber.Ctx) error
	DeleteAlbumPhoto(c *fiber.Ctx) error
}

type AlbumPhotoHandlerImpl struct {
	client *ent.Client
}

func NewAlbumPhotoHandlerImpl(client *ent.Client) *AlbumPhotoHandlerImpl {
	return &AlbumPhotoHandlerImpl{
		client: client,
	}
}

func (h *AlbumPhotoHandlerImpl) ListAlbumPhoto(c *fiber.Ctx) error {
	albumPhotos, err := h.client.AlbumPhoto.Query().
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhotos))
}

func (h *AlbumPhotoHandlerImpl) ListAlbumPhotoPage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	count, err := h.client.AlbumPhoto.Query().
		Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	albumPhotos, err := h.client.AlbumPhoto.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	pageResult := model.PageResult[*ent.AlbumPhoto]{
		Total:   int64(count),
		Records: albumPhotos,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

func (h *AlbumPhotoHandlerImpl) CreateAlbumPhoto(c *fiber.Ctx) error {
	var albumPhoto *model.AlbumPhotoCreateReq
	if err := c.BodyParser(&albumPhoto); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	if err := h.client.AlbumPhoto.Create().
		SetAlbumID(albumPhoto.AlbumID).
		SetName(albumPhoto.Name).
		SetDescription(albumPhoto.Description).
		SetImageURL(albumPhoto.ImageURL).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhoto))
}
func (h *AlbumPhotoHandlerImpl) UpdateAlbumPhoto(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var albumPhoto *model.AlbumPhotoUpdateReq
	if err := c.BodyParser(&albumPhoto); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	if err := h.client.AlbumPhoto.UpdateOneID(id).
		SetImageURL(albumPhoto.ImageURL).
		SetName(albumPhoto.Name).
		SetDescription(albumPhoto.Description).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhoto))
}
func (h *AlbumPhotoHandlerImpl) QueryAlbumPhoto(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	albumPhoto, err := h.client.AlbumPhoto.Query().
		Where(albumphoto.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhoto))
}

func (h *AlbumPhotoHandlerImpl) DeleteAlbumPhoto(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	if err := h.client.AlbumPhoto.DeleteOneID(id).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
