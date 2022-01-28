package orderbookv1beta1

import (
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	v1beta1 "github.com/regen-ledger/regen-network/api/regen/ecocredit/marketplace/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_BuyOrderSellOrderMatch                      protoreflect.MessageDescriptor
	fd_BuyOrderSellOrderMatch_bid_denom_id         protoreflect.FieldDescriptor
	fd_BuyOrderSellOrderMatch_buy_order_id         protoreflect.FieldDescriptor
	fd_BuyOrderSellOrderMatch_sell_order_id        protoreflect.FieldDescriptor
	fd_BuyOrderSellOrderMatch_bid_price_complement protoreflect.FieldDescriptor
	fd_BuyOrderSellOrderMatch_ask_price            protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_init()
	md_BuyOrderSellOrderMatch = File_regen_ecocredit_orderbook_v1beta1_memory_proto.Messages().ByName("BuyOrderSellOrderMatch")
	fd_BuyOrderSellOrderMatch_bid_denom_id = md_BuyOrderSellOrderMatch.Fields().ByName("bid_denom_id")
	fd_BuyOrderSellOrderMatch_buy_order_id = md_BuyOrderSellOrderMatch.Fields().ByName("buy_order_id")
	fd_BuyOrderSellOrderMatch_sell_order_id = md_BuyOrderSellOrderMatch.Fields().ByName("sell_order_id")
	fd_BuyOrderSellOrderMatch_bid_price_complement = md_BuyOrderSellOrderMatch.Fields().ByName("bid_price_complement")
	fd_BuyOrderSellOrderMatch_ask_price = md_BuyOrderSellOrderMatch.Fields().ByName("ask_price")
}

var _ protoreflect.Message = (*fastReflection_BuyOrderSellOrderMatch)(nil)

type fastReflection_BuyOrderSellOrderMatch BuyOrderSellOrderMatch

func (x *BuyOrderSellOrderMatch) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrderSellOrderMatch)(x)
}

func (x *BuyOrderSellOrderMatch) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrderSellOrderMatch_messageType fastReflection_BuyOrderSellOrderMatch_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrderSellOrderMatch_messageType{}

type fastReflection_BuyOrderSellOrderMatch_messageType struct{}

