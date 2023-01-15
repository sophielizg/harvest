package harvest

import "time"

type CrawlConfig struct {
	allowedDomains     []string
	maxDepth           int
	userAgent          string
	allowRevisit       bool
	proxies            []string
	randomDelaySeconds float32
}

type CrawlFields struct {
	name    string
	running bool
	config  CrawlConfig
}

type Crawl struct {
	crawlId          int
	createdTimestamp time.Time
	CrawlFields
}

type CrawlService interface {
	Crawl(crawlId int) (*Crawl, error)
	CrawlByName(name string) (*Crawl, error)
	Crawls() ([]Crawl, error)
	AddCrawl(crawl CrawlFields) (int, error)
	DeleteCrawl(crawlId int) error
	StartCrawl(crawlId int) error
	StopCrawl(crawlId int) error
	PauseCrawl(crawlId int) error
	UnpauseCrawl(crawlId int) error
}
