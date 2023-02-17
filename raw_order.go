package binance

import (
	"fmt"
	"strconv"
)

func (o RawOrder) query(c *Client) (string, error) {
	var trade string
	pair, ok := c.Symbols[o.Symbol]
	if !ok {
		return trade, fmt.Errorf("invalid symbol %v or improperly initialized client", o.Symbol)
	}
	// required symbol, side, and timestamp
	var qty, qoqty, price, iqty, stopP string
	var ncoi, stpm, ort, recWin, tdelta string
	var tif, sym, side, typ, timestamp string
	if o.Symbol != "" {
		sym = "symbol=" + o.Symbol
	}
	if o.Side != "" {
		side = "&side=" + string(o.Side)
	}
	if o.OrderType != "" {
		typ = "&type=" + string(o.OrderType)
	}
	if !o.Timestamp.IsZero() {
		timestamp = "&timestamp=" + strconv.Itoa(int(o.Timestamp.UnixMilli()))
	}
	if o.Quantity != 0 {
		qty = "&quantity=" + strconv.FormatFloat(o.Quantity, 'f', pair.StepSize, 64)
	}
	if o.QuoteOrderQty != 0 {
		qoqty = "&quoteOrderQty=" + strconv.FormatFloat(o.QuoteOrderQty, 'f', pair.TickSize, 64)
	}
	if o.NewClientOrderId != "" {
		ncoi = "&newClientOrderId=" + o.NewClientOrderId
	}
	if o.SelfTradePreventionMode != "" {
		stpm = "&selfTradePreventionMode=" + o.SelfTradePreventionMode
	}
	if o.OrderRespType != "" {
		ort = "&newOrderRespType=" + string(o.OrderRespType)
	}
	if o.RecvWindow != 0 {
		recWin = "&recvWindow=" + strconv.Itoa(int(o.RecvWindow))
	}
	if o.IcebergQty != 0 {
		// not sure if this is the correct precision
		iqty = "&icebergQty=" + strconv.FormatFloat(o.IcebergQty, 'f', pair.StepSize, 64)
	}
	if o.Price != 0 {
		price = "&price=" + strconv.FormatFloat(o.Price, 'f', pair.TickSize, 64)
	}
	if o.TimeInForce != "" {
		tif = "&timeInForce=" + string(o.TimeInForce)
	}
	if o.StopPrice != 0 {
		stopP = "stopPrice=" + strconv.FormatFloat(o.StopPrice, 'f', pair.TickSize, 64)
	}
	if o.TrailingDelta != 0 {
		tdelta = recWin + "&trailingDelta=" + strconv.Itoa(int(o.TrailingDelta))
	}

	trade = sym + side + typ + qty + qoqty +
		price + iqty + stopP + ncoi +
		stpm + ort + recWin + tdelta +
		tif + timestamp

	return trade, nil
}
