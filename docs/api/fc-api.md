# API Specs

### API List

<table><thead><tr><th width="191.33333333333331">API</th><th width="128">Method</th><th>Path</th></tr></thead><tbody><tr><td>AccessToken</td><td>POST</td><td>Market/AccessToken</td></tr><tr><td>Securities</td><td>GET</td><td>Market/Securities</td></tr><tr><td>SecuritiesDetails</td><td>GET</td><td>Market/SecuritiesDetails</td></tr><tr><td>IndexComponents</td><td>GET</td><td>Market/IndexComponents</td></tr><tr><td>IndexList</td><td>GET</td><td>Market/IndexList</td></tr><tr><td>DailyOhlc</td><td>GET</td><td>Market/DailyOhlc</td></tr><tr><td>IntradayOhlc</td><td>GET</td><td>Market/IntradayOhlc</td></tr><tr><td>DailyIndex</td><td>GET</td><td>Market/DailyIndex</td></tr><tr><td>DailyStockPrice</td><td>GET</td><td>Market/DailyStockPrice</td></tr></tbody></table>

### **POST** AccessToken

```
https://fc-data.ssi.com.vn/api/v2/Market/AccessToken
```

To get access token used to run APIs or connect streaming of FastConnect Data.

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="181">Field</th><th width="103">Type</th><th width="113">Required?</th><th>Description</th></tr></thead><tbody><tr><td>consumerID</td><td>string</td><td>Yes</td><td>ConsumerID</td></tr><tr><td>consumerSecret</td><td>string</td><td>Yes</td><td>ConsumerSecret </td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="163.66666666666666">Field</th><th width="112">Type</th><th>Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td>Message</td></tr><tr><td>status</td><td>number</td><td>Status</td></tr><tr><td>accessToken</td><td>string</td><td>Token to access API and connect streaming</td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
{
  "consumerID": "c058f55761814787882b2c8df1336e25",
  "consumerSecret": "144cac45770949519d2dfd20edb5b6ab",
}
Output:
{
  "message": "Success",
  "status": 200,
  "data": {
    "accessToken": "eyJhbGciOiJSUzI1NiIsI"
	}
}
```

### **GET** Securities

```
https://fc-data.ssi.com.vn/api/v2/Market/Securities
```

To get securities list by exchange.

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="88">Type</th><th width="113">Required?</th><th>Description</th></tr></thead><tbody><tr><td>market</td><td>string</td><td>No</td><td><p>HOSE | HNX | UPCOM | DER</p><p>If not set, returns securities of all markets </p></td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td>From 1 to 10. Default value is 1. </td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td><p>10; 20; 50; 100; 1000</p><p>Default value is 10.</p></td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="167.66666666666666">Field</th><th width="114">Type</th><th>Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td>Message</td></tr><tr><td>status</td><td>number</td><td>Status</td></tr><tr><td>totalRecord</td><td>number</td><td>Total records</td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>market</td><td>string</td><td>HOSE | HNX | UPCOM | DER</td></tr><tr><td>symbol</td><td>string</td><td>Securities symbol</td></tr><tr><td>StockName</td><td>string</td><td>Stock Name in Vietnamese</td></tr><tr><td>StockEnName</td><td>string</td><td>Stock Name in English</td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
{
  "pageIndex" : "1",
  "pageSize": "10",
  "market": "hose"
}
Output:
{
  "data": [
    {
      "Market": "HOSE",
      "Symbol": "AAA",
      "StockName": "CTCP NHUA&MT XANH AN PHAT",
      "StockEnName": "An Phat Bioplastics Joint Stock Company"
    },
    {
      "Market": "HOSE",
            "Symbol": "AAM",
            "StockName": "CTCP THUY SAN MEKONG",
            "StockEnName": "Mekong Fisheries Joint Stock Company"
    }
    ],
    "message": "Success",
    "status": "Success",
    "totalRecord":2
}
```

