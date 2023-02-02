package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type RequestService struct {
	Db *sql.DB
}

func (r *RequestService) AddRequest(request harvest.RequestFields) (int, error) {
	rows, err := r.Db.Query("CALL addRequest(?, ?, ?, ?);", request.RunId, request.Blob,
		request.ParentRequestId, request.OriginatorRequestId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var requestId int
		err = rows.Scan(&requestId)
		if err != nil {
			return 0, err
		}
		return requestId, nil
	}

	return 0, errors.New("Request added but no requestId returned")
}

func (r *RequestService) UpdateRequest(requestId int, request harvest.RequestFields) error {
	stmt, err := r.Db.Prepare("CALL updateRequest(?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(requestId, request.Blob, request.ParentRequestId,
		request.OriginatorRequestId)
	return err
}
