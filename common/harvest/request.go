package harvest

import "time"

type RequestFields struct {
	Blob                []byte `json:"blob"`
	ParentRequestId     int    `json:"parentRequestId"`
	OriginatorRequestId int    `json:"originatorRequestId"`
}

type Request struct {
	RunId            int       `json:"runId"`
	RequestId        uint64    `json:"requestId"`
	VisitedTimestamp time.Time `json:"visitedTimestamp"`
	RequestFields
}

type RequestService interface {
	IsRequestVisited(runId int, requestId uint64) (bool, error)
	AddRequestIsVisited(runId int, requestId uint64) (int, error)
	UpdateRequest(runId int, requestId uint64, request RequestFields) (int, error)
}
