package harvest

import "time"

type RequestToScrape struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	Headers string `json:"headers"`
	Body    string `json:"body"`
}

type QueuedRequestFields struct {
	Request RequestToScrape `json:"request"`
}

type QueuedRequest struct {
	RequestId        int       `json:"requestId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	QueuedRequestFields
}

type RequestQueueService interface {
	Requests(crawlId int) ([]QueuedRequest, error)
	AddRequest(request QueuedRequestFields) (int, error)
	DeleteRequest(requestId int) error
}
