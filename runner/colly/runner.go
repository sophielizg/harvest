package colly

import (
	"github.com/sophielizg/harvest/common"
	collyCommon "github.com/sophielizg/harvest/runner/colly/common"
	"github.com/sophielizg/harvest/runner/colly/parsers"
	"github.com/sophielizg/harvest/runner/colly/storage"
)

type Runner struct {
	collyCommon.SharedFields
	ScraperService     common.ScraperService
	RunnerQueueService common.RunnerQueueService
	RequestService     common.RequestService
	storage.StorageServices
	parsers.ParsersServices
}

func (r *Runner) dequeue() error {
	runner, err := r.RunnerQueueService.DequeueRunner()
	if err != nil {
		return err
	}

	r.ScraperId = runner.ScraperId
	r.RunId = runner.RunId
	r.RunnerId = runner.RunnerId
	return nil
}

func (r *Runner) end() {
	err := r.RunnerQueueService.EndRunner(r.RunnerId)
	if err != nil {
		r.Logger.WithFields(common.LogFields{
			"error": err,
			"ids":   r.SharedIds,
		}).Warn("An error ocurred in EndRunner while ending the runner")
	}
}

func (r *Runner) Run() error {
	err := r.dequeue()
	if err != nil {
		return err
	}
	defer r.end()

	collector, queue, err := r.collector()
	if err != nil {
		return err
	}

	return queue.Run(collector)
}
