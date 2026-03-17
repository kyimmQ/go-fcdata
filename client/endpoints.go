package client

import (
	"fmt"
	"net/url"

	"github.com/kyimmQ/go_fcdata/models"
)

const (
	SecuritiesListEndpoint    = "api/v2/Market/Securities"
	SecuritiesDetailsEndpoint = "api/v2/Market/SecuritiesDetails"
	IndexListEndpoint         = "api/v2/Market/IndexList"
	IndexComponentsEndpoint   = "api/v2/Market/IndexComponents"
	DailyOHLCEndpoint         = "api/v2/Market/DailyOhlc"
	IntradayOHLCEndpoint      = "api/v2/Market/IntradayOhlc"
	DailyIndexEndpoint        = "api/v2/Market/DailyIndex"
	DailyStockPriceEndpoint   = "api/v2/Market/DailyStockPrice"
)

func setQueryParam(q url.Values, paramName string, value interface{}) {
	q.Set(fmt.Sprintf("lookupRequest.%s", paramName), fmt.Sprintf("%v", value))
}

func (c *FCDataClient) GetSecuritiesList(market string, pageIndex, pageSize int) (*models.SecuritiesListResponse, error) {
	q := url.Values{}
	setQueryParam(q, "market", market)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)

	var result models.SecuritiesListResponse
	if err := c.doRequest("GET", SecuritiesListEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *FCDataClient) GetSecuritiesDetails(market, symbol string, pageIndex, pageSize int) (*models.SecuritiesDetailsResponse, error) {
	q := url.Values{}
	setQueryParam(q, "market", market)
	if symbol != "" {
		setQueryParam(q, "symbol", symbol)
	}
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)

	var result models.SecuritiesDetailsResponse
	if err := c.doRequest("GET", SecuritiesDetailsEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *FCDataClient) GetIndexList(exchange string, pageIndex, pageSize int) (*models.IndexListResponse, error) {
	q := url.Values{}
	setQueryParam(q, "exchange", exchange)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)

	var result models.IndexListResponse
	if err := c.doRequest("GET", IndexListEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *FCDataClient) GetIndexComponents(indexCode string, pageIndex, pageSize int) (*models.IndexComponentsResponse, error) {
	q := url.Values{}
	setQueryParam(q, "indexCode", indexCode)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)

	var result models.IndexComponentsResponse
	if err := c.doRequest("GET", IndexComponentsEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *FCDataClient) GetDailyOhlc(symbol, fromDate, toDate string, pageIndex, pageSize int, ascending bool) (*models.OHLCResponse, error) {
	q := url.Values{}
	setQueryParam(q, "symbol", symbol)
	setQueryParam(q, "fromDate", fromDate)
	setQueryParam(q, "toDate", toDate)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)
	setQueryParam(q, "ascending", ascending)

	var result models.OHLCResponse
	if err := c.doRequest("GET", DailyOHLCEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *FCDataClient) GetIntradayOhlc(symbol, fromDate, toDate string, pageIndex, pageSize int, ascending bool) (*models.OHLCResponse, error) {
	q := url.Values{}
	setQueryParam(q, "symbol", symbol)
	setQueryParam(q, "fromDate", fromDate)
	setQueryParam(q, "toDate", toDate)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)
	setQueryParam(q, "ascending", ascending)

	var result models.OHLCResponse
	if err := c.doRequest("GET", IntradayOHLCEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Reusing OHLCResponse since the original node example does the same or similar for DailyIndex.
func (c *FCDataClient) GetDailyIndex(indexId, fromDate, toDate string, pageIndex, pageSize int, ascending bool) (*models.OHLCResponse, error) {
	q := url.Values{}
	setQueryParam(q, "indexId", indexId)
	setQueryParam(q, "fromDate", fromDate)
	setQueryParam(q, "toDate", toDate)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)
	setQueryParam(q, "ascending", ascending)

	var result models.OHLCResponse
	if err := c.doRequest("GET", DailyIndexEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *FCDataClient) GetDailyStockPrice(symbol, market, fromDate, toDate string, pageIndex, pageSize int) (*models.StockPriceResponse, error) {
	q := url.Values{}
	setQueryParam(q, "symbol", symbol)
	if market != "" {
		setQueryParam(q, "market", market)
	}
	setQueryParam(q, "fromDate", fromDate)
	setQueryParam(q, "toDate", toDate)
	setQueryParam(q, "pageIndex", pageIndex)
	setQueryParam(q, "pageSize", pageSize)

	var result models.StockPriceResponse
	if err := c.doRequest("GET", DailyStockPriceEndpoint, q, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
