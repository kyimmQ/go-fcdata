# Streaming data

FastConnect Data provides you realtime data through streaming. As soon as the client is connected, the system will send the update data if any changes. Update data includes:

<table><thead><tr><th width="126.33333333333331">Data type</th><th width="257">Description</th><th>Note</th></tr></thead><tbody><tr><td>F</td><td>Securities status</td><td>Refer details <a href="#securities-status">here</a>. </td></tr><tr><td>X</td><td>Snapshot of latest best bid/ask and trade</td><td>Refer details <a href="#securities-snapshot">here</a>.</td></tr><tr><td>X-Quote</td><td>Best Bid/ask</td><td>HOSE supports 3 best prices. HNX, UPCOM and Derivatives supports 10 best prices. Refer details <a href="#quote">here</a>. </td></tr><tr><td>X-Trade</td><td>Matched volume and price</td><td>Refer details <a href="#trade">here.</a> </td></tr><tr><td>B</td><td>Realtime open, high, low, close, volume</td><td>Refer <a href="#ohlcv">OHLCV</a></td></tr><tr><td>R</td><td>Foreign Room</td><td>Refer details <a href="#foreign-room">here.</a> </td></tr><tr><td>MI</td><td>Index data</td><td>Refer details <a href="#index">here.</a> </td></tr></tbody></table>

### Securities status

Rtype F returns trading session and trading status of securities.&#x20;

{% tabs %}
{% tab title="Input" %}
F: \<Securities list>

**Note:** Securities are separated by "-". Input ALL to get status of all securities. &#x20;

