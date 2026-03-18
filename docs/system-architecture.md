# System Architecture

## Overview

go-fcdata is designed with a clean separation between REST API access and real-time streaming. The architecture consists of three main components: the REST client, the SignalR streaming client, and shared data models.

## High-Level Architecture

```
                    +------------------+
                    |   go-fcdata      |
                    |    Application   |
                    +--------+---------+
                             |
        +--------------------+--------------------+
        |                                         |
        v                                         v
+---------------+                         +----------------+
|  REST Client  |                         | SignalR Client |
|   (client/)   |                         |   (signalr/)  |
+---------------+                         +----------------+
        |                                         |
        | HTTP + JSON                            | WebSocket
        v                                         v
+------------------+                    +------------------+
|  FCData REST API |                    |  FCData SignalR  |
| fc-data.ssi.com.vn|                   |fc-datahub.ssi...|
+------------------+                    +------------------+
```

## Component Diagram

### 1. REST API Client (`client/`)

```
client/
├── fcdata.go           Core client with HTTP transport
├── auth.go             Authentication (Login)
├── endpoints.go        Market data endpoints
└── debug.go           Debug/raw request helpers
```

**Responsibilities:**
- HTTP request/response handling
- JSON serialization/deserialization
- Authentication token management
- Error handling and status code checking

**Data Flow:**
```
Application
    |
    v
FCDataClient.Login() --> POST /api/v2/Market/AccessToken
    |                      |
    v                      v
FCDataClient.GetXxx() --> GET /api/v2/Market/Xxx
    |                      |
    v                      v
Response<T> <---------- JSON
```

### 2. SignalR Streaming Client (`signalr/`)

```
signalr/
└── client.go    WebSocket client with SignalR protocol and subscription helpers
```

**Responsibilities:**
- SignalR protocol handshake (negotiate, connect, start)
- WebSocket connection management
- Message parsing and dispatching
- Channel subscription (generic and typed helpers)

**Connection Flow:**
```
1. Client.NewClient(url, token)
2. Client.StartWithLoop()
   a. negotiate() --> POST /negotiate
      Returns: ConnectionToken
   b. connect()  --> WS /connect?token=...
      Establishes WebSocket
   c. start()    --> GET /start
      Activates SignalR
   d. readLoop() --> Background goroutine
      Reads and dispatches messages
3. Client.SwitchChannel("X-QUOTE:SSI")
   Invokes hub method to subscribe
4. Messages delivered via OnData callback
```

### 3. Data Models (`models/`)

```
models/
└── models.go    All request/response structures
```

**Responsibilities:**
- Define data structures for API communication
- JSON tags for serialization
- Shared types between REST and SignalR

---

## Data Flow Diagrams

### Authentication Flow

```
User Application
       |
       v
fcClient.Login(consumerID, consumerSecret)
       |
       v
+--------------------------+
| Construct AuthRequest    |
| {consumerID, consumerSecret} |
+--------------------------+
       |
       v
+--------------------------+
| POST /api/v2/Market/     |
| AccessToken              |
| Content-Type: application/json
+--------------------------+
       |
       v
FCData Server
       |
       v
+--------------------------+
| AuthResponse             |
| {status, message, data: |
|   {accessToken}}         |
+--------------------------+
       |
       v
Store token in fcClient.Token
       |
       v
Return token to caller
```

### Real-Time Data Flow

```
SignalR Client (Connected)
       |
       v
Client.SwitchChannel("X-QUOTE:SSI")
       |
       v
+--------------------------+
| Invoke:                  |
| Hub: FcMarketDataV2Hub   |
| Method: SwitchChannels  |
| Args: ["X-QUOTE:SSI"]    |
+--------------------------+
       |
       v
Server pushes messages...
       |
       v
readLoop() receives message
       |
       v
+--------------------------+
| Parse SignalRMessage     |
| {M: [{H: "...", M: "Broadcast", |
|    A: [...]}]}            |
+--------------------------+
       |
       v
+--------------------------+
| Parse BroadcastMessage   |
| {DataType, Content}      |
+--------------------------+
       |
       v
OnData callback invoked
```

---

## API Endpoints

