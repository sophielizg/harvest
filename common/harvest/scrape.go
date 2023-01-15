package harvest

type Scrape struct {
	ScrapeId int `json:"scrapeId"`
	CrawlId  int `json:"crawlId"`
}

type ScrapeService interface {
	EnqueueScrape(crawlId int) (int, error)
	DequeueScrape() (*Scrape, error)
	EndScrape(scrapeId int) error
}
