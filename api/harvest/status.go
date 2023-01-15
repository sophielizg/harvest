package harvest

import "time"

type Status struct {
	queued               int
	successes            int
	errors               int
	missing              int
	lastUpdatedTimestamp time.Time
}

type StatusService interface {
	CrawlStatus(crawlId int) (Status, error)
	RunStatus(runId int) (Status, error)
}
