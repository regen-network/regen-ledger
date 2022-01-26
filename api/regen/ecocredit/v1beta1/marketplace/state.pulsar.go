package marketplace

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_SellOrder                     protoreflect.MessageDescriptor
	fd_SellOrder_order_id            protoreflect.FieldDescriptor
	fd_SellOrder_owner               protoreflect.FieldDescriptor
	fd_SellOrder_batch_denom         protoreflect.FieldDescriptor
	fd_SellOrder_quantity            protoreflect.FieldDescriptor
	fd_SellOrder_ask_price           protoreflect.FieldDescriptor
	fd_SellOrder_disable_auto_retire protoreflect.FieldDescriptor
	fd_SellOrder_expiration          protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_marketplace_state_proto_init()
	md_SellOrder = File_regen_ecocredit_v1beta1_marketplace_state_proto.Messages().ByName("SellOrder")
	fd_SellOrder_order_id = md_SellOrder.Fields().ByName("order_id")
	fd_SellOrder_owner = md_SellOrder.Fields().ByName("owner")
	fd_SellOrder_batch_denom = md_SellOrder.Fields().ByName("batch_denom")
	fd_SellOrder_quantity = md_SellOrder.Fields().ByName("quantity")
	fd_SellOrder_ask_price = md_SellOrder.Fields().ByName("ask_price")
	fd_SellOrder_disable_auto_retire = md_SellOrder.Fields().ByName("disable_auto_retire")
	fd_SellOrder_expiration = md_SellOrder.Fields().ByName("expiration")
}

var _ protoreflect.Message = (*fastReflection_SellOrder)(nil)

type fastReflection_SellOrder SellOrder

func (x *SellOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SellOrder)(x)
}

