package schedule

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shuTwT/gobee/ent"
	"github.com/shuTwT/gobee/ent/schedulejob"
	friendcircle_job "github.com/shuTwT/gobee/internal/job/friendcircle"
	"github.com/shuTwT/gobee/pkg"
	schedule_model "github.com/shuTwT/gobee/pkg/domain/model/schedule"

	"github.com/go-co-op/gocron/v2"
)

var jobCache map[string]schedule_model.Job

func ClearCache() {
	for key := range jobCache {
		delete(jobCache, key)
	}
}

func GetJob(jobName string) (schedule_model.Job, bool) {
	job, ok := jobCache[jobName]
	return job, ok
}

func InitializeSchedule(db *ent.Client, serviceMap pkg.ServiceMap) (gocron.Scheduler, error) {
	jobCache = map[string]schedule_model.Job{
		"friendCircle": friendcircle_job.FriendCircleJob{
			DbClient:            db,
			FriendCircleService: serviceMap.FriendCircleService,
		},
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	jobs, err := db.ScheduleJob.Query().
		Where(schedulejob.Enabled(true)).
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("查询定时任务失败: %w", err)
	}

	for _, jobEntity := range jobs {
		job, ok := jobCache[jobEntity.JobName]
		if !ok {
			log.Printf("警告: 找不到任务 '%s' 的实现", jobEntity.JobName)
			continue
		}

		var taskJob gocron.Job
		switch jobEntity.Type {
		case "interval":
			durationJob, ok := job.(schedule_model.DurationJob)
			if !ok {
				log.Printf("警告: 任务 '%s' 不是 DurationJob 类型", jobEntity.JobName)
				continue
			}

			duration, err := time.ParseDuration(jobEntity.Expression)
			if err != nil {
				log.Printf("警告: 任务 '%s' 的表达式 '%s' 解析失败: %v", jobEntity.JobName, jobEntity.Expression, err)
				continue
			}

			taskJob, err = scheduler.NewJob(
				gocron.DurationJob(duration),
				gocron.NewTask(
					func(ctx context.Context, j schedule_model.DurationJob) {
						if err := j.Execute(ctx); err != nil {
							log.Printf("任务 '%s' 执行失败: %v", jobEntity.JobName, err)
						}
					},
					context.Background(),
					durationJob,
				),
			)

		case "cron":
			cronJob, ok := job.(schedule_model.CronJob)
			if !ok {
				log.Printf("警告: 任务 '%s' 不是 CronJob 类型", jobEntity.JobName)
				continue
			}

			taskJob, err = scheduler.NewJob(
				gocron.CronJob(jobEntity.Expression, true),
				gocron.NewTask(
					func(ctx context.Context, j schedule_model.CronJob) {
						if err := j.Execute(ctx); err != nil {
							log.Printf("任务 '%s' 执行失败: %v", jobEntity.JobName, err)
						}
					},
					context.Background(),
					cronJob,
				),
			)

		default:
			log.Printf("警告: 任务 '%s' 的类型 '%s' 不支持", jobEntity.JobName, jobEntity.Type)
			continue
		}

		if err != nil {
			log.Printf("添加任务 '%s' 失败: %v", jobEntity.JobName, err)
			continue
		}

		log.Printf("已添加任务: %s (ID: %s, 类型: %s, 表达式: %s)", jobEntity.Name, taskJob.ID(), jobEntity.Type, jobEntity.Expression)
	}

	scheduler.Start()

	return scheduler, nil
}
