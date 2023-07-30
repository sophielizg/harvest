package connections

import (
	"github.com/sophielizg/harvest/cache"
	"github.com/sophielizg/harvest/queue"
)

type Connections struct {
	VisitedCache cache.Cache
	CookiesCache cache.Cache
	RequestQueue queue.Queue
	ResultQueue  queue.Queue
	ErrorQueue   queue.Queue
}

func NewConnections(options ...func(*Connections)) *Connections {
	c := &Connections{}

	for _, option := range options {
		option(c)
	}

	return c
}

func WithVisitedCache(visitedCache cache.Cache) func(*Connections) {
	return func(c *Connections) {
		c.VisitedCache = visitedCache
	}
}

func WithCookiesCache(cookiesCache cache.Cache) func(*Connections) {
	return func(c *Connections) {
		c.CookiesCache = cookiesCache
	}
}

func WithRequestQueue(requestQueue queue.Queue) func(*Connections) {
	return func(c *Connections) {
		c.RequestQueue = requestQueue
	}
}

func WithResultQueue(resultQueue queue.Queue) func(*Connections) {
	return func(c *Connections) {
		c.ResultQueue = resultQueue
	}
}

func WithErrorQueue(errorQueue queue.Queue) func(*Connections) {
	return func(c *Connections) {
		c.ErrorQueue = errorQueue
	}
}
