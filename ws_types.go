// Defines types for responses sent by the binance websocket api.
// The actual websocket client is not included in this package.
package binance

type WsKline struct {
	StartTime             int     `json:"t"`        // Kline start time
	CloseTime             int     `json:"T"`        // Kline close time
	Symbol                string  `json:"s"`        // Symbol
	Interval              string  `json:"i"`        // Interval
	FirstTradeId          int     `json:"f"`        // First trade ID
	LastTradeId           int     `json:"L"`        // Last trade ID
	OpenPrice             float64 `json:"o,string"` // Open price
	ClosePrice            float64 `json:"c,string"` // Close price
	HighPrice             float64 `json:"h,string"` // High price
	LowPrice              float64 `json:"l,string"` // Low price
	BaseAssetVol          float64 `json:"v,string"` // Base asset volume
	NumTrades             int     `json:"n"`        // Number of trades
	IsClosed              bool    `json:"x"`        // Is this kline closed?
	QuoteAssetVol         float64 `json:"q,string"` // Quote asset volume
	TakerBuyBaseAssetVol  float64 `json:"V,string"` // Taker buy base asset volume
	TakerBuyQuoteAssetVol float64 `json:"Q,string"` // Taker buy quote asset volume
}

type KlineStream struct {
	EventType string  `json:"e"` // Event type
	EventTime int     `json:"E"` // Event time
	Symbol    string  `json:"s"` // Symbol
	K         WsKline `json:"k"`
}

type StreamResp struct {
	Stream string      `json:"stream"`
	Data   KlineStream `json:"data"`
}

type AggTrade struct {
	EventType    string  `json:"e"`        // Event type
	EventTime    int     `json:"E"`        // Event time
	Symbol       string  `json:"s"`        // Symbol
	AggTradeId   int     `json:"a"`        // Aggregate trade ID
	Price        float64 `json:"p,string"` // Price
	Quantity     float64 `json:"q,string"` // Quantity
	FirstTradeId int     `json:"f"`        // First trade ID
	LastTradeId  int     `json:"l"`        // Last trade ID
	TradeTime    int     `json:"T"`        // Trade time
	IsBuyerMaker bool    `json:"m"`        // Is the buyer the market maker?
	//"M": true         // Ignore
}

type MiniTicker struct {
	EventType     string  `json:"e"`
	EventTime     int     `json:"E"`
	Symbol        string  `json:"s"`
	Price         float64 `json:"c,string"`
	OpenPrice     float64 `json:"o,string"`
	HighPrice     float64 `json:"h,string"`
	LowPrice      float64 `json:"l,string"`
	BaseAssetVol  float64 `json:"v,string"`
	QuoteAssetVol float64 `json:"q,string"`
}

type Ticker struct {
	EventType        string  `json:"e"`
	EventTime        int     `json:"E"`
	Symbol           string  `json:"s"`
	PriceChange      float64 `json:"p,string"`
	PriceChangePct   float64 `json:"P,string"`
	OpenPrice        float64 `json:"o,string"`
	HighPrice        float64 `json:"h,string"`
	LowPrice         float64 `json:"l,string"`
	Price            float64 `json:"c,string"`
	WeightedAvgPrice float64 `json:"w,string"`
	BaseAssetVol     float64 `json:"v,string"`
	QuoteAssetVol    float64 `json:"q,string"`
	OpenTime         int     `json:"O"`
	CloseTime        int     `json:"C"`
	FirstTradeId     int     `json:"F"`
	LastTradeId      int     `json:"L"`
	NumTrades        int     `json:"n"`
}