**Example:** F: SSI or F:SSI-PAN or F:ALL
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="175">Field</th><th width="101">Type</th><th>Description</th></tr></thead><tbody><tr><td>Rtype</td><td>string</td><td>Value = F</td></tr><tr><td>MarketID</td><td>string</td><td>Exchange ID, includes: HOSE|HNX|HNXBOND|UPCOM|DER</td></tr><tr><td>TradingDate</td><td>Date</td><td>Trading date in ddmmyyy format</td></tr><tr><td>Time</td><td>Time</td><td>Time in format HHMMSS</td></tr><tr><td>Symbol</td><td>String</td><td>Securities symbol</td></tr><tr><td>TradingSession</td><td>String</td><td><a href="data-mapping/trading-session">Phiên giao dịch</a></td></tr><tr><td>TradingStatus</td><td>String</td><td><a href="data-mapping/securities-trading-status">Trạng thái giao dịch</a></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
F: SSI
Output:
{"DataType":"F",
"Content":"{
\"RType\":\"F\",
\"MarketId\":\"HOSE\",
\"TradingDate\":\"14/08/2023\",
\"Time\":\"13:00:00\",
\"Symbol\":\"SSI\",
\"TradingSession\":\"LO\",
\"TradingStatus\":\"N\",
\"Exchange\":\"HOSE\"}"
}
```

### Quote

RType X-Quote provides best bid/ask.&#x20;

- For HOSE, 3 best bid/ask prices are provided.&#x20;
- For HNX, UPCOM, DER, 10 best bid/ask prices are provided.&#x20;

{% tabs %}
{% tab title="Input" %}
X-Quote: \<Securities list> or ALL&#x20;

**Note**: Securities list including symbols separated by "-"

ALL: All securities
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="192.66666666666666">Field</th><th width="143">Type</th><th>Description</th></tr></thead><tbody><tr><td>Rtype</td><td>string</td><td>Value = X-QUOTE</td></tr><tr><td>TradingDate</td><td>Date</td><td>Format DD/MM/YYYY</td></tr><tr><td>Time</td><td>Time</td><td>Format HH:MM:SS</td></tr><tr><td>Symbol</td><td>string</td><td></td></tr><tr><td>BidPrice1</td><td>number</td><td></td></tr><tr><td>BidVol1</td><td>number</td><td></td></tr><tr><td>BidPrice2</td><td>number</td><td></td></tr><tr><td>BidVol2</td><td>number</td><td></td></tr><tr><td>BidPrice3</td><td>number</td><td></td></tr><tr><td>BidVol3</td><td>number</td><td></td></tr><tr><td>BidPrice4</td><td>number</td><td></td></tr><tr><td>BidVol4</td><td>number</td><td></td></tr><tr><td>BidPrice5</td><td>number</td><td></td></tr><tr><td>BidVol5</td><td>number</td><td></td></tr><tr><td>BidPrice6</td><td>number</td><td></td></tr><tr><td>BidVol6</td><td>number</td><td></td></tr><tr><td>BidPrice7</td><td>number</td><td></td></tr><tr><td>BidVol7</td><td>number</td><td></td></tr><tr><td>BidPrice8</td><td>number</td><td></td></tr><tr><td>BidVol8</td><td>number</td><td></td></tr><tr><td>BidPrice9</td><td>number</td><td></td></tr><tr><td>BidVol9</td><td>number</td><td></td></tr><tr><td>BidPrice10</td><td>number</td><td></td></tr><tr><td>BidVol10</td><td>number</td><td></td></tr><tr><td>AskPrice1</td><td>number</td><td></td></tr><tr><td>AskVol1</td><td>number</td><td></td></tr><tr><td>AskPrice2</td><td>number</td><td></td></tr><tr><td>AskVol2</td><td>number</td><td></td></tr><tr><td>AskPrice3</td><td>number</td><td></td></tr><tr><td>AskVol3</td><td>number</td><td></td></tr><tr><td>AskPrice4</td><td>number</td><td></td></tr><tr><td>AskVol4</td><td>number</td><td></td></tr><tr><td>AskPrice5</td><td>number</td><td></td></tr><tr><td>AskVol5</td><td>number</td><td></td></tr><tr><td>AskPrice6</td><td>number</td><td></td></tr><tr><td>AskVol6</td><td>number</td><td></td></tr><tr><td>AskPrice7</td><td>number</td><td></td></tr><tr><td>AskVol7</td><td>number</td><td></td></tr><tr><td>AskPrice8</td><td>number</td><td></td></tr><tr><td>AskVol8</td><td>number</td><td></td></tr><tr><td>AskPrice9</td><td>number</td><td></td></tr><tr><td>AskVol9</td><td>number</td><td></td></tr><tr><td>AskPrice10</td><td>number</td><td>10th best ask price</td></tr><tr><td>AskVol10</td><td>number</td><td>10th best ask vol </td></tr><tr><td>MarketID</td><td> string</td><td>Market</td></tr><tr><td>Exchange</td><td> string</td><td> Exchange</td></tr><tr><td>TradingSession</td><td>string</td><td><a href="data-mapping/trading-session">Phiên giao dịch</a></td></tr><tr><td>TradingStatus</td><td>string</td><td><a href="data-mapping/securities-trading-status">Trạng thái giao dịch</a></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
X-QUOTE: ACB
Output:
{"DataType":"X-QUOTE",
"Content":"{
\"TradingDate\":\"14/08/2023\",
\"Time\":\"14:00:28\",
\"Exchange\":\"HOSE\",
\"Symbol\":\"ACB\",
\"RType\":\"X-QUOTE\",
\"AskPrice1\":22950.0,
\"AskPrice2\":23000.0,
\"AskPrice3\":23050.0,
\"AskPrice4\":0.0,
\"AskPrice5\":0.0,
\"AskPrice6\":0.0,
\"AskPrice7\":0.0,
\"AskPrice8\":0.0,
\"AskPrice9\":0.0,
\"AskPrice10\":0.0,
\"AskVol1\":109100.0,
\"AskVol2\":208600.0,
\"AskVol3\":207200.0,
\"AskVol4\":0.0,
\"AskVol5\":0.0,
\"AskVol6\":0.0,
\"AskVol7\":0.0,
\"AskVol8\":0.0,
\"AskVol9\":0.0,
\"AskVol10\":0.0,
\"BidPrice1\":22900.0,
\"BidPrice2\":22850.0,
\"BidPrice3\":22800.0,
\"BidPrice4\":0.0,
\"BidPrice5\":0.0,
\"BidPrice6\":0.0,
\"BidPrice7\":0.0,
\"BidPrice8\":0.0,
\"BidPrice9\":0.0,
\"BidPrice10\":0.0,
\"BidVol1\":290900.0,
\"BidVol2\":454200.0,
\"BidVol3\":678800.0,
\"BidVol4\":0.0,
\"BidVol5\":0.0,
\"BidVol6\":0.0,
\"BidVol7\":0.0,
\"BidVol8\":0.0,
\"BidVol9\":0.0,
\"BidVol10\":0.0,
\"TradingSession\":\"LO\"}"}
```

