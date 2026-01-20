package schedulejob

import (
	"context"
	"errors"

	"github.com/shuTwT/gobee/ent"
	"github.com/shuTwT/gobee/ent/schedulejob"
	"github.com/shuTwT/gobee/pkg/domain/model"
)

type ScheduleJobService interface {
	ListScheduleJob(ctx context.Context) ([]*ent.ScheduleJob, error)
	ListScheduleJobPage(ctx context.Context, page, size int) (int, []*ent.ScheduleJob, error)
	QueryScheduleJob(ctx context.Context, id int) (*ent.ScheduleJob, error)
	CreateScheduleJob(ctx context.Context, req *model.CreateScheduleJobReq) (*ent.ScheduleJob, error)
	UpdateScheduleJob(ctx context.Context, id int, req *model.UpdateScheduleJobReq) (*ent.ScheduleJob, error)
	DeleteScheduleJob(ctx context.Context, id int) error
}

type ScheduleJobServiceImpl struct {
	client *ent.Client
}

func NewScheduleJobServiceImpl(client *ent.Client) *ScheduleJobServiceImpl {
	return &ScheduleJobServiceImpl{client: client}
}

func (s *ScheduleJobServiceImpl) ListScheduleJob(ctx context.Context) ([]*ent.ScheduleJob, error) {
	jobs, err := s.client.ScheduleJob.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (s *ScheduleJobServiceImpl) ListScheduleJobPage(ctx context.Context, page, size int) (int, []*ent.ScheduleJob, error) {
	count, err := s.client.ScheduleJob.Query().Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	jobs, err := s.client.ScheduleJob.Query().
		Order(ent.Desc(schedulejob.FieldID)).
		Offset((page - 1) * size).
		Limit(size).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, jobs, nil
}

func (s *ScheduleJobServiceImpl) QueryScheduleJob(ctx context.Context, id int) (*ent.ScheduleJob, error) {
	job, err := s.client.ScheduleJob.Query().
		Where(schedulejob.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (s *ScheduleJobServiceImpl) CreateScheduleJob(ctx context.Context, req *model.CreateScheduleJobReq) (*ent.ScheduleJob, error) {
	if err := validateCreateScheduleJobReq(req); err != nil {
		return nil, err
	}

	builder := s.client.ScheduleJob.Create().
		SetName(req.Name).
		SetType(req.Type).
		SetExpression(req.Expression).
		SetExecutionType(req.ExecutionType).
		SetEnabled(req.Enabled).
		SetMaxRetries(req.MaxRetries).
		SetFailureNotification(req.FailureNotification)

	if req.Description != nil {
		builder.SetDescription(*req.Description)
	}

	if req.HTTPMethod != nil {
		builder.SetHTTPMethod(*req.HTTPMethod)
	}

	if req.HTTPURL != nil {
		builder.SetHTTPURL(*req.HTTPURL)
	}

	if req.HTTPHeaders != nil {
		builder.SetHTTPHeaders(req.HTTPHeaders)
	}

	if req.HTTPBody != nil {
		builder.SetHTTPBody(*req.HTTPBody)
	}

	if req.HTTPTimeout != nil {
		builder.SetHTTPTimeout(*req.HTTPTimeout)
	}

	job, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *ScheduleJobServiceImpl) UpdateScheduleJob(ctx context.Context, id int, req *model.UpdateScheduleJobReq) (*ent.ScheduleJob, error) {
	if err := validateUpdateScheduleJobReq(req); err != nil {
		return nil, err
	}

	exists, err := s.client.ScheduleJob.Query().Where(schedulejob.ID(id)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("定时任务不存在")
	}

	builder := s.client.ScheduleJob.UpdateOneID(id)

	if req.Name != nil {
		builder.SetName(*req.Name)
	}

	if req.Type != nil {
		builder.SetType(*req.Type)
	}

	if req.Expression != nil {
		builder.SetExpression(*req.Expression)
	}

	if req.Description != nil {
		builder.SetDescription(*req.Description)
	}

	if req.Enabled != nil {
		builder.SetEnabled(*req.Enabled)
	}

	if req.ExecutionType != nil {
		builder.SetExecutionType(*req.ExecutionType)
	}

	if req.HTTPMethod != nil {
		builder.SetHTTPMethod(*req.HTTPMethod)
	}

	if req.HTTPURL != nil {
		builder.SetHTTPURL(*req.HTTPURL)
	}

	if req.HTTPHeaders != nil {
		builder.SetHTTPHeaders(req.HTTPHeaders)
	}

	if req.HTTPBody != nil {
		builder.SetHTTPBody(*req.HTTPBody)
	}

	if req.HTTPTimeout != nil {
		builder.SetHTTPTimeout(*req.HTTPTimeout)
	}

	if req.MaxRetries != nil {
		builder.SetMaxRetries(*req.MaxRetries)
	}

	if req.FailureNotification != nil {
		builder.SetFailureNotification(*req.FailureNotification)
	}

	job, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *ScheduleJobServiceImpl) DeleteScheduleJob(ctx context.Context, id int) error {
	err := s.client.ScheduleJob.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func validateCreateScheduleJobReq(req *model.CreateScheduleJobReq) error {
	if req.Name == "" {
		return errors.New("任务名称不能为空")
	}
	if req.Type == "" {
		return errors.New("任务类型不能为空")
	}
	if req.Expression == "" {
		return errors.New("调度表达式不能为空")
	}
	if req.ExecutionType == "" {
		return errors.New("执行类型不能为空")
	}
	return nil
}

func validateUpdateScheduleJobReq(req *model.UpdateScheduleJobReq) error {
	return nil
}
