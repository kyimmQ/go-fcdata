package models

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalFData(t *testing.T) {
	jsonStr := `{
		"DataType": "F",
		"Content": "{\"RType\":\"F\",\"MarketId\":\"HOSE\",\"TradingDate\":\"14/08/2023\",\"Time\":\"13:00:00\",\"Symbol\":\"SSI\",\"TradingSession\":\"LO\",\"TradingStatus\":\"N\",\"Exchange\":\"HOSE\"}"
	}`

	var msg BroadcastMessage
	if err := json.Unmarshal([]byte(jsonStr), &msg); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if msg.DataType != "F" {
		t.Errorf("Expected DataType 'F', got '%s'", msg.DataType)
	}

	fData, ok := msg.Data.(FData)
	if !ok {
		t.Fatalf("Expected FData, got %T", msg.Data)
	}

	if fData.Symbol != "SSI" {
		t.Errorf("Expected Symbol 'SSI', got '%s'", fData.Symbol)
	}
	if fData.MarketId != "HOSE" {
		t.Errorf("Expected MarketId 'HOSE', got '%s'", fData.MarketId)
	}
}

func TestUnmarshalMIData(t *testing.T) {
	// Sample from docs (using single quotes in JSON for keys which is invalid, but Content is a string inside JSON)
	// Content string itself uses double quotes.
	// Sample: {"DataType":"MI","Content":'{...}'}
	// Wait, the sample JSON output in docs uses single quotes for keys which is invalid JSON.
	// I must use valid JSON for testing. The sample in docs:
	/*
	   Input:
	   MI:VN30
	   Output:
	   {'DataType': 'MI',
	   'Content': '{
	   "IndexId":"VN30",
	   "IndexValEst":1200.03,
	   ...
	   }'}
	}
	*/
	// I will construct valid JSON.

	jsonStr := `{
		"DataType": "MI",
		"Content": "{\"IndexId\":\"VN30\",\"IndexValEst\":1200.03,\"IndexValue\":1238.76,\"PriorIndexValue\":1226.16,\"TradingDate\":\"02/04/2021\",\"Time\":\"11:28:13\",\"TotalTrade\":0.0,\"TotalQtty\":191838100.0,\"TotalValue\":7289093000000.0,\"IndexName\":\"VN30\",\"Advances\":25,\"NoChanges\":2,\"Declines\":3,\"Ceilings\":0,\"Floors\":0,\"Change\":12.6,\"RatioChange\":1.03,\"TotalQttyPt\":2064000.0,\"TotalValuePt\":244251000000.0,\"Exchange\":\"HOSE\",\"AllQty\":193902100.0,\"AllValue\":7533344000000.0,\"IndexType\":\"Main\",\"TradingSession\":null,\"MarketId\":null,\"RType\":\"MI\",\"TotalQttyOd\":0.0,\"TotalValueOd\":0.0}"
	}`

	var msg BroadcastMessage
	if err := json.Unmarshal([]byte(jsonStr), &msg); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if msg.DataType != "MI" {
		t.Errorf("Expected DataType 'MI', got '%s'", msg.DataType)
	}

	miData, ok := msg.Data.(MIData)
	if !ok {
		t.Fatalf("Expected MIData, got %T", msg.Data)
	}

	if miData.IndexId != "VN30" {
		t.Errorf("Expected IndexId 'VN30', got '%s'", miData.IndexId)
	}
	if miData.Advances != 25 {
		t.Errorf("Expected Advances 25, got %d", miData.Advances)
	}
}

