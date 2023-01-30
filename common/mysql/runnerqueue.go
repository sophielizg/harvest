package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type RunnerQueueService struct {
	Db *sql.DB
}

func (q *RunnerQueueService) enqueueRunner(scraperId *int, runId *int) (int, error) {
	rows, err := q.Db.Query("CALL enqueueRunner(?, ?);", scraperId, runId)
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
	return 0, errors.New("Record created but no runnerId returned")
}

func (q *RunnerQueueService) EnqueueRunnerForRun(runId int) (int, error) {
	return q.enqueueRunner(nil, &runId)
}

func (q *RunnerQueueService) EnqueueRunnerForCurrentRun(scraperId int) (int, error) {
	return q.enqueueRunner(&scraperId, nil)
}

func (q *RunnerQueueService) DequeueRunner() (*harvest.Runner, error) {
	rows, err := q.Db.Query("CALL dequeueRunner(1);")
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
	return nil, errors.New("No runners found to dequeue")
}

func (q *RunnerQueueService) EndRunner(runnerId int) error {
	stmt, err := q.Db.Prepare("CALL endRunner(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(runnerId)
	return err
}
