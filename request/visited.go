package request

import (
	"github.com/sophielizg/harvest/cache"
)

func generateKey(cacheId string, req *Request) string {
	// TODO: generate actual key
	return ""
}

func ShouldVisit(cache cache.Cache, cacheId string, req *Request) (bool, error) {
	isVisited, err := cache.Exists(generateKey(cacheId, req))
	if err != nil {
		return false, err
	}
	return !isVisited, nil
}

func MarkVisited(cache cache.Cache, cacheId string, req *Request) error {
	return cache.Put(generateKey(cacheId, req), true, req.VisitSettings.Ttl)
}
