package harvest

import "time"

type ErrorFields struct {
	RunId               int    `json:"runId"`
	RequestId           int    `json:"requestId"`
	ParserId            int    `json:"parserId"`
	StatusCode          int    `json:"statusCode"`
	Response            []byte `json:"response"`
	IsMissngParseResult bool   `json:"isMissngParseResult"`
	ErrorMessage        string `json:"errorMessage"`
}

type Error struct {
	ScrapedTimestamp time.Time `json:"scrapedTimestamp"`
	ErrorFields
}

type ErrorService interface {
	// ScraperErrors(scraperId int, tags []string) ([]Error, error)
	// RunErrors(runId int, tags []string) ([]Error, error)
	AddError(runnerId int, parseError ErrorFields) error
}
