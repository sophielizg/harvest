package request

import (
	"bytes"
	"fmt"
	"net/http"

	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xmlquery"
)

type SelectorsResult map[string][]string

func applyHtmlSelectors(resp *http.Response, selectors []*Selector) (SelectorsResult, error) {
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	results := SelectorsResult{}
	for _, selector := range selectors {
		if selector.PageType != Html {
			continue
		}

		nodes, err := htmlquery.QueryAll(doc, selector.Xpath)
		if err != nil {
			return nil, err
		}

		selectorResults := make([]string, len(nodes))
		for i, node := range nodes {
			var b bytes.Buffer
			err := html.Render(&b, node)

			if err != nil {
				return nil, err
			}

			selectorResults[i] = b.String()
		}

		results[selector.Name] = selectorResults
	}

	return results, nil
}

func applyJsonSelectors(resp *http.Response, selectors []*Selector) (SelectorsResult, error) {
	doc, err := jsonquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	results := SelectorsResult{}
	for _, selector := range selectors {
		if selector.PageType != Json {
			continue
		}

		nodes, err := jsonquery.QueryAll(doc, selector.Xpath)
		if err != nil {
			return nil, err
		}

		selectorResults := make([]string, len(nodes))
		for i, node := range nodes {
			selectorResults[i] = fmt.Sprintf("%s", node.Value())
		}

		results[selector.Name] = selectorResults
	}

	return results, nil
}

func applyXmlSelectors(resp *http.Response, selectors []*Selector) (SelectorsResult, error) {
	doc, err := xmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	results := SelectorsResult{}
	for _, selector := range selectors {
		if selector.PageType != Json {
			continue
		}

		nodes, err := xmlquery.QueryAll(doc, selector.Xpath)
		if err != nil {
			return nil, err
		}

		selectorResults := make([]string, len(nodes))
		for i, node := range nodes {
			selectorResults[i] = node.InnerText()
		}

		results[selector.Name] = selectorResults
	}

	return results, nil
}

func applySelectors(resp *http.Response, selectors []*Selector) (SelectorsResult, error) {
	switch resp.Header.Get("Content-Type") {
	case "text/html":
		return applyHtmlSelectors(resp, selectors)
	case "application/json":
		return applyJsonSelectors(resp, selectors)
	case "text/xml", "application/xml":
		return applyXmlSelectors(resp, selectors)
	default:
		return nil, nil
	}
}
