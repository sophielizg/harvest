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
	RequestQueueId   int       `json:"requestQueueId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	QueuedRequestFields
}

type RequestQueueService interface {
	InitialRequests(crawlId int) ([]QueuedRequest, error)
	AddInitialRequest(crawlId int, request QueuedRequestFields) (int, error)
	DeleteInitialRequest(requestQueueId int) error
	EnqueueRequest(crawlId int, request QueuedRequestFields) (int, error)
	DequeueRequest(crawlId int) (QueuedRequest, error)
}
