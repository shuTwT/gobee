package friendcircle

import (
	"context"
	"gobee/ent"
	"gobee/ent/friendcirclerecord"
)

type FriendCircleService interface {
	ExistsRecord(c context.Context, linkUrl string) (bool, error)
	InsertRecord(c context.Context, author string, avatarUrl string, title string, linkUrl string, publishedAt string) error
}

type FriendCircleServiceImpl struct {
	client *ent.Client
}

func NewFriendCircleServiceImpl(client *ent.Client) *FriendCircleServiceImpl {
	return &FriendCircleServiceImpl{client: client}
}

// 判断是否已存在
func (s *FriendCircleServiceImpl) ExistsRecord(c context.Context, linkUrl string) (bool, error) {
	client := s.client
	return client.FriendCircleRecord.Query().Where(friendcirclerecord.LinkURLEQ(linkUrl)).Exist(c)
}

// 插入文章
func (s *FriendCircleServiceImpl) InsertRecord(c context.Context, author string, avatarUrl string, title string, linkUrl string, publishedAt string) error {
	client := s.client
	_, err := client.FriendCircleRecord.Create().
		SetAuthor(author).
		SetAvatarURL(avatarUrl).
		SetTitle(title).
		SetLinkURL(linkUrl).
		SetPublishedAt(publishedAt).
		Save(c)
	return err
}