func (x *SellOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SellOrder_messageType fastReflection_SellOrder_messageType
var _ protoreflect.MessageType = fastReflection_SellOrder_messageType{}

type fastReflection_SellOrder_messageType struct{}

func (x fastReflection_SellOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SellOrder)(nil)
}
func (x fastReflection_SellOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_SellOrder)
}
func (x fastReflection_SellOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SellOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SellOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_SellOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SellOrder) Type() protoreflect.MessageType {
	return _fastReflection_SellOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SellOrder) New() protoreflect.Message {
	return new(fastReflection_SellOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SellOrder) Interface() protoreflect.ProtoMessage {
	return (*SellOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SellOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.OrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.OrderId)
		if !f(fd_SellOrder_order_id, value) {
			return
		}
	}
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_SellOrder_owner, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_SellOrder_batch_denom, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_SellOrder_quantity, value) {
			return
		}
	}
	if x.AskPrice != nil {
		value := protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
		if !f(fd_SellOrder_ask_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_SellOrder_disable_auto_retire, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_SellOrder_expiration, value) {
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
func (x *fastReflection_SellOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.order_id":
		return x.OrderId != uint64(0)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.ask_price":
		return x.AskPrice != nil
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.expiration":
		return x.Expiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.SellOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.order_id":
		x.OrderId = uint64(0)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.owner":
		x.Owner = ""
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.ask_price":
		x.AskPrice = nil
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.expiration":
		x.Expiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.SellOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SellOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.order_id":
		value := x.OrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.ask_price":
		value := x.AskPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.SellOrder does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SellOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.order_id":
		x.OrderId = value.Uint()
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.ask_price":
		x.AskPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.SellOrder does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SellOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.ask_price":
		if x.AskPrice == nil {
			x.AskPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.order_id":
		panic(fmt.Errorf("field order_id of message regen.ecocredit.v1beta1.marketplace.SellOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1beta1.marketplace.SellOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1beta1.marketplace.SellOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1beta1.marketplace.SellOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1beta1.marketplace.SellOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.SellOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SellOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.ask_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1beta1.marketplace.SellOrder.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.SellOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SellOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.marketplace.SellOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SellOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SellOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SellOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SellOrder)
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
		if x.OrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.OrderId))
		}
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AskPrice != nil {
			l = options.Size(x.AskPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.Expiration != nil {
			l = options.Size(x.Expiration)
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
		x := input.Message.Interface().(*SellOrder)
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
		if x.Expiration != nil {
			encoded, err := options.Marshal(x.Expiration)
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
			dAtA[i] = 0x3a
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x30
		}
		if x.AskPrice != nil {
			encoded, err := options.Marshal(x.AskPrice)
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
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0x12
		}
		if x.OrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.OrderId))
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
		x := input.Message.Interface().(*SellOrder)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SellOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SellOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
				}
				x.OrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.OrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
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
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
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
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AskPrice", wireType)
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
				if x.AskPrice == nil {
					x.AskPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AskPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
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
				if x.Expiration == nil {
					x.Expiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expiration); err != nil {
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

var (
	md_BuyOrder                      protoreflect.MessageDescriptor
	fd_BuyOrder_buy_order_id         protoreflect.FieldDescriptor
	fd_BuyOrder_buyer                protoreflect.FieldDescriptor
	fd_BuyOrder_selection            protoreflect.FieldDescriptor
	fd_BuyOrder_quantity             protoreflect.FieldDescriptor
	fd_BuyOrder_bid_price            protoreflect.FieldDescriptor
	fd_BuyOrder_disable_auto_retire  protoreflect.FieldDescriptor
	fd_BuyOrder_disable_partial_fill protoreflect.FieldDescriptor
	fd_BuyOrder_expiration           protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_marketplace_state_proto_init()
	md_BuyOrder = File_regen_ecocredit_v1beta1_marketplace_state_proto.Messages().ByName("BuyOrder")
	fd_BuyOrder_buy_order_id = md_BuyOrder.Fields().ByName("buy_order_id")
	fd_BuyOrder_buyer = md_BuyOrder.Fields().ByName("buyer")
	fd_BuyOrder_selection = md_BuyOrder.Fields().ByName("selection")
	fd_BuyOrder_quantity = md_BuyOrder.Fields().ByName("quantity")
	fd_BuyOrder_bid_price = md_BuyOrder.Fields().ByName("bid_price")
	fd_BuyOrder_disable_auto_retire = md_BuyOrder.Fields().ByName("disable_auto_retire")
	fd_BuyOrder_disable_partial_fill = md_BuyOrder.Fields().ByName("disable_partial_fill")
	fd_BuyOrder_expiration = md_BuyOrder.Fields().ByName("expiration")
}

var _ protoreflect.Message = (*fastReflection_BuyOrder)(nil)

type fastReflection_BuyOrder BuyOrder

func (x *BuyOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrder)(x)
}

func (x *BuyOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrder_messageType fastReflection_BuyOrder_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrder_messageType{}

type fastReflection_BuyOrder_messageType struct{}

func (x fastReflection_BuyOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrder)(nil)
}
func (x fastReflection_BuyOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrder)
}
func (x fastReflection_BuyOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrder) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrder) New() protoreflect.Message {
	return new(fastReflection_BuyOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrder) Interface() protoreflect.ProtoMessage {
	return (*BuyOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BuyOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BuyOrderId)
		if !f(fd_BuyOrder_buy_order_id, value) {
			return
		}
	}
	if x.Buyer != "" {
		value := protoreflect.ValueOfString(x.Buyer)
		if !f(fd_BuyOrder_buyer, value) {
			return
		}
	}
	if x.Selection != nil {
		value := protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
		if !f(fd_BuyOrder_selection, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_BuyOrder_quantity, value) {
			return
		}
	}
	if x.BidPrice != nil {
		value := protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
		if !f(fd_BuyOrder_bid_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_BuyOrder_disable_auto_retire, value) {
			return
		}
	}
	if x.DisablePartialFill != false {
		value := protoreflect.ValueOfBool(x.DisablePartialFill)
		if !f(fd_BuyOrder_disable_partial_fill, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_BuyOrder_expiration, value) {
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
func (x *fastReflection_BuyOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buy_order_id":
		return x.BuyOrderId != uint64(0)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buyer":
		return x.Buyer != ""
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.selection":
		return x.Selection != nil
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.bid_price":
		return x.BidPrice != nil
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_partial_fill":
		return x.DisablePartialFill != false
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.expiration":
		return x.Expiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buy_order_id":
		x.BuyOrderId = uint64(0)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buyer":
		x.Buyer = ""
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.selection":
		x.Selection = nil
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.bid_price":
		x.BidPrice = nil
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_partial_fill":
		x.DisablePartialFill = false
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.expiration":
		x.Expiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buy_order_id":
		value := x.BuyOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buyer":
		value := x.Buyer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.selection":
		value := x.Selection
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.bid_price":
		value := x.BidPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_partial_fill":
		value := x.DisablePartialFill
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BuyOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buy_order_id":
		x.BuyOrderId = value.Uint()
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buyer":
		x.Buyer = value.Interface().(string)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.selection":
		x.Selection = value.Message().Interface().(*BuyOrder_Selection)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.bid_price":
		x.BidPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_partial_fill":
		x.DisablePartialFill = value.Bool()
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BuyOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.selection":
		if x.Selection == nil {
			x.Selection = new(BuyOrder_Selection)
		}
		return protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.bid_price":
		if x.BidPrice == nil {
			x.BidPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buy_order_id":
		panic(fmt.Errorf("field buy_order_id of message regen.ecocredit.v1beta1.marketplace.BuyOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buyer":
		panic(fmt.Errorf("field buyer of message regen.ecocredit.v1beta1.marketplace.BuyOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1beta1.marketplace.BuyOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1beta1.marketplace.BuyOrder is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_partial_fill":
		panic(fmt.Errorf("field disable_partial_fill of message regen.ecocredit.v1beta1.marketplace.BuyOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buy_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.buyer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.selection":
		m := new(BuyOrder_Selection)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.bid_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.disable_partial_fill":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.marketplace.BuyOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BuyOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrder)
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
		l = len(x.Buyer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Selection != nil {
			l = options.Size(x.Selection)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BidPrice != nil {
			l = options.Size(x.BidPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.DisablePartialFill {
			n += 2
		}
		if x.Expiration != nil {
			l = options.Size(x.Expiration)
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
		x := input.Message.Interface().(*BuyOrder)
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
		if x.Expiration != nil {
			encoded, err := options.Marshal(x.Expiration)
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
			dAtA[i] = 0x42
		}
		if x.DisablePartialFill {
			i--
			if x.DisablePartialFill {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x38
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x30
		}
		if x.BidPrice != nil {
			encoded, err := options.Marshal(x.BidPrice)
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
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if x.Selection != nil {
			encoded, err := options.Marshal(x.Selection)
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
			dAtA[i] = 0x1a
		}
		if len(x.Buyer) > 0 {
			i -= len(x.Buyer)
			copy(dAtA[i:], x.Buyer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Buyer)))
			i--
			dAtA[i] = 0x12
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
		x := input.Message.Interface().(*BuyOrder)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder: illegal tag %d (wire type %d)", fieldNum, wire)
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
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
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
				x.Buyer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Selection", wireType)
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
				if x.Selection == nil {
					x.Selection = &BuyOrder_Selection{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Selection); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
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
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BidPrice", wireType)
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
				if x.BidPrice == nil {
					x.BidPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BidPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisablePartialFill", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisablePartialFill = bool(v != 0)
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
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
				if x.Expiration == nil {
					x.Expiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expiration); err != nil {
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

var (
	md_BuyOrder_Selection               protoreflect.MessageDescriptor
	fd_BuyOrder_Selection_sell_order_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_marketplace_state_proto_init()
	md_BuyOrder_Selection = File_regen_ecocredit_v1beta1_marketplace_state_proto.Messages().ByName("BuyOrder").Messages().ByName("Selection")
	fd_BuyOrder_Selection_sell_order_id = md_BuyOrder_Selection.Fields().ByName("sell_order_id")
}

var _ protoreflect.Message = (*fastReflection_BuyOrder_Selection)(nil)

type fastReflection_BuyOrder_Selection BuyOrder_Selection

func (x *BuyOrder_Selection) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrder_Selection)(x)
}

func (x *BuyOrder_Selection) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrder_Selection_messageType fastReflection_BuyOrder_Selection_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrder_Selection_messageType{}

type fastReflection_BuyOrder_Selection_messageType struct{}

func (x fastReflection_BuyOrder_Selection_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrder_Selection)(nil)
}
func (x fastReflection_BuyOrder_Selection_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrder_Selection)
}
func (x fastReflection_BuyOrder_Selection_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder_Selection
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrder_Selection) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder_Selection
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrder_Selection) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrder_Selection_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrder_Selection) New() protoreflect.Message {
	return new(fastReflection_BuyOrder_Selection)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrder_Selection) Interface() protoreflect.ProtoMessage {
	return (*BuyOrder_Selection)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrder_Selection) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			v := o.SellOrderId
			value := protoreflect.ValueOfUint64(v)
			if !f(fd_BuyOrder_Selection_sell_order_id, value) {
				return
			}
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
func (x *fastReflection_BuyOrder_Selection) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection.sell_order_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*BuyOrder_Selection_SellOrderId); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection.sell_order_id":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrder_Selection) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection.sell_order_id":
		if x.Sum == nil {
			return protoreflect.ValueOfUint64(uint64(0))
		} else if v, ok := x.Sum.(*BuyOrder_Selection_SellOrderId); ok {
			return protoreflect.ValueOfUint64(v.SellOrderId)
		} else {
			return protoreflect.ValueOfUint64(uint64(0))
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BuyOrder_Selection) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection.sell_order_id":
		cv := value.Uint()
		x.Sum = &BuyOrder_Selection_SellOrderId{SellOrderId: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BuyOrder_Selection) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrder_Selection) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrder_Selection) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			return x.Descriptor().Fields().ByName("sell_order_id")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrder_Selection) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BuyOrder_Selection) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrder_Selection) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrder_Selection)
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
		switch x := x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			if x == nil {
				break
			}
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
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
		x := input.Message.Interface().(*BuyOrder_Selection)
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
		switch x := x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SellOrderId))
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
		x := input.Message.Interface().(*BuyOrder_Selection)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder_Selection: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder_Selection: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderId", wireType)
				}
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Sum = &BuyOrder_Selection_SellOrderId{v}
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
	md_AskDenom               protoreflect.MessageDescriptor
	fd_AskDenom_denom         protoreflect.FieldDescriptor
	fd_AskDenom_display_denom protoreflect.FieldDescriptor
	fd_AskDenom_exponent      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_marketplace_state_proto_init()
	md_AskDenom = File_regen_ecocredit_v1beta1_marketplace_state_proto.Messages().ByName("AskDenom")
	fd_AskDenom_denom = md_AskDenom.Fields().ByName("denom")
	fd_AskDenom_display_denom = md_AskDenom.Fields().ByName("display_denom")
	fd_AskDenom_exponent = md_AskDenom.Fields().ByName("exponent")
}

var _ protoreflect.Message = (*fastReflection_AskDenom)(nil)

type fastReflection_AskDenom AskDenom

func (x *AskDenom) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AskDenom)(x)
}

func (x *AskDenom) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AskDenom_messageType fastReflection_AskDenom_messageType
var _ protoreflect.MessageType = fastReflection_AskDenom_messageType{}

type fastReflection_AskDenom_messageType struct{}

func (x fastReflection_AskDenom_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AskDenom)(nil)
}
func (x fastReflection_AskDenom_messageType) New() protoreflect.Message {
	return new(fastReflection_AskDenom)
}
func (x fastReflection_AskDenom_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AskDenom
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AskDenom) Descriptor() protoreflect.MessageDescriptor {
	return md_AskDenom
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AskDenom) Type() protoreflect.MessageType {
	return _fastReflection_AskDenom_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AskDenom) New() protoreflect.Message {
	return new(fastReflection_AskDenom)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AskDenom) Interface() protoreflect.ProtoMessage {
	return (*AskDenom)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AskDenom) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_AskDenom_denom, value) {
			return
		}
	}
	if x.DisplayDenom != "" {
		value := protoreflect.ValueOfString(x.DisplayDenom)
		if !f(fd_AskDenom_display_denom, value) {
			return
		}
	}
	if x.Exponent != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Exponent)
		if !f(fd_AskDenom_exponent, value) {
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
func (x *fastReflection_AskDenom) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.denom":
		return x.Denom != ""
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.display_denom":
		return x.DisplayDenom != ""
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.exponent":
		return x.Exponent != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.AskDenom does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AskDenom) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.denom":
		x.Denom = ""
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.display_denom":
		x.DisplayDenom = ""
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.exponent":
		x.Exponent = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.AskDenom does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AskDenom) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.display_denom":
		value := x.DisplayDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.exponent":
		value := x.Exponent
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.AskDenom does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_AskDenom) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.denom":
		x.Denom = value.Interface().(string)
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.display_denom":
		x.DisplayDenom = value.Interface().(string)
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.exponent":
		x.Exponent = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.AskDenom does not contain field %s", fd.FullName()))
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
func (x *fastReflection_AskDenom) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.denom":
		panic(fmt.Errorf("field denom of message regen.ecocredit.v1beta1.marketplace.AskDenom is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.display_denom":
		panic(fmt.Errorf("field display_denom of message regen.ecocredit.v1beta1.marketplace.AskDenom is not mutable"))
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.exponent":
		panic(fmt.Errorf("field exponent of message regen.ecocredit.v1beta1.marketplace.AskDenom is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.AskDenom does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AskDenom) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.display_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.marketplace.AskDenom.exponent":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.marketplace.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.marketplace.AskDenom does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AskDenom) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.marketplace.AskDenom", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AskDenom) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AskDenom) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_AskDenom) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AskDenom) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AskDenom)
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
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DisplayDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Exponent != 0 {
			n += 1 + runtime.Sov(uint64(x.Exponent))
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
		x := input.Message.Interface().(*AskDenom)
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
		if x.Exponent != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exponent))
			i--
			dAtA[i] = 0x18
		}
		if len(x.DisplayDenom) > 0 {
			i -= len(x.DisplayDenom)
			copy(dAtA[i:], x.DisplayDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DisplayDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
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
		x := input.Message.Interface().(*AskDenom)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AskDenom: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AskDenom: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisplayDenom", wireType)
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
				x.DisplayDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
				}
				x.Exponent = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exponent |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: regen/ecocredit/v1beta1/marketplace/state.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SellOrder represents the information for a sell order.
type SellOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// order_id is the unique ID of sell order.
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// owner is the address of the owner of the credits being sold.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// batch_denom is the credit batch being sold.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// quantity is the quantity of credits being sold.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// ask_price is the price the seller is asking for each unit of the
	// batch_denom. Each credit unit of the batch will be sold for at least the
	// ask_price or more.
	AskPrice *v1beta1.Coin `protobuf:"bytes,5,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	// disable_auto_retire disables auto-retirement of credits which allows a
	// buyer to disable auto-retirement in their buy order enabling them to
	// resell the credits to another buyer.
	DisableAutoRetire bool `protobuf:"varint,6,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// expiration is an optional timestamp when the sell order expires. When the
	// expiration time is reached, the sell order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *SellOrder) Reset() {
	*x = SellOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellOrder) ProtoMessage() {}

// Deprecated: Use SellOrder.ProtoReflect.Descriptor instead.
func (*SellOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescGZIP(), []int{0}
}

func (x *SellOrder) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SellOrder) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SellOrder) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *SellOrder) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *SellOrder) GetAskPrice() *v1beta1.Coin {
	if x != nil {
		return x.AskPrice
	}
	return nil
}

