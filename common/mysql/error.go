package mysql

import (
	"database/sql"

	"github.com/sophielizg/harvest/common/harvest"
)

type ErrorService struct {
	Db *sql.DB
}

func (e *ErrorService) AddError(runnerId int, parseError harvest.ErrorFields) error {
	stmt, err := e.Db.Prepare("CALL addError(?, ?, ?, ?, ?, ?, ?, ?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(parseError.RunId, runnerId, parseError.RequestId, parseError.ParserId,
		parseError.StatusCode, parseError.Response, parseError.IsMissngParseResult,
		parseError.ErrorMessage)
	return err
}
