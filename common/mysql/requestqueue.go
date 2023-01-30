package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type RequestQueueService struct {
	Db *sql.DB
}

func (q *RequestQueueService) enqueueRequest(scraperId int,
	request harvest.QueuedRequestFields, isInitialRequest bool) (int, error) {
	rows, err := q.Db.Query("CALL enqueueRequest(?, ?, ?, ?, ?, 1);", scraperId,
		request.RunId, request.RunnerId, request.Blob, isInitialRequest)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var requestQueueId int
		err = rows.Scan(&requestQueueId)
		if err != nil {
			return 0, err
		}
		return requestQueueId, nil
	}
	return 0, errors.New("Record created but no requestQueueId returned")
}

func (q *RequestQueueService) AddStartingRequest(scraperId int,
	requestBlob []byte) (int, error) {
	request := harvest.QueuedRequestFields{
		Blob: requestBlob,
	}
	return q.enqueueRequest(scraperId, request, true)
}

func (q *RequestQueueService) EnqueueRequest(scraperId int,
	request harvest.QueuedRequestFields) (int, error) {
	return q.enqueueRequest(scraperId, request, false)
}
