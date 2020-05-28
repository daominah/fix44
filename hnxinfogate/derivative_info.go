package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
)

// DerivativeInfo MsgType = DI.
// giá tạm khớp trong phiên khớp lệnh định kỳ sẽ gửi cho các thành viên giao dịch
type DerivativeInfo struct {
	*StockInfo
}

// FromMessage creates a DerivativeInfo from a quickfix.Message instance
func FromMessageToDerivativeInfo(m *quickfix.Message) DerivativeInfo {
	return DerivativeInfo{StockInfo: &StockInfo{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}}
}

// Route returns the beginstring, message type, and MessageRoute for this Message type
func RouteDerivativeInfo(router func(msg DerivativeInfo, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToDerivativeInfo(msg), sessionID)
	}
	return fix44.BeginString, "DI", r
}

// GetUnderlying Tag 800
// Mã tài sản cơ sở
func (m StockInfo) GetUnderlying() (v string, err quickfix.MessageRejectError) {
	var f UnderlyingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type UnderlyingField struct{ quickfix.FIXString }

func (f UnderlyingField) Tag() quickfix.Tag { return 800 }
func (f UnderlyingField) Value() string     { return f.String() }

// GetOpenInterest Tag 801
// Khối lượng mở OI, cuối ngày mới có giá trị
func (m StockInfo) GetOpenInterest() (v float64, err quickfix.MessageRejectError) {
	var f OpenInterestField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type OpenInterestField struct{ quickfix.FIXDecimal }

func (f OpenInterestField) Tag() quickfix.Tag      { return 801 }
func (f OpenInterestField) Value() decimal.Decimal { return f.Decimal }

// GetOpenInterestChange Tag 8011
// Thay đổi khối lượng mở OI (%), cuối ngày mới có giá trị
func (m StockInfo) GetOpenInterestChange() (v float64, err quickfix.MessageRejectError) {
	var f OpenInterestChangeField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type OpenInterestChangeField struct{ quickfix.FIXDecimal }

func (f OpenInterestChangeField) Tag() quickfix.Tag      { return 8011 }
func (f OpenInterestChangeField) Value() decimal.Decimal { return f.Decimal }

// GetFirstTradingDate Tag 802.
// Ngày giao dịch đầu tiên theo định dạng dd/MM/yyyy :v
func (m StockInfo) GetFirstTradingDate() (v string, err quickfix.MessageRejectError) {
	var f FirstTradingDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type FirstTradingDateField struct{ quickfix.FIXString }

func (f FirstTradingDateField) Tag() quickfix.Tag { return 802 }
func (f FirstTradingDateField) Value() string     { return f.String() }

// GetLastTradingDate Tag 803.
// Ngày giao dịch cuối cùng theo định dạng dd/MM/yyyy
func (m StockInfo) GetLastTradingDate() (v string, err quickfix.MessageRejectError) {
	var f LastTradingDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type LastTradingDateField struct{ quickfix.FIXString }

func (f LastTradingDateField) Tag() quickfix.Tag { return 803 }
func (f LastTradingDateField) Value() string     { return f.String() }

// GetTradingSessionID Tag 336.
//	case "AVAILABLE":
//	case "CALL_AUCTION_OPENING":
//	case "OPEN":
//	case "CALL_AUCTION_CLOSING":
//	case "CLOSED":
func (m StockInfo) GetTradingSessionID() (v string, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.String()
	}
	return
}

// GetTradSesStatus Tag 340.
func (m StockInfo) GetTradSesStatus() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesStatusField
	if err = m.Get(&f); err == nil {
		v = f.String()
	}
	return
}
