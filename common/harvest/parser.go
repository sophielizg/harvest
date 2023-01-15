package harvest

import (
	"time"
)

type ParserAutoIncrementRules struct {
	BodyPath []string `json:"bodyPath"`
	UrlRegex string   `json:"urlRegex"`
}

type ParserTag struct {
	ParserTagId      int       `json:"parserTagId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	Name             string    `json:"name"`
}

type ParserFields struct {
	PageType           string                   `json:"pageType"`
	Selector           string                   `json:"selector"`
	Attr               string                   `json:"attr"`
	Xpath              string                   `json:"xpath"`
	JsonPath           []string                 `json:"jsonPath"`
	EnqueueCrawlId     int                      `json:"enqueueCrawlId"`
	AutoIncrementRules ParserAutoIncrementRules `json:"autoIncrementRules"`
}

type Parser struct {
	ParserId         int         `json:"parserId"`
	CreatedTimestamp time.Time   `json:"createdTimestamp"`
	Tags             []ParserTag `json:"tags"`
	ParserFields
}

type ParserService interface {
	ParserTypes() ([]string, error)
	Parsers(crawlId int) ([]Parser, error)
	AddParser(parser ParserFields) (int, error)
	DeleteParser(parserId int) error
	AddParserTag(parserId int, tag string) (int, error)
	DeleteParserTag(parserTagId int) error
}
