package colly

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/sophielizg/harvest/common/harvest"
)

func (app *App) Collector() (*colly.Collector, error) {
	collector := colly.NewCollector()

	collector.OnRequest(func(request *colly.Request) {
		var err error
		newRequest := harvest.RequestFields{
			RunId: app.RunId,
		}

		newRequest.Blob, err = request.Marshal()
		if err != nil {
			fmt.Fprintf(os.Stderr, "request.Marshal error: %s", err)
		}

		if id, ok := request.Ctx.GetAny("parentRequestId").(int); ok {
			*newRequest.ParentRequestId = id
		}

		if id, ok := request.Ctx.GetAny("originatorRequestId").(int); ok {
			*newRequest.OriginatorRequestId = id
		}

		newRequestId, err := app.RequestService.AddRequest(newRequest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "UpdateRequest error: %s", err)
		}

		request.Ctx.Put("parentRequestId", newRequestId)
		if newRequest.OriginatorRequestId == nil {
			request.Ctx.Put("originatorRequestId", newRequestId)
		}
	})

	app.Configure(collector)
	app.AddParsers(collector)

	storage := &Storage{
		ScraperId:           app.ScraperId,
		RunId:               app.RunId,
		RunnerId:            app.RunnerId,
		RequestService:      app.RequestService,
		RequestQueueService: app.RequestQueueService,
	}

	err := collector.SetStorage(storage)
	if err != nil {
		return nil, err
	}

	app.RequestQueue, err = queue.New(16, storage)
	if err != nil {
		return nil, err
	}

	return collector, nil
}
