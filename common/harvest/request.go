package harvest

import "time"

type RequestFields struct {
	ScrapeId           int    `json:"scrapeId"`
	Id                 uint64 `json:"id"`
	Blob               []byte `json:"blob"`
	CreatedByRequestId int    `json:"createdByRequestId"`
}

type Request struct {
	VisitedTimestamp time.Time `json:"visitedTimestamp"`
	RequestFields
}

type RequestService interface {
	IsRequestVisited(crawlRunId int, requestId uint64) (bool, error)
	AddRequestIsVisited(request RequestFields) (int, error)
}