### Trade

{% tabs %}
{% tab title="Input" %}
X-TRADE: \<Securities list> or ALL&#x20;

**Note:** Securities list including symbols separated by "-"&#x20;

ALL: All securities
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="192.66666666666666">Field</th><th width="143">Data</th><th>Description</th></tr></thead><tbody><tr><td>Rtype</td><td>string</td><td>Value = X-TRADE</td></tr><tr><td>TradingDate</td><td>Date</td><td>In DD/MM/YYYY format</td></tr><tr><td>Time</td><td>Time</td><td>In HH:MM:SS format</td></tr><tr><td>ISIN</td><td>string</td><td>ISIN</td></tr><tr><td>Symbol</td><td>string</td><td>Symbol</td></tr><tr><td>Ceiling</td><td>number</td><td></td></tr><tr><td>Floor</td><td>number</td><td></td></tr><tr><td>RefPrice</td><td>number</td><td></td></tr><tr><td>Highest</td><td>number</td><td></td></tr><tr><td>Lowest</td><td>number</td><td></td></tr><tr><td>AvgPrice</td><td>number</td><td></td></tr><tr><td><p>PriorVal</p><p> </p></td><td>number</td><td></td></tr><tr><td>LastPrice</td><td>number</td><td></td></tr><tr><td>Change</td><td>number</td><td></td></tr><tr><td>RatioChange</td><td>number</td><td></td></tr><tr><td>EstMatchedPrice</td><td>number</td><td>Estimated matched price in ATO/ATC</td></tr><tr><td>LastVol</td><td>number</td><td>Lats matched vol</td></tr><tr><td>TotalVal</td><td>number</td><td>Total matched value </td></tr><tr><td>TotalVol</td><td><p>number</p><p> </p></td><td>Total matched vol</td></tr><tr><td>MarketID</td><td> string</td><td> </td></tr><tr><td>Exchange</td><td> string</td><td> </td></tr><tr><td>TradingSession</td><td>string</td><td><a href="data-mapping/trading-session">Phiên giao dịch</a></td></tr><tr><td>TradingStatus</td><td>string</td><td><a href="data-mapping/securities-trading-status">Trạng thái giao dịch </a></td></tr><tr><td>Side</td><td>string</td><td>Buy up/ Sell down <br>BU: Buy up <br>SD: Sell down <br>Unknow</td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample:**&#x20;

```json
Input:
X-TRADE: ACB
Output:
{"DataType":"X-TRADE",
"Content":"{
\"RType\":\"X-TRADE\",
\"TradingDate\":\"14/08/2023\",
\"Time\":\"14:07:53\",
\"Isin\":\"ACB\",
\"Symbol\":\"ACB\",
\"Ceiling\":24500.0,
\"Floor\":21300.0,
\"RefPrice\":22900.0,
\"AvgPrice\":22927.93,
\"PriorVal\":22900.0,
\"LastPrice\":22950.0,
\"LastVol\":100.0,
\"TotalVal\":201713015000.0,
\"TotalVol\":8797700.0,
\"MarketId\":\"HOSE\",
\"Exchange\":\"HOSE\",
\"TradingSession\":\"LO\",
\"TradingStatus\":\"N\",
\"Change\":50.0,
\"RatioChange\":0.22,
\"EstMatchedPrice\":22900.0,
\"Highest\":23050,
\"Lowest\":22800,
\"Side\":\"SD\",
}"}
```

