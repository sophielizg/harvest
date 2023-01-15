package harvest

type ScraperService interface {
	AddScraperToCrawl(crawlId int) error
}
