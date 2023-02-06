package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type ResultService struct {
	Db *sql.DB
}

func (r *ResultService) AddResult(runnerId int, result harvest.ResultFields) (int, error) {
	rows, err := r.Db.Query("CALL addResult(?, ?, ?, ?, ?, ?, 1);", result.RunId,
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

	return 0, errors.New("Result added but no resultId returned")
}
