package colly

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
)

func (app *App) Crawler() (*colly.Collector, error) {
	scrape, err := app.ScrapeService.DequeueScrape()
	if err != nil {
		return nil, err
	}

	crawl, err := app.CrawlService.Crawl(scrape.CrawlId)
	if err != nil {
		return nil, err
	}

	config := crawl.Config
	collector := colly.NewCollector(
		colly.MaxDepth(config.MaxDepth),
	)

	if config.AllowedDomains != nil {
		collector.AllowedDomains = config.AllowedDomains
	}

	if config.UserAgent == "" {
		extensions.RandomUserAgent(collector)
	} else {
		collector.UserAgent = config.UserAgent
	}

	if config.AllowRevisit {
		collector.AllowURLRevisit = true
	}

	if config.Proxies != nil {
		proxySwitcher, err := proxy.RoundRobinProxySwitcher(config.Proxies...)
		if err != nil {
			return nil, err
		}
		collector.SetProxyFunc(proxySwitcher)
	}

	if config.RandomDelaySeconds != 0 {
		collector.Limit(&colly.LimitRule{
			DomainGlob:  "*",
			Parallelism: 2,
			RandomDelay: time.Duration(config.RandomDelaySeconds) * time.Second,
		})
	}

	return collector, nil
}
