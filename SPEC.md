# FCData Go Client Specification

## 1. Overview

**go-fcdata** is a Go client library for the SSI FCData API (Vietnam Stock Exchange). It provides access to both REST API (historical data) and SignalR WebSocket API (real-time streaming).

### Features
- **REST API**: Access to securities, indices, OHLC, and stock price data
- **SignalR Streaming**: Real-time market data (quotes, trades, index, etc.)
- **Type-safe Models**: All data types are modeled in Go structs
- **Easy Subscription**: Helper methods for subscribing to streaming channels

### Module Information

```go
module github.com/kyimmQ/go-fcdata
go 1.24.3

require github.com/gorilla/websocket v1.5.3
```

---

## 2. REST API

### 2.1 Authentication

To use the REST API, you need a `consumerID` and `consumerSecret`. Obtain these from SSI.

**Endpoint**: `POST /api/v2/Market/AccessToken`

**Usage**:

```go
fcClient := client.NewFCDataClient("")
token, err := fcClient.Login(consumerID, consumerSecret)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Token:", token)
```

### 2.2 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GetSecuritiesList` | `GET /api/v2/Market/Securities` | List securities by market |
| `GetSecuritiesDetails` | `GET /api/v2/Market/SecuritiesDetails` | Get security details |
| `GetIndexList` | `GET /api/v2/Market/IndexList` | List indices |
| `GetIndexComponents` | `GET /api/v2/Market/IndexComponents` | Get index constituents |
| `GetDailyOhlc` | `GET /api/v2/Market/DailyOhlc` | Get daily OHLC |
| `GetIntradayOhlc` | `GET /api/v2/Market/IntradayOhlc` | Get intraday OHLC |
| `GetDailyIndex` | `GET /api/v2/Market/DailyIndex` | Get daily index |
| `GetDailyStockPrice` | `GET /api/v2/Market/DailyStockPrice` | Get daily stock price |

### 2.3 Request Parameters

All endpoints accept query parameters prefixed with `lookupRequest.`.

```go
q := url.Values{}
q.Set("lookupRequest.market", "HOSE")
q.Set("lookupRequest.pageIndex", "1")
q.Set("lookupRequest.pageSize", "10")
```

---

## 3. SignalR Streaming

### 3.1 Connection

Connect to the SignalR hub using the token obtained from REST API login.

```go
streamClient := signalr.NewClient("https://fc-datahub.ssi.com.vn/v2.0/signalr", token)

streamClient.OnConnected = func() {
    fmt.Println("Connected!")
}

streamClient.OnData = func(msg models.BroadcastMessage) {
    // Handle message
}

streamClient.OnError = func(err error) {
    fmt.Println("Error:", err)
}

if err := streamClient.StartWithLoop(); err != nil {
    log.Fatal(err)
}
```

### 3.2 Data Types

The SignalR hub supports the following data types:

| Code | Data Type | Description |
|------|-----------|-------------|
| `F` | Securities Status | Trading session and status |
| `X-QUOTE` | Order Book | Best bid/ask (10 levels) |
| `X-TRADE` | Trade Execution | Matched trades |
| `X` | Snapshot | Full market snapshot (OHLCV + order book) |
| `B` | OHLCV | Bar/Candlestick data |
| `R` | Foreign Room | Foreign investor activity |
| `MI` | Index | Index values |
| `OL` | Odd Lot | Odd lot trading data |

### 3.3 Subscription

Subscribe to data channels using helper methods or raw `SwitchChannel`.

**Helper Methods** (Recommended):

```go
// Subscribe to multiple symbols using slice
streamClient.SubscribeTrade([]string{"41I1G3000", "41I1G4000"})
streamClient.SubscribeSnapshot([]string{"41I1G3000", "41I1G4000"})
streamClient.SubscribeQuote([]string{"SSI", "VCB"})
streamClient.SubscribeIndex([]string{"VN30", "HNX30"})
streamClient.SubscribeOHLCV([]string{"SSI"})
streamClient.SubscribeForeignRoom([]string{"SSI"})
streamClient.SubscribeSecurityStatus([]string{"SSI"})
streamClient.SubscribeOddLot([]string{"SSI"})
```

**Raw SwitchChannel**:

```go
streamClient.SwitchChannel("X-TRADE:41I1G3000-41I1G4000")
```

### 3.4 Handling Data

The `BroadcastMessage` struct contains `DataType` and `Data`. The `Data` field is automatically unmarshaled into the correct type based on `DataType`.

