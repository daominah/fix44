package hnxinfogate

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
)

// StockInfo MsgType = SI.
// Gửi thông tin chi tiết về chứng khoán mà hệ thống HNX tính toán
type StockInfo struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a StockInfo from a quickfix.Message instance
func FromMessageToStockInfo(m *quickfix.Message) StockInfo {
	return StockInfo{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m StockInfo) ToMessage() *quickfix.Message {
	return m.Message
}

//Route returns the beginstring, message type, and MessageRoute for this Message type
func RouteStockInfo(router func(msg StockInfo, sessionID quickfix.SessionID) quickfix.MessageRejectError) (
	string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessageToStockInfo(msg), sessionID)
	}
	return fix44.BeginString, "SI", r
}

// GetSymbol Tag 55
func (m StockInfo) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBoardCode Tag 425
// Mã bảng của chứng khoán:
//	-LIS_BRD_01,..: bảng niêm yết
//	-UPC_BRD_01,…: bảng upcom
func (m StockInfo) GetBoardCode() (v string, err quickfix.MessageRejectError) {
	var f BoardCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type BoardCodeField struct{ quickfix.FIXString }

func (f BoardCodeField) Tag() quickfix.Tag { return 425 }
func (f BoardCodeField) Value() string     { return f.String() }

// GetSecurityTradingStatus Tag 326.
// Trạng thái chứng khoán:
//	= 0: Bình thường
//	= 1: Chứng khoán không được giao dịch trong ngày
//	= 2: Ngừng giao dịch
//	= 6: Hủy niêm yết
//	= 7: Niêm yết mới
//	= 8: Sắp hủy niêm yếtGetTradingSessionID
//	= 10: Tạm ngừng giao dịch giữa phiên
//	= 25: Giao dịch đặc biệt
func (m StockInfo) GetSecurityTradingStatus() (v int, err quickfix.MessageRejectError) {
	var f SecurityTradingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type SecurityTradingStatusField struct{ quickfix.FIXInt }

func (f SecurityTradingStatusField) Tag() quickfix.Tag { return 326 }
func (f SecurityTradingStatusField) Value() int        { return f.Int() }

// GetSecurityType Tag 167.
// Loại chứng khoán:
//	ST: Cổ phiếu
//	BO: Trái phiếu
//	MF: Chứng chỉ quỹ
//	EF: Exchange-Traded Funds
//	FU: Future
//	OP: Option
func (m StockInfo) GetSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = string(f.Value())
	}
	return
}

// GetIssueDate Tag 225.
// Ngày phát hành theo định dạng yyyyMMdd
func (m StockInfo) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer Tag 106.
// Tổ chức phát hành
func (m StockInfo) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc Tag 107.
// Mô tả thêm về chứng khoán
func (m StockInfo) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBestBidPrice Tag 132.
// Giá đặt mua tốt nhất của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetBestBidPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.BidPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetBestBidQtty Tag 1321.
// Khối lượng đặt mua tốt nhất của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetBestBidQtty() (v float64, err quickfix.MessageRejectError) {
	var f field.DerivativeCapPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetBestOfferPrice Tag 133.
// Giá đặt bán tốt nhất của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetBestOfferPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.OfferPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetBestOfferQtty Tag 1331.
// Khối lượng đặt bán tốt nhất của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetBestOfferQtty() (v float64, err quickfix.MessageRejectError) {
	var f BestOfferQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type BestOfferQttyField struct{ quickfix.FIXDecimal }

func (f BestOfferQttyField) Tag() quickfix.Tag      { return 1331 }
func (f BestOfferQttyField) Value() decimal.Decimal { return f.Decimal }

