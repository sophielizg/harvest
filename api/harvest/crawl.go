package harvest

import "time"

type CrawlConfig struct {
	AllowedDomains     []string `json:"allowedDomains"`
	MaxDepth           int      `json:"maxDepth"`
	UserAgent          string   `json:"userAgent"`
	AllowRevisit       bool     `json:"allowRevisit"`
	Proxies            []string `json:"proxies"`
	RandomDelaySeconds float32  `json:"randomDelaySeconds"`
}

type CrawlFields struct {
	Name    string      `json:"name"`
	Running bool        `json:"running"`
	Config  CrawlConfig `json:"config"`
}

type Crawl struct {
	CrawlId          int       `json:"crawlId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	CrawlFields
}

type CrawlService interface {
	Crawl(crawlId int) (*Crawl, error)
	// CrawlByName(name string) (*Crawl, error)
	// Crawls() ([]Crawl, error)
	// AddCrawl(crawl CrawlFields) (int, error)
	// DeleteCrawl(crawlId int) error
	// StartCrawl(crawlId int) error
	// StopCrawl(crawlId int) error
	// PauseCrawl(crawlId int) error
	// UnpauseCrawl(crawlId int) error
}
