package colly

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
)

func (app *App) Configure(collector *colly.Collector) error {
	scraper, err := app.ScraperService.Scraper(app.ScraperId)
	if err != nil {
		return err
	}

	config := scraper.Config

	if config.MaxDepth != 0 {
		collector.MaxDepth = config.MaxDepth
	}

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
			return err
		}
		collector.SetProxyFunc(proxySwitcher)
	}

	if config.DomainRules != nil {
		for domain, rules := range config.DomainRules {
			collector.Limit(&colly.LimitRule{
				DomainGlob:  domain,
				RandomDelay: time.Duration(rules.RandomDelaySeconds) * time.Second,
				Parallelism: rules.Parallelism,
			})
		}
	}

	if config.RequestTimeout != 0 {
		collector.SetRequestTimeout(time.Duration(config.RequestTimeout) * time.Second)
	}

	if config.GlobalCookies != nil {
		collector.OnRequest(func(request *colly.Request) {
			for key, values := range config.GlobalCookies {
				for _, value := range values {
					request.Headers.Add(key, value)
				}
			}
		})
	}

	return nil
}
