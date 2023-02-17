package common

import (
	"time"
)

type ParserAutoIncrementRules struct {
	BodyPath []string `json:"bodyPath"`
	UrlRegex *string  `json:"urlRegex"`
}

type ParserFields struct {
	PageType           *string                   `json:"pageType"`
	Selector           *string                   `json:"selector"`
	Attr               *string                   `json:"attr"`
	Xpath              *string                   `json:"xpath"`
	EnqueueScraperId   *int                      `json:"enqueueScraperId"`
	AutoIncrementRules *ParserAutoIncrementRules `json:"autoIncrementRules"`
}

type Parser struct {
	ParserId         int       `json:"parserId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	Tags             []string  `json:"tags"`
	ParserFields
}

type ParserService interface {
	ParserTypes() ([]string, error)
	Parsers(scraperId int) ([]Parser, error)
	AddParser(scraperId int, parser ParserFields) (int, error)
	DeleteParser(parserId int) error
	AddParserTag(parserId int, tag string) error
	DeleteParserTag(parserId int, tag string) error
}
