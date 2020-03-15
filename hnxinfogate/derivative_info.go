package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
)

// DerivativeInfo MsgType = DI.
// giá tạm khớp trong phiên khớp lệnh định kỳ sẽ gửi cho các thành viên giao dịch
type DerivativeInfo struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a DerivativeInfo from a quickfix.Message instance
func FromMessageToDerivativeInfo(m *quickfix.Message) DerivativeInfo {
	return DerivativeInfo{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m DerivativeInfo) ToMessage() *quickfix.Message {
	return m.Message
}

//Route returns the beginstring, message type, and MessageRoute for this Message type
func RouteDerivativeInfo(router func(msg DerivativeInfo, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToDerivativeInfo(msg), sessionID)
	}
	return fix44.BeginString, "EP", r
}

// GetSymbol Tag 55
func (m DerivativeInfo) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}
