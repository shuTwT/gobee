package flink

import (
	"gobee/ent"
	"gobee/ent/flink"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FlinkHandler interface {
	ListFlink(c *fiber.Ctx) error
	ListFlinkPage(c *fiber.Ctx) error
	CreateFlink(c *fiber.Ctx) error
	UpdateFlink(c *fiber.Ctx) error
	QueryFlink(c *fiber.Ctx) error
	DeleteFlink(c *fiber.Ctx) error
}

type FlinkHandlerImpl struct {
	client *ent.Client
}

func NewFlinkHandlerImpl(client *ent.Client) *FlinkHandlerImpl {
	return &FlinkHandlerImpl{
		client: client,
	}
}

// @Summary 获取所有Flink
// @Description 获取所有Flink
// @Tags flinks
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]model.FlinkResp}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flinks [get]
func (h *FlinkHandlerImpl) ListFlink(c *fiber.Ctx) error {
	flinks, err := h.client.FLink.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	result := []model.FlinkResp{}
	for _, flink := range flinks {
		result = append(result, model.FlinkResp{
			ID:                 flink.ID,
			CreatedAt:          &flink.CreatedAt,
			UpdatedAt:          &flink.UpdatedAt,
			Name:               flink.Name,
			URL:                flink.URL,
			AvatarURL:          flink.AvatarURL,
			Description:        flink.Description,
			CoverURL:           flink.CoverURL,
			Status:             flink.Status,
			SnapshotUrl:        flink.SnapshotURL,
			Email:              flink.Email,
			EnableFriendCircle: flink.EnableFriendCircle,
			FriendCircleRuleID: flink.FriendCircleRuleID,
		})
	}
	return c.JSON(model.NewSuccess("success", result))
}

// @Summary 获取Flink分页列表
// @Description 获取Flink分页列表
// @Tags flinks
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Success 200 {object} model.HttpSuccess{data=model.PageResult[model.FlinkResp]}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flinks/page [get]
func (h *FlinkHandlerImpl) ListFlinkPage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	count, err := h.client.FLink.Query().Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	flinks, err := h.client.FLink.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	records := []model.FlinkResp{}
	for _, flink := range flinks {
		records = append(records, model.FlinkResp{
			ID:                 flink.ID,
			CreatedAt:          &flink.CreatedAt,
			UpdatedAt:          &flink.UpdatedAt,
			Name:               flink.Name,
			URL:                flink.URL,
			AvatarURL:          flink.AvatarURL,
			Description:        flink.Description,
			CoverURL:           flink.CoverURL,
			Status:             flink.Status,
			SnapshotUrl:        flink.SnapshotURL,
			Email:              flink.Email,
			EnableFriendCircle: flink.EnableFriendCircle,
			FriendCircleRuleID: flink.FriendCircleRuleID,
		})
	}
	pageResult := model.PageResult[model.FlinkResp]{
		Total:   int64(count),
		Records: records,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

// @Summary 创建Flink
// @Description 创建Flink
// @Tags flinks
// @Accept json
// @Produce json
// @Param flink_create_req body model.FlinkCreateReq true "Flink创建请求体"
// @Success 200 {object} model.HttpSuccess{data=ent.FLink}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flinks [post]
func (h *FlinkHandlerImpl) CreateFlink(c *fiber.Ctx) error {
	var createReq *model.FlinkCreateReq
	if err := c.BodyParser(&createReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	flink, err := h.client.FLink.Create().
		SetName(createReq.Name).
		SetURL(createReq.URL).
		SetAvatarURL(createReq.AvatarURL).
		SetDescription(createReq.Description).
		SetCoverURL(createReq.CoverURL).
		SetSnapshotURL(createReq.SnapshotURL).
		SetEmail(createReq.Email).
		SetEnableFriendCircle(createReq.EnableFriendCircle).
		SetNillableFriendCircleRuleID(createReq.FriendCircleRuleID).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flink))
}

// @Summary 更新Flink
// @Description 更新Flink
// @Tags flinks
// @Accept json
// @Produce json
// @Param id path int true "Flink ID"
// @Param flink_update_req body model.FlinkUpdateReq true "Flink更新请求体"
// @Success 200 {object} model.HttpSuccess{data=ent.FLink}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flinks/{id} [put]
func (h *FlinkHandlerImpl) UpdateFlink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid ID format"))
	}
	var updateReq *model.FlinkUpdateReq
	if err = c.BodyParser(&updateReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	flink, err := h.client.FLink.UpdateOneID(id).
		SetName(updateReq.Name).
		SetURL(updateReq.URL).
		SetAvatarURL(updateReq.AvatarURL).
		SetDescription(updateReq.Description).
		SetCoverURL(updateReq.CoverURL).
		SetSnapshotURL(updateReq.SnapshotURL).
		SetEmail(updateReq.Email).
		SetEnableFriendCircle(updateReq.EnableFriendCircle).
		SetFriendCircleRuleID(updateReq.FriendCircleRuleID).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flink))
}

// @Summary 查询Flink
// @Description 查询Flink
// @Tags flinks
// @Accept json
// @Produce json
// @Param id path int true "Flink ID"
// @Success 200 {object} model.HttpSuccess{data=ent.FLink}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flinks/{id} [get]
func (h *FlinkHandlerImpl) QueryFlink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	flink, err := h.client.FLink.Query().
		Where(flink.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flink))
}

// @Summary 删除Flink
// @Description 删除Flink
// @Tags flinks
// @Accept json
// @Produce json
// @Param id path int true "Flink ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/flinks/{id} [delete]
func (h *FlinkHandlerImpl) DeleteFlink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	if err := h.client.FLink.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