func TestUnmarshalOLData(t *testing.T) {
	jsonStr := `{
		"DataType": "OL",
		"Content": "{\"RType\":\"OL\",\"TradingDate\":\"18/02/2025\",\"Time\":\"13:55:03\",\"StockNo\":2027,\"Symbol\":\"MBB\",\"Ceiling\":24200,\"Floor\":21100,\"RefPrice\":22650,\"Open\":22650,\"High\":22950,\"Low\":22600,\"LastPrice\":22750,\"LastVol\":9193,\"TotalVal\":185028289999.99728,\"TotalVol\":8135000,\"BidPrice1\":22700,\"BidPrice2\":22650,\"BidPrice3\":22600,\"BidVol1\":1108,\"BidVol2\":1630,\"BidVol3\":2016,\"AskPrice1\":22750,\"AskPrice2\":22800,\"AskPrice3\":22850,\"AskVol1\":132,\"AskVol2\":548,\"AskVol3\":297,\"Exchange\":\"HOSE\",\"TradingSession\":\"LO\",\"TradingStatus\":\"H\",\"Change\":100,\"RatioChange\":0.44,\"StockType\":\"Stock\"}"
	}`

	var msg BroadcastMessage
	if err := json.Unmarshal([]byte(jsonStr), &msg); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if msg.DataType != "OL" {
		t.Errorf("Expected DataType 'OL', got '%s'", msg.DataType)
	}

	olData, ok := msg.Data.(OLData)
	if !ok {
		t.Fatalf("Expected OLData, got %T", msg.Data)
	}

	if olData.Symbol != "MBB" {
		t.Errorf("Expected Symbol 'MBB', got '%s'", olData.Symbol)
	}
	if olData.LastPrice != 22750 {
		t.Errorf("Expected LastPrice 22750, got %f", olData.LastPrice)
	}
}

func TestUnmarshalXSnapshotDataWithNaN(t *testing.T) {
	jsonStr := `{
		"DataType": "X",
		"Content": "{\"RType\":\"X\",\"TradingDate\":\"19/03/2026\",\"Time\":\"08:00:07\",\"Isin\":\"41I1G4000\",\"Symbol\":\"41I1G4000\",\"Ceiling\":2011.6,\"Floor\":1748.4,\"RefPrice\":1880.0,\"Open\":0.0,\"High\":0.0,\"Low\":0.0,\"Close\":0.0,\"AvgPrice\":\"NaN\",\"PriorVal\":1880.0,\"LastPrice\":0.0,\"LastVol\":0.0,\"TotalVal\":0.0,\"TotalVol\":0.0,\"BidPrice1\":0.0,\"BidPrice2\":0.0,\"BidPrice3\":0.0,\"BidPrice4\":0.0,\"BidPrice5\":0.0,\"BidPrice6\":0.0,\"BidPrice7\":0.0,\"BidPrice8\":0.0,\"BidPrice9\":0.0,\"BidPrice10\":0.0,\"BidVol1\":0.0,\"BidVol2\":0.0,\"BidVol3\":0.0,\"BidVol4\":0.0,\"BidVol5\":0.0,\"BidVol6\":0.0,\"BidVol7\":0.0,\"BidVol8\":0.0,\"BidVol9\":0.0,\"BidVol10\":0.0,\"AskPrice1\":0.0,\"AskPrice2\":0.0,\"AskPrice3\":0.0,\"AskPrice4\":0.0,\"AskPrice5\":0.0,\"AskPrice6\":0.0,\"AskPrice7\":0.0,\"AskPrice8\":0.0,\"AskPrice9\":0.0,\"AskPrice10\":0.0,\"AskVol1\":0.0,\"AskVol2\":0.0,\"AskVol3\":0.0,\"AskVol4\":0.0,\"AskVol5\":0.0,\"AskVol6\":0.0,\"AskVol7\":0.0,\"AskVol8\":0.0,\"AskVol9\":0.0,\"AskVol10\":0.0,\"MarketId\":\"DERIVATIVES\",\"Exchange\":\"DERIVATIVES\",\"TradingSession\":\"C\",\"TradingStatus\":\"Normal\",\"Change\":-1880.0,\"RatioChange\":-100.0,\"EstMatchedPrice\":0.0,\"Side\":null,\"CloseQtty\":0.0}"
	}`

	var msg BroadcastMessage
	if err := json.Unmarshal([]byte(jsonStr), &msg); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if msg.DataType != "X" {
		t.Errorf("Expected DataType 'X', got '%s'", msg.DataType)
	}

	xData, ok := msg.Data.(XSnapshotData)
	if !ok {
		t.Fatalf("Expected XSnapshotData, got %T", msg.Data)
	}

	if xData.Symbol != "41I1G4000" {
		t.Errorf("Expected Symbol '41I1G4000', got '%s'", xData.Symbol)
	}

	// Check that AvgPrice defaults to 0.0 when "NaN" is received
	if xData.AvgPrice != 0.0 {
		t.Errorf("Expected AvgPrice 0.0, got %f", xData.AvgPrice)
	}
}
