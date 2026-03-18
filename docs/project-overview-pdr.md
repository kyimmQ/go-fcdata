# Project Overview and Product Development Requirements (PDR)

## 1. Project Overview

### Project Name

**go-fcdata** - Go Client for SSI FCData API

### Project Type

Go library (client library)

### Purpose

go-fcdata provides a Go client library for accessing Vietnamese stock market data through the SSI FCData API. It enables developers to integrate Vietnam Stock Exchange data into Go applications with support for both batch data retrieval via REST API and real-time streaming via SignalR.

### Target Users

- Go developers building financial applications
- Algorithmic trading systems requiring Vietnam market data
- Investment research platforms
- Financial analytics dashboards

### Repository

- GitHub: https://github.com/kyimmQ/go-fcdata

---

## 2. Product Development Requirements

### 2.1 Functional Requirements

#### Authentication

| ID     | Requirement                                           | Priority |
| ------ | ----------------------------------------------------- | -------- |
| FR-001 | Client must support consumer ID/secret authentication | Required |
| FR-002 | Client must obtain and store Bearer access token      | Required |
| FR-003 | Token must be included in subsequent API requests     | Required |

#### REST API Endpoints

| ID     | Requirement                                        | Priority |
| ------ | -------------------------------------------------- | -------- |
| FR-101 | Get securities list by market (HNX, HOSE, UPCOM)   | Required |
| FR-102 | Get securities details with trading specifications | Required |
| FR-103 | Get index list by exchange                         | Required |
| FR-104 | Get index components (e.g., VN30 constituents)     | Required |
| FR-105 | Get daily OHLC (Open, High, Low, Close) data       | Required |
| FR-106 | Get intraday OHLC data                             | Required |
| FR-107 | Get daily index values                             | Required |
| FR-108 | Get daily stock prices with foreign room data      | Required |
| FR-109 | Support pagination for all list endpoints          | Required |

#### Real-Time Streaming (SignalR)

| ID     | Requirement                               | Priority |
| ------ | ----------------------------------------- | -------- |
| FR-201 | Connect to SignalR hub via WebSocket      | Required |
| FR-202 | Authenticate with Bearer token            | Required |
| FR-203 | Subscribe to market data channels         | Required |
| FR-204 | Receive real-time broadcast messages      | Required |
| FR-205 | Handle connection errors and reconnection | Optional |

### 2.2 Non-Functional Requirements

#### Performance

| ID      | Requirement                     | Target           |
| ------- | ------------------------------- | ---------------- |
| NFR-001 | HTTP request timeout            | 10 seconds       |
| NFR-002 | Support concurrent API requests | Yes              |
| NFR-003 | Efficient JSON parsing          | Standard library |

#### Reliability

| ID      | Requirement                           | Priority |
| ------- | ------------------------------------- | -------- |
| NFR-101 | Handle HTTP error status codes        | Required |
| NFR-102 | Proper resource cleanup (defer close) | Required |
| NFR-103 | Error messages include status code    | Required |

#### Compatibility

| ID      | Requirement                             | Priority |
| ------- | --------------------------------------- | -------- |
| NFR-201 | Go version 1.24+                        | Required |
| NFR-202 | Support HTTPS only                      | Required |
| NFR-203 | WebSocket support via gorilla/websocket | Required |

---

## 3. API Endpoints Reference

### Base URLs

- REST API: `https://fc-data.ssi.com.vn/`
- SignalR Hub: `https://fc-datahub.ssi.com.vn/v2.0/signalr`

### Authentication

```
POST /api/v2/Market/AccessToken
Body: { "consumerID": "...", "consumerSecret": "..." }
```

### Market Data

| Endpoint                           | Method | Description       |
| ---------------------------------- | ------ | ----------------- |
| `/api/v2/Market/Securities`        | GET    | Securities list   |
| `/api/v2/Market/SecuritiesDetails` | GET    | Security details  |
| `/api/v2/Market/IndexList`         | GET    | Index list        |
| `/api/v2/Market/IndexComponents`   | GET    | Index components  |
| `/api/v2/Market/DailyOhlc`         | GET    | Daily OHLC        |
| `/api/v2/Market/IntradayOhlc`      | GET    | Intraday OHLC     |
| `/api/v2/Market/DailyIndex`        | GET    | Daily index       |
| `/api/v2/Market/DailyStockPrice`   | GET    | Daily stock price |

### SignalR

| Method           | Description               |
| ---------------- | ------------------------- |
| `/negotiate`     | Get connection token      |
| `/connect`       | WebSocket connection      |
| `/start`         | Start SignalR session     |
| `SwitchChannels` | Subscribe to data channel |

---

## 4. Data Models

### Authentication

- `AuthRequest` - Consumer ID and secret
- `AuthResponse` - Access token response
- `AuthResponseData` - Token data container

### Market Data

- `SecuritiesListResponse` - Security list with symbol, market, names
- `SecuritiesDetailsResponse` - Detailed security info (lot size, tick price)
- `IndexListResponse` - Index list with code and name
- `IndexComponentsResponse` - Index constituents
- `OHLCResponse` - OHLC price data
- `StockPriceResponse` - Stock prices with foreign room

### Streaming

- `BroadcastMessage` - Real-time data message (with auto-unmarshaling)
- `NegotiationResponse` - SignalR connection details
- `SignalRMessage` - SignalR protocol message
- `HubMessage` - Hub method invocation
- `XQuoteData` - Order book data (X-QUOTE)
- `XTradeData` - Trade execution data (X-TRADE)
- `FData` - Securities status (F)
- `XSnapshotData` - Snapshot data (X)
- `BData` - OHLCV data (B)
- `RData` - Foreign room data (R)
- `MIData` - Index data (MI)
- `OLData` - Odd lot data (OL)

---

## 5. Acceptance Criteria

### Authentication

- [ ] Login with valid credentials returns access token
- [ ] Login with invalid credentials returns error
- [ ] Token is stored and used in subsequent requests

### REST API

- [ ] GetSecuritiesList returns paginated security list
- [ ] GetIndexComponents returns VN30 (or other index) constituents
- [ ] GetDailyOhlc returns historical price data
- [ ] All endpoints properly handle errors

### SignalR

- [ ] Client connects to SignalR hub successfully
- [ ] Authentication via Bearer token works
- [ ] SwitchChannel subscribes to data streams
- [ ] OnData callback receives real-time messages
- [ ] Client handles disconnection gracefully

### Documentation

- [ ] README provides quick start guide
- [ ] Code is commented for godoc
- [ ] All public types and functions have documentation
