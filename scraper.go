package harvest

import (
	"context"
	"time"

	"github.com/sophielizg/harvest/common"
	"github.com/sophielizg/harvest/connections"
	"github.com/sophielizg/harvest/request"
)

type Scraper struct {
	conn            *connections.Connections
	delayCache      *request.DelayCache
	CacheId         string
	RequestSettings *request.RequestSettings
}

func (s *Scraper) setContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, common.CacheIdKey, s.CacheId)
}

func (s *Scraper) SendRequests(ctx context.Context, reqs ...*request.Request) error {
	for _, req := range reqs {
		req.FillInNulls(s.RequestSettings)
	}

	return request.SendRequests(s.setContext(ctx), s.conn, s.delayCache, reqs...)
}

func NewScraper(conn *connections.Connections, options ...func(*Scraper)) *Scraper {
	s := &Scraper{
		conn: conn,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

func WithCacheId(id string) func(*Scraper) {
	return func(s *Scraper) {
		s.CacheId = id
	}
}

func WithDelayedRequestPollingInterval(interval time.Duration) func(*Scraper) {
	return func(s *Scraper) {
		s.delayCache = request.NewDelayCache(s.conn.RequestQueue, interval)
	}
}

func WithRequestSetting(setting func(*request.RequestSettings)) func(*Scraper) {
	return func(s *Scraper) {
		setting(s.RequestSettings)
	}
}
