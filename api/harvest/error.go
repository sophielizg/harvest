package harvest

import "time"

type Error struct {
	parserId            int
	scrapedTimestamp    time.Time
	statusCode          int
	isMissngParseResult bool
	response            string
}

type ErrorService interface {
	CrawlErrors(crawlId int, tags []string) ([]Error, error)
	RunErrors(runId int, tags []string) ([]Error, error)
}
