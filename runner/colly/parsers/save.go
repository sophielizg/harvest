package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
)

func (p *Parsers) saveResult(response *colly.Response, parserId int, parsedValue string) {
	requestId, err := getRequestId(response.Request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getRequestId error: %s", err)
	}

	resultFields := harvest.ResultFields{
		RunId:     p.RunId,
		RequestId: requestId,
		ParserId:  parserId,
		Value:     parsedValue,
	}

	err = p.ResultService.AddResult(p.RunnerId, resultFields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "AddResult error: %s", err)
	}
}

func (p *Parsers) saveError(response *colly.Response, parserId int, parseError error,
	isMissingParseResult bool) {
	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.Marshal error: %s", err)
		return
	}

	requestId, err := getRequestId(response.Request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getRequestId error: %s", err)
	}

	errorFields := harvest.ErrorFields{
		RunId:               p.RunId,
		RequestId:           requestId,
		ParserId:            parserId,
		StatusCode:          response.StatusCode,
		Response:            marshaledResponse,
		IsMissngParseResult: isMissingParseResult,
		ErrorMessage:        parseError.Error(),
	}

	err = p.ErrorService.AddError(p.RunnerId, errorFields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "AddError error: %s", err)
	}
}

func (p *Parsers) saveAndEnqueue(response *colly.Response, parser harvest.Parser,
	parsedValue string) {
	if parsedValue == "" {
		missingResultErr := errors.New("No result found from parser")
		p.saveError(response, parser.ParserId, missingResultErr, true)
	} else {
		p.saveResult(response, parser.ParserId, parsedValue)

		if parser.EnqueueScraperId != nil {
			err := p.enqueueRequest(response.Request, parser, parsedValue)
			if err != nil {
				p.saveError(response, parser.ParserId, err, false)
			}
		}
	}
}
