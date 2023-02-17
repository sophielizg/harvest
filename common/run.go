package common

import "time"

type Run struct {
	RunId          int       `json:"runId"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
}

type RunService interface {
	// IsRunning(runId int) (bool, error)
	// Runs(scraperId int) ([]Run, error)
	// DeleteRun(runId int) error
	CreateRun(scraperId int) (int, error)
	// StopRun(runId int) error
	// PauseRun(runId int) error
	// UnpauseRun(runId int) error
	// StopCurrentRun(scraperId int) error
	// PauseCurrentRun(scraperId int) error
	// UnpauseCurrentRun(scraperId int) error
}
