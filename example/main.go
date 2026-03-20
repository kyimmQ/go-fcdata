package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

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
		// Subscribe to X-TRADE and X-SNAPSHOT for symbols 41I1G3000 and 41I1G4000
		// streamClient.SubscribeTrade([]string{"41I1G3000", "41I1G4000"})
		streamClient.SubscribeSnapshot([]string{"41I1G3000", "41I1G4000"})
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
		case models.XSnapshotData:
			fmt.Printf("X-SNAPSHOT -> \n")
			//  print all fields of XSnapshotData
			fmt.Printf("Symbol: %s\n", data.Symbol)
			fmt.Printf("LastPrice: %.2f\n", data.LastPrice)
			fmt.Printf("BidPrice1: %.2f, BidVol1: %.2f\n", data.BidPrice1, data.BidVol1)
			fmt.Printf("AskPrice1: %.2f, AskVol1: %.2f\n", data.AskPrice1, data.AskVol1)
			fmt.Printf("BidPrice2: %.2f, BidVol2: %.2f\n", data.BidPrice2, data.BidVol2)
			fmt.Printf("AskPrice2: %.2f, AskVol2: %.2f\n", data.AskPrice2, data.AskVol2)
			fmt.Printf("BidPrice3: %.2f, BidVol3: %.2f\n", data.BidPrice3, data.BidVol3)
			fmt.Printf("AskPrice3: %.2f, AskVol3: %.2f\n", data.AskPrice3, data.AskVol3)
			fmt.Printf("BidPrice4: %.2f, BidVol4: %.2f\n", data.BidPrice4, data.BidVol4)
			fmt.Printf("AskPrice4: %.2f, AskVol4: %.2f\n", data.AskPrice4, data.AskVol4)
			fmt.Printf("BidPrice5: %.2f, BidVol5: %.2f\n", data.BidPrice5, data.BidVol5)
			fmt.Printf("AskPrice5: %.2f, AskVol5: %.2f\n", data.AskPrice5, data.AskVol5)
			fmt.Printf("TotalVal: %.2f, TotalVol: %.2f\n", data.TotalVal, data.TotalVol)
		default:
			fmt.Printf("STREAM DATA -> Type: %s, Content: %s\n", msg.DataType, msg.Content)
		}
	}

	streamClient.OnError = func(err error) {
		fmt.Printf("STREAM ERROR -> %v at %s\n", err, time.Now().Format("2006-01-02 15:04:05"))
	}

	fmt.Println("Starting stream client at", time.Now().Format("2006-01-02 15:04:05"))
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
