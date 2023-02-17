package mysql

import (
	"database/sql"
	"errors"

	harvest "github.com/sophielizg/harvest/common"
)

type RequestQueueService struct {
	Db *sql.DB
}

func (q *RequestQueueService) enqueueRequest(request harvest.QueuedRequestFields,
	isInitialRequest bool) (int, error) {
	rows, err := q.Db.Query("CALL enqueueRequest(?, ?, ?, ?, ?, 1);", request.ScraperId,
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

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return 0, errors.New("Record created but no requestQueueId returned")
}

func (q *RequestQueueService) GetQueueSize(runId int) (int, error) {
	rows, err := q.Db.Query("CALL getQueueSize(?);", runId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var queueSize int
		err = rows.Scan(&queueSize)
		if err != nil {
			return 0, err
		}
		return queueSize, nil
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return 0, errors.New("No rows returned for queue size for runId")
}

func (q *RequestQueueService) AddStartingRequest(scraperId int,
	requestBlob []byte) (int, error) {
	request := harvest.QueuedRequestFields{
		ScraperId: scraperId,
		Blob:      requestBlob,
	}
	return q.enqueueRequest(request, true)
}

func (q *RequestQueueService) EnqueueRequest(request harvest.QueuedRequestFields) (int, error) {
	return q.enqueueRequest(request, false)
}

func (q *RequestQueueService) DequeueRequests(runId int, runnerId int,
	numToDequeue int) ([][]byte, error) {
	rows, err := q.Db.Query("CALL dequeueRequests(?, ?, ?, 1);", runId, runnerId,
		numToDequeue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests [][]byte
	for rows.Next() {
		var requestBlob []byte
		err = rows.Scan(&requestBlob)
		if err != nil {
			return requests, err
		}
		requests = append(requests, requestBlob)
	}
	return requests, rows.Err()
}
