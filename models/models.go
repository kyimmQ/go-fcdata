package models

import "encoding/json"

// REST Models

type AuthRequest struct {
	ConsumerID     string `json:"consumerID"`
	ConsumerSecret string `json:"consumerSecret"`
}

type AuthResponseData struct {
	AccessToken string `json:"accessToken"`
}

type AuthResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    AuthResponseData `json:"data"`
}

// Common pagination response structure
type PaginatedResponse struct {
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
}

type SecuritiesListResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	TotalRecord int    `json:"totalRecord"`
	Data        []struct {
		Symbol           string `json:"Symbol"`
		Market           string `json:"Market"`
		StockName        string `json:"StockName"`
		StockEnName      string `json:"StockEnName"`
	} `json:"data"`
}

type SecuritiesDetailsResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	TotalRecord int    `json:"totalRecord"`
	Data        []struct {
		Symbol        string `json:"Symbol"`
		Market        string `json:"Market"`
		SecType       string `json:"SecType"`
		LotSize       int    `json:"LotSize"`
		TickPrice     int    `json:"TickPrice"`
		TickIncrement int    `json:"TickIncrement"`
		// Other fields as needed
	} `json:"data"`
}

type IndexListResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	TotalRecord int    `json:"totalRecord"`
	Data        []struct {
		IndexCode string `json:"IndexCode"`
		IndexName string `json:"IndexName"`
		Exchange  string `json:"Exchange"`
	} `json:"data"`
}

type IndexComponentsResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	TotalRecord int    `json:"totalRecord"`
	Data        []struct {
		IndexCode      string `json:"IndexCode"`
		IndexName      string `json:"IndexName"`
		Exchange       string `json:"Exchange"`
		TotalSymbolNo  string `json:"TotalSymbolNo"`
		IndexComponent []struct {
			Isin        string `json:"Isin"`
			StockSymbol string `json:"StockSymbol"`
		} `json:"IndexComponent"`
	} `json:"data"`
}

// Custom types for OHLC to handle string parsing directly
// The API returns strings for Open, High, Low, Close, Volume, Value
type OHLCResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	TotalRecord int    `json:"totalRecord"`
	Data        []struct {
		Symbol      string `json:"Symbol"`
		Market      string `json:"Market"`
		TradingDate string `json:"TradingDate"`
		Open        string `json:"Open"`
		High        string `json:"High"`
		Low         string `json:"Low"`
		Close       string `json:"Close"`
		Volume      string `json:"Volume"`
		Value       string `json:"Value"`
	} `json:"data"`
}

type StockPriceResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	TotalRecord int    `json:"totalRecord"`
	Data        []struct {
		Symbol             string `json:"Symbol"`
		Market             string `json:"Market"`
		TradingDate        string `json:"TradingDate"`
		ForeignCurrentRoom string `json:"ForeignCurrentRoom"`
		TotalTradedVol     string `json:"TotalTradedVol"`
	} `json:"data"`
}

// Streaming Models

// XQuoteData represents the content of X-QUOTE messages (Order Book)
type XQuoteData struct {
	TradingDate    string  `json:"TradingDate"`
	Time          string  `json:"Time"`
	Exchange      string  `json:"Exchange"`
	Symbol        string  `json:"Symbol"`
	RType         string  `json:"RType"`
	AskPrice1     float64 `json:"AskPrice1"`
	AskPrice2     float64 `json:"AskPrice2"`
	AskPrice3     float64 `json:"AskPrice3"`
	AskPrice4     float64 `json:"AskPrice4"`
	AskPrice5     float64 `json:"AskPrice5"`
	AskPrice6     float64 `json:"AskPrice6"`
	AskPrice7     float64 `json:"AskPrice7"`
	AskPrice8     float64 `json:"AskPrice8"`
	AskPrice9     float64 `json:"AskPrice9"`
	AskPrice10    float64 `json:"AskPrice10"`
	AskVol1       float64 `json:"AskVol1"`
	AskVol2       float64 `json:"AskVol2"`
	AskVol3       float64 `json:"AskVol3"`
	AskVol4       float64 `json:"AskVol4"`
	AskVol5       float64 `json:"AskVol5"`
	AskVol6       float64 `json:"AskVol6"`
	AskVol7       float64 `json:"AskVol7"`
	AskVol8       float64 `json:"AskVol8"`
	AskVol9       float64 `json:"AskVol9"`
	AskVol10      float64 `json:"AskVol10"`
	BidPrice1     float64 `json:"BidPrice1"`
	BidPrice2     float64 `json:"BidPrice2"`
	BidPrice3     float64 `json:"BidPrice3"`
	BidPrice4     float64 `json:"BidPrice4"`
	BidPrice5     float64 `json:"BidPrice5"`
	BidPrice6     float64 `json:"BidPrice6"`
	BidPrice7     float64 `json:"BidPrice7"`
	BidPrice8     float64 `json:"BidPrice8"`
	BidPrice9     float64 `json:"BidPrice9"`
	BidPrice10    float64 `json:"BidPrice10"`
	BidVol1       float64 `json:"BidVol1"`
	BidVol2       float64 `json:"BidVol2"`
	BidVol3       float64 `json:"BidVol3"`
	BidVol4       float64 `json:"BidVol4"`
	BidVol5       float64 `json:"BidVol5"`
	BidVol6       float64 `json:"BidVol6"`
	BidVol7       float64 `json:"BidVol7"`
	BidVol8       float64 `json:"BidVol8"`
	BidVol9       float64 `json:"BidVol9"`
	BidVol10      float64 `json:"BidVol10"`
	TradingSession string  `json:"TradingSession"`
}

