# Deployment Guide

## Overview

This guide covers deployment considerations for applications using go-fcdata.

---

## Prerequisites

### Environment Requirements
- Go 1.24.3 or later
- Internet access to FCData API endpoints

### API Credentials
You must obtain credentials from SSI:
- **Consumer ID**: Your API consumer identifier
- **Consumer Secret**: Your API secret key

---

## Installation

### Using Go Modules

Add to your project:

```bash
go get github.com/kyimmQ/go_fcdata
```

Or in your Go code:

```go
import (
    "github.com/kyimmQ/go_fcdata/client"
    "github.com/kyimmQ/go_fcdata/signalr"
    "github.com/kyimmQ/go_fcdata/models"
)
```

---

## Configuration

### Environment Variables

Set credentials before running your application:

```bash
# Bash/Zsh
export ConsumerID="your-consumer-id"
export ConsumerSecret="your-consumer-secret"

# Windows PowerShell
$env:ConsumerID="your-consumer-id"
$env:ConsumerSecret="your-consumer-secret"
```

### Using .env File

The project includes a `.env` example file:

```
CONSUMER_ID=your-consumer-id
CONSUMER_SECRET=your-consumer-secret
```

For Go applications, use a library like `godotenv`:

```go
import (
    "github.com/joho/godotenv"
    "os"
)

func init() {
    godotenv.Load()
}
```

---

## Basic Usage

### REST API

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

    // Create client
    fcClient := client.NewFCDataClient("")

    // Authenticate
    token, err := fcClient.Login(consumerID, consumerSecret)
    if err != nil {
        fmt.Printf("Login failed: %v\n", err)
        return
    }
    fmt.Printf("Logged in with token: %s\n", token[:20]+"...")

    // Fetch data
    securities, err := fcClient.GetSecuritiesList("HO", 1, 10)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Found %d securities\n", securities.TotalRecord)
}
```

### Real-Time Streaming

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/kyimmQ/go_fcdata/client"
    "github.com/kyimmQ/go_fcdata/signalr"
    "github.com/kyimmQ/go_fcdata/models"
)

func main() {
    consumerID := os.Getenv("ConsumerID")
    consumerSecret := os.Getenv("ConsumerSecret")

    // REST authentication
    fcClient := client.NewFCDataClient("")
    token, err := fcClient.Login(consumerID, consumerSecret)
    if err != nil {
        fmt.Printf("Login failed: %v\n", err)
        return
    }

    // SignalR streaming
    streamClient := signalr.NewClient(
        "https://fc-datahub.ssi.com.vn/v2.0/signalr",
        token,
    )

    streamClient.OnConnected = func() {
        fmt.Println("Connected to SignalR")
        streamClient.SwitchChannel("X-QUOTE:SSI")
    }

    streamClient.OnData = func(msg models.BroadcastMessage) {
        fmt.Printf("Data: %s\n", msg.Content)
    }

    streamClient.OnError = func(err error) {
        fmt.Printf("Error: %v\n", err)
    }

    if err := streamClient.StartWithLoop(); err != nil {
        fmt.Printf("Failed to start: %v\n", err)
        return
    }

    // Wait for shutdown signal
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    <-sigs

    streamClient.Close()
}
```

---

## Production Deployment

### Timeout Configuration

The default HTTP timeout is 10 seconds. Adjust as needed:

```go
fcClient := client.NewFCDataClient("")
fcClient.HTTPClient.Timeout = 30 * time.Second
```

### Connection Pooling

For high-volume applications, configure the HTTP transport:

```go
fcClient := client.NewFCDataClient("")
fcClient.HTTPClient.Transport = &http.Transport{
    MaxIdleConns:        10,
    MaxIdleConnsPerHost:  10,
    IdleConnTimeout:      30 * time.Second,
}
```

### Error Handling Best Practices

```go
func fetchWithRetry(client *client.FCDataClient, maxRetries int) error {
    var lastErr error
    for i := 0; i < maxRetries; i++ {
        _, err := client.GetSecuritiesList("HO", 1, 10)
        if err == nil {
            return nil
        }
        lastErr = err
        time.Sleep(time.Duration(i+1) * time.Second)
    }
    return lastErr
}
```

---

## Security Considerations

### Credential Storage

- Never commit credentials to version control
- Use environment variables or secure secret management
- Consider using a vault for production

### HTTPS Only

The library defaults to HTTPS. Ensure no HTTP endpoints are used in production:

```go
// Default is already HTTPS
// https://fc-data.ssi.com.vn/
// https://fc-datahub.ssi.com.vn/
```

### Token Management

- Tokens have expiration - implement refresh logic
- Don't log tokens in production

---

## Monitoring

### Logging

Add request logging for debugging:

```go
// Example: Use a custom HTTP client with logging
fcClient := client.NewFCDataClient("")
// Integrate with your logging framework
```

### Error Metrics

Track error rates:

- Authentication failures
- API request failures
- SignalR disconnections

---

## Troubleshooting

### Common Issues

| Issue | Solution |
|-------|----------|
| Login fails | Check ConsumerID and ConsumerSecret |
| 401 Unauthorized | Token expired, re-authenticate |
| 429 Too Many Requests | Implement rate limiting |
| Connection timeout | Check network, increase timeout |
| SignalR disconnect | Implement reconnection logic |

### Debug Mode

Use raw response for debugging:

```go
raw, err := fcClient.GetSecuritiesListRaw("HO", 1, 10)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(raw))
```

---

## Performance Tips

1. **Reuse Client**: Create client once, reuse for all requests
2. **Close Response Bodies**: Always defer `resp.Body.Close()`
3. **Limit Concurrent Requests**: Use goroutine pools if needed
4. **Cache Responses**: Consider caching for frequently accessed data

---

## Example Deployment

### Docker

```dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main ./example

FROM alpine
COPY --from=builder /app/main .
ENV ConsumerID=xxx
ENV Consumer_SECRET=xxx
CMD ["./main"]
```

### Kubernetes

```yaml
apiDeployment:
spec:
  containers:
  - name: app
    image: your-app:latest
    env:
    - name: ConsumerID
      valueFrom:
        secretKeyRef:
          name: fcdata-credentials
          key: consumer-id
```

---

## Testing

### Unit Tests

Run tests:

```bash
go test ./...
```

### Integration Tests

Set credentials for integration tests:

```bash
ConsumerID=xxx ConsumerSecret=xxx go test -tags=integration ./...
```

---

## Support

For issues and questions:
- GitHub Issues: https://github.com/kyimmQ/go_fcdata/issues
- SSI FCData API Support: https://fc-data.ssi.com.vn/