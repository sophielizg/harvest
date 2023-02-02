package colly

func (r *Runner) Dequeue() error {
	runner, err := r.RunnerQueueService.DequeueRunner()
	if err != nil {
		return err
	}

	r.ScraperId = runner.ScraperId
	r.RunId = runner.RunId
	r.RunnerId = runner.RunnerId
	return nil
}