```go
streamClient.OnData = func(msg models.BroadcastMessage) {
    switch data := msg.Data.(type) {
    case models.XQuoteData:
        fmt.Printf("Quote: %s, Bid: %.2f, Ask: %.2f\n",
            data.Symbol, data.BidPrice1, data.AskPrice1)
    case models.XTradeData:
        fmt.Printf("Trade: %s, Price: %.2f, Vol: %.2f\n",
            data.Symbol, data.LastPrice, data.LastVol)
    case models.XSnapshotData:
        fmt.Printf("Snapshot: %s, Last: %.2f\n",
            data.Symbol, data.LastPrice)
    case models.MIData:
        fmt.Printf("Index: %s, Value: %.2f\n",
            data.IndexId, data.IndexValue)
    case models.BData:
        fmt.Printf("OHLCV: %s, Close: %.2f\n",
            data.Symbol, data.Close)
    case models.RData:
        fmt.Printf("Foreign Room: %s, Room: %.2f\n",
            data.Symbol, data.CurrentRoom)
    case models.OLData:
        fmt.Printf("Odd Lot: %s, Last: %.2f\n",
            data.Symbol, data.LastPrice)
    case models.FData:
        fmt.Printf("Status: %s, Session: %s\n",
            data.Symbol, data.TradingSession)
    default:
        fmt.Printf("Unknown: %s\n", msg.Content)
    }
}
```

---

## 4. Data Models

### 4.1 REST Models

| Struct | Description |
|--------|-------------|
| `AuthRequest` | Authentication request |
| `AuthResponse` | Authentication response |
| `SecuritiesListResponse` | List of securities |
| `SecuritiesDetailsResponse` | Security details |
| `IndexListResponse` | List of indices |
| `IndexComponentsResponse` | Index constituents |
| `OHLCResponse` | OHLC data |
| `StockPriceResponse` | Stock price data |

### 4.2 Streaming Models

| Struct | DataType | Description |
|--------|----------|-------------|
| `XQuoteData` | `X-QUOTE` | Order book (10 bid/ask levels) |
| `XTradeData` | `X-TRADE` | Trade execution |
| `XSnapshotData` | `X` | Full snapshot |
| `BData` | `B` | OHLCV bar |
| `RData` | `R` | Foreign room |
| `MIData` | `MI` | Index data |
| `OLData` | `OL` | Odd lot |
| `FData` | `F` | Securities status |

---

## 5. Examples

### 5.1 Full Example

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kyimmQ/go-fcdata/client"
	"github.com/kyimmQ/go-fcdata/models"
	"github.com/kyimmQ/go-fcdata/signalr"
)

func main() {
	// Load credentials
	_ = godotenv.Load()
	consumerID := os.Getenv("CONSUMER_ID")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	// 1. REST API
	fcClient := client.NewFCDataClient("")
	token, _ := fcClient.Login(consumerID, consumerSecret)

	// Get index components
	resp, _ := fcClient.GetIndexComponents("VN30", 1, 10)
	fmt.Println("VN30 components:", resp.Data[0].TotalSymbolNo)

	// 2. Streaming
	streamClient := signalr.NewClient("https://fc-datahub.ssi.com.vn/v2.0/signalr", token)

	streamClient.OnConnected = func() {
		streamClient.SubscribeTrade([]string{"41I1G3000", "41I1G4000"})
	}

	streamClient.OnData = func(msg models.BroadcastMessage) {
		if trade, ok := msg.Data.(models.XTradeData); ok {
			fmt.Printf("Trade: %s @ %.2f\n", trade.Symbol, trade.LastPrice)
		}
	}

	streamClient.StartWithLoop()

	// Wait for exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
```

---

## 6. Configuration

### 6.1 Default Values

| Config | Default |
|--------|---------|
| Base URL | `https://fc-data.ssi.com.vn/` |
| SignalR URL | `https://fc-datahub.ssi.com.vn/v2.0/signalr` |
| HTTP Timeout | 10 seconds |
| Protocol Version | 1.3 |

### 6.2 Environment Variables

| Variable | Description |
|----------|-------------|
| `CONSUMER_ID` | API Consumer ID |
| `CONSUMER_SECRET` | API Consumer Secret |

---

## 7. Error Handling

### 7.1 REST API

Errors return as standard Go `error`. Check for specific conditions:

```go
resp, err := fcClient.GetSecuritiesList("HOSE", 1, 10)
if err != nil {
    if strings.Contains(err.Error(), "401") {
        fmt.Println("Unauthorized - check token")
    } else {
        fmt.Println("Error:", err)
    }
}
```

### 7.2 SignalR

Errors are passed to the `OnError` callback. Reconnection logic should be implemented in the callback if needed.

```go
streamClient.OnError = func(err error) {
    fmt.Println("Stream error:", err)
    // Implement reconnection logic here
}
```

---

## 8. Package Structure

```
go-fcdata/
├── client/
│   ├── fcdata.go          # Core client
│   ├── auth.go            # Authentication
│   ├── endpoints.go       # REST endpoints
│   └── debug.go          # Debug helpers
├── signalr/
│   └── client.go          # SignalR client
├── models/
│   ├── models.go          # Data structures
│   └── models_test.go     # Tests
└── example/
    └── main.go            # Usage example
```
