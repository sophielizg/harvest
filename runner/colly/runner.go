package colly

import (
	"github.com/sophielizg/harvest/common/harvest"
	"github.com/sophielizg/harvest/runner/colly/common"
	"github.com/sophielizg/harvest/runner/colly/parsers"
	"github.com/sophielizg/harvest/runner/colly/storage"
)

type Runner struct {
	common.RunnerIds
	ScraperService     harvest.ScraperService
	RunnerQueueService harvest.RunnerQueueService
	RequestService     harvest.RequestService
	storage.StorageServices
	parsers.ParsersServices
}

func (r *Runner) Run() error {
	err := r.Dequeue()
	if err != nil {
		return err
	}

	collector, queue, err := r.Collector()
	if err != nil {
		return err
	}

	return queue.Run(collector)
}
