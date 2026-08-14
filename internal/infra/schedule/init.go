package schedule

import (
	"context"
	"fmt"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/schedulejob"
	"github.com/shuTwT/hoshikuzu/internal/infra/schedule/manager"
	friendcircle_job "github.com/shuTwT/hoshikuzu/internal/job/friendcircle"
	payorder_job "github.com/shuTwT/hoshikuzu/internal/job/payorder"
	friend_circle_service "github.com/shuTwT/hoshikuzu/internal/services/content/friendcircle"
	payorder_service "github.com/shuTwT/hoshikuzu/internal/services/mall/payorder"
)

func InitializeSchedule(db *ent.Client, scheduleManager *manager.ScheduleManager, friendCircleService friend_circle_service.FriendCircleService, payOrderService payorder_service.PayOrderService) error {

	scheduleManager.AddJobToCache("friendCircle", friendcircle_job.FriendCircleJob{
		DbClient:            db,
		FriendCircleService: friendCircleService,
	})

	scheduleManager.AddJobToCache("closeTimeoutOrders", payorder_job.CloseTimeoutOrdersJob{
		PayOrderService: payOrderService,
	})

	jobs, err := db.ScheduleJob.Query().
		Where(schedulejob.Enabled(true)).
		All(context.Background())
	if err != nil {
		return fmt.Errorf("查询定时任务失败: %w", err)
	}

	for _, jobEntity := range jobs {
		err := scheduleManager.AddJobToScheduler(jobEntity)
		if err != nil {
			return fmt.Errorf("添加定时任务失败: %w", err)
		}
	}

	scheduleManager.Start()

	return nil
}
