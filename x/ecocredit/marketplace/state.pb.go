// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/marketplace/v1/state.proto

package marketplace

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SellOrder represents the information for a sell order.
type SellOrder struct {
	// id is the unique ID of sell order.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// seller is the address of the account that is selling credits.
	Seller []byte `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	// batch_key is the table row identifier of the credit batch used internally
	// for efficient lookups. This links a sell order to a credit batch.
	BatchKey uint64 `protobuf:"varint,3,opt,name=batch_key,json=batchKey,proto3" json:"batch_key,omitempty"`
	// quantity is the decimal quantity of credits being sold.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// market_id is the market in which this sell order exists and specifies
	// the bank_denom that ask_amount corresponds to forming the ask_price.
	MarketId uint64 `protobuf:"varint,5,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// ask_amount is the integer amount (encoded as a string) that the seller is
	// asking for each credit unit of the batch. Each credit unit of the batch
	// will be sold for at least the ask_amount. The ask_amount corresponds to the
	// Market.denom to form the ask price.
	AskAmount string `protobuf:"bytes,6,opt,name=ask_amount,json=askAmount,proto3" json:"ask_amount,omitempty"`
	// disable_auto_retire disables auto-retirement of credits which allows a
	// buyer to disable auto-retirement in their buy order enabling them to
	// resell the credits to another buyer.
	DisableAutoRetire bool `protobuf:"varint,7,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// expiration is an optional timestamp when the sell order expires. When the
	// expiration time is reached, the sell order is removed from state.
	Expiration *types.Timestamp `protobuf:"bytes,9,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// maker indicates that this is a maker order, meaning that when it hit
	// the order book, there were no matching buy orders.
	Maker bool `protobuf:"varint,10,opt,name=maker,proto3" json:"maker,omitempty"`
}

func (m *SellOrder) Reset()         { *m = SellOrder{} }
func (m *SellOrder) String() string { return proto.CompactTextString(m) }
func (*SellOrder) ProtoMessage()    {}
func (*SellOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_718b9cb8f10a9f3c, []int{0}
}
func (m *SellOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SellOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SellOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SellOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellOrder.Merge(m, src)
}
func (m *SellOrder) XXX_Size() int {
	return m.Size()
}
func (m *SellOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SellOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SellOrder proto.InternalMessageInfo

func (m *SellOrder) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SellOrder) GetSeller() []byte {
	if m != nil {
		return m.Seller
	}
	return nil
}

func (m *SellOrder) GetBatchKey() uint64 {
	if m != nil {
		return m.BatchKey
	}
	return 0
}

func (m *SellOrder) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *SellOrder) GetMarketId() uint64 {
	if m != nil {
		return m.MarketId
	}
	return 0
}

func (m *SellOrder) GetAskAmount() string {
	if m != nil {
		return m.AskAmount
	}
	return ""
}

func (m *SellOrder) GetDisableAutoRetire() bool {
	if m != nil {
		return m.DisableAutoRetire
	}
	return false
}

func (m *SellOrder) GetExpiration() *types.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *SellOrder) GetMaker() bool {
	if m != nil {
		return m.Maker
	}
	return false
}

