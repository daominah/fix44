package hnxinfogate

import (
	"github.com/quickfixgo/field"
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

// GetBoardCode Tag 425
// Mã bảng của chứng khoán:
//	-LIS_BRD_01,..: bảng niêm yết
//	-UPC_BRD_01,…: bảng upcom
func (m BoardInfo) GetBoardCode() (v string, err quickfix.MessageRejectError) {
	var f BoardCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBoardStatus Tag 426
// Trạng thái của bảng:
//	-A: Đang hoạt động
//	-C: Ngừng hoạt động
//  -P: Tạm thời dừng hoạt động
func (m BoardInfo) GetBoardStatus() (v string, err quickfix.MessageRejectError) {
	var f BoardStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type BoardStatusField struct{ quickfix.FIXString }

func (f BoardStatusField) Tag() quickfix.Tag { return 426 }
func (f BoardStatusField) Value() string     { return f.String() }

// GetTradingSessionID Tag 336
// Mã trạng thái giao dịch.
// Các mã sử dụng trong bảng LIST:
//	LIS_AUC_O_NML: Phiên mở cửa(chưa áp dụng)
//	LIS_AUC_O_NML_LOC: Phiên mở cửa BL(chưa áp dụng)
//	LIS_CON_NML: Phiên liên tục
// 	LIS_AUC_C_NML: Phiên đóng cửa
//	LIS_AUC_C_NML_LOC: Phiên đóng cửa BL
//	LIS_PTH_P_NML: Phiên sau đóng cửa
// Các mã sử dụng trong bảng UPCOM:
//	UPC_AUC_O_NML: Phiên mở cửa(chưa áp dụng):
//	UPC_AUC_O_NML_LOC: Phiên mở cửa BL(chưa áp dụng)
//	UPC_CON_NML: Phiên liên tục
//	UPC_AUC_C_NML: Phiên đóng cửa
//	UPC_AUC_C_NML_LOC: Phiên đóng cửa BL
//	UPC_PTH_P_NML: Phiên sau đóng cửa(chưa áp dụng)
func (m BoardInfo) GetTradingSessionID() (v string, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.String()
	}
	return
}

// GetTradSesStatus Tag 340
// Trạng thái giao dịch (áp dụng cho Cổ phiếu):
//	= 0 Chưa bắt đầu.
//	= 1 Bình thường
//	= 2 Tạm dừng
//	= 3 Kết thúc nhận lệnh phiên hiện tại do RandomEnd
//	= 4 Tạm dừng do CircuitBreak
//	= 5 Phiên định kỳ sau CB
//	= 6 Chứng khoán đang Prolong
//	= 13 Kết thúc nhận lệnh của ngày giao dịch hiện tại
//	= 90 Thị trường đang ở trạng thái chờ nhận lệnh
//	= 97 Đóng cửa thị trường
func (m BoardInfo) GetTradSesStatus() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesStatusField
	if err = m.Get(&f); err == nil {
		v = f.String()
	}
	return
}

// GetName Tag 421
func (m BoardInfo) GetName() (v string, err quickfix.MessageRejectError) {
	var f NameField
	if err = m.Get(&f); err == nil {
		v = f.String()
	}
	return
}

type NameField struct{ quickfix.FIXString }

func (f NameField) Tag() quickfix.Tag { return 421 }
func (f NameField) Value() string     { return f.String() }

// GetNumSymbolAdvances Tag 251
func (m BoardInfo) GetNumSymbolAdvances() (v int, err quickfix.MessageRejectError) {
	var f NumSymbolAdvancesField
	if err = m.Get(&f); err == nil {
		v = f.Int()
	}
	return
}

type NumSymbolAdvancesField struct{ quickfix.FIXInt }

func (f NumSymbolAdvancesField) Tag() quickfix.Tag { return 251 }
func (f NumSymbolAdvancesField) Value() int        { return f.Int() }

// GetNumSymbolNoChange Tag 252
func (m BoardInfo) GetNumSymbolNoChange() (v int, err quickfix.MessageRejectError) {
	var f NumSymbolNoChangeField
	if err = m.Get(&f); err == nil {
		v = f.Int()
	}
	return
}

type NumSymbolNoChangeField struct{ quickfix.FIXInt }

func (f NumSymbolNoChangeField) Tag() quickfix.Tag { return 252 }
func (f NumSymbolNoChangeField) Value() int        { return f.Int() }

// GetNumSymbolDeclines Tag 253
func (m BoardInfo) GetNumSymbolDeclines() (v int, err quickfix.MessageRejectError) {
	var f NumSymbolDeclinesField
	if err = m.Get(&f); err == nil {
		v = f.Int()
	}
	return
}

type NumSymbolDeclinesField struct{ quickfix.FIXInt }

func (f NumSymbolDeclinesField) Tag() quickfix.Tag { return 253 }
func (f NumSymbolDeclinesField) Value() int        { return f.Int() }

func (m BoardInfo) GetTime() (v string, err quickfix.MessageRejectError) {
	var f TimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}
