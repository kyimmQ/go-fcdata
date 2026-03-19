package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Custom type to handle unmarshaling float64 that can be "NaN" or numeric string
type Float64FromString float64

func (f *Float64FromString) UnmarshalJSON(data []byte) error {
	// Remove quotes if present
	s := string(data)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}

	// Handle non-numeric values (like "NaN") by defaulting to 0
	if s == "NaN" || s == "" {
		*f = 0.0
		return nil
	}

	// Try to parse as float64
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		// Default to 0 on error
		*f = 0.0
		return nil
	}

	*f = Float64FromString(val)
	return nil
}

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
	TradingDate    string            `json:"TradingDate"`
	Time           string            `json:"Time"`
	Exchange       string            `json:"Exchange"`
	Symbol         string            `json:"Symbol"`
	RType          string            `json:"RType"`
	AskPrice1      Float64FromString `json:"AskPrice1"`
	AskPrice2      Float64FromString `json:"AskPrice2"`
	AskPrice3      Float64FromString `json:"AskPrice3"`
	AskPrice4      Float64FromString `json:"AskPrice4"`
	AskPrice5      Float64FromString `json:"AskPrice5"`
	AskPrice6      Float64FromString `json:"AskPrice6"`
	AskPrice7      Float64FromString `json:"AskPrice7"`
	AskPrice8      Float64FromString `json:"AskPrice8"`
	AskPrice9      Float64FromString `json:"AskPrice9"`
	AskPrice10     Float64FromString `json:"AskPrice10"`
	AskVol1        Float64FromString `json:"AskVol1"`
	AskVol2        Float64FromString `json:"AskVol2"`
	AskVol3        Float64FromString `json:"AskVol3"`
	AskVol4        Float64FromString `json:"AskVol4"`
	AskVol5        Float64FromString `json:"AskVol5"`
	AskVol6        Float64FromString `json:"AskVol6"`
	AskVol7        Float64FromString `json:"AskVol7"`
	AskVol8        Float64FromString `json:"AskVol8"`
	AskVol9        Float64FromString `json:"AskVol9"`
	AskVol10       Float64FromString `json:"AskVol10"`
	BidPrice1      Float64FromString `json:"BidPrice1"`
	BidPrice2      Float64FromString `json:"BidPrice2"`
	BidPrice3      Float64FromString `json:"BidPrice3"`
	BidPrice4      Float64FromString `json:"BidPrice4"`
	BidPrice5      Float64FromString `json:"BidPrice5"`
	BidPrice6      Float64FromString `json:"BidPrice6"`
	BidPrice7      Float64FromString `json:"BidPrice7"`
	BidPrice8      Float64FromString `json:"BidPrice8"`
	BidPrice9      Float64FromString `json:"BidPrice9"`
	BidPrice10     Float64FromString `json:"BidPrice10"`
	BidVol1        Float64FromString `json:"BidVol1"`
	BidVol2        Float64FromString `json:"BidVol2"`
	BidVol3        Float64FromString `json:"BidVol3"`
	BidVol4        Float64FromString `json:"BidVol4"`
	BidVol5        Float64FromString `json:"BidVol5"`
	BidVol6        Float64FromString `json:"BidVol6"`
	BidVol7        Float64FromString `json:"BidVol7"`
	BidVol8        Float64FromString `json:"BidVol8"`
	BidVol9        Float64FromString `json:"BidVol9"`
	BidVol10       Float64FromString `json:"BidVol10"`
	TradingSession string            `json:"TradingSession"`
}

// XTradeData represents the content of X-TRADE messages (Trade execution)
type XTradeData struct {
	TradingDate      string            `json:"TradingDate"`
	Time             string            `json:"Time"`
	Isin             string            `json:"Isin"`
	Symbol           string            `json:"Symbol"`
	Ceiling          Float64FromString `json:"Ceiling"`
	Floor            Float64FromString `json:"Floor"`
	RefPrice         Float64FromString `json:"RefPrice"`
	AvgPrice         Float64FromString `json:"AvgPrice"`
	PriorVal         Float64FromString `json:"PriorVal"`
	LastPrice        Float64FromString `json:"LastPrice"`
	LastVol          Float64FromString `json:"LastVol"`
	TotalVal         Float64FromString `json:"TotalVal"`
	TotalVol         Float64FromString `json:"TotalVol"`
	MarketId         string            `json:"MarketId"`
	Exchange         string            `json:"Exchange"`
	TradingSession   string            `json:"TradingSession"`
	TradingStatus    string            `json:"TradingStatus"`
	Change           Float64FromString `json:"Change"`
	RatioChange      Float64FromString `json:"RatioChange"`
	EstMatchedPrice  Float64FromString `json:"EstMatchedPrice"`
	Highest          Float64FromString `json:"Highest"`
	Lowest           Float64FromString `json:"Lowest"`
	Side             string            `json:"Side"`
}

