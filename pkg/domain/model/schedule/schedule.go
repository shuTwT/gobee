package schedule

import (
	"context"
	"time"
)

type JobType int

const (
	DurationJobType JobType = iota
	CronJobType
)

type Job interface {
	Execute(context.Context) error
	Type() JobType
	Description() string
}

type DurationJob interface {
	Execute(context.Context) error
	Duration() time.Duration
	Type() JobType
	Description() string
}

type CronJob interface {
	Execute(context.Context) error
	// Crontab() string
	Type() JobType
	Description() string
}
