package colly

import "github.com/sophielizg/harvest/common/harvest"

type App struct {
	CrawlService  harvest.CrawlService
	ScrapeService harvest.ScrapeService
}
