package colly

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/sophielizg/harvest/runner/colly/parsers"
	"github.com/sophielizg/harvest/runner/colly/storage"
)

func (r *Runner) collector() (*colly.Collector, *queue.Queue, error) {
	collector := colly.NewCollector()

	r.configure(collector)
	r.addCallbacks(collector)

	var storage storage.Storage
	storage.SharedFields = r.SharedFields
	storage.StorageServices = r.StorageServices

	err := collector.SetStorage(&storage)
	if err != nil {
		return nil, nil, err
	}

	requestQueue, err := queue.New(16, &storage)
	if err != nil {
		return nil, nil, err
	}

	parsers := &parsers.Parsers{
		SharedFields:    r.SharedFields,
		Queue:           requestQueue,
		ParsersServices: r.ParsersServices,
	}
	err = parsers.Add(collector)
	return collector, requestQueue, err
}
