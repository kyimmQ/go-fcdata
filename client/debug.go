package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *FCDataClient) GetSecuritiesListRaw(market string, pageIndex, pageSize int) ([]byte, error) {
	q := url.Values{}
	setQueryParam(q, "market", market)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)

	reqURL, _ := url.Parse(c.BaseURL)
	reqURL.Path = SecuritiesListEndpoint
	reqURL.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", reqURL.String(), nil)
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
