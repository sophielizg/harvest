package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type RunnerQueueService struct {
	Db *sql.DB
}

func (s *RunnerQueueService) EnqueueRunner(scraperId int) (int, error) {
	rows, err := s.Db.Query("CALL enqueueScrape(?);", scraperId)
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

func (s *RunnerQueueService) DequeueRunner() (*harvest.Scrape, error) {
	rows, err := s.Db.Query("CALL dequeueScrape(1);")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var scrape harvest.Scrape
		err = rows.Scan(&scrape.RunnerId, &scrape.RunId, &scrape.ScraperId)
		if err != nil {
			return nil, err
		}
		return &scrape, nil
	}
	return nil, errors.New("No scrapes found to dequeue")
}

func (s *RunnerQueueService) EndRunner(runnerId int) error {
	stmt, err := s.Db.Prepare("CALL endScrape(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(runnerId)
	return err
}
