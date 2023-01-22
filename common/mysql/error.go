package mysql

import (
	"database/sql"

	"github.com/sophielizg/harvest/common/harvest"
)

type ErrorService struct {
	Db *sql.DB
}

func (e *ErrorService) AddError(crawlId int, scrapeId int,
	parseError harvest.ErrorFields) error {
	stmt, err := e.Db.Prepare("CALL addError(?, ?, ?, ?, ?, ?, ?, ?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(crawlId, scrapeId, parseError.RequestId, parseError.ParserId,
		parseError.StatusCode, parseError.IsMissngParseResult, parseError.ErrorMessage)
	return err
}