### **GET** SecuritiesDetails

<pre><code><strong>https://fc-data.ssi.com.vn/api/v2/Market/SecuritiesDetails
</strong></code></pre>

To get securities details.&#x20;

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="88">Type</th><th width="113">Required?</th><th>Description</th></tr></thead><tbody><tr><td>market</td><td>string</td><td>No</td><td><p>HOSE | HNX | UPCOM | DER</p><p>If not set, returns securities of all markets </p></td></tr><tr><td>symbol</td><td>string</td><td>No</td><td>If not set, returns all securities of defined markets</td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td>From 1 to 10. Default value is 1</td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td><p>10; 20; 50; 100; 1000</p><p>Default value is 10</p></td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="203">Data</th><th width="135">Type</th><th width="287">Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td>Message</td></tr><tr><td>status</td><td>number</td><td>Status</td></tr><tr><td>totalRecord</td><td>number</td><td>Total records</td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>RType</td><td>string</td><td>y</td></tr><tr><td>ReportDate</td><td>number</td><td>Date of report<br>Format: dd/mm/yyyy</td></tr><tr><td>TotalNoSym</td><td>number</td><td>Total number of securities</td></tr><tr><td><strong>repeatedinfoList</strong></td><td>list</td><td>Data list</td></tr><tr><td>Isin</td><td>string</td><td>ISIN code </td></tr><tr><td>Symbol</td><td>string</td><td>Securities symbol. Used as key in trading activities. </td></tr><tr><td>SymbolName</td><td>string</td><td>Stock name in Vietnamese</td></tr><tr><td>SymbolEngName</td><td>string</td><td>Stock name in English</td></tr><tr><td>SecType</td><td>string</td><td><p>Securties Type:<br>ST: Stock</p><p>CW: Covered Warrant</p><p>FU: Futures</p><p>EF: ETF</p><p>BO: BOND</p><p>OF: OEF</p><p>MF: Mutual Fund</p></td></tr><tr><td>Exchange</td><td>string</td><td><p>Exchange: <br>HOSE</p><p>HNX</p><p>HNXBOND</p><p>UPCOM</p><p>DER</p></td></tr><tr><td>Issuer</td><td>string</td><td></td></tr><tr><td>LotSize</td><td>string</td><td></td></tr><tr><td>IssueDate</td><td>number</td><td> </td></tr><tr><td>MaturityDate</td><td>Date</td><td> </td></tr><tr><td>FirstTradingDate</td><td>Date</td><td> </td></tr><tr><td>LastTradingDate</td><td>Date</td><td> </td></tr><tr><td>ContractMultiplier</td><td>Date</td><td></td></tr><tr><td>SettlMethod</td><td>number</td><td></td></tr><tr><td>Underlying</td><td>string</td><td></td></tr><tr><td>PutOrCall</td><td>string</td><td></td></tr><tr><td>ExercisePrice</td><td>string</td><td></td></tr><tr><td>ExerciseStyle</td><td>number</td><td></td></tr><tr><td>ExcerciseRatio</td><td>string</td><td></td></tr><tr><td>ListedShare</td><td>string</td><td></td></tr><tr><td>TickPrice1</td><td>number</td><td></td></tr><tr><td>TickIncrement1</td><td>number</td><td></td></tr><tr><td>TickPrice2</td><td>number</td><td></td></tr><tr><td>TickIncrement2</td><td>number</td><td></td></tr><tr><td>TickPrice3</td><td>number</td><td></td></tr><tr><td>TickIncrement3</td><td>number</td><td></td></tr><tr><td>TickPrice4</td><td>number</td><td></td></tr><tr><td>TickIncrement4</td><td>number</td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
{
  "pageIndex" : "1",
  "pageSize": "10",
  "market": "hose",
  "symbol": "SSI"
}
Output:
{
  "data": [
    {
      "RType": "y",
      "ReportDate": "19/01/2023",
      "TotalNoSym": "1",
      "RepeatedInfo": [
        {
          "Isin": null,
          "Symbol": "SSI",
          "SymbolName": "CTCP CHUNG KHOAN SSI",
          "SymbolEngName": "SSI Securities Corporation",
          "SecType": "S",
          "MarketId": "HOSE",
          "Exchange": "HOSE",
          "Issuer": null,
          "LotSize": "100",
          "IssueDate": "",
          "MaturityDate": "",
          "FirstTradingDate": "",
          "LastTradingDate": "",
          "ContractMultiplier": "0",
          "SettlMethod": "",
          "Underlying": null,
          "PutOrCall": null,
          "ExercisePrice": "0",
          "ExerciseStyle": "",
          "ExcerciseRatio": "0",
          "ListedShare": "1501130137",
          "TickPrice1": null,
          "TickIncrement1": null,
          "TickPrice2": null,
          "TickIncrement2": null,
          "TickPrice3": null,
          "TickIncrement3": null,
          "TickPrice4": null,
          "TickIncrement4": null
        }
      ]
    }
  ],
  "message": "Success",
  "status": "Success",
  "totalRecord": 1
}
```

### **GET** IndexComponents

```
https://fc-data.ssi.com.vn/api/v2/Market/IndexComponents
```

To get securities list of an index.&#x20;

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="98">Type</th><th width="110">Required?</th><th>Description</th></tr></thead><tbody><tr><td>Indexcode</td><td>string</td><td>Yes</td><td>Input one index code to get securities list</td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td>From 1 to 10. Default value is 1</td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td><p>10; 20; 50; 100; 1000</p><p>Default value is 10</p></td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="187">Field</th><th width="109">Type</th><th width="399">Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td> </td></tr><tr><td>status</td><td>number</td><td> </td></tr><tr><td>totalRecord</td><td>number</td><td> </td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>IndexCode</td><td>string</td><td> </td></tr><tr><td>IndexName</td><td>string</td><td> </td></tr><tr><td>Exchange</td><td>string</td><td> HOSE|HNX</td></tr><tr><td>TotalSymbolNo</td><td>number</td><td>Total number of symbols in the index </td></tr><tr><td><strong>IndexComponent</strong></td><td>list</td><td> </td></tr><tr><td>Isin</td><td>string</td><td> </td></tr><tr><td>StockSymbol</td><td>string</td><td> </td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample:**

```json
Input:
{
"pageIndex": "1",
"pageSize": "10",
"indexCode": "VN30"
}
Output:
{
  "data": [
    {
      "IndexCode": "VN30",
      "IndexName": "VN30",
      "Exchange": "HOSE",
      "TotalSymbolNo": "30",
      "IndexComponent": [
        {
          "Isin": "ACB",
          "StockSymbol": "ACB"
        },
        {
          "Isin": "BCM",
          "StockSymbol": "BCM"
        }
      ]
    }
  ],
  "message": "Success",
  "status": "Success",
  "totalRecord": 1
}
```

### **GET** IndexList

<pre><code><strong>https://fc-data.ssi.com.vn/api/v2/Market/IndexList
</strong></code></pre>

To get index list. &#x20;

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="98">Type</th><th width="108">Required?</th><th>Description</th></tr></thead><tbody><tr><td>exchange</td><td>string</td><td>Yes</td><td>HOSE | HNX</td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td>From 1 to 10. Default 1</td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td>10; 20; 50; 100; 1000. Default 10</td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="159">Field</th><th width="109">Type</th><th width="399">Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td></td></tr><tr><td>status</td><td>number</td><td></td></tr><tr><td>totalRecord</td><td>number</td><td></td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>IndexCode</td><td>string</td><td>Index code</td></tr><tr><td>IndexName</td><td>string</td><td>Index name</td></tr><tr><td>Exchange</td><td>string</td><td>Exchange: HOSE|HNX</td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
{
  "pageIndex": "1",
  "pageSize": "10",
  "exchange": "HOSE"
}
Output:
{
  "data": [
    {
      "IndexCode": "VN100",
      "IndexName": "VN100",
      "Exchange": "HOSE"
    },
    {
      "IndexCode": "VN30",
      "IndexName": "VN30",
      "Exchange": "HOSE"
    }
  ],
  "message": "Success",
  "status": "Success",
  "totalRecord": 2
}
```

