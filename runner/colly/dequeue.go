package colly

func (app *App) Dequeue() error {
	scrape, err := app.ScrapeService.DequeueScrape()
	if err != nil {
		return err
	}

	app.CrawlId = scrape.CrawlId
	app.CrawlRunId = scrape.CrawlRunId
	app.ScrapeId = scrape.ScrapeId
	return nil
}