// FData represents securities status (F)
type FData struct {
	RType         string `json:"RType"`
	MarketId      string `json:"MarketId"`
	TradingDate   string `json:"TradingDate"`
	Time          string `json:"Time"`
	Symbol        string `json:"Symbol"`
	TradingSession string `json:"TradingSession"`
	TradingStatus string `json:"TradingStatus"`
	Exchange      string `json:"Exchange"`
}

// XSnapshotData represents the snapshot (X)
type XSnapshotData struct {
	RType             string            `json:"RType"`
	TradingDate       string            `json:"TradingDate"`
	Time              string            `json:"Time"`
	Isin              string            `json:"Isin"`
	Symbol            string            `json:"Symbol"`
	Ceiling           Float64FromString `json:"Ceiling"`
	Floor             Float64FromString `json:"Floor"`
	RefPrice          Float64FromString `json:"RefPrice"`
	Open              Float64FromString `json:"Open"`
	Close             Float64FromString `json:"Close"`
	High              Float64FromString `json:"High"`
	Low               Float64FromString `json:"Low"`
	AvgPrice          Float64FromString `json:"AvgPrice"`
	PriorVal          Float64FromString `json:"PriorVal"`
	LastPrice         Float64FromString `json:"LastPrice"`
	Change            Float64FromString `json:"Change"`
	RatioChange       Float64FromString `json:"RatioChange"`
	EstMatchedPrice   Float64FromString `json:"EstMatchedPrice"`
	LastVol           Float64FromString `json:"LastVol"`
	TotalVal          Float64FromString `json:"TotalVal"`
	TotalVol          Float64FromString `json:"TotalVol"`
	BidPrice1         Float64FromString `json:"BidPrice1"`
	BidVol1           Float64FromString `json:"BidVol1"`
	BidPrice2         Float64FromString `json:"BidPrice2"`
	BidVol2           Float64FromString `json:"BidVol2"`
	BidPrice3         Float64FromString `json:"BidPrice3"`
	BidVol3           Float64FromString `json:"BidVol3"`
	BidPrice4         Float64FromString `json:"BidPrice4"`
	BidVol4           Float64FromString `json:"BidVol4"`
	BidPrice5         Float64FromString `json:"BidPrice5"`
	BidVol5           Float64FromString `json:"BidVol5"`
	BidPrice6         Float64FromString `json:"BidPrice6"`
	BidVol6           Float64FromString `json:"BidVol6"`
	BidPrice7         Float64FromString `json:"BidPrice7"`
	BidVol7           Float64FromString `json:"BidVol7"`
	BidPrice8         Float64FromString `json:"BidPrice8"`
	BidVol8           Float64FromString `json:"BidVol8"`
	BidPrice9         Float64FromString `json:"BidPrice9"`
	BidVol9           Float64FromString `json:"BidVol9"`
	BidPrice10        Float64FromString `json:"BidPrice10"`
	BidVol10          Float64FromString `json:"BidVol10"`
	AskPrice1         Float64FromString `json:"AskPrice1"`
	AskVol1           Float64FromString `json:"AskVol1"`
	AskPrice2         Float64FromString `json:"AskPrice2"`
	AskVol2           Float64FromString `json:"AskVol2"`
	AskPrice3         Float64FromString `json:"AskPrice3"`
	AskVol3           Float64FromString `json:"AskVol3"`
	AskPrice4         Float64FromString `json:"AskPrice4"`
	AskVol4           Float64FromString `json:"AskVol4"`
	AskPrice5         Float64FromString `json:"AskPrice5"`
	AskVol5           Float64FromString `json:"AskVol5"`
	AskPrice6         Float64FromString `json:"AskPrice6"`
	AskVol6           Float64FromString `json:"AskVol6"`
	AskPrice7         Float64FromString `json:"AskPrice7"`
	AskVol7           Float64FromString `json:"AskVol7"`
	AskPrice8         Float64FromString `json:"AskPrice8"`
	AskVol8           Float64FromString `json:"AskVol8"`
	AskPrice9         Float64FromString `json:"AskPrice9"`
	AskVol9           Float64FromString `json:"AskVol9"`
	AskPrice10        Float64FromString `json:"AskPrice10"`
	AskVol10          Float64FromString `json:"AskVol10"`
	MarketId          string            `json:"MarketId"`
	Exchange          string            `json:"Exchange"`
	TradingSession    string            `json:"TradingSession"`
	TradingStatus     string            `json:"TradingStatus"`
}

