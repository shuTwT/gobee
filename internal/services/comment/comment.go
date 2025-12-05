package comment

import (
	"context"
	"gobee/ent"
	"gobee/ent/comment"
	"gobee/internal/database"
	"gobee/pkg/domain/model"

	"github.com/medama-io/go-useragent"
)

func ListCommentPage(c context.Context, pageQuery model.PageQuery) (*model.PageResult[*ent.Comment], error) {
	client := database.DB

	count, err := client.Comment.Query().Count(c)
	if err != nil {
		return nil, err
	}
	comments, err := client.Comment.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c)
	if err != nil {
		return nil, err
	}
	pageResult := &model.PageResult[*ent.Comment]{
		Total:   int64(count),
		Records: comments,
	}
	return pageResult, nil
}

func ListComment(c context.Context, url string) ([]*ent.Comment, error) {
	client := database.DB
	comments, err := client.Comment.Query().
		Where(comment.URLEQ(url)).
		All(c)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func CountComment(c context.Context, includeReply bool, urls []string) (int64, error) {
	client := database.DB
	if includeReply {
		count, err := client.Comment.Query().
			Where(comment.URLIn(urls...)).
			Where(comment.Or(comment.ParentIDIsNil(), comment.ParentIDNotIn(0))).
			Count(c)
		if err != nil {
			return 0, err
		}
		return int64(count), nil
	}
	count, err := client.Comment.Query().
		Where(comment.URLIn(urls...)).
		Where(comment.ParentIDIsNil()).
		Count(c)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

func CreateComment(c context.Context, comment string, href string, link string, mail string, nick string, ua string, url string, ipAddress string) (*int, error) {
	client := database.DB
	entity, err := client.Comment.Create().
		SetContent(comment).
		SetUserAgent(ua).
		SetURL(url).
		SetIPAddress(ipAddress).
		Save(c)
	if err != nil {
		return nil, err
	}
	return &entity.ID, nil
}

func ParseUserAgent(ua string) (browser string, os string) {
	// Create a new parser. Initialize only once during application startup.
	parser := useragent.NewParser()
	agent := parser.Parse(ua)
	return agent.BrowserVersion(), agent.OS().String()
}

func GetRecentComment(c context.Context, pageSize int) ([]*ent.Comment, error) {
	client := database.DB
	comments, err := client.Comment.Query().
		Order(ent.Desc(comment.FieldCreatedAt)).
		Limit(pageSize).
		All(c)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