// AllowedDenom represents the information for an allowed ask/bid denom.
type AllowedDenom struct {
	// denom is the bank denom to allow (ex. ibc/GLKHDSG423SGS)
	BankDenom string `protobuf:"bytes,1,opt,name=bank_denom,json=bankDenom,proto3" json:"bank_denom,omitempty"`
	// display_denom is the denom to display to the user and is informational.
	// Because the denom is likely an IBC denom, this should be chosen by
	// governance to represent the consensus trusted name of the denom.
	DisplayDenom string `protobuf:"bytes,2,opt,name=display_denom,json=displayDenom,proto3" json:"display_denom,omitempty"`
	// exponent is the exponent that relates the denom to the display_denom and is
	// informational
	Exponent uint32 `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (m *AllowedDenom) Reset()         { *m = AllowedDenom{} }
func (m *AllowedDenom) String() string { return proto.CompactTextString(m) }
func (*AllowedDenom) ProtoMessage()    {}
func (*AllowedDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_718b9cb8f10a9f3c, []int{1}
}
func (m *AllowedDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedDenom.Merge(m, src)
}
func (m *AllowedDenom) XXX_Size() int {
	return m.Size()
}
func (m *AllowedDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedDenom.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedDenom proto.InternalMessageInfo

func (m *AllowedDenom) GetBankDenom() string {
	if m != nil {
		return m.BankDenom
	}
	return ""
}

func (m *AllowedDenom) GetDisplayDenom() string {
	if m != nil {
		return m.DisplayDenom
	}
	return ""
}

func (m *AllowedDenom) GetExponent() uint32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

// Market describes a distinctly processed market between a credit type and
// allowed bank denom. Each market has its own precision in the order book
// and is processed independently of other markets. Governance must enable
// markets one by one. Every additional enabled market potentially adds more
// processing overhead to the blockchain and potentially weakens liquidity in
// competing markets. For instance, enabling side by side USD/Carbon and
// EUR/Carbon markets may have the end result that each market individually has
// less liquidity and longer settlement times. Such decisions should be taken
// with care.
type Market struct {
	// id is the unique ID of the market.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// credit_type_abbrev is the abbreviation of the credit type.
	CreditTypeAbbrev string `protobuf:"bytes,2,opt,name=credit_type_abbrev,json=creditTypeAbbrev,proto3" json:"credit_type_abbrev,omitempty"`
	// bank_denom is an allowed bank denom.
	BankDenom string `protobuf:"bytes,3,opt,name=bank_denom,json=bankDenom,proto3" json:"bank_denom,omitempty"`
	// precision_modifier is an optional modifier used to convert arbitrary
	// precision integer bank amounts to uint32 values used for sorting in the
	// order book. Given an arbitrary precision integer x, its uint32 conversion
	// will be x / 10^precision_modifier using round half away from zero
	// rounding.
	//
	// uint32 values range from 0 to 4,294,967,295.
	// This allows for a full 9 digits of precision. In most real world markets
	// this amount of precision is sufficient and most common downside -
	// that some orders with very miniscule price differences may be ordered
	// equivalently (because of rounding) - is acceptable.
	// Note that this rounding will not affect settlement price which will
	// always be done exactly.
	//
	// Given a USD stable coin with 6 decimal digits, a precision_modifier
	// of 0 is probably acceptable as long as credits are always less than
	// $4,294/unit. With precision down to $0.001 (a precision_modifier of 3
	// in this case), prices can rise up to $4,294,000/unit. Either scenario
	// is probably quite acceptable given that carbon prices are unlikely to
	// rise above $1000/ton any time in the near future.
	//
	// If credit prices, exceed the maximum range of uint32 with this
	// precision_modifier, orders with high prices will fail and governance
	// will need to adjust precision_modifier to allow for higher prices in
	// exchange for less precision at the lower end.
	PrecisionModifier uint32 `protobuf:"varint,4,opt,name=precision_modifier,json=precisionModifier,proto3" json:"precision_modifier,omitempty"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_718b9cb8f10a9f3c, []int{2}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Market) GetCreditTypeAbbrev() string {
	if m != nil {
		return m.CreditTypeAbbrev
	}
	return ""
}

func (m *Market) GetBankDenom() string {
	if m != nil {
		return m.BankDenom
	}
	return ""
}

func (m *Market) GetPrecisionModifier() uint32 {
	if m != nil {
		return m.PrecisionModifier
	}
	return 0
}

func init() {
	proto.RegisterType((*SellOrder)(nil), "regen.ecocredit.marketplace.v1.SellOrder")
	proto.RegisterType((*AllowedDenom)(nil), "regen.ecocredit.marketplace.v1.AllowedDenom")
	proto.RegisterType((*Market)(nil), "regen.ecocredit.marketplace.v1.Market")
}

func init() {
	proto.RegisterFile("regen/ecocredit/marketplace/v1/state.proto", fileDescriptor_718b9cb8f10a9f3c)
}

