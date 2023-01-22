package colly

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
)

func (app *App) Collector() (*colly.Collector, error) {
	crawl, err := app.CrawlService.Crawl(app.CrawlId)
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
			RandomDelay: time.Duration(config.RandomDelaySeconds) * time.Second,
		})
	}

	if config.RequestTimeout != 0 {
		collector.SetRequestTimeout(time.Duration(config.RequestTimeout) * time.Second)
	}

	if config.Cookies != nil {
		collector.OnRequest(func(request *colly.Request) {
			for key, values := range config.Cookies {
				for _, value := range values {
					request.Headers.Add(key, value)
				}
			}
		})
	}

	return collector, nil
}