### **GET** DailyOhlc

<pre><code><strong>https://fc-data.ssi.com.vn/api/v2/Market/DailyOhlc
</strong></code></pre>

To get daily open, high, low, close, volume, value.

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="98">Type</th><th width="101">Required?</th><th>Description</th></tr></thead><tbody><tr><td>symbol</td><td>string</td><td>No</td><td>Securities/Index code</td></tr><tr><td>fromDate</td><td>string</td><td>Yes</td><td>Default is today if not set. In dd/mm/yyyy format</td></tr><tr><td>toDate</td><td>string</td><td>Yes</td><td>Default is today if not set. In dd/mm/yyyy format</td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td><p>From 1 to 10</p><p>Default 1</p></td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td><p>10; 20; 50; 100; 1000</p><p>Default 10</p></td></tr><tr><td>ascending</td><td>boolean</td><td>No</td><td>true/ false</td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="170">Field</th><th width="137">Type</th><th>Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td></td></tr><tr><td>status</td><td>number</td><td></td></tr><tr><td>totalRecord</td><td>number</td><td></td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>Symbol</td><td>String</td><td></td></tr><tr><td>Market</td><td>String</td><td>HOSE | HNX | UPCOM</td></tr><tr><td>TradingDate</td><td>Date</td><td>dd/mm/yyyy</td></tr><tr><td>Time</td><td>Timestamp</td><td></td></tr><tr><td>Open</td><td>Number</td><td> </td></tr><tr><td>High</td><td>Number</td><td> </td></tr><tr><td>Low</td><td>Number</td><td> </td></tr><tr><td>Close</td><td>Number</td><td></td></tr><tr><td>Volume</td><td>Number</td><td>Total normal matched volume</td></tr><tr><td>Value</td><td>Number</td><td>Total normal matched value</td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
{
  "pageIndex": "1",
  "pageSize": "10",
  "Symbol": "SSI",
  "Fromdate": "10/08/2023",
  "Todate": "13/08/2023"
}
Output:
{
  "data": [
    {
      "Symbol": "SSI",
      "Market": "HOSE",
      "TradingDate": "10/08/2023",
      "Time": null,
      "Open": "28600",
      "High": "28850",
      "Low": "28100",
      "Close": "28100",
      "Volume": "23382100",
      "Value": "663258204999.9850"
    },
    {
      "Symbol": "SSI",
      "Market": "HOSE",
      "TradingDate": "11/08/2023",
      "Time": null,
      "Open": "28250",
      "High": "28300",
      "Low": "27650",
      "Close": "28150",
      "Volume": "27536000",
      "Value": "769411290000.0090"
    }
  ],
  "message": "Success",
  "status": "Success",
  "totalRecord": 2
}
```

### **GET** IntradayOhlc

```
https://fc-data.ssi.com.vn/api/v2/Market/IntradayOhlc
```

To get realtime open, high, low, close, volume of securities.&#x20;

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="98">Type</th><th width="101">Required?</th><th>Description</th></tr></thead><tbody><tr><td>symbol</td><td>string</td><td>No</td><td>Securities symbols, indexes</td></tr><tr><td>fromDate</td><td>string</td><td>Yes</td><td>If not set, default to today<br>Format dd/mm/yyyy</td></tr><tr><td>toDate</td><td>string</td><td>Yes</td><td>If not set, default to today<br>Format dd/mm/yyyy</td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td><p>From 1 to 10</p><p>Default 1</p></td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td><p>10; 20; 50; 100; 1000</p><p>Default 10</p></td></tr><tr><td>ascending</td><td>boolean</td><td>No</td><td>true/ false</td></tr><tr><td>resollution</td><td>integer</td><td>No</td><td>Default 1 minute</td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="179">Field</th><th width="155">Type</th><th>Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td></td></tr><tr><td>status</td><td>number</td><td></td></tr><tr><td>totalRecord</td><td>number</td><td></td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>Symbol</td><td>String</td><td></td></tr><tr><td>Market</td><td>String</td><td></td></tr><tr><td>TradingDate</td><td>Date</td><td>Format dd/mm/yyyy</td></tr><tr><td>Time</td><td>Timestamp</td><td> </td></tr><tr><td>Open</td><td>Number</td><td>  </td></tr><tr><td>High</td><td>Number</td><td>  </td></tr><tr><td>Low</td><td>Number</td><td>  </td></tr><tr><td>Close</td><td>Number</td><td>  </td></tr><tr><td>Volume</td><td>Number</td><td> </td></tr><tr><td>Value</td><td>Number</td><td> </td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**&#x20;

```json
Input:
{
  "pageIndex": "1",
  "pageSize": "10",
  "Symbol": "SSI",
  "Fromdate": "14/08/2023",
  "Todate": "14/08/2023"
}
Output:
{
  "data": [
    {
      "Symbol": "SSI",
      "Value": "29150",
      "TradingDate": "14/08/2023",
      "Time": "14:45:04",
      "Open": "29150",
      "High": "29150",
      "Low": "29150",
      "Close": "29150",
      "Volume": "529200"
    },
    {
      "Symbol": "SSI",
      "Value": "29100",
      "TradingDate": "14/08/2023",
      "Time": "14:29:59",
      "Open": "29050",
      "High": "29150",
      "Low": "29050",
      "Close": "29100",
      "Volume": "166400"
    }
  ],
  "message": "Success",
  "status": "Success",
  "totalRecord": 2
}
```

### **GET** DailyIndex

```
https://fc-data.ssi.com.vn/api/v2/Market/DailyIndex
```

To get daily trading data of Index

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="98">Type</th><th width="110">Required?</th><th>Description</th></tr></thead><tbody><tr><td>indexId</td><td>string</td><td>Yes</td><td>IndexID. The list of indexes can be retrieved by api getIndexList, or refer to <a href="data-mapping/index-list">this page</a>. </td></tr><tr><td>fromDate</td><td>string</td><td>Yes</td><td>Default is today if not set. In dd/mm/yyyy format</td></tr><tr><td>toDate</td><td>string</td><td>Yes</td><td>Default is today if not set. In dd/mm/yyyy format</td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td><p>From 1 to 10</p><p>Default 1</p></td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td><p>10; 20; 50; 100; 1000</p><p>Default 10</p></td></tr><tr><td>ascending</td><td>boolean</td><td>No</td><td>true/ false</td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="182">Field</th><th width="134">Type</th><th>Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td> </td></tr><tr><td>status</td><td>number</td><td> </td></tr><tr><td>totalRecord</td><td>number</td><td> </td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>Indexcode</td><td>string</td><td> </td></tr><tr><td>IndexValue</td><td>number</td><td> </td></tr><tr><td>Trading Date</td><td>Date</td><td> Format dd/mm/yyyy</td></tr><tr><td>Time</td><td>Timestamp</td><td> </td></tr><tr><td>Change</td><td>number</td><td> </td></tr><tr><td>RatioChange</td><td>number</td><td> </td></tr><tr><td>TotalTrade</td><td>number</td><td>Total trades (both normal and put-through)</td></tr><tr><td>Totalmatchvol</td><td>number</td><td>Total matched vol</td></tr><tr><td>Totalmatchval</td><td>number</td><td>Total matched value</td></tr><tr><td>TypeIndex</td><td>string</td><td></td></tr><tr><td>IndexName</td><td>string</td><td></td></tr><tr><td>Advances</td><td>number</td><td>Total number of symbols with increased price</td></tr><tr><td>Nochanges</td><td>number</td><td>Total number of symbols with unchanged price</td></tr><tr><td>Declines</td><td>number</td><td>Total number of declined symbols</td></tr><tr><td>Ceiling</td><td>number</td><td>Total number of symbols with ceiling price</td></tr><tr><td>Floor</td><td>number</td><td>Total number of symbols with floor price</td></tr><tr><td>Totaldealvol</td><td>number</td><td>Total put-through matched quantity</td></tr><tr><td>Totaldealval</td><td>number</td><td>Total put-through matched value</td></tr><tr><td>Totalvol</td><td>number</td><td>Total matched quantity (both normal and put-through)</td></tr><tr><td>Totalval</td><td>number</td><td>Total matched value (both normal and put-through)</td></tr><tr><td>TradingSession</td><td>string</td><td><a href="data-mapping/trading-session">Trading session</a></td></tr><tr><td>Market</td><td>string</td><td> </td></tr><tr><td>Exchange</td><td>string</td><td>HOSE | HNX</td></tr></tbody></table>
{% endtab %}
{% endtabs %}

**Sample**

```json
Input:
{
  "pageIndex": "1",
  "pageSize": "10",
  "indexID": "HNX30",
  "Fromdate": "14/08/2023",
  "Todate": "14/08/2023"
}
Output:
{
  "data": [
    {
      "IndexId": "HNX30",
      "IndexValue": "510.56",
      "TradingDate": "14/08/2023",
      "Time": null,
      "Change": "19.09",
      "RatioChange": "3.89",
      "TotalTrade": "0",
      "TotalMatchVol": "84693600",
      "TotalMatchVal": "1836008470000",
      "TypeIndex": null,
      "IndexName": "HNX30",
      "Advances": "21",
      "NoChanges": "4",
      "Declines": "5",
      "Ceilings": "2",
      "Floors": "0",
      "TotalDealVol": "2504000",
      "TotalDealVal": "60256000000",
      "TotalVol": "87197600",
      "TotalVal": "1896264470000",
      "TradingSession": "C"
    }
  ],
  "message": "Success",
  "status": "Success",
  "totalRecord": 1
}
```

### **GET** DailyStockPrice

```
 https://fc-data.ssi.com.vn/api/v2/Market/DailyStockPrice