### Securities snapshot

{% tabs %}
{% tab title="Input" %}
X: \<securities list> or ALL
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="192.66666666666666">Field</th><th width="143">Type</th><th>Description</th></tr></thead><tbody><tr><td>Rtype</td><td>string</td><td>Value = X</td></tr><tr><td>TradingDate</td><td>Date</td><td>Trading date in DD/MM/YYYY format</td></tr><tr><td>Time</td><td>Time</td><td>Time in HH:MM:SS</td></tr><tr><td>ISIN</td><td>string</td><td>ISIN</td></tr><tr><td>Symbol</td><td>string</td><td></td></tr><tr><td>Ceiling</td><td>number</td><td></td></tr><tr><td>Floor</td><td>number</td><td></td></tr><tr><td>RefPrice</td><td>number</td><td></td></tr><tr><td>Open</td><td>number</td><td></td></tr><tr><td>Close</td><td>number</td><td></td></tr><tr><td>High</td><td>number</td><td></td></tr><tr><td>Low</td><td>number</td><td></td></tr><tr><td>Avg</td><td>number</td><td></td></tr><tr><td><p>PriorVal</p><p> </p></td><td>number</td><td></td></tr><tr><td>LastPrice</td><td>number</td><td>Last matched price</td></tr><tr><td>Change</td><td>number</td><td></td></tr><tr><td>RatioChange</td><td>number</td><td></td></tr><tr><td>EstMatchedPrice</td><td>number</td><td>Estimated matched price during ATO/ATC</td></tr><tr><td>LastVol</td><td>number</td><td>Last matched vol</td></tr><tr><td>TotalVal</td><td>number</td><td>Total matched value</td></tr><tr><td>TotalVol</td><td><p>number</p><p> </p></td><td>Total matched vol</td></tr><tr><td>BidPrice1</td><td>number</td><td></td></tr><tr><td>BidVol1</td><td>number</td><td></td></tr><tr><td>BidPrice2</td><td>number</td><td></td></tr><tr><td>BidVol2</td><td>number</td><td></td></tr><tr><td>BidPrice3</td><td>number</td><td></td></tr><tr><td>BidVol3</td><td>number</td><td></td></tr><tr><td>BidPrice4</td><td>number</td><td></td></tr><tr><td>BidVol4</td><td>number</td><td></td></tr><tr><td>BidPrice5</td><td>number</td><td></td></tr><tr><td>BidVol5</td><td>number</td><td></td></tr><tr><td>BidPrice6</td><td>number</td><td></td></tr><tr><td>BidVol6</td><td>number</td><td></td></tr><tr><td>BidPrice7</td><td>number</td><td></td></tr><tr><td>BidVol7</td><td>number</td><td></td></tr><tr><td>BidPrice8</td><td>number</td><td></td></tr><tr><td>BidVol8</td><td>number</td><td></td></tr><tr><td>BidPrice9</td><td>number</td><td></td></tr><tr><td>BidVol9</td><td>number</td><td></td></tr><tr><td>BidPrice10</td><td>number</td><td></td></tr><tr><td>BidVol10</td><td>number</td><td></td></tr><tr><td>AskPrice1</td><td>number</td><td></td></tr><tr><td>AskVol1</td><td>number</td><td></td></tr><tr><td>AskPrice2</td><td>number</td><td></td></tr><tr><td>AskVol2</td><td>number</td><td></td></tr><tr><td>AskPrice3</td><td>number</td><td></td></tr><tr><td>AskVol3</td><td>number</td><td></td></tr><tr><td>AskPrice4</td><td>number</td><td></td></tr><tr><td>AskVol4</td><td>number</td><td></td></tr><tr><td>AskPrice5</td><td>number</td><td></td></tr><tr><td>AskVol5</td><td>number</td><td></td></tr><tr><td>AskPrice6</td><td>number</td><td></td></tr><tr><td>AskVol6</td><td>number</td><td></td></tr><tr><td>AskPrice7</td><td>number</td><td></td></tr><tr><td>AskVol7</td><td>number</td><td></td></tr><tr><td>AskPrice8</td><td>number</td><td></td></tr><tr><td>AskVol8</td><td>number</td><td></td></tr><tr><td>AskPrice9</td><td>number</td><td></td></tr><tr><td>AskVol9</td><td>number</td><td></td></tr><tr><td>AskPrice10</td><td>number</td><td></td></tr><tr><td>AskVol10</td><td>number</td><td></td></tr><tr><td>MarketID</td><td> string</td><td> Sàn giao dịch</td></tr><tr><td>Exchange</td><td> string</td><td> Sàn giao dịch</td></tr><tr><td>TradingSession</td><td>string</td><td><a href="data-mapping/trading-session">Phiên giao dịch</a></td></tr><tr><td>TradingStatus</td><td>string</td><td><a href="data-mapping/securities-trading-status">Trạng thái giao dịch </a></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
X: ALL
Output:
{"DataType":"X",
"Content":"{
\"RType\":\"X\",
\"TradingDate\":\"15/08/2023\",
\"Time\":\"10:28:41\",
\"Isin\":\"TAR\",
\"Symbol\":\"TAR\",
\"Ceiling\":23400.0,
\"Floor\":19200.0,
\"RefPrice\":21300.0,
\"Open\":21300.0,
\"High\":22300.0,
\"Low\":21300.0,
\"Close\":21800.0,
\"AvgPrice\":21784.0,
\"PriorVal\":21300.0,
\"LastPrice\":21800.0,
\"LastVol\":100.0,
\"TotalVal\":27946100000.0,
\"TotalVol\":1282900.0,
\"BidPrice1\":21800.0,
\"BidPrice2\":21700.0,
\"BidPrice3\":21600.0,
\"BidPrice4\":21500.0,
\"BidPrice5\":21400.0,
\"BidPrice6\":21300.0,
\"BidPrice7\":21200.0,
\"BidPrice8\":21100.0,
\"BidPrice9\":21000.0,
\"BidPrice10\":20900.0,
\"BidVol1\":18800.0,
\"BidVol2\":35300.0,
\"BidVol3\":93600.0,
\"BidVol4\":60000.0,
\"BidVol5\":45800.0,
\"BidVol6\":94500.0,
\"BidVol7\":42900.0,
\"BidVol8\":48400.0,
\"BidVol9\":165900.0,
\"BidVol10\":26100.0,
\"AskPrice1\":21900.0,
\"AskPrice2\":22000.0,
\"AskPrice3\":22100.0,
\"AskPrice4\":22200.0,
\"AskPrice5\":22300.0,
\"AskPrice6\":22400.0,
\"AskPrice7\":22500.0,
\"AskPrice8\":22600.0,
\"AskPrice9\":22700.0,
\"AskPrice10\":22800.0,
\"AskVol1\":50300.0,
\"AskVol2\":109800.0,
\"AskVol3\":58500.0,
\"AskVol4\":152600.0,
\"AskVol5\":92200.0,
\"AskVol6\":129800.0,
\"AskVol7\":178700.0,
\"AskVol8\":60100.0,
\"AskVol9\":29400.0,
\"AskVol10\":17000.0,
\"MarketId\":\"HNX\",
\"Exchange\":\"HNX\",
\"TradingSession\":\"LO\",
\"TradingStatus\":\"N\",
\"Change\":500.0,
\"RatioChange\":2.35,
\"EstMatchedPrice\":0.0}"}
```

### Foreign room

{% tabs %}
{% tab title="Input" %}
R:\<Securities list>

**Note:** Securities in the list are separated by "-". Input R to get room of all securities.&#x20;

**Example:** R:SSI or R:SSI-PAN or R:ALL
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="159">Field</th><th width="120">Type</th><th>Description</th></tr></thead><tbody><tr><td>Rtype</td><td>string</td><td>Value = R</td></tr><tr><td>TradingDate</td><td>Date</td><td>Trading date in dd/mm/yyyy format</td></tr><tr><td>Time</td><td>Time</td><td>Time in HH:MM:SS format</td></tr><tr><td>ISIN</td><td>string</td><td>ISIN</td></tr><tr><td>Symbol</td><td>string</td><td></td></tr><tr><td>TotalRoom</td><td>number</td><td></td></tr><tr><td><p>CurrentRoom</p><p> </p></td><td>number</td><td>Current room</td></tr><tr><td>FBuyVol</td><td>number</td><td>Total matched buy vol of foreign customers</td></tr><tr><td>FSellVol</td><td> number</td><td>Total matched sell vol of foreign customers</td></tr><tr><td>FBuyVal</td><td>number</td><td><p>Total matched buy value of foreign customers. For HOSE, the value is approximately calculated by</p><p>BuyVal = BuyVol * LastPrice</p></td></tr><tr><td>FSellVal</td><td>number</td><td><p>Total matched sell value of foreign customers. For HOSE, the value is approximately calculated by</p><p>SellVal = SellVol * LastPrice</p></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
R:SSI
Output:
{"DataType":"R",
"Content":"{
\"RType\":\"R\",
\"TradingDate\":\"14/08/2023\",
\"Time\":\"14:24:02\",
\"Isin\":\"SSI\",
\"Symbol\":\"SSI\",
\"TotalRoom\":1501130137.0,
\"CurrentRoom\":806887173.0,
\"BuyVol\":863352.0,
\"SellVol\":825308.0,
\"BuyVal\":25123543200.0,
\"SellVal\":24016462800.0,
\"MarketId\":\"HOSE\",
\"Exchange\":\"HOSE\"}"
}
```

