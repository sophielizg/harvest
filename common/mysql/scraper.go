package mysql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/sophielizg/harvest/common"
)

type ScraperService struct {
	db *sql.DB
}

type ScraperConfig common.ScraperConfig

func (scraperConfig *ScraperConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("incompatible type for ScraperConfig")
	}
	return json.Unmarshal(b, &scraperConfig)
}

func (scraperConfig *ScraperConfig) Value() (driver.Value, error) {
	if scraperConfig == nil {
		return nil, nil
	}

	b, err := json.Marshal(scraperConfig)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func scanScraper(rows *sql.Rows) (*common.Scraper, error) {
	var scraperConfig *ScraperConfig
	var scraper common.Scraper

	err := rows.Scan(&scraper.ScraperId, &scraper.Name, &scraper.CreatedTimestamp,
		&scraperConfig)

	if scraperConfig != nil {
		convertedConfig := common.ScraperConfig(*scraperConfig)
		scraper.Config = &convertedConfig
	}

	return &scraper, err
}

func (c *ScraperService) Scraper(scraperId int) (*common.Scraper, error) {
	rows, err := c.db.Query("CALL getScraperById(?);", scraperId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanScraper(rows)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("No scrapers with specified scraperId found")
}

func (c *ScraperService) Scrapers() ([]common.Scraper, error) {
	rows, err := c.db.Query("CALL getAllScrapers();")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scrapers []common.Scraper
	for rows.Next() {
		scraper, err := scanScraper(rows)
		if err != nil {
			return nil, err
		}
		scrapers = append(scrapers, *scraper)
	}

	return scrapers, rows.Err()
}

func (c *ScraperService) AddScraper(scraper common.ScraperFields) (int, error) {
	var scraperConfig *ScraperConfig
	if scraper.Config != nil {
		convertedConfig := ScraperConfig(*scraper.Config)
		scraperConfig = &convertedConfig
	}

	rows, err := c.db.Query("CALL addScraper(?, ?);", scraper.Name, &scraperConfig)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var scraperId int
		err = rows.Scan(&scraperId)
		if err != nil {
			return 0, err
		}
		return scraperId, nil
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return 0, errors.New("Record created but no scraperId returned")
}

func (c *ScraperService) UpdateScraper(scraperId int, scraper common.ScraperFields) error {
	var scraperConfig *ScraperConfig
	if scraper.Config != nil {
		convertedConfig := ScraperConfig(*scraper.Config)
		scraperConfig = &convertedConfig
	}

	_, err := c.db.Exec("CALL updateScraper(?, ?, ?);", scraperId, scraper.Name, &scraperConfig)
	return err
}

func (c *ScraperService) DeleteScraper(scraperId int) error {
	_, err := c.db.Exec("CALL deleteScraper(?, 1);", scraperId)
	return err
}
