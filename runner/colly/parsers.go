package colly

import (
	"bytes"
	"errors"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
	"github.com/sophielizg/harvest/runner/utils"
)

type parserFunc func(*colly.Collector, harvest.Parser) error

func (app *App) enqueueRequest(parentRequest *colly.Request, parser harvest.Parser,
	parsedValue string) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(parentRequest.Body)
	body := buf.Bytes()
	url := parentRequest.URL.String()

	var err error
	newUrl := url
	newBody := body

	if parser.AutoIncrementRules != nil {
		if parser.AutoIncrementRules.UrlRegex != nil {
			newUrl, err = utils.IncrementUrl(url, *parser.AutoIncrementRules.UrlRegex)
		}
		if err == nil && parser.AutoIncrementRules.BodyPath != nil {
			newBody, err = utils.IncrementBody(body, parser.AutoIncrementRules.BodyPath)
		}
		if err != nil {
			return err
		}
	} else {
		newUrl = parsedValue
		newBody = nil
	}

	// TODO Enqueue request with newUrl and newBody
}

func (app *App) saveResult(response *colly.Response, parserId int, parsedValue string) error {
	// TODO: implement this
	// get app.CrawlId, app.ScrapeId, ctx.Get("requestId")

	// add result in result service
}

func (app *App) saveError(response *colly.Response, parserId int, err error,
	isMissingResult bool) error {
	// TODO: implement this
	// get app.CrawlId, app.ScrapeId, ctx.Get("requestId")

	// add error in error service
}

func (app *App) saveAndEnqueue(response *colly.Response, parser harvest.Parser,
	parsedValue string) {
	if parsedValue == "" {
		err := errors.New("No result found from parser")
		app.saveError(response, parser.ParserId, err, true)
	} else {
		app.saveResult(response, parser.ParserId, parsedValue)

		if parser.EnqueueCrawlId != nil {
			err := app.enqueueRequest(response.Request, parser, parsedValue)
			if err != nil {
				app.saveError(response, parser.ParserId, err, false)
			}
		}
	}
}

func (app *App) htmlParser(collector *colly.Collector, parser harvest.Parser) error {
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

		app.saveAndEnqueue(e.Response, parser, val)
	})

	return nil
}

func (app *App) xmlParser(collector *colly.Collector, parser harvest.Parser) error {
	if parser.Xpath == nil {
		return errors.New("The parser.Xpath field is required for type xml")
	}

	collector.OnXML(*parser.Xpath, func(e *colly.XMLElement) {
		val := e.Text
		app.saveAndEnqueue(e.Response, parser, val)
	})

	return nil
}

func (app *App) jsonParser(collector *colly.Collector, parser harvest.Parser) error {
	if parser.JsonPath == nil {
		return errors.New("The parser.JsonPath field is required for type json")
	}

	collector.OnResponse(func(r *colly.Response) {
		if r.Headers.Get("Content-Type") == "application/json" {
			val, err := utils.GetFromJson(r.Body, parser.JsonPath)
			if err != nil {
				app.saveError(r, parser.ParserId, err, false)
			}
			app.saveAndEnqueue(r, parser, val)
		}
	})

	return nil
}

func (app *App) AddParsers(collector *colly.Collector) error {
	parsers, err := app.ParserService.Parsers(app.CrawlId)
	if err != nil {
		return err
	}

	parserFuncByType := map[string]parserFunc{
		"html": app.htmlParser,
		"xml":  app.xmlParser,
		"json": app.jsonParser,
	}

	var parserIds []int
	for _, parser := range parsers {
		f := parserFuncByType[*parser.PageType]
		if err = f(collector, parser); err != nil {
			return err
		}

		parserIds = append(parserIds, parser.ParserId)
	}

	// Add error handler
	collector.OnError(func(r *colly.Response, err error) {
		for _, parserId := range parserIds {
			app.saveError(r, parserId, err, false)
		}
	})

	return nil
}
