package flink

import (
	"gobee/ent/flink"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListFlink(c *fiber.Ctx) error {
	client := database.DB
	flinks, err := client.FLink.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	result := []model.FlinkResp{}
	for _, flink := range flinks {
		result = append(result, model.FlinkResp{
			ID:          flink.ID,
			CreatedAt:   &flink.CreatedAt,
			UpdatedAt:   &flink.UpdatedAt,
			Name:        flink.Name,
			URL:         flink.URL,
			Logo:        flink.Logo,
			Description: flink.Description,
			Status:      flink.Status,
			Snapshot:    flink.Snapshot,
			Email:       flink.Email,
		})
	}
	return c.JSON(model.NewSuccess("success", result))
}

func ListFlinkPage(c *fiber.Ctx) error {
	client := database.DB
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	count, err := client.FLink.Query().Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	flinks, err := client.FLink.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	records := []model.FlinkResp{}
	for _, flink := range flinks {
		records = append(records, model.FlinkResp{
			ID:          flink.ID,
			CreatedAt:   &flink.CreatedAt,
			UpdatedAt:   &flink.UpdatedAt,
			Name:        flink.Name,
			URL:         flink.URL,
			Logo:        flink.Logo,
			Description: flink.Description,
			Status:      flink.Status,
			Snapshot:    flink.Snapshot,
			Email:       flink.Email,
		})
	}
	pageResult := model.PageResult[model.FlinkResp]{
		Total:   int64(count),
		Records: records,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

func CreateFlink(c *fiber.Ctx) error {
	var createReq *model.FlinkCreateReq
	if err := c.BodyParser(&createReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	flink, err := client.FLink.Create().
		SetName(createReq.Name).
		SetURL(createReq.URL).
		SetLogo(createReq.Logo).
		SetDescription(createReq.Description).
		SetSnapshot(createReq.Snapshot).
		SetEmail(createReq.Email).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flink))
}

func UpdateFlink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "Invalid ID format"))
	}
	var updateReq *model.FlinkUpdateReq
	if err = c.BodyParser(&updateReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	client := database.DB
	flink, err := client.FLink.UpdateOneID(id).
		SetName(updateReq.Name).
		SetURL(updateReq.URL).
		SetLogo(updateReq.Logo).
		SetDescription(updateReq.Description).
		SetSnapshot(updateReq.Snapshot).
		SetEmail(updateReq.Email).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flink))
}

func QueryFlink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := database.DB
	flink, err := client.FLink.Query().
		Where(flink.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", flink))
}

func DeleteFlink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	client := database.DB
	if err := client.FLink.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
