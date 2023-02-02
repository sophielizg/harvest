package colly

func (app *App) DequeueRunner() error {
	runner, err := app.RunnerQueueService.DequeueRunner()
	if err != nil {
		return err
	}

	app.ScraperId = runner.ScraperId
	app.RunId = runner.RunId
	app.RunnerId = runner.RunnerId
	return nil
}
