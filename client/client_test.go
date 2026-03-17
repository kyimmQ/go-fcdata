package client

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getTestCredentials(t *testing.T) (string, string) {
	_ = godotenv.Load() // Load .env file if it exists
	consumerID := os.Getenv("CONSUMER_ID")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	if consumerID == "" || consumerSecret == "" {
		t.Skip("Skipping test because ConsumerID and/or ConsumerSecret environment variables are not set")
	}
	return consumerID, consumerSecret
}

func getAuthenticatedClient(t *testing.T) *FCDataClient {
	consumerID, consumerSecret := getTestCredentials(t)
	fcClient := NewFCDataClient("")
	_, err := fcClient.Login(consumerID, consumerSecret)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}
	return fcClient
}

func TestLogin(t *testing.T) {
	consumerID, consumerSecret := getTestCredentials(t)
	fcClient := NewFCDataClient("")

	token, err := fcClient.Login(consumerID, consumerSecret)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if token == "" {
		t.Errorf("Expected token, got empty string")
	}
}

func TestGetSecuritiesList(t *testing.T) {
	client := getAuthenticatedClient(t)

	resp, err := client.GetSecuritiesList("HOSE", 1, 10)
	if err != nil {
		t.Fatalf("GetSecuritiesList failed: %v", err)
	}

	if resp.Status != "Success" {
		t.Errorf("Expected status 200, got %s", resp.Status)
	}

	if len(resp.Data) == 0 {
		t.Errorf("Expected data, got 0 items")
	}
}

func TestGetSecuritiesDetails(t *testing.T) {
	client := getAuthenticatedClient(t)

	resp, err := client.GetSecuritiesDetails("HOSE", "SSI", 1, 10)
	if err != nil {
		t.Fatalf("GetSecuritiesDetails failed: %v", err)
	}

	if resp.Status != "Success" {
		t.Errorf("Expected status 200, got %s", resp.Status)
	}
}

func TestGetIndexList(t *testing.T) {
	client := getAuthenticatedClient(t)

	resp, err := client.GetIndexList("HOSE", 1, 10)
	if err != nil {
		t.Fatalf("GetIndexList failed: %v", err)
	}

	if resp.Status != "Success" {
		t.Errorf("Expected status 200, got %s", resp.Status)
	}
}

func TestGetIndexComponents(t *testing.T) {
	client := getAuthenticatedClient(t)

	resp, err := client.GetIndexComponents("VN30", 1, 10)
	if err != nil {
		t.Fatalf("GetIndexComponents failed: %v", err)
	}

	if resp.Status != "Success" {
		t.Errorf("Expected status 200, got %s", resp.Status)
	}
}

func TestGetDailyOhlc(t *testing.T) {
	client := getAuthenticatedClient(t)

	resp, err := client.GetDailyOhlc("SSI", "01/01/2023", "10/01/2023", 1, 10, true)
	if err != nil {
		t.Fatalf("GetDailyOhlc failed: %v", err)
	}

	if resp.Status != "Success" {
		t.Errorf("Expected status 200, got %s", resp.Status)
	}
}
