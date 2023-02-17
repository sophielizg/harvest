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
	defer rows.Close()

	for rows.Next() {
		var isVisited bool
		err = rows.Scan(&isVisited)
		if err != nil {
			return false, err
		}
		return isVisited, nil
	}
	return false, rows.Err()
}

func (v *VisitedService) SetIsVisited(runId int, requestHash uint64) error {
	_, err := v.Db.Exec("CALL setIsVisited(?, ?);", runId, requestHash)
	return err
}
