package colly

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/sophielizg/harvest/common/harvest"
)

type App struct {
	ScraperId           int
	RunId               int
	RunnerId            int
	RequestQueue        *queue.Queue
	ScraperService      harvest.ScraperService
	RunnerQueueService  harvest.RunnerQueueService
	ParserService       harvest.ParserService
	ResultService       harvest.ResultService
	ErrorService        harvest.ErrorService
	RequestService      harvest.RequestService
	RequestQueueService harvest.RequestQueueService
}

func (app *App) Scraper() (*colly.Collector, error) {
	err := app.DequeueRunner()
	if err != nil {
		return nil, err
	}

	return app.Collector()
}
