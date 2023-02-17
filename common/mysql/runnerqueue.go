package mysql

import (
	"database/sql"
	"errors"

	harvest "github.com/sophielizg/harvest/common"
)

type RunnerQueueService struct {
	db *sql.DB
}

func (q *RunnerQueueService) enqueueRunner(scraperId *int, runId *int) (int, error) {
	rows, err := q.db.Query("CALL enqueueRunner(?, ?);", scraperId, runId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var runnerId int
		err = rows.Scan(&runnerId)
		if err != nil {
			return 0, err
		}
		return runnerId, nil
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return 0, errors.New("Record created but no runnerId returned")
}

func (q *RunnerQueueService) EnqueueRunnerForRun(runId int) (int, error) {
	return q.enqueueRunner(nil, &runId)
}

func (q *RunnerQueueService) EnqueueRunnerForCurrentRun(scraperId int) (int, error) {
	return q.enqueueRunner(&scraperId, nil)
}

func (q *RunnerQueueService) DequeueRunner() (*harvest.Runner, error) {
	rows, err := q.db.Query("CALL dequeueRunner(1);")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var runner harvest.Runner
		err = rows.Scan(&runner.ScraperId, &runner.RunId, &runner.RunnerId)
		if err != nil {
			return nil, err
		}
		return &runner, nil
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("No runners found to dequeue")
}

func (q *RunnerQueueService) EndRunner(runnerId int) error {
	_, err := q.db.Exec("CALL endRunner(?);", runnerId)
	return err
}
