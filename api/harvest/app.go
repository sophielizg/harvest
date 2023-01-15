package harvest

type App struct {
	configService       ConfigService
	crawlService        CrawlService
	errorService        ErrorService
	parserService       ParserService
	requestQueueService RequestQueueService
	resultService       ResultService
	runService          RunService
	scraperService      ScraperService
	statusService       StatusService
}
