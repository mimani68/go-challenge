package job

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunJobs(interval int) {
	s := gocron.NewScheduler(time.Local)
	s.Every(interval).Seconds().Do(analyzeData)
	s.StartAsync()
}
