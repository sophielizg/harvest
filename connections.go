package harvest

import (
	"github.com/sophielizg/harvest/cache"
	"github.com/sophielizg/harvest/queue"
)

type Connections struct {
	visitedCache cache.Cache
	cookiesCache cache.Cache
	requestQueue queue.Queue
	resultQueue  queue.Queue
	errorQueue   queue.Queue
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
		c.visitedCache = visitedCache
	}
}

func WithCookiesCache(cookiesCache cache.Cache) func(*Connections) {
	return func(c *Connections) {
		c.cookiesCache = cookiesCache
	}
}

func WithRequestQueue(requestQueue queue.Queue) func(*Connections) {
	return func(c *Connections) {
		c.requestQueue = requestQueue
	}
}

func WithResultQueue(resultQueue queue.Queue) func(*Connections) {
	return func(c *Connections) {
		c.resultQueue = resultQueue
	}
}

func WithErrorQueue(errorQueue queue.Queue) func(*Connections) {
	return func(c *Connections) {
		c.errorQueue = errorQueue
	}
}
