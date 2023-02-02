package colly

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/sophielizg/harvest/runner/colly/parsers"
	"github.com/sophielizg/harvest/runner/colly/storage"
)

func (r *Runner) Collector() (*colly.Collector, *queue.Queue, error) {
	collector := colly.NewCollector()

	r.Configure(collector)
	r.AddCallbacks(collector)

	storage := &storage.Storage{
		r.RunnerIds,
		r.StorageServices,
	}

	err := collector.SetStorage(storage)
	if err != nil {
		return nil, nil, err
	}

	requestQueue, err := queue.New(16, storage)
	if err != nil {
		return nil, nil, err
	}

	parsers := &parsers.Parsers{
		r.RunnerIds,
		requestQueue,
		r.ParsersServices,
	}
	err = parsers.Add(collector)
	return collector, requestQueue, err
}
