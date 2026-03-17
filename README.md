# go-fcdata

Go client library for the SSI FCData API - Vietnam Stock Exchange data.

## Overview

go-fcdata provides a Go client for accessing Vietnamese stock market data through the SSI FCData API. It supports both REST API for historical/batch data and SignalR for real-time streaming.

## Features

- **REST API Client**: Access securities, indices, OHLC prices, and more
- **SignalR Client**: Real-time streaming of market data via WebSocket
- **Authentication**: Bearer token-based authentication
- **Pagination Support**: Built-in pagination for list endpoints

## Installation

```bash
go get github.com/kyimmQ/go_fcdata
```

## Quick Start

```go
package main

import (
    "fmt"
    "os"

    "github.com/kyimmQ/go_fcdata/client"
)

func main() {
    consumerID := os.Getenv("ConsumerID")
    consumerSecret := os.Getenv("ConsumerSecret")

    // Create client and authenticate
    fcClient := client.NewFCDataClient("")
    token, err := fcClient.Login(consumerID, consumerSecret)
    if err != nil {
        fmt.Printf("Login failed: %v\n", err)
        return
    }
    fmt.Println("Login successful!")

    // Fetch securities list
    resp, err := fcClient.GetSecuritiesList("HO", 1, 10)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Total securities: %d\n", resp.TotalRecord)
}
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `ConsumerID` | SSI FCData API Consumer ID |
| `ConsumerSecret` | SSI FCData API Consumer Secret |

## REST API Reference

### Authentication

```go
token, err := fcClient.Login(consumerID, consumerSecret)
```

### Market Data Endpoints

| Method | Description |
|--------|-------------|
| `GetSecuritiesList(market, pageIndex, pageSize)` | Get list of securities |
| `GetSecuritiesDetails(market, symbol, pageIndex, pageSize)` | Get security details |
| `GetIndexList(exchange, pageIndex, pageSize)` | Get list of indices |
| `GetIndexComponents(indexCode, pageIndex, pageSize)` | Get index components |
| `GetDailyOhlc(symbol, fromDate, toDate, pageIndex, pageSize, ascending)` | Get daily OHLC data |
| `GetIntradayOhlc(symbol, fromDate, toDate, pageIndex, pageSize, ascending)` | Get intraday OHLC data |
| `GetDailyIndex(indexId, fromDate, toDate, pageIndex, pageSize, ascending)` | Get daily index data |
| `GetDailyStockPrice(symbol, market, fromDate, toDate, pageIndex, pageSize)` | Get daily stock prices |

## Real-Time Streaming (SignalR)

```go
import "github.com/kyimmQ/go_fcdata/signalr"
import "github.com/kyimmQ/go_fcdata/models"

// Create SignalR client
streamClient := signalr.NewClient("https://fc-datahub.ssi.com.vn/v2.0/signalr", token)

// Set up callbacks
streamClient.OnConnected = func() {
    fmt.Println("Connected!")
    streamClient.SwitchChannel("X-QUOTE:SSI")
    streamClient.SwitchChannel("X-TRADE:41I1G3000")
}

streamClient.OnData = func(msg models.BroadcastMessage) {
    // Handle typed messages automatically unmarshaled
    switch data := msg.Data.(type) {
    case models.XQuoteData:
        fmt.Printf("X-QUOTE -> Symbol: %s, Bid: %.2f, Ask: %.2f\n",
            data.Symbol, data.BidPrice1, data.AskPrice1)
    case models.XTradeData:
        fmt.Printf("X-TRADE -> Symbol: %s, Price: %.2f, Vol: %.2f\n",
            data.Symbol, data.LastPrice, data.LastVol)
    default:
        fmt.Printf("Received: %s\n", msg.Content)
    }
}

streamClient.OnError = func(err error) {
    fmt.Printf("Error: %v\n", err)
}

// Start streaming
streamClient.StartWithLoop()
```

## Project Structure

```
go-fcdata/
├── client/           # REST API client
│   ├── auth.go       # Authentication
│   ├── endpoints.go  # Market data endpoints
│   ├── fcdata.go    # Core client implementation
│   └── debug.go     # Raw request helpers
├── signalr/          # SignalR client
│   └── client.go    # WebSocket connection
├── models/          # Data models
│   └── models.go    # Request/response structures
├── example/         # Example usage
│   └── main.go      # Demo application
└── docs/            # Documentation
```

## Documentation

- [Project Overview & PDR](./docs/project-overview-pdr.md)
- [Codebase Summary](./docs/codebase-summary.md)
- [Code Standards](./docs/code-standards.md)
- [System Architecture](./docs/system-architecture.md)
- [Project Roadmap](./docs/project-roadmap.md)
- [Deployment Guide](./docs/deployment-guide.md)

## License

MIT License

## References

- [SSI FCData API Documentation](https://fc-data.ssi.com.vn/)
- [Vietnam Stock Exchange](https://www.ssi.com.vn/)
