package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kyimmQ/go-fcdata/client"
	"github.com/kyimmQ/go-fcdata/models"
	"github.com/kyimmQ/go-fcdata/signalr"
)

func main() {
	// get project root directory and load .env file
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}
	
	_ = godotenv.Load(filepath.Join(rootDir, ".env"))
	consumerID := os.Getenv("CONSUMER_ID")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	if consumerID == "" || consumerSecret == "" {
		fmt.Println("Please set ConsumerID and ConsumerSecret environment variables")
		return
	}

	// 1. Initialize REST Client and Login
	fcClient := client.NewFCDataClient("")
	token, err := fcClient.Login(consumerID, consumerSecret)
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		return
	}
	fmt.Println("Login successful, token obtained.")

	// 2. Fetch some REST Data
	fmt.Println("Fetching Index Components for VN30...")
	indexResp, err := fcClient.GetIndexComponents("VN30", 1, 10)
	if err != nil {
		fmt.Printf("Error fetching VN30 components: %v\n", err)
	} else {
		fmt.Printf("VN30 has %s components\n", indexResp.Data[0].TotalSymbolNo)
	}

	// 3. Initialize SignalR Streaming Client
	streamClient := signalr.NewClient("https://fc-datahub.ssi.com.vn/v2.0/signalr", token)

	streamClient.OnConnected = func() {
		fmt.Println("SignalR Connected!")
		// Subscribe to a few channels
		streamClient.SwitchChannel("X-TRADE:41I1G3000")
	}

	streamClient.OnData = func(msg models.BroadcastMessage) {

		// Use typed data
		switch data := msg.Data.(type) {
		case models.XQuoteData:
			fmt.Printf("X-QUOTE -> Symbol: %s, LastPrice: %.2f, Bid: %.2f/%.2f, Ask: %.2f/%.2f\n",
				data.Symbol, 0.0, data.BidPrice1, data.AskPrice1, data.BidVol1, data.AskVol1)
		case models.XTradeData:
			fmt.Printf("X-TRADE -> Symbol: %s, Price: %.2f, Vol: %.2f\n",
				data.Symbol, data.LastPrice, data.LastVol)
		default:
			fmt.Printf("STREAM DATA -> Type: %s, Content: %s\n", msg.DataType, msg.Content)
		}
	}

	streamClient.OnError = func(err error) {
		fmt.Printf("STREAM ERROR -> %v\n", err)
	}

	fmt.Println("Starting stream client...")
	if err := streamClient.StartWithLoop(); err != nil {
		fmt.Printf("Stream client failed to start: %v\n", err)
		return
	}

	// 4. Wait for termination
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("Shutting down...")
	streamClient.Close()
}
