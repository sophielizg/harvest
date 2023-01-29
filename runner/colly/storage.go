package colly

import "github.com/sophielizg/harvest/common/harvest"

type Storage struct {
	CrawlRunId          int
	ScrapeId            int
	RequestService      harvest.RequestService
	RequestQueueService harvest.RequestQueueService
}

func (s *Storage) Visited(requestId uint64) error {
	_, err := s.RequestService.AddRequestIsVisited(harvest.RequestFields{
		ScrapeId: s.ScrapeId,
		Id:       requestId,
	})
	return err
}

func (s *Storage) IsVisited(requestId uint64) (bool, error) {
	return s.RequestService.IsRequestVisited(s.CrawlRunId, requestId)
}
