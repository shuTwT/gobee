package model

import "time"

type CreateScheduleJobReq struct {
	Name                string  `json:"name" validate:"required"`
	Type                string  `json:"type" validate:"required,oneof=cron interval"`
	Expression          string  `json:"expression" validate:"required"`
	Description         *string `json:"description,omitempty"`
	Enabled             bool    `json:"enabled"`
	JobName             string  `json:"job_name" validate:"required"`
	MaxRetries          int     `json:"max_retries" validate:"min=0,max=10"`
	FailureNotification bool    `json:"failure_notification"`
}

type UpdateScheduleJobReq struct {
	Name                *string `json:"name,omitempty"`
	Type                *string `json:"type,omitempty" validate:"omitempty,oneof=cron interval"`
	Expression          *string `json:"expression,omitempty"`
	Description         *string `json:"description,omitempty"`
	Enabled             *bool   `json:"enabled,omitempty"`
	JobName             *string `json:"job_name,omitempty"`
	MaxRetries          *int    `json:"max_retries,omitempty" validate:"omitempty,min=0,max=10"`
	FailureNotification *bool   `json:"failure_notification,omitempty"`
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
	JobName             string    `json:"job_name"`
	MaxRetries          int       `json:"max_retries"`
	FailureNotification bool      `json:"failure_notification"`
}
