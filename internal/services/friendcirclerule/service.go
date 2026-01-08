package friendcirclerule

import (
	"context"
	"gobee/ent"
	"gobee/ent/friendcirclerule"
)

type FriendCircleRuleService interface {
	ListFriendCircleRule(ctx context.Context) ([]*ent.FriendCircleRule, error)
	ListFriendCircleRulePage(ctx context.Context, page, size int) (int, []*ent.FriendCircleRule, error)
	CreateFriendCircleRule(ctx context.Context, name, titleSelector, linkSelector, createdSelector, updatedSelector string) (*ent.FriendCircleRule, error)
	UpdateFriendCircleRule(ctx context.Context, id int, name, titleSelector, linkSelector, createdSelector, updatedSelector string) (*ent.FriendCircleRule, error)
	QueryFriendCircleRule(ctx context.Context, id int) (*ent.FriendCircleRule, error)
	DeleteFriendCircleRule(ctx context.Context, id int) error
}

type FriendCircleRuleServiceImpl struct {
	client *ent.Client
}

func NewFriendCircleRuleServiceImpl(client *ent.Client) *FriendCircleRuleServiceImpl {
	return &FriendCircleRuleServiceImpl{client: client}
}

func (s *FriendCircleRuleServiceImpl) ListFriendCircleRule(ctx context.Context) ([]*ent.FriendCircleRule, error) {
	rules, err := s.client.FriendCircleRule.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (s *FriendCircleRuleServiceImpl) ListFriendCircleRulePage(ctx context.Context, page, size int) (int, []*ent.FriendCircleRule, error) {
	count, err := s.client.FriendCircleRule.Query().Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	rules, err := s.client.FriendCircleRule.Query().
		Limit(size).
		Offset((page - 1) * size).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, rules, nil
}

func (s *FriendCircleRuleServiceImpl) CreateFriendCircleRule(ctx context.Context, name, titleSelector, linkSelector, createdSelector, updatedSelector string) (*ent.FriendCircleRule, error) {
	builder := s.client.FriendCircleRule.Create().
		SetName(name)

	if titleSelector != "" {
		builder.SetTitleSelector(titleSelector)
	}
	if linkSelector != "" {
		builder.SetLinkSelector(linkSelector)
	}
	if createdSelector != "" {
		builder.SetCreatedSelector(createdSelector)
	}
	if updatedSelector != "" {
		builder.SetUpdatedSelector(updatedSelector)
	}

	newRule, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return newRule, nil
}

func (s *FriendCircleRuleServiceImpl) UpdateFriendCircleRule(ctx context.Context, id int, name, titleSelector, linkSelector, createdSelector, updatedSelector string) (*ent.FriendCircleRule, error) {
	builder := s.client.FriendCircleRule.UpdateOneID(id).
		SetName(name)

	if titleSelector != "" {
		builder.SetTitleSelector(titleSelector)
	} else {
		builder.ClearTitleSelector()
	}
	if linkSelector != "" {
		builder.SetLinkSelector(linkSelector)
	} else {
		builder.ClearLinkSelector()
	}
	if createdSelector != "" {
		builder.SetCreatedSelector(createdSelector)
	} else {
		builder.ClearCreatedSelector()
	}
	if updatedSelector != "" {
		builder.SetUpdatedSelector(updatedSelector)
	} else {
		builder.ClearUpdatedSelector()
	}

	updatedRule, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedRule, nil
}

func (s *FriendCircleRuleServiceImpl) QueryFriendCircleRule(ctx context.Context, id int) (*ent.FriendCircleRule, error) {
	rule, err := s.client.FriendCircleRule.Query().
		Where(friendcirclerule.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func (s *FriendCircleRuleServiceImpl) DeleteFriendCircleRule(ctx context.Context, id int) error {
	err := s.client.FriendCircleRule.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
