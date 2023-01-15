package mysql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/sophielizg/harvest/api/harvest"
)

type CrawlService struct{}

type CrawlConfig harvest.CrawlConfig

func (crawlConfig *CrawlConfig) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("incompatible type for CrawlConfig")
	}
	return json.Unmarshal(b, &crawlConfig)
}

func (crawlConfig *CrawlConfig) Value() (driver.Value, error) {
	j, err := json.Marshal(crawlConfig)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}

func scanCrawl(rows *sql.Rows) (*harvest.Crawl, error) {
	var crawl harvest.Crawl

	err := rows.Scan(&crawl.CrawlId, &crawl.Name, &crawl.CreatedTimestamp,
		&crawl.Running, &crawl.Config)

	return &crawl, err
}

func (c *CrawlService) Crawl(crawlId int) (*harvest.Crawl, error) {
	row, err := db.Query("CALL getCrawlByCrawlId(?);", crawlId)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	return scanCrawl(row)
}
