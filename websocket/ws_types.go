// Package websocket defines types for responses sent by the binance websocket api.
// An actual websocket client is not included in this package.
package websocket

type StreamType interface {
	AggTrade | KlineStream | TradeData | Ticker | MiniTicker |
		WSPriceChange | []MiniTicker | []Ticker | BookTicker | BookDepth |
		BookDepthUpdate | AccountUpdate | BalanceUpdate | OrderUpdatePayload
}

type Kline struct {
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
	EventType string `json:"e"` // Event type
	EventTime int    `json:"E"` // Event time
	Symbol    string `json:"s"` // Symbol
	K         Kline  `json:"k"`
}

type StreamResp[S StreamType] struct {
	Stream string `json:"stream"`
	Data   S      `json:"data"`
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

type TradeData struct {
	EventType     string  `json:"e"`
	EventTime     int     `json:"E"`
	Symbol        string  `json:"s"`
	TradeId       int     `json:"t"`
	Price         float64 `json:"p,string"`
	Quantity      float64 `json:"q,string"`
	BuyerOrderId  int     `json:"b"`
	SellerOrderId int     `json:"a"`
	TradeTime     int     `json:"T"`
	IsBuyerMaker  bool    `json:"m"`
}

type WSPriceChange struct {
	EventType        string  `json:"e"`
	EventTime        int     `json:"E"`
	Symbol           string  `json:"s"`
	PriceChange      float64 `json:"p,string"`
	PriceChangePct   float64 `json:"P,string"`
	WeightedAvgPrice float64 `json:"w,string"`
	FirstTradePrice  float64 `json:"x,string"`
	LastPrice        float64 `json:"c,string"`
	LastQty          float64 `json:"Q,string"`
	BestBidPrice     float64 `json:"b,string"`
	BestBidQty       float64 `json:"B,string"`
	BestAskPrice     float64 `json:"a,string"`
	BestAskQty       float64 `json:"A,string"`
	OpenPrice        float64 `json:"o,string"`
	HighPrice        float64 `json:"h,string"`
	LowPrice         float64 `json:"l,string"`
	BaseAssetVol     float64 `json:"v,string"`
	QuoteAssetVol    float64 `json:"q,string"`
	OpenTime         int     `json:"O"`
	CloseTime        int     `json:"C"`
	FirstTradeId     int     `json:"F"`
	LastTradeId      int     `json:"L"`
	NumTrades        int     `json:"n"`
}

type BookTicker struct {
	BookUpdateId int     `json:"u"`
	Symbol       string  `json:"s"`
	BestBidPrice float64 `json:"b,string"`
	BestBidQty   float64 `json:"B,string"`
	BestAskPrice float64 `json:"a,string"`
	BestAskQty   float64 `json:"A,string"`
}

type BookDepth struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type BookDepthUpdate struct {
	EventType     string     `json:"e"`
	EventTime     int        `json:"E"`
	Symbol        string     `json:"s"`
	FirstUpdateId int        `json:"U"`
	FinalUpdateId int        `json:"u"`
	UpdatedBids   [][]string `json:"b"`
	UpdatedAsks   [][]string `json:"a"`
}

type AccountUpdate struct {
	EventType      string `json:"e"`
	EventTime      int    `json:"E"`
	LastUpdateTime int    `json:"u"`
	Balances       []struct {
		Asset  string  `json:"a"`
		Free   float64 `json:"f"`
		Locked float64 `json:"l"`
	} `json:"B"`
}

type OrderUpdatePayload interface {
	ExecutionReport | ListStatus
}

type ExecutionReport struct {
	EventType         int     `json:"e"`
	EventTime         int     `json:"E"`
	Symbol            string  `json:"s"`
	ClientOrderId     string  `json:"c"`
	Side              string  `json:"S"`
	OrderType         string  `json:"o"`
	TimeInForce       string  `json:"f"`
	OrderQty          float64 `json:"q,string"`
	OrderPrice        float64 `json:"p,string"`
	StopPrice         float64 `json:"P,string"`
	TrailingDelta     int     `json:"d"`
	IcebergQty        float64 `json:"F,string"`
	OrderListId       int     `json:"g"`
	OrigClientOrderId string  `json:"C"`
	ExecType          string  `json:"x"`
	OrderStatus       string  `json:"X"`
	RejectReason      string  `json:"r"`
	OrderID           int     `json:"i"`
	LastExecQty       float64 `json:"l,string"`
	CumQty            float64 `json:"z,string"`
	LastExecPrc       float64 `json:"L,string"`
	CommissionAmt     float64 `json:"n,string"`
	CommissionAsset   string  `json:"N"`
	TransactTime      int     `json:"T"`
	TradeID           int     `json:"t"`

	OnBook       bool `json:"w"`
	IsTradeMaker bool `json:"m"`

	OrderCreationTime       int     `json:"O"`
	CumQuoteQty             float64 `json:"Z,string"` // Cumulative quote asset transacted quantity
	LastQuoteAmt            float64 `json:"Y,string"` // Last quote asset transacted quantity (i.e. lastPrice * lastQty)
	QuoteOrderQty           float64 `json:"Q,string"`
	SelfTradePreventionMode string  `json:"V"`
	TrailingTime            string  `json:"D"`        // (Appears if the trailing stop order is active)
	WorkingTime             string  `json:"W"`        // (Appears if the order is working on the order book)
	TradeGroupID            int     `json:"u"`        // (Appears if the order is working on the order book)
	PreventMatchId          int     `json:"v"`        // (Appears if the order has expired due to STP trigger)
	CounterOrderId          int     `json:"U"`        //  (Appears if the order has expired due to STP trigger)
	PreventedQty            float64 `json:"A,string"` // (Appears if the order has expired due to STP trigger)
	LastPreventedQty        float64 `json:"B,string"` // (Appears if the order has expired due to STP trigger)
}

type BalanceUpdate struct {
	EventType    string  `json:"e"` // Event type
	EventTime    int     `json:"E"` // Event time
	Asset        string  `json:"a"` // Asset
	BalanceDelta float64 `json:"d,string"`
	ClearTime    int     `json:"T"`
}

type ListStatus struct {
	EventType       int    `json:"e"`
	EventTime       int    `json:"E"`
	Symbol          string `json:"s"`
	OrderListId     int    `json:"g"`
	ContingencyType string `json:"c"`
	ListStatusType  string `json:"l"`
	ListOrderStatus string `json:"L"`
	RejectReason    string `json:"r"`
	ClientOrderId   string `json:"C"`
	TransactTime    int    `json:"T"`
	Orders          []struct {
		Symbol        string `json:"s"`
		OrderId       int    `json:"i"`
		ClientOrderId string `json:"c"`
	} `json:"O"`
}
