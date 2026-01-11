package post

import (
	"context"
	"gobee/ent"
	"gobee/ent/category"
	"gobee/ent/post"
	"gobee/ent/tag"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
)

type PostService interface {
	QueryPostList(c context.Context) ([]*ent.Post, error)
	QueryPostPage(c context.Context, req model.PostPageReq) ([]*ent.Post, int, error)
	CreatePost(c context.Context, title string, content string) (*ent.Post, error)
	UpdatePostContent(c context.Context, id int, content string) (*ent.Post, error)
	UpdatePostSetting(c context.Context, id int, post model.PostUpdateReq) (*ent.Post, error)
	GetPostCount(c context.Context) (int, error)
}

type PostServiceImpl struct {
	client *ent.Client
}

func NewPostServiceImpl(client *ent.Client) *PostServiceImpl {
	return &PostServiceImpl{client: client}
}

func (s *PostServiceImpl) QueryPostList(c context.Context) ([]*ent.Post, error) {
	posts, err := s.client.Post.Query().
		WithCategories().
		WithTags().
		Order(ent.Desc(post.FieldID)).
		All(c)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostServiceImpl) QueryPostPage(c context.Context, req model.PostPageReq) ([]*ent.Post, int, error) {
	query := s.client.Post.Query()

	if req.CategoryID != nil {
		query.Where(post.HasCategoriesWith(category.ID(*req.CategoryID)))
	}

	if req.TagID != nil {
		query.Where(post.HasTagsWith(tag.ID(*req.TagID)))
	}

	if req.Title != "" {
		query.Where(post.TitleContains(req.Title))
	}

	count, err := query.Count(c)
	if err != nil {
		return nil, 0, err
	}
	posts, err := query.
		WithCategories().
		WithTags().
		Order(ent.Desc(post.FieldID)).
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	return posts, count, nil
}

func (s *PostServiceImpl) CreatePost(c context.Context, title string, content string) (*ent.Post, error) {

	newPost, err := s.client.Post.Create().
		SetTitle(title).
		SetContent(content).
		Save(c)
	return newPost, err
}

func (s *PostServiceImpl) UpdatePostContent(c context.Context, id int, content string) (*ent.Post, error) {
	newPost, err := s.client.Post.UpdateOneID(id).
		SetContent(content).
		Save(c)
	return newPost, err
}

func (s *PostServiceImpl) UpdatePostSetting(c context.Context, id int, post model.PostUpdateReq) (*ent.Post, error) {
	client := s.client
	var summary string
	if post.IsAutogenSummary {
		summary = "生成失败"
	} else {
		summary = post.Summary
	}
	newPost, err := client.Post.UpdateOneID(id).
		SetTitle(post.Title).
		SetNillableAlias(post.Alias).
		SetCover(post.Cover).
		SetKeywords(post.Keywords).
		SetCopyright(post.Copyright).
		SetAuthor(post.Author).
		SetIsAutogenSummary(post.IsAutogenSummary).
		SetIsVisible(post.IsVisible).
		SetIsPinToTop(post.IsPinToTop).
		SetIsAllowComment(post.IsAllowComment).
		SetIsVisibleAfterComment(post.IsVisibleAfterComment).
		SetIsVisibleAfterPay(post.IsVisibleAfterPay).
		SetPrice(int(post.Price * 100)).
		SetSummary(summary).
		AddCategoryIDs(post.Categories...).
		AddTagIDs(post.Tags...).
		Save(c)
	return newPost, err
}

func (s *PostServiceImpl) GetPostCount(c context.Context) (int, error) {
	client := database.DB
	count, err := client.Post.Query().Count(c)
	if err != nil {
		return 0, err
	}
	return count, nil
}
