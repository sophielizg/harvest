package request

import (
	"bytes"
	"net/http"

	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
)

type SelectorsResult map[string][]string

func applyHtmlSelectors(resp *http.Response, selectors []*Selector) (SelectorsResult, error) {
	defer resp.Body.Close()
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

func applySelectors(resp *http.Response, selectors []*Selector) (SelectorsResult, error) {
	switch resp.Header.Get("Content-Type") {
	case "text/html":
		return applyHtmlSelectors(resp, selectors)
	default:
		return nil, nil
	}
}
