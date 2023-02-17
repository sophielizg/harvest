package common

import "time"

type RequestFields struct {
	RunId               int    `json:"runId"`
	Url                 string `json:"url"`
	Method              string `json:"method"`
	Blob                []byte `json:"blob"`
	ParentRequestId     *int   `json:"parentRequestId"`
	OriginatorRequestId *int   `json:"originatorRequestId"`
}

type Request struct {
	RequestId        int       `json:"requestId"`
	VisitedTimestamp time.Time `json:"visitedTimestamp"`
	RequestFields
}

type RequestService interface {
	AddRequest(request RequestFields) (int, error)
}
