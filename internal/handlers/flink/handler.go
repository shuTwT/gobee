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
