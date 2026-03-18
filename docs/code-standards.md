# Code Standards

## Overview

This document defines the coding standards, conventions, and best practices for the go-fcdata project.

## Language Version

- **Minimum Go Version**: 1.24.3
- **Standard Library Only**: Uses Go standard library for HTTP and JSON
- **External Dependency**: `gorilla/websocket` v1.5.3

---

## Coding Conventions

### Naming Conventions

| Element         | Convention              | Example                       |
| --------------- | ----------------------- | ----------------------------- |
| Package         | Lowercase, short        | `client`, `signalr`, `models` |
| Types (structs) | PascalCase              | `FCDataClient`, `Client`      |
| Functions       | PascalCase              | `GetSecuritiesList`, `Login`  |
| Methods         | PascalCase              | `doRequest`, `SetToken`       |
| Variables       | CamelCase               | `httpClient`, `baseURL`       |
| Constants       | PascalCase or CamelCase | `DefaultBaseURL`, `hubName`   |
| Private fields  | CamelCase               | `connToken`, `messageID`      |

### File Organization

- One package per directory
- Related types grouped in same file
- Test files in same package with `_test.go` suffix

### Import Organization

```go
import (
    "fmt"
    "io"
    "net/http"
    "net/url"

    "github.com/gorilla/websocket"
    "github.com/kyimmQ/go-fcdata/models"
)
```

Order:

1. Standard library (alphabetical)
2. External packages
3. Project packages

---

## API Design Standards

### Client Pattern

```go
// NewFCDataClient creates a new client with default settings
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
```

### Method Signatures

- Public methods: Document with godoc comments
- Return error as last return value
- Pointer receivers for methods modifying state

```go
// GetSecuritiesList retrieves a paginated list of securities.
// Parameters:
//   - market: Market filter (HO, HNX, UPCOM)
//   - pageIndex: Page number (1-based)
//   - pageSize: Number of items per page
//
// Returns:
//   - *SecuritiesListResponse: The response data
//   - error: Any error that occurred
func (c *FCDataClient) GetSecuritiesList(market string, pageIndex, pageSize int) (*models.SecuritiesListResponse, error)
```

---

## Error Handling

### HTTP Errors

```go
if resp.StatusCode < 200 || resp.StatusCode >= 300 {
    respBody, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
}
```

### Validation Errors

```go
if respBody.Status != 200 {
    return "", fmt.Errorf("login failed: %s", respBody.Message)
}
```

### Resource Cleanup

```go
resp, err := c.HTTPClient.Do(req)
if err != nil {
    return nil, err
}
defer resp.Body.Close()  // Always defer close
```

---

## Data Models

### JSON Tags

Use explicit JSON tags for all exported fields:

```go
type AuthRequest struct {
    ConsumerID     string `json:"consumerID"`
    ConsumerSecret string `json:"consumerSecret"`
}
```

### Response Wrappers

REST responses follow common pattern:

```go
type SecuritiesListResponse struct {
    Status      string `json:"status"`
    Message     string `json:"message"`
    TotalRecord int    `json:"totalRecord"`
    Data        []struct {
        Symbol    string `json:"Symbol"`
        Market    string `json:"Market"`
        StockName string `json:"StockName"`
    } `json:"data"`
}
```

---

## SignalR Client Standards

### Connection State

Use mutex for thread-safe state management:

```go
type Client struct {
    mu          sync.Mutex
    connected   bool
    Conn        *websocket.Conn
    // ...
}
```

### Callback Pattern

Provide optional callbacks for async events:

```go
type Client struct {
    OnData      func(models.BroadcastMessage)
    OnConnected func()
    OnError     func(error)
}
```

---

## Testing Standards

### Test File Naming

```
client_test.go            -> tests for client package
endpoints_debug_test.go   -> endpoint debug tests
signalr/client_test.go   -> signalr client tests
models/models_test.go     -> models tests
```

### Test Functions

```go
func TestGetSecuritiesList(t *testing.T) {
    // Test implementation
}
```

---

## Configuration

### Constants

Group related constants:

```go
const (
    DefaultBaseURL = "https://fc-data.ssi.com.vn/"
)

const (
    SecuritiesListEndpoint    = "api/v2/Market/Securities"
    SecuritiesDetailsEndpoint = "api/v2/Market/SecuritiesDetails"
    // ...
)
```

### Environment Variables

Document required environment variables:

| Variable         | Description             |
| ---------------- | ----------------------- |
| `ConsumerID`     | SSI API Consumer ID     |
| `ConsumerSecret` | SSI API Consumer Secret |

---

## Documentation Standards

### Godoc Comments

Every public type and function must have documentation:

```go
// FCDataClient is the main client for FCData REST API.
type FCDataClient struct {
    // ...
}

// Login authenticates with the FCData API using consumer credentials.
// Returns the access token on success.
func (c *FCDataClient) Login(consumerID, consumerSecret string) (string, error)
```

---

## Versioning

### Semantic Importing

- Module path: `github.com/kyimmQ/go-fcdata`
- Go version declared in `go.mod`

---

## Security Considerations

### Token Handling

- Store token in client struct (not global)
- Include in Authorization header with "Bearer " prefix
- Do not log sensitive credentials

### HTTPS Only

- Default to HTTPS endpoints
- Production should never use HTTP
