package harvest

import "time"

type QueuedRequestFields struct {
	ScraperId int    `json:"scraperId"`
	RunId     int    `json:"runId"`
	RunnerId  int    `json:"runnerId"`
	Blob      []byte `json:"blob"`
}

type QueuedRequest struct {
	RequestQueueId   int       `json:"requestQueueId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	IsInitialRequest bool      `json:"isInitialRequest"`
	QueuedRequestFields
}

type RequestQueueService interface {
	GetQueueSize(runId int) (int, error)
	// QueuedRequests(runId int) ([]QueuedRequest, error)
	AddStartingRequest(scraperId int, requestBlob []byte) (int, error)
	// DeleteQueuedRequest(requestQueueId int) error
	EnqueueRequest(request QueuedRequestFields) (int, error)
	DequeueRequests(runId int, runnerId int, numToDequeue int) ([][]byte, error)
}
