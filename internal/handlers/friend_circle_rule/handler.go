package friendcirclerule

import (
	"gobee/internal/services/friendcirclerule"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FriendCircleRuleHandler interface {
	ListFriendCircleRule(c *fiber.Ctx) error
	ListFriendCircleRulePage(c *fiber.Ctx) error
	CreateFriendCircleRule(c *fiber.Ctx) error
	UpdateFriendCircleRule(c *fiber.Ctx) error
	DeleteFriendCircleRule(c *fiber.Ctx) error
}

type FriendCircleRuleHandlerImpl struct {
	friendCircleRuleService friendcirclerule.FriendCircleRuleService
}

func NewFriendCircleRuleHandlerImpl(friendCircleRuleService friendcirclerule.FriendCircleRuleService) *FriendCircleRuleHandlerImpl {
	return &FriendCircleRuleHandlerImpl{
		friendCircleRuleService: friendCircleRuleService,
	}
}

// @Summary 获取朋友圈规则列表
// @Description 获取所有朋友圈规则
// @Tags friend_circle_rules
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]model.FriendCircleRuleResp}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_rules [get]
func (h *FriendCircleRuleHandlerImpl) ListFriendCircleRule(c *fiber.Ctx) error {
	rules, err := h.friendCircleRuleService.ListFriendCircleRule(c.Context())
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

// @Summary 获取朋友圈规则分页列表
// @Description 获取朋友圈规则分页列表
// @Tags friend_circle_rules
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Success 200 {object} model.HttpSuccess{data=model.PageResult[model.FriendCircleRuleResp]}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_rules/page [get]
func (h *FriendCircleRuleHandlerImpl) ListFriendCircleRulePage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	count, rules, err := h.friendCircleRuleService.ListFriendCircleRulePage(c.Context(), pageQuery.Page, pageQuery.Size)
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

// @Summary 创建朋友圈规则
// @Description 创建新的朋友圈规则
// @Tags friend_circle_rules
// @Accept json
// @Produce json
// @Param friend_circle_rule_save_req body model.FriendCircleRuleSaveReq true "朋友圈规则保存请求体"
// @Success 200 {object} model.HttpSuccess{data=ent.FriendCircleRule}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_rules [post]
func (h *FriendCircleRuleHandlerImpl) CreateFriendCircleRule(c *fiber.Ctx) error {
	var createReq *model.FriendCircleRuleSaveReq
	if err := c.BodyParser(&createReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	rule, err := h.friendCircleRuleService.CreateFriendCircleRule(c.Context(), createReq.Name, createReq.TitleSelector, createReq.LinkSelector, createReq.CreatedSelector, createReq.UpdatedSelector)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", rule))
}

// @Summary 更新朋友圈规则
// @Description 更新指定 ID 的朋友圈规则
// @Tags friend_circle_rules
// @Accept json
// @Produce json
// @Param id path int true "朋友圈规则 ID"
// @Param friend_circle_rule_save_req body model.FriendCircleRuleSaveReq true "朋友圈规则保存请求体"
// @Success 200 {object} model.HttpSuccess{data=ent.FriendCircleRule}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_rules/{id} [put]
func (h *FriendCircleRuleHandlerImpl) UpdateFriendCircleRule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid ID format"))
	}

	var updateReq *model.FriendCircleRuleSaveReq
	if err = c.BodyParser(&updateReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	rule, err := h.friendCircleRuleService.UpdateFriendCircleRule(c.Context(), id, updateReq.Name, updateReq.TitleSelector, updateReq.LinkSelector, updateReq.CreatedSelector, updateReq.UpdatedSelector)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", rule))
}

// @Summary 删除朋友圈规则
// @Description 删除指定 ID 的朋友圈规则
// @Tags friend_circle_rules
// @Accept json
// @Produce json
// @Param id path int true "朋友圈规则 ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/friend_circle_rules/{id} [delete]
func (h *FriendCircleRuleHandlerImpl) DeleteFriendCircleRule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid ID format"))
	}

	err = h.friendCircleRuleService.DeleteFriendCircleRule(c.Context(), id)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
