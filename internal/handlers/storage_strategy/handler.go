package storagestrategy

import (
	"gobee/ent"
	"gobee/ent/storagestrategy"
	"gobee/internal/services/storage"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StorageStrategyHandler interface {
	ListStorageStrategy(c *fiber.Ctx) error
	ListStorageStrategyAll(c *fiber.Ctx) error
	CreateStorageStrategy(c *fiber.Ctx) error
	UpdateStorageStrategy(c *fiber.Ctx) error
	QueryStorageStrategy(c *fiber.Ctx) error
	DeleteStorageStrategy(c *fiber.Ctx) error
	SetDefaultStorageStrategy(c *fiber.Ctx) error
}

type StorageStrategyHandlerImpl struct {
	client *ent.Client
}

func NewStorageStrategyHandlerImpl(client *ent.Client) *StorageStrategyHandlerImpl {
	return &StorageStrategyHandlerImpl{
		client: client,
	}
}

// @Summary 获取存储策略列表
// @Description 获取所有存储策略的列表
// @Tags storageStrategy
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]ent.StorageStrategy}
// @Failure 500 {object} model.HttpError
// @Router /api/v1/storage-strategy/list [get]
func (h *StorageStrategyHandlerImpl) ListStorageStrategy(c *fiber.Ctx) error {
	strategies, err := h.client.StorageStrategy.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", strategies))
}

// @Summary 获取所有存储策略列表
// @Description 获取所有存储策略的列表，包括默认策略
// @Tags storageStrategy
// @Accept json
// @Produce json
// @Success 200 {object} model.HttpSuccess{data=[]model.StorageStrategyListResp}
// @Failure 500 {object} model.HttpError
// @Router /api/v1/storage-strategy/list-all [get]
func (h *StorageStrategyHandlerImpl) ListStorageStrategyAll(c *fiber.Ctx) error {
	strategies, err := h.client.StorageStrategy.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	var strategyList []model.StorageStrategyListResp
	for _, strategy := range strategies {
		strategyList = append(strategyList, model.StorageStrategyListResp{
			ID:     strategy.ID,
			Name:   strategy.Name,
			Type:   string(strategy.Type),
			Master: strategy.Master,
		})
	}
	return c.JSON(model.NewSuccess("success", strategyList))
}

// @Summary 创建存储策略
// @Description 创建一个新的存储策略
// @Tags storageStrategy
// @Accept json
// @Produce json
// @Param strategy body model.StorageStrategyCreateReq true "Storage Strategy Create Request"
// @Success 200 {object} model.HttpSuccess{data=ent.StorageStrategy}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/storage-strategy/create [post]
func (h *StorageStrategyHandlerImpl) CreateStorageStrategy(c *fiber.Ctx) error {
	var strategy *ent.StorageStrategy
	if err := c.BodyParser(&strategy); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newStrategy, err := h.client.StorageStrategy.Create().
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

// @Summary 更新存储策略
// @Description 更新指定ID的存储策略
// @Tags storageStrategy
// @Accept json
// @Produce json
// @Param id path int true "Storage Strategy ID"
// @Param strategy body model.StorageStrategyUpdateReq true "Storage Strategy Update Request"
// @Success 200 {object} model.HttpSuccess{data=ent.StorageStrategy}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/storage-strategy/update/{id} [put]
func (h *StorageStrategyHandlerImpl) UpdateStorageStrategy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	var strategy *model.StorageStrategyUpdateReq
	if err = c.BodyParser(&strategy); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	newStrategy, err := h.client.StorageStrategy.UpdateOneID(id).
		SetName(strategy.Name).
		SetType(storagestrategy.Type(strategy.Type)).
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

// @Summary 查询存储策略
// @Description 查询指定ID的存储策略
// @Tags storageStrategy
// @Accept json
// @Produce json
// @Param id path int true "Storage Strategy ID"
// @Success 200 {object} model.HttpSuccess{data=ent.StorageStrategy}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/storage-strategy/query/{id} [get]
func (h *StorageStrategyHandlerImpl) QueryStorageStrategy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	strategy, err := h.client.StorageStrategy.Query().
		Where(storagestrategy.ID(id)).
		First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", strategy))
}

// @Summary 删除存储策略
// @Description 删除指定ID的存储策略
// @Tags storageStrategy
// @Accept json
// @Produce json
// @Param id path int true "Storage Strategy ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/storage-strategy/delete/{id} [delete]
func (h *StorageStrategyHandlerImpl) DeleteStorageStrategy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	if err := h.client.StorageStrategy.DeleteOneID(id).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	// 清除缓存
	storage.ClearCache()
	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 设置默认存储策略
// @Description 设置指定ID的存储策略为默认策略
// @Tags storageStrategy
// @Accept json
// @Produce json
// @Param id path int true "Storage Strategy ID"
// @Success 200 {object} model.HttpSuccess{data=nil}
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/storage-strategy/default/{id} [put]
func (h *StorageStrategyHandlerImpl) SetDefaultStorageStrategy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format"))
	}
	// 先将所有策略的 master 字段设置为 false
	if err := h.client.StorageStrategy.Update().SetMaster(false).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	// 将指定 ID 的策略的 master 字段设置为 true
	if err := h.client.StorageStrategy.UpdateOneID(id).SetMaster(true).Exec(c.Context()); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	// 清除缓存
	storage.ClearCache()
	return c.JSON(model.NewSuccess("success", nil))
}
