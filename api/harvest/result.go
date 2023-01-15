package harvest

import "time"

type Result struct {
	ParserId         int       `json:"parserId"`
	ScrapedTimestamp time.Time `json:"scrapedTimestamp"`
	Value            string    `json:"value"`
}

type ResultService interface {
	CrawlResults(crawlId int, tags []string) ([]Result, error)
	RunResults(runId int, tags []string) ([]Result, error)
}
