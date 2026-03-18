package signalr

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kyimmQ/go-fcdata/client"
	"github.com/kyimmQ/go-fcdata/models"
)

func getTestCredentials(t *testing.T) (string, string) {
	consumerID := os.Getenv("ConsumerID")
	consumerSecret := os.Getenv("ConsumerSecret")

	if consumerID == "" || consumerSecret == "" {
		t.Skip("Skipping test because ConsumerID and/or ConsumerSecret environment variables are not set")
	}
	return consumerID, consumerSecret
}

func getAuthenticatedToken(t *testing.T) string {
	consumerID, consumerSecret := getTestCredentials(t)
	fcClient := client.NewFCDataClient("")
	token, err := fcClient.Login(consumerID, consumerSecret)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}
	return token
}

func TestSignalRConnectAndStream(t *testing.T) {
	token := getAuthenticatedToken(t)
	c := NewClient("https://fc-datahub.ssi.com.vn/v2.0/signalr", token)

	connectedCh := make(chan bool)
	dataCh := make(chan models.BroadcastMessage, 1)

	c.OnConnected = func() {
		connectedCh <- true
		fmt.Println("Connected, switching channel to X-TRADE:SSI")
		if err := c.SwitchChannel("X-TRADE:SSI"); err != nil {
			t.Errorf("SwitchChannel failed: %v", err)
		}
	}

	c.OnData = func(msg models.BroadcastMessage) {
		fmt.Printf("Received data: %s\n", msg.DataType)
		select {
		case dataCh <- msg:
		default:
		}
	}

	if err := c.StartWithLoop(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer c.Close()

	select {
	case <-connectedCh:
	case <-time.After(10 * time.Second):
		t.Fatalf("Timeout waiting for connection")
	}

	select {
	case msg := <-dataCh:
		if msg.DataType == "" {
			t.Errorf("Expected populated message, got empty DataType")
		}
	case <-time.After(15 * time.Second):
		t.Logf("Timeout waiting for data (market might be closed or symbol inactive)")
	}
}
