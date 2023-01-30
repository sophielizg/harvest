package mysql

import (
	"database/sql"

	"github.com/sophielizg/harvest/common/harvest"
)

type RequestService struct {
	Db *sql.DB
}

func (r *RequestService) addOrUpdateRequest(runId int, requestId uint64,
	request harvest.RequestFields) error {
	stmt, err := r.Db.Prepare("CALL addRequestIsVisited(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(runId, requestId, request.Blob, request.ParentRequestId,
		request.OriginatorRequestId)
	return err
}

func (r *RequestService) IsRequestVisited(runId int, requestId uint64) (bool, error) {
	rows, err := r.Db.Query("CALL getRequestIsVisited(?, ?);", runId, requestId)
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

func (r *RequestService) AddRequestIsVisited(runId int, requestId uint64) error {
	return r.addOrUpdateRequest(runId, requestId, harvest.RequestFields{})
}

func (r *RequestService) UpdateRequest(runId int, requestId uint64,
	request harvest.RequestFields) error {
	return r.addOrUpdateRequest(runId, requestId, request)
}
