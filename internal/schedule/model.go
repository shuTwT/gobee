package schedule

import "context"

type Job interface {
	Execute(context.Context) error
	Description() string
}
