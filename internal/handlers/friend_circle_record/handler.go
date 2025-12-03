package friendcirclerecord

import (
	"gobee/internal/database"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

func ListFriendCircleRecordPage(c *fiber.Ctx) error {
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

func CreateFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

func UpdateFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}

func DeleteFriendCircleRecord(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", nil))
}
