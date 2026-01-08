package friendcirclerecord

import (
	"context"
	"gobee/ent"
	"gobee/ent/friendcirclerecord"
)

type FriendCircleRecordService interface {
	ListFriendCircleRecord(ctx context.Context) ([]*ent.FriendCircleRecord, error)
	ListFriendCircleRecordPage(ctx context.Context, page, size int) (int, []*ent.FriendCircleRecord, error)
	CreateFriendCircleRecord(ctx context.Context, author, title, linkURL, avatarURL, siteURL, publishedAt string) (*ent.FriendCircleRecord, error)
	UpdateFriendCircleRecord(ctx context.Context, id int, author, title, linkURL, avatarURL, siteURL, publishedAt string) (*ent.FriendCircleRecord, error)
	QueryFriendCircleRecord(ctx context.Context, id int) (*ent.FriendCircleRecord, error)
	DeleteFriendCircleRecord(ctx context.Context, id int) error
}

type FriendCircleRecordServiceImpl struct {
	client *ent.Client
}

func NewFriendCircleRecordServiceImpl(client *ent.Client) *FriendCircleRecordServiceImpl {
	return &FriendCircleRecordServiceImpl{client: client}
}

func (s *FriendCircleRecordServiceImpl) ListFriendCircleRecord(ctx context.Context) ([]*ent.FriendCircleRecord, error) {
	records, err := s.client.FriendCircleRecord.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *FriendCircleRecordServiceImpl) ListFriendCircleRecordPage(ctx context.Context, page, size int) (int, []*ent.FriendCircleRecord, error) {
	count, err := s.client.FriendCircleRecord.Query().Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	records, err := s.client.FriendCircleRecord.Query().
		Limit(size).
		Offset((page - 1) * size).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, records, nil
}

func (s *FriendCircleRecordServiceImpl) CreateFriendCircleRecord(ctx context.Context, author, title, linkURL, avatarURL, siteURL, publishedAt string) (*ent.FriendCircleRecord, error) {
	builder := s.client.FriendCircleRecord.Create().
		SetAuthor(author).
		SetTitle(title).
		SetLinkURL(linkURL).
		SetAvatarURL(avatarURL)

	if siteURL != "" {
		builder.SetSiteURL(siteURL)
	}
	if publishedAt != "" {
		builder.SetPublishedAt(publishedAt)
	}

	newRecord, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return newRecord, nil
}

func (s *FriendCircleRecordServiceImpl) UpdateFriendCircleRecord(ctx context.Context, id int, author, title, linkURL, avatarURL, siteURL, publishedAt string) (*ent.FriendCircleRecord, error) {
	builder := s.client.FriendCircleRecord.UpdateOneID(id).
		SetAuthor(author).
		SetTitle(title).
		SetLinkURL(linkURL).
		SetAvatarURL(avatarURL)

	if siteURL != "" {
		builder.SetSiteURL(siteURL)
	} else {
		builder.ClearSiteURL()
	}
	if publishedAt != "" {
		builder.SetPublishedAt(publishedAt)
	} else {
		builder.ClearPublishedAt()
	}

	updatedRecord, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedRecord, nil
}

func (s *FriendCircleRecordServiceImpl) QueryFriendCircleRecord(ctx context.Context, id int) (*ent.FriendCircleRecord, error) {
	record, err := s.client.FriendCircleRecord.Query().
		Where(friendcirclerecord.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *FriendCircleRecordServiceImpl) DeleteFriendCircleRecord(ctx context.Context, id int) error {
	err := s.client.FriendCircleRecord.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