// BData represents OHLCV (B)
type BData struct {
	RType       string            `json:"RType"`
	Symbol      string            `json:"Symbol"`
	TradingTime string            `json:"TradingTime"` // "Time" in docs but "TradingTime" in sample
	Open        Float64FromString `json:"Open"`
	High        Float64FromString `json:"High"`
	Low         Float64FromString `json:"Low"`
	Close       Float64FromString `json:"Close"`
	Volume      Float64FromString `json:"Volume"`
	Value       Float64FromString `json:"Value"`
}

// RData represents Foreign Room (R)
type RData struct {
	RType       string            `json:"RType"`
	TradingDate string            `json:"TradingDate"`
	Time        string            `json:"Time"`
	Isin        string            `json:"Isin"`
	Symbol      string            `json:"Symbol"`
	TotalRoom   Float64FromString `json:"TotalRoom"`
	CurrentRoom Float64FromString `json:"CurrentRoom"`
	BuyVol      Float64FromString `json:"BuyVol"`
	SellVol     Float64FromString `json:"SellVol"`
	BuyVal      Float64FromString `json:"BuyVal"`
	SellVal     Float64FromString `json:"SellVal"`
	MarketId    string            `json:"MarketId"`
	Exchange    string            `json:"Exchange"`
}

// MIData represents Index data (MI)
type MIData struct {
	RType             string            `json:"RType"`
	IndexId           string            `json:"IndexId"`
	IndexValEst       Float64FromString `json:"IndexValEst"`
	IndexValue        Float64FromString `json:"IndexValue"`
	PriorIndexValue   Float64FromString `json:"PriorIndexValue"` // Note: sample has "PriorIndexValue" but doc says "Time". Checking sample.
	TradingDate       string            `json:"TradingDate"`
	Time              string            `json:"Time"`
	Change            Float64FromString `json:"Change"`
	RatioChange       Float64FromString `json:"RatioChange"`
	TotalTrade        Float64FromString `json:"TotalTrade"`
	TotalQtty         Float64FromString `json:"TotalQtty"`
	TotalValue        Float64FromString `json:"TotalValue"`
	IndexName         string            `json:"IndexName"`
	Advances          int               `json:"Advances"`
	NoChanges         int               `json:"NoChanges"`
	Declines          int               `json:"Declines"`
	Ceiling           int               `json:"Ceiling"`
	Floor             int               `json:"Floor"`
	TotalQttyPt       Float64FromString `json:"TotalQttyPt"`
	TotalValuePt      Float64FromString `json:"TotalValuePt"`
	Exchange          string            `json:"Exchange"`
	AllQty            Float64FromString `json:"AllQty"`
	AllValue          Float64FromString `json:"AllValue"`
	IndexType         string            `json:"IndexType"` // "TypeIndex" in doc, "IndexType" in sample
	TradingSession    string            `json:"TradingSession"`
	MarketId          string            `json:"MarketId"`
	TotalQttyOd       Float64FromString `json:"TotalQttyOd"`
	TotalValueOd      Float64FromString `json:"TotalValueOd"`
}

