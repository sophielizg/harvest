package harvest

import "net/http"

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

type Request struct {
	Request       *http.Request
	Selector      *Selector
	VisitSettings *VisitSettings
}

func NewRequest(req *http.Request, options ...func(*Request)) *Request {
	r := &Request{
		Request: req,
	}

	for _, option := range options {
		option(r)
	}

	return r
}

func WithSelector(selector *Selector) func(*Request) {
	return func(r *Request) {
		r.Selector = selector
	}
}

func WithVisitSettings(settings *VisitSettings) func(*Request) {
	return func(r *Request) {
		r.VisitSettings = settings
	}
}