### Index

Rtype MI provides realtime update index values of HOSE, HNX, UPCOM.&#x20;

{% tabs %}
{% tab title="Input" %}
MI:\<Index list>

**Note:** Indexes in the list are separated by "-". Input ALL to get update of all indexes. &#x20;

**Example:** MI:VN30 or MI:VN30-HNXindex or MI:ALL
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="146">Field</th><th width="126">Type</th><th width="320">Description</th></tr></thead><tbody><tr><td>Rtype</td><td>String</td><td>Value = MI</td></tr><tr><td>IndexID</td><td>String</td><td>Index ID</td></tr><tr><td>IndexValEst</td><td>Number</td><td>Estimated Index value during ATO/ATC</td></tr><tr><td>IndexValue</td><td>Number</td><td>Index value</td></tr><tr><td>Trading Date</td><td>Date</td><td></td></tr><tr><td>Time</td><td>Timestamp</td><td></td></tr><tr><td>Change</td><td>Number</td><td></td></tr><tr><td>RatioChange</td><td>Number</td><td></td></tr><tr><td>TotalTrade</td><td>Number</td><td>Total trade including normal and putthrough transactions</td></tr><tr><td>TotalQtty</td><td>Number</td><td>Total matched quantity</td></tr><tr><td>TotalValue</td><td>Number</td><td>Total matched value</td></tr><tr><td>TypeIndex</td><td>String</td><td><p>Index: <br>“Main” – main index like VN30 </p><p>“Industry” – Industrial indexes </p><p>“Khác” – Others</p></td></tr><tr><td>IndexName</td><td>String</td><td>Index name</td></tr><tr><td>Advances</td><td>Number</td><td>Total number of securities with price increase</td></tr><tr><td>Nochanges</td><td>Number</td><td>Total number of securities with price unchanged. </td></tr><tr><td>Declines</td><td>Number</td><td>Total number of securities with price decrease</td></tr><tr><td>Ceiling</td><td>Number</td><td>Total number of securities with last price = ceiling</td></tr><tr><td>Floor</td><td>Number</td><td>Total number of securities with last price = floor</td></tr><tr><td>TotalQttyPT</td><td>Number</td><td>Total matched quantity by putthrough</td></tr><tr><td>TotalValuePT</td><td>Number</td><td>Total matched value by putthrough</td></tr><tr><td>TotalQttyOd</td><td>Number</td><td>Total matched quantity by odd lot orders</td></tr><tr><td>TotalValueOd</td><td>Number</td><td> Total matched value by odd lot orders</td></tr><tr><td>AllQty</td><td>Number</td><td>Total matched quantity including normal and putthrough orders</td></tr><tr><td>AllValue</td><td>Number</td><td>Total matched value including normal and putthrough orders</td></tr><tr><td>TradingSession</td><td>String</td><td><a href="data-mapping/trading-session">Phiên giao dịch</a></td></tr><tr><td>Exchange</td><td>String</td><td><a href="data-mapping/exchanges">Sàn giao dịch</a></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample:**&#x20;

