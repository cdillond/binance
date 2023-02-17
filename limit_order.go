package binance

import (
	"fmt"
	"strconv"
)

func (l LimitOrder) query(c *Client) (string, error) {
	var trade string
	pair, ok := c.Symbols[l.Symbol]
	if !ok {
		return trade, fmt.Errorf("invalid symbol %v or improperly initialized client", l.Symbol)
	}
	var ncoi, iqty, stpm, nort, recWin string
	if l.NewClientOrderId != "" {
		ncoi = "&newClientOrderId=" + l.NewClientOrderId
	}
	if l.IcebergQty != 0 {
		// not sure if this is the correct precision
		iqty = "&icebergQty=" + strconv.FormatFloat(l.IcebergQty, 'f', pair.StepSize, 64)
	}
	if l.SelfTradePreventionMode != "" {
		stpm = "&selfTradePreventionMode=" + l.SelfTradePreventionMode
	}
	if l.OrderRespType != "" {
		nort = "&newOrderRespType=" + string(l.OrderRespType)
	}
	if l.RecvWindow != 0 {
		recWin = "&recvWindow=" + strconv.Itoa(int(l.RecvWindow))
	}
	trade = "symbol=" + l.Symbol +
		"&side=" + string(l.Side) +
		"&type=" + string(LIMIT) +
		"&timeInForce=" + string(l.TimeInForce) +
		"&quantity=" + strconv.FormatFloat(l.Quantity, 'f', pair.StepSize, 64) +
		"&price=" + strconv.FormatFloat(l.Price, 'f', pair.TickSize, 64) +
		ncoi + iqty + stpm + nort + recWin +
		"&timestamp=" + strconv.Itoa(int(l.Timestamp.UnixMilli()))
	return trade, nil
}
