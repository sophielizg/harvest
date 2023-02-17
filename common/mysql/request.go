package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common"
)

type RequestService struct {
	db *sql.DB
}

func (r *RequestService) AddRequest(request common.RequestFields) (int, error) {
	rows, err := r.db.Query("CALL addRequest(?, ?, ?, ?, ?, ?);", request.RunId, request.Url,
		request.Method, request.Blob, request.ParentRequestId, request.OriginatorRequestId)
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
