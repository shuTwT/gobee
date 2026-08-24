package common

import (
	"context"
	"fmt"
	"unicode"

	"github.com/shuTwT/hoshikuzu/ent/post"

	comment_service "github.com/shuTwT/hoshikuzu/internal/services/content/comment"
	post_service "github.com/shuTwT/hoshikuzu/internal/services/content/post"
	user_service "github.com/shuTwT/hoshikuzu/internal/services/system/user"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
)

type CommonService interface {
	GetHomeStatistic(c context.Context) model.HomeStatistic
	GetSiteStatistic(c context.Context) (*model.SiteStatistic, error)
}

type CommonServiceImpl struct {
	client         *ent.Client
	userService    user_service.UserService
	postService    post_service.PostService
	commentService comment_service.CommentService
}

func NewCommonServiceImpl(client *ent.Client, userService user_service.UserService, postService post_service.PostService, commentService comment_service.CommentService) *CommonServiceImpl {
	return &CommonServiceImpl{client: client, userService: userService, postService: postService, commentService: commentService}
}

func (s *CommonServiceImpl) GetHomeStatistic(c context.Context) model.HomeStatistic {
	userCount, _ := s.userService.GetUserCount(c)
	postCount, _ := s.postService.GetPostCount(c)
	commentCount, _ := s.commentService.GetCommentCount(c)
	homeStatistic := model.HomeStatistic{
		PostCount:    postCount,
		UserCount:    userCount,
		CommentCount: commentCount,
		VisitCount:   0,
	}
	return homeStatistic
}

func (s *CommonServiceImpl) GetSiteStatistic(c context.Context) (*model.SiteStatistic, error) {
	publishedPosts := s.client.Post.Query().Where(
		post.StatusEQ("published"),
		post.IsVisible(true),
	)

	postCount, err := publishedPosts.Count(c)
	if err != nil {
		return nil, fmt.Errorf("统计文章数量失败: %w", err)
	}

	categoryCount, err := s.client.Category.Query().Count(c)
	if err != nil {
		return nil, fmt.Errorf("统计分类数量失败: %w", err)
	}

	tagCount, err := s.client.Tag.Query().Count(c)
	if err != nil {
		return nil, fmt.Errorf("统计标签数量失败: %w", err)
	}

	posts, err := publishedPosts.Select(post.FieldMdContent, post.FieldContent).All(c)
	if err != nil {
		return nil, fmt.Errorf("统计总字数失败: %w", err)
	}
	totalWordCount := 0
	for _, p := range posts {
		content := p.Content
		if p.MdContent != nil && *p.MdContent != "" {
			content = *p.MdContent
		}
		totalWordCount += countTextLength(content)
	}

	return &model.SiteStatistic{
		PostCount:      postCount,
		CategoryCount:  categoryCount,
		TagCount:       tagCount,
		TotalWordCount: totalWordCount,
	}, nil
}

// countTextLength 统计去除空白字符后的字符数，作为文章字数
func countTextLength(text string) int {
	count := 0
	for _, r := range text {
		if !unicode.IsSpace(r) {
			count++
		}
	}
	return count
}
