package storage

import (
	"net/url"
	"sync"

	"github.com/sophielizg/harvest/common"
	collyCommon "github.com/sophielizg/harvest/runner/colly/common"
)

type StorageServices struct {
	CookieService       common.CookieService
	VisitedService      common.VisitedService
	RequestQueueService common.RequestQueueService
}

type Storage struct {
	collyCommon.SharedFields
	StorageServices
	mu sync.RWMutex // Only used for dequeue method
}

func (s *Storage) Init() error { return nil }

func (s *Storage) Visited(requestId uint64) error {
	err := s.VisitedService.SetIsVisited(s.RunId, requestId)
	if err != nil {
		s.Logger.WithFields(common.LogFields{
			"error":     err,
			"requestId": requestId,
			"ids":       s.SharedIds,
		}).Warn("An error ocurred within SetIsVisited while setting request as visited")
	}
	return err
}

func (s *Storage) IsVisited(requestId uint64) (bool, error) {
	isVisited, err := s.VisitedService.GetIsVisited(s.RunId, requestId)
	if err != nil {
		s.Logger.WithFields(common.LogFields{
			"error":     err,
			"requestId": requestId,
			"ids":       s.SharedIds,
		}).Warn("An error ocurred within GetIsVisited while getting request visited status")
	}
	return isVisited, err
}

func (s *Storage) Cookies(u *url.URL) string {
	cookies, err := s.CookieService.GetCookies(s.RunId, u.Host)
	if err != nil {
		s.Logger.WithFields(common.LogFields{
			"error": err,
			"url":   u,
			"ids":   s.SharedIds,
		}).Warn("An error ocurred within GetCookies while getting cookies for url")
	}
	return cookies
}

func (s *Storage) SetCookies(u *url.URL, cookies string) {
	err := s.CookieService.SetCookies(s.RunId, u.Host, cookies)
	if err != nil {
		s.Logger.WithFields(common.LogFields{
			"error": err,
			"url":   u,
			"ids":   s.SharedIds,
		}).Warn("An error ocurred within SetCookies while setting cookies for url")
	}
}

func (s *Storage) QueueSize() (int, error) {
	size, err := s.RequestQueueService.GetQueueSize(s.RunId)
	if err != nil {
		s.Logger.WithFields(common.LogFields{
			"error": err,
			"ids":   s.SharedIds,
		}).Warn("An error ocurred within GetQueueSize while getting queue size")
	}
	return size, err
}

func (s *Storage) GetRequest() ([]byte, error) {
	// Prevent deadlock when multiple threads try to dequeue at once
	s.mu.Lock()
	defer s.mu.Unlock()

	reqs, err := s.RequestQueueService.DequeueRequests(s.RunId, s.RunnerId, 1)
	if err != nil {
		s.Logger.WithFields(common.LogFields{
			"error": err,
			"ids":   s.SharedIds,
		}).Error("An error ocurred within DequeueRequests while getting a new request")
		return nil, err
	}

	if len(reqs) == 1 {
		return reqs[0], nil
	}
	return nil, nil
}

func (s *Storage) AddRequest(requestBlob []byte) error {
	runId, runnerId := s.RunId, s.RunnerId
	_, err := s.RequestQueueService.EnqueueRequest(common.QueuedRequestFields{
		ScraperId: s.ScraperId,
		RunId:     &runId,
		RunnerId:  &runnerId,
		Blob:      requestBlob,
	})
	if err != nil {
		s.Logger.WithFields(common.LogFields{
			"error": err,
			"ids":   s.SharedIds,
		}).Error("An error ocurred within EnqueueRequest while adding a new request")
	}
	return err
}
