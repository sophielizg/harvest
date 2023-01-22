package colly

import (
	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
)

type App struct {
	CrawlId       int
	CrawlRunId    int
	ScrapeId      int
	CrawlService  harvest.CrawlService
	ScrapeService harvest.ScrapeService
	ParserService harvest.ParserService
}

func (app *App) Crawler() (*colly.Collector, error) {
	err := app.Dequeue()
	if err != nil {
		return nil, err
	}

	collector, err := app.Collector()
	if err != nil {
		return nil, err
	}

	app.AddParsers(collector)

	return collector, nil
}
