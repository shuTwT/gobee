package friendcirclerule

import (
	"gobee/internal/database"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListFriendCircleRule(c *fiber.Ctx) error {
	client := database.DB
	rules, err := client.FriendCircleRule.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	records := []model.FriendCircleRuleResp{}
	for _, rule := range rules {
		records = append(records, model.FriendCircleRuleResp{
			ID:              rule.ID,
			Name:            rule.Name,
			TitleSelector:   &rule.TitleSelector,
			LinkSelector:    &rule.LinkSelector,
			CreatedSelector: &rule.CreatedSelector,
			UpdatedSelector: &rule.UpdatedSelector,
		})
	}
	return c.JSON(model.NewSuccess("success", records))
}

func ListFriendCircleRulePage(c *fiber.Ctx) error {
	client := database.DB
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	count, err := client.FriendCircleRule.Query().Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	rules, err := client.FriendCircleRule.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	records := []model.FriendCircleRuleResp{}
	for _, rule := range rules {
		records = append(records, model.FriendCircleRuleResp{
			ID:              rule.ID,
			Name:            rule.Name,
			TitleSelector:   &rule.TitleSelector,
			LinkSelector:    &rule.LinkSelector,
			CreatedSelector: &rule.CreatedSelector,
			UpdatedSelector: &rule.UpdatedSelector,
		})
	}
	pageResult := model.PageResult[model.FriendCircleRuleResp]{
		Total:   int64(count),
		Records: records,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

func CreateFriendCircleRule(c *fiber.Ctx) error {
	var createReq *model.FriendCircleRuleSaveReq
	if err := c.BodyParser(&createReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	rule, err := client.FriendCircleRule.Create().
		SetName(createReq.Name).
		SetNillableTitleSelector(&createReq.TitleSelector).
		SetNillableLinkSelector(&createReq.LinkSelector).
		SetNillableCreatedSelector(&createReq.CreatedSelector).
		SetNillableUpdatedSelector(&createReq.UpdatedSelector).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", rule))
}

func UpdateFriendCircleRule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid ID format"))
	}
	var updateReq *model.FriendCircleRuleSaveReq
	if err = c.BodyParser(&updateReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	rule, err := client.FriendCircleRule.UpdateOneID(id).
		SetName(updateReq.Name).
		SetNillableTitleSelector(&updateReq.TitleSelector).
		SetNillableLinkSelector(&updateReq.LinkSelector).
		SetNillableCreatedSelector(&updateReq.CreatedSelector).
		SetNillableUpdatedSelector(&updateReq.UpdatedSelector).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", rule))
}

func DeleteFriendCircleRule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := database.DB
	if err := client.FriendCircleRule.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
