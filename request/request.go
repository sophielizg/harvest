package request

import (
	"context"
	"net/http"

	"github.com/sophielizg/harvest/common"
	"github.com/sophielizg/harvest/connections"
)

type PageType int8

const (
	Html PageType = iota
	Json
	Xml
)

type Selector struct {
	PageType PageType
	Xpath    string
}

type VisitSettings struct {
	AllowRevisit bool
	Ttl          int
}

type CookiesSettings struct {
	EnableCookies bool
	Ttl           int
}

type RandomDelaySeconds struct {
	Min float32
	Max float32
}

type RequestSettings struct {
	Selectors          []*Selector
	VisitSettings      *VisitSettings
	CookiesSettings    *CookiesSettings
	RandomDelaySeconds *RandomDelaySeconds
}

func (s *RequestSettings) FillInNulls(fillSettings *RequestSettings) {
	if s.Selectors == nil {
		s.Selectors = fillSettings.Selectors
	}
	if s.VisitSettings == nil {
		s.VisitSettings = fillSettings.VisitSettings
	}
	if s.CookiesSettings == nil {
		s.CookiesSettings = fillSettings.CookiesSettings
	}
	if s.RandomDelaySeconds == nil {
		s.RandomDelaySeconds = fillSettings.RandomDelaySeconds
	}
}

type Request struct {
	Request *http.Request
	*RequestSettings
}

func NewRequest(req *http.Request, options ...func(*RequestSettings)) *Request {
	r := &Request{
		Request: req,
	}

	for _, option := range options {
		option(r.RequestSettings)
	}

	return r
}

func WithSelectors(selectors ...*Selector) func(*RequestSettings) {
	return func(r *RequestSettings) {
		r.Selectors = selectors
	}
}

func WithVisitSettings(settings *VisitSettings) func(*RequestSettings) {
	return func(r *RequestSettings) {
		r.VisitSettings = settings
	}
}

func WithCookiesSettings(settings *CookiesSettings) func(*RequestSettings) {
	return func(r *RequestSettings) {
		r.CookiesSettings = settings
	}
}

func WithRandomDelaySeconds(delay *RandomDelaySeconds) func(*RequestSettings) {
	return func(r *RequestSettings) {
		r.RandomDelaySeconds = delay
	}
}

type requestMessage struct {
	ctx context.Context
	req *Request
}

func SendRequests(ctx context.Context, conn *connections.Connections, delayCache *DelayCache, reqs ...*Request) error {
	cacheId, useCache := ctx.Value(common.CacheIdKey).(string)

	reqsToSendNow := make([]*requestMessage, 0, len(reqs))
	reqsToSendDelayed := make([]*requestMessage, 0, len(reqs))

	for _, req := range reqs {
		if !useCache || req.VisitSettings == nil || req.VisitSettings.AllowRevisit {
			addRequestForSend(&reqsToSendNow, &reqsToSendDelayed, &requestMessage{ctx, req})
			continue
		}

		shouldSend, err := ShouldVisit(conn.VisitedCache, cacheId, req)
		if err != nil {
			return err
		}
		if shouldSend {
			addRequestForSend(&reqsToSendNow, &reqsToSendDelayed, &requestMessage{ctx, req})
		}
	}

	// TODO: check for request queue and delay cache existence
	if len(reqsToSendDelayed) > 0 {
		delayCache.SendMessages(reqsToSendDelayed)
	}

	return conn.RequestQueue.SendMessages(reqsToSendNow)
}

func addRequestForSend(reqsToSendNow *[]*requestMessage, reqsToSendDelayed *[]*requestMessage, message *requestMessage) {
	if message.req.RandomDelaySeconds == nil {
		*reqsToSendNow = append(*reqsToSendNow, message)
	} else {
		*reqsToSendDelayed = append(*reqsToSendDelayed, message)
	}
}