// OLData represents Odd lot message (OL)
type OLData struct {
	RType          string            `json:"RType"`
	TradingDate    string            `json:"TradingDate"`
	Time           string            `json:"Time"`
	StockNo        int               `json:"StockNo"`
	Symbol         string            `json:"Symbol"`
	Ceiling        Float64FromString `json:"Ceiling"`
	Floor          Float64FromString `json:"Floor"`
	RefPrice       Float64FromString `json:"RefPrice"`
	Open           Float64FromString `json:"Open"`
	High           Float64FromString `json:"High"`
	Low            Float64FromString `json:"Low"`
	LastPrice      Float64FromString `json:"LastPrice"`
	LastVol        Float64FromString `json:"LastVol"`
	TotalVal       Float64FromString `json:"TotalVal"`
	TotalVol       Float64FromString `json:"TotalVol"`
	BidPrice1      Float64FromString `json:"BidPrice1"`
	BidPrice2      Float64FromString `json:"BidPrice2"`
	BidPrice3      Float64FromString `json:"BidPrice3"`
	BidVol1        Float64FromString `json:"BidVol1"`
	BidVol2        Float64FromString `json:"BidVol2"`
	BidVol3        Float64FromString `json:"BidVol3"`
	AskPrice1      Float64FromString `json:"AskPrice1"`
	AskPrice2      Float64FromString `json:"AskPrice2"`
	AskPrice3      Float64FromString `json:"AskPrice3"`
	AskVol1        Float64FromString `json:"AskVol1"`
	AskVol2        Float64FromString `json:"AskVol2"`
	AskVol3        Float64FromString `json:"AskVol3"`
	Exchange       string            `json:"Exchange"`
	TradingSession string            `json:"TradingSession"`
	TradingStatus  string            `json:"TradingStatus"`
	Change         Float64FromString `json:"Change"`
	RatioChange    Float64FromString `json:"RatioChange"`
	StockType      string            `json:"StockType"`
}

// BroadcastMessage represents a real-time message from the SignalR hub
type BroadcastMessage struct {
	DataType string      `json:"DataType"`
	Content  string      `json:"Content"`           // Raw JSON string
	Data     interface{} `json:"-"`                // Unmarshalled content (XQuoteData, XTradeData, etc.)
}

// UnmarshalJSON implements custom unmarshaling for BroadcastMessage
func (b *BroadcastMessage) UnmarshalJSON(data []byte) error {
	// First, unmarshal into a map to extract DataType
	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	// Extract DataType
	if dt, ok := tmp["DataType"].(string); ok {
		b.DataType = dt
	}

	// Extract Content
	if content, ok := tmp["Content"].(string); ok {
		b.Content = content

		// Unmarshal content based on DataType
		contentBytes := []byte(content)
		switch b.DataType {
		case "X-QUOTE":
			var quoteData XQuoteData
			if err := json.Unmarshal(contentBytes, &quoteData); err != nil {
				return fmt.Errorf("failed to unmarshal X-QUOTE: %w", err)
			}
			b.Data = quoteData
		case "X-TRADE":
			var tradeData XTradeData
			if err := json.Unmarshal(contentBytes, &tradeData); err != nil {
				return fmt.Errorf("failed to unmarshal X-TRADE: %w", err)
			}
			b.Data = tradeData
		case "F":
			var fData FData
			if err := json.Unmarshal(contentBytes, &fData); err != nil {
				return fmt.Errorf("failed to unmarshal F: %w", err)
			}
			b.Data = fData
		case "X":
			var xData XSnapshotData
			if err := json.Unmarshal(contentBytes, &xData); err != nil {
				return fmt.Errorf("failed to unmarshal X: %w", err)
			}
			b.Data = xData
		case "B":
			var bData BData
			if err := json.Unmarshal(contentBytes, &bData); err != nil {
				return fmt.Errorf("failed to unmarshal B: %w", err)
			}
			b.Data = bData
		case "R":
			var rData RData
			if err := json.Unmarshal(contentBytes, &rData); err != nil {
				return fmt.Errorf("failed to unmarshal R: %w", err)
			}
			b.Data = rData
		case "MI":
			var miData MIData
			if err := json.Unmarshal(contentBytes, &miData); err != nil {
				return fmt.Errorf("failed to unmarshal MI: %w", err)
			}
			b.Data = miData
		case "OL":
			var olData OLData
			if err := json.Unmarshal(contentBytes, &olData); err != nil {
				return fmt.Errorf("failed to unmarshal OL: %w", err)
			}
			b.Data = olData
		default:
			// For unknown types, keep as raw string
			b.Data = content
		}
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
