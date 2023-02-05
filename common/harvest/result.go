package harvest

import "time"

type ResultFields struct {
	RunId     int    `json:"runId"`
	RequestId int    `json:"requestId"`
	ParserId  int    `json:"parserId"`
	Value     string `json:"value"`
}

type Result struct {
	ResultId         int       `json:"resultId"`
	ScrapedTimestamp time.Time `json:"scrapedTimestamp"`
	ResultFields
}

type ResultService interface {
	// ScraperResults(scraperId int, tags []string) ([]Result, error)
	// RunResults(runId int, tags []string) ([]Result, error)
	AddResult(runnerId int, result ResultFields) (int, error)
}
