package job

import (
	"context"
	"time"

	payorder_service "github.com/shuTwT/hoshikuzu/internal/services/mall/payorder"
	schedule_model "github.com/shuTwT/hoshikuzu/pkg/domain/model/schedule"
)

// CloseTimeoutOrdersJob 定时关闭超时未支付订单。
type CloseTimeoutOrdersJob struct {
	PayOrderService payorder_service.PayOrderService
}

func (job CloseTimeoutOrdersJob) Execute(ctx context.Context) error {
	return job.PayOrderService.CloseTimeoutOrders(ctx)
}

func (CloseTimeoutOrdersJob) Type() schedule_model.JobType {
	return schedule_model.DurationJobType
}

func (CloseTimeoutOrdersJob) Duration() time.Duration {
	return 5 * time.Minute
}

func (CloseTimeoutOrdersJob) Description() string {
	return "关闭超时未支付订单"
}
