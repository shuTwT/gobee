package friendcircle

import (
	"context"
	"gobee/ent/friendcirclerecord"
	"gobee/internal/database"
)

// 判断是否已存在
func ExistsRecord(c context.Context, linkUrl string) (bool, error) {
	client := database.DB
	return client.FriendCircleRecord.Query().Where(friendcirclerecord.LinkURLEQ(linkUrl)).Exist(c)
}

// 插入文章
func InsertRecord(c context.Context, author string, avatarUrl string, title string, linkUrl string, publishedAt string) error {
	client := database.DB
	_, err := client.FriendCircleRecord.Create().
		SetAuthor(author).
		SetAvatarURL(avatarUrl).
		SetTitle(title).
		SetLinkURL(linkUrl).
		SetPublishedAt(publishedAt).
		Save(c)
	return err
}
