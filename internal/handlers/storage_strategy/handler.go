package storagestrategy

import (
	"gobee/ent"
	"gobee/ent/storagestrategy"
	"gobee/internal/database"
	"gobee/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListStorageStrategy(c *fiber.Ctx) error {
	client := database.DB
	strategies, err := client.StorageStrategy.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", strategies))
}

func CreateStorageStrategy(c *fiber.Ctx) error {
	client := database.DB
	var strategy *ent.StorageStrategy
	if err := c.BodyParser(&strategy); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newStrategy, err := client.StorageStrategy.Create().
		SetName(strategy.Name).
		SetType(strategy.Type).
		SetNodeID(strategy.NodeID).
		SetBucket(strategy.Bucket).
		SetAccessKey(strategy.AccessKey).
		SetSecretKey(strategy.SecretKey).
		SetBasePath(strategy.BasePath).
		SetDomain(strategy.Domain).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newStrategy))
}

func UpdateStorageStrategy(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var strategy *ent.StorageStrategy
	if err := c.BodyParser(&strategy); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newStrategy, err := client.StorageStrategy.UpdateOneID(id).
		SetName(strategy.Name).
		SetType(strategy.Type).
		SetNodeID(strategy.NodeID).
		SetBucket(strategy.Bucket).
		SetAccessKey(strategy.AccessKey).
		SetSecretKey(strategy.SecretKey).
		SetBasePath(strategy.BasePath).
		SetDomain(strategy.Domain).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", newStrategy))
}

func QueryStorageStrategy(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	strategy, err := client.StorageStrategy.Query().
		Where(storagestrategy.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", strategy))
}

func DeleteStorageStrategy(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	if err := client.StorageStrategy.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}

func SetDefaultStorageStrategy(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	// 先将所有策略的 master 字段设置为 false
	if err := client.StorageStrategy.Update().SetMaster(false).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	// 将指定 ID 的策略的 master 字段设置为 true
	if err := client.StorageStrategy.UpdateOneID(id).SetMaster(true).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", nil))
}
