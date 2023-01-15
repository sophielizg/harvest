package harvest

import "time"

type Error struct {
	ParserId            int       `json:"parserId"`
	ScrapedTimestamp    time.Time `json:"scrapedTimestamp"`
	StatusCode          int       `json:"statusCode"`
	IsMissngParseResult bool      `json:"isMissngParseResult"`
	Response            string    `json:"response"`
}

type ErrorService interface {
	CrawlErrors(crawlId int, tags []string) ([]Error, error)
	RunErrors(runId int, tags []string) ([]Error, error)
}
