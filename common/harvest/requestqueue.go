package harvest

import "time"

type QueuedRequestFields struct {
	RunId int    `json:"runId"`
	Blob  []byte `json:"blob"`
}

type QueuedRequest struct {
	RequestQueueId   int       `json:"requestQueueId"`
	RunnerId         int       `json:"runnerId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	IsInitialRequest bool      `json:"isInitialRequest"`
	QueuedRequestFields
}

type RequestQueueService interface {
	// QueuedRequests(runId int) ([]QueuedRequest, error)
	AddStartingRequest(scraperId int, requestBlob []byte) (int, error)
	// DeleteQueuedRequest(requestQueueId int) error
	EnqueueRequest(scraperId int, request QueuedRequestFields) (int, error)
	DequeueRequests(runId int, runnerId int, numToDequeue int) ([]QueuedRequest, error)
}
