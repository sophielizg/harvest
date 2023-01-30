package mysql

import (
	"database/sql"

	"github.com/sophielizg/harvest/common/harvest"
)

type ResultService struct {
	Db *sql.DB
}

func (r *ResultService) AddResult(runnerId int, result harvest.ResultFields) error {
	stmt, err := r.Db.Prepare("CALL addResult(?, ?, ?, ?, ?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(result.RunId, runnerId, result.RequestId, result.ParserId, result.Value)
	return err
}