var fileDescriptor_718b9cb8f10a9f3c = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xee, 0x24, 0x6d, 0xcc, 0x8e, 0x49, 0x49, 0xa7, 0x22, 0x43, 0xa4, 0x6b, 0x68, 0x11, 0x82,
	0xb6, 0xbb, 0x44, 0xe9, 0x25, 0xe2, 0x21, 0xe2, 0x45, 0xa4, 0x08, 0x6b, 0x41, 0xf0, 0xb2, 0xcc,
	0xee, 0xbc, 0xa6, 0xc3, 0xce, 0xee, 0xac, 0xb3, 0x93, 0xb4, 0xf9, 0x13, 0xe2, 0x5d, 0xf0, 0xf7,
	0x08, 0x5e, 0x0a, 0x5e, 0x3c, 0x4a, 0x7b, 0xf0, 0xee, 0x2f, 0x90, 0x9d, 0xdd, 0xc6, 0xd0, 0x7a,
	0xfc, 0xde, 0xf7, 0xbd, 0x99, 0xf7, 0x7d, 0xf3, 0x06, 0x3f, 0xd6, 0x30, 0x85, 0xcc, 0x87, 0x58,
	0xc5, 0x1a, 0xb8, 0x30, 0x7e, 0xca, 0x74, 0x02, 0x26, 0x97, 0x2c, 0x06, 0x7f, 0x3e, 0xf2, 0x0b,
	0xc3, 0x0c, 0x78, 0xb9, 0x56, 0x46, 0x11, 0xd7, 0x6a, 0xbd, 0xa5, 0xd6, 0x5b, 0xd1, 0x7a, 0xf3,
	0x51, 0x7f, 0x27, 0x56, 0x45, 0xaa, 0x0a, 0x5f, 0xe9, 0xd4, 0x9f, 0x8f, 0x98, 0xcc, 0x4f, 0xd9,
	0xa8, 0x04, 0x55, 0x7b, 0xff, 0xe1, 0x54, 0xa9, 0xa9, 0x04, 0xdf, 0xa2, 0x68, 0x76, 0xe2, 0x1b,
	0x91, 0x42, 0x61, 0x58, 0x9a, 0x57, 0x82, 0xdd, 0xdf, 0x0d, 0xec, 0xbc, 0x03, 0x29, 0xdf, 0x6a,
	0x0e, 0x9a, 0x6c, 0xe2, 0x86, 0xe0, 0x14, 0x0d, 0xd0, 0x70, 0x3d, 0x68, 0x08, 0x4e, 0xee, 0xe3,
	0x56, 0x01, 0x52, 0x82, 0xa6, 0x8d, 0x01, 0x1a, 0x76, 0x82, 0x1a, 0x91, 0x07, 0xd8, 0x89, 0x98,
	0x89, 0x4f, 0xc3, 0x04, 0x16, 0xb4, 0x69, 0xe5, 0x6d, 0x5b, 0x78, 0x03, 0x0b, 0xd2, 0xc7, 0xed,
	0x8f, 0x33, 0x96, 0x19, 0x61, 0x16, 0x74, 0x7d, 0x80, 0x86, 0x4e, 0xb0, 0xc4, 0x65, 0x63, 0x65,
	0x20, 0x14, 0x9c, 0x6e, 0x54, 0x8d, 0x55, 0xe1, 0x35, 0x27, 0x3b, 0x18, 0xb3, 0x22, 0x09, 0x59,
	0xaa, 0x66, 0x99, 0xa1, 0x2d, 0xdb, 0xea, 0xb0, 0x22, 0x99, 0xd8, 0x02, 0xf1, 0xf0, 0x36, 0x17,
	0x05, 0x8b, 0x24, 0x84, 0x6c, 0x66, 0x54, 0xa8, 0xc1, 0x08, 0x0d, 0xf4, 0xce, 0x00, 0x0d, 0xdb,
	0xc1, 0x56, 0x4d, 0x4d, 0x66, 0x46, 0x05, 0x96, 0x20, 0x63, 0x8c, 0xe1, 0x3c, 0x17, 0x9a, 0x19,
	0xa1, 0x32, 0xea, 0x0c, 0xd0, 0xf0, 0xee, 0xd3, 0xbe, 0x57, 0x05, 0xe2, 0x5d, 0x07, 0xe2, 0x1d,
	0x5f, 0x07, 0x12, 0xac, 0xa8, 0xc9, 0x3d, 0xbc, 0x91, 0xb2, 0x04, 0x34, 0xc5, 0xf6, 0xf4, 0x0a,
	0x8c, 0x9f, 0xff, 0xf9, 0xfa, 0xe3, 0x53, 0xf3, 0x10, 0xb7, 0xca, 0x98, 0x7a, 0x88, 0x74, 0x57,
	0x62, 0xe8, 0x21, 0x82, 0xaf, 0xd3, 0xea, 0x35, 0xc8, 0xe6, 0xea, 0xe5, 0xbd, 0x26, 0x45, 0xbb,
	0x5f, 0x10, 0xee, 0x4c, 0xa4, 0x54, 0x67, 0xc0, 0x5f, 0x41, 0xa6, 0xd2, 0xd2, 0x6e, 0xc4, 0xb2,
	0x24, 0xe4, 0x25, 0xb2, 0xa1, 0x3b, 0x81, 0x53, 0x56, 0x2a, 0x7a, 0x0f, 0x77, 0xb9, 0x28, 0x72,
	0xc9, 0x16, 0xb5, 0xa2, 0x61, 0x15, 0x9d, 0xba, 0x58, 0x89, 0xfa, 0xb8, 0x0d, 0xe7, 0xb9, 0xca,
	0x20, 0x33, 0xf6, 0x1d, 0xba, 0xc1, 0x12, 0x8f, 0x9f, 0xd8, 0x69, 0x1f, 0xe1, 0xce, 0xea, 0x3d,
	0x64, 0xfb, 0xc6, 0xb1, 0x3d, 0x44, 0x11, 0x6d, 0xee, 0x7e, 0x47, 0xb8, 0x75, 0x64, 0x1f, 0xe2,
	0xd6, 0x12, 0xec, 0x63, 0x52, 0xed, 0x5e, 0x68, 0x16, 0x39, 0x84, 0x2c, 0x8a, 0x34, 0xcc, 0xeb,
	0x69, 0x7a, 0x15, 0x73, 0xbc, 0xc8, 0x61, 0x62, 0xeb, 0x37, 0x5c, 0x35, 0x6f, 0xba, 0x3a, 0xc0,
	0x24, 0xd7, 0x10, 0x8b, 0x42, 0xa8, 0x2c, 0x4c, 0x15, 0x17, 0x27, 0x02, 0xb4, 0x5d, 0x93, 0x6e,
	0xb0, 0xb5, 0x64, 0x8e, 0x6a, 0x62, 0x7c, 0x68, 0x3d, 0xf8, 0xcb, 0xc4, 0xf7, 0xf0, 0xce, 0xed,
	0x59, 0xf6, 0xff, 0x5d, 0x68, 0xdd, 0xac, 0xbf, 0x7c, 0xff, 0xed, 0xd2, 0x45, 0x17, 0x97, 0x2e,
	0xfa, 0x75, 0xe9, 0xa2, 0xcf, 0x57, 0xee, 0xda, 0xc5, 0x95, 0xbb, 0xf6, 0xf3, 0xca, 0x5d, 0xfb,
	0xf0, 0x62, 0x2a, 0xcc, 0xe9, 0x2c, 0xf2, 0x62, 0x95, 0xfa, 0xf6, 0x6b, 0x1d, 0x64, 0x60, 0xce,
	0x94, 0x4e, 0x6a, 0x24, 0x81, 0x4f, 0x41, 0xfb, 0xe7, 0xff, 0xff, 0x9d, 0x51, 0xcb, 0xee, 0xcd,
	0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x77, 0x4d, 0xa3, 0xc3, 0x03, 0x00, 0x00,
}

