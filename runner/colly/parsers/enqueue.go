package parsers

import (
	"bytes"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common"
	"github.com/sophielizg/harvest/common/utils"
)

func (p *Parsers) enqueueRequest(parentRequest *colly.Request, parser common.Parser,
	parsedValue string) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(parentRequest.Body)
	body := buf.Bytes()
	url := parentRequest.URL.String()

	method := "GET"
	newUrl := url
	newBody := body

	if parser.AutoIncrementRules != nil {
		var err error
		method = parentRequest.Method
		if parser.AutoIncrementRules.UrlRegex != nil {
			newUrl, err = utils.IncrementUrl(url, *parser.AutoIncrementRules.UrlRegex)
		}
		if err == nil && parser.AutoIncrementRules.BodyPath != nil {
			newBody, err = utils.IncrementBody(body, parser.AutoIncrementRules.BodyPath)
		}
		if err != nil {
			return err
		}
	} else {
		newUrl = parsedValue
		newBody = nil
	}

	newRequest, err := parentRequest.New(method, newUrl, bytes.NewReader(newBody))
	if err != nil {
		return err
	}

	err = addParentRequestIds(newRequest)
	if err != nil {
		return err
	}

	p.Logger.WithFields(common.LogFields{
		"ids": p.SharedIds,
		"url": newUrl,
	}).Debug("Adding new request to queue")
	return p.Queue.AddRequest(newRequest)
}