func (x fastReflection_BuyOrderSellOrderMatch_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrderSellOrderMatch)(nil)
}
func (x fastReflection_BuyOrderSellOrderMatch_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrderSellOrderMatch)
}
func (x fastReflection_BuyOrderSellOrderMatch_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrderSellOrderMatch
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrderSellOrderMatch) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrderSellOrderMatch
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrderSellOrderMatch) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrderSellOrderMatch_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrderSellOrderMatch) New() protoreflect.Message {
	return new(fastReflection_BuyOrderSellOrderMatch)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrderSellOrderMatch) Interface() protoreflect.ProtoMessage {
	return (*BuyOrderSellOrderMatch)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrderSellOrderMatch) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BidDenomId != uint32(0) {
		value := protoreflect.ValueOfUint32(x.BidDenomId)
		if !f(fd_BuyOrderSellOrderMatch_bid_denom_id, value) {
			return
		}
	}
	if x.BuyOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BuyOrderId)
		if !f(fd_BuyOrderSellOrderMatch_buy_order_id, value) {
			return
		}
	}
	if x.SellOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SellOrderId)
		if !f(fd_BuyOrderSellOrderMatch_sell_order_id, value) {
			return
		}
	}
	if x.BidPriceComplement != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BidPriceComplement)
		if !f(fd_BuyOrderSellOrderMatch_bid_price_complement, value) {
			return
		}
	}
	if x.AskPrice != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AskPrice)
		if !f(fd_BuyOrderSellOrderMatch_ask_price, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BuyOrderSellOrderMatch) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_denom_id":
		return x.BidDenomId != uint32(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.buy_order_id":
		return x.BuyOrderId != uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.sell_order_id":
		return x.SellOrderId != uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_price_complement":
		return x.BidPriceComplement != uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.ask_price":
		return x.AskPrice != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrderSellOrderMatch) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_denom_id":
		x.BidDenomId = uint32(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.buy_order_id":
		x.BuyOrderId = uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.sell_order_id":
		x.SellOrderId = uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_price_complement":
		x.BidPriceComplement = uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.ask_price":
		x.AskPrice = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrderSellOrderMatch) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_denom_id":
		value := x.BidDenomId
		return protoreflect.ValueOfUint32(value)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.buy_order_id":
		value := x.BuyOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.sell_order_id":
		value := x.SellOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_price_complement":
		value := x.BidPriceComplement
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.ask_price":
		value := x.AskPrice
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrderSellOrderMatch) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_denom_id":
		x.BidDenomId = uint32(value.Uint())
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.buy_order_id":
		x.BuyOrderId = value.Uint()
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.sell_order_id":
		x.SellOrderId = value.Uint()
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_price_complement":
		x.BidPriceComplement = value.Uint()
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.ask_price":
		x.AskPrice = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrderSellOrderMatch) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_denom_id":
		panic(fmt.Errorf("field bid_denom_id of message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch is not mutable"))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.buy_order_id":
		panic(fmt.Errorf("field buy_order_id of message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch is not mutable"))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch is not mutable"))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_price_complement":
		panic(fmt.Errorf("field bid_price_complement of message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch is not mutable"))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.ask_price":
		panic(fmt.Errorf("field ask_price of message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrderSellOrderMatch) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_denom_id":
		return protoreflect.ValueOfUint32(uint32(0))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.buy_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.bid_price_complement":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch.ask_price":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrderSellOrderMatch) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrderSellOrderMatch) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrderSellOrderMatch) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BuyOrderSellOrderMatch) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrderSellOrderMatch) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrderSellOrderMatch)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BidDenomId != 0 {
			n += 1 + runtime.Sov(uint64(x.BidDenomId))
		}
		if x.BuyOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.BuyOrderId))
		}
		if x.SellOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
		}
		if x.BidPriceComplement != 0 {
			n += 9
		}
		if x.AskPrice != 0 {
			n += 9
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BuyOrderSellOrderMatch)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.AskPrice != 0 {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.AskPrice))
			i--
			dAtA[i] = 0x29
		}
		if x.BidPriceComplement != 0 {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.BidPriceComplement))
			i--
			dAtA[i] = 0x21
		}
		if x.SellOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SellOrderId))
			i--
			dAtA[i] = 0x18
		}
		if x.BuyOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BuyOrderId))
			i--
			dAtA[i] = 0x10
		}
		if x.BidDenomId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BidDenomId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BuyOrderSellOrderMatch)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrderSellOrderMatch: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrderSellOrderMatch: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BidDenomId", wireType)
				}
				x.BidDenomId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BidDenomId |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuyOrderId", wireType)
				}
				x.BuyOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BuyOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderId", wireType)
				}
				x.SellOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SellOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BidPriceComplement", wireType)
				}
				x.BidPriceComplement = 0
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BidPriceComplement = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
			case 5:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AskPrice", wireType)
				}
				x.AskPrice = 0
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AskPrice = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_UInt64SelectorBuyOrder                  protoreflect.MessageDescriptor
	fd_UInt64SelectorBuyOrder_buy_order_id     protoreflect.FieldDescriptor
	fd_UInt64SelectorBuyOrder_selector_type    protoreflect.FieldDescriptor
	fd_UInt64SelectorBuyOrder_value            protoreflect.FieldDescriptor
	fd_UInt64SelectorBuyOrder_project_location protoreflect.FieldDescriptor
	fd_UInt64SelectorBuyOrder_min_start_date   protoreflect.FieldDescriptor
	fd_UInt64SelectorBuyOrder_max_end_date     protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_init()
	md_UInt64SelectorBuyOrder = File_regen_ecocredit_orderbook_v1beta1_memory_proto.Messages().ByName("UInt64SelectorBuyOrder")
	fd_UInt64SelectorBuyOrder_buy_order_id = md_UInt64SelectorBuyOrder.Fields().ByName("buy_order_id")
	fd_UInt64SelectorBuyOrder_selector_type = md_UInt64SelectorBuyOrder.Fields().ByName("selector_type")
	fd_UInt64SelectorBuyOrder_value = md_UInt64SelectorBuyOrder.Fields().ByName("value")
	fd_UInt64SelectorBuyOrder_project_location = md_UInt64SelectorBuyOrder.Fields().ByName("project_location")
	fd_UInt64SelectorBuyOrder_min_start_date = md_UInt64SelectorBuyOrder.Fields().ByName("min_start_date")
	fd_UInt64SelectorBuyOrder_max_end_date = md_UInt64SelectorBuyOrder.Fields().ByName("max_end_date")
}

var _ protoreflect.Message = (*fastReflection_UInt64SelectorBuyOrder)(nil)

type fastReflection_UInt64SelectorBuyOrder UInt64SelectorBuyOrder

