package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
)

// TopNPrice MsgType = TP.
// Gồm nhiều bước giá do hệ thống HNX tính toán
type TopNPrice struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a TopNPrice from a quickfix.Message instance
func FromMessageToTopNPrice(m *quickfix.Message) TopNPrice {
	return TopNPrice{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m TopNPrice) ToMessage() *quickfix.Message {
	return m.Message
}

//Route returns the beginstring, message type, and MessageRoute for this Message type
func RouteTopNPrice(router func(msg TopNPrice, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToTopNPrice(msg), sessionID)
	}
	return fix44.BeginString, "TP", r
}

// GetSymbol Tag 55
func (m TopNPrice) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBoardCode Tag 425
//Mã bảng của chứng khoán:
//	-LIS_BRD_01,..: bảng niêm yết
//	-UPC_BRD_01,…: bảng upcom
func (m TopNPrice) GetBoardCode() (v string, err quickfix.MessageRejectError) {
	var f BoardCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// Tag 555.
// Số mức giá tốt nhất
func (m StockInfo) GetNOTopPrice() (v float64, err quickfix.MessageRejectError) {
	var f NOTopPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type NOTopPriceField struct{ quickfix.FIXDecimal }

func (f NOTopPriceField) Tag() quickfix.Tag      { return 555 }
func (f NOTopPriceField) Value() decimal.Decimal { return f.Decimal }

type BidAskRepeatingGroup struct{ *quickfix.RepeatingGroup }

func (m BidAskRepeatingGroup) Add() BidAskGroup {
	g := m.RepeatingGroup.Add()
	return BidAskGroup{g}
}
func (m BidAskRepeatingGroup) Get(i int) BidAskGroup {
	return BidAskGroup{m.RepeatingGroup.Get(i)}
}

type BidAskGroup struct{ *quickfix.Group }

// Tag 556
// Số thứ tự của mức giá: 1, 2, 3, 4, ..
func (m BidAskGroup) GetNumTopPrice() (v float64, err quickfix.MessageRejectError) {
	var f NumTopPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type NumTopPriceField struct{ quickfix.FIXDecimal }

func (f NumTopPriceField) Tag() quickfix.Tag      { return 556 }
func (f NumTopPriceField) Value() decimal.Decimal { return f.Decimal }

// Tag 132
// Giá mua
func (m BidAskGroup) GetBestBidPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.BidPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// Tag 1321
// Khối lượng mua
func (m BidAskGroup) GetBestBidQtty() (v float64, err quickfix.MessageRejectError) {
	var f field.DerivativeCapPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// Tag 133
// Giá bán
func (m BidAskGroup) GetBestOfferPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.OfferPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// Tag 1331
// Khối lượng bán
func (m BidAskGroup) GetBestOfferQtty() (v float64, err quickfix.MessageRejectError) {
	var f BestOfferQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}
