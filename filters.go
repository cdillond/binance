package binance

type Filter string

const (
	PRICE_FILTER           = "PRICE_FILTER"
	PERCENT_PRICE          = "PERCENT_PRICE"
	PERCENT_PRICE_BY_SIDE  = "PERCENT_PRICE_BY_SIDE"
	LOT_SIZE               = "LOT_SIZE"
	NOTIONAL               = "NOTIONAL"
	MIN_NOTIONAL           = "MIN_NOTIONAL"
	ICEBERG_PARTS          = "ICEBERG_PARTS"
	MARKET_LOT_SIZE        = "MARKET_LOT_SIZE"
	MAX_NUM_ORDERS         = "MAX_NUM_ORDERS"
	MAX_NUM_ALGO_ORDERS    = "MAX_NUM_ALGO_ORDERS"
	MAX_NUM_ICEBERG_ORDERS = "MAX_NUM_ICEBERG_ORDERS"
	MAX_POSITION           = "MAX_POSITION"
	TRAILING_DELTA         = "TRAILING_DELTA"
)
const (
	EXCHANGE_MAX_NUM_ORDERS         = "EXCHANGE_MAX_NUM_ORDERS"
	EXCHANGE_MAX_ALGO_ORDERS        = "EXCHANGE_MAX_ALGO_ORDERS"
	EXCHANGE_MAX_NUM_ICEBERG_ORDERS = "EXCHANGE_MAX_NUM_ICEBERG_ORDERS"
)

type RawFilter struct {
	FilterType            string `json:"filterType"`
	MinPrice              string `json:"minPrice"`
	MaxPrice              string `json:"maxPrice"`
	TickSize              string `json:"tickSize"`
	MultiplierUp          string `json:"multiplierUp"`
	MultiplierDown        string `json:"multiplierDown"`
	AvgPriceMins          int    `json:"avgPriceMins"`
	BidMultiplierUp       string `json:"bidMultiplierUp"`
	BidMultiplierDown     string `json:"bidMultiplierDown"`
	AskMultiplierUp       string `json:"askMultiplierUp"`
	AskMultiplierDown     string `json:"askMultiplierDown"`
	MinQty                string `json:"minQty"`
	MaxQty                string `json:"maxQty"`
	StepSize              string `json:"stepSize"`
	MinNotional           string `json:"minNotional"`
	ApplyMinToMarket      bool   `json:"applyMinToMarket"`
	MaxNotional           string `json:"maxNotional"`
	ApplyMaxToMarket      bool   `json:"applyMaxToMarket"`
	Limit                 int    `json:"limit"`
	MaxNumAlgoOrders      int    `json:"maxNumAlgoOrders"`
	MaxNumIcebergOrders   int    `json:"maxNumIcebergOrders"`
	MaxPosition           string `json:"maxPosition"`
	MinTrailingAboveDelta int    `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int    `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta"`
	MaxNumOrders          int    `json:"maxNumOrders"`
}

// Price Filter
type PriceFilter struct {
	FilterType string `json:"filterType"`
	MinPrice   string `json:"minPrice"`
	MaxPrice   string `json:"maxPrice"`
	TickSize   string `json:"tickSize"`
}

// Percent Price
type PercentPrice struct {
	FilterType     string `json:"filterType"`
	MultiplierUp   string `json:"multiplierUp"`
	MultiplierDown string `json:"multiplierDown"`
	AvgPriceMins   int    `json:"avgPriceMins"`
}

// Percent Price By Side
type PercentPriceBySide struct {
	FilterType        string `json:"filterType"`
	BidMultiplierUp   string `json:"bidMultiplierUp"`
	BidMultiplierDown string `json:"bidMultiplierDown"`
	AskMultiplierUp   string `json:"askMultiplierUp"`
	AskMultiplierDown string `json:"askMultiplierDown"`
	AvgPriceMins      int    `json:"avgPriceMins"`
}

// Lot Size
type LotSize struct {
	FilterType string `json:"filterType"`
	MinQty     string `json:"minQty"`
	MaxQty     string `json:"maxQty"`
	StepSize   string `json:"stepSize"`
}

// Type: Notional
type Notional struct {
	FilterType       string `json:"filterType"`
	MinNotional      string `json:"minNotional"`
	ApplyMinToMarket bool   `json:"applyMinToMarket"`
	MaxNotional      string `json:"maxNotional"`
	ApplyMaxToMarket bool   `json:"applyMaxToMarket"`
	AvgPriceMins     int    `json:"avgPriceMins"`
}

// Type: MinNotional
type MinNotional struct {
	FilterType    string `json:"filterType"`
	MinNotional   string `json:"minNotional"`
	ApplyToMarket bool   `json:"applyToMarket"`
	AvgPriceMins  int    `json:"avgPriceMins"`
}

type IcebergParts struct {
	FilterType string `json:"filterType"`
	Limit      int    `json:"limit"`
}

type MarketLotSize struct {
	FilterType string `json:"filterType"`
	MinQty     string `json:"minQty"`
	MaxQty     string `json:"maxQty"`
	StepSize   string `json:"stepSize"`
}

type MaxNumOrders struct {
	FilterType string `json:"filterType"`
	Limit      int    `json:"limit"`
}

type MaxNumAlgoOrders struct {
	FilterType       string `json:"filterType"`
	MaxNumAlgoOrders int    `json:"maxNumAlgoOrders"`
}

type MaxNumIcebergOrders struct {
	FilterType          string `json:"filterType"`
	MaxNumIcebergOrders int    `json:"maxNumIcebergOrders"`
}

type MaxPosition struct {
	FilterType  string `json:"filterType"`
	MaxPosition string `json:"maxPosition"`
}

type TrailingDelta struct {
	FilterType            string `json:"filterType"`
	MinTrailingAboveDelta int    `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int    `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta"`
}

type ExchangeMaxNumOrders struct {
	FilterType   string `json:"filterType"`
	MaxNumOrders int    `json:"maxNumOrders"`
}

type ExchangeMaxAlgoOrders struct {
	FilterType       string `json:"filterType"`
	MaxNumAlgoOrders int    `json:"maxNumAlgoOrders"`
}

type ExchangeMaxNumIcebergOrders struct {
	FilterType          string `json:"filterType"`
	MaxNumIcebergOrders int    `json:"maxNumIcebergOrders"`
}