```json
Input:
MI:VN30
Output:
{'DataType': 'MI',
'Content': {
"IndexId":"VN30",
“IndexValEst”:1200.03,
"IndexValue":1238.76,
"PriorIndexValue":1226.16,
"TradingDate":"02/04/2021",
"Time":"11:28:13",
"TotalTrade":0.0,
"TotalQtty":191838100.0,
"TotalValue":7289093000000.0,
"IndexName":"VN30",
"Advances":25,
"NoChanges":2,
"Declines":3,"Ceilings":0,
"Floors":0,"Change":12.6,
"RatioChange":1.03,
"TotalQttyPt":2064000.0,
"TotalValuePt":244251000000.0,
"Exchange":"HOSE",
"AllQty":193902100.0,
"AllValue":7533344000000.0,
"IndexType":"Main",
"TradingSession":null,
"MarketId":null,
"RType":"MI",
"TotalQttyOd":0.0,
"TotalValueOd":0.0}'
}
```

### OHLCV

Rtype B returns open, high, low, close, volume of securities/indexes by tick.&#x20;

{% tabs %}
{% tab title="Input" %}
B:\<Securities or Indexes>

**Note:** Securities/Indexes are separated by "-". Input ALL to get realtime OHLCV realtime of all symbols/indexes.&#x20;

