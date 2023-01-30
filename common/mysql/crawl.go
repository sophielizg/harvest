package mysql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type ScraperService struct {
	Db *sql.DB
}

type ScraperConfig harvest.ScraperConfig

func (crawlConfig *ScraperConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("incompatible type for ScraperConfig")
	}
	return json.Unmarshal(b, &crawlConfig)
}

func (crawlConfig *ScraperConfig) Value() (driver.Value, error) {
	if crawlConfig == nil {
		return nil, nil
	}

	b, err := json.Marshal(crawlConfig)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func scanScraper(rows *sql.Rows) (*harvest.Scraper, error) {
	var crawlConfig *ScraperConfig
	var crawl harvest.Scraper

	err := rows.Scan(&crawl.ScraperId, &crawl.Name, &crawl.CreatedTimestamp,
		&crawlConfig)

	if crawlConfig != nil {
		convertedConfig := harvest.ScraperConfig(*crawlConfig)
		crawl.Config = &convertedConfig
	}

	return &crawl, err
}

func (c *ScraperService) Scraper(scraperId int) (*harvest.Scraper, error) {
	rows, err := c.Db.Query("CALL getScraperByScraperId(?);", scraperId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanScraper(rows)
	}
	return nil, errors.New("No crawls with specified scraperId found")
}

func (c *ScraperService) ScraperByName(name string) (*harvest.Scraper, error) {
	rows, err := c.Db.Query("CALL getScraperByName(?);", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanScraper(rows)
	}
	return nil, errors.New("No crawls with specified name found")
}

func (c *ScraperService) Scrapers() ([]harvest.Scraper, error) {
	rows, err := c.Db.Query("CALL getAllScrapers();")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var crawls []harvest.Scraper
	for rows.Next() {
		crawl, err := scanScraper(rows)
		if err != nil {
			return nil, err
		}
		crawls = append(crawls, *crawl)
	}

	return crawls, nil
}

func (c *ScraperService) AddScraper(crawl harvest.ScraperFields) (int, error) {
	var crawlConfig *ScraperConfig
	if crawl.Config != nil {
		convertedConfig := ScraperConfig(*crawl.Config)
		crawlConfig = &convertedConfig
	}

	rows, err := c.Db.Query("CALL addScraper(?, ?);", crawl.Name, &crawlConfig)
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
	return 0, errors.New("Record created but no scraperId returned")
}

func (c *ScraperService) UpdateScraper(scraperId int, crawl harvest.ScraperFields) error {
	var crawlConfig *ScraperConfig
	if crawl.Config != nil {
		convertedConfig := ScraperConfig(*crawl.Config)
		crawlConfig = &convertedConfig
	}

	stmt, err := c.Db.Prepare("CALL updateScraper(?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(scraperId, crawl.Name, &crawlConfig)
	return err
}

func (c *ScraperService) DeleteScraper(scraperId int) error {
	stmt, err := c.Db.Prepare("CALL deleteScraper(?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(scraperId)
	return err
}

func (c *ScraperService) StartScraper(scraperId int) error {
	stmt, err := c.Db.Prepare("CALL startScraper(?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(scraperId)
	return err
}

func (c *ScraperService) StopScraper(scraperId int) error {
	stmt, err := c.Db.Prepare("CALL stopScraper(?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(scraperId)
	return err
}

func (c *ScraperService) PauseScraper(scraperId int) error {
	stmt, err := c.Db.Prepare("CALL pauseScraper(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(scraperId)
	return err
}

func (c *ScraperService) UnpauseScraper(scraperId int) error {
	stmt, err := c.Db.Prepare("CALL unpauseScraper(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(scraperId)
	return err
}
