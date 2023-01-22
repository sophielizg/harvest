package harvest

import "time"

type ResultFields struct {
	RequestId int    `json:"requestId"`
	ParserId  int    `json:"parserId"`
	Value     string `json:"value"`
}

type Result struct {
	ScrapedTimestamp time.Time `json:"scrapedTimestamp"`
	ResultFields
}

type ResultService interface {
	// CrawlResults(crawlId int, tags []string) ([]Result, error)
	// RunResults(runId int, tags []string) ([]Result, error)
	AddResult(crawlId int, scrapeId int, result ResultFields) error
}
