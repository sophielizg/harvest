package colly

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/sophielizg/harvest/common/harvest"
)

func valueFromJson(data *interface{}, jsonPath []string) (*interface{}, error) {
	curr := data
	for _, key := range jsonPath {
		if m, ok := (*curr).(map[string]interface{}); ok {
			val := m[key]
			curr = &val
			continue
		}

		if l, ok := (*curr).([]interface{}); ok {
			idx, err := strconv.Atoi(key)
			if err != nil {
				return nil, err
			}

			val := l[idx]
			curr = &val
			continue
		}

		return nil, errors.New("Could not cast data as valid json")
	}

	return curr, nil
}

func getJson(data []byte, jsonPath []string) (string, error) {
	var j interface{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return "", err
	}

	val, err := valueFromJson(&j, jsonPath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", *val), nil
}

func incrementBytes(bytes []byte) []byte {
	num, err := strconv.Atoi(string(bytes))
	if err != nil {
		return nil
	}

	newNum := num + 1
	return []byte(strconv.Itoa(newNum))
}

func incrementNum(val interface{}) (int, error) {
	if num, ok := val.(float64); ok {
		return int(num) + 1, nil
	} else if num, ok := val.(int); ok {
		return num + 1, nil
	}
	return 0, errors.New("Could not cast value as valid float")
}

func incrementJson(data []byte, jsonPath []string) ([]byte, error) {
	var j interface{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return nil, err
	}

	parentVal, err := valueFromJson(&j, jsonPath[:len(jsonPath)-1])
	if err != nil {
		return nil, err
	}

	key := jsonPath[len(jsonPath)-1]
	if m, ok := (*parentVal).(map[string]interface{}); ok {
		m[key], err = incrementNum(m[key])
	} else if l, ok := (*parentVal).([]interface{}); ok {
		idx, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		l[idx], err = incrementNum(m[key])
	}

	if err != nil {
		return nil, err
	}

	return json.Marshal(j)
}

func (app *App) saveAndEnqueue(parentRequest *colly.Request, parser harvest.Parser,
	result harvest.Result) error {
	// TODO: implement this
	// get app.CrawlId, app.ScrapeId, ctx.Get("requestId")
	nMissing := 0
	nSuccess := 0
	if result.Value == "" {
		// add error in ErrorService
	} else {
		// add result in ResultService

		if parser.EnqueueCrawlId != nil && parser.AutoIncrementRules != nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(parentRequest.Body)
			body := buf.Bytes()
			url := parentRequest.URL.String()

			newUrl := url
			newBody := body

			if parser.AutoIncrementRules.UrlRegex != nil {
				re := regexp.MustCompile(*parser.AutoIncrementRules.UrlRegex)
				newUrl = string(re.ReplaceAllFunc([]byte(url), incrementBytes))
			}

			if parser.AutoIncrementRules.BodyPath != nil {
				var err error
				newBody, err = incrementJson(body, parser.AutoIncrementRules.BodyPath)
				if err != nil {
					return err
				}
			}

			// Enqueue request with newUrl and newBody
		} else if parser.EnqueueCrawlId != nil {
			// Queue request with result.Value as URL
		}
	}

	// update status in StatusService
}

func (app *App) htmlParser(collector *colly.Collector, parser harvest.Parser) error {
	if parser.Selector == nil {
		return errors.New("The parser.Selector field is required for type html")
	}

	collector.OnHTML(*parser.Selector, func(e *colly.HTMLElement) {
		var val string
		if parser.Attr == nil {
			val = e.Text
		} else {
			val = e.Attr(*parser.Attr)
		}

		app.saveResult(e.Request, harvest.Result{ParserId: parser.ParserId, Value: val})
	})

	return nil
}

func (app *App) xmlParser(collector *colly.Collector, parser harvest.Parser) error {
	if parser.Xpath == nil {
		return errors.New("The parser.Xpath field is required for type xml")
	}

	collector.OnXML(*parser.Xpath, func(e *colly.XMLElement) {
		val := e.Text
		app.saveResult(e.Request, harvest.Result{ParserId: parser.ParserId, Value: val})
	})

	return nil
}

func (app *App) jsonParser(collector *colly.Collector, parser harvest.Parser) error {
	if parser.JsonPath == nil {
		return errors.New("The parser.JsonPath field is required for type json")
	}

	collector.OnResponse(func(r *colly.Response) {
		if r.Headers.Get("Content-Type") == "application/json" {
			val, err := getJson(r.Body, parser.JsonPath)
			if err != nil {
				// TODO call error service
			}
			app.saveResult(r.Request, harvest.Result{ParserId: parser.ParserId, Value: val})
		}
	})

	return nil
}

func (app *App) AddParsers(collector *colly.Collector) error {

}
