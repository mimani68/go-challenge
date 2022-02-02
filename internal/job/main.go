package job

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunJobs() {
	s := gocron.NewScheduler(time.Local)
	s.Every( /* config.Config.CronJobTime */ 5).Seconds().Do(analyzeData)
	s.StartAsync()
}
