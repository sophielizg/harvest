package parsers

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
	"github.com/sophielizg/harvest/common/utils"
)

func (p *Parsers) htmlCallback(collector *colly.Collector, parser harvest.Parser) error {
	if parser.Selector == nil {
		return errors.New("The parser.Selector field is required for type html")
	}

	collector.OnHTML(*parser.Selector, func(e *colly.HTMLElement) {
		var val string
		if parser.Attr == nil {
			val = e.Text
		} else {
			val = e.Attr(*parser.Attr)
		}

		p.saveAndEnqueue(e.Response, parser, val)
	})

	return nil
}

func (p *Parsers) xmlCallback(collector *colly.Collector, parser harvest.Parser) error {
	if parser.Xpath == nil {
		return errors.New("The parser.Xpath field is required for type xml")
	}

	collector.OnXML(*parser.Xpath, func(e *colly.XMLElement) {
		val := e.Text
		p.saveAndEnqueue(e.Response, parser, val)
	})

	return nil
}

func (p *Parsers) jsonCallback(collector *colly.Collector, parser harvest.Parser) error {
	if parser.JsonPath == nil {
		return errors.New("The parser.JsonPath field is required for type json")
	}

	collector.OnResponse(func(r *colly.Response) {
		if r.Headers.Get("Content-Type") == "application/json" {
			val, jsonParseErr := utils.GetFromJson(r.Body, parser.JsonPath)
			if jsonParseErr != nil {
				p.saveError(r, parser.ParserId, jsonParseErr, false)
			}
			p.saveAndEnqueue(r, parser, val)
		}
	})

	return nil
}
