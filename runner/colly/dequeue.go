package colly

func (app *App) Dequeue() error {
	scrape, err := app.RunnerQueueService.DequeueRunner()
	if err != nil {
		return err
	}

	app.ScraperId = scrape.ScraperId
	app.RunId = scrape.RunId
	app.RunnerId = scrape.RunnerId
	return nil
}
