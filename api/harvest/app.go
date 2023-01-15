package harvest

type App struct {
	ConfigService       ConfigService
	CrawlService        CrawlService
	ErrorService        ErrorService
	ParserService       ParserService
	RequestQueueService RequestQueueService
	ResultService       ResultService
	RunService          RunService
	ScraperService      ScraperService
	StatusService       StatusService
}
