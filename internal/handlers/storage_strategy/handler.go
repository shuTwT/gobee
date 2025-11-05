package storagestrategy

import (
	"gobee/ent"
	"gobee/ent/storagestrategy"
	"gobee/internal/database"
	"gobee/internal/model"
	"gobee/internal/services/storage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StorageStrategyList struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Master bool   `json:"master"`
}

func ListStorageStrategy(c *fiber.Ctx) error {
	client := database.DB
	strategies, err := client.StorageStrategy.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", strategies))
}

func ListStorageStrategyAll(c *fiber.Ctx) error {
	client := database.DB
	strategies, err := client.StorageStrategy.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	var strategyList []StorageStrategyList
	for _, strategy := range strategies {
		strategyList = append(strategyList, StorageStrategyList{
			ID:     strategy.ID,
			Name:   strategy.Name,
			Type:   string(strategy.Type),
			Master: strategy.Master,
		})
	}
	return c.JSON(model.NewSuccess("success", strategyList))
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
		SetRegion(strategy.Region).
		SetEndpoint(strategy.Endpoint).
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
		SetRegion(strategy.Region).
		SetEndpoint(strategy.Endpoint).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	// 清除缓存
	storage.ClearCache()
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
	// 清除缓存
	storage.ClearCache()
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
	// 清除缓存
	storage.ClearCache()
	return c.JSON(model.NewSuccess("success", nil))
}
