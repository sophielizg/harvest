package harvest

import "context"

type Scraper struct {
	conn    *Connections
	CacheId string
}

func (s *Scraper) SendRequests(ctx context.Context, reqs ...*Request) error {
	reqs, err := s.filterVisited(reqs)
	if err != nil {
		return err
	}

	return s.sendRequestsWithoutValidation(ctx, reqs)
}

func (s *Scraper) sendRequestsWithoutValidation(ctx context.Context, reqs []*Request) error {
	reqsToSend := make([]struct {
		context.Context
		*Request
	}, len(reqs))

	for i, req := range reqs {
		reqsToSend[i] = struct {
			context.Context
			*Request
		}{ctx, req}
	}

	return s.conn.requestQueue.SendMessages(reqsToSend)
}

func (s *Scraper) filterVisited(reqs []*Request) ([]*Request, error) {
	filtered := reqs[:0]

	for _, req := range reqs {
		if req.VisitSettings == nil || req.VisitSettings.AllowRevisit {
			filtered = append(filtered, req)
			continue
		}

		visitedCacheKey := "" // TODO: generate actual key
		isVisited, err := s.conn.visitedCache.Exists(visitedCacheKey)
		if err != nil {
			return nil, err
		}
		if !isVisited {
			filtered = append(filtered, req)
		}
	}

	// make sure discarded items can be garbage collected
	for i := len(filtered); i < len(reqs); i++ {
		reqs[i] = nil
	}

	return filtered, nil
}

func NewScraper(conn *Connections, options ...func(*Scraper)) *Scraper {
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
