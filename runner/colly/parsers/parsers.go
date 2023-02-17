package parsers

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/sophielizg/harvest/common"
	collyCommon "github.com/sophielizg/harvest/runner/colly/common"
)

type ParsersServices struct {
	ParserService common.ParserService
	ErrorService  common.ErrorService
	ResultService common.ResultService
}

type Parsers struct {
	collyCommon.SharedFields
	Queue *queue.Queue
	ParsersServices
}

func (p *Parsers) Add(collector *colly.Collector) error {
	parsers, err := p.ParserService.Parsers(p.ScraperId)
	if err != nil {
		return err
	}

	callbackByType := map[string]func(*colly.Collector, common.Parser) error{
		"html": p.htmlCallback,
		"xml":  p.xmlCallback,
		"json": p.jsonCallback,
	}

	var parserIds []int
	for _, parser := range parsers {
		f := callbackByType[*parser.PageType]
		if err = f(collector, parser); err != nil {
			return err
		}

		parserIds = append(parserIds, parser.ParserId)
	}

	// Add error handler
	collector.OnError(func(r *colly.Response, err error) {
		for _, parserId := range parserIds {
			p.saveError(r, parserId, err, nil, false)
		}
	})

	return nil
}