func (x *UInt64SelectorBuyOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_UInt64SelectorBuyOrder)(x)
}

func (x *UInt64SelectorBuyOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_UInt64SelectorBuyOrder_messageType fastReflection_UInt64SelectorBuyOrder_messageType
var _ protoreflect.MessageType = fastReflection_UInt64SelectorBuyOrder_messageType{}

type fastReflection_UInt64SelectorBuyOrder_messageType struct{}

func (x fastReflection_UInt64SelectorBuyOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_UInt64SelectorBuyOrder)(nil)
}
func (x fastReflection_UInt64SelectorBuyOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_UInt64SelectorBuyOrder)
}
func (x fastReflection_UInt64SelectorBuyOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_UInt64SelectorBuyOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_UInt64SelectorBuyOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_UInt64SelectorBuyOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_UInt64SelectorBuyOrder) Type() protoreflect.MessageType {
	return _fastReflection_UInt64SelectorBuyOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_UInt64SelectorBuyOrder) New() protoreflect.Message {
	return new(fastReflection_UInt64SelectorBuyOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_UInt64SelectorBuyOrder) Interface() protoreflect.ProtoMessage {
	return (*UInt64SelectorBuyOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_UInt64SelectorBuyOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BuyOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BuyOrderId)
		if !f(fd_UInt64SelectorBuyOrder_buy_order_id, value) {
			return
		}
	}
	if x.SelectorType != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.SelectorType))
		if !f(fd_UInt64SelectorBuyOrder_selector_type, value) {
			return
		}
	}
	if x.Value != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Value)
		if !f(fd_UInt64SelectorBuyOrder_value, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_UInt64SelectorBuyOrder_project_location, value) {
			return
		}
	}
	if x.MinStartDate != nil {
		value := protoreflect.ValueOfMessage(x.MinStartDate.ProtoReflect())
		if !f(fd_UInt64SelectorBuyOrder_min_start_date, value) {
			return
		}
	}
	if x.MaxEndDate != nil {
		value := protoreflect.ValueOfMessage(x.MaxEndDate.ProtoReflect())
		if !f(fd_UInt64SelectorBuyOrder_max_end_date, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_UInt64SelectorBuyOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.buy_order_id":
		return x.BuyOrderId != uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.selector_type":
		return x.SelectorType != 0
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.value":
		return x.Value != uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.project_location":
		return x.ProjectLocation != ""
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.min_start_date":
		return x.MinStartDate != nil
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.max_end_date":
		return x.MaxEndDate != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UInt64SelectorBuyOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.buy_order_id":
		x.BuyOrderId = uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.selector_type":
		x.SelectorType = 0
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.value":
		x.Value = uint64(0)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.project_location":
		x.ProjectLocation = ""
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.min_start_date":
		x.MinStartDate = nil
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.max_end_date":
		x.MaxEndDate = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_UInt64SelectorBuyOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.buy_order_id":
		value := x.BuyOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.selector_type":
		value := x.SelectorType
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.value":
		value := x.Value
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.min_start_date":
		value := x.MinStartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.max_end_date":
		value := x.MaxEndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UInt64SelectorBuyOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.buy_order_id":
		x.BuyOrderId = value.Uint()
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.selector_type":
		x.SelectorType = (v1beta1.SelectorType)(value.Enum())
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.value":
		x.Value = value.Uint()
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.project_location":
		x.ProjectLocation = value.Interface().(string)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.min_start_date":
		x.MinStartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.max_end_date":
		x.MaxEndDate = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UInt64SelectorBuyOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.min_start_date":
		if x.MinStartDate == nil {
			x.MinStartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.MinStartDate.ProtoReflect())
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.max_end_date":
		if x.MaxEndDate == nil {
			x.MaxEndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.MaxEndDate.ProtoReflect())
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.buy_order_id":
		panic(fmt.Errorf("field buy_order_id of message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder is not mutable"))
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.selector_type":
		panic(fmt.Errorf("field selector_type of message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder is not mutable"))
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.value":
		panic(fmt.Errorf("field value of message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder is not mutable"))
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_UInt64SelectorBuyOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.buy_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.selector_type":
		return protoreflect.ValueOfEnum(0)
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.value":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.project_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.min_start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.max_end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_UInt64SelectorBuyOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_UInt64SelectorBuyOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UInt64SelectorBuyOrder) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_UInt64SelectorBuyOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_UInt64SelectorBuyOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*UInt64SelectorBuyOrder)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BuyOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.BuyOrderId))
		}
		if x.SelectorType != 0 {
			n += 1 + runtime.Sov(uint64(x.SelectorType))
		}
		if x.Value != 0 {
			n += 1 + runtime.Sov(uint64(x.Value))
		}
		l = len(x.ProjectLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MinStartDate != nil {
			l = options.Size(x.MinStartDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxEndDate != nil {
			l = options.Size(x.MaxEndDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*UInt64SelectorBuyOrder)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.MaxEndDate != nil {
			encoded, err := options.Marshal(x.MaxEndDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x32
		}
		if x.MinStartDate != nil {
			encoded, err := options.Marshal(x.MinStartDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.ProjectLocation) > 0 {
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x22
		}
		if x.Value != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Value))
			i--
			dAtA[i] = 0x18
		}
		if x.SelectorType != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SelectorType))
			i--
			dAtA[i] = 0x10
		}
		if x.BuyOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BuyOrderId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*UInt64SelectorBuyOrder)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: UInt64SelectorBuyOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: UInt64SelectorBuyOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuyOrderId", wireType)
				}
				x.BuyOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BuyOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SelectorType", wireType)
				}
				x.SelectorType = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SelectorType |= v1beta1.SelectorType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
				}
				x.Value = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Value |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectLocation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectLocation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinStartDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MinStartDate == nil {
					x.MinStartDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MinStartDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxEndDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MaxEndDate == nil {
					x.MaxEndDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MaxEndDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/ecocredit/orderbook/v1beta1/memory.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BuyOrderSellOrderMatch defines the data the FIFO/price-time-priority matching
// algorithm used to actually match buy and sell orders.
type BuyOrderSellOrderMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bid_denom_id defines the bid denom being used by the buy and sell orders. Matching always happens within a single bid denom.
	BidDenomId uint32 `protobuf:"varint,1,opt,name=bid_denom_id,json=bidDenomId,proto3" json:"bid_denom_id,omitempty"`
	// buy_order_id is the buy order ID.
	BuyOrderId uint64 `protobuf:"varint,2,opt,name=buy_order_id,json=buyOrderId,proto3" json:"buy_order_id,omitempty"`
	// sell_order_id is the sell order ID.
	SellOrderId uint64 `protobuf:"varint,3,opt,name=sell_order_id,json=sellOrderId,proto3" json:"sell_order_id,omitempty"`
	// bid_price_complement is the the complement (~ operator) of the bid price encoded as a uint64 (which should have sufficient precision) - effectively ~price * 10^exponent (usually 10^6). The complement is used so that bids can be sorted high to low.
	BidPriceComplement uint64 `protobuf:"fixed64,4,opt,name=bid_price_complement,json=bidPriceComplement,proto3" json:"bid_price_complement,omitempty"`
	// ask_price is the ask price encoded to a uint64. Ask prices are sorted low to high.
	AskPrice uint64 `protobuf:"fixed64,5,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
}

func (x *BuyOrderSellOrderMatch) Reset() {
	*x = BuyOrderSellOrderMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrderSellOrderMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrderSellOrderMatch) ProtoMessage() {}

// Deprecated: Use BuyOrderSellOrderMatch.ProtoReflect.Descriptor instead.
func (*BuyOrderSellOrderMatch) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescGZIP(), []int{0}
}

func (x *BuyOrderSellOrderMatch) GetBidDenomId() uint32 {
	if x != nil {
		return x.BidDenomId
	}
	return 0
}

func (x *BuyOrderSellOrderMatch) GetBuyOrderId() uint64 {
	if x != nil {
		return x.BuyOrderId
	}
	return 0
}

func (x *BuyOrderSellOrderMatch) GetSellOrderId() uint64 {
	if x != nil {
		return x.SellOrderId
	}
	return 0
}

func (x *BuyOrderSellOrderMatch) GetBidPriceComplement() uint64 {
	if x != nil {
		return x.BidPriceComplement
	}
	return 0
}

func (x *BuyOrderSellOrderMatch) GetAskPrice() uint64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

// UInt64SelectorBuyOrder indexes a buy order against uint64 selectors in its criteria.
// For example, for a buy order with a selector for a credit class, should insert
/// an entry with type class.
type UInt64SelectorBuyOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buy_order_id is the buy order ID.
	BuyOrderId uint64 `protobuf:"varint,1,opt,name=buy_order_id,json=buyOrderId,proto3" json:"buy_order_id,omitempty"`
	// type is the selector type.
	SelectorType v1beta1.SelectorType `protobuf:"varint,2,opt,name=selector_type,json=selectorType,proto3,enum=regen.ecocredit.marketplace.v1beta1.SelectorType" json:"selector_type,omitempty"`
	// value is the uint64 selector value.
	Value uint64 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	// project_location is the project location in the selector's criteria.
	ProjectLocation string `protobuf:"bytes,4,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
	// min_start_date is the minimum start date in the selector's criteria.
	MinStartDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=min_start_date,json=minStartDate,proto3" json:"min_start_date,omitempty"`
	// max_end_date is the maximum end date in the selector's criteria.
	MaxEndDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=max_end_date,json=maxEndDate,proto3" json:"max_end_date,omitempty"`
}

func (x *UInt64SelectorBuyOrder) Reset() {
	*x = UInt64SelectorBuyOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt64SelectorBuyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt64SelectorBuyOrder) ProtoMessage() {}

// Deprecated: Use UInt64SelectorBuyOrder.ProtoReflect.Descriptor instead.
func (*UInt64SelectorBuyOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescGZIP(), []int{1}
}

func (x *UInt64SelectorBuyOrder) GetBuyOrderId() uint64 {
	if x != nil {
		return x.BuyOrderId
	}
	return 0
}

func (x *UInt64SelectorBuyOrder) GetSelectorType() v1beta1.SelectorType {
	if x != nil {
		return x.SelectorType
	}
	return v1beta1.SelectorType(0)
}

func (x *UInt64SelectorBuyOrder) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *UInt64SelectorBuyOrder) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

