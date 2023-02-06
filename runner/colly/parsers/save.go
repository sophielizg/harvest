package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
)

func (p *Parsers) saveResult(response *colly.Response, parserId int, parsedValue string,
	elementIndex *int) {
	requestId, err := getRequestId(response.Request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getRequestId error: %s\n", err)
	}

	resultFields := harvest.ResultFields{
		RunId:        p.RunId,
		RequestId:    requestId,
		ParserId:     parserId,
		ElementIndex: elementIndex,
		Value:        parsedValue,
	}

	_, err = p.ResultService.AddResult(p.RunnerId, resultFields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "AddResult error: %s\n", err)
	}
}

func (p *Parsers) saveError(response *colly.Response, parserId int, parseError error,
	elementIndex *int, isMissingParseResult bool) {
	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.Marshal error: %s\n", err)
		return
	}

	requestId, err := getRequestId(response.Request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getRequestId error: %s\n", err)
	}

	errorFields := harvest.ErrorFields{
		RunId:               p.RunId,
		RequestId:           requestId,
		ParserId:            parserId,
		ElementIndex:        elementIndex,
		StatusCode:          response.StatusCode,
		Response:            string(marshaledResponse),
		IsMissngParseResult: isMissingParseResult,
		ErrorMessage:        parseError.Error(),
	}

	_, err = p.ErrorService.AddError(p.RunnerId, errorFields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "AddError error: %s\n", err)
	}
}

func (p *Parsers) saveAndEnqueue(response *colly.Response, parser harvest.Parser,
	parsedValue string, elementIndex *int) {
	if parsedValue == "" {
		missingResultErr := errors.New("No result found from parser")
		p.saveError(response, parser.ParserId, missingResultErr, elementIndex, true)
	} else {
		p.saveResult(response, parser.ParserId, parsedValue, elementIndex)

		if parser.EnqueueScraperId != nil {
			err := p.enqueueRequest(response.Request, parser, parsedValue)
			if err != nil {
				p.saveError(response, parser.ParserId, err, elementIndex, false)
			}
		}
	}
}
