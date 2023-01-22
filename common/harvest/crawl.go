package harvest

import "time"

type CrawlConfig struct {
	AllowedDomains     []string          `json:"allowedDomains"`
	MaxDepth           int               `json:"maxDepth"`
	UserAgent          string            `json:"userAgent"`
	AllowRevisit       bool              `json:"allowRevisit"`
	Proxies            []string          `json:"proxies"`
	RandomDelaySeconds float32           `json:"randomDelaySeconds"`
	RequestTimeout     float32           `json:"requestTimeout"`
	Cookies            map[string]string `json:"cookies"`
}

type CrawlFields struct {
	Name   *string      `json:"name"`
	Config *CrawlConfig `json:"config"`
}

type Crawl struct {
	CrawlId          int       `json:"crawlId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	CrawlFields
}

type CrawlService interface {
	Crawl(crawlId int) (*Crawl, error)
	CrawlByName(name string) (*Crawl, error)
	Crawls() ([]Crawl, error)
	AddCrawl(crawl CrawlFields) (int, error)
	UpdateCrawl(crawlId int, crawl CrawlFields) error
	DeleteCrawl(crawlId int) error
	StartCrawl(crawlId int) error
	StopCrawl(crawlId int) error
	PauseCrawl(crawlId int) error
	UnpauseCrawl(crawlId int) error
}
