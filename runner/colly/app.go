package colly

import (
	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
)

type App struct {
	ScraperId           int
	RunId               int
	RunnerId            int
	ScraperService      harvest.ScraperService
	RunnerQueueService  harvest.RunnerQueueService
	ParserService       harvest.ParserService
	ResultService       harvest.ResultService
	ErrorService        harvest.ErrorService
	RequestQueueService harvest.RequestQueueService
}

func (app *App) Scraperer() (*colly.Collector, error) {
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