// XTradeData represents the content of X-TRADE messages (Trade execution)
type XTradeData struct {
	TradingDate      string  `json:"TradingDate"`
	Time             string  `json:"Time"`
	Isin             string  `json:"Isin"`
	Symbol           string  `json:"Symbol"`
	Ceiling          float64 `json:"Ceiling"`
	Floor            float64 `json:"Floor"`
	RefPrice         float64 `json:"RefPrice"`
	AvgPrice         float64 `json:"AvgPrice"`
	PriorVal         float64 `json:"PriorVal"`
	LastPrice        float64 `json:"LastPrice"`
	LastVol          float64 `json:"LastVol"`
	TotalVal         float64 `json:"TotalVal"`
	TotalVol         float64 `json:"TotalVol"`
	MarketId         string  `json:"MarketId"`
	Exchange         string  `json:"Exchange"`
	TradingSession   string  `json:"TradingSession"`
	TradingStatus    string  `json:"TradingStatus"`
	Change           float64 `json:"Change"`
	RatioChange      float64 `json:"RatioChange"`
	EstMatchedPrice  float64 `json:"EstMatchedPrice"`
	Highest          float64 `json:"Highest"`
	Lowest           float64 `json:"Lowest"`
	Side             string  `json:"Side"`
}

// BroadcastMessage represents a real-time message from the SignalR hub
type BroadcastMessage struct {
	DataType string      `json:"DataType"`
	Content  string      `json:"Content"`           // Raw JSON string
	Data     interface{} `json:"-"`                 // Unmarshalted content (XQuoteData, XTradeData, etc.)
}

// UnmarshalJSON implements custom unmarshaling for BroadcastMessage
func (b *BroadcastMessage) UnmarshalJSON(data []byte) error {
	type Alias BroadcastMessage
	aux := &struct {
		Content json.RawMessage `json:"Content"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Store raw content
	b.Content = string(aux.Content)

	// Unmarshal content based on DataType
	switch b.DataType {
	case "X-QUOTE":
		var quoteData XQuoteData
		if err := json.Unmarshal(aux.Content, &quoteData); err != nil {
			return err
		}
		b.Data = quoteData
	case "X-TRADE":
		var tradeData XTradeData
		if err := json.Unmarshal(aux.Content, &tradeData); err != nil {
			return err
		}
		b.Data = tradeData
	default:
		// For unknown types, keep as raw string
		b.Data = string(aux.Content)
	}

	return nil
}

// SignalR models
type NegotiationResponse struct {
	Url                     string `json:"Url"`
	ConnectionToken         string `json:"ConnectionToken"`
	ConnectionId            string `json:"ConnectionId"`
	KeepAliveTimeout        float64 `json:"KeepAliveTimeout"`
	DisconnectTimeout       float64 `json:"DisconnectTimeout"`
	ConnectionTimeout       float64 `json:"ConnectionTimeout"`
	TryWebSockets           bool   `json:"TryWebSockets"`
	ProtocolVersion         string `json:"ProtocolVersion"`
	TransportConnectTimeout float64 `json:"TransportConnectTimeout"`
	LongPollDelay           float64 `json:"LongPollDelay"`
}

type SignalRMessage struct {
	C string        `json:"C"`
	M []HubMessage  `json:"M"`
	I string        `json:"I"`
	E string        `json:"E"`
	R interface{}   `json:"R"`
}

type HubMessage struct {
	H string        `json:"H"`
	M string        `json:"M"`
	A []interface{} `json:"A"`
}
