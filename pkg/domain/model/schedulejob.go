package model

import "time"

type CreateScheduleJobReq struct {
	Name                 string            `json:"name" validate:"required"`
	Type                 string            `json:"type" validate:"required,oneof=cron interval once"`
	Expression           string            `json:"expression" validate:"required"`
	Description          *string           `json:"description,omitempty"`
	Enabled              bool              `json:"enabled"`
	ExecutionType        string            `json:"execution_type" validate:"required,oneof=http internal command mq"`
	HTTPMethod           *string           `json:"http_method,omitempty" validate:"omitempty,oneof=GET POST PUT DELETE"`
	HTTPURL              *string           `json:"http_url,omitempty" validate:"omitempty,url"`
	HTTPHeaders          map[string]string `json:"http_headers,omitempty"`
	HTTPBody             *string           `json:"http_body,omitempty"`
	HTTPTimeout          *int              `json:"http_timeout,omitempty" validate:"omitempty,min=1,max=300"`
	MaxRetries           int               `json:"max_retries" validate:"min=0,max=10"`
	FailureNotification  bool              `json:"failure_notification"`
}

type UpdateScheduleJobReq struct {
	Name                *string           `json:"name,omitempty"`
	Type                *string           `json:"type,omitempty" validate:"omitempty,oneof=cron interval once"`
	Expression          *string           `json:"expression,omitempty"`
	Description         *string           `json:"description,omitempty"`
	Enabled             *bool             `json:"enabled,omitempty"`
	ExecutionType       *string           `json:"execution_type,omitempty" validate:"omitempty,oneof=http internal command mq"`
	HTTPMethod          *string           `json:"http_method,omitempty" validate:"omitempty,oneof=GET POST PUT DELETE"`
	HTTPURL             *string           `json:"http_url,omitempty" validate:"omitempty,url"`
	HTTPHeaders         map[string]string `json:"http_headers,omitempty"`
	HTTPBody            *string           `json:"http_body,omitempty"`
	HTTPTimeout         *int              `json:"http_timeout,omitempty" validate:"omitempty,min=1,max=300"`
	MaxRetries          *int              `json:"max_retries,omitempty" validate:"omitempty,min=0,max=10"`
	FailureNotification *bool             `json:"failure_notification,omitempty"`
}

type ScheduleJobResp struct {
	ID                  int       `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Name                string    `json:"name"`
	Type                string    `json:"type"`
	Expression          string    `json:"expression"`
	Description         string    `json:"description"`
	Enabled             bool      `json:"enabled"`
	NextRunTime         time.Time `json:"next_run_time"`
	LastRunTime         time.Time `json:"last_run_time"`
	ExecutionType       string    `json:"execution_type"`
	HTTPMethod          string    `json:"http_method"`
	HTTPURL             string    `json:"http_url"`
	HTTPHeaders         map[string]string `json:"http_headers"`
	HTTPBody            string    `json:"http_body"`
	HTTPTimeout         int       `json:"http_timeout"`
	MaxRetries          int       `json:"max_retries"`
	FailureNotification bool      `json:"failure_notification"`
}
