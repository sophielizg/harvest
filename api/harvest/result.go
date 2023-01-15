package harvest

import "time"

type Result struct {
	parserId         int
	scrapedTimestamp time.Time
	value            string
}

type ResultService interface {
	CrawlResults(crawlId int, tags []string) ([]Result, error)
	RunResults(runId int, tags []string) ([]Result, error)
}
