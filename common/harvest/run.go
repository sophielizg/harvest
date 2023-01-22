package harvest

import "time"

type Run struct {
	CrawlRunId     int       `json:"crawlRunId"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
}

type RunService interface {
	IsRunning(crawlRunId int) (bool, error)
	// Runs(crawlId int) ([]Run, error)
	// DeleteRun(runId int) error
}
