package schedule

import (
	"context"
	"fmt"

	"github.com/shuTwT/gobee/ent"
	"github.com/shuTwT/gobee/ent/schedulejob"
	"github.com/shuTwT/gobee/internal/infra/schedule/manager"
	friendcircle_job "github.com/shuTwT/gobee/internal/job/friendcircle"
	friend_circle_service "github.com/shuTwT/gobee/internal/services/content/friendcircle"
)

func InitializeSchedule(db *ent.Client, scheduleManager *manager.ScheduleManager, friendCircleService friend_circle_service.FriendCircleService) error {

	scheduleManager.AddJobToCache("friendCircle", friendcircle_job.FriendCircleJob{
		DbClient:            db,
		FriendCircleService: friendCircleService,
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
