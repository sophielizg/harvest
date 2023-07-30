package request

import (
	"net/http"

	"github.com/sophielizg/harvest/connections"
)

func DoRequest(conn *connections.Connections, cacheId string, req *Request) (SelectorsResult, error) {
	var client *http.Client
	if req.CookiesSettings != nil && req.CookiesSettings.EnableCookies {
		// TODO: add check for null conn
		var err error
		client, err = getClientWithCookies(conn.CookiesCache, cacheId, req.CookiesSettings.Ttl)
		if err != nil {
			return nil, err
		}
		defer setCookiesFromClient(client)
	} else {
		client = http.DefaultClient
	}

	resp, err := client.Do(req.Request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil, BadStatusCodeError
	}

	return applySelectors(resp, req.Selectors)
}
