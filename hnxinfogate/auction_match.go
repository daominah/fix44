package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
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
