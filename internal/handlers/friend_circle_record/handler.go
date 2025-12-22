package friendcirclerecord

import (
	"gobee/ent"
	"gobee/internal/database"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

type FriendCircleRecordHandler interface {
	ListFriendCircleRecordPage(c *fiber.Ctx) error
	CreateFriendCircleRecord(c *fiber.Ctx) error
	UpdateFriendCircleRecord(c *fiber.Ctx) error
	DeleteFriendCircleRecord(c *fiber.Ctx) error
}

type FriendCircleRecordHandlerImpl struct {
	client *ent.Client
}

func NewFriendCircleRecordHandlerImpl(client *ent.Client) *FriendCircleRecordHandlerImpl {
	return &FriendCircleRecordHandlerImpl{client: client}
}

// @Summary 获取朋友圈记录分页列表
// @Description 获取朋友圈记录分页列表
// @Tags friend_circle_records
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Success 200 {object} model.HttpSuccess{data=model.PageResult[model.FriendCircleRecordResp]}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_records [get]
func (h *FriendCircleRecordHandlerImpl) ListFriendCircleRecordPage(c *fiber.Ctx) error {
	client := database.DB
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	count, err := client.FriendCircleRecord.Query().Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	records, err := client.FriendCircleRecord.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	resps := []model.FriendCircleRecordResp{}
	for _, record := range records {
		resps = append(resps, model.FriendCircleRecordResp{
			ID:          record.ID,
			Title:       record.Title,
			Author:      record.Author,
			LinkURL:     record.LinkURL,
			AvatarURL:   record.AvatarURL,
			PublishedAt: record.PublishedAt,
		})
	}
	pageResult := model.PageResult[model.FriendCircleRecordResp]{
		Total:   int64(count),
		Records: resps,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

// @Summary 创建朋友圈记录
// @Description 创建朋友圈记录
// @Tags friend_circle_records
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=ent.FriendCircleRecord}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_records [post]
func (h *FriendCircleRecordHandlerImpl) CreateFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 更新朋友圈记录
// @Description 更新朋友圈记录
// @Tags friend_circle_records
// @Accept json
// @Produce json
// @Param id path int true "朋友圈记录 ID"
// @Success 200 {object} model.HttpSuccess{data=ent.FriendCircleRecord}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_records/{id} [put]
func (h *FriendCircleRecordHandlerImpl) UpdateFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 删除朋友圈记录
// @Description 删除朋友圈记录
// @Tags friend_circle_records
// @Accept json
// @Produce json
// @Param id path int true "朋友圈记录 ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_records/{id} [delete]
func (h *FriendCircleRecordHandlerImpl) DeleteFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}
