package mysql

import (
	"database/sql"
	"errors"

	harvest "github.com/sophielizg/harvest/common"
)

type ErrorService struct {
	Db *sql.DB
}

func (e *ErrorService) AddError(runnerId int, parseError harvest.ErrorFields) (int, error) {
	rows, err := e.Db.Query("CALL addError(?, ?, ?, ?, ?, ?, ?, ?, ?, 1);", parseError.RunId,
		runnerId, parseError.RequestId, parseError.ParserId, parseError.ElementIndex,
		parseError.StatusCode, parseError.Response, parseError.IsMissngParseResult,
		parseError.ErrorMessage)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var errorId int
		err = rows.Scan(&errorId)
		if err != nil {
			return 0, err
		}
		return errorId, nil
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return 0, errors.New("Error added but no errorId returned")
}
