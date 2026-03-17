# Codebase Summary

## Overview

go-fcdata is a Go client library for the SSI FCData API (Vietnam Stock Exchange). It consists of three main packages: `client` for REST API, `signalr` for real-time streaming, and `models` for data structures.

## Module Information

```go
module github.com/kyimmQ/go_fcdata
go 1.24.3

require github.com/gorilla/websocket v1.5.3
```

## Package Structure

### `client/` Package

The `client` package provides REST API access to FCData services.

#### `client/fcdata.go`
- **FCDataClient**: Main struct with `BaseURL`, `HTTPClient`, and `Token` fields
- **NewFCDataClient(baseURL string)**: Constructor with default base URL fallback
- **SetToken(token string)**: Sets authentication token
- **doRequest()**: Core method for making HTTP requests with JSON marshaling/unmarshaling

#### `client/auth.go`
- **Login(consumerID, consumerSecret string)**: Authenticates and retrieves access token
- Returns token string or error
- Sets token on client automatically

#### `client/endpoints.go`
Market data endpoint methods:

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GetSecuritiesList` | `/api/v2/Market/Securities` | List securities by market |
| `GetSecuritiesDetails` | `/api/v2/Market/SecuritiesDetails` | Security details |
| `GetIndexList` | `/api/v2/Market/IndexList` | List indices |
| `GetIndexComponents` | `/api/v2/Market/IndexComponents` | Index constituents |
| `GetDailyOhlc` | `/api/v2/Market/DailyOhlc` | Daily OHLC prices |
| `GetIntradayOhlc` | `/api/v2/Market/IntradayOhlc` | Intraday OHLC |
| `GetDailyIndex` | `/api/v2/Market/DailyIndex` | Daily index values |
| `GetDailyStockPrice` | `/api/v2/Market/DailyStockPrice` | Stock prices |

#### `client/debug.go`
- **GetSecuritiesListRaw()**: Returns raw bytes for debugging

### `signalr/` Package

The `signalr` package provides WebSocket-based real-time data streaming.

#### `signalr/client.go`
- **Client**: Main struct with connection state and callbacks
- **NewClient(baseURL, token string)**: Constructor
- **Start()**: Establishes SignalR connection
- **StartWithLoop()**: Starts connection with background read loop
- **SwitchChannel(channel string)**: Subscribes to data channel
- **Invoke()**: Generic method invocation on hub
- **Close()**: Closes connection

**Callbacks:**
- `OnData func(BroadcastMessage)`: Received data handler
- `OnConnected func()`: Connection established handler
- `OnError func(error)`: Error handler

### `models/` Package

Shared data structures for REST API and SignalR communication.

#### REST Models
- `AuthRequest` / `AuthResponse` / `AuthResponseData`: Authentication
- `SecuritiesListResponse`: Securities list with Symbol, Market, StockName
- `SecuritiesDetailsResponse`: Security details with SecType, LotSize, TickPrice
- `IndexListResponse`: Index list with IndexCode, IndexName, Exchange
- `IndexComponentsResponse`: Index constituents
- `OHLCResponse`: OHLC data (Open, High, Low, Close, Volume, Value)
- `StockPriceResponse`: Stock prices with ForeignCurrentRoom

#### Streaming Models
- `BroadcastMessage`: Real-time message with DataType and Content
- `NegotiationResponse`: SignalR handshake response
- `SignalRMessage` / `HubMessage`: SignalR protocol messages

### `example/` Package

Demonstration application in `example/main.go`:
1. Creates REST client and authenticates
2. Fetches index components via REST
3. Connects to SignalR for real-time data
4. Subscribes to channels (X-QUOTE:SSI, MI:VN30)
5. Handles signals for graceful shutdown

---

## Key Implementation Details

### HTTP Client Configuration
- Default timeout: 10 seconds
- Bearer token authentication
- JSON content type for requests

### SignalR Protocol
- Protocol version: 1.3
- Hub name: FcMarketDataV2Hub
- Transport: WebSocket
- Authentication: Bearer token in header

### Query Parameter Format
Uses `lookupRequest.{paramName}` prefix for query parameters:
```go
q.Set("lookupRequest.market", "HO")
q.Set("lookupRequest.pageIndex", "1")
```

### Channel Subscription
SignalR channels follow format:
- `X-QUOTE:{symbol}` - Quote data (e.g., X-QUOTE:SSI)
- `MI:{indexCode}` - Index data (e.g., MI:VN30)

---

## Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| gorilla/websocket | v1.5.3 | WebSocket client for SignalR |

---

## File Listing

```
go-fcdata/
├── go.mod                          # Module definition
├── go.sum                          # Dependency checksums
├── .env                            # Environment variables (example)
├── .gitignore                      # Git ignore patterns
├── client/
│   ├── fcdata.go                   # Core client (89 lines)
│   ├── auth.go                     # Authentication (29 lines)
│   ├── endpoints.go                # API endpoints (145 lines)
│   ├── debug.go                    # Debug helpers (36 lines)
│   ├── client_test.go              # Client tests
│   └── endpoints_debug_test.go     # Endpoint tests
├── signalr/
│   ├── client.go                   # SignalR client (262 lines)
│   └── client_test.go              # SignalR tests
├── models/
│   └── models.go                   # Data models (146 lines)
├── example/
│   └── main.go                     # Example application (73 lines)
└── docs/                           # Documentation
    └── *.md                        # Various documentation files
```
