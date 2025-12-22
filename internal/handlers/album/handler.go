package album

import (
	"gobee/ent"
	"gobee/ent/album"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlbumHandler interface {
	ListAlbum(c *fiber.Ctx) error
	ListAlbumPage(c *fiber.Ctx) error
	CreateAlbum(c *fiber.Ctx) error
	UpdateAlbum(c *fiber.Ctx) error
	QueryAlbum(c *fiber.Ctx) error
	DeleteAlbum(c *fiber.Ctx) error
}

type AlbumHandlerImpl struct {
	client *ent.Client
}

func NewAlbumHandlerImpl(client *ent.Client) *AlbumHandlerImpl {
	return &AlbumHandlerImpl{
		client: client,
	}
}

// @Summary 查询相册列表
// @Description 查询所有相册
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]ent.Album}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/albums [get]
func (h *AlbumHandlerImpl) ListAlbum(c *fiber.Ctx) error {
	client := h.client

	albums, err := client.Album.Query().
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", albums))
}

// @Summary 查询相册列表分页
// @Description 查询相册列表分页
// @Tags albums
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Success 200 {object} model.HttpSuccess{data=model.PageResult[ent.Album]}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/albums/page [get]
func (h *AlbumHandlerImpl) ListAlbumPage(c *fiber.Ctx) error {
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

// @Summary 创建相册
// @Description 创建一个新相册
// @Tags albums
// @Accept json
// @Produce json
// @Param req body model.AlbumCreateReq true "相册创建请求"
// @Success 200 {object} model.HttpSuccess{data=model.AlbumCreateReq}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/albums [post]
func (h *AlbumHandlerImpl) CreateAlbum(c *fiber.Ctx) error {
	var album *model.AlbumCreateReq
	if err := c.BodyParser(&album); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := h.client
	if err := client.Album.Create().
		SetName(album.Name).
		SetDescription(album.Description).
		SetSort(album.Sort).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", album))
}

// @Summary 更新相册
// @Description 更新指定相册的信息
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "相册ID"
// @Param req body model.AlbumUpdateReq true "相册更新请求"
// @Success 200 {object} model.HttpSuccess{data=model.AlbumUpdateReq}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/albums/{id} [put]
func (h *AlbumHandlerImpl) UpdateAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var album *model.AlbumUpdateReq
	if err := c.BodyParser(&album); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := h.client
	if err := client.Album.UpdateOneID(id).
		SetName(album.Name).
		SetDescription(album.Description).
		SetSort(album.Sort).
		Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", album))
}

// @Summary 查询相册
// @Description 查询指定相册的信息
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "相册ID"
// @Success 200 {object} model.HttpSuccess{data=ent.Album}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/albums/{id} [get]
func (h *AlbumHandlerImpl) QueryAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := h.client
	album, err := client.Album.Query().
		Where(album.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", album))
}

// @Summary 删除相册
// @Description 删除指定相册
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "相册ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/albums/{id} [delete]
func (h *AlbumHandlerImpl) DeleteAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := h.client
	if err := client.Album.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