func (x *SellOrder) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *SellOrder) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// BuyOrder represents the information for a buy order.
type BuyOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buy_order_id is the unique ID of buy order.
	BuyOrderId uint64 `protobuf:"varint,1,opt,name=buy_order_id,json=buyOrderId,proto3" json:"buy_order_id,omitempty"`
	// buyer is the address that created the buy order
	Buyer string `protobuf:"bytes,2,opt,name=buyer,proto3" json:"buyer,omitempty"`
	// selection is the buy order selection.
	Selection *BuyOrder_Selection `protobuf:"bytes,3,opt,name=selection,proto3" json:"selection,omitempty"`
	// quantity is the quantity of credits to buy. If the quantity of credits
	// available is less than this amount the order will be partially filled
	// unless disable_partial_fill is true.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// bid price is the bid price for this buy order. A credit unit will be
	// settled at a purchase price that is no more than the bid price. The
	// buy order will fail if the buyer does not have enough funds available
	// to complete the purchase.
	BidPrice *v1beta1.Coin `protobuf:"bytes,5,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	// disable_auto_retire allows auto-retirement to be disabled. If it is set to true
	// the credits will not auto-retire and can be resold assuming that the
	// corresponding sell order has auto-retirement disabled. If the sell order
	// hasn't disabled auto-retirement and the buy order tries to disable it,
	// that buy order will fail.
	DisableAutoRetire bool `protobuf:"varint,6,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// disable_partial_fill disables the default behavior of partially filling
	// buy orders if the requested quantity is not available.
	DisablePartialFill bool `protobuf:"varint,7,opt,name=disable_partial_fill,json=disablePartialFill,proto3" json:"disable_partial_fill,omitempty"`
	// expiration is the optional timestamp when the buy order expires. When the
	// expiration time is reached, the buy order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *BuyOrder) Reset() {
	*x = BuyOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrder) ProtoMessage() {}

