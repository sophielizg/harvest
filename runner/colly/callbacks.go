package colly

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
)

func (r *Runner) trackRequest(request *colly.Request) {
	var err error
	newRequest := harvest.RequestFields{
		RunId:  r.RunId,
		Url:    request.URL.String(),
		Method: request.Method,
	}

	newRequest.Blob, err = request.Marshal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "request.Marshal error: %s\n", err)
	}

	if id, ok := request.Ctx.GetAny("parentRequestId").(int); ok {
		*newRequest.ParentRequestId = id
	}

	if id, ok := request.Ctx.GetAny("originatorRequestId").(int); ok {
		*newRequest.OriginatorRequestId = id
	}

	newRequestId, err := r.RequestService.AddRequest(newRequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "UpdateRequest error: %s\n", err)
	}

	request.Ctx.Put("requestId", newRequestId)
}

func (r *Runner) AddCallbacks(collector *colly.Collector) {
	collector.OnRequest(r.trackRequest)
}
