package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
)

// AuctionMatch MsgType = EP.
// giá tạm khớp trong phiên khớp lệnh định kỳ sẽ gửi cho các thành viên giao dịch
type AuctionMatch struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a AuctionMatch from a quickfix.Message instance
func FromMessageToAuctionMatch(m *quickfix.Message) AuctionMatch {
	return AuctionMatch{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m AuctionMatch) ToMessage() *quickfix.Message {
	return m.Message
}

//Route returns the beginstring, message type, and MessageRoute for this Message type
func RouteAuctionMatch(router func(msg AuctionMatch, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToAuctionMatch(msg), sessionID)
	}
	return fix44.BeginString, "EP", r
}

// GetSymbol Tag 55
func (m AuctionMatch) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetActionType Tag 33
//Loại khớp : A : khớp chính (không có ở phái sinh), M : tạm khớp
func (m AuctionMatch) GetActionType() (v string, err quickfix.MessageRejectError) {
	var f ActionTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type ActionTypeField struct{ quickfix.FIXString }

func (f ActionTypeField) Tag() quickfix.Tag { return 33 }
func (f ActionTypeField) Value() string     { return f.String() }

// GetPrice Tag 31,
// Giá khớp (định kỳ)
func (m AuctionMatch) GetPrice() (v float64, err quickfix.MessageRejectError) {
	var f PriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PriceField struct{ quickfix.FIXDecimal }

func (f PriceField) Tag() quickfix.Tag      { return 31 }
func (f PriceField) Value() decimal.Decimal { return f.Decimal }

// GetQtty Tag 32,
// Khối lượng khớp (định kỳ), với dữ liệu Phái sinh thì không có tag này
func (m AuctionMatch) GetQtty() (v float64, err quickfix.MessageRejectError) {
	var f QttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type QttyField struct{ quickfix.FIXDecimal }

func (f QttyField) Tag() quickfix.Tag      { return 32 }
func (f QttyField) Value() decimal.Decimal { return f.Decimal }
