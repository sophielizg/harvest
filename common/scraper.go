package harvest

import "time"

type DomainRule struct {
	RandomDelaySeconds float32 `json:"randomDelaySeconds"`
	Parallelism        int     `json:"parallelism"`
}

type ScraperConfig struct {
	AllowedDomains []string              `json:"allowedDomains"`
	MaxDepth       int                   `json:"maxDepth"`
	UserAgent      string                `json:"userAgent"`
	AllowRevisit   bool                  `json:"allowRevisit"`
	Proxies        []string              `json:"proxies"`
	DomainRules    map[string]DomainRule `json:"domainRules"`
	RequestTimeout float32               `json:"requestTimeout"`
	GlobalCookies  map[string][]string   `json:"globalCookies"`
}

type ScraperFields struct {
	Name   *string        `json:"name"`
	Config *ScraperConfig `json:"config"`
}

type Scraper struct {
	ScraperId        int       `json:"scraperId"`
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	ScraperFields
}

type ScraperService interface {
	Scraper(scraperId int) (*Scraper, error)
	Scrapers() ([]Scraper, error)
	AddScraper(scraper ScraperFields) (int, error)
	UpdateScraper(scraperId int, scraper ScraperFields) error
	DeleteScraper(scraperId int) error
}