// Deprecated: Use BuyOrder.ProtoReflect.Descriptor instead.
func (*BuyOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescGZIP(), []int{1}
}

func (x *BuyOrder) GetBuyOrderId() uint64 {
	if x != nil {
		return x.BuyOrderId
	}
	return 0
}

func (x *BuyOrder) GetBuyer() string {
	if x != nil {
		return x.Buyer
	}
	return ""
}

func (x *BuyOrder) GetSelection() *BuyOrder_Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

func (x *BuyOrder) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *BuyOrder) GetBidPrice() *v1beta1.Coin {
	if x != nil {
		return x.BidPrice
	}
	return nil
}

func (x *BuyOrder) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *BuyOrder) GetDisablePartialFill() bool {
	if x != nil {
		return x.DisablePartialFill
	}
	return false
}

func (x *BuyOrder) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// AskDenom represents the information for an ask denom.
type AskDenom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// denom is the denom to allow (ex. ibc/GLKHDSG423SGS)
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// display_denom is the denom to display to the user and is informational
	DisplayDenom string `protobuf:"bytes,2,opt,name=display_denom,json=displayDenom,proto3" json:"display_denom,omitempty"`
	// exponent is the exponent that relates the denom to the display_denom and is
	// informational
	Exponent uint32 `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (x *AskDenom) Reset() {
	*x = AskDenom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskDenom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskDenom) ProtoMessage() {}

// Deprecated: Use AskDenom.ProtoReflect.Descriptor instead.
func (*AskDenom) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescGZIP(), []int{2}
}

func (x *AskDenom) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *AskDenom) GetDisplayDenom() string {
	if x != nil {
		return x.DisplayDenom
	}
	return ""
}

func (x *AskDenom) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

// Selection defines a buy order selection.
type BuyOrder_Selection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum defines the type of selection.
	//
	// Types that are assignable to Sum:
	//	*BuyOrder_Selection_SellOrderId
	Sum isBuyOrder_Selection_Sum `protobuf_oneof:"sum"`
}

func (x *BuyOrder_Selection) Reset() {
	*x = BuyOrder_Selection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrder_Selection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrder_Selection) ProtoMessage() {}

// Deprecated: Use BuyOrder_Selection.ProtoReflect.Descriptor instead.
func (*BuyOrder_Selection) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescGZIP(), []int{1, 0}
}

func (x *BuyOrder_Selection) GetSum() isBuyOrder_Selection_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *BuyOrder_Selection) GetSellOrderId() uint64 {
	if x, ok := x.GetSum().(*BuyOrder_Selection_SellOrderId); ok {
		return x.SellOrderId
	}
	return 0
}

type isBuyOrder_Selection_Sum interface {
	isBuyOrder_Selection_Sum()
}

type BuyOrder_Selection_SellOrderId struct {
	// sell_order_id is the sell order ID against which the buyer is trying to buy.
	// When sell_order_id is set, this is known as a direct buy order because it
	// is placed directly against a specific sell order.
	SellOrderId uint64 `protobuf:"varint,1,opt,name=sell_order_id,json=sellOrderId,proto3,oneof"`
}

func (*BuyOrder_Selection_SellOrderId) isBuyOrder_Selection_Sum() {}

var File_regen_ecocredit_v1beta1_marketplace_state_proto protoreflect.FileDescriptor

var file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x1a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f,
	0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a,
	0x09, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65,
	0x12, 0x40, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x40, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x3a, 0x0a, 0x0a, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x03, 0x18, 0x01, 0x22, 0x80, 0x04, 0x0a, 0x08, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x09, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x09,
	0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x74, 0x69, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x46, 0x69, 0x6c, 0x6c, 0x12, 0x40, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x38, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x73,
	0x75, 0x6d, 0x3a, 0x33, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x2d, 0x0a, 0x0e, 0x0a, 0x0c, 0x62, 0x75,
	0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x75,
	0x79, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x10, 0x02, 0x18, 0x02, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x41, 0x73, 0x6b, 0x44,
	0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x3a, 0x24, 0xf2, 0x9e, 0xd3,
	0x8e, 0x03, 0x1e, 0x0a, 0x07, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x11, 0x0a, 0x0d,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x10, 0x01, 0x18,
	0x03, 0x42, 0xb4, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0xa2, 0x02, 0x04, 0x52, 0x45, 0x56,
	0x4d, 0xaa, 0x02, 0x23, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0xca, 0x02, 0x23, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c,
	0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x5c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0xe2, 0x02, 0x2f,
	0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x26, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3a, 0x3a, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescOnce sync.Once
	file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescData = file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDesc
)

func file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescData)
	})
	return file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDescData
}

var file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_regen_ecocredit_v1beta1_marketplace_state_proto_goTypes = []interface{}{
	(*SellOrder)(nil),             // 0: regen.ecocredit.v1beta1.marketplace.SellOrder
	(*BuyOrder)(nil),              // 1: regen.ecocredit.v1beta1.marketplace.BuyOrder
	(*AskDenom)(nil),              // 2: regen.ecocredit.v1beta1.marketplace.AskDenom
	(*BuyOrder_Selection)(nil),    // 3: regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection
	(*v1beta1.Coin)(nil),          // 4: cosmos.base.v1beta1.Coin
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_regen_ecocredit_v1beta1_marketplace_state_proto_depIdxs = []int32{
	4, // 0: regen.ecocredit.v1beta1.marketplace.SellOrder.ask_price:type_name -> cosmos.base.v1beta1.Coin
	5, // 1: regen.ecocredit.v1beta1.marketplace.SellOrder.expiration:type_name -> google.protobuf.Timestamp
	3, // 2: regen.ecocredit.v1beta1.marketplace.BuyOrder.selection:type_name -> regen.ecocredit.v1beta1.marketplace.BuyOrder.Selection
	4, // 3: regen.ecocredit.v1beta1.marketplace.BuyOrder.bid_price:type_name -> cosmos.base.v1beta1.Coin
	5, // 4: regen.ecocredit.v1beta1.marketplace.BuyOrder.expiration:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_v1beta1_marketplace_state_proto_init() }
func file_regen_ecocredit_v1beta1_marketplace_state_proto_init() {
	if File_regen_ecocredit_v1beta1_marketplace_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellOrder); i {
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
		file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrder); i {
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
		file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskDenom); i {
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
		file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrder_Selection); i {
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
	file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*BuyOrder_Selection_SellOrderId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_v1beta1_marketplace_state_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_v1beta1_marketplace_state_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_v1beta1_marketplace_state_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_v1beta1_marketplace_state_proto = out.File
	file_regen_ecocredit_v1beta1_marketplace_state_proto_rawDesc = nil
	file_regen_ecocredit_v1beta1_marketplace_state_proto_goTypes = nil
	file_regen_ecocredit_v1beta1_marketplace_state_proto_depIdxs = nil
}
