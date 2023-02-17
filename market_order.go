package binance

import (
	"fmt"
	"strconv"
)

func (m MarketOrder) query(c *Client) (string, error) {
	var trade string
	pair, ok := c.Symbols[m.Symbol]
	if !ok {
		return trade, fmt.Errorf("invalid symbol %v or improperly initialized client", m.Symbol)
	}
	var qty, orderId, stPrev, respType, recWin string
	if m.Quantity != 0 {
		qty = "&quantity=" + strconv.FormatFloat(m.Quantity, 'f', pair.StepSize, 64)
	} else {
		qty = "&quoteOrderQty=" + strconv.FormatFloat(m.QuoteOrderQty, 'f', pair.QuoteAssetPrecision, 64)
	}
	if m.NewClientOrderId != "" {
		orderId = "&newClientOrderId=" + m.NewClientOrderId
	}
	if m.SelfTradePreventionMode != "" {
		stPrev = "&selfTradePreventionMode=" + m.SelfTradePreventionMode
	}
	if m.OrderRespType != "" {
		respType = "&newOrderRespType=" + string(m.OrderRespType)
	}
	if m.RecvWindow != 0 {
		recWin = "&recvWindow=" + strconv.Itoa(int(m.RecvWindow))
	}

	trade = "symbol=" + m.Symbol +
		"&side=" + string(m.Side) +
		"&type=" + string(MARKET) +
		qty +
		orderId + stPrev + respType + recWin +
		"&timestamp=" + strconv.Itoa(int(m.Timestamp.UnixMilli()))
	return trade, nil
}