func (x *UInt64SelectorBuyOrder) GetMinStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MinStartDate
	}
	return nil
}

func (x *UInt64SelectorBuyOrder) GetMaxEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MaxEndDate
	}
	return nil
}

var File_regen_ecocredit_orderbook_v1beta1_memory_proto protoreflect.FileDescriptor

var file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x16, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x20, 0x0a, 0x0c, 0x62, 0x69, 0x64, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x69, 0x64, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x62, 0x69, 0x64, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x12, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x61, 0x73,
	0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x3a, 0x85, 0x01, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x7f, 0x0a,
	0x1c, 0x0a, 0x1a, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c,
	0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x4a, 0x0a,
	0x46, 0x62, 0x69, 0x64, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x2c, 0x62, 0x69,
	0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2c, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c,
	0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2c, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x73, 0x65, 0x6c,
	0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x10, 0x02, 0x18, 0x01, 0x22, 0x9a,
	0x03, 0x0a, 0x16, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x75, 0x79,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x0d, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x45, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x3a, 0x45, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x3f, 0x0a, 0x22, 0x0a, 0x20,
	0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x17, 0x0a, 0x13, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x01, 0x18, 0x02, 0x42, 0xb9, 0x02, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x4f, 0xaa, 0x02, 0x21, 0x52, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x21, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x2d, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x24, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x3a, 0x3a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescOnce sync.Once
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescData = file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDesc
)

