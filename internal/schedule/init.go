package schedule

import (
	"context"
	"fmt"
	friendcircle_job "gobee/internal/job/friendcircle"
	schedule_model "gobee/pkg/domain/model/schedule"
	"log"

	"github.com/go-co-op/gocron/v2"
)

var jobCache map[string]schedule_model.Job

func ClearCache() {
	for key := range jobCache {
		delete(jobCache, key)
	}
}

// 初始化调度
func InitializeSchedule() (gocron.Scheduler, error) {
	jobCache = map[string]schedule_model.Job{
		"friendCircle": friendcircle_job.FriendCircleJob{},
	}
	// 创建调度器
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}
	for _, task := range jobCache {
		var taskJob gocron.Job
		switch t := task.(type) {
		case schedule_model.DurationJob:
			taskJob, err = scheduler.NewJob(
				gocron.DurationJob(
					t.Duration(),
				),
				gocron.NewTask(
					func(a string, b int) {
						_ = t.Execute(context.TODO())
					},
					"hello",
					1,
				),
			)

		case schedule_model.CronJob:
			taskJob, err = scheduler.NewJob(
				gocron.CronJob(
					t.Crontab(),
					true,
				),
				gocron.NewTask(
					func(a string, b int) {
						t.Execute(context.TODO())
					},
					"hello",
					1,
				),
			)

		default:
			fmt.Println("位置 Job类型")
			continue
		}

		if err != nil {
			return nil, err
		}
		log.Println("已添加任务:", taskJob.ID())
	}

	scheduler.Start()

	return scheduler, nil
}
