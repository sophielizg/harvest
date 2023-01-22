package harvest

import "time"

type ErrorFields struct {
	RequestId           int    `json:"requestId"`
	ParserId            int    `json:"parserId"`
	StatusCode          int    `json:"statusCode"`
	IsMissngParseResult bool   `json:"isMissngParseResult"`
	ErrorMessage        string `json:"errorMessage"`
}

type Error struct {
	ScrapedTimestamp time.Time `json:"scrapedTimestamp"`
	ErrorFields
}

type ErrorService interface {
	// CrawlErrors(crawlId int, tags []string) ([]Error, error)
	// RunErrors(runId int, tags []string) ([]Error, error)
	AddError(crawlId int, scrapeId int, parseError ErrorFields) error
}
