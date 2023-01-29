package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type RequestService struct {
	Db *sql.DB
}

func (r *RequestService) IsRequestVisited(crawlRunId int, requestId uint64) (bool, error) {
	rows, err := r.Db.Query("CALL getRequestIsVisited(?, ?);", crawlRunId, requestId)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var isVisited bool
		err = rows.Scan(&isVisited)
		if err != nil {
			return false, err
		}
		return isVisited, nil
	}
	return false, nil
}

func (r *RequestService) AddRequestIsVisited(request harvest.RequestFields) (int, error) {
	rows, err := r.Db.Query("CALL addRequestIsVisited(?, ?, ?, ?);", request.ScrapeId,
		request.Id, request.Blob, request.CreatedByRequestId)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		var requestId int
		err = rows.Scan(&requestId)
		if err != nil {
			return 0, err
		}
		return requestId, nil
	}
	return 0, errors.New("Request added as visited but no requestId returned")
}
