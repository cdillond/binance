package binance

import "time"

// Types that satisfy the Order interface include: MarketOrder, LimitOrder, LimitMakerOrder
// StopLossOrder, StopLossLimitOrder, TakeProfitOrder, and TakeProfitLimitOrder.
type Order interface {
	query(c *Client) (string, error) // returns the query string for the trade
}

type MarketOrder struct {
	Symbol                  string  // REQUIRED
	Side                    Side    // REQUIRED
	Quantity                float64 // DEFAULT; MUST BE USED IFF QuoteOrderQty == 0
	QuoteOrderQty           float64 // 0 BY DEFAULT; MUST BE USED IFF Quantity == 0
	NewClientOrderId        string
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

func NewMarketOrder(symbol string, side Side, quantity float64) MarketOrder {
	return MarketOrder{Symbol: symbol, Side: side, Quantity: quantity, Timestamp: time.Now()}
}

type LimitOrder struct {
	Symbol                  string      // REQUIRED
	Side                    Side        // REQUIRED
	TimeInForce             TimeInForce // REQUIRED
	Quantity                float64     // REQUIRED
	Price                   float64     // REQUIRED
	NewClientOrderId        string
	IcebergQty              float64
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

func NewLimitOrder(symbol string, side Side, timeInForce TimeInForce, quantity, price float64) LimitOrder {
	return LimitOrder{
		Symbol:      symbol,
		Side:        side,
		TimeInForce: timeInForce,
		Quantity:    quantity,
		Price:       price,
		Timestamp:   time.Now(),
	}
}

type LimitMakerOrder struct {
	Symbol                  string      // REQUIRED
	Side                    Side        // REQUIRED
	TimeInForce             TimeInForce // REQUIRED
	Quantity                float64     // REQUIRED
	Price                   float64     // REQUIRED
	NewClientOrderId        string
	IcebergQty              float64
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

func NewLimitMakerOrder(symbol string, side Side, timeInForce TimeInForce, quantity, price float64) LimitMakerOrder {
	return LimitMakerOrder{
		Symbol:      symbol,
		Side:        side,
		TimeInForce: timeInForce,
		Quantity:    quantity,
		Price:       price,
		Timestamp:   time.Now(),
	}
}

type StopLossOrder struct {
	Symbol                  string  // REQUIRED
	Quantity                float64 // REQUIRED
	NewClientOrderId        string
	StopPrice               float64 // REQUIRED
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

func NewStopLossOrder(symbol string, quantity, stopPrice float64) StopLossOrder {
	return StopLossOrder{
		Symbol:    symbol,
		Quantity:  quantity,
		StopPrice: stopPrice,
		Timestamp: time.Now(),
	}
}

type StopLossLimitOrder struct {
	Symbol                  string      // REQUIRED
	TimeInForce             TimeInForce // REQUIRED
	Quantity                float64     // REQUIRED
	Price                   float64     // REQUIRED
	NewClientOrderId        string
	StopPrice               float64 // DEFAULT; MUST BE USED IFF TrailingDelta == 0
	TrailingDelta           int     // 0 BY DEFAULT; MUST BE USED IFF StopPrice == 0
	IcebergQty              float64
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

func NewStopLossLimitOrder(symbol string, timeInForce TimeInForce, quantity, price, stopPrice float64) StopLossLimitOrder {
	return StopLossLimitOrder{
		Symbol:      symbol,
		Quantity:    quantity,
		Price:       price,
		TimeInForce: timeInForce,
		StopPrice:   stopPrice,
		Timestamp:   time.Now(),
	}
}

type TakeProfitOrder struct {
	Symbol                  string  // REQUIRED
	Quantity                float64 // REQUIRED
	NewClientOrderId        string
	StopPrice               float64 // DEFAULT; MUST BE USED IFF TrailingDelta == 0
	TrailingDelta           int     // 0 BY DEFAULT; MUST BE USED IFF StopPrice == 0
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

func NewTakeProfitOrder(symbol string, quantity, stopPrice float64) TakeProfitOrder {
	return TakeProfitOrder{
		Symbol:    symbol,
		Quantity:  quantity,
		StopPrice: stopPrice,
		Timestamp: time.Now(),
	}
}

type TakeProfitLimitOrder struct {
	Symbol                  string      // REQUIRED
	TimeInForce             TimeInForce // REQUIRED
	Quantity                float64     // REQUIRED
	Price                   float64     // REQUIRED
	NewClientOrderId        string
	StopPrice               float64 // DEFAULT; MUST BE USED IFF TrailingDelta == 0
	TrailingDelta           int     // 0 BY DEFAULT; MUST BE USED IFF StopPrice == 0
	IcebergQty              float64
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time // REQUIRED
}

func NewTakeProfitLimitOrder(symbol string, timeInForce TimeInForce, quantity, price, stopPrice float64) TakeProfitLimitOrder {
	return TakeProfitLimitOrder{
		Symbol:      symbol,
		Quantity:    quantity,
		Price:       price,
		TimeInForce: timeInForce,
		StopPrice:   stopPrice,
		Timestamp:   time.Now(),
	}
}

type RawOrder struct {
	OrderType               OrderType
	Symbol                  string
	Side                    Side
	Quantity                float64
	QuoteOrderQty           float64
	NewClientOrderId        string
	SelfTradePreventionMode string
	OrderRespType           OrderRespType
	RecvWindow              int
	Timestamp               time.Time
	TimeInForce             TimeInForce
	Price                   float64
	IcebergQty              float64
	StopPrice               float64
	TrailingDelta           int
}

func NewRawOrder() RawOrder {
	return RawOrder{}
}
