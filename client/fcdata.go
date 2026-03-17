package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

)

const (
	DefaultBaseURL = "https://fc-data.ssi.com.vn/"
)

type FCDataClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func NewFCDataClient(baseURL string) *FCDataClient {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return &FCDataClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *FCDataClient) SetToken(token string) {
	c.Token = token
}

func (c *FCDataClient) doRequest(method, endpoint string, queryParams url.Values, body interface{}, response interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	reqURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}
	reqURL.Path = endpoint
	if queryParams != nil {
		reqURL.RawQuery = queryParams.Encode()
	}

	req, err := http.NewRequest(method, reqURL.String(), bodyReader)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	if response != nil {
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return err
		}
	}

	return nil
}
