package mysql

import (
	"database/sql"
	"errors"
)

type RunService struct {
	db *sql.DB
}

func (r *RunService) CreateRun(scraperId int) (int, error) {
	rows, err := r.db.Query("CALL createRun(?, 1);", scraperId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var runId int
		err = rows.Scan(&runId)
		if err != nil {
			return 0, err
		}
		return runId, nil
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return 0, errors.New("Record created but no runId returned")
}