// GetTotalBidQtty Tag 134.
// Tổng KL đặt mua của GD khớp lệnh lô chẵn (trừ kl sửa, hủy)
func (m StockInfo) GetTotalBidQtty() (v float64, err quickfix.MessageRejectError) {
	var f field.BidSizeField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetTotalOfferQtty Tag 135.
// Tổng KL đặt bán của GD khớp lệnh lô chẵn (trừ kl sửa, hủy)
func (m StockInfo) GetTotalOfferQtty() (v float64, err quickfix.MessageRejectError) {
	var f field.OfferSizeField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetBasicPrice Tag 260.
// Giá tham chiếu (nghiệp vụ)
func (m StockInfo) GetBasicPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.BasisFeaturePriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetFloorPrice Tag 333.
// Giá sàn (nghiệp vụ)
func (m StockInfo) GetFloorPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.LowPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetCeilingPrice Tag 332.
// Giá trần (nghiệp vụ)
func (m StockInfo) GetCeilingPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.HighPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetFloorPricePT Tag 3331.
// Giá sàn cho giao dịch thỏa thuận ngoài biên độ (nghiệp vụ)
func (m StockInfo) GetFloorPricePT() (v float64, err quickfix.MessageRejectError) {
	var f FloorPricePTField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type FloorPricePTField struct{ quickfix.FIXDecimal }

func (f FloorPricePTField) Tag() quickfix.Tag      { return 3331 }
func (f FloorPricePTField) Value() decimal.Decimal { return f.Decimal }

// GetCeilingPricePT Tag 3321.
// Giá trần cho giao dịch thỏa thuận ngoài biên độ (nghiệp vụ)
func (m StockInfo) GetCeilingPricePT() (v float64, err quickfix.MessageRejectError) {
	var f CeilingPricePTField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type CeilingPricePTField struct{ quickfix.FIXDecimal }

func (f CeilingPricePTField) Tag() quickfix.Tag      { return 3321 }
func (f CeilingPricePTField) Value() decimal.Decimal { return f.Decimal }

// GetParValue Tag 334.
// Mệnh giá chứng khoán
func (m StockInfo) GetParValue() (v float64, err quickfix.MessageRejectError) {
	var f ParValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type ParValueField struct{ quickfix.FIXDecimal }

func (f ParValueField) Tag() quickfix.Tag      { return 334 }
func (f ParValueField) Value() decimal.Decimal { return f.Decimal }

// GetMatchPrice Tag 31.
// Giá khớp gần nhất của GD khớp lệnh lô chẵn
func (m StockInfo) GetMatchPrice() (v float64, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetMatchQtty Tag 32.
// KL khớp gần của GD khớp lệnh lô chăn
func (m StockInfo) GetMatchQtty() (v float64, err quickfix.MessageRejectError) {
	var f field.LastQtyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetOpenPrice Tag 137.
// Giá mở cửa (nghiệp vụ)
func (m StockInfo) GetOpenPrice() (v float64, err quickfix.MessageRejectError) {
	var f OpenPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type OpenPriceField struct{ quickfix.FIXDecimal }

func (f OpenPriceField) Tag() quickfix.Tag      { return 137 }
func (f OpenPriceField) Value() decimal.Decimal { return f.Decimal }

// GetPriorOpenPrice Tag 138.
// Giá mở cửa phiên giao dịch trước phiên giao dịch hiện tại
func (m StockInfo) GetPriorOpenPrice() (v float64, err quickfix.MessageRejectError) {
	var f PriorOpenPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PriorOpenPriceField struct{ quickfix.FIXDecimal }

func (f PriorOpenPriceField) Tag() quickfix.Tag      { return 138 }
func (f PriorOpenPriceField) Value() decimal.Decimal { return f.Decimal }

// GetClosePrice Tag 139.
// Giá đóng cửa (nghiệp vụ)
func (m StockInfo) GetClosePrice() (v float64, err quickfix.MessageRejectError) {
	var f ClosePriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type ClosePriceField struct{ quickfix.FIXDecimal }

func (f ClosePriceField) Tag() quickfix.Tag      { return 139 }
func (f ClosePriceField) Value() decimal.Decimal { return f.Decimal }

// GetPriorClosePrice Tag 140.
// Giá đóng cửa phiên trước phiên giao dịch hiện tại
func (m StockInfo) GetPriorClosePrice() (v float64, err quickfix.MessageRejectError) {
	var f PriorClosePriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PriorClosePriceField struct{ quickfix.FIXDecimal }

func (f PriorClosePriceField) Tag() quickfix.Tag      { return 140 }
func (f PriorClosePriceField) Value() decimal.Decimal { return f.Decimal }

// GetTotalVolumeTraded Tag 387.
// Tổng KL giao dịch của GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetTotalVolumeTraded() (v float64, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetTotalValueTraded Tag 3871.
// Tổng giá trị giao dịch của GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetTotalValueTraded() (v float64, err quickfix.MessageRejectError) {
	var f TotalValueTradedField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalValueTradedField struct{ quickfix.FIXDecimal }

func (f TotalValueTradedField) Tag() quickfix.Tag      { return 3871 }
func (f TotalValueTradedField) Value() decimal.Decimal { return f.Decimal }

// GetMidPx Tag 631.
// Giá bình quân (nghiệp vụ)
func (m StockInfo) GetMidPx() (v float64, err quickfix.MessageRejectError) {
	var f field.MidPxField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// GetTradingDate Tag 388.
// Ngày giao dịch hiện tại theo định dạng yyyyMMdd
func (m StockInfo) GetTradingDate() (v string, err quickfix.MessageRejectError) {
	var f TradingDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type TradingDateField struct{ quickfix.FIXString }

func (f TradingDateField) Tag() quickfix.Tag { return 388 }
func (f TradingDateField) Value() string     { return f.String() }

// GetTime Tag 399.
// Thời gian theo định dạng HH:mm:ss
func (m StockInfo) GetTime() (v string, err quickfix.MessageRejectError) {
	var f TimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type TimeField struct{ quickfix.FIXString }

func (f TimeField) Tag() quickfix.Tag { return 399 }
func (f TimeField) Value() string     { return f.String() }

// GetTradingUnit Tag 400.
// Đơn vị giao dịch nhỏ nhất
func (m StockInfo) GetTradingUnit() (v float64, err quickfix.MessageRejectError) {
	var f TradingUnit
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TradingUnit struct{ quickfix.FIXDecimal }

func (f TradingUnit) Tag() quickfix.Tag      { return 400 }
func (f TradingUnit) Value() decimal.Decimal { return f.Decimal }

// GetTotalListingQtty Tag 109.
// Khối lượng niêm yết
func (m StockInfo) GetTotalListingQtty() (v float64, err quickfix.MessageRejectError) {
	var f TotalListingQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalListingQttyField struct{ quickfix.FIXDecimal }

func (f TotalListingQttyField) Tag() quickfix.Tag      { return 109 }
func (f TotalListingQttyField) Value() decimal.Decimal { return f.Decimal }

// GetDateNo Tag 17.
// Phiên giao dịch thứ ( kể từ ngày niêm yết)
func (m StockInfo) GetDateNo() (v float64, err quickfix.MessageRejectError) {
	var f DateNoField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type DateNoField struct{ quickfix.FIXDecimal }

func (f DateNoField) Tag() quickfix.Tag      { return 17 }
func (f DateNoField) Value() decimal.Decimal { return f.Decimal }

// GetAdjustQtty Tag .
// Dự phòng, không dùng
func (m StockInfo) GetAdjustQtty() (v float64, err quickfix.MessageRejectError) {
	var f AdjustQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type AdjustQttyField struct{ quickfix.FIXDecimal }

func (f AdjustQttyField) Tag() quickfix.Tag      { return 230 }
func (f AdjustQttyField) Value() decimal.Decimal { return f.Decimal }

// GetReferenceStatus Tag 232.
// Trạng thái thực hiện quyền ảnh hưởng tới giá chứng khoán:
//	0: Không xảy ra
//	1:Trả CT bằng tiền
//	2:Trả cổ tức bằng CP/CP thưởng
//	3: Phát hành CP cho cổ đông hiện hữu
//	4: Trả cổ tức bằng CP/CP thưởng,phát hành CP cho cổ đông hiện hữu
//	5: Trả cổ tức bằng tiền, bằng CP/CP thưởng, phát hành CP cho cổ đông hiện hữu
//	6: Niêm yết bổ sung
//	7: Giảm vốn
//	8: Trả cổ tức bằng tiền, trả cổ tức bằng CP/CP thưởng
//	9: Trả cổ tức bằng tiền, phát hành CP cho cổ đông hiện hữu
//	10: Thay đổi tỷ lệ Free Float
//	11: Họp đại cổ đông
func (m StockInfo) GetReferenceStatus() (v string, err quickfix.MessageRejectError) {
	var f ReferenceStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type ReferenceStatusField struct{ quickfix.FIXString }

func (f ReferenceStatusField) Tag() quickfix.Tag { return 232 }
func (f ReferenceStatusField) Value() string     { return f.String() }

// GetCurrentPrice Tag 255.
// Giá khớp dự kiến của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetCurrentPrice() (v float64, err quickfix.MessageRejectError) {
	var f CurrentPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type CurrentPriceField struct{ quickfix.FIXDecimal }

func (f CurrentPriceField) Tag() quickfix.Tag      { return 255 }
func (f CurrentPriceField) Value() decimal.Decimal { return f.Decimal }

// GetCurrentQtty Tag 2551.
// Khối lượng khớp dự kiến của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetCurrentQtty() (v float64, err quickfix.MessageRejectError) {
	var f CurrentQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type CurrentQttyField struct{ quickfix.FIXDecimal }

func (f CurrentQttyField) Tag() quickfix.Tag      { return 2551 }
func (f CurrentQttyField) Value() decimal.Decimal { return f.Decimal }

// GetHighestPrice Tag 266.
// Giá thực hiện cao nhất của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetHighestPrice() (v float64, err quickfix.MessageRejectError) {
	var f HighestPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type HighestPriceField struct{ quickfix.FIXDecimal }

func (f HighestPriceField) Tag() quickfix.Tag      { return 266 }
func (f HighestPriceField) Value() decimal.Decimal { return f.Decimal }

// GetLowestPrice Tag 2661.
// Giá thực hiện thấp nhất của GD khớp lệnh (lô chẵn)
func (m StockInfo) GetLowestPrice() (v float64, err quickfix.MessageRejectError) {
	var f LowestPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type LowestPriceField struct{ quickfix.FIXDecimal }

func (f LowestPriceField) Tag() quickfix.Tag      { return 2661 }
func (f LowestPriceField) Value() decimal.Decimal { return f.Decimal }

// GetPriorPrice Tag 277.
// Giá khớp lệnh của phiên trước đó. Chỉ tính với khớp lệnh thông thường.
func (m StockInfo) GetPriorPrice() (v float64, err quickfix.MessageRejectError) {
	var f PriorPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PriorPriceField struct{ quickfix.FIXDecimal }

func (f PriorPriceField) Tag() quickfix.Tag      { return 277 }
func (f PriorPriceField) Value() decimal.Decimal { return f.Decimal }

// GetMatchValue Tag 310.
// Giá trị khớp lệnh gần nhất của GD khớp lệnh lô chẵn
func (m StockInfo) GetMatchValue() (v float64, err quickfix.MessageRejectError) {
	var f MatchValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type MatchValueField struct{ quickfix.FIXDecimal }

func (f MatchValueField) Tag() quickfix.Tag      { return 310 }
func (f MatchValueField) Value() decimal.Decimal { return f.Decimal }

// GetOfferCount Tag 320.
// Tổng số lệnh đặt bán của GD khớp lệnh lô chẵn
func (m StockInfo) GetOfferCount() (v float64, err quickfix.MessageRejectError) {
	var f OfferCountField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type OfferCountField struct{ quickfix.FIXDecimal }

func (f OfferCountField) Tag() quickfix.Tag      { return 320 }
func (f OfferCountField) Value() decimal.Decimal { return f.Decimal }

// GetBidCount Tag 321.
// Tổng số lệnh đặt mua của GD khớp lệnh lô chẵn
func (m StockInfo) GetBidCount() (v float64, err quickfix.MessageRejectError) {
	var f BidCountField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type BidCountField struct{ quickfix.FIXDecimal }

func (f BidCountField) Tag() quickfix.Tag      { return 321 }
func (f BidCountField) Value() decimal.Decimal { return f.Decimal }

// GetNormalTotalTradedQtty Tag 391.
// Tổng khối lượng giao dịch thông thường của GD khớp lệnh lô chẵn
func (m StockInfo) GetNormalTotalTradedQtty() (v float64, err quickfix.MessageRejectError) {
	var f NormalTotalTradedQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type NormalTotalTradedQttyField struct{ quickfix.FIXDecimal }

func (f NormalTotalTradedQttyField) Tag() quickfix.Tag      { return 391 }
func (f NormalTotalTradedQttyField) Value() decimal.Decimal { return f.Decimal }

// GetNormalTotalTradedValue Tag 392.
//
func (m StockInfo) GetNormalTotalTradedValue() (v float64, err quickfix.MessageRejectError) {
	var f NormalTotalTradedValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type NormalTotalTradedValueField struct{ quickfix.FIXDecimal }

func (f NormalTotalTradedValueField) Tag() quickfix.Tag      { return 392 }
func (f NormalTotalTradedValueField) Value() decimal.Decimal { return f.Decimal }

// GetPutThroughMatchQtty Tag 393.
// Khối lượng thực hiện gần nhất của giao dịch thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetPutThroughMatchQtty() (v float64, err quickfix.MessageRejectError) {
	var f PutThroughMatchQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PutThroughMatchQttyField struct{ quickfix.FIXDecimal }

func (f PutThroughMatchQttyField) Tag() quickfix.Tag      { return 393 }
func (f PutThroughMatchQttyField) Value() decimal.Decimal { return f.Decimal }

// GetPutThroughMatchPrice Tag 3931.
// Giá thực hiện gần nhất của giao dịch thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetPutThroughMatchPrice() (v float64, err quickfix.MessageRejectError) {
	var f PutThroughMatchPriceField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PutThroughMatchPriceField struct{ quickfix.FIXDecimal }

func (f PutThroughMatchPriceField) Tag() quickfix.Tag      { return 3931 }
func (f PutThroughMatchPriceField) Value() decimal.Decimal { return f.Decimal }

// GetPutThroughTotalTradedQtty Tag 394.
// Tổng khối lượng của giao dịch thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetPutThroughTotalTradedQtty() (v float64, err quickfix.MessageRejectError) {
	var f PutThroughTotalTradedQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PutThroughTotalTradedQttyField struct{ quickfix.FIXDecimal }

func (f PutThroughTotalTradedQttyField) Tag() quickfix.Tag      { return 394 }
func (f PutThroughTotalTradedQttyField) Value() decimal.Decimal { return f.Decimal }

// GetPutThroughTotalTradedValue Tag 3941.
// Tổng giá trị của giao dịch thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetPutThroughTotalTradedValue() (v float64, err quickfix.MessageRejectError) {
	var f PutThroughTotalTradedValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type PutThroughTotalTradedValueField struct{ quickfix.FIXDecimal }

func (f PutThroughTotalTradedValueField) Tag() quickfix.Tag      { return 3941 }
func (f PutThroughTotalTradedValueField) Value() decimal.Decimal { return f.Decimal }

// GetTotalBuyTradingQtty Tag 395.
// Tổng khối lượng mua khớp của GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetTotalBuyTradingQtty() (v float64, err quickfix.MessageRejectError) {
	var f TotalBuyTradingQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalBuyTradingQttyField struct{ quickfix.FIXDecimal }

func (f TotalBuyTradingQttyField) Tag() quickfix.Tag      { return 395 }
func (f TotalBuyTradingQttyField) Value() decimal.Decimal { return f.Decimal }

// GetBuyCount Tag 3951.
//Tổng số lệnh mua khớp của GD khớp lệnh và thỏa thuân (lô chẵn và lẻ)
func (m StockInfo) GetBuyCount() (v float64, err quickfix.MessageRejectError) {
	var f BuyCountField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type BuyCountField struct{ quickfix.FIXDecimal }

func (f BuyCountField) Tag() quickfix.Tag      { return 3951 }
func (f BuyCountField) Value() decimal.Decimal { return f.Decimal }

// GetTotalBuyTradingValue Tag 3952.
// Tổng giá trị mua khớp của GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetTotalBuyTradingValue() (v float64, err quickfix.MessageRejectError) {
	var f TotalBuyTradingValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalBuyTradingValueField struct{ quickfix.FIXDecimal }

func (f TotalBuyTradingValueField) Tag() quickfix.Tag      { return 3952 }
func (f TotalBuyTradingValueField) Value() decimal.Decimal { return f.Decimal }

// GetTotalSellTradingQtty Tag 396.
// Tổng khối lượng bán khớp của GD khớp lệnh và thỏa thuân (lô chẵn và lẻ)
func (m StockInfo) GetTotalSellTradingQtty() (v float64, err quickfix.MessageRejectError) {
	var f TotalSellTradingQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalSellTradingQttyField struct{ quickfix.FIXDecimal }

func (f TotalSellTradingQttyField) Tag() quickfix.Tag      { return 396 }
func (f TotalSellTradingQttyField) Value() decimal.Decimal { return f.Decimal }

// GetSellCount Tag 3961.
// Tổng số lệnh bán khớp của GD khớp lệnh và thỏa thuân (lô chẵn và lẻ)
func (m StockInfo) GetSellCount() (v float64, err quickfix.MessageRejectError) {
	var f SellCountField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type SellCountField struct{ quickfix.FIXDecimal }

func (f SellCountField) Tag() quickfix.Tag      { return 3961 }
func (f SellCountField) Value() decimal.Decimal { return f.Decimal }

// GetTotalSellTradingValue Tag 3962.
// Tổng giá trị bán khớp của GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetTotalSellTradingValue() (v float64, err quickfix.MessageRejectError) {
	var f TotalSellTradingValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalSellTradingValueField struct{ quickfix.FIXDecimal }

func (f TotalSellTradingValueField) Tag() quickfix.Tag      { return 3962 }
func (f TotalSellTradingValueField) Value() decimal.Decimal { return f.Decimal }

// GetBuyForeignQtty Tag 397.
// Tổng khối lượng mua khớp của NĐT NN. Áp dụng cho GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetBuyForeignQtty() (v float64, err quickfix.MessageRejectError) {
	var f BuyForeignQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type BuyForeignQttyField struct{ quickfix.FIXDecimal }

func (f BuyForeignQttyField) Tag() quickfix.Tag      { return 397 }
func (f BuyForeignQttyField) Value() decimal.Decimal { return f.Decimal }

// GetBuyForeignValue Tag 3971.
// Tổng giá trị mua khớp của NĐTNN. Áp dụng cho GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetBuyForeignValue() (v float64, err quickfix.MessageRejectError) {
	var f BuyForeignValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type BuyForeignValueField struct{ quickfix.FIXDecimal }

func (f BuyForeignValueField) Tag() quickfix.Tag      { return 3971 }
func (f BuyForeignValueField) Value() decimal.Decimal { return f.Decimal }

// Tag 398.
// Tổng khối lượng bán khớp của NĐT NN. Áp dụng cho GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetSellForeignQtty() (v float64, err quickfix.MessageRejectError) {
	var f SellForeignQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type SellForeignQttyField struct{ quickfix.FIXDecimal }

func (f SellForeignQttyField) Tag() quickfix.Tag      { return 398 }
func (f SellForeignQttyField) Value() decimal.Decimal { return f.Decimal }

// Tag 3981.
// Tổng giá trị bán khớp của NĐT NN. Áp dụng cho GD khớp lệnh và thỏa thuận (lô chẵn và lẻ)
func (m StockInfo) GetSellForeignValue() (v float64, err quickfix.MessageRejectError) {
	var f SellForeignValueField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type SellForeignValueField struct{ quickfix.FIXDecimal }

func (f SellForeignValueField) Tag() quickfix.Tag      { return 3981 }
func (f SellForeignValueField) Value() decimal.Decimal { return f.Decimal }

// Tag 3301.
// Số lượng còn lại cho phép NDTNN đặt lệnh mua
func (m StockInfo) GetRemainForeignQtty() (v float64, err quickfix.MessageRejectError) {
	var f RemainForeignQttyField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type RemainForeignQttyField struct{ quickfix.FIXDecimal }

func (f RemainForeignQttyField) Tag() quickfix.Tag      { return 3301 }
func (f RemainForeignQttyField) Value() decimal.Decimal { return f.Decimal }

// Tag 541.
// Dự phòng, không dùng
func (m StockInfo) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.String()
	}
	return
}

// Tag 223.
// Dự phòng, không dùng
func (m StockInfo) GetCouponRate() (v float64, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

// Tag 1341.
// Tổng KL đặt mua của GD khớp lệnh lô lẻ (trừ sửa, hủy)
func (m StockInfo) GetTotalBidQttyOdd() (v float64, err quickfix.MessageRejectError) {
	var f TotalBidQttyOddField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalBidQttyOddField struct{ quickfix.FIXDecimal }

func (f TotalBidQttyOddField) Tag() quickfix.Tag      { return 1341 }
func (f TotalBidQttyOddField) Value() decimal.Decimal { return f.Decimal }

// Tag 1351.
// Tổng KL đặt bán của GD khớp lệnh lô lẻ (trừ sửa hủy)
func (m StockInfo) GetTotalOfferQttyOdd() (v float64, err quickfix.MessageRejectError) {
	var f TotalOfferQttyOddField
	if err = m.Get(&f); err == nil {
		v, _ = f.Value().Float64()
	}
	return
}

type TotalOfferQttyOddField struct{ quickfix.FIXDecimal }

func (f TotalOfferQttyOddField) Tag() quickfix.Tag      { return 1351 }
func (f TotalOfferQttyOddField) Value() decimal.Decimal { return f.Decimal }
