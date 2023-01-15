package harvest

import "time"

type Status struct {
	Queued               int       `json:"queued"`
	Successes            int       `json:"successes"`
	Errors               int       `json:"errors"`
	Missing              int       `json:"missing"`
	LastUpdatedTimestamp time.Time `json:"lastUpdatedTimestamp"`
}

type StatusService interface {
	CrawlStatus(crawlId int) (Status, error)
	RunStatus(runId int) (Status, error)
}
