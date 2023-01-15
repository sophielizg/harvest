package harvest

import "time"

type RequestToScrape struct {
	url     string
	method  string
	headers string
	body    string
}

type QueuedRequestFields struct {
	request RequestToScrape
}

type QueuedRequest struct {
	requestId        int
	createdTimestamp time.Time
	QueuedRequestFields
}

type RequestQueueService interface {
	Requests(crawlId int) ([]QueuedRequest, error)
	AddRequest(request QueuedRequestFields) (int, error)
	DeleteRequest(requestId int) error
}
