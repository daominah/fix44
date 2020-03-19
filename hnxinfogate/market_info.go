package hnxinfogate

import (
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
)

// MarketInfo MsgType = MI,
// là nội dung thông tin về một chỉ số mà hệ thống chỉ số gửi ra
type MarketInfo struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessageToMarketInfo creates a MarketInfo msg from a quickfix.Message instance
func FromMessageToMarketInfo(m *quickfix.Message) MarketInfo {
	return MarketInfo{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m MarketInfo) ToMessage() *quickfix.Message {
	return m.Message
}

//RouteIndex returns the begin string, message type, and MessageRoute for this Message type
func RouteMarketInfo(router func(msg MarketInfo, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToMarketInfo(msg), sessionID)
	}
	return fix44.BeginString, "MI", r
}

// GetMarketCode Tag 341
func (m MarketInfo) GetMarketCode() (v string, err quickfix.MessageRejectError) {
	var f MarketCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type MarketCodeField struct{ quickfix.FIXString }

func (f MarketCodeField) Tag() quickfix.Tag { return 341 }
func (f MarketCodeField) Value() string     { return f.String() }

// TODO: complete this MsgType
