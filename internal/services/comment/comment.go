package comment

import (
	"gobee/ent"
	"gobee/internal/database"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

func ListCommentPage(c *fiber.Ctx) error {
	client := database.DB
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	count, err := client.Comment.Query().Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	comments, err := client.Comment.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	pageResult := model.PageResult[*ent.Comment]{
		Total:   int64(count),
		Records: comments,
	}
	return c.JSON(model.NewSuccess("评论列表获取成功", pageResult))
}
