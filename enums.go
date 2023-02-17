package binance

type BaseUrl string

const (
	BINANCE_US   BaseUrl = "https://api.binance.us"
	BINANCE_COM  BaseUrl = "https://api.binance.com"
	BINANCE_COM1 BaseUrl = "https://api1.binance.com"
	BINANCE_COM2 BaseUrl = "https://api2.binance.com"
	BINANCE_COM3 BaseUrl = "https://api3.binance.com"
	BINANCE_COM4 BaseUrl = "https://api4.binance.com"
)

type OrderType string

const (
	LIMIT             OrderType = "LIMIT"
	LIMIT_MAKER       OrderType = "LIMIT_MAKER"
	MARKET            OrderType = "MARKET"
	STOP_LOSS         OrderType = "STOP_LOSS"
	STOP_LOSS_LIMIT   OrderType = "STOP_LOSS_LIMIT"
	TAKE_PROFIT       OrderType = "TAKE_PROFIT"
	TAKE_PROFIT_LIMIT OrderType = "TAKE_PROFIT_LIMIT"
)

type OrderRespType string

const (
	ACK    OrderRespType = "ACK"
	FULL   OrderRespType = "FULL"
	RESULT OrderRespType = "RESULT"
)

type SelfTradePreventionMode string

const (
	EXPIRE_BOTH  SelfTradePreventionMode = "EXPIRE_BOTH"
	EXPIRE_MAKER SelfTradePreventionMode = "EXPIRE_MAKER"
	EXPIRE_TAKER SelfTradePreventionMode = "EXPIRE_TAKER"
)

type Side string

const (
	BUY  Side = "BUY"
	SELL Side = "SELL"
)

type TimeInForce string

const (
	GTC TimeInForce = "GTC"
	FOK TimeInForce = "FOK"
	IOC TimeInForce = "IOC"
)

type Permission string

const (
	SPOT        Permission = "SPOT"
	MARGIN      Permission = "MARGIN"
	LEVERAGED   Permission = "LEVERAGED"
	TRD_GRP_002 Permission = "TRD_GRP_002" // BINANCE.COM ONLY
	TRD_GRP_003 Permission = "TRD_GRP_003" // BINANCE.COM ONLY
	TRD_GRP_004 Permission = "TRD_GRP_004" // BINANCE.COM ONLY
	TRD_GRP_005 Permission = "TRD_GRP_005" // BINANCE.COM ONLY
	TRD_GRP_006 Permission = "TRD_GRP_006" // BINANCE.COM ONLY
	TRD_GRP_007 Permission = "TRD_GRP_007" // BINANCE.COM ONLY
)

type IntervalUnit string

const (
	MINUTE IntervalUnit = "m"
	HOUR   IntervalUnit = "h"
	DAY    IntervalUnit = "d"
)
