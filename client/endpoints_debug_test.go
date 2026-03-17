package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestRawRequestSuccess(t *testing.T) {
	client := getAuthenticatedClient(t)

	q := url.Values{}
	setQueryParam(q, "indexCode", "VN30")
	setQueryParam(q, "pageIndex", 1)
	setQueryParam(q, "pageSize", 10)
	reqURL, _ := url.Parse(client.BaseURL)
	reqURL.Path = IndexComponentsEndpoint
	reqURL.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", reqURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+client.Token)
	resp, err := client.HTTPClient.Do(req)
	if err == nil {
		raw, _ := io.ReadAll(resp.Body)
		fmt.Printf("IndexComponents Raw: %s\n", string(raw))
		resp.Body.Close()
	}

	q = url.Values{}
	setQueryParam(q, "symbol", "SSI")
	setQueryParam(q, "fromDate", "01/01/2023")
	setQueryParam(q, "toDate", "10/01/2023")
	setQueryParam(q, "pageIndex", 1)
	setQueryParam(q, "pageSize", 10)
	setQueryParam(q, "ascending", true)
	reqURL.Path = DailyOHLCEndpoint
	reqURL.RawQuery = q.Encode()
	req, _ = http.NewRequest("GET", reqURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+client.Token)
	resp, err = client.HTTPClient.Do(req)
	if err == nil {
		raw, _ := io.ReadAll(resp.Body)
		fmt.Printf("DailyOHLC Raw: %s\n", string(raw))
		resp.Body.Close()
	}
}
