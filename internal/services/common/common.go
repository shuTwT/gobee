package common

import (
	"context"
	"gobee/ent"
	comment_service "gobee/internal/services/comment"
	post_service "gobee/internal/services/post"
	user_service "gobee/internal/services/user"
	"gobee/pkg/domain/model"
)

type CommonService interface {
	GetHomeStatistic(c context.Context) model.HomeStatistic
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
