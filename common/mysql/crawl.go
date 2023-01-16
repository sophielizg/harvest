package mysql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type CrawlService struct {
	Db *sql.DB
}

type CrawlConfig harvest.CrawlConfig

func (crawlConfig *CrawlConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("incompatible type for CrawlConfig")
	}
	return json.Unmarshal(b, &crawlConfig)
}

func (crawlConfig *CrawlConfig) Value() (driver.Value, error) {
	if crawlConfig == nil {
		return nil, nil
	}

	b, err := json.Marshal(crawlConfig)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func scanCrawl(rows *sql.Rows) (*harvest.Crawl, error) {
	var crawlConfig *CrawlConfig
	var crawl harvest.Crawl

	err := rows.Scan(&crawl.CrawlId, &crawl.Name, &crawl.CreatedTimestamp,
		&crawl.Running, &crawlConfig)

	if crawlConfig != nil {
		convertedConfig := harvest.CrawlConfig(*crawlConfig)
		crawl.Config = &convertedConfig
	}

	return &crawl, err
}

func (c *CrawlService) Crawl(crawlId int) (*harvest.Crawl, error) {
	rows, err := c.Db.Query("CALL getCrawlByCrawlId(?);", crawlId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanCrawl(rows)
	}
	return nil, errors.New("No crawls with specified crawlId found")
}

func (c *CrawlService) CrawlByName(name string) (*harvest.Crawl, error) {
	rows, err := c.Db.Query("CALL getCrawlByName(?);", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanCrawl(rows)
	}
	return nil, errors.New("No crawls with specified name found")
}

func (c *CrawlService) Crawls() ([]harvest.Crawl, error) {
	rows, err := c.Db.Query("CALL getAllCrawls();")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var crawls []harvest.Crawl
	for rows.Next() {
		crawl, err := scanCrawl(rows)
		if err != nil {
			return nil, err
		}
		crawls = append(crawls, *crawl)
	}

	return crawls, nil
}

func (c *CrawlService) AddCrawl(crawl harvest.CrawlFields) (int, error) {
	var crawlConfig *CrawlConfig
	if crawl.Config != nil {
		convertedConfig := CrawlConfig(*crawl.Config)
		crawlConfig = &convertedConfig
	}

	rows, err := c.Db.Query("CALL addCrawl(?, ?);", crawl.Name, &crawlConfig)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var crawlId int
		err = rows.Scan(&crawlId)
		if err != nil {
			return 0, err
		}
		return crawlId, nil
	}
	return 0, errors.New("Record created but no crawlId returned")
}

func (c *CrawlService) UpdateCrawl(crawlId int, crawl harvest.CrawlFields) error {
	var crawlConfig *CrawlConfig
	if crawl.Config != nil {
		convertedConfig := CrawlConfig(*crawl.Config)
		crawlConfig = &convertedConfig
	}

	stmt, err := c.Db.Prepare("CALL updateCrawl(?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(crawlId, crawl.Name, &crawlConfig)
	return err
}

func (c *CrawlService) DeleteCrawl(crawlId int) error {
	stmt, err := c.Db.Prepare("CALL deleteCrawl(?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(crawlId)
	return err
}

func (c *CrawlService) StartCrawl(crawlId int) error {
	stmt, err := c.Db.Prepare("CALL startCrawl(?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(crawlId)
	return err
}

func (c *CrawlService) StopCrawl(crawlId int) error {
	stmt, err := c.Db.Prepare("CALL stopCrawl(?, 1);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(crawlId)
	return err
}

func (c *CrawlService) PauseCrawl(crawlId int) error {
	stmt, err := c.Db.Prepare("CALL pauseCrawl(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(crawlId)
	return err
}

func (c *CrawlService) UnpauseCrawl(crawlId int) error {
	stmt, err := c.Db.Prepare("CALL unpauseCrawl(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(crawlId)
	return err
}
