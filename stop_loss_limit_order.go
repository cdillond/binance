package binance

import (
	"fmt"
	"strconv"
)

func (s StopLossLimitOrder) query(c *Client) (string, error) {
	var trade string
	pair, ok := c.Symbols[s.Symbol]
	if !ok {
		return trade, fmt.Errorf("invalid symbol %v or improperly initialized client", s.Symbol)
	}
	var stopPrice, td string
	var ncoi, stpm, nort, recWin, iqty string
	if s.NewClientOrderId != "" {
		ncoi = "&newClientOrderId=" + s.NewClientOrderId
	}
	if s.SelfTradePreventionMode != "" {
		stpm = "&selfTradePreventionMode=" + s.SelfTradePreventionMode
	}
	if s.OrderRespType != "" {
		nort = "&newOrderRespType=" + string(s.OrderRespType)
	}
	if s.RecvWindow != 0 {
		recWin = "&recvWindow=" + strconv.Itoa(int(s.RecvWindow))
	}
	if s.IcebergQty != 0 {
		// not sure if this is the correct precision
		iqty = "&icebergQty=" + strconv.FormatFloat(s.IcebergQty, 'f', pair.StepSize, 64)
	}
	if s.StopPrice != 0 {
		stopPrice = "&stopPrice=" + strconv.FormatFloat(s.StopPrice, 'f', pair.TickSize, 64)
	} else if s.TrailingDelta != 0 {
		td = "&trailingDelta=" + strconv.Itoa(int(s.TrailingDelta))
	}
	trade = "symbol=" + s.Symbol +
		"&type=" + string(STOP_LOSS_LIMIT) +
		"&side=" + string(SELL) + // ODD THAT THIS IS REQUIRED?
		"&quantity=" + strconv.FormatFloat(s.Quantity, 'f', pair.StepSize, 64) +
		"&price=" + strconv.FormatFloat(s.Price, 'f', pair.TickSize, 64) +
		"&timeInForce=" + string(s.TimeInForce) +
		stopPrice + td +
		iqty + ncoi + stpm + nort + recWin +
		"&timestamp=" + strconv.Itoa(int(s.Timestamp.UnixMilli()))
	return trade, nil
}
