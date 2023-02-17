package colly

import (
	harvest "github.com/sophielizg/harvest/common"
	"github.com/sophielizg/harvest/runner/colly/common"
	"github.com/sophielizg/harvest/runner/colly/parsers"
	"github.com/sophielizg/harvest/runner/colly/storage"
)

type Runner struct {
	common.SharedFields
	ScraperService     harvest.ScraperService
	RunnerQueueService harvest.RunnerQueueService
	RequestService     harvest.RequestService
	storage.StorageServices
	parsers.ParsersServices
}

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

func (r *Runner) End() {
	err := r.RunnerQueueService.EndRunner(r.RunnerId)
	if err != nil {
		r.Logger.WithFields(harvest.LogFields{
			"error": err,
			"ids":   r.SharedIds,
		}).Warn("An error ocurred in EndRunner while ending the runner")
	}
}

func (r *Runner) Run() error {
	err := r.Dequeue()
	if err != nil {
		return err
	}
	defer r.End()

	collector, queue, err := r.Collector()
	if err != nil {
		return err
	}

	return queue.Run(collector)
}