```

To get daily price of securities.&#x20;

**Details**

{% tabs %}
{% tab title="Input" %}

<table><thead><tr><th width="131">Field</th><th width="98">Type</th><th width="113">Required?</th><th>Description</th></tr></thead><tbody><tr><td>Symbol</td><td>string</td><td>No</td><td> </td></tr><tr><td>fromDate</td><td>string</td><td>Yes</td><td>Default is today if not set. In dd/mm/yyyy format</td></tr><tr><td>toDate</td><td>string</td><td>Yes</td><td>Default is today if not set. In dd/mm/yyyy format</td></tr><tr><td>pageIndex</td><td>integer</td><td>Yes</td><td><p>From 1 to 10</p><p>Default 1</p></td></tr><tr><td>pageSize</td><td>integer</td><td>Yes</td><td><p>10; 20; 50; 100; 1000</p><p>Default 10</p></td></tr><tr><td>market</td><td>string</td><td>No</td><td>HOSE|HNX|UPCOM|DER|BOND</td></tr></tbody></table>
{% endtab %}

{% tab title="Output" %}

<table><thead><tr><th width="203">Field</th><th width="125">Type</th><th width="377">Description</th></tr></thead><tbody><tr><td>message</td><td>string</td><td></td></tr><tr><td>status</td><td>number</td><td> </td></tr><tr><td>totalRecord</td><td>number</td><td></td></tr><tr><td><strong>data</strong></td><td>list</td><td>List data</td></tr><tr><td>Tradingdate</td><td>string</td><td></td></tr><tr><td>Symbol</td><td>string</td><td></td></tr><tr><td>Pricechange</td><td>string</td><td></td></tr><tr><td>Perpricechange</td><td>string</td><td></td></tr><tr><td>Ceilingprice</td><td>string</td><td></td></tr><tr><td>Floorprice</td><td>string</td><td></td></tr><tr><td>Refprice</td><td>string</td><td></td></tr><tr><td>Openprice</td><td>string</td><td></td></tr><tr><td>Highestprice</td><td>string</td><td></td></tr><tr><td>Lowestprice</td><td>string</td><td></td></tr><tr><td>Closeprice</td><td>string</td><td></td></tr><tr><td>Averageprice</td><td>string</td><td></td></tr><tr><td>Closepriceadjusted</td><td>string</td><td></td></tr><tr><td>Totalmatchvol</td><td>string</td><td></td></tr><tr><td>Totalmatchval</td><td>string</td><td></td></tr><tr><td>Totaldealval</td><td>string</td><td></td></tr><tr><td>Totaldealvol</td><td>string</td><td></td></tr><tr><td>Foreignbuyvoltotal</td><td>string</td><td></td></tr><tr><td>Foreigncurrentroom</td><td> string</td><td> </td></tr><tr><td>Foreignsellvoltotal</td><td>string</td><td></td></tr><tr><td>Foreignbuyvaltotal</td><td>string</td><td></td></tr><tr><td>Toreignsellvaltotal</td><td>string</td><td></td></tr><tr><td>Totalbuytrade</td><td>string</td><td></td></tr><tr><td>Totalbuytradevol</td><td>string</td><td></td></tr><tr><td>Totalselltrade</td><td>string</td><td></td></tr><tr><td>Totalselltradevol</td><td>string</td><td></td></tr><tr><td>Netforeivol</td><td>string</td><td></td></tr><tr><td>Netforeignval</td><td>string</td><td></td></tr><tr><td>Totaltradedvol</td><td>string</td><td>Total traded vol, including odd lot</td></tr><tr><td>Totaltradedvalue</td><td>string</td><td>Total traded value, including odd lot</td></tr><tr><td>Time</td><td>string</td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

```json
Input:
{
  "pageIndex": "1",
  "pageSize": "10",
  "symbol": "SSI",
  "market": "HOSE",
  "Fromdate": "19/07/2023",
  "Todate": "19/07/2023"
}
Output:
{
"data": [
{
"TradingDate": "19/07/2023",
"PriceChange": "-150",
"PerPriceChange": "-0.70",
"CeilingPrice": "21550",
"FloorPrice": "18750",
"RefPrice": "20150",
"OpenPrice": "20950",
"HighestPrice": "20950",
"LowestPrice": "20000",
"ClosePrice": "20000",
"AveragePrice": "20118",
"ClosePriceAdjusted": "17392",
"TotalMatchVol": "18900",
"TotalMatchVal": "380230000",
"TotalDealVal": "0",
"TotalDealVol": "0",
"ForeignBuyVolTotal": "0",
"ForeignCurrentRoom": "0",
"ForeignSellVolTotal": "0",
"ForeignBuyValTotal": "0",
"ForeignSellValTotal": "0",
"TotalBuyTrade": "0",
"TotalBuyTradeVol": "0",
"TotalSellTrade": "0",
"TotalSellTradeVol": "0",
"NetBuySellVol": "0",
"NetBuySellVal": "0",
"TotalTradedVol": "18900",
"TotalTradedValue": "380230000",
"Symbol": "HUB",
"Time": null
}
],
"message": "Success",
"status": "Success",
"totalRecord": 1
}
```

**Sample**
