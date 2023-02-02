package mysql

import (
	"database/sql"
)

type VisitedService struct {
	Db *sql.DB
}

func (v *VisitedService) GetIsVisited(runId int, requestHash uint64) (bool, error) {
	rows, err := v.Db.Query("CALL getIsVisited(?, ?);", runId, requestHash)
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

func (v *VisitedService) SetIsVisited(runId int, requestHash uint64) error {
	stmt, err := v.Db.Prepare("CALL setIsVisited(?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(runId, requestHash)
	return err
}
