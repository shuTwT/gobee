package albumphoto

import (
	"gobee/ent"
	"gobee/ent/albumphoto"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListAlbumPhoto(c *fiber.Ctx) error {
	client := database.DB
	albumPhotos, err := client.AlbumPhoto.Query().
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhotos))
}

func CreateAlbumPhoto(c *fiber.Ctx) error {
	var albumPhoto *ent.AlbumPhoto
	if err := c.BodyParser(&albumPhoto); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	if err := client.AlbumPhoto.Create().
		SetAlbumID(albumPhoto.AlbumID).
		SetImageURL(albumPhoto.ImageURL).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhoto))
}
func UpdateAlbumPhoto(c *fiber.Ctx) error {
	var albumPhoto *ent.AlbumPhoto
	if err := c.BodyParser(&albumPhoto); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	if err := client.AlbumPhoto.UpdateOneID(albumPhoto.ID).
		SetImageURL(albumPhoto.ImageURL).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhoto))
}
func QueryAlbumPhoto(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := database.DB
	albumPhoto, err := client.AlbumPhoto.Query().
		Where(albumphoto.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albumPhoto))
}

func DeleteAlbumPhoto(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := database.DB
	if err := client.AlbumPhoto.DeleteOneID(id).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