func file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescData)
	})
	return file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDescData
}

var file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_regen_ecocredit_orderbook_v1beta1_memory_proto_goTypes = []interface{}{
	(*BuyOrderSellOrderMatch)(nil), // 0: regen.ecocredit.orderbook.v1beta1.BuyOrderSellOrderMatch
	(*UInt64SelectorBuyOrder)(nil), // 1: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder
	(v1beta1.SelectorType)(0),      // 2: regen.ecocredit.marketplace.v1beta1.SelectorType
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_regen_ecocredit_orderbook_v1beta1_memory_proto_depIdxs = []int32{
	2, // 0: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.selector_type:type_name -> regen.ecocredit.marketplace.v1beta1.SelectorType
	3, // 1: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.min_start_date:type_name -> google.protobuf.Timestamp
	3, // 2: regen.ecocredit.orderbook.v1beta1.UInt64SelectorBuyOrder.max_end_date:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_orderbook_v1beta1_memory_proto_init() }
func file_regen_ecocredit_orderbook_v1beta1_memory_proto_init() {
	if File_regen_ecocredit_orderbook_v1beta1_memory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrderSellOrderMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt64SelectorBuyOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_orderbook_v1beta1_memory_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_orderbook_v1beta1_memory_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_orderbook_v1beta1_memory_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_orderbook_v1beta1_memory_proto = out.File
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_rawDesc = nil
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_goTypes = nil
	file_regen_ecocredit_orderbook_v1beta1_memory_proto_depIdxs = nil
}
