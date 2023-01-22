package harvest

type Scrape struct {
	ScrapeId   int `json:"scrapeId"`
	CrawlRunId int `json:"crawlRunId"`
	CrawlId    int `json:"crawlId"`
}

type ScrapeService interface {
	EnqueueScrape(crawlId int) (int, error)
	DequeueScrape() (*Scrape, error)
	EndScrape(scrapeId int) error
}
