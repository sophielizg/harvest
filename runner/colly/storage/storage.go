package storage

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/sophielizg/harvest/common/harvest"
	"github.com/sophielizg/harvest/runner/colly/common"
)

type StorageServices struct {
	CookieService       harvest.CookieService
	VisitedService      harvest.VisitedService
	RequestQueueService harvest.RequestQueueService
}

type Storage struct {
	common.RunnerIds
	StorageServices
	mu sync.RWMutex // Only used for dequeue method
}

func (s *Storage) Init() error { return nil }

func (s *Storage) Visited(requestId uint64) error {
	err := s.VisitedService.SetIsVisited(s.RunId, requestId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SetIsVisited error: %s\n", err)
	}
	return err
}

func (s *Storage) IsVisited(requestId uint64) (bool, error) {
	isVisited, err := s.VisitedService.GetIsVisited(s.RunId, requestId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetIsVisited error: %s\n", err)
	}
	return isVisited, err
}

func (s *Storage) Cookies(u *url.URL) string {
	cookies, err := s.CookieService.GetCookies(s.RunId, u.Host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetCookies error: %s\n", err)
	}
	return cookies
}

func (s *Storage) SetCookies(u *url.URL, cookies string) {
	err := s.CookieService.SetCookies(s.RunId, u.Host, cookies)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SetCookies error: %s\n", err)
	}
}

func (s *Storage) QueueSize() (int, error) {
	size, err := s.RequestQueueService.GetQueueSize(s.RunId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetQueueSize error: %s\n", err)
	}
	return size, err
}

func (s *Storage) GetRequest() ([]byte, error) {
	// Prevent deadlock when multiple threads try to dequeue at once
	s.mu.Lock()
	defer s.mu.Unlock()

	reqs, err := s.RequestQueueService.DequeueRequests(s.RunId, s.RunnerId, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DequeueRequests error: %s\n", err)
		return nil, err
	}

	if len(reqs) == 1 {
		return reqs[0], nil
	}
	return nil, nil
}

func (s *Storage) AddRequest(requestBlob []byte) error {
	runId, runnerId := s.RunId, s.RunnerId
	_, err := s.RequestQueueService.EnqueueRequest(harvest.QueuedRequestFields{
		ScraperId: s.ScraperId,
		RunId:     &runId,
		RunnerId:  &runnerId,
		Blob:      requestBlob,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "EnqueueRequest error: %s\n", err)
	}
	return err
}
