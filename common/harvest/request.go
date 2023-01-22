package harvest

import "time"

type RequestFields struct {
	RequestId          int    `json:"requestId"`
	ScrapeId           int    `json:"scrapeId"`
	RequestHash        string `json:"requestHash"`
	Request            []byte `json:"request"`
	CreatedByRequestId int    `json:"createdByRequestId"`
}

type Request struct {
	VisitedTimestamp time.Time `json:"visitedTimestamp"`
	RequestFields
}

type RequestQueueService interface {
	IsRequestVisited(requestHash string) (bool, error)
	AddRequestIsVisited(request RequestFields) (int, error)
}
