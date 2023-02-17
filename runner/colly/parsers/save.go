package parsers

import (
	"encoding/json"
	"errors"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common"
)

func (p *Parsers) saveResult(response *colly.Response, parserId int, parsedValue string,
	elementIndex *int) {
	requestId, err := getRequestId(response.Request)
	if err != nil {
		p.Logger.WithFields(common.LogFields{
			"error":   err,
			"ids":     p.SharedIds,
			"request": response.Request,
		}).Warn("An error ocurred in getRequestId while saving result")
		return
	}

	resultFields := common.ResultFields{
		RunId:        p.RunId,
		RequestId:    requestId,
		ParserId:     parserId,
		ElementIndex: elementIndex,
		Value:        parsedValue,
	}

	_, err = p.ResultService.AddResult(p.RunnerId, resultFields)
	if err != nil {
		p.Logger.WithFields(common.LogFields{
			"error":        err,
			"ids":          p.SharedIds,
			"resultFields": resultFields,
		}).Warn("An error ocurred in AddResult while saving result")
	}
}

func (p *Parsers) saveError(response *colly.Response, parserId int, parseError error,
	elementIndex *int, isMissingParseResult bool) {
	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		p.Logger.WithFields(common.LogFields{
			"error":    err,
			"ids":      p.SharedIds,
			"response": response,
		}).Warn("An error ocurred in json.Marshal while saving error")
		return
	}

	requestId, err := getRequestId(response.Request)
	if err != nil {
		p.Logger.WithFields(common.LogFields{
			"error":   err,
			"ids":     p.SharedIds,
			"request": response.Request,
		}).Warn("An error ocurred in getRequestId while saving error")
		return
	}

	errorFields := common.ErrorFields{
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
		p.Logger.WithFields(common.LogFields{
			"error":       err,
			"ids":         p.SharedIds,
			"errorFields": errorFields,
		}).Warn("An error ocurred in AddError while saving error")
	}
}

func (p *Parsers) saveAndEnqueue(response *colly.Response, parser common.Parser,
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
