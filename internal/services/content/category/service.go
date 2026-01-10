package category

import (
	"context"
	"gobee/ent"
	"gobee/ent/category"
	"gobee/internal/database"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	QueryCategory(c *fiber.Ctx, id int) (*ent.Category, error)
	QueryCategoryList(c *fiber.Ctx) ([]*ent.Category, error)
	QueryCategoryPage(c *fiber.Ctx, pageQuery model.PageQuery) (int, []*ent.Category, error)
	CreateCategory(c context.Context, createReq model.CategoryCreateReq) (*ent.Category, error)
	UpdateCategory(c *fiber.Ctx, id int, updateReq *model.CategoryUpdateReq) (*ent.Category, error)
	DeleteCategory(c *fiber.Ctx, id int) error
}

type CategoryServiceImpl struct {
	client *ent.Client
}

func NewCategoryServiceImpl(client *ent.Client) *CategoryServiceImpl {
	return &CategoryServiceImpl{client: client}
}

func (s *CategoryServiceImpl) QueryCategory(c *fiber.Ctx, id int) (*ent.Category, error) {
	client := database.DB
	category, err := client.Category.Query().
		Where(category.ID(id)).
		Only(c.Context())
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryServiceImpl) QueryCategoryList(c *fiber.Ctx) ([]*ent.Category, error) {
	client := database.DB
	categories, err := client.Category.Query().
		All(c.Context())
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryServiceImpl) QueryCategoryPage(c *fiber.Ctx, pageQuery model.PageQuery) (int, []*ent.Category, error) {
	client := database.DB
	count, err := client.Category.Query().Count(c.UserContext())
	if err != nil {
		return 0, nil, err
	}

	categories, err := client.Category.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		return 0, nil, err
	}

	return count, categories, nil
}

func (s *CategoryServiceImpl) CreateCategory(c context.Context, createReq model.CategoryCreateReq) (*ent.Category, error) {
	client := database.DB
	category, err := client.Category.Create().
		SetName(createReq.Name).
		SetNillableDescription(createReq.Description).
		SetNillableSlug(createReq.Slug).
		SetSortOrder(createReq.SortOrder).
		SetNillableActive(createReq.Active).
		Save(c)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryServiceImpl) UpdateCategory(c *fiber.Ctx, id int, updateReq *model.CategoryUpdateReq) (*ent.Category, error) {
	client := database.DB
	update := client.Category.UpdateOneID(id)

	if updateReq.Name != nil {
		update.SetName(*updateReq.Name)
	}

	if updateReq.Description != nil {
		update.SetNillableDescription(updateReq.Description)
	}

	if updateReq.Slug != nil {
		update.SetNillableSlug(updateReq.Slug)
	}

	if updateReq.SortOrder != nil {
		update.SetNillableSortOrder(updateReq.SortOrder)
	}

	if updateReq.Active != nil {
		update.SetNillableActive(updateReq.Active)
	}

	updatedCategory, err := update.Save(c.Context())
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (s *CategoryServiceImpl) DeleteCategory(c *fiber.Ctx, id int) error {
	client := database.DB
	err := client.Category.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		return err
	}
	return nil
}