**Example:** B:SSI; B:SSI-VN30 or B:ALL
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="153">Field</th><th width="143">Type</th><th>Description</th></tr></thead><tbody><tr><td>Rtype</td><td>String</td><td>Value = B</td></tr><tr><td>Time</td><td>Timestamp</td><td>Including date and time</td></tr><tr><td>Symbol</td><td>String</td><td>Securities symbol/index</td></tr><tr><td>Open</td><td>Number</td><td></td></tr><tr><td>High</td><td>Number</td><td></td></tr><tr><td>Low</td><td>Number</td><td></td></tr><tr><td>Close</td><td>Number</td><td></td></tr><tr><td>Volume</td><td>Number</td><td>Last matched volume</td></tr><tr><td>Value</td><td>Number</td><td>Last matched value. Not used yet, temporarily set to 0. </td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample:**&#x20;

```json
Input:
B:X26
Output:
{‘Datatype’: 'B',
 ‘Content’: '{
 "RType":"B",
 "Symbol":"X26",
 "TradingTime":"14:28:33",
 "Open":16000,
 "High":16000,
 "Low":16000,
 "Close":16000,
 "Volume":5000,
 "Value":0}'
 }
```

### Odlot message

Rtype OL: return odlot message including open, high, low, close, volume .. of stocks&#x20;

{% tabs %}
{% tab title="Input" %}
**OL:All** : return odlot data of all stocks in 3 exchange HOSE, HNX, UPCOM

**OL:\<Stock>**: return registered stock odlot data&#x20;

**Note:** Securitiesare separated by "-". Input ALL to get realtime Odlot data realtime of all symbols