func (m *SellOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SellOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SellOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Maker {
		i--
		if m.Maker {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.Expiration != nil {
		{
			size, err := m.Expiration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.DisableAutoRetire {
		i--
		if m.DisableAutoRetire {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.AskAmount) > 0 {
		i -= len(m.AskAmount)
		copy(dAtA[i:], m.AskAmount)
		i = encodeVarintState(dAtA, i, uint64(len(m.AskAmount)))
		i--
		dAtA[i] = 0x32
	}
	if m.MarketId != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Quantity) > 0 {
		i -= len(m.Quantity)
		copy(dAtA[i:], m.Quantity)
		i = encodeVarintState(dAtA, i, uint64(len(m.Quantity)))
		i--
		dAtA[i] = 0x22
	}
	if m.BatchKey != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.BatchKey))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintState(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AllowedDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exponent != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Exponent))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DisplayDenom) > 0 {
		i -= len(m.DisplayDenom)
		copy(dAtA[i:], m.DisplayDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.DisplayDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BankDenom) > 0 {
		i -= len(m.BankDenom)
		copy(dAtA[i:], m.BankDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.BankDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PrecisionModifier != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.PrecisionModifier))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BankDenom) > 0 {
		i -= len(m.BankDenom)
		copy(dAtA[i:], m.BankDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.BankDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CreditTypeAbbrev) > 0 {
		i -= len(m.CreditTypeAbbrev)
		copy(dAtA[i:], m.CreditTypeAbbrev)
		i = encodeVarintState(dAtA, i, uint64(len(m.CreditTypeAbbrev)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SellOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.BatchKey != 0 {
		n += 1 + sovState(uint64(m.BatchKey))
	}
	l = len(m.Quantity)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.MarketId != 0 {
		n += 1 + sovState(uint64(m.MarketId))
	}
	l = len(m.AskAmount)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.DisableAutoRetire {
		n += 2
	}
	if m.Expiration != nil {
		l = m.Expiration.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Maker {
		n += 2
	}
	return n
}

func (m *AllowedDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BankDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.DisplayDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Exponent != 0 {
		n += 1 + sovState(uint64(m.Exponent))
	}
	return n
}

func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.CreditTypeAbbrev)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.BankDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.PrecisionModifier != 0 {
		n += 1 + sovState(uint64(m.PrecisionModifier))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SellOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SellOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SellOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = append(m.Seller[:0], dAtA[iNdEx:postIndex]...)
			if m.Seller == nil {
				m.Seller = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchKey", wireType)
			}
			m.BatchKey = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchKey |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quantity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AskAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableAutoRetire = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = &types.Timestamp{}
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maker", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Maker = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AllowedDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllowedDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BankDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BankDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			m.Exponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exponent |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditTypeAbbrev", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditTypeAbbrev = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BankDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BankDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrecisionModifier", wireType)
			}
			m.PrecisionModifier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrecisionModifier |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)