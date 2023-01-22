package harvest

import "time"

type RequestToScrape struct {
	Url     string              `json:"url"`
	Method  string              `json:"method"`
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
}

type QueuedRequestFields struct {
	CrawlId            int             `json:"crawlId"`
	Request            RequestToScrape `json:"request"`
	IsInitialRequest   bool            `json:"isInitialRequest"`
	CreatedByRequestId int             `json:"createdByRequestId"`
}

type QueuedRequest struct {
	RequestQueueId   int       `json:"requestQueueId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	ScrapeId         int       `json:"scrapeId"`
	QueuedRequestFields
}

type RequestQueueService interface {
	QueuedRequests(crawlId int) ([]QueuedRequest, error)
	DeleteQueuedRequest(requestQueueId int) error
	EnqueueRequest(request QueuedRequestFields) (int, error)
	DequeueRequest(crawlId int) (QueuedRequest, error)
}
