package storage

import (
	"fmt"
	"net/url"
	"os"

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
}

func (s *Storage) Init() error { return nil }

func (s *Storage) Visited(requestId uint64) error {
	return s.VisitedService.SetIsVisited(s.RunId, requestId)
}

func (s *Storage) IsVisited(requestId uint64) (bool, error) {
	return s.VisitedService.GetIsVisited(s.RunId, requestId)
}

func (s *Storage) Cookies(u *url.URL) string {
	cookies, err := s.CookieService.GetCookies(s.RunId, u.Host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetCookies error: %s", err)
	}
	return cookies
}

func (s *Storage) SetCookies(u *url.URL, cookies string) {
	err := s.CookieService.SetCookies(s.RunId, u.Host, cookies)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SetCookies error: %s", err)
	}
}

func (s *Storage) QueueSize() (int, error) {
	return s.RequestQueueService.GetQueueSize(s.RunId)
}

func (s *Storage) GetRequest() ([]byte, error) {
	reqs, err := s.RequestQueueService.DequeueRequests(s.RunId, s.RunnerId, 1)
	if err != nil {
		return nil, err
	}
	return reqs[0], nil
}

func (s *Storage) AddRequest(requestBlob []byte) error {
	_, err := s.RequestQueueService.EnqueueRequest(harvest.QueuedRequestFields{
		ScraperId: s.ScraperId,
		RunId:     s.RunId,
		RunnerId:  s.RunnerId,
		Blob:      requestBlob,
	})
	return err
}
