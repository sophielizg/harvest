package parsers

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/antchfx/jsonquery"
	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
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

		p.saveAndEnqueue(e.Response, parser, val, &e.Index)
	})

	return nil
}

func (p *Parsers) xmlCallback(collector *colly.Collector, parser harvest.Parser) error {
	if parser.Xpath == nil {
		return errors.New("The parser.Xpath field is required for type xml")
	}

	collector.OnXML(*parser.Xpath, func(e *colly.XMLElement) {
		val := e.Text
		p.saveAndEnqueue(e.Response, parser, val, nil)
	})

	return nil
}

func (p *Parsers) jsonCallback(collector *colly.Collector, parser harvest.Parser) error {
	if parser.Xpath == nil {
		return errors.New("The parser.Xpath field is required for type json")
	}

	collector.OnResponse(func(r *colly.Response) {
		if r.Headers.Get("Content-Type") == "application/json" {
			doc, err := jsonquery.Parse(bytes.NewReader(r.Body))
			if err != nil {
				p.saveError(r, parser.ParserId, err, nil, false)
				return
			}

			vals, err := jsonquery.QueryAll(doc, *parser.Xpath)
			if err != nil {
				p.saveError(r, parser.ParserId, err, nil, false)
				return
			}

			for i, val := range vals {
				valStr := fmt.Sprintf("%v", val.Value())
				p.saveAndEnqueue(r, parser, valStr, &i)
			}
		}
	})

	return nil
}