**Example:** B:SSI; B:SSI-VND hoặc B:ALL
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="157">Tên trường</th><th width="117">Kiểu dữ liệu</th><th>Mô tả</th></tr></thead><tbody><tr><td>Rtype</td><td>String</td><td>Value OL</td></tr><tr><td>TradingDate</td><td>String</td><td>Trading date in format dd/MM/yyyy</td></tr><tr><td>Time</td><td>Timestamp</td><td>Time in format HH24:MI:SS</td></tr><tr><td>StockNo</td><td>Number</td><td>Stock Number</td></tr><tr><td>Symbol</td><td>String</td><td>Stock symbol</td></tr><tr><td>Ceiling</td><td>Number</td><td>Ceilling price</td></tr><tr><td>Floor</td><td>Number</td><td>Floor price</td></tr><tr><td>RefPrice</td><td>Number</td><td>Reference price</td></tr><tr><td>Open</td><td>Number</td><td>Open</td></tr><tr><td>High</td><td>Number</td><td>High</td></tr><tr><td>Low</td><td>Number</td><td>Low</td></tr><tr><td>LastPrice</td><td>Number</td><td>Last matched price</td></tr><tr><td>LastVol</td><td>Number</td><td>Last matched volume</td></tr><tr><td>TotalVal</td><td>Number</td><td>Total value including normal and odlot value</td></tr><tr><td>TotalVol</td><td>Number</td><td>Total volume including normal and odlot volume </td></tr><tr><td>BidPrice1 </td><td>Number</td><td>Bid price 1</td></tr><tr><td>BidPrice2 </td><td>Number</td><td>Bid price 2</td></tr><tr><td>BidPrice3 </td><td>Number</td><td>Bid price 3</td></tr><tr><td>BidVol1 </td><td>Number</td><td>Bid volume 1</td></tr><tr><td>BidVol2 </td><td>Number</td><td>Bid volume 2</td></tr><tr><td>BidVol3 </td><td>Number</td><td>Bid volume 3</td></tr><tr><td>AskPrice1 </td><td>Number</td><td>Ask price 1</td></tr><tr><td>AskPrice2 </td><td>Number</td><td>Ask price 2</td></tr><tr><td>AskPrice3 </td><td>Number</td><td>Ask price 3</td></tr><tr><td>AskVol1 </td><td>Number</td><td>Ask volume 1</td></tr><tr><td>AskVol2 </td><td>Number</td><td>Ask volume 2</td></tr><tr><td>AskVol3</td><td>Number</td><td>Ask volume 3</td></tr><tr><td>Exchange</td><td>String</td><td>Exchange (HOSE/HNX/UPCOM)</td></tr><tr><td>TradingSession</td><td>String</td><td>Trading session (LO/ATO/ATC/CLOSE/Break)</td></tr><tr><td>TradingStatus</td><td>String</td><td>Stock status N/H (Normal/Halt)</td></tr><tr><td>Change</td><td>Number</td><td>Change in price compared to the reference price</td></tr><tr><td>RatioChange</td><td>Number</td><td>Percentage change in price compared to the reference price</td></tr><tr><td>StockType</td><td>String</td><td>Stock type (Stock / Bond/ ETF...)</td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample:**&#x20;

```json

Input:
OL:MBB
Output:
{
  "DataType": "OL",
  "Content": {
    "RType": "OL",
    "TradingDate": "18/02/2025",
    "Time": "13:55:03",
    "StockNo": 2027,
    "Symbol": "MBB",
    "Ceiling": 24200,
    "Floor": 21100,
    "RefPrice": 22650,
    "Open": 22650,
    "High": 22950,
    "Low": 22600,
    "LastPrice": 22750,
    "LastVol": 9193,
    "TotalVal": 185028289999.99728,
    "TotalVol": 8135000,
    "BidPrice1": 22700,
    "BidPrice2": 22650,
    "BidPrice3": 22600,
    "BidVol1": 1108,
    "BidVol2": 1630,
    "BidVol3": 2016,
    "AskPrice1": 22750,
    "AskPrice2": 22800,
    "AskPrice3": 22850,
    "AskVol1": 132,
    "AskVol2": 548,
    "AskVol3": 297,
    "Exchange": "HOSE",
    "TradingSession": "LO",
    "TradingStatus": "H",
    "Change": 100,
    "RatioChange": 0.44,
    "StockType": "Stock"
  }
}
```
