package common

import (
	"context"
	comment_service "gobee/internal/services/comment"
	post_service "gobee/internal/services/post"
	user_service "gobee/internal/services/user"
	"gobee/pkg/domain/model"
)

func GetHomeStatistic(c context.Context) model.HomeStatistic {
	userCount, _ := user_service.GetUserCount(c)
	postCount, _ := post_service.GetPostCount(c)
	commentCount, _ := comment_service.GetCommentCount(c)
	homeStatistic := model.HomeStatistic{
		PostCount:    postCount,
		UserCount:    userCount,
		CommentCount: commentCount,
		VisitCount:   0,
	}
	return homeStatistic
}
