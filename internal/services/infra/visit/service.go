package visit

import (
	"context"
	"gobee/ent"
	"gobee/pkg/domain/model"
)

type VisitService interface {
	CreateVisitLog(ctx context.Context, req model.VisitLogReq) error
}

type VisitServiceImpl struct {
	client *ent.Client
}

func NewVisitServiceImpl(client *ent.Client) VisitService {
	return &VisitServiceImpl{client: client}
}

func (s *VisitServiceImpl) CreateVisitLog(ctx context.Context, req model.VisitLogReq) error {
	_, err := s.client.VisitLog.Create().
		SetIP(req.IP).
		SetUserAgent(req.UserAgent).
		SetPath(req.Path).
		SetOs(req.OS).
		SetBrowser(req.Browser).
		SetDevice(req.Device).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
