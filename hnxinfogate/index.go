package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
)

// Index MsgType = I,
// là nội dung thông tin về một chỉ số mà hệ thống chỉ số gửi ra
type Index struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessageToIndex creates a Index msg from a quickfix.Message instance
func FromMessageToIndex(m *quickfix.Message) Index {
	return Index{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m Index) ToMessage() *quickfix.Message {
	return m.Message
}

//RouteIndex returns the begin string, message type, and MessageRoute for this Message type
func RouteIndex(router func(msg Index, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToIndex(msg), sessionID)
	}
	return fix44.BeginString, "I", r
}

// GetIndexCode Tag 2
func (m Index) GetIndexCode() (v string, err quickfix.MessageRejectError) {
	var f field.AdvIdField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// TODO: complete this MsgType
