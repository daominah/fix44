package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
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

// GetValue Tag 3
// Giá trị chỉ số tại thời điểm hiện tại
// Giá trị TRI tại thời điểm hiện tại
// Giá trị DPI trong ngày
func (m Index) GetValue() (v float64, err quickfix.MessageRejectError) {
	var f ValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type ValueField struct{ quickfix.FIXDecimal }

func (f ValueField) Tag() quickfix.Tag      { return 3 }
func (f ValueField) Value() decimal.Decimal { return f.Decimal }

// GetChange Tag 5
// Giá trị thay đổi chỉ số hoặc TRI so với ngày hôm trước
func (m Index) GetChange() (v float64, err quickfix.MessageRejectError) {
	var f ChangeField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type ChangeField struct{ quickfix.FIXDecimal }

func (f ChangeField) Tag() quickfix.Tag      { return 5 }
func (f ChangeField) Value() decimal.Decimal { return f.Decimal }

// GetRatioChange Tag 6
// Tỷ lệ (%) thay đổi chỉ số hoặc TRI
func (m Index) GetRatioChange() (v float64, err quickfix.MessageRejectError) {
	var f RatioChangeField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type RatioChangeField struct{ quickfix.FIXDecimal }

func (f RatioChangeField) Tag() quickfix.Tag      { return 6 }
func (f RatioChangeField) Value() decimal.Decimal { return f.Decimal }

// GetTotalQtty Tag 7
// Tổng khối lượng giao dịch của khớp lệnh thông thường (lô chẵn)
func (m Index) GetTotalQtty() (v float64, err quickfix.MessageRejectError) {
	var f TotalQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalQttyField struct{ quickfix.FIXDecimal }

func (f TotalQttyField) Tag() quickfix.Tag      { return 7 }
func (f TotalQttyField) Value() decimal.Decimal { return f.Decimal }

// GetTotalValue Tag 14
// Tổng giá trị giao dịch của khớp lệnh thông thường (lô chẵn)
func (m Index) GetTotalValue() (v float64, err quickfix.MessageRejectError) {
	var f TotalValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalValueField struct{ quickfix.FIXDecimal }

func (f TotalValueField) Tag() quickfix.Tag      { return 14 }
func (f TotalValueField) Value() decimal.Decimal { return f.Decimal }

// GetPriorIndexVal Tag 23
//
func (m Index) GetPriorIndexVal() (v float64, err quickfix.MessageRejectError) {
	var f PriorIndexValField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PriorIndexValField struct{ quickfix.FIXDecimal }

func (f PriorIndexValField) Tag() quickfix.Tag      { return 23 }
func (f PriorIndexValField) Value() decimal.Decimal { return f.Decimal }

// GetHighestIndex Tag 24
//
func (m Index) GetHighestIndex() (v float64, err quickfix.MessageRejectError) {
	var f HighestIndexField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type HighestIndexField struct{ quickfix.FIXDecimal }

func (f HighestIndexField) Tag() quickfix.Tag      { return 24 }
func (f HighestIndexField) Value() decimal.Decimal { return f.Decimal }

// GetTotalValue Tag 25
//
func (m Index) GetLowestIndex() (v float64, err quickfix.MessageRejectError) {
	var f TotalValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type LowestIndexField struct{ quickfix.FIXDecimal }

func (f LowestIndexField) Tag() quickfix.Tag      { return 25 }
func (f LowestIndexField) Value() decimal.Decimal { return f.Decimal }