### REST API (fc-data.ssi.com.vn)

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/api/v2/Market/AccessToken` | POST | Authentication |
| `/api/v2/Market/Securities` | GET | List securities |
| `/api/v2/Market/SecuritiesDetails` | GET | Security details |
| `/api/v2/Market/IndexList` | GET | List indices |
| `/api/v2/Market/IndexComponents` | GET | Index constituents |
| `/api/v2/Market/DailyOhlc` | GET | Daily OHLC |
| `/api/v2/Market/IntradayOhlc` | GET | Intraday OHLC |
| `/api/v2/Market/DailyIndex` | GET | Daily index |
| `/api/v2/Market/DailyStockPrice` | GET | Daily stock price |

### SignalR Hub (fc-datahub.ssi.com.vn)

| Path | Method | Purpose |
|------|--------|---------|
| `/negotiate` | GET | Get connection token |
| `/connect` | WS | WebSocket connection |
| `/start` | GET | Start session |
| Hub: `FcMarketDataV2Hub` | | |
| - `SwitchChannels` | Invoke | Subscribe to channels |

#### Subscription Channels

| Channel    | Data Type | Description |
|------------|-----------|-------------|
| `F:{symbols}` | Securities Status | Trading status (HALT, OPEN, etc.) |
| `X-QUOTE:{symbols}` | Order Book | Best 10 bid/ask prices and volumes |
| `X-TRADE:{symbols}` | Trade Execution | Matched trades with price and volume |
| `X:{symbols}` | Snapshot | Full market snapshot (OHLCV + order book) |
| `B:{symbols}` | OHLCV | Bar/Candlestick data |
| `R:{symbols}` | Foreign Room | Foreign investor trading activity |
| `MI:{indexes}` | Index | Index values (VN30, VN100, etc.) |
| `OL:{symbols}` | Odd Lot | Odd lot trading data |

#### Subscription Helper Methods

The SignalR client provides typed subscription methods:

```go
// Subscribe to different data types
client.SubscribeSecurityStatus("SSI,VCB")  // F channel
client.SubscribeQuote("SSI,VCB")           // X-QUOTE channel
client.SubscribeTrade("SSI,VCB")            // X-TRADE channel
client.SubscribeSnapshot("SSI,VCB")        // X channel
client.SubscribeOHLCV("SSI,VCB")            // B channel
client.SubscribeForeignRoom("SSI,VCB")      // R channel
client.SubscribeIndex("VN30,VN100")         // MI channel
client.SubscribeOddLot("SSI,VCB")           // OL channel
```

---

## Sequence Diagrams

### REST Data Retrieval

```
App                    FCDataClient           FCData Server
 |                         |                      |
 | Login(id, secret)       |                      |
 |------------------------>| POST /AccessToken    |
 |                         |--------------------->|
 |                         |                      |
 |                         |    {accessToken}     |
 |                         |<---------------------|
 | token                   |                      |
 |<------------------------|                      |
 |                         |                      |
 | GetSecuritiesList(...)  |                      |
 |------------------------>| GET /Securities      |
 |                         | Auth: Bearer token   |
 |                         |--------------------->|
 |                         |                      |
 |                         |   SecuritiesList     |
 |                         |<---------------------|
 | *SecuritiesList         |                      |
 |<------------------------|                      |
 |                         |                      |
```

### Real-Time Streaming

```
App                  SignalR Client          FCData Hub
 |                         |                      |
 | NewClient(url, token)   |                      |
 |------------------------>|                      |
 |                         |                      |
 | StartWithLoop()         |                      |
 |------------------------>|                      |
 |                         | negotiate()          |
 |                         |--------------------->|
 |                         |   {ConnectionToken}  |
 |                         |<---------------------|
 |                         |                      |
 |                         | connect() (WebSocket)|
 |                         |--------------------->|
 |                         |     WS Upgrade       |
 |                         |<--------------------->|
 |                         |                      |
 |                         | start()              |
 |                         |--------------------->|
 |                         |    {OK}              |
 |                         |<---------------------|
 |                         |                      |
 | OnConnected callback    |                      |
 |<------------------------|                      |
 |                         |                      |
 | SwitchChannel("X-QUOTE")|                      |
 |------------------------>| Invoke(SwitchChannels)|
 |                         |--------------------->|
 |                         |                      |
 |                         |       ...            |
 |                         |  Message loop       |
 |                         |<---------------------|
 |                         |                      |
 | OnData callback         |                      |
 |<------------------------| Broadcast(...)       |
 |                         |                      |
```

---

## Technology Stack

| Layer | Technology |
|-------|------------|
| Language | Go 1.24+ |
| HTTP Client | `net/http` |
| JSON | `encoding/json` |
| WebSocket | `gorilla/websocket` |
| Data Format | JSON |

---

## Configuration

### Default Configuration

```go
const DefaultBaseURL = "https://fc-data.ssi.com.vn/"

FCDataClient{
    BaseURL: DefaultBaseURL,
    HTTPClient: &http.Client{
        Timeout: 10 * time.Second,
    },
    Token: "",
}
```

### SignalR Configuration

```go
const (
    ClientProtocolVersion = "1.3"
    HubName              = "FcMarketDataV2Hub"
)
```

---

## Error Handling

### HTTP Status Codes

| Status | Handling |
|--------|----------|
| 200-299 | Success, parse response |
| 400-499 | Client error, return error with status |
| 500-599 | Server error, return error with status |

### SignalR Errors

| Error | Handling |
|-------|----------|
| Negotiation failed | Return error with status |
| WebSocket dial failed | Return error |
| Start failed | Return error |
| Read error | Invoke OnError callback |