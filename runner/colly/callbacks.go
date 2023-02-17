package colly

import (
	"github.com/gocolly/colly"
	harvest "github.com/sophielizg/harvest/common"
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
		r.Logger.WithFields(harvest.LogFields{
			"error":   err,
			"ids":     r.SharedIds,
			"request": request,
		}).Warn("An error ocurred in request.Marshal while making request")
	}

	if id, ok := request.Ctx.GetAny("parentRequestId").(int); ok {
		*newRequest.ParentRequestId = id
	}

	if id, ok := request.Ctx.GetAny("originatorRequestId").(int); ok {
		*newRequest.OriginatorRequestId = id
	}

	newRequestId, err := r.RequestService.AddRequest(newRequest)
	if err != nil {
		r.Logger.WithFields(harvest.LogFields{
			"error":   err,
			"ids":     r.SharedIds,
			"request": newRequest,
		}).Error("An error ocurred in AddRequest while making request")
		request.Abort()
	}

	request.Ctx.Put("requestId", newRequestId)
}

func (r *Runner) AddCallbacks(collector *colly.Collector) {
	collector.OnRequest(r.trackRequest)
}
