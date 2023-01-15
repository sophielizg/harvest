package harvest

import "time"

type Run struct {
	runId          int
	startTimestamp time.Time
	endTimestamp   time.Time
}

type RunService interface {
	Runs(crawlId int) ([]Run, error)
	DeleteRun(runId int) error
}
