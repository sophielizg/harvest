package mysql

import (
	"database/sql"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type ScrapeService struct {
	Db *sql.DB
}

func (s *ScrapeService) EnqueueScrape(crawlId int) (int, error) {
	rows, err := s.Db.Query("CALL enqueueScrape(?);", crawlId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var scrapeId int
		err = rows.Scan(&scrapeId)
		if err != nil {
			return 0, err
		}
		return scrapeId, nil
	}
	return 0, errors.New("Record created but no scrapeId returned")
}

func (s *ScrapeService) DequeueScrape() (*harvest.Scrape, error) {
	rows, err := s.Db.Query("CALL dequeueScrape(1);")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var scrape harvest.Scrape
		err = rows.Scan(&scrape.ScrapeId, &scrape.CrawlId)
		if err != nil {
			return nil, err
		}
		return &scrape, nil
	}
	return nil, errors.New("No scrapes found to dequeue")
}

func (s *ScrapeService) EndScrape(scrapeId int) error {
	stmt, err := s.Db.Prepare("CALL endScrape(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(scrapeId)
	return err
}
