package album

import (
	"gobee/ent"
	"gobee/ent/album"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/**
 * @api {get} /album/list ListAlbum
 * @apiName ListAlbum
 * @apiGroup Album
 *
 * @apiSuccess {Object} data 相册列表
 * @apiSuccess {String} message 成功消息
 * @apiSuccess {Number} status 状态码
 *
 * @apiError {String} message 错误消息
 * @apiError {Number} status 状态码
 */
func ListAlbum(c *fiber.Ctx) error {
	client := database.DB

	albums, err := client.Album.Query().
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albums))
}

func ListAlbumPage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	count, err := client.Album.Query().Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	albums, err := client.Album.Query().
		Limit(pageQuery.Size).
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	pageResult := model.PageResult[*ent.Album]{
		Total:   int64(count),
		Records: albums,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

func CreateAlbum(c *fiber.Ctx) error {
	var album *model.AlbumCreateReq
	if err := c.BodyParser(&album); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	if err := client.Album.Create().
		SetName(album.Name).
		SetDescription(album.Description).
		SetSort(album.Sort).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", album))
}

func UpdateAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var album *model.AlbumUpdateReq
	if err := c.BodyParser(&album); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	if err := client.Album.UpdateOneID(id).
		SetName(album.Name).
		SetDescription(album.Description).
		SetSort(album.Sort).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", album))
}

func QueryAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := database.DB
	album, err := client.Album.Query().
		Where(album.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", album))
}

func DeleteAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := database.DB
	if err := client.Album.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
