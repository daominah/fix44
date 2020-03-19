package hnxinfogate

import (
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
)

// BoardInfo MsgType = BI,
// là thông tin về bảng giao dịch
type BoardInfo struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessageToBoardInfo creates a BoardInfo msg from a quickfix.Message instance
func FromMessageToBoardInfo(m *quickfix.Message) BoardInfo {
	return BoardInfo{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m BoardInfo) ToMessage() *quickfix.Message {
	return m.Message
}

//RouteIndex returns the begin string, message type, and MessageRoute for this Message type
func RouteBoardInfo(router func(msg BoardInfo, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToBoardInfo(msg), sessionID)
	}
	return fix44.BeginString, "BI", r
}

// GetMarketCode Tag 341
func (m BoardInfo) GetMarketCode() (v string, err quickfix.MessageRejectError) {
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
