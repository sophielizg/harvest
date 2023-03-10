package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common"
)

type ResultService struct {
	db *sql.DB
}

func (r *ResultService) AddResult(runnerId int, result common.ResultFields) (int, error) {
	rows, err := r.db.Query("CALL addResult(?, ?, ?, ?, ?, ?, 1);", result.RunId,
		runnerId, result.RequestId, result.ParserId, result.ElementIndex, result.Value)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var resultId int
		err = rows.Scan(&resultId)
		if err != nil {
			return 0, err
		}
		return resultId, nil
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return 0, errors.New("Result added but no resultId returned")
}
