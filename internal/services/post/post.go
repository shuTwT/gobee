package post

import (
	"context"
	"gobee/ent"
	"gobee/internal/database"
)

func UpdatePostSetting(c context.Context, id int, post *ent.Post) (*ent.Post, error) {
	client := database.DB
	newPost, err := client.Post.UpdateOneID(id).
		SetTitle(post.Title).
		SetCover(post.Cover).
		SetKeywords(post.Keywords).
		SetCopyright(post.Copyright).
		SetAuthor(post.Author).
		SetStatus(post.Status).
		SetIsAutogenSummary(post.IsAutogenSummary).
		SetIsVisible(post.IsVisible).
		SetIsTipToTop(post.IsTipToTop).
		SetIsAllowComment(post.IsAllowComment).
		Save(c)
	return newPost, err
}

func GetPostCount(c context.Context) (int, error) {
	client := database.DB
	count, err := client.Post.Query().Count(c)
	if err != nil {
		return 0, err
	}
	return count, nil
}
