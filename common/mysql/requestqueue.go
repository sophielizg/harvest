package mysql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/sophielizg/harvest/common/harvest"
)

type RequestQueueService struct {
	Db *sql.DB
}

type RequestToScrape harvest.RequestToScrape

func (requestToScrape *RequestToScrape) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("incompatible type for RequestToScrape")
	}
	return json.Unmarshal(b, &requestToScrape)
}

func (requestToScrape *RequestToScrape) Value() (driver.Value, error) {
	if requestToScrape == nil {
		return nil, nil
	}

	b, err := json.Marshal(requestToScrape)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (rq *RequestQueueService) EnqueueRequest(request harvest.QueuedRequestFields) (int, error) {
	requestToScrape := RequestToScrape(request.Request)
	rows, err := rq.Db.Query("CALL enqueueRequest(?, ?, ?, ?, ?, 1);", request.CrawlId,
		request.ScrapeId, requestToScrape, request.CreatedByRequestId,
		request.IsInitialRequest)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var requestQueueId int
		err = rows.Scan(&requestQueueId)
		if err != nil {
			return 0, err
		}
		return requestQueueId, nil
	}
	return 0, errors.New("Record created but no requestQueueId returned")
}
