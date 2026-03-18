package signalr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/kyimmQ/go-fcdata/models"
)

const (
	ClientProtocolVersion = "1.3"
	HubName               = "FcMarketDataV2Hub"
)

// Client handles the SignalR connection to FCData Hub.
type Client struct {
	BaseURL     string
	Token       string
	Conn        *websocket.Conn
	mu          sync.Mutex
	messageID   int
	connected   bool
	hubData     string // JSON encoded array of hub names, e.g., [{"name": "fcmarketdatav2hub"}]
	connToken   string

	OnData      func(models.BroadcastMessage)
	OnConnected func()
	OnError     func(error)
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: baseURL,
		Token:   token,
		hubData: `[{"name":"fcmarketdatav2hub"}]`,
	}
}

func (c *Client) getAuthHeader() http.Header {
	header := http.Header{}
	if c.Token != "" {
		if strings.HasPrefix(c.Token, "Bearer ") {
			header.Set("Authorization", c.Token)
		} else {
			header.Set("Authorization", "Bearer "+c.Token)
		}
	}
	return header
}

func (c *Client) negotiate() (*models.NegotiationResponse, error) {
	negotiateURL := fmt.Sprintf("%s/negotiate?connectionData=%s&clientProtocol=%s",
		c.BaseURL, url.QueryEscape(c.hubData), ClientProtocolVersion)

	req, err := http.NewRequest("GET", negotiateURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header = c.getAuthHeader()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("negotiate failed with status %d: %s", resp.StatusCode, string(b))
	}

	var negResp models.NegotiationResponse
	if err := json.NewDecoder(resp.Body).Decode(&negResp); err != nil {
		return nil, err
	}

	c.connToken = negResp.ConnectionToken
	return &negResp, nil
}

func (c *Client) connect() error {
	wsBaseURL := strings.Replace(c.BaseURL, "https://", "wss://", 1)
	wsBaseURL = strings.Replace(wsBaseURL, "http://", "ws://", 1)

	connectURL := fmt.Sprintf("%s/connect?clientProtocol=%s&transport=webSockets&connectionToken=%s&connectionData=%s&tid=10",
		wsBaseURL, ClientProtocolVersion, url.QueryEscape(c.connToken), url.QueryEscape(c.hubData))

	conn, resp, err := websocket.DefaultDialer.Dial(connectURL, c.getAuthHeader())
	if err != nil {
		if resp != nil {
			b, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("websocket dial failed: %v, body: %s", err, string(b))
		}
		return fmt.Errorf("websocket dial failed: %v", err)
	}

	c.Conn = conn
	return nil
}

func (c *Client) startRequest() error {
	startURL := fmt.Sprintf("%s/start?clientProtocol=%s&transport=webSockets&connectionData=%s&connectionToken=%s",
		c.BaseURL, ClientProtocolVersion, url.QueryEscape(c.hubData), url.QueryEscape(c.connToken))

	req, err := http.NewRequest("GET", startURL, nil)
	if err != nil {
		return err
	}
	req.Header = c.getAuthHeader()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("start failed with status %d: %s", resp.StatusCode, string(b))
	}
	return nil
}

func (c *Client) Start() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected {
		return nil
	}

	negResp, err := c.negotiate()
	if err != nil {
		return fmt.Errorf("negotiation error: %w", err)
	}

	if !negResp.TryWebSockets {
		return fmt.Errorf("server does not support WebSockets")
	}

	if err := c.connect(); err != nil {
		return fmt.Errorf("connect error: %w", err)
	}

	if err := c.startRequest(); err != nil {
		c.Conn.Close()
		return fmt.Errorf("start error: %w", err)
	}

	c.connected = true
	if c.OnConnected != nil {
		go c.OnConnected()
	}

	return nil
}

func (c *Client) readLoop() {
	defer func() {
		c.mu.Lock()
		c.connected = false
		if c.Conn != nil {
			c.Conn.Close()
		}
		c.mu.Unlock()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if c.OnError != nil {
				c.OnError(fmt.Errorf("read error: %w", err))
			}
			return
		}

		if len(message) == 0 || string(message) == "{}" {
			continue // Keep-alive or empty message
		}

		var sigMsg models.SignalRMessage
		if err := json.Unmarshal(message, &sigMsg); err != nil {
			if c.OnError != nil {
				c.OnError(fmt.Errorf("json unmarshal error: %w, payload: %s", err, string(message)))
			}
			continue
		}

		if len(sigMsg.M) > 0 {
			for _, m := range sigMsg.M {
				if strings.EqualFold(m.H, HubName) && strings.EqualFold(m.M, "Broadcast") {
					for _, arg := range m.A {
						var argBytes []byte

						// Handle if arg is already a string (JSON string)
						if str, ok := arg.(string); ok {
							argBytes = []byte(str)
						} else {
							argBytes, _ = json.Marshal(arg) //nolint:errchkjson
						}

						var bMsg models.BroadcastMessage
						if err := json.Unmarshal(argBytes, &bMsg); err == nil {
							if c.OnData != nil {
								c.OnData(bMsg)
							}
						} else {
							fmt.Printf("Failed to unmarshal Broadcast argument: %s, error: %v\n", string(argBytes), err)
						}
					}
				} else if strings.EqualFold(m.H, HubName) && strings.EqualFold(m.M, "Error") {
					if c.OnError != nil {
						c.OnError(fmt.Errorf("hub error: %v", m.A))
					}
				}
			}
		}
	}
}

// Modify Start to run readLoop
func (c *Client) StartWithLoop() error {
	if err := c.Start(); err != nil {
		return err
	}
	go c.readLoop()
	return nil
}

func (c *Client) Invoke(hub, method string, args ...interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return fmt.Errorf("client is not connected")
	}
	c.messageID++
	msgID := c.messageID

	payload := map[string]interface{}{
		"H": hub,
		"M": method,
		"A": args,
		"I": fmt.Sprintf("%d", msgID),
	}

	return c.Conn.WriteJSON(payload)
}

func (c *Client) SwitchChannel(channel string) error {
	return c.Invoke(HubName, "SwitchChannels", channel)
}

func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.connected {
		return nil
	}
	c.connected = false
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}
