package binance

import (
	"fmt"
	"strconv"
)

func (s StopLossOrder) query(c *Client) (string, error) {
	var trade string
	pair, ok := c.Symbols[s.Symbol]
	if !ok {
		return trade, fmt.Errorf("invalid symbol %v or improperly initialized client", s.Symbol)
	}
	var ncoi, stpm, nort, recWin string
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
	trade = "symbol=" + s.Symbol +
		"&type=" + string(STOP_LOSS) +
		"&side=" + string(SELL) + // ODD THAT THIS IS REQUIRED?
		"&quantity=" + strconv.FormatFloat(s.Quantity, 'f', pair.StepSize, 64) +
		"&stopPrice=" + strconv.FormatFloat(s.StopPrice, 'f', pair.TickSize, 64) +
		ncoi + stpm + nort + recWin +
		"&timestamp=" + strconv.Itoa(int(s.Timestamp.UnixMilli()))
	return trade, nil
}
