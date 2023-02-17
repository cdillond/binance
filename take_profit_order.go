package binance

import (
	"fmt"
	"strconv"
)

func (t TakeProfitOrder) query(c *Client) (string, error) {
	var trade string
	pair, ok := c.Symbols[t.Symbol]
	if !ok {
		return trade, fmt.Errorf("invalid symbol %v or improperly initialized client", t.Symbol)
	}
	var ncoi, stpm, nort, recWin, stopPrice, td string
	if t.NewClientOrderId != "" {
		ncoi = "&newClientOrderId=" + t.NewClientOrderId
	}
	if t.SelfTradePreventionMode != "" {
		stpm = "&selfTradePreventionMode=" + t.SelfTradePreventionMode
	}
	if t.OrderRespType != "" {
		nort = "&newOrderRespType=" + string(t.OrderRespType)
	}
	if t.RecvWindow != 0 {
		recWin = "&recvWindow=" + strconv.Itoa(int(t.RecvWindow))
	}
	if t.StopPrice != 0 {
		stopPrice = "&stopPrice=" + strconv.FormatFloat(t.StopPrice, 'f', pair.TickSize, 64)
	} else if t.TrailingDelta != 0 {
		td = "&trailingDelta=" + strconv.Itoa(int(t.TrailingDelta))
	}
	trade = "symbol=" + t.Symbol +
		"&type=" + string(TAKE_PROFIT) +
		"&side=" + string(SELL) + // ODD THAT THIS IS REQUIRED?
		"&quantity=" + strconv.FormatFloat(t.Quantity, 'f', pair.StepSize, 64) +
		stopPrice + td +
		ncoi + stpm + nort + recWin +
		"&timestamp=" + strconv.Itoa(int(t.Timestamp.UnixMilli()))
	return trade, nil
}
