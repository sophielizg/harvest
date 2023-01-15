package harvest

import "time"

type Run struct {
	RunId          int       `json:"runId"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
}

type RunService interface {
	Runs(crawlId int) ([]Run, error)
	DeleteRun(runId int) error
}
