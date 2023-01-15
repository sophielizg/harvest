package harvest

import (
	"time"
)

type ParserAutoIncrementRules struct {
	bodyPath []string
	urlRegex string
}

type ParserTag struct {
	parserTagId      int
	createdTimestamp time.Time
	name             string
}

type ParserFields struct {
	pageType           string
	selector           string
	attr               string
	xpath              string
	jsonPath           []string
	enqueueCrawlId     int
	autoIncrementRules ParserAutoIncrementRules
}

type Parser struct {
	parserId         int
	createdTimestamp time.Time
	tags             []ParserTag
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
