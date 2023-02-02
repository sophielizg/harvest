package parsers

import (
	"errors"

	"github.com/gocolly/colly"
)

func getRequestId(request *colly.Request) (int, error) {
	if requestId, ok := request.Ctx.GetAny("requestId").(int); ok {
		return requestId, nil
	}
	return 0, errors.New("Could not fetch requestId from context")
}

func addParentRequestIds(request *colly.Request) error {
	requestId, err := getRequestId(request)
	if err != nil {
		return err
	}

	request.Ctx.Put("parentRequestId", requestId)
	if request.Ctx.GetAny("originatorRequestId") == nil {
		request.Ctx.Put("originatorRequestId", requestId)
	}
	return nil
}
