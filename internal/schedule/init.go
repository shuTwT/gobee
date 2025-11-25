package schedule

var jobCache map[string]Job

func ClearCache() {
	for key := range jobCache {
		delete(jobCache, key)
	}
}

// 初始化调度
func InitializeSchedule() error {
	jobCache = map[string]Job{}
	return nil
}
