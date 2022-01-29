package ecocreditv1alpha2

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_MsgCreateClass_2_list)(nil)

type _MsgCreateClass_2_list struct {
	list *[]string
}

func (x *_MsgCreateClass_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateClass_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_MsgCreateClass_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateClass_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateClass_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgCreateClass at list field Issuers as it is not of Message kind"))
}

func (x *_MsgCreateClass_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateClass_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_MsgCreateClass_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateClass                  protoreflect.MessageDescriptor
	fd_MsgCreateClass_admin            protoreflect.FieldDescriptor
	fd_MsgCreateClass_issuers          protoreflect.FieldDescriptor
	fd_MsgCreateClass_metadata         protoreflect.FieldDescriptor
	fd_MsgCreateClass_credit_type_name protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateClass = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateClass")
	fd_MsgCreateClass_admin = md_MsgCreateClass.Fields().ByName("admin")
	fd_MsgCreateClass_issuers = md_MsgCreateClass.Fields().ByName("issuers")
	fd_MsgCreateClass_metadata = md_MsgCreateClass.Fields().ByName("metadata")
	fd_MsgCreateClass_credit_type_name = md_MsgCreateClass.Fields().ByName("credit_type_name")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateClass)(nil)

type fastReflection_MsgCreateClass MsgCreateClass

func (x *MsgCreateClass) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateClass)(x)
}

func (x *MsgCreateClass) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateClass_messageType fastReflection_MsgCreateClass_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateClass_messageType{}

type fastReflection_MsgCreateClass_messageType struct{}

func (x fastReflection_MsgCreateClass_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateClass)(nil)
}
func (x fastReflection_MsgCreateClass_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClass)
}
func (x fastReflection_MsgCreateClass_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClass
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateClass) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClass
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateClass) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateClass_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateClass) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClass)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateClass) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateClass)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateClass) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgCreateClass_admin, value) {
			return
		}
	}
	if len(x.Issuers) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateClass_2_list{list: &x.Issuers})
		if !f(fd_MsgCreateClass_issuers, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateClass_metadata, value) {
			return
		}
	}
	if x.CreditTypeName != "" {
		value := protoreflect.ValueOfString(x.CreditTypeName)
		if !f(fd_MsgCreateClass_credit_type_name, value) {
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
func (x *fastReflection_MsgCreateClass) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClass.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha2.MsgCreateClass.issuers":
		return len(x.Issuers) != 0
	case "regen.ecocredit.v1alpha2.MsgCreateClass.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1alpha2.MsgCreateClass.credit_type_name":
		return x.CreditTypeName != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClass) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClass.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha2.MsgCreateClass.issuers":
		x.Issuers = nil
	case "regen.ecocredit.v1alpha2.MsgCreateClass.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1alpha2.MsgCreateClass.credit_type_name":
		x.CreditTypeName = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateClass) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClass.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateClass.issuers":
		if len(x.Issuers) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateClass_2_list{})
		}
		listValue := &_MsgCreateClass_2_list{list: &x.Issuers}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.MsgCreateClass.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1alpha2.MsgCreateClass.credit_type_name":
		value := x.CreditTypeName
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClass does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateClass) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClass.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateClass.issuers":
		lv := value.List()
		clv := lv.(*_MsgCreateClass_2_list)
		x.Issuers = *clv.list
	case "regen.ecocredit.v1alpha2.MsgCreateClass.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1alpha2.MsgCreateClass.credit_type_name":
		x.CreditTypeName = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClass does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateClass) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClass.issuers":
		if x.Issuers == nil {
			x.Issuers = []string{}
		}
		value := &_MsgCreateClass_2_list{list: &x.Issuers}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgCreateClass.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha2.MsgCreateClass is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateClass.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha2.MsgCreateClass is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateClass.credit_type_name":
		panic(fmt.Errorf("field credit_type_name of message regen.ecocredit.v1alpha2.MsgCreateClass is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateClass) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClass.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateClass.issuers":
		list := []string{}
		return protoreflect.ValueOfList(&_MsgCreateClass_2_list{list: &list})
	case "regen.ecocredit.v1alpha2.MsgCreateClass.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1alpha2.MsgCreateClass.credit_type_name":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateClass) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateClass", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateClass) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClass) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateClass) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateClass) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateClass)
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
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Issuers) > 0 {
			for _, s := range x.Issuers {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.CreditTypeName)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCreateClass)
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
		if len(x.CreditTypeName) > 0 {
			i -= len(x.CreditTypeName)
			copy(dAtA[i:], x.CreditTypeName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CreditTypeName)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Issuers) > 0 {
			for iNdEx := len(x.Issuers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Issuers[iNdEx])
				copy(dAtA[i:], x.Issuers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuers[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgCreateClass)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClass: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClass: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
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
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuers", wireType)
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
				x.Issuers = append(x.Issuers, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditTypeName", wireType)
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
				x.CreditTypeName = string(dAtA[iNdEx:postIndex])
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
	md_MsgCreateClassResponse          protoreflect.MessageDescriptor
	fd_MsgCreateClassResponse_class_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateClassResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateClassResponse")
	fd_MsgCreateClassResponse_class_id = md_MsgCreateClassResponse.Fields().ByName("class_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateClassResponse)(nil)

type fastReflection_MsgCreateClassResponse MsgCreateClassResponse

func (x *MsgCreateClassResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateClassResponse)(x)
}

func (x *MsgCreateClassResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateClassResponse_messageType fastReflection_MsgCreateClassResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateClassResponse_messageType{}

type fastReflection_MsgCreateClassResponse_messageType struct{}

func (x fastReflection_MsgCreateClassResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateClassResponse)(nil)
}
func (x fastReflection_MsgCreateClassResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClassResponse)
}
func (x fastReflection_MsgCreateClassResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClassResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateClassResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClassResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateClassResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateClassResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateClassResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClassResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateClassResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateClassResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateClassResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgCreateClassResponse_class_id, value) {
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
func (x *fastReflection_MsgCreateClassResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClassResponse.class_id":
		return x.ClassId != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClassResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClassResponse.class_id":
		x.ClassId = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateClassResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClassResponse.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClassResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateClassResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClassResponse.class_id":
		x.ClassId = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClassResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateClassResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClassResponse.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.MsgCreateClassResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateClassResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateClassResponse.class_id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateClassResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateClassResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateClassResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClassResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateClassResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateClassResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateClassResponse)
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
		l = len(x.ClassId)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCreateClassResponse)
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
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
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
		x := input.Message.Interface().(*MsgCreateClassResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClassResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClassResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
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
				x.ClassId = string(dAtA[iNdEx:postIndex])
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
	md_MsgCreateProject                  protoreflect.MessageDescriptor
	fd_MsgCreateProject_issuer           protoreflect.FieldDescriptor
	fd_MsgCreateProject_class_id         protoreflect.FieldDescriptor
	fd_MsgCreateProject_metadata         protoreflect.FieldDescriptor
	fd_MsgCreateProject_project_location protoreflect.FieldDescriptor
	fd_MsgCreateProject_project_id       protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateProject = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateProject")
	fd_MsgCreateProject_issuer = md_MsgCreateProject.Fields().ByName("issuer")
	fd_MsgCreateProject_class_id = md_MsgCreateProject.Fields().ByName("class_id")
	fd_MsgCreateProject_metadata = md_MsgCreateProject.Fields().ByName("metadata")
	fd_MsgCreateProject_project_location = md_MsgCreateProject.Fields().ByName("project_location")
	fd_MsgCreateProject_project_id = md_MsgCreateProject.Fields().ByName("project_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateProject)(nil)

type fastReflection_MsgCreateProject MsgCreateProject

func (x *MsgCreateProject) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateProject)(x)
}

func (x *MsgCreateProject) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateProject_messageType fastReflection_MsgCreateProject_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateProject_messageType{}

type fastReflection_MsgCreateProject_messageType struct{}

func (x fastReflection_MsgCreateProject_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateProject)(nil)
}
func (x fastReflection_MsgCreateProject_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProject)
}
func (x fastReflection_MsgCreateProject_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProject
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateProject) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProject
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateProject) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateProject_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateProject) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProject)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateProject) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateProject)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateProject) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Issuer != "" {
		value := protoreflect.ValueOfString(x.Issuer)
		if !f(fd_MsgCreateProject_issuer, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgCreateProject_class_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateProject_metadata, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_MsgCreateProject_project_location, value) {
			return
		}
	}
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_MsgCreateProject_project_id, value) {
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
func (x *fastReflection_MsgCreateProject) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProject.issuer":
		return x.Issuer != ""
	case "regen.ecocredit.v1alpha2.MsgCreateProject.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.MsgCreateProject.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_location":
		return x.ProjectLocation != ""
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_id":
		return x.ProjectId != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProject does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProject) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProject.issuer":
		x.Issuer = ""
	case "regen.ecocredit.v1alpha2.MsgCreateProject.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.MsgCreateProject.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_location":
		x.ProjectLocation = ""
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_id":
		x.ProjectId = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProject does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateProject) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProject.issuer":
		value := x.Issuer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProject does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateProject) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProject.issuer":
		x.Issuer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_location":
		x.ProjectLocation = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_id":
		x.ProjectId = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProject does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateProject) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProject.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1alpha2.MsgCreateProject is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateProject.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.MsgCreateProject is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateProject.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha2.MsgCreateProject is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.v1alpha2.MsgCreateProject is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.MsgCreateProject is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProject does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateProject) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProject.issuer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateProject.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateProject.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateProject.project_id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProject does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateProject) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateProject", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateProject) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProject) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateProject) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateProject) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateProject)
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
		l = len(x.Issuer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProjectLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProjectId)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCreateProject)
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
		if len(x.ProjectId) > 0 {
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
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
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Issuer) > 0 {
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
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
		x := input.Message.Interface().(*MsgCreateProject)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProject: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProject: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
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
				x.Issuer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
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
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
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
				x.ProjectId = string(dAtA[iNdEx:postIndex])
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
	md_MsgCreateProjectResponse            protoreflect.MessageDescriptor
	fd_MsgCreateProjectResponse_project_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateProjectResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateProjectResponse")
	fd_MsgCreateProjectResponse_project_id = md_MsgCreateProjectResponse.Fields().ByName("project_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateProjectResponse)(nil)

type fastReflection_MsgCreateProjectResponse MsgCreateProjectResponse

func (x *MsgCreateProjectResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateProjectResponse)(x)
}

func (x *MsgCreateProjectResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateProjectResponse_messageType fastReflection_MsgCreateProjectResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateProjectResponse_messageType{}

type fastReflection_MsgCreateProjectResponse_messageType struct{}

func (x fastReflection_MsgCreateProjectResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateProjectResponse)(nil)
}
func (x fastReflection_MsgCreateProjectResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProjectResponse)
}
func (x fastReflection_MsgCreateProjectResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProjectResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateProjectResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProjectResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateProjectResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateProjectResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateProjectResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProjectResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateProjectResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateProjectResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateProjectResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_MsgCreateProjectResponse_project_id, value) {
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
func (x *fastReflection_MsgCreateProjectResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProjectResponse.project_id":
		return x.ProjectId != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProjectResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProjectResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProjectResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProjectResponse.project_id":
		x.ProjectId = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProjectResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProjectResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateProjectResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProjectResponse.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProjectResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProjectResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateProjectResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProjectResponse.project_id":
		x.ProjectId = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProjectResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProjectResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateProjectResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProjectResponse.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.MsgCreateProjectResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProjectResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProjectResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateProjectResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateProjectResponse.project_id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateProjectResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateProjectResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateProjectResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateProjectResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateProjectResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProjectResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateProjectResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateProjectResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateProjectResponse)
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
		l = len(x.ProjectId)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCreateProjectResponse)
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
		if len(x.ProjectId) > 0 {
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
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
		x := input.Message.Interface().(*MsgCreateProjectResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProjectResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProjectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
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
				x.ProjectId = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_MsgCreateBatch_3_list)(nil)

type _MsgCreateBatch_3_list struct {
	list *[]*MsgCreateBatch_BatchIssuance
}

func (x *_MsgCreateBatch_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateBatch_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateBatch_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCreateBatch_BatchIssuance)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateBatch_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCreateBatch_BatchIssuance)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateBatch_3_list) AppendMutable() protoreflect.Value {
	v := new(MsgCreateBatch_BatchIssuance)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateBatch_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateBatch_3_list) NewElement() protoreflect.Value {
	v := new(MsgCreateBatch_BatchIssuance)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateBatch_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateBatch            protoreflect.MessageDescriptor
	fd_MsgCreateBatch_issuer     protoreflect.FieldDescriptor
	fd_MsgCreateBatch_project_id protoreflect.FieldDescriptor
	fd_MsgCreateBatch_issuance   protoreflect.FieldDescriptor
	fd_MsgCreateBatch_metadata   protoreflect.FieldDescriptor
	fd_MsgCreateBatch_start_date protoreflect.FieldDescriptor
	fd_MsgCreateBatch_end_date   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateBatch = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateBatch")
	fd_MsgCreateBatch_issuer = md_MsgCreateBatch.Fields().ByName("issuer")
	fd_MsgCreateBatch_project_id = md_MsgCreateBatch.Fields().ByName("project_id")
	fd_MsgCreateBatch_issuance = md_MsgCreateBatch.Fields().ByName("issuance")
	fd_MsgCreateBatch_metadata = md_MsgCreateBatch.Fields().ByName("metadata")
	fd_MsgCreateBatch_start_date = md_MsgCreateBatch.Fields().ByName("start_date")
	fd_MsgCreateBatch_end_date = md_MsgCreateBatch.Fields().ByName("end_date")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBatch)(nil)

type fastReflection_MsgCreateBatch MsgCreateBatch

func (x *MsgCreateBatch) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch)(x)
}

func (x *MsgCreateBatch) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBatch_messageType fastReflection_MsgCreateBatch_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBatch_messageType{}

type fastReflection_MsgCreateBatch_messageType struct{}

func (x fastReflection_MsgCreateBatch_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch)(nil)
}
func (x fastReflection_MsgCreateBatch_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch)
}
func (x fastReflection_MsgCreateBatch_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBatch) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBatch) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBatch_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBatch) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBatch) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBatch)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBatch) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Issuer != "" {
		value := protoreflect.ValueOfString(x.Issuer)
		if !f(fd_MsgCreateBatch_issuer, value) {
			return
		}
	}
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_MsgCreateBatch_project_id, value) {
			return
		}
	}
	if len(x.Issuance) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateBatch_3_list{list: &x.Issuance})
		if !f(fd_MsgCreateBatch_issuance, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateBatch_metadata, value) {
			return
		}
	}
	if x.StartDate != nil {
		value := protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
		if !f(fd_MsgCreateBatch_start_date, value) {
			return
		}
	}
	if x.EndDate != nil {
		value := protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
		if !f(fd_MsgCreateBatch_end_date, value) {
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
func (x *fastReflection_MsgCreateBatch) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuer":
		return x.Issuer != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.project_id":
		return x.ProjectId != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuance":
		return len(x.Issuance) != 0
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.start_date":
		return x.StartDate != nil
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.end_date":
		return x.EndDate != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuer":
		x.Issuer = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.project_id":
		x.ProjectId = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuance":
		x.Issuance = nil
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.start_date":
		x.StartDate = nil
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.end_date":
		x.EndDate = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBatch) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuer":
		value := x.Issuer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuance":
		if len(x.Issuance) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateBatch_3_list{})
		}
		listValue := &_MsgCreateBatch_3_list{list: &x.Issuance}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.start_date":
		value := x.StartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.end_date":
		value := x.EndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBatch) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuer":
		x.Issuer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.project_id":
		x.ProjectId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuance":
		lv := value.List()
		clv := lv.(*_MsgCreateBatch_3_list)
		x.Issuance = *clv.list
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.start_date":
		x.StartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.end_date":
		x.EndDate = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBatch) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuance":
		if x.Issuance == nil {
			x.Issuance = []*MsgCreateBatch_BatchIssuance{}
		}
		value := &_MsgCreateBatch_3_list{list: &x.Issuance}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.start_date":
		if x.StartDate == nil {
			x.StartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.end_date":
		if x.EndDate == nil {
			x.EndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1alpha2.MsgCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.MsgCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha2.MsgCreateBatch is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBatch) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.project_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.issuance":
		list := []*MsgCreateBatch_BatchIssuance{}
		return protoreflect.ValueOfList(&_MsgCreateBatch_3_list{list: &list})
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBatch) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateBatch", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBatch) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBatch) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBatch) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBatch)
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
		l = len(x.Issuer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProjectId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Issuance) > 0 {
			for _, e := range x.Issuance {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartDate != nil {
			l = options.Size(x.StartDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.EndDate != nil {
			l = options.Size(x.EndDate)
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
		x := input.Message.Interface().(*MsgCreateBatch)
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
		if x.EndDate != nil {
			encoded, err := options.Marshal(x.EndDate)
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
		if x.StartDate != nil {
			encoded, err := options.Marshal(x.StartDate)
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
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Issuance) > 0 {
			for iNdEx := len(x.Issuance) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Issuance[iNdEx])
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
		}
		if len(x.ProjectId) > 0 {
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Issuer) > 0 {
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
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
		x := input.Message.Interface().(*MsgCreateBatch)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
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
				x.Issuer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
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
				x.ProjectId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuance", wireType)
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
				x.Issuance = append(x.Issuance, &MsgCreateBatch_BatchIssuance{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Issuance[len(x.Issuance)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
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
				if x.StartDate == nil {
					x.StartDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.StartDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
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
				if x.EndDate == nil {
					x.EndDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.EndDate); err != nil {
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
	md_MsgCreateBatch_BatchIssuance                     protoreflect.MessageDescriptor
	fd_MsgCreateBatch_BatchIssuance_recipient           protoreflect.FieldDescriptor
	fd_MsgCreateBatch_BatchIssuance_tradable_amount     protoreflect.FieldDescriptor
	fd_MsgCreateBatch_BatchIssuance_retired_amount      protoreflect.FieldDescriptor
	fd_MsgCreateBatch_BatchIssuance_retirement_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateBatch_BatchIssuance = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateBatch").Messages().ByName("BatchIssuance")
	fd_MsgCreateBatch_BatchIssuance_recipient = md_MsgCreateBatch_BatchIssuance.Fields().ByName("recipient")
	fd_MsgCreateBatch_BatchIssuance_tradable_amount = md_MsgCreateBatch_BatchIssuance.Fields().ByName("tradable_amount")
	fd_MsgCreateBatch_BatchIssuance_retired_amount = md_MsgCreateBatch_BatchIssuance.Fields().ByName("retired_amount")
	fd_MsgCreateBatch_BatchIssuance_retirement_location = md_MsgCreateBatch_BatchIssuance.Fields().ByName("retirement_location")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBatch_BatchIssuance)(nil)

type fastReflection_MsgCreateBatch_BatchIssuance MsgCreateBatch_BatchIssuance

func (x *MsgCreateBatch_BatchIssuance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch_BatchIssuance)(x)
}

func (x *MsgCreateBatch_BatchIssuance) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[34]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBatch_BatchIssuance_messageType fastReflection_MsgCreateBatch_BatchIssuance_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBatch_BatchIssuance_messageType{}

type fastReflection_MsgCreateBatch_BatchIssuance_messageType struct{}

func (x fastReflection_MsgCreateBatch_BatchIssuance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch_BatchIssuance)(nil)
}
func (x fastReflection_MsgCreateBatch_BatchIssuance_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch_BatchIssuance)
}
func (x fastReflection_MsgCreateBatch_BatchIssuance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch_BatchIssuance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch_BatchIssuance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBatch_BatchIssuance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch_BatchIssuance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBatch_BatchIssuance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Recipient != "" {
		value := protoreflect.ValueOfString(x.Recipient)
		if !f(fd_MsgCreateBatch_BatchIssuance_recipient, value) {
			return
		}
	}
	if x.TradableAmount != "" {
		value := protoreflect.ValueOfString(x.TradableAmount)
		if !f(fd_MsgCreateBatch_BatchIssuance_tradable_amount, value) {
			return
		}
	}
	if x.RetiredAmount != "" {
		value := protoreflect.ValueOfString(x.RetiredAmount)
		if !f(fd_MsgCreateBatch_BatchIssuance_retired_amount, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_MsgCreateBatch_BatchIssuance_retirement_location, value) {
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.recipient":
		return x.Recipient != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.tradable_amount":
		return x.TradableAmount != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retired_amount":
		return x.RetiredAmount != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retirement_location":
		return x.RetirementLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.recipient":
		x.Recipient = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.tradable_amount":
		x.TradableAmount = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retired_amount":
		x.RetiredAmount = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retirement_location":
		x.RetirementLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.recipient":
		value := x.Recipient
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.tradable_amount":
		value := x.TradableAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retired_amount":
		value := x.RetiredAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.recipient":
		x.Recipient = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.tradable_amount":
		x.TradableAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retired_amount":
		x.RetiredAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.recipient":
		panic(fmt.Errorf("field recipient of message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.tradable_amount":
		panic(fmt.Errorf("field tradable_amount of message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retired_amount":
		panic(fmt.Errorf("field retired_amount of message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.recipient":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.tradable_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retired_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance.retirement_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBatch_BatchIssuance)
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
		l = len(x.Recipient)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TradableAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetiredAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetirementLocation)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCreateBatch_BatchIssuance)
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
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.RetiredAmount) > 0 {
			i -= len(x.RetiredAmount)
			copy(dAtA[i:], x.RetiredAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetiredAmount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TradableAmount) > 0 {
			i -= len(x.TradableAmount)
			copy(dAtA[i:], x.TradableAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TradableAmount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Recipient) > 0 {
			i -= len(x.Recipient)
			copy(dAtA[i:], x.Recipient)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Recipient)))
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
		x := input.Message.Interface().(*MsgCreateBatch_BatchIssuance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch_BatchIssuance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch_BatchIssuance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
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
				x.Recipient = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TradableAmount", wireType)
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
				x.TradableAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetiredAmount", wireType)
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
				x.RetiredAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
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
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
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
	md_MsgCreateBatchResponse             protoreflect.MessageDescriptor
	fd_MsgCreateBatchResponse_batch_denom protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateBatchResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateBatchResponse")
	fd_MsgCreateBatchResponse_batch_denom = md_MsgCreateBatchResponse.Fields().ByName("batch_denom")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBatchResponse)(nil)

type fastReflection_MsgCreateBatchResponse MsgCreateBatchResponse

func (x *MsgCreateBatchResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBatchResponse)(x)
}

func (x *MsgCreateBatchResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBatchResponse_messageType fastReflection_MsgCreateBatchResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBatchResponse_messageType{}

type fastReflection_MsgCreateBatchResponse_messageType struct{}

func (x fastReflection_MsgCreateBatchResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBatchResponse)(nil)
}
func (x fastReflection_MsgCreateBatchResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatchResponse)
}
func (x fastReflection_MsgCreateBatchResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatchResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBatchResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatchResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBatchResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBatchResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBatchResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatchResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBatchResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBatchResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBatchResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgCreateBatchResponse_batch_denom, value) {
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
func (x *fastReflection_MsgCreateBatchResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatchResponse.batch_denom":
		return x.BatchDenom != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatchResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatchResponse.batch_denom":
		x.BatchDenom = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBatchResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatchResponse.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatchResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBatchResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatchResponse.batch_denom":
		x.BatchDenom = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBatchResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatchResponse.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.MsgCreateBatchResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBatchResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBatchResponse.batch_denom":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBatchResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateBatchResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBatchResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatchResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBatchResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBatchResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBatchResponse)
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
		l = len(x.BatchDenom)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCreateBatchResponse)
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
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
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
		x := input.Message.Interface().(*MsgCreateBatchResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatchResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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

var _ protoreflect.List = (*_MsgSend_3_list)(nil)

type _MsgSend_3_list struct {
	list *[]*MsgSend_SendCredits
}

func (x *_MsgSend_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgSend_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgSend_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSend_SendCredits)
	(*x.list)[i] = concreteValue
}

func (x *_MsgSend_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSend_SendCredits)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgSend_3_list) AppendMutable() protoreflect.Value {
	v := new(MsgSend_SendCredits)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSend_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgSend_3_list) NewElement() protoreflect.Value {
	v := new(MsgSend_SendCredits)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSend_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgSend           protoreflect.MessageDescriptor
	fd_MsgSend_sender    protoreflect.FieldDescriptor
	fd_MsgSend_recipient protoreflect.FieldDescriptor
	fd_MsgSend_credits   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgSend = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgSend")
	fd_MsgSend_sender = md_MsgSend.Fields().ByName("sender")
	fd_MsgSend_recipient = md_MsgSend.Fields().ByName("recipient")
	fd_MsgSend_credits = md_MsgSend.Fields().ByName("credits")
}

var _ protoreflect.Message = (*fastReflection_MsgSend)(nil)

type fastReflection_MsgSend MsgSend

func (x *MsgSend) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSend)(x)
}

func (x *MsgSend) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSend_messageType fastReflection_MsgSend_messageType
var _ protoreflect.MessageType = fastReflection_MsgSend_messageType{}

type fastReflection_MsgSend_messageType struct{}

func (x fastReflection_MsgSend_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSend)(nil)
}
func (x fastReflection_MsgSend_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}
func (x fastReflection_MsgSend_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSend) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSend) Type() protoreflect.MessageType {
	return _fastReflection_MsgSend_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSend) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSend) Interface() protoreflect.ProtoMessage {
	return (*MsgSend)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSend) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sender != "" {
		value := protoreflect.ValueOfString(x.Sender)
		if !f(fd_MsgSend_sender, value) {
			return
		}
	}
	if x.Recipient != "" {
		value := protoreflect.ValueOfString(x.Recipient)
		if !f(fd_MsgSend_recipient, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgSend_3_list{list: &x.Credits})
		if !f(fd_MsgSend_credits, value) {
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
func (x *fastReflection_MsgSend) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.sender":
		return x.Sender != ""
	case "regen.ecocredit.v1alpha2.MsgSend.recipient":
		return x.Recipient != ""
	case "regen.ecocredit.v1alpha2.MsgSend.credits":
		return len(x.Credits) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.sender":
		x.Sender = ""
	case "regen.ecocredit.v1alpha2.MsgSend.recipient":
		x.Recipient = ""
	case "regen.ecocredit.v1alpha2.MsgSend.credits":
		x.Credits = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSend) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.sender":
		value := x.Sender
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSend.recipient":
		value := x.Recipient
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSend.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgSend_3_list{})
		}
		listValue := &_MsgSend_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSend) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.sender":
		x.Sender = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSend.recipient":
		x.Recipient = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSend.credits":
		lv := value.List()
		clv := lv.(*_MsgSend_3_list)
		x.Credits = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSend) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.credits":
		if x.Credits == nil {
			x.Credits = []*MsgSend_SendCredits{}
		}
		value := &_MsgSend_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgSend.sender":
		panic(fmt.Errorf("field sender of message regen.ecocredit.v1alpha2.MsgSend is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgSend.recipient":
		panic(fmt.Errorf("field recipient of message regen.ecocredit.v1alpha2.MsgSend is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSend) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.sender":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSend.recipient":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSend.credits":
		list := []*MsgSend_SendCredits{}
		return protoreflect.ValueOfList(&_MsgSend_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSend) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgSend", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSend) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSend) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSend) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSend)
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
		l = len(x.Sender)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Recipient)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgSend)
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
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
		}
		if len(x.Recipient) > 0 {
			i -= len(x.Recipient)
			copy(dAtA[i:], x.Recipient)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Recipient)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Sender) > 0 {
			i -= len(x.Sender)
			copy(dAtA[i:], x.Sender)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sender)))
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
		x := input.Message.Interface().(*MsgSend)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
				x.Sender = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
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
				x.Recipient = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &MsgSend_SendCredits{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
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
	md_MsgSend_SendCredits                     protoreflect.MessageDescriptor
	fd_MsgSend_SendCredits_batch_denom         protoreflect.FieldDescriptor
	fd_MsgSend_SendCredits_tradable_amount     protoreflect.FieldDescriptor
	fd_MsgSend_SendCredits_retired_amount      protoreflect.FieldDescriptor
	fd_MsgSend_SendCredits_retirement_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgSend_SendCredits = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgSend").Messages().ByName("SendCredits")
	fd_MsgSend_SendCredits_batch_denom = md_MsgSend_SendCredits.Fields().ByName("batch_denom")
	fd_MsgSend_SendCredits_tradable_amount = md_MsgSend_SendCredits.Fields().ByName("tradable_amount")
	fd_MsgSend_SendCredits_retired_amount = md_MsgSend_SendCredits.Fields().ByName("retired_amount")
	fd_MsgSend_SendCredits_retirement_location = md_MsgSend_SendCredits.Fields().ByName("retirement_location")
}

var _ protoreflect.Message = (*fastReflection_MsgSend_SendCredits)(nil)

type fastReflection_MsgSend_SendCredits MsgSend_SendCredits

func (x *MsgSend_SendCredits) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSend_SendCredits)(x)
}

func (x *MsgSend_SendCredits) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[35]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSend_SendCredits_messageType fastReflection_MsgSend_SendCredits_messageType
var _ protoreflect.MessageType = fastReflection_MsgSend_SendCredits_messageType{}

type fastReflection_MsgSend_SendCredits_messageType struct{}

func (x fastReflection_MsgSend_SendCredits_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSend_SendCredits)(nil)
}
func (x fastReflection_MsgSend_SendCredits_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSend_SendCredits)
}
func (x fastReflection_MsgSend_SendCredits_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend_SendCredits
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSend_SendCredits) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend_SendCredits
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSend_SendCredits) Type() protoreflect.MessageType {
	return _fastReflection_MsgSend_SendCredits_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSend_SendCredits) New() protoreflect.Message {
	return new(fastReflection_MsgSend_SendCredits)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSend_SendCredits) Interface() protoreflect.ProtoMessage {
	return (*MsgSend_SendCredits)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSend_SendCredits) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgSend_SendCredits_batch_denom, value) {
			return
		}
	}
	if x.TradableAmount != "" {
		value := protoreflect.ValueOfString(x.TradableAmount)
		if !f(fd_MsgSend_SendCredits_tradable_amount, value) {
			return
		}
	}
	if x.RetiredAmount != "" {
		value := protoreflect.ValueOfString(x.RetiredAmount)
		if !f(fd_MsgSend_SendCredits_retired_amount, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_MsgSend_SendCredits_retirement_location, value) {
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
func (x *fastReflection_MsgSend_SendCredits) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.tradable_amount":
		return x.TradableAmount != ""
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retired_amount":
		return x.RetiredAmount != ""
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retirement_location":
		return x.RetirementLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend_SendCredits) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.tradable_amount":
		x.TradableAmount = ""
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retired_amount":
		x.RetiredAmount = ""
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retirement_location":
		x.RetirementLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSend_SendCredits) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.tradable_amount":
		value := x.TradableAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retired_amount":
		value := x.RetiredAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend.SendCredits does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSend_SendCredits) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.tradable_amount":
		x.TradableAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retired_amount":
		x.RetiredAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend.SendCredits does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSend_SendCredits) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.MsgSend.SendCredits is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.tradable_amount":
		panic(fmt.Errorf("field tradable_amount of message regen.ecocredit.v1alpha2.MsgSend.SendCredits is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retired_amount":
		panic(fmt.Errorf("field retired_amount of message regen.ecocredit.v1alpha2.MsgSend.SendCredits is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha2.MsgSend.SendCredits is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSend_SendCredits) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.tradable_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retired_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSend.SendCredits.retirement_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSend_SendCredits) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgSend.SendCredits", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSend_SendCredits) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend_SendCredits) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSend_SendCredits) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSend_SendCredits) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSend_SendCredits)
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
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TradableAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetiredAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetirementLocation)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgSend_SendCredits)
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
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.RetiredAmount) > 0 {
			i -= len(x.RetiredAmount)
			copy(dAtA[i:], x.RetiredAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetiredAmount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TradableAmount) > 0 {
			i -= len(x.TradableAmount)
			copy(dAtA[i:], x.TradableAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TradableAmount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
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
		x := input.Message.Interface().(*MsgSend_SendCredits)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend_SendCredits: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend_SendCredits: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TradableAmount", wireType)
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
				x.TradableAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetiredAmount", wireType)
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
				x.RetiredAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
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
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
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
	md_MsgSendResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgSendResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgSendResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgSendResponse)(nil)

type fastReflection_MsgSendResponse MsgSendResponse

func (x *MsgSendResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSendResponse)(x)
}

func (x *MsgSendResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSendResponse_messageType fastReflection_MsgSendResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgSendResponse_messageType{}

type fastReflection_MsgSendResponse_messageType struct{}

func (x fastReflection_MsgSendResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSendResponse)(nil)
}
func (x fastReflection_MsgSendResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSendResponse)
}
func (x fastReflection_MsgSendResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSendResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSendResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSendResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSendResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgSendResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSendResponse) New() protoreflect.Message {
	return new(fastReflection_MsgSendResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSendResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgSendResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSendResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgSendResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSendResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSendResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSendResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSendResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSendResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSendResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSendResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSendResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgSendResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSendResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSendResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSendResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSendResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSendResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgSendResponse)
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
		x := input.Message.Interface().(*MsgSendResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSendResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgRetire_2_list)(nil)

type _MsgRetire_2_list struct {
	list *[]*MsgRetire_RetireCredits
}

func (x *_MsgRetire_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgRetire_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgRetire_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgRetire_RetireCredits)
	(*x.list)[i] = concreteValue
}

func (x *_MsgRetire_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgRetire_RetireCredits)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgRetire_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgRetire_RetireCredits)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgRetire_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgRetire_2_list) NewElement() protoreflect.Value {
	v := new(MsgRetire_RetireCredits)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgRetire_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgRetire          protoreflect.MessageDescriptor
	fd_MsgRetire_holder   protoreflect.FieldDescriptor
	fd_MsgRetire_credits  protoreflect.FieldDescriptor
	fd_MsgRetire_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgRetire = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgRetire")
	fd_MsgRetire_holder = md_MsgRetire.Fields().ByName("holder")
	fd_MsgRetire_credits = md_MsgRetire.Fields().ByName("credits")
	fd_MsgRetire_location = md_MsgRetire.Fields().ByName("location")
}

var _ protoreflect.Message = (*fastReflection_MsgRetire)(nil)

type fastReflection_MsgRetire MsgRetire

func (x *MsgRetire) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRetire)(x)
}

func (x *MsgRetire) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRetire_messageType fastReflection_MsgRetire_messageType
var _ protoreflect.MessageType = fastReflection_MsgRetire_messageType{}

type fastReflection_MsgRetire_messageType struct{}

func (x fastReflection_MsgRetire_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRetire)(nil)
}
func (x fastReflection_MsgRetire_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRetire)
}
func (x fastReflection_MsgRetire_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRetire) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRetire) Type() protoreflect.MessageType {
	return _fastReflection_MsgRetire_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRetire) New() protoreflect.Message {
	return new(fastReflection_MsgRetire)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRetire) Interface() protoreflect.ProtoMessage {
	return (*MsgRetire)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRetire) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Holder != "" {
		value := protoreflect.ValueOfString(x.Holder)
		if !f(fd_MsgRetire_holder, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgRetire_2_list{list: &x.Credits})
		if !f(fd_MsgRetire_credits, value) {
			return
		}
	}
	if x.Location != "" {
		value := protoreflect.ValueOfString(x.Location)
		if !f(fd_MsgRetire_location, value) {
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
func (x *fastReflection_MsgRetire) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.holder":
		return x.Holder != ""
	case "regen.ecocredit.v1alpha2.MsgRetire.credits":
		return len(x.Credits) != 0
	case "regen.ecocredit.v1alpha2.MsgRetire.location":
		return x.Location != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.holder":
		x.Holder = ""
	case "regen.ecocredit.v1alpha2.MsgRetire.credits":
		x.Credits = nil
	case "regen.ecocredit.v1alpha2.MsgRetire.location":
		x.Location = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRetire) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.holder":
		value := x.Holder
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgRetire.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgRetire_2_list{})
		}
		listValue := &_MsgRetire_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.MsgRetire.location":
		value := x.Location
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRetire) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.holder":
		x.Holder = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgRetire.credits":
		lv := value.List()
		clv := lv.(*_MsgRetire_2_list)
		x.Credits = *clv.list
	case "regen.ecocredit.v1alpha2.MsgRetire.location":
		x.Location = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRetire) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.credits":
		if x.Credits == nil {
			x.Credits = []*MsgRetire_RetireCredits{}
		}
		value := &_MsgRetire_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgRetire.holder":
		panic(fmt.Errorf("field holder of message regen.ecocredit.v1alpha2.MsgRetire is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgRetire.location":
		panic(fmt.Errorf("field location of message regen.ecocredit.v1alpha2.MsgRetire is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRetire) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.holder":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgRetire.credits":
		list := []*MsgRetire_RetireCredits{}
		return protoreflect.ValueOfList(&_MsgRetire_2_list{list: &list})
	case "regen.ecocredit.v1alpha2.MsgRetire.location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRetire) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgRetire", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRetire) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRetire) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRetire) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRetire)
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
		l = len(x.Holder)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Location)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgRetire)
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
		if len(x.Location) > 0 {
			i -= len(x.Location)
			copy(dAtA[i:], x.Location)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Location)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
				dAtA[i] = 0x12
			}
		}
		if len(x.Holder) > 0 {
			i -= len(x.Holder)
			copy(dAtA[i:], x.Holder)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Holder)))
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
		x := input.Message.Interface().(*MsgRetire)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Holder", wireType)
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
				x.Holder = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &MsgRetire_RetireCredits{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
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
				x.Location = string(dAtA[iNdEx:postIndex])
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
	md_MsgRetire_RetireCredits             protoreflect.MessageDescriptor
	fd_MsgRetire_RetireCredits_batch_denom protoreflect.FieldDescriptor
	fd_MsgRetire_RetireCredits_amount      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgRetire_RetireCredits = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgRetire").Messages().ByName("RetireCredits")
	fd_MsgRetire_RetireCredits_batch_denom = md_MsgRetire_RetireCredits.Fields().ByName("batch_denom")
	fd_MsgRetire_RetireCredits_amount = md_MsgRetire_RetireCredits.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_MsgRetire_RetireCredits)(nil)

type fastReflection_MsgRetire_RetireCredits MsgRetire_RetireCredits

func (x *MsgRetire_RetireCredits) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRetire_RetireCredits)(x)
}

func (x *MsgRetire_RetireCredits) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[36]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRetire_RetireCredits_messageType fastReflection_MsgRetire_RetireCredits_messageType
var _ protoreflect.MessageType = fastReflection_MsgRetire_RetireCredits_messageType{}

type fastReflection_MsgRetire_RetireCredits_messageType struct{}

func (x fastReflection_MsgRetire_RetireCredits_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRetire_RetireCredits)(nil)
}
func (x fastReflection_MsgRetire_RetireCredits_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRetire_RetireCredits)
}
func (x fastReflection_MsgRetire_RetireCredits_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire_RetireCredits
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRetire_RetireCredits) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire_RetireCredits
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRetire_RetireCredits) Type() protoreflect.MessageType {
	return _fastReflection_MsgRetire_RetireCredits_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRetire_RetireCredits) New() protoreflect.Message {
	return new(fastReflection_MsgRetire_RetireCredits)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRetire_RetireCredits) Interface() protoreflect.ProtoMessage {
	return (*MsgRetire_RetireCredits)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRetire_RetireCredits) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgRetire_RetireCredits_batch_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_MsgRetire_RetireCredits_amount, value) {
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
func (x *fastReflection_MsgRetire_RetireCredits) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire_RetireCredits) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRetire_RetireCredits) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRetire_RetireCredits) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.amount":
		x.Amount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRetire_RetireCredits) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.amount":
		panic(fmt.Errorf("field amount of message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRetire_RetireCredits) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgRetire.RetireCredits.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRetire_RetireCredits) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgRetire.RetireCredits", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRetire_RetireCredits) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire_RetireCredits) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRetire_RetireCredits) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRetire_RetireCredits) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRetire_RetireCredits)
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
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgRetire_RetireCredits)
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
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
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
		x := input.Message.Interface().(*MsgRetire_RetireCredits)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire_RetireCredits: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire_RetireCredits: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = string(dAtA[iNdEx:postIndex])
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
	md_MsgRetireResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgRetireResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgRetireResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgRetireResponse)(nil)

type fastReflection_MsgRetireResponse MsgRetireResponse

func (x *MsgRetireResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRetireResponse)(x)
}

func (x *MsgRetireResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRetireResponse_messageType fastReflection_MsgRetireResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgRetireResponse_messageType{}

type fastReflection_MsgRetireResponse_messageType struct{}

func (x fastReflection_MsgRetireResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRetireResponse)(nil)
}
func (x fastReflection_MsgRetireResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRetireResponse)
}
func (x fastReflection_MsgRetireResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetireResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRetireResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetireResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRetireResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgRetireResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRetireResponse) New() protoreflect.Message {
	return new(fastReflection_MsgRetireResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRetireResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgRetireResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRetireResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgRetireResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetireResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRetireResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetireResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRetireResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetireResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRetireResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRetireResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRetireResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgRetireResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRetireResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetireResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRetireResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRetireResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRetireResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgRetireResponse)
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
		x := input.Message.Interface().(*MsgRetireResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetireResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetireResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgCancel_2_list)(nil)

type _MsgCancel_2_list struct {
	list *[]*MsgCancel_CancelCredits
}

func (x *_MsgCancel_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCancel_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCancel_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCancel_CancelCredits)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCancel_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCancel_CancelCredits)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCancel_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgCancel_CancelCredits)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCancel_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCancel_2_list) NewElement() protoreflect.Value {
	v := new(MsgCancel_CancelCredits)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCancel_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCancel         protoreflect.MessageDescriptor
	fd_MsgCancel_holder  protoreflect.FieldDescriptor
	fd_MsgCancel_credits protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCancel = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCancel")
	fd_MsgCancel_holder = md_MsgCancel.Fields().ByName("holder")
	fd_MsgCancel_credits = md_MsgCancel.Fields().ByName("credits")
}

var _ protoreflect.Message = (*fastReflection_MsgCancel)(nil)

type fastReflection_MsgCancel MsgCancel

func (x *MsgCancel) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCancel)(x)
}

func (x *MsgCancel) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCancel_messageType fastReflection_MsgCancel_messageType
var _ protoreflect.MessageType = fastReflection_MsgCancel_messageType{}

type fastReflection_MsgCancel_messageType struct{}

func (x fastReflection_MsgCancel_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCancel)(nil)
}
func (x fastReflection_MsgCancel_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCancel)
}
func (x fastReflection_MsgCancel_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCancel) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCancel) Type() protoreflect.MessageType {
	return _fastReflection_MsgCancel_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCancel) New() protoreflect.Message {
	return new(fastReflection_MsgCancel)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCancel) Interface() protoreflect.ProtoMessage {
	return (*MsgCancel)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCancel) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Holder != "" {
		value := protoreflect.ValueOfString(x.Holder)
		if !f(fd_MsgCancel_holder, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgCancel_2_list{list: &x.Credits})
		if !f(fd_MsgCancel_credits, value) {
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
func (x *fastReflection_MsgCancel) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.holder":
		return x.Holder != ""
	case "regen.ecocredit.v1alpha2.MsgCancel.credits":
		return len(x.Credits) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.holder":
		x.Holder = ""
	case "regen.ecocredit.v1alpha2.MsgCancel.credits":
		x.Credits = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCancel) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.holder":
		value := x.Holder
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCancel.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgCancel_2_list{})
		}
		listValue := &_MsgCancel_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCancel) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.holder":
		x.Holder = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCancel.credits":
		lv := value.List()
		clv := lv.(*_MsgCancel_2_list)
		x.Credits = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCancel) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.credits":
		if x.Credits == nil {
			x.Credits = []*MsgCancel_CancelCredits{}
		}
		value := &_MsgCancel_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgCancel.holder":
		panic(fmt.Errorf("field holder of message regen.ecocredit.v1alpha2.MsgCancel is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCancel) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.holder":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCancel.credits":
		list := []*MsgCancel_CancelCredits{}
		return protoreflect.ValueOfList(&_MsgCancel_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCancel) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCancel", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCancel) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCancel) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCancel) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCancel)
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
		l = len(x.Holder)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgCancel)
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
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
				dAtA[i] = 0x12
			}
		}
		if len(x.Holder) > 0 {
			i -= len(x.Holder)
			copy(dAtA[i:], x.Holder)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Holder)))
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
		x := input.Message.Interface().(*MsgCancel)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Holder", wireType)
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
				x.Holder = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &MsgCancel_CancelCredits{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
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
	md_MsgCancel_CancelCredits             protoreflect.MessageDescriptor
	fd_MsgCancel_CancelCredits_batch_denom protoreflect.FieldDescriptor
	fd_MsgCancel_CancelCredits_amount      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCancel_CancelCredits = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCancel").Messages().ByName("CancelCredits")
	fd_MsgCancel_CancelCredits_batch_denom = md_MsgCancel_CancelCredits.Fields().ByName("batch_denom")
	fd_MsgCancel_CancelCredits_amount = md_MsgCancel_CancelCredits.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_MsgCancel_CancelCredits)(nil)

type fastReflection_MsgCancel_CancelCredits MsgCancel_CancelCredits

func (x *MsgCancel_CancelCredits) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCancel_CancelCredits)(x)
}

func (x *MsgCancel_CancelCredits) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[37]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCancel_CancelCredits_messageType fastReflection_MsgCancel_CancelCredits_messageType
var _ protoreflect.MessageType = fastReflection_MsgCancel_CancelCredits_messageType{}

type fastReflection_MsgCancel_CancelCredits_messageType struct{}

func (x fastReflection_MsgCancel_CancelCredits_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCancel_CancelCredits)(nil)
}
func (x fastReflection_MsgCancel_CancelCredits_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCancel_CancelCredits)
}
func (x fastReflection_MsgCancel_CancelCredits_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel_CancelCredits
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCancel_CancelCredits) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel_CancelCredits
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCancel_CancelCredits) Type() protoreflect.MessageType {
	return _fastReflection_MsgCancel_CancelCredits_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCancel_CancelCredits) New() protoreflect.Message {
	return new(fastReflection_MsgCancel_CancelCredits)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCancel_CancelCredits) Interface() protoreflect.ProtoMessage {
	return (*MsgCancel_CancelCredits)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCancel_CancelCredits) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgCancel_CancelCredits_batch_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_MsgCancel_CancelCredits_amount, value) {
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
func (x *fastReflection_MsgCancel_CancelCredits) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel_CancelCredits) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCancel_CancelCredits) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCancel_CancelCredits) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.amount":
		x.Amount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCancel_CancelCredits) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.amount":
		panic(fmt.Errorf("field amount of message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCancel_CancelCredits) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCancel.CancelCredits.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCancel_CancelCredits) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCancel.CancelCredits", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCancel_CancelCredits) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel_CancelCredits) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCancel_CancelCredits) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCancel_CancelCredits) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCancel_CancelCredits)
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
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCancel_CancelCredits)
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
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
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
		x := input.Message.Interface().(*MsgCancel_CancelCredits)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel_CancelCredits: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel_CancelCredits: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = string(dAtA[iNdEx:postIndex])
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
	md_MsgCancelResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCancelResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCancelResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgCancelResponse)(nil)

type fastReflection_MsgCancelResponse MsgCancelResponse

func (x *MsgCancelResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCancelResponse)(x)
}

func (x *MsgCancelResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCancelResponse_messageType fastReflection_MsgCancelResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCancelResponse_messageType{}

type fastReflection_MsgCancelResponse_messageType struct{}

func (x fastReflection_MsgCancelResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCancelResponse)(nil)
}
func (x fastReflection_MsgCancelResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCancelResponse)
}
func (x fastReflection_MsgCancelResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancelResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCancelResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancelResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCancelResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCancelResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCancelResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCancelResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCancelResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCancelResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCancelResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgCancelResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancelResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCancelResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancelResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCancelResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancelResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCancelResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCancelResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCancelResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCancelResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCancelResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancelResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCancelResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCancelResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCancelResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCancelResponse)
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
		x := input.Message.Interface().(*MsgCancelResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancelResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_MsgUpdateClassAdmin           protoreflect.MessageDescriptor
	fd_MsgUpdateClassAdmin_admin     protoreflect.FieldDescriptor
	fd_MsgUpdateClassAdmin_class_id  protoreflect.FieldDescriptor
	fd_MsgUpdateClassAdmin_new_admin protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateClassAdmin = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateClassAdmin")
	fd_MsgUpdateClassAdmin_admin = md_MsgUpdateClassAdmin.Fields().ByName("admin")
	fd_MsgUpdateClassAdmin_class_id = md_MsgUpdateClassAdmin.Fields().ByName("class_id")
	fd_MsgUpdateClassAdmin_new_admin = md_MsgUpdateClassAdmin.Fields().ByName("new_admin")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassAdmin)(nil)

type fastReflection_MsgUpdateClassAdmin MsgUpdateClassAdmin

func (x *MsgUpdateClassAdmin) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdmin)(x)
}

func (x *MsgUpdateClassAdmin) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassAdmin_messageType fastReflection_MsgUpdateClassAdmin_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassAdmin_messageType{}

type fastReflection_MsgUpdateClassAdmin_messageType struct{}

func (x fastReflection_MsgUpdateClassAdmin_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdmin)(nil)
}
func (x fastReflection_MsgUpdateClassAdmin_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdmin)
}
func (x fastReflection_MsgUpdateClassAdmin_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdmin
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassAdmin) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdmin
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassAdmin) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassAdmin_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassAdmin) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdmin)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassAdmin) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassAdmin)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassAdmin) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateClassAdmin_admin, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgUpdateClassAdmin_class_id, value) {
			return
		}
	}
	if x.NewAdmin != "" {
		value := protoreflect.ValueOfString(x.NewAdmin)
		if !f(fd_MsgUpdateClassAdmin_new_admin, value) {
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
func (x *fastReflection_MsgUpdateClassAdmin) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.new_admin":
		return x.NewAdmin != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdmin) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.new_admin":
		x.NewAdmin = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassAdmin) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.new_admin":
		value := x.NewAdmin
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdmin) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.new_admin":
		x.NewAdmin = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdmin) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.new_admin":
		panic(fmt.Errorf("field new_admin of message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassAdmin) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateClassAdmin.new_admin":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassAdmin) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateClassAdmin", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassAdmin) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdmin) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassAdmin) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassAdmin) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassAdmin)
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
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NewAdmin)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgUpdateClassAdmin)
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
		if len(x.NewAdmin) > 0 {
			i -= len(x.NewAdmin)
			copy(dAtA[i:], x.NewAdmin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewAdmin)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgUpdateClassAdmin)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdmin: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
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
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
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
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
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
				x.NewAdmin = string(dAtA[iNdEx:postIndex])
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
	md_MsgUpdateClassAdminResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateClassAdminResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateClassAdminResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassAdminResponse)(nil)

type fastReflection_MsgUpdateClassAdminResponse MsgUpdateClassAdminResponse

func (x *MsgUpdateClassAdminResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdminResponse)(x)
}

func (x *MsgUpdateClassAdminResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassAdminResponse_messageType fastReflection_MsgUpdateClassAdminResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassAdminResponse_messageType{}

type fastReflection_MsgUpdateClassAdminResponse_messageType struct{}

func (x fastReflection_MsgUpdateClassAdminResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdminResponse)(nil)
}
func (x fastReflection_MsgUpdateClassAdminResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdminResponse)
}
func (x fastReflection_MsgUpdateClassAdminResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdminResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassAdminResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdminResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassAdminResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassAdminResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassAdminResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdminResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassAdminResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassAdminResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassAdminResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgUpdateClassAdminResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdminResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassAdminResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdminResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdminResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassAdminResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassAdminResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassAdminResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdminResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassAdminResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassAdminResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassAdminResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateClassAdminResponse)
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
		x := input.Message.Interface().(*MsgUpdateClassAdminResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdminResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgUpdateClassIssuers_3_list)(nil)

type _MsgUpdateClassIssuers_3_list struct {
	list *[]string
}

func (x *_MsgUpdateClassIssuers_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgUpdateClassIssuers_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_MsgUpdateClassIssuers_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgUpdateClassIssuers_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgUpdateClassIssuers_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgUpdateClassIssuers at list field Issuers as it is not of Message kind"))
}

func (x *_MsgUpdateClassIssuers_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgUpdateClassIssuers_3_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_MsgUpdateClassIssuers_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgUpdateClassIssuers          protoreflect.MessageDescriptor
	fd_MsgUpdateClassIssuers_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateClassIssuers_class_id protoreflect.FieldDescriptor
	fd_MsgUpdateClassIssuers_issuers  protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateClassIssuers = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateClassIssuers")
	fd_MsgUpdateClassIssuers_admin = md_MsgUpdateClassIssuers.Fields().ByName("admin")
	fd_MsgUpdateClassIssuers_class_id = md_MsgUpdateClassIssuers.Fields().ByName("class_id")
	fd_MsgUpdateClassIssuers_issuers = md_MsgUpdateClassIssuers.Fields().ByName("issuers")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassIssuers)(nil)

type fastReflection_MsgUpdateClassIssuers MsgUpdateClassIssuers

func (x *MsgUpdateClassIssuers) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuers)(x)
}

func (x *MsgUpdateClassIssuers) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassIssuers_messageType fastReflection_MsgUpdateClassIssuers_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassIssuers_messageType{}

type fastReflection_MsgUpdateClassIssuers_messageType struct{}

func (x fastReflection_MsgUpdateClassIssuers_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuers)(nil)
}
func (x fastReflection_MsgUpdateClassIssuers_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuers)
}
func (x fastReflection_MsgUpdateClassIssuers_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuers
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassIssuers) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuers
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassIssuers) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassIssuers_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassIssuers) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuers)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassIssuers) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassIssuers)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassIssuers) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateClassIssuers_admin, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgUpdateClassIssuers_class_id, value) {
			return
		}
	}
	if len(x.Issuers) != 0 {
		value := protoreflect.ValueOfList(&_MsgUpdateClassIssuers_3_list{list: &x.Issuers})
		if !f(fd_MsgUpdateClassIssuers_issuers, value) {
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
func (x *fastReflection_MsgUpdateClassIssuers) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.issuers":
		return len(x.Issuers) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuers) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.issuers":
		x.Issuers = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassIssuers) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.issuers":
		if len(x.Issuers) == 0 {
			return protoreflect.ValueOfList(&_MsgUpdateClassIssuers_3_list{})
		}
		listValue := &_MsgUpdateClassIssuers_3_list{list: &x.Issuers}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuers) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.issuers":
		lv := value.List()
		clv := lv.(*_MsgUpdateClassIssuers_3_list)
		x.Issuers = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuers) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.issuers":
		if x.Issuers == nil {
			x.Issuers = []string{}
		}
		value := &_MsgUpdateClassIssuers_3_list{list: &x.Issuers}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassIssuers) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateClassIssuers.issuers":
		list := []string{}
		return protoreflect.ValueOfList(&_MsgUpdateClassIssuers_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassIssuers) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateClassIssuers", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassIssuers) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuers) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassIssuers) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassIssuers) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassIssuers)
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
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Issuers) > 0 {
			for _, s := range x.Issuers {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgUpdateClassIssuers)
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
		if len(x.Issuers) > 0 {
			for iNdEx := len(x.Issuers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Issuers[iNdEx])
				copy(dAtA[i:], x.Issuers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuers[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgUpdateClassIssuers)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuers: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuers: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
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
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
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
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuers", wireType)
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
				x.Issuers = append(x.Issuers, string(dAtA[iNdEx:postIndex]))
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
	md_MsgUpdateClassIssuersResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateClassIssuersResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateClassIssuersResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassIssuersResponse)(nil)

type fastReflection_MsgUpdateClassIssuersResponse MsgUpdateClassIssuersResponse

func (x *MsgUpdateClassIssuersResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuersResponse)(x)
}

func (x *MsgUpdateClassIssuersResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassIssuersResponse_messageType fastReflection_MsgUpdateClassIssuersResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassIssuersResponse_messageType{}

type fastReflection_MsgUpdateClassIssuersResponse_messageType struct{}

func (x fastReflection_MsgUpdateClassIssuersResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuersResponse)(nil)
}
func (x fastReflection_MsgUpdateClassIssuersResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuersResponse)
}
func (x fastReflection_MsgUpdateClassIssuersResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuersResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuersResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassIssuersResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassIssuersResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuersResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassIssuersResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassIssuersResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassIssuersResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassIssuersResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuersResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassIssuersResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassIssuersResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateClassIssuersResponse)
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
		x := input.Message.Interface().(*MsgUpdateClassIssuersResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuersResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_MsgUpdateClassMetadata          protoreflect.MessageDescriptor
	fd_MsgUpdateClassMetadata_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateClassMetadata_class_id protoreflect.FieldDescriptor
	fd_MsgUpdateClassMetadata_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateClassMetadata = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateClassMetadata")
	fd_MsgUpdateClassMetadata_admin = md_MsgUpdateClassMetadata.Fields().ByName("admin")
	fd_MsgUpdateClassMetadata_class_id = md_MsgUpdateClassMetadata.Fields().ByName("class_id")
	fd_MsgUpdateClassMetadata_metadata = md_MsgUpdateClassMetadata.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassMetadata)(nil)

type fastReflection_MsgUpdateClassMetadata MsgUpdateClassMetadata

func (x *MsgUpdateClassMetadata) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadata)(x)
}

func (x *MsgUpdateClassMetadata) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassMetadata_messageType fastReflection_MsgUpdateClassMetadata_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassMetadata_messageType{}

type fastReflection_MsgUpdateClassMetadata_messageType struct{}

func (x fastReflection_MsgUpdateClassMetadata_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadata)(nil)
}
func (x fastReflection_MsgUpdateClassMetadata_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadata)
}
func (x fastReflection_MsgUpdateClassMetadata_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadata
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassMetadata) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadata
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassMetadata) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassMetadata_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassMetadata) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadata)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassMetadata) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassMetadata)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassMetadata) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateClassMetadata_admin, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgUpdateClassMetadata_class_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgUpdateClassMetadata_metadata, value) {
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
func (x *fastReflection_MsgUpdateClassMetadata) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadata) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassMetadata) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadata) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadata) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassMetadata) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateClassMetadata.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassMetadata) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateClassMetadata", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassMetadata) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadata) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassMetadata) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassMetadata) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassMetadata)
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
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgUpdateClassMetadata)
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
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgUpdateClassMetadata)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadata: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
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
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
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
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
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
	md_MsgUpdateClassMetadataResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateClassMetadataResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateClassMetadataResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassMetadataResponse)(nil)

type fastReflection_MsgUpdateClassMetadataResponse MsgUpdateClassMetadataResponse

func (x *MsgUpdateClassMetadataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadataResponse)(x)
}

func (x *MsgUpdateClassMetadataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassMetadataResponse_messageType fastReflection_MsgUpdateClassMetadataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassMetadataResponse_messageType{}

type fastReflection_MsgUpdateClassMetadataResponse_messageType struct{}

func (x fastReflection_MsgUpdateClassMetadataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadataResponse)(nil)
}
func (x fastReflection_MsgUpdateClassMetadataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadataResponse)
}
func (x fastReflection_MsgUpdateClassMetadataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassMetadataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassMetadataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassMetadataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassMetadataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassMetadataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassMetadataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadataResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassMetadataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassMetadataResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateClassMetadataResponse)
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
		x := input.Message.Interface().(*MsgUpdateClassMetadataResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgSell_2_list)(nil)

type _MsgSell_2_list struct {
	list *[]*MsgSell_Order
}

func (x *_MsgSell_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgSell_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgSell_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSell_Order)
	(*x.list)[i] = concreteValue
}

func (x *_MsgSell_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSell_Order)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgSell_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgSell_Order)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSell_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgSell_2_list) NewElement() protoreflect.Value {
	v := new(MsgSell_Order)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSell_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgSell        protoreflect.MessageDescriptor
	fd_MsgSell_owner  protoreflect.FieldDescriptor
	fd_MsgSell_orders protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgSell = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgSell")
	fd_MsgSell_owner = md_MsgSell.Fields().ByName("owner")
	fd_MsgSell_orders = md_MsgSell.Fields().ByName("orders")
}

var _ protoreflect.Message = (*fastReflection_MsgSell)(nil)

type fastReflection_MsgSell MsgSell

func (x *MsgSell) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSell)(x)
}

func (x *MsgSell) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSell_messageType fastReflection_MsgSell_messageType
var _ protoreflect.MessageType = fastReflection_MsgSell_messageType{}

type fastReflection_MsgSell_messageType struct{}

func (x fastReflection_MsgSell_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSell)(nil)
}
func (x fastReflection_MsgSell_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSell)
}
func (x fastReflection_MsgSell_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSell
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSell) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSell
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSell) Type() protoreflect.MessageType {
	return _fastReflection_MsgSell_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSell) New() protoreflect.Message {
	return new(fastReflection_MsgSell)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSell) Interface() protoreflect.ProtoMessage {
	return (*MsgSell)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSell) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_MsgSell_owner, value) {
			return
		}
	}
	if len(x.Orders) != 0 {
		value := protoreflect.ValueOfList(&_MsgSell_2_list{list: &x.Orders})
		if !f(fd_MsgSell_orders, value) {
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
func (x *fastReflection_MsgSell) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1alpha2.MsgSell.orders":
		return len(x.Orders) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSell) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.owner":
		x.Owner = ""
	case "regen.ecocredit.v1alpha2.MsgSell.orders":
		x.Orders = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSell) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSell.orders":
		if len(x.Orders) == 0 {
			return protoreflect.ValueOfList(&_MsgSell_2_list{})
		}
		listValue := &_MsgSell_2_list{list: &x.Orders}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSell) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSell.orders":
		lv := value.List()
		clv := lv.(*_MsgSell_2_list)
		x.Orders = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSell) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.orders":
		if x.Orders == nil {
			x.Orders = []*MsgSell_Order{}
		}
		value := &_MsgSell_2_list{list: &x.Orders}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgSell.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.MsgSell is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSell) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSell.orders":
		list := []*MsgSell_Order{}
		return protoreflect.ValueOfList(&_MsgSell_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSell) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgSell", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSell) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSell) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSell) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSell) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSell)
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
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Orders) > 0 {
			for _, e := range x.Orders {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgSell)
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
		if len(x.Orders) > 0 {
			for iNdEx := len(x.Orders) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Orders[iNdEx])
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
				dAtA[i] = 0x12
			}
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
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
		x := input.Message.Interface().(*MsgSell)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSell: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSell: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Orders", wireType)
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
				x.Orders = append(x.Orders, &MsgSell_Order{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Orders[len(x.Orders)-1]); err != nil {
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
	md_MsgSell_Order                     protoreflect.MessageDescriptor
	fd_MsgSell_Order_batch_denom         protoreflect.FieldDescriptor
	fd_MsgSell_Order_quantity            protoreflect.FieldDescriptor
	fd_MsgSell_Order_ask_price           protoreflect.FieldDescriptor
	fd_MsgSell_Order_disable_auto_retire protoreflect.FieldDescriptor
	fd_MsgSell_Order_expiration          protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgSell_Order = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgSell").Messages().ByName("Order")
	fd_MsgSell_Order_batch_denom = md_MsgSell_Order.Fields().ByName("batch_denom")
	fd_MsgSell_Order_quantity = md_MsgSell_Order.Fields().ByName("quantity")
	fd_MsgSell_Order_ask_price = md_MsgSell_Order.Fields().ByName("ask_price")
	fd_MsgSell_Order_disable_auto_retire = md_MsgSell_Order.Fields().ByName("disable_auto_retire")
	fd_MsgSell_Order_expiration = md_MsgSell_Order.Fields().ByName("expiration")
}

var _ protoreflect.Message = (*fastReflection_MsgSell_Order)(nil)

type fastReflection_MsgSell_Order MsgSell_Order

func (x *MsgSell_Order) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSell_Order)(x)
}

func (x *MsgSell_Order) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[38]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSell_Order_messageType fastReflection_MsgSell_Order_messageType
var _ protoreflect.MessageType = fastReflection_MsgSell_Order_messageType{}

type fastReflection_MsgSell_Order_messageType struct{}

func (x fastReflection_MsgSell_Order_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSell_Order)(nil)
}
func (x fastReflection_MsgSell_Order_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSell_Order)
}
func (x fastReflection_MsgSell_Order_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSell_Order
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSell_Order) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSell_Order
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSell_Order) Type() protoreflect.MessageType {
	return _fastReflection_MsgSell_Order_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSell_Order) New() protoreflect.Message {
	return new(fastReflection_MsgSell_Order)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSell_Order) Interface() protoreflect.ProtoMessage {
	return (*MsgSell_Order)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSell_Order) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgSell_Order_batch_denom, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_MsgSell_Order_quantity, value) {
			return
		}
	}
	if x.AskPrice != nil {
		value := protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
		if !f(fd_MsgSell_Order_ask_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_MsgSell_Order_disable_auto_retire, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_MsgSell_Order_expiration, value) {
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
func (x *fastReflection_MsgSell_Order) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.Order.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.MsgSell.Order.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1alpha2.MsgSell.Order.ask_price":
		return x.AskPrice != nil
	case "regen.ecocredit.v1alpha2.MsgSell.Order.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.MsgSell.Order.expiration":
		return x.Expiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell.Order does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSell_Order) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.Order.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.MsgSell.Order.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1alpha2.MsgSell.Order.ask_price":
		x.AskPrice = nil
	case "regen.ecocredit.v1alpha2.MsgSell.Order.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.MsgSell.Order.expiration":
		x.Expiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell.Order does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSell_Order) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.Order.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSell.Order.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgSell.Order.ask_price":
		value := x.AskPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgSell.Order.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.MsgSell.Order.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell.Order does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSell_Order) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.Order.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSell.Order.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgSell.Order.ask_price":
		x.AskPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.MsgSell.Order.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.MsgSell.Order.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell.Order does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSell_Order) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.Order.ask_price":
		if x.AskPrice == nil {
			x.AskPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgSell.Order.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgSell.Order.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.MsgSell.Order is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgSell.Order.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1alpha2.MsgSell.Order is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgSell.Order.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.MsgSell.Order is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell.Order does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSell_Order) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSell.Order.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSell.Order.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgSell.Order.ask_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgSell.Order.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.MsgSell.Order.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSell.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSell.Order does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSell_Order) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgSell.Order", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSell_Order) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSell_Order) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSell_Order) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSell_Order) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSell_Order)
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
		x := input.Message.Interface().(*MsgSell_Order)
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
			dAtA[i] = 0x2a
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x20
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
			dAtA[i] = 0x1a
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
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
		x := input.Message.Interface().(*MsgSell_Order)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSell_Order: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSell_Order: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
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
			case 3:
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
			case 4:
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
			case 5:
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

var _ protoreflect.List = (*_MsgSellResponse_1_list)(nil)

type _MsgSellResponse_1_list struct {
	list *[]uint64
}

func (x *_MsgSellResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgSellResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfUint64((*x.list)[i])
}

func (x *_MsgSellResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgSellResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgSellResponse_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgSellResponse at list field SellOrderIds as it is not of Message kind"))
}

func (x *_MsgSellResponse_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgSellResponse_1_list) NewElement() protoreflect.Value {
	v := uint64(0)
	return protoreflect.ValueOfUint64(v)
}

func (x *_MsgSellResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgSellResponse                protoreflect.MessageDescriptor
	fd_MsgSellResponse_sell_order_ids protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgSellResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgSellResponse")
	fd_MsgSellResponse_sell_order_ids = md_MsgSellResponse.Fields().ByName("sell_order_ids")
}

var _ protoreflect.Message = (*fastReflection_MsgSellResponse)(nil)

type fastReflection_MsgSellResponse MsgSellResponse

func (x *MsgSellResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSellResponse)(x)
}

func (x *MsgSellResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSellResponse_messageType fastReflection_MsgSellResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgSellResponse_messageType{}

type fastReflection_MsgSellResponse_messageType struct{}

func (x fastReflection_MsgSellResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSellResponse)(nil)
}
func (x fastReflection_MsgSellResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSellResponse)
}
func (x fastReflection_MsgSellResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSellResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSellResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSellResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSellResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgSellResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSellResponse) New() protoreflect.Message {
	return new(fastReflection_MsgSellResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSellResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgSellResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSellResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.SellOrderIds) != 0 {
		value := protoreflect.ValueOfList(&_MsgSellResponse_1_list{list: &x.SellOrderIds})
		if !f(fd_MsgSellResponse_sell_order_ids, value) {
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
func (x *fastReflection_MsgSellResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSellResponse.sell_order_ids":
		return len(x.SellOrderIds) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSellResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSellResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSellResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSellResponse.sell_order_ids":
		x.SellOrderIds = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSellResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSellResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSellResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSellResponse.sell_order_ids":
		if len(x.SellOrderIds) == 0 {
			return protoreflect.ValueOfList(&_MsgSellResponse_1_list{})
		}
		listValue := &_MsgSellResponse_1_list{list: &x.SellOrderIds}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSellResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSellResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSellResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSellResponse.sell_order_ids":
		lv := value.List()
		clv := lv.(*_MsgSellResponse_1_list)
		x.SellOrderIds = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSellResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSellResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSellResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSellResponse.sell_order_ids":
		if x.SellOrderIds == nil {
			x.SellOrderIds = []uint64{}
		}
		value := &_MsgSellResponse_1_list{list: &x.SellOrderIds}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSellResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSellResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSellResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgSellResponse.sell_order_ids":
		list := []uint64{}
		return protoreflect.ValueOfList(&_MsgSellResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgSellResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgSellResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSellResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgSellResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSellResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSellResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSellResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSellResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSellResponse)
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
		if len(x.SellOrderIds) > 0 {
			l = 0
			for _, e := range x.SellOrderIds {
				l += runtime.Sov(uint64(e))
			}
			n += 1 + runtime.Sov(uint64(l)) + l
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
		x := input.Message.Interface().(*MsgSellResponse)
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
		if len(x.SellOrderIds) > 0 {
			var pksize2 int
			for _, num := range x.SellOrderIds {
				pksize2 += runtime.Sov(uint64(num))
			}
			i -= pksize2
			j1 := i
			for _, num := range x.SellOrderIds {
				for num >= 1<<7 {
					dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j1++
				}
				dAtA[j1] = uint8(num)
				j1++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize2))
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
		x := input.Message.Interface().(*MsgSellResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSellResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSellResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType == 0 {
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
					x.SellOrderIds = append(x.SellOrderIds, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.SellOrderIds) == 0 {
						x.SellOrderIds = make([]uint64, 0, elementCount)
					}
					for iNdEx < postIndex {
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
						x.SellOrderIds = append(x.SellOrderIds, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderIds", wireType)
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

var _ protoreflect.List = (*_MsgUpdateSellOrders_2_list)(nil)

type _MsgUpdateSellOrders_2_list struct {
	list *[]*MsgUpdateSellOrders_Update
}

func (x *_MsgUpdateSellOrders_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgUpdateSellOrders_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgUpdateSellOrders_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgUpdateSellOrders_Update)
	(*x.list)[i] = concreteValue
}

func (x *_MsgUpdateSellOrders_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgUpdateSellOrders_Update)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgUpdateSellOrders_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgUpdateSellOrders_Update)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgUpdateSellOrders_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgUpdateSellOrders_2_list) NewElement() protoreflect.Value {
	v := new(MsgUpdateSellOrders_Update)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgUpdateSellOrders_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgUpdateSellOrders         protoreflect.MessageDescriptor
	fd_MsgUpdateSellOrders_owner   protoreflect.FieldDescriptor
	fd_MsgUpdateSellOrders_updates protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateSellOrders = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateSellOrders")
	fd_MsgUpdateSellOrders_owner = md_MsgUpdateSellOrders.Fields().ByName("owner")
	fd_MsgUpdateSellOrders_updates = md_MsgUpdateSellOrders.Fields().ByName("updates")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateSellOrders)(nil)

type fastReflection_MsgUpdateSellOrders MsgUpdateSellOrders

func (x *MsgUpdateSellOrders) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateSellOrders)(x)
}

func (x *MsgUpdateSellOrders) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateSellOrders_messageType fastReflection_MsgUpdateSellOrders_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateSellOrders_messageType{}

type fastReflection_MsgUpdateSellOrders_messageType struct{}

func (x fastReflection_MsgUpdateSellOrders_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateSellOrders)(nil)
}
func (x fastReflection_MsgUpdateSellOrders_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateSellOrders)
}
func (x fastReflection_MsgUpdateSellOrders_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateSellOrders
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateSellOrders) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateSellOrders
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateSellOrders) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateSellOrders_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateSellOrders) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateSellOrders)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateSellOrders) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateSellOrders)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateSellOrders) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_MsgUpdateSellOrders_owner, value) {
			return
		}
	}
	if len(x.Updates) != 0 {
		value := protoreflect.ValueOfList(&_MsgUpdateSellOrders_2_list{list: &x.Updates})
		if !f(fd_MsgUpdateSellOrders_updates, value) {
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
func (x *fastReflection_MsgUpdateSellOrders) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.updates":
		return len(x.Updates) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateSellOrders) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.owner":
		x.Owner = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.updates":
		x.Updates = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateSellOrders) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.updates":
		if len(x.Updates) == 0 {
			return protoreflect.ValueOfList(&_MsgUpdateSellOrders_2_list{})
		}
		listValue := &_MsgUpdateSellOrders_2_list{list: &x.Updates}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateSellOrders) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.updates":
		lv := value.List()
		clv := lv.(*_MsgUpdateSellOrders_2_list)
		x.Updates = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateSellOrders) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.updates":
		if x.Updates == nil {
			x.Updates = []*MsgUpdateSellOrders_Update{}
		}
		value := &_MsgUpdateSellOrders_2_list{list: &x.Updates}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.MsgUpdateSellOrders is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateSellOrders) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.updates":
		list := []*MsgUpdateSellOrders_Update{}
		return protoreflect.ValueOfList(&_MsgUpdateSellOrders_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateSellOrders) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateSellOrders", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateSellOrders) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateSellOrders) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateSellOrders) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateSellOrders) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateSellOrders)
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
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Updates) > 0 {
			for _, e := range x.Updates {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgUpdateSellOrders)
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
		if len(x.Updates) > 0 {
			for iNdEx := len(x.Updates) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Updates[iNdEx])
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
				dAtA[i] = 0x12
			}
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
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
		x := input.Message.Interface().(*MsgUpdateSellOrders)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateSellOrders: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateSellOrders: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
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
				x.Updates = append(x.Updates, &MsgUpdateSellOrders_Update{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Updates[len(x.Updates)-1]); err != nil {
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
	md_MsgUpdateSellOrders_Update                     protoreflect.MessageDescriptor
	fd_MsgUpdateSellOrders_Update_sell_order_id       protoreflect.FieldDescriptor
	fd_MsgUpdateSellOrders_Update_new_quantity        protoreflect.FieldDescriptor
	fd_MsgUpdateSellOrders_Update_new_ask_price       protoreflect.FieldDescriptor
	fd_MsgUpdateSellOrders_Update_disable_auto_retire protoreflect.FieldDescriptor
	fd_MsgUpdateSellOrders_Update_new_expiration      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateSellOrders_Update = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateSellOrders").Messages().ByName("Update")
	fd_MsgUpdateSellOrders_Update_sell_order_id = md_MsgUpdateSellOrders_Update.Fields().ByName("sell_order_id")
	fd_MsgUpdateSellOrders_Update_new_quantity = md_MsgUpdateSellOrders_Update.Fields().ByName("new_quantity")
	fd_MsgUpdateSellOrders_Update_new_ask_price = md_MsgUpdateSellOrders_Update.Fields().ByName("new_ask_price")
	fd_MsgUpdateSellOrders_Update_disable_auto_retire = md_MsgUpdateSellOrders_Update.Fields().ByName("disable_auto_retire")
	fd_MsgUpdateSellOrders_Update_new_expiration = md_MsgUpdateSellOrders_Update.Fields().ByName("new_expiration")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateSellOrders_Update)(nil)

type fastReflection_MsgUpdateSellOrders_Update MsgUpdateSellOrders_Update

func (x *MsgUpdateSellOrders_Update) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateSellOrders_Update)(x)
}

func (x *MsgUpdateSellOrders_Update) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[39]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateSellOrders_Update_messageType fastReflection_MsgUpdateSellOrders_Update_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateSellOrders_Update_messageType{}

type fastReflection_MsgUpdateSellOrders_Update_messageType struct{}

func (x fastReflection_MsgUpdateSellOrders_Update_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateSellOrders_Update)(nil)
}
func (x fastReflection_MsgUpdateSellOrders_Update_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateSellOrders_Update)
}
func (x fastReflection_MsgUpdateSellOrders_Update_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateSellOrders_Update
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateSellOrders_Update) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateSellOrders_Update
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateSellOrders_Update) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateSellOrders_Update_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateSellOrders_Update) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateSellOrders_Update)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateSellOrders_Update) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateSellOrders_Update)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateSellOrders_Update) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.SellOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SellOrderId)
		if !f(fd_MsgUpdateSellOrders_Update_sell_order_id, value) {
			return
		}
	}
	if x.NewQuantity != "" {
		value := protoreflect.ValueOfString(x.NewQuantity)
		if !f(fd_MsgUpdateSellOrders_Update_new_quantity, value) {
			return
		}
	}
	if x.NewAskPrice != nil {
		value := protoreflect.ValueOfMessage(x.NewAskPrice.ProtoReflect())
		if !f(fd_MsgUpdateSellOrders_Update_new_ask_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_MsgUpdateSellOrders_Update_disable_auto_retire, value) {
			return
		}
	}
	if x.NewExpiration != nil {
		value := protoreflect.ValueOfMessage(x.NewExpiration.ProtoReflect())
		if !f(fd_MsgUpdateSellOrders_Update_new_expiration, value) {
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
func (x *fastReflection_MsgUpdateSellOrders_Update) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.sell_order_id":
		return x.SellOrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_quantity":
		return x.NewQuantity != ""
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_ask_price":
		return x.NewAskPrice != nil
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_expiration":
		return x.NewExpiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateSellOrders_Update) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.sell_order_id":
		x.SellOrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_quantity":
		x.NewQuantity = ""
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_ask_price":
		x.NewAskPrice = nil
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_expiration":
		x.NewExpiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateSellOrders_Update) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.sell_order_id":
		value := x.SellOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_quantity":
		value := x.NewQuantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_ask_price":
		value := x.NewAskPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_expiration":
		value := x.NewExpiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateSellOrders_Update) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.sell_order_id":
		x.SellOrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_quantity":
		x.NewQuantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_ask_price":
		x.NewAskPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_expiration":
		x.NewExpiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateSellOrders_Update) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_ask_price":
		if x.NewAskPrice == nil {
			x.NewAskPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.NewAskPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_expiration":
		if x.NewExpiration == nil {
			x.NewExpiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.NewExpiration.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_quantity":
		panic(fmt.Errorf("field new_quantity of message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateSellOrders_Update) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_ask_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateSellOrders_Update) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateSellOrders_Update) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateSellOrders_Update) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateSellOrders_Update) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateSellOrders_Update) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateSellOrders_Update)
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
		if x.SellOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
		}
		l = len(x.NewQuantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.NewAskPrice != nil {
			l = options.Size(x.NewAskPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.NewExpiration != nil {
			l = options.Size(x.NewExpiration)
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
		x := input.Message.Interface().(*MsgUpdateSellOrders_Update)
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
		if x.NewExpiration != nil {
			encoded, err := options.Marshal(x.NewExpiration)
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
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x20
		}
		if x.NewAskPrice != nil {
			encoded, err := options.Marshal(x.NewAskPrice)
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
		if len(x.NewQuantity) > 0 {
			i -= len(x.NewQuantity)
			copy(dAtA[i:], x.NewQuantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewQuantity)))
			i--
			dAtA[i] = 0x12
		}
		if x.SellOrderId != 0 {
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
		x := input.Message.Interface().(*MsgUpdateSellOrders_Update)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateSellOrders_Update: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateSellOrders_Update: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewQuantity", wireType)
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
				x.NewQuantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAskPrice", wireType)
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
				if x.NewAskPrice == nil {
					x.NewAskPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.NewAskPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
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
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewExpiration", wireType)
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
				if x.NewExpiration == nil {
					x.NewExpiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.NewExpiration); err != nil {
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
	md_MsgUpdateSellOrdersResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgUpdateSellOrdersResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgUpdateSellOrdersResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateSellOrdersResponse)(nil)

type fastReflection_MsgUpdateSellOrdersResponse MsgUpdateSellOrdersResponse

func (x *MsgUpdateSellOrdersResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateSellOrdersResponse)(x)
}

func (x *MsgUpdateSellOrdersResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateSellOrdersResponse_messageType fastReflection_MsgUpdateSellOrdersResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateSellOrdersResponse_messageType{}

type fastReflection_MsgUpdateSellOrdersResponse_messageType struct{}

func (x fastReflection_MsgUpdateSellOrdersResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateSellOrdersResponse)(nil)
}
func (x fastReflection_MsgUpdateSellOrdersResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateSellOrdersResponse)
}
func (x fastReflection_MsgUpdateSellOrdersResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateSellOrdersResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateSellOrdersResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateSellOrdersResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateSellOrdersResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateSellOrdersResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateSellOrdersResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateSellOrdersResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateSellOrdersResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateSellOrdersResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateSellOrdersResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgUpdateSellOrdersResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateSellOrdersResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateSellOrdersResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateSellOrdersResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateSellOrdersResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateSellOrdersResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateSellOrdersResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateSellOrdersResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateSellOrdersResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateSellOrdersResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateSellOrdersResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateSellOrdersResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateSellOrdersResponse)
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
		x := input.Message.Interface().(*MsgUpdateSellOrdersResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateSellOrdersResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateSellOrdersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgBuy_2_list)(nil)

type _MsgBuy_2_list struct {
	list *[]*MsgBuy_Order
}

func (x *_MsgBuy_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgBuy_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgBuy_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgBuy_Order)
	(*x.list)[i] = concreteValue
}

func (x *_MsgBuy_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgBuy_Order)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgBuy_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgBuy_Order)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgBuy_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgBuy_2_list) NewElement() protoreflect.Value {
	v := new(MsgBuy_Order)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgBuy_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgBuy        protoreflect.MessageDescriptor
	fd_MsgBuy_buyer  protoreflect.FieldDescriptor
	fd_MsgBuy_orders protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgBuy = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgBuy")
	fd_MsgBuy_buyer = md_MsgBuy.Fields().ByName("buyer")
	fd_MsgBuy_orders = md_MsgBuy.Fields().ByName("orders")
}

var _ protoreflect.Message = (*fastReflection_MsgBuy)(nil)

type fastReflection_MsgBuy MsgBuy

func (x *MsgBuy) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgBuy)(x)
}

func (x *MsgBuy) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgBuy_messageType fastReflection_MsgBuy_messageType
var _ protoreflect.MessageType = fastReflection_MsgBuy_messageType{}

type fastReflection_MsgBuy_messageType struct{}

func (x fastReflection_MsgBuy_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgBuy)(nil)
}
func (x fastReflection_MsgBuy_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgBuy)
}
func (x fastReflection_MsgBuy_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuy
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgBuy) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuy
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgBuy) Type() protoreflect.MessageType {
	return _fastReflection_MsgBuy_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgBuy) New() protoreflect.Message {
	return new(fastReflection_MsgBuy)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgBuy) Interface() protoreflect.ProtoMessage {
	return (*MsgBuy)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgBuy) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Buyer != "" {
		value := protoreflect.ValueOfString(x.Buyer)
		if !f(fd_MsgBuy_buyer, value) {
			return
		}
	}
	if len(x.Orders) != 0 {
		value := protoreflect.ValueOfList(&_MsgBuy_2_list{list: &x.Orders})
		if !f(fd_MsgBuy_orders, value) {
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
func (x *fastReflection_MsgBuy) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.buyer":
		return x.Buyer != ""
	case "regen.ecocredit.v1alpha2.MsgBuy.orders":
		return len(x.Orders) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuy) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.buyer":
		x.Buyer = ""
	case "regen.ecocredit.v1alpha2.MsgBuy.orders":
		x.Orders = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgBuy) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.buyer":
		value := x.Buyer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgBuy.orders":
		if len(x.Orders) == 0 {
			return protoreflect.ValueOfList(&_MsgBuy_2_list{})
		}
		listValue := &_MsgBuy_2_list{list: &x.Orders}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgBuy) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.buyer":
		x.Buyer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgBuy.orders":
		lv := value.List()
		clv := lv.(*_MsgBuy_2_list)
		x.Orders = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgBuy) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.orders":
		if x.Orders == nil {
			x.Orders = []*MsgBuy_Order{}
		}
		value := &_MsgBuy_2_list{list: &x.Orders}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgBuy.buyer":
		panic(fmt.Errorf("field buyer of message regen.ecocredit.v1alpha2.MsgBuy is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgBuy) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.buyer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgBuy.orders":
		list := []*MsgBuy_Order{}
		return protoreflect.ValueOfList(&_MsgBuy_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgBuy) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgBuy", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgBuy) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuy) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgBuy) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgBuy) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgBuy)
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
		l = len(x.Buyer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Orders) > 0 {
			for _, e := range x.Orders {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgBuy)
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
		if len(x.Orders) > 0 {
			for iNdEx := len(x.Orders) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Orders[iNdEx])
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
				dAtA[i] = 0x12
			}
		}
		if len(x.Buyer) > 0 {
			i -= len(x.Buyer)
			copy(dAtA[i:], x.Buyer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Buyer)))
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
		x := input.Message.Interface().(*MsgBuy)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuy: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuy: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Orders", wireType)
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
				x.Orders = append(x.Orders, &MsgBuy_Order{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Orders[len(x.Orders)-1]); err != nil {
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
	md_MsgBuy_Order                      protoreflect.MessageDescriptor
	fd_MsgBuy_Order_selection            protoreflect.FieldDescriptor
	fd_MsgBuy_Order_quantity             protoreflect.FieldDescriptor
	fd_MsgBuy_Order_bid_price            protoreflect.FieldDescriptor
	fd_MsgBuy_Order_disable_auto_retire  protoreflect.FieldDescriptor
	fd_MsgBuy_Order_disable_partial_fill protoreflect.FieldDescriptor
	fd_MsgBuy_Order_retirement_location  protoreflect.FieldDescriptor
	fd_MsgBuy_Order_expiration           protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgBuy_Order = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgBuy").Messages().ByName("Order")
	fd_MsgBuy_Order_selection = md_MsgBuy_Order.Fields().ByName("selection")
	fd_MsgBuy_Order_quantity = md_MsgBuy_Order.Fields().ByName("quantity")
	fd_MsgBuy_Order_bid_price = md_MsgBuy_Order.Fields().ByName("bid_price")
	fd_MsgBuy_Order_disable_auto_retire = md_MsgBuy_Order.Fields().ByName("disable_auto_retire")
	fd_MsgBuy_Order_disable_partial_fill = md_MsgBuy_Order.Fields().ByName("disable_partial_fill")
	fd_MsgBuy_Order_retirement_location = md_MsgBuy_Order.Fields().ByName("retirement_location")
	fd_MsgBuy_Order_expiration = md_MsgBuy_Order.Fields().ByName("expiration")
}

var _ protoreflect.Message = (*fastReflection_MsgBuy_Order)(nil)

type fastReflection_MsgBuy_Order MsgBuy_Order

func (x *MsgBuy_Order) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgBuy_Order)(x)
}

func (x *MsgBuy_Order) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[40]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgBuy_Order_messageType fastReflection_MsgBuy_Order_messageType
var _ protoreflect.MessageType = fastReflection_MsgBuy_Order_messageType{}

type fastReflection_MsgBuy_Order_messageType struct{}

func (x fastReflection_MsgBuy_Order_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgBuy_Order)(nil)
}
func (x fastReflection_MsgBuy_Order_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgBuy_Order)
}
func (x fastReflection_MsgBuy_Order_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuy_Order
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgBuy_Order) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuy_Order
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgBuy_Order) Type() protoreflect.MessageType {
	return _fastReflection_MsgBuy_Order_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgBuy_Order) New() protoreflect.Message {
	return new(fastReflection_MsgBuy_Order)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgBuy_Order) Interface() protoreflect.ProtoMessage {
	return (*MsgBuy_Order)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgBuy_Order) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Selection != nil {
		value := protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
		if !f(fd_MsgBuy_Order_selection, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_MsgBuy_Order_quantity, value) {
			return
		}
	}
	if x.BidPrice != nil {
		value := protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
		if !f(fd_MsgBuy_Order_bid_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_MsgBuy_Order_disable_auto_retire, value) {
			return
		}
	}
	if x.DisablePartialFill != false {
		value := protoreflect.ValueOfBool(x.DisablePartialFill)
		if !f(fd_MsgBuy_Order_disable_partial_fill, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_MsgBuy_Order_retirement_location, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_MsgBuy_Order_expiration, value) {
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
func (x *fastReflection_MsgBuy_Order) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.selection":
		return x.Selection != nil
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.bid_price":
		return x.BidPrice != nil
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_partial_fill":
		return x.DisablePartialFill != false
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.retirement_location":
		return x.RetirementLocation != ""
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.expiration":
		return x.Expiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuy_Order) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.selection":
		x.Selection = nil
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.bid_price":
		x.BidPrice = nil
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_partial_fill":
		x.DisablePartialFill = false
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.retirement_location":
		x.RetirementLocation = ""
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.expiration":
		x.Expiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgBuy_Order) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.selection":
		value := x.Selection
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.bid_price":
		value := x.BidPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_partial_fill":
		value := x.DisablePartialFill
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgBuy_Order) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.selection":
		x.Selection = value.Message().Interface().(*MsgBuy_Order_Selection)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.bid_price":
		x.BidPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_partial_fill":
		x.DisablePartialFill = value.Bool()
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgBuy_Order) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.selection":
		if x.Selection == nil {
			x.Selection = new(MsgBuy_Order_Selection)
		}
		return protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.bid_price":
		if x.BidPrice == nil {
			x.BidPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1alpha2.MsgBuy.Order is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.MsgBuy.Order is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_partial_fill":
		panic(fmt.Errorf("field disable_partial_fill of message regen.ecocredit.v1alpha2.MsgBuy.Order is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha2.MsgBuy.Order is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgBuy_Order) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.selection":
		m := new(MsgBuy_Order_Selection)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.bid_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.disable_partial_fill":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.retirement_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgBuy_Order) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgBuy.Order", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgBuy_Order) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuy_Order) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgBuy_Order) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgBuy_Order) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgBuy_Order)
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
		l = len(x.RetirementLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*MsgBuy_Order)
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
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x32
		}
		if x.DisablePartialFill {
			i--
			if x.DisablePartialFill {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x20
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
			dAtA[i] = 0x1a
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x12
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
		x := input.Message.Interface().(*MsgBuy_Order)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuy_Order: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuy_Order: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
					x.Selection = &MsgBuy_Order_Selection{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Selection); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
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
			case 3:
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
			case 4:
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
			case 5:
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
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
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
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
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
	md_MsgBuy_Order_Selection               protoreflect.MessageDescriptor
	fd_MsgBuy_Order_Selection_sell_order_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgBuy_Order_Selection = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgBuy").Messages().ByName("Order").Messages().ByName("Selection")
	fd_MsgBuy_Order_Selection_sell_order_id = md_MsgBuy_Order_Selection.Fields().ByName("sell_order_id")
}

var _ protoreflect.Message = (*fastReflection_MsgBuy_Order_Selection)(nil)

type fastReflection_MsgBuy_Order_Selection MsgBuy_Order_Selection

func (x *MsgBuy_Order_Selection) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgBuy_Order_Selection)(x)
}

func (x *MsgBuy_Order_Selection) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[41]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgBuy_Order_Selection_messageType fastReflection_MsgBuy_Order_Selection_messageType
var _ protoreflect.MessageType = fastReflection_MsgBuy_Order_Selection_messageType{}

type fastReflection_MsgBuy_Order_Selection_messageType struct{}

func (x fastReflection_MsgBuy_Order_Selection_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgBuy_Order_Selection)(nil)
}
func (x fastReflection_MsgBuy_Order_Selection_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgBuy_Order_Selection)
}
func (x fastReflection_MsgBuy_Order_Selection_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuy_Order_Selection
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgBuy_Order_Selection) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuy_Order_Selection
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgBuy_Order_Selection) Type() protoreflect.MessageType {
	return _fastReflection_MsgBuy_Order_Selection_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgBuy_Order_Selection) New() protoreflect.Message {
	return new(fastReflection_MsgBuy_Order_Selection)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgBuy_Order_Selection) Interface() protoreflect.ProtoMessage {
	return (*MsgBuy_Order_Selection)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgBuy_Order_Selection) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *MsgBuy_Order_Selection_SellOrderId:
			v := o.SellOrderId
			value := protoreflect.ValueOfUint64(v)
			if !f(fd_MsgBuy_Order_Selection_sell_order_id, value) {
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
func (x *fastReflection_MsgBuy_Order_Selection) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.Selection.sell_order_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*MsgBuy_Order_Selection_SellOrderId); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order.Selection does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuy_Order_Selection) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.Selection.sell_order_id":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order.Selection does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgBuy_Order_Selection) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.Selection.sell_order_id":
		if x.Sum == nil {
			return protoreflect.ValueOfUint64(uint64(0))
		} else if v, ok := x.Sum.(*MsgBuy_Order_Selection_SellOrderId); ok {
			return protoreflect.ValueOfUint64(v.SellOrderId)
		} else {
			return protoreflect.ValueOfUint64(uint64(0))
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order.Selection does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgBuy_Order_Selection) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.Selection.sell_order_id":
		cv := value.Uint()
		x.Sum = &MsgBuy_Order_Selection_SellOrderId{SellOrderId: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order.Selection does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgBuy_Order_Selection) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.Selection.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.v1alpha2.MsgBuy.Order.Selection is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order.Selection does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgBuy_Order_Selection) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.Selection.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuy.Order.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuy.Order.Selection does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgBuy_Order_Selection) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuy.Order.Selection.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *MsgBuy_Order_Selection_SellOrderId:
			return x.Descriptor().Fields().ByName("sell_order_id")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgBuy.Order.Selection", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgBuy_Order_Selection) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuy_Order_Selection) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgBuy_Order_Selection) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgBuy_Order_Selection) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgBuy_Order_Selection)
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
		case *MsgBuy_Order_Selection_SellOrderId:
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
		x := input.Message.Interface().(*MsgBuy_Order_Selection)
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
		case *MsgBuy_Order_Selection_SellOrderId:
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
		x := input.Message.Interface().(*MsgBuy_Order_Selection)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuy_Order_Selection: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuy_Order_Selection: illegal tag %d (wire type %d)", fieldNum, wire)
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
				x.Sum = &MsgBuy_Order_Selection_SellOrderId{v}
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

var _ protoreflect.List = (*_MsgBuyResponse_1_list)(nil)

type _MsgBuyResponse_1_list struct {
	list *[]uint64
}

func (x *_MsgBuyResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgBuyResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfUint64((*x.list)[i])
}

func (x *_MsgBuyResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgBuyResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgBuyResponse_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgBuyResponse at list field BuyOrderIds as it is not of Message kind"))
}

func (x *_MsgBuyResponse_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgBuyResponse_1_list) NewElement() protoreflect.Value {
	v := uint64(0)
	return protoreflect.ValueOfUint64(v)
}

func (x *_MsgBuyResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgBuyResponse               protoreflect.MessageDescriptor
	fd_MsgBuyResponse_buy_order_ids protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgBuyResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgBuyResponse")
	fd_MsgBuyResponse_buy_order_ids = md_MsgBuyResponse.Fields().ByName("buy_order_ids")
}

var _ protoreflect.Message = (*fastReflection_MsgBuyResponse)(nil)

type fastReflection_MsgBuyResponse MsgBuyResponse

func (x *MsgBuyResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgBuyResponse)(x)
}

func (x *MsgBuyResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgBuyResponse_messageType fastReflection_MsgBuyResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgBuyResponse_messageType{}

type fastReflection_MsgBuyResponse_messageType struct{}

func (x fastReflection_MsgBuyResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgBuyResponse)(nil)
}
func (x fastReflection_MsgBuyResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgBuyResponse)
}
func (x fastReflection_MsgBuyResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuyResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgBuyResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgBuyResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgBuyResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgBuyResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgBuyResponse) New() protoreflect.Message {
	return new(fastReflection_MsgBuyResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgBuyResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgBuyResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgBuyResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.BuyOrderIds) != 0 {
		value := protoreflect.ValueOfList(&_MsgBuyResponse_1_list{list: &x.BuyOrderIds})
		if !f(fd_MsgBuyResponse_buy_order_ids, value) {
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
func (x *fastReflection_MsgBuyResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuyResponse.buy_order_ids":
		return len(x.BuyOrderIds) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuyResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuyResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuyResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuyResponse.buy_order_ids":
		x.BuyOrderIds = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuyResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuyResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgBuyResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuyResponse.buy_order_ids":
		if len(x.BuyOrderIds) == 0 {
			return protoreflect.ValueOfList(&_MsgBuyResponse_1_list{})
		}
		listValue := &_MsgBuyResponse_1_list{list: &x.BuyOrderIds}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuyResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuyResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgBuyResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuyResponse.buy_order_ids":
		lv := value.List()
		clv := lv.(*_MsgBuyResponse_1_list)
		x.BuyOrderIds = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuyResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuyResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgBuyResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuyResponse.buy_order_ids":
		if x.BuyOrderIds == nil {
			x.BuyOrderIds = []uint64{}
		}
		value := &_MsgBuyResponse_1_list{list: &x.BuyOrderIds}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuyResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuyResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgBuyResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgBuyResponse.buy_order_ids":
		list := []uint64{}
		return protoreflect.ValueOfList(&_MsgBuyResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgBuyResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgBuyResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgBuyResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgBuyResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgBuyResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgBuyResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgBuyResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgBuyResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgBuyResponse)
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
		if len(x.BuyOrderIds) > 0 {
			l = 0
			for _, e := range x.BuyOrderIds {
				l += runtime.Sov(uint64(e))
			}
			n += 1 + runtime.Sov(uint64(l)) + l
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
		x := input.Message.Interface().(*MsgBuyResponse)
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
		if len(x.BuyOrderIds) > 0 {
			var pksize2 int
			for _, num := range x.BuyOrderIds {
				pksize2 += runtime.Sov(uint64(num))
			}
			i -= pksize2
			j1 := i
			for _, num := range x.BuyOrderIds {
				for num >= 1<<7 {
					dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j1++
				}
				dAtA[j1] = uint8(num)
				j1++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize2))
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
		x := input.Message.Interface().(*MsgBuyResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuyResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgBuyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType == 0 {
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
					x.BuyOrderIds = append(x.BuyOrderIds, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.BuyOrderIds) == 0 {
						x.BuyOrderIds = make([]uint64, 0, elementCount)
					}
					for iNdEx < postIndex {
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
						x.BuyOrderIds = append(x.BuyOrderIds, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuyOrderIds", wireType)
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

var (
	md_MsgAllowAskDenom               protoreflect.MessageDescriptor
	fd_MsgAllowAskDenom_root_address  protoreflect.FieldDescriptor
	fd_MsgAllowAskDenom_denom         protoreflect.FieldDescriptor
	fd_MsgAllowAskDenom_display_denom protoreflect.FieldDescriptor
	fd_MsgAllowAskDenom_exponent      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgAllowAskDenom = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgAllowAskDenom")
	fd_MsgAllowAskDenom_root_address = md_MsgAllowAskDenom.Fields().ByName("root_address")
	fd_MsgAllowAskDenom_denom = md_MsgAllowAskDenom.Fields().ByName("denom")
	fd_MsgAllowAskDenom_display_denom = md_MsgAllowAskDenom.Fields().ByName("display_denom")
	fd_MsgAllowAskDenom_exponent = md_MsgAllowAskDenom.Fields().ByName("exponent")
}

var _ protoreflect.Message = (*fastReflection_MsgAllowAskDenom)(nil)

type fastReflection_MsgAllowAskDenom MsgAllowAskDenom

func (x *MsgAllowAskDenom) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAllowAskDenom)(x)
}

func (x *MsgAllowAskDenom) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAllowAskDenom_messageType fastReflection_MsgAllowAskDenom_messageType
var _ protoreflect.MessageType = fastReflection_MsgAllowAskDenom_messageType{}

type fastReflection_MsgAllowAskDenom_messageType struct{}

func (x fastReflection_MsgAllowAskDenom_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAllowAskDenom)(nil)
}
func (x fastReflection_MsgAllowAskDenom_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAllowAskDenom)
}
func (x fastReflection_MsgAllowAskDenom_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAllowAskDenom
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAllowAskDenom) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAllowAskDenom
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAllowAskDenom) Type() protoreflect.MessageType {
	return _fastReflection_MsgAllowAskDenom_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAllowAskDenom) New() protoreflect.Message {
	return new(fastReflection_MsgAllowAskDenom)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAllowAskDenom) Interface() protoreflect.ProtoMessage {
	return (*MsgAllowAskDenom)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAllowAskDenom) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.RootAddress != "" {
		value := protoreflect.ValueOfString(x.RootAddress)
		if !f(fd_MsgAllowAskDenom_root_address, value) {
			return
		}
	}
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_MsgAllowAskDenom_denom, value) {
			return
		}
	}
	if x.DisplayDenom != "" {
		value := protoreflect.ValueOfString(x.DisplayDenom)
		if !f(fd_MsgAllowAskDenom_display_denom, value) {
			return
		}
	}
	if x.Exponent != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Exponent)
		if !f(fd_MsgAllowAskDenom_exponent, value) {
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
func (x *fastReflection_MsgAllowAskDenom) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.root_address":
		return x.RootAddress != ""
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.denom":
		return x.Denom != ""
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.display_denom":
		return x.DisplayDenom != ""
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.exponent":
		return x.Exponent != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAllowAskDenom) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.root_address":
		x.RootAddress = ""
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.denom":
		x.Denom = ""
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.display_denom":
		x.DisplayDenom = ""
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.exponent":
		x.Exponent = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAllowAskDenom) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.root_address":
		value := x.RootAddress
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.display_denom":
		value := x.DisplayDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.exponent":
		value := x.Exponent
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenom does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAllowAskDenom) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.root_address":
		x.RootAddress = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.denom":
		x.Denom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.display_denom":
		x.DisplayDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.exponent":
		x.Exponent = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenom does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAllowAskDenom) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.root_address":
		panic(fmt.Errorf("field root_address of message regen.ecocredit.v1alpha2.MsgAllowAskDenom is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.denom":
		panic(fmt.Errorf("field denom of message regen.ecocredit.v1alpha2.MsgAllowAskDenom is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.display_denom":
		panic(fmt.Errorf("field display_denom of message regen.ecocredit.v1alpha2.MsgAllowAskDenom is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.exponent":
		panic(fmt.Errorf("field exponent of message regen.ecocredit.v1alpha2.MsgAllowAskDenom is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAllowAskDenom) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.root_address":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.display_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgAllowAskDenom.exponent":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAllowAskDenom) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgAllowAskDenom", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAllowAskDenom) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAllowAskDenom) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAllowAskDenom) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAllowAskDenom) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAllowAskDenom)
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
		l = len(x.RootAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
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
		x := input.Message.Interface().(*MsgAllowAskDenom)
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
			dAtA[i] = 0x20
		}
		if len(x.DisplayDenom) > 0 {
			i -= len(x.DisplayDenom)
			copy(dAtA[i:], x.DisplayDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DisplayDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.RootAddress) > 0 {
			i -= len(x.RootAddress)
			copy(dAtA[i:], x.RootAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RootAddress)))
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
		x := input.Message.Interface().(*MsgAllowAskDenom)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAllowAskDenom: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAllowAskDenom: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RootAddress", wireType)
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
				x.RootAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
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
			case 3:
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
			case 4:
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

var (
	md_MsgAllowAskDenomResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgAllowAskDenomResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgAllowAskDenomResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgAllowAskDenomResponse)(nil)

type fastReflection_MsgAllowAskDenomResponse MsgAllowAskDenomResponse

func (x *MsgAllowAskDenomResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAllowAskDenomResponse)(x)
}

func (x *MsgAllowAskDenomResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAllowAskDenomResponse_messageType fastReflection_MsgAllowAskDenomResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgAllowAskDenomResponse_messageType{}

type fastReflection_MsgAllowAskDenomResponse_messageType struct{}

func (x fastReflection_MsgAllowAskDenomResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAllowAskDenomResponse)(nil)
}
func (x fastReflection_MsgAllowAskDenomResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAllowAskDenomResponse)
}
func (x fastReflection_MsgAllowAskDenomResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAllowAskDenomResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAllowAskDenomResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAllowAskDenomResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAllowAskDenomResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgAllowAskDenomResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAllowAskDenomResponse) New() protoreflect.Message {
	return new(fastReflection_MsgAllowAskDenomResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAllowAskDenomResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgAllowAskDenomResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAllowAskDenomResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgAllowAskDenomResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAllowAskDenomResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAllowAskDenomResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAllowAskDenomResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAllowAskDenomResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAllowAskDenomResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAllowAskDenomResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAllowAskDenomResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAllowAskDenomResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAllowAskDenomResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAllowAskDenomResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAllowAskDenomResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgAllowAskDenomResponse)
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
		x := input.Message.Interface().(*MsgAllowAskDenomResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAllowAskDenomResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAllowAskDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgCreateBasket_5_list)(nil)

type _MsgCreateBasket_5_list struct {
	list *[]*BasketCriteria
}

func (x *_MsgCreateBasket_5_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateBasket_5_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateBasket_5_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCriteria)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateBasket_5_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCriteria)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateBasket_5_list) AppendMutable() protoreflect.Value {
	v := new(BasketCriteria)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateBasket_5_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateBasket_5_list) NewElement() protoreflect.Value {
	v := new(BasketCriteria)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateBasket_5_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateBasket                     protoreflect.MessageDescriptor
	fd_MsgCreateBasket_curator             protoreflect.FieldDescriptor
	fd_MsgCreateBasket_name                protoreflect.FieldDescriptor
	fd_MsgCreateBasket_display_name        protoreflect.FieldDescriptor
	fd_MsgCreateBasket_exponent            protoreflect.FieldDescriptor
	fd_MsgCreateBasket_basket_criteria     protoreflect.FieldDescriptor
	fd_MsgCreateBasket_disable_auto_retire protoreflect.FieldDescriptor
	fd_MsgCreateBasket_allow_picking       protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateBasket = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateBasket")
	fd_MsgCreateBasket_curator = md_MsgCreateBasket.Fields().ByName("curator")
	fd_MsgCreateBasket_name = md_MsgCreateBasket.Fields().ByName("name")
	fd_MsgCreateBasket_display_name = md_MsgCreateBasket.Fields().ByName("display_name")
	fd_MsgCreateBasket_exponent = md_MsgCreateBasket.Fields().ByName("exponent")
	fd_MsgCreateBasket_basket_criteria = md_MsgCreateBasket.Fields().ByName("basket_criteria")
	fd_MsgCreateBasket_disable_auto_retire = md_MsgCreateBasket.Fields().ByName("disable_auto_retire")
	fd_MsgCreateBasket_allow_picking = md_MsgCreateBasket.Fields().ByName("allow_picking")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBasket)(nil)

type fastReflection_MsgCreateBasket MsgCreateBasket

func (x *MsgCreateBasket) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBasket)(x)
}

func (x *MsgCreateBasket) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBasket_messageType fastReflection_MsgCreateBasket_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBasket_messageType{}

type fastReflection_MsgCreateBasket_messageType struct{}

func (x fastReflection_MsgCreateBasket_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBasket)(nil)
}
func (x fastReflection_MsgCreateBasket_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBasket)
}
func (x fastReflection_MsgCreateBasket_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBasket
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBasket) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBasket
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBasket) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBasket_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBasket) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBasket)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBasket) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBasket)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBasket) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Curator != "" {
		value := protoreflect.ValueOfString(x.Curator)
		if !f(fd_MsgCreateBasket_curator, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_MsgCreateBasket_name, value) {
			return
		}
	}
	if x.DisplayName != "" {
		value := protoreflect.ValueOfString(x.DisplayName)
		if !f(fd_MsgCreateBasket_display_name, value) {
			return
		}
	}
	if x.Exponent != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Exponent)
		if !f(fd_MsgCreateBasket_exponent, value) {
			return
		}
	}
	if len(x.BasketCriteria) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateBasket_5_list{list: &x.BasketCriteria})
		if !f(fd_MsgCreateBasket_basket_criteria, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_MsgCreateBasket_disable_auto_retire, value) {
			return
		}
	}
	if x.AllowPicking != false {
		value := protoreflect.ValueOfBool(x.AllowPicking)
		if !f(fd_MsgCreateBasket_allow_picking, value) {
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
func (x *fastReflection_MsgCreateBasket) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.curator":
		return x.Curator != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.name":
		return x.Name != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.display_name":
		return x.DisplayName != ""
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.exponent":
		return x.Exponent != uint32(0)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.basket_criteria":
		return len(x.BasketCriteria) != 0
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.allow_picking":
		return x.AllowPicking != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasket does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBasket) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.curator":
		x.Curator = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.name":
		x.Name = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.display_name":
		x.DisplayName = ""
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.exponent":
		x.Exponent = uint32(0)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.basket_criteria":
		x.BasketCriteria = nil
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.allow_picking":
		x.AllowPicking = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasket does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBasket) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.curator":
		value := x.Curator
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.display_name":
		value := x.DisplayName
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.exponent":
		value := x.Exponent
		return protoreflect.ValueOfUint32(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.basket_criteria":
		if len(x.BasketCriteria) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateBasket_5_list{})
		}
		listValue := &_MsgCreateBasket_5_list{list: &x.BasketCriteria}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.allow_picking":
		value := x.AllowPicking
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasket does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBasket) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.curator":
		x.Curator = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.name":
		x.Name = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.display_name":
		x.DisplayName = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.exponent":
		x.Exponent = uint32(value.Uint())
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.basket_criteria":
		lv := value.List()
		clv := lv.(*_MsgCreateBasket_5_list)
		x.BasketCriteria = *clv.list
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.allow_picking":
		x.AllowPicking = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasket does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBasket) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.basket_criteria":
		if x.BasketCriteria == nil {
			x.BasketCriteria = []*BasketCriteria{}
		}
		value := &_MsgCreateBasket_5_list{list: &x.BasketCriteria}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.curator":
		panic(fmt.Errorf("field curator of message regen.ecocredit.v1alpha2.MsgCreateBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.name":
		panic(fmt.Errorf("field name of message regen.ecocredit.v1alpha2.MsgCreateBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.display_name":
		panic(fmt.Errorf("field display_name of message regen.ecocredit.v1alpha2.MsgCreateBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.exponent":
		panic(fmt.Errorf("field exponent of message regen.ecocredit.v1alpha2.MsgCreateBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.MsgCreateBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.allow_picking":
		panic(fmt.Errorf("field allow_picking of message regen.ecocredit.v1alpha2.MsgCreateBasket is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasket does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBasket) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.curator":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.name":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.display_name":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.exponent":
		return protoreflect.ValueOfUint32(uint32(0))
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.basket_criteria":
		list := []*BasketCriteria{}
		return protoreflect.ValueOfList(&_MsgCreateBasket_5_list{list: &list})
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.MsgCreateBasket.allow_picking":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasket does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBasket) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateBasket", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBasket) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBasket) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBasket) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBasket) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBasket)
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
		l = len(x.Curator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DisplayName)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Exponent != 0 {
			n += 1 + runtime.Sov(uint64(x.Exponent))
		}
		if len(x.BasketCriteria) > 0 {
			for _, e := range x.BasketCriteria {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.AllowPicking {
			n += 2
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
		x := input.Message.Interface().(*MsgCreateBasket)
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
		if x.AllowPicking {
			i--
			if x.AllowPicking {
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
		if len(x.BasketCriteria) > 0 {
			for iNdEx := len(x.BasketCriteria) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.BasketCriteria[iNdEx])
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
		}
		if x.Exponent != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exponent))
			i--
			dAtA[i] = 0x20
		}
		if len(x.DisplayName) > 0 {
			i -= len(x.DisplayName)
			copy(dAtA[i:], x.DisplayName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DisplayName)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Curator) > 0 {
			i -= len(x.Curator)
			copy(dAtA[i:], x.Curator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Curator)))
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
		x := input.Message.Interface().(*MsgCreateBasket)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBasket: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBasket: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Curator", wireType)
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
				x.Curator = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
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
				x.DisplayName = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
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
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BasketCriteria", wireType)
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
				x.BasketCriteria = append(x.BasketCriteria, &BasketCriteria{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BasketCriteria[len(x.BasketCriteria)-1]); err != nil {
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AllowPicking", wireType)
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
				x.AllowPicking = bool(v != 0)
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
	md_MsgCreateBasketResponse              protoreflect.MessageDescriptor
	fd_MsgCreateBasketResponse_basket_denom protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgCreateBasketResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgCreateBasketResponse")
	fd_MsgCreateBasketResponse_basket_denom = md_MsgCreateBasketResponse.Fields().ByName("basket_denom")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBasketResponse)(nil)

type fastReflection_MsgCreateBasketResponse MsgCreateBasketResponse

func (x *MsgCreateBasketResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBasketResponse)(x)
}

func (x *MsgCreateBasketResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBasketResponse_messageType fastReflection_MsgCreateBasketResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBasketResponse_messageType{}

type fastReflection_MsgCreateBasketResponse_messageType struct{}

func (x fastReflection_MsgCreateBasketResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBasketResponse)(nil)
}
func (x fastReflection_MsgCreateBasketResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBasketResponse)
}
func (x fastReflection_MsgCreateBasketResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBasketResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBasketResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBasketResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBasketResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBasketResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBasketResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBasketResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBasketResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBasketResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBasketResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BasketDenom != "" {
		value := protoreflect.ValueOfString(x.BasketDenom)
		if !f(fd_MsgCreateBasketResponse_basket_denom, value) {
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
func (x *fastReflection_MsgCreateBasketResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasketResponse.basket_denom":
		return x.BasketDenom != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBasketResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasketResponse.basket_denom":
		x.BasketDenom = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBasketResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasketResponse.basket_denom":
		value := x.BasketDenom
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasketResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBasketResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasketResponse.basket_denom":
		x.BasketDenom = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasketResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBasketResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasketResponse.basket_denom":
		panic(fmt.Errorf("field basket_denom of message regen.ecocredit.v1alpha2.MsgCreateBasketResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasketResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBasketResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgCreateBasketResponse.basket_denom":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgCreateBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgCreateBasketResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBasketResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgCreateBasketResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBasketResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBasketResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBasketResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBasketResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBasketResponse)
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
		l = len(x.BasketDenom)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgCreateBasketResponse)
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
		if len(x.BasketDenom) > 0 {
			i -= len(x.BasketDenom)
			copy(dAtA[i:], x.BasketDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BasketDenom)))
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
		x := input.Message.Interface().(*MsgCreateBasketResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBasketResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBasketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
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
				x.BasketDenom = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_MsgAddToBasket_3_list)(nil)

type _MsgAddToBasket_3_list struct {
	list *[]*BasketCredit
}

func (x *_MsgAddToBasket_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgAddToBasket_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgAddToBasket_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCredit)
	(*x.list)[i] = concreteValue
}

func (x *_MsgAddToBasket_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCredit)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgAddToBasket_3_list) AppendMutable() protoreflect.Value {
	v := new(BasketCredit)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgAddToBasket_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgAddToBasket_3_list) NewElement() protoreflect.Value {
	v := new(BasketCredit)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgAddToBasket_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgAddToBasket              protoreflect.MessageDescriptor
	fd_MsgAddToBasket_owner        protoreflect.FieldDescriptor
	fd_MsgAddToBasket_basket_denom protoreflect.FieldDescriptor
	fd_MsgAddToBasket_credits      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgAddToBasket = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgAddToBasket")
	fd_MsgAddToBasket_owner = md_MsgAddToBasket.Fields().ByName("owner")
	fd_MsgAddToBasket_basket_denom = md_MsgAddToBasket.Fields().ByName("basket_denom")
	fd_MsgAddToBasket_credits = md_MsgAddToBasket.Fields().ByName("credits")
}

var _ protoreflect.Message = (*fastReflection_MsgAddToBasket)(nil)

type fastReflection_MsgAddToBasket MsgAddToBasket

func (x *MsgAddToBasket) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAddToBasket)(x)
}

func (x *MsgAddToBasket) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[28]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAddToBasket_messageType fastReflection_MsgAddToBasket_messageType
var _ protoreflect.MessageType = fastReflection_MsgAddToBasket_messageType{}

type fastReflection_MsgAddToBasket_messageType struct{}

func (x fastReflection_MsgAddToBasket_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAddToBasket)(nil)
}
func (x fastReflection_MsgAddToBasket_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAddToBasket)
}
func (x fastReflection_MsgAddToBasket_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAddToBasket
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAddToBasket) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAddToBasket
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAddToBasket) Type() protoreflect.MessageType {
	return _fastReflection_MsgAddToBasket_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAddToBasket) New() protoreflect.Message {
	return new(fastReflection_MsgAddToBasket)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAddToBasket) Interface() protoreflect.ProtoMessage {
	return (*MsgAddToBasket)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAddToBasket) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_MsgAddToBasket_owner, value) {
			return
		}
	}
	if x.BasketDenom != "" {
		value := protoreflect.ValueOfString(x.BasketDenom)
		if !f(fd_MsgAddToBasket_basket_denom, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgAddToBasket_3_list{list: &x.Credits})
		if !f(fd_MsgAddToBasket_credits, value) {
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
func (x *fastReflection_MsgAddToBasket) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.basket_denom":
		return x.BasketDenom != ""
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.credits":
		return len(x.Credits) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasket does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAddToBasket) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.owner":
		x.Owner = ""
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.basket_denom":
		x.BasketDenom = ""
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.credits":
		x.Credits = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasket does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAddToBasket) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.basket_denom":
		value := x.BasketDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgAddToBasket_3_list{})
		}
		listValue := &_MsgAddToBasket_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasket does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAddToBasket) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.basket_denom":
		x.BasketDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.credits":
		lv := value.List()
		clv := lv.(*_MsgAddToBasket_3_list)
		x.Credits = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasket does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAddToBasket) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.credits":
		if x.Credits == nil {
			x.Credits = []*BasketCredit{}
		}
		value := &_MsgAddToBasket_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.MsgAddToBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.basket_denom":
		panic(fmt.Errorf("field basket_denom of message regen.ecocredit.v1alpha2.MsgAddToBasket is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasket does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAddToBasket) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.basket_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgAddToBasket.credits":
		list := []*BasketCredit{}
		return protoreflect.ValueOfList(&_MsgAddToBasket_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasket does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAddToBasket) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgAddToBasket", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAddToBasket) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAddToBasket) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAddToBasket) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAddToBasket) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAddToBasket)
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
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BasketDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgAddToBasket)
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
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
		}
		if len(x.BasketDenom) > 0 {
			i -= len(x.BasketDenom)
			copy(dAtA[i:], x.BasketDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BasketDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
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
		x := input.Message.Interface().(*MsgAddToBasket)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAddToBasket: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAddToBasket: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
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
				x.BasketDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &BasketCredit{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
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
	md_MsgAddToBasketResponse                 protoreflect.MessageDescriptor
	fd_MsgAddToBasketResponse_amount_received protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgAddToBasketResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgAddToBasketResponse")
	fd_MsgAddToBasketResponse_amount_received = md_MsgAddToBasketResponse.Fields().ByName("amount_received")
}

var _ protoreflect.Message = (*fastReflection_MsgAddToBasketResponse)(nil)

type fastReflection_MsgAddToBasketResponse MsgAddToBasketResponse

func (x *MsgAddToBasketResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAddToBasketResponse)(x)
}

func (x *MsgAddToBasketResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[29]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAddToBasketResponse_messageType fastReflection_MsgAddToBasketResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgAddToBasketResponse_messageType{}

type fastReflection_MsgAddToBasketResponse_messageType struct{}

func (x fastReflection_MsgAddToBasketResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAddToBasketResponse)(nil)
}
func (x fastReflection_MsgAddToBasketResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAddToBasketResponse)
}
func (x fastReflection_MsgAddToBasketResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAddToBasketResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAddToBasketResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAddToBasketResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAddToBasketResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgAddToBasketResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAddToBasketResponse) New() protoreflect.Message {
	return new(fastReflection_MsgAddToBasketResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAddToBasketResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgAddToBasketResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAddToBasketResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.AmountReceived != "" {
		value := protoreflect.ValueOfString(x.AmountReceived)
		if !f(fd_MsgAddToBasketResponse_amount_received, value) {
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
func (x *fastReflection_MsgAddToBasketResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasketResponse.amount_received":
		return x.AmountReceived != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAddToBasketResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasketResponse.amount_received":
		x.AmountReceived = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAddToBasketResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasketResponse.amount_received":
		value := x.AmountReceived
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasketResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAddToBasketResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasketResponse.amount_received":
		x.AmountReceived = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasketResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAddToBasketResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasketResponse.amount_received":
		panic(fmt.Errorf("field amount_received of message regen.ecocredit.v1alpha2.MsgAddToBasketResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasketResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAddToBasketResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgAddToBasketResponse.amount_received":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgAddToBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgAddToBasketResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAddToBasketResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgAddToBasketResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAddToBasketResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAddToBasketResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAddToBasketResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAddToBasketResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAddToBasketResponse)
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
		l = len(x.AmountReceived)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgAddToBasketResponse)
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
		if len(x.AmountReceived) > 0 {
			i -= len(x.AmountReceived)
			copy(dAtA[i:], x.AmountReceived)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AmountReceived)))
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
		x := input.Message.Interface().(*MsgAddToBasketResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAddToBasketResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAddToBasketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AmountReceived", wireType)
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
				x.AmountReceived = string(dAtA[iNdEx:postIndex])
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
	md_MsgTakeFromBasket                     protoreflect.MessageDescriptor
	fd_MsgTakeFromBasket_owner               protoreflect.FieldDescriptor
	fd_MsgTakeFromBasket_basket_denom        protoreflect.FieldDescriptor
	fd_MsgTakeFromBasket_amount              protoreflect.FieldDescriptor
	fd_MsgTakeFromBasket_retirement_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgTakeFromBasket = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgTakeFromBasket")
	fd_MsgTakeFromBasket_owner = md_MsgTakeFromBasket.Fields().ByName("owner")
	fd_MsgTakeFromBasket_basket_denom = md_MsgTakeFromBasket.Fields().ByName("basket_denom")
	fd_MsgTakeFromBasket_amount = md_MsgTakeFromBasket.Fields().ByName("amount")
	fd_MsgTakeFromBasket_retirement_location = md_MsgTakeFromBasket.Fields().ByName("retirement_location")
}

var _ protoreflect.Message = (*fastReflection_MsgTakeFromBasket)(nil)

type fastReflection_MsgTakeFromBasket MsgTakeFromBasket

func (x *MsgTakeFromBasket) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgTakeFromBasket)(x)
}

func (x *MsgTakeFromBasket) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[30]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgTakeFromBasket_messageType fastReflection_MsgTakeFromBasket_messageType
var _ protoreflect.MessageType = fastReflection_MsgTakeFromBasket_messageType{}

type fastReflection_MsgTakeFromBasket_messageType struct{}

func (x fastReflection_MsgTakeFromBasket_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgTakeFromBasket)(nil)
}
func (x fastReflection_MsgTakeFromBasket_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgTakeFromBasket)
}
func (x fastReflection_MsgTakeFromBasket_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTakeFromBasket
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgTakeFromBasket) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTakeFromBasket
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgTakeFromBasket) Type() protoreflect.MessageType {
	return _fastReflection_MsgTakeFromBasket_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgTakeFromBasket) New() protoreflect.Message {
	return new(fastReflection_MsgTakeFromBasket)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgTakeFromBasket) Interface() protoreflect.ProtoMessage {
	return (*MsgTakeFromBasket)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgTakeFromBasket) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_MsgTakeFromBasket_owner, value) {
			return
		}
	}
	if x.BasketDenom != "" {
		value := protoreflect.ValueOfString(x.BasketDenom)
		if !f(fd_MsgTakeFromBasket_basket_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_MsgTakeFromBasket_amount, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_MsgTakeFromBasket_retirement_location, value) {
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
func (x *fastReflection_MsgTakeFromBasket) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.basket_denom":
		return x.BasketDenom != ""
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.amount":
		return x.Amount != ""
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.retirement_location":
		return x.RetirementLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasket does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTakeFromBasket) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.owner":
		x.Owner = ""
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.basket_denom":
		x.BasketDenom = ""
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.amount":
		x.Amount = ""
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.retirement_location":
		x.RetirementLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasket does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgTakeFromBasket) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.basket_denom":
		value := x.BasketDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasket does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgTakeFromBasket) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.basket_denom":
		x.BasketDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.amount":
		x.Amount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasket does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgTakeFromBasket) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.MsgTakeFromBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.basket_denom":
		panic(fmt.Errorf("field basket_denom of message regen.ecocredit.v1alpha2.MsgTakeFromBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.amount":
		panic(fmt.Errorf("field amount of message regen.ecocredit.v1alpha2.MsgTakeFromBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha2.MsgTakeFromBasket is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasket does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgTakeFromBasket) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.basket_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasket.retirement_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasket does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgTakeFromBasket) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgTakeFromBasket", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgTakeFromBasket) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTakeFromBasket) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgTakeFromBasket) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgTakeFromBasket) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgTakeFromBasket)
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
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BasketDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetirementLocation)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgTakeFromBasket)
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
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.BasketDenom) > 0 {
			i -= len(x.BasketDenom)
			copy(dAtA[i:], x.BasketDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BasketDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
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
		x := input.Message.Interface().(*MsgTakeFromBasket)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTakeFromBasket: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTakeFromBasket: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
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
				x.BasketDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
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
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_MsgTakeFromBasketResponse_1_list)(nil)

type _MsgTakeFromBasketResponse_1_list struct {
	list *[]*BasketCredit
}

func (x *_MsgTakeFromBasketResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgTakeFromBasketResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgTakeFromBasketResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCredit)
	(*x.list)[i] = concreteValue
}

func (x *_MsgTakeFromBasketResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCredit)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgTakeFromBasketResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(BasketCredit)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgTakeFromBasketResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgTakeFromBasketResponse_1_list) NewElement() protoreflect.Value {
	v := new(BasketCredit)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgTakeFromBasketResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgTakeFromBasketResponse         protoreflect.MessageDescriptor
	fd_MsgTakeFromBasketResponse_credits protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgTakeFromBasketResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgTakeFromBasketResponse")
	fd_MsgTakeFromBasketResponse_credits = md_MsgTakeFromBasketResponse.Fields().ByName("credits")
}

var _ protoreflect.Message = (*fastReflection_MsgTakeFromBasketResponse)(nil)

type fastReflection_MsgTakeFromBasketResponse MsgTakeFromBasketResponse

func (x *MsgTakeFromBasketResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgTakeFromBasketResponse)(x)
}

func (x *MsgTakeFromBasketResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[31]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgTakeFromBasketResponse_messageType fastReflection_MsgTakeFromBasketResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgTakeFromBasketResponse_messageType{}

type fastReflection_MsgTakeFromBasketResponse_messageType struct{}

func (x fastReflection_MsgTakeFromBasketResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgTakeFromBasketResponse)(nil)
}
func (x fastReflection_MsgTakeFromBasketResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgTakeFromBasketResponse)
}
func (x fastReflection_MsgTakeFromBasketResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTakeFromBasketResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgTakeFromBasketResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTakeFromBasketResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgTakeFromBasketResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgTakeFromBasketResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgTakeFromBasketResponse) New() protoreflect.Message {
	return new(fastReflection_MsgTakeFromBasketResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgTakeFromBasketResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgTakeFromBasketResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgTakeFromBasketResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgTakeFromBasketResponse_1_list{list: &x.Credits})
		if !f(fd_MsgTakeFromBasketResponse_credits, value) {
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
func (x *fastReflection_MsgTakeFromBasketResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse.credits":
		return len(x.Credits) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTakeFromBasketResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse.credits":
		x.Credits = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgTakeFromBasketResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgTakeFromBasketResponse_1_list{})
		}
		listValue := &_MsgTakeFromBasketResponse_1_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgTakeFromBasketResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse.credits":
		lv := value.List()
		clv := lv.(*_MsgTakeFromBasketResponse_1_list)
		x.Credits = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgTakeFromBasketResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse.credits":
		if x.Credits == nil {
			x.Credits = []*BasketCredit{}
		}
		value := &_MsgTakeFromBasketResponse_1_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgTakeFromBasketResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse.credits":
		list := []*BasketCredit{}
		return protoreflect.ValueOfList(&_MsgTakeFromBasketResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgTakeFromBasketResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgTakeFromBasketResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTakeFromBasketResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgTakeFromBasketResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgTakeFromBasketResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgTakeFromBasketResponse)
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
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgTakeFromBasketResponse)
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
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*MsgTakeFromBasketResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTakeFromBasketResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTakeFromBasketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &BasketCredit{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
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

var _ protoreflect.List = (*_MsgPickFromBasket_3_list)(nil)

type _MsgPickFromBasket_3_list struct {
	list *[]*BasketCredit
}

func (x *_MsgPickFromBasket_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgPickFromBasket_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgPickFromBasket_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCredit)
	(*x.list)[i] = concreteValue
}

func (x *_MsgPickFromBasket_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*BasketCredit)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgPickFromBasket_3_list) AppendMutable() protoreflect.Value {
	v := new(BasketCredit)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgPickFromBasket_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgPickFromBasket_3_list) NewElement() protoreflect.Value {
	v := new(BasketCredit)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgPickFromBasket_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgPickFromBasket                     protoreflect.MessageDescriptor
	fd_MsgPickFromBasket_owner               protoreflect.FieldDescriptor
	fd_MsgPickFromBasket_basket_denom        protoreflect.FieldDescriptor
	fd_MsgPickFromBasket_credits             protoreflect.FieldDescriptor
	fd_MsgPickFromBasket_retirement_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgPickFromBasket = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgPickFromBasket")
	fd_MsgPickFromBasket_owner = md_MsgPickFromBasket.Fields().ByName("owner")
	fd_MsgPickFromBasket_basket_denom = md_MsgPickFromBasket.Fields().ByName("basket_denom")
	fd_MsgPickFromBasket_credits = md_MsgPickFromBasket.Fields().ByName("credits")
	fd_MsgPickFromBasket_retirement_location = md_MsgPickFromBasket.Fields().ByName("retirement_location")
}

var _ protoreflect.Message = (*fastReflection_MsgPickFromBasket)(nil)

type fastReflection_MsgPickFromBasket MsgPickFromBasket

func (x *MsgPickFromBasket) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgPickFromBasket)(x)
}

func (x *MsgPickFromBasket) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[32]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgPickFromBasket_messageType fastReflection_MsgPickFromBasket_messageType
var _ protoreflect.MessageType = fastReflection_MsgPickFromBasket_messageType{}

type fastReflection_MsgPickFromBasket_messageType struct{}

func (x fastReflection_MsgPickFromBasket_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgPickFromBasket)(nil)
}
func (x fastReflection_MsgPickFromBasket_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgPickFromBasket)
}
func (x fastReflection_MsgPickFromBasket_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgPickFromBasket
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgPickFromBasket) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgPickFromBasket
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgPickFromBasket) Type() protoreflect.MessageType {
	return _fastReflection_MsgPickFromBasket_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgPickFromBasket) New() protoreflect.Message {
	return new(fastReflection_MsgPickFromBasket)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgPickFromBasket) Interface() protoreflect.ProtoMessage {
	return (*MsgPickFromBasket)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgPickFromBasket) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_MsgPickFromBasket_owner, value) {
			return
		}
	}
	if x.BasketDenom != "" {
		value := protoreflect.ValueOfString(x.BasketDenom)
		if !f(fd_MsgPickFromBasket_basket_denom, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgPickFromBasket_3_list{list: &x.Credits})
		if !f(fd_MsgPickFromBasket_credits, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_MsgPickFromBasket_retirement_location, value) {
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
func (x *fastReflection_MsgPickFromBasket) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.basket_denom":
		return x.BasketDenom != ""
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.credits":
		return len(x.Credits) != 0
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.retirement_location":
		return x.RetirementLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasket does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgPickFromBasket) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.owner":
		x.Owner = ""
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.basket_denom":
		x.BasketDenom = ""
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.credits":
		x.Credits = nil
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.retirement_location":
		x.RetirementLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasket does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgPickFromBasket) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.basket_denom":
		value := x.BasketDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgPickFromBasket_3_list{})
		}
		listValue := &_MsgPickFromBasket_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasket does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgPickFromBasket) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.basket_denom":
		x.BasketDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.credits":
		lv := value.List()
		clv := lv.(*_MsgPickFromBasket_3_list)
		x.Credits = *clv.list
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasket does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgPickFromBasket) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.credits":
		if x.Credits == nil {
			x.Credits = []*BasketCredit{}
		}
		value := &_MsgPickFromBasket_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.MsgPickFromBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.basket_denom":
		panic(fmt.Errorf("field basket_denom of message regen.ecocredit.v1alpha2.MsgPickFromBasket is not mutable"))
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha2.MsgPickFromBasket is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasket does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgPickFromBasket) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.basket_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.credits":
		list := []*BasketCredit{}
		return protoreflect.ValueOfList(&_MsgPickFromBasket_3_list{list: &list})
	case "regen.ecocredit.v1alpha2.MsgPickFromBasket.retirement_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasket"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasket does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgPickFromBasket) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgPickFromBasket", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgPickFromBasket) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgPickFromBasket) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgPickFromBasket) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgPickFromBasket) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgPickFromBasket)
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
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BasketDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.RetirementLocation)
		if l > 0 {
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
		x := input.Message.Interface().(*MsgPickFromBasket)
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
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
		}
		if len(x.BasketDenom) > 0 {
			i -= len(x.BasketDenom)
			copy(dAtA[i:], x.BasketDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BasketDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
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
		x := input.Message.Interface().(*MsgPickFromBasket)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgPickFromBasket: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgPickFromBasket: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
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
				x.BasketDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &BasketCredit{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
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
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
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
	md_MsgPickFromBasketResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_tx_proto_init()
	md_MsgPickFromBasketResponse = File_regen_ecocredit_v1alpha2_tx_proto.Messages().ByName("MsgPickFromBasketResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgPickFromBasketResponse)(nil)

type fastReflection_MsgPickFromBasketResponse MsgPickFromBasketResponse

func (x *MsgPickFromBasketResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgPickFromBasketResponse)(x)
}

func (x *MsgPickFromBasketResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[33]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgPickFromBasketResponse_messageType fastReflection_MsgPickFromBasketResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgPickFromBasketResponse_messageType{}

type fastReflection_MsgPickFromBasketResponse_messageType struct{}

func (x fastReflection_MsgPickFromBasketResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgPickFromBasketResponse)(nil)
}
func (x fastReflection_MsgPickFromBasketResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgPickFromBasketResponse)
}
func (x fastReflection_MsgPickFromBasketResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgPickFromBasketResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgPickFromBasketResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgPickFromBasketResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgPickFromBasketResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgPickFromBasketResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgPickFromBasketResponse) New() protoreflect.Message {
	return new(fastReflection_MsgPickFromBasketResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgPickFromBasketResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgPickFromBasketResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgPickFromBasketResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgPickFromBasketResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgPickFromBasketResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgPickFromBasketResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasketResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgPickFromBasketResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasketResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgPickFromBasketResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgPickFromBasketResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.MsgPickFromBasketResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.MsgPickFromBasketResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgPickFromBasketResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.MsgPickFromBasketResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgPickFromBasketResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgPickFromBasketResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgPickFromBasketResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgPickFromBasketResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgPickFromBasketResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgPickFromBasketResponse)
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
		x := input.Message.Interface().(*MsgPickFromBasketResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgPickFromBasketResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgPickFromBasketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
// source: regen/ecocredit/v1alpha2/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgCreateClass is the Msg/CreateClass request type.
type MsgCreateClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that created the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// issuers are the account addresses of the approved issuers.
	Issuers []string `protobuf:"bytes,2,rep,name=issuers,proto3" json:"issuers,omitempty"`
	// metadata is any arbitrary metadata to attached to the credit class.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// credit_type_name describes the type of credit (e.g. "carbon", "biodiversity").
	CreditTypeName string `protobuf:"bytes,4,opt,name=credit_type_name,json=creditTypeName,proto3" json:"credit_type_name,omitempty"`
}

func (x *MsgCreateClass) Reset() {
	*x = MsgCreateClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateClass) ProtoMessage() {}

// Deprecated: Use MsgCreateClass.ProtoReflect.Descriptor instead.
func (*MsgCreateClass) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateClass) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgCreateClass) GetIssuers() []string {
	if x != nil {
		return x.Issuers
	}
	return nil
}

func (x *MsgCreateClass) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateClass) GetCreditTypeName() string {
	if x != nil {
		return x.CreditTypeName
	}
	return ""
}

// MsgCreateClassResponse is the Msg/CreateClass response type.
type MsgCreateClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// class_id is the unique ID of the newly created credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *MsgCreateClassResponse) Reset() {
	*x = MsgCreateClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateClassResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateClassResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateClassResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCreateClassResponse) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

// MsgCreateProjectResponse is the Msg/CreateProject request type.
type MsgCreateProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// issuer is the address of an approved issuer for the credit class through
	// which batches will be issued. It is not required, however, that this same
	// issuer issue all batches for a project.
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// class_id is the unique ID of the class within which the project is created.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// metadata is any arbitrary metadata attached to the project.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// project_location is the location of the project backing the credits in this
	// batch. It is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters. country-code is required, while sub-national-code
	// and postal-code can be added for increasing precision.
	ProjectLocation string `protobuf:"bytes,4,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
	// project_id is an optional user-specified project ID which can be used
	// instead of an auto-generated ID. If project_id is provided, it must be
	// unique within the credit class and match the regex [A-Za-z0-9]{2,16}
	// or else the operation will fail. If project_id is omitted an ID will
	// automatically be generated.
	ProjectId string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *MsgCreateProject) Reset() {
	*x = MsgCreateProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateProject) ProtoMessage() {}

// Deprecated: Use MsgCreateProject.ProtoReflect.Descriptor instead.
func (*MsgCreateProject) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgCreateProject) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *MsgCreateProject) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgCreateProject) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateProject) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

func (x *MsgCreateProject) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// MsgCreateProjectResponse is the Msg/CreateProject response type.
type MsgCreateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the ID of the newly created project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *MsgCreateProjectResponse) Reset() {
	*x = MsgCreateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateProjectResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateProjectResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{3}
}

func (x *MsgCreateProjectResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// MsgCreateBatch is the Msg/CreateBatch request type.
type MsgCreateBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// issuer is the address of the batch issuer.
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// project_id is the unique ID of the project this batch belongs to.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// issuance are the credits issued in the batch.
	Issuance []*MsgCreateBatch_BatchIssuance `protobuf:"bytes,3,rep,name=issuance,proto3" json:"issuance,omitempty"`
	// metadata is any arbitrary metadata attached to the credit batch.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// start_date is the beginning of the period during which this credit batch
	// was quantified and verified.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the end of the period during which this credit batch was
	// quantified and verified.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *MsgCreateBatch) Reset() {
	*x = MsgCreateBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBatch) ProtoMessage() {}

// Deprecated: Use MsgCreateBatch.ProtoReflect.Descriptor instead.
func (*MsgCreateBatch) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgCreateBatch) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *MsgCreateBatch) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *MsgCreateBatch) GetIssuance() []*MsgCreateBatch_BatchIssuance {
	if x != nil {
		return x.Issuance
	}
	return nil
}

func (x *MsgCreateBatch) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateBatch) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *MsgCreateBatch) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

// MsgCreateBatchResponse is the Msg/CreateBatch response type.
type MsgCreateBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique denomination ID of the newly created batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
}

func (x *MsgCreateBatchResponse) Reset() {
	*x = MsgCreateBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBatchResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateBatchResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateBatchResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{5}
}

func (x *MsgCreateBatchResponse) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

// MsgSend is the Msg/Send request type.
type MsgSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sender is the address of the account sending credits.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// sender is the address of the account receiving credits.
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// credits are the credits being sent.
	Credits []*MsgSend_SendCredits `protobuf:"bytes,3,rep,name=credits,proto3" json:"credits,omitempty"`
}

func (x *MsgSend) Reset() {
	*x = MsgSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSend) ProtoMessage() {}

// Deprecated: Use MsgSend.ProtoReflect.Descriptor instead.
func (*MsgSend) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{6}
}

func (x *MsgSend) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgSend) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *MsgSend) GetCredits() []*MsgSend_SendCredits {
	if x != nil {
		return x.Credits
	}
	return nil
}

// MsgSendResponse is the Msg/Send response type.
type MsgSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSendResponse) Reset() {
	*x = MsgSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSendResponse) ProtoMessage() {}

// Deprecated: Use MsgSendResponse.ProtoReflect.Descriptor instead.
func (*MsgSendResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{7}
}

// MsgRetire is the Msg/Retire request type.
type MsgRetire struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// holder is the credit holder address.
	Holder string `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	// credits are the credits being retired.
	Credits []*MsgRetire_RetireCredits `protobuf:"bytes,2,rep,name=credits,proto3" json:"credits,omitempty"`
	// location is the location of the beneficiary or buyer of the retired
	// credits. It is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters.
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *MsgRetire) Reset() {
	*x = MsgRetire{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRetire) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRetire) ProtoMessage() {}

// Deprecated: Use MsgRetire.ProtoReflect.Descriptor instead.
func (*MsgRetire) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{8}
}

func (x *MsgRetire) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

func (x *MsgRetire) GetCredits() []*MsgRetire_RetireCredits {
	if x != nil {
		return x.Credits
	}
	return nil
}

func (x *MsgRetire) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

// MsgRetire is the Msg/Retire response type.
type MsgRetireResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgRetireResponse) Reset() {
	*x = MsgRetireResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRetireResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRetireResponse) ProtoMessage() {}

// Deprecated: Use MsgRetireResponse.ProtoReflect.Descriptor instead.
func (*MsgRetireResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{9}
}

// MsgCancel is the Msg/Cancel request type.
type MsgCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// holder is the credit holder address.
	Holder string `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	// credits are the credits being cancelled.
	Credits []*MsgCancel_CancelCredits `protobuf:"bytes,2,rep,name=credits,proto3" json:"credits,omitempty"`
}

func (x *MsgCancel) Reset() {
	*x = MsgCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCancel) ProtoMessage() {}

// Deprecated: Use MsgCancel.ProtoReflect.Descriptor instead.
func (*MsgCancel) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{10}
}

func (x *MsgCancel) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

func (x *MsgCancel) GetCredits() []*MsgCancel_CancelCredits {
	if x != nil {
		return x.Credits
	}
	return nil
}

// MsgCancelResponse is the Msg/Cancel response type.
type MsgCancelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCancelResponse) Reset() {
	*x = MsgCancelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCancelResponse) ProtoMessage() {}

// Deprecated: Use MsgCancelResponse.ProtoReflect.Descriptor instead.
func (*MsgCancelResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{11}
}

// MsgUpdateClassAdmin is the Msg/UpdateClassAdmin request type.
type MsgUpdateClassAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that is the admin of the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// class_id is the unique ID of the credit class.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// new_admin is the address of the new admin of the credit class.
	NewAdmin string `protobuf:"bytes,3,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (x *MsgUpdateClassAdmin) Reset() {
	*x = MsgUpdateClassAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassAdmin) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassAdmin.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassAdmin) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{12}
}

func (x *MsgUpdateClassAdmin) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateClassAdmin) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgUpdateClassAdmin) GetNewAdmin() string {
	if x != nil {
		return x.NewAdmin
	}
	return ""
}

// MsgUpdateClassAdminResponse is the MsgUpdateClassAdmin response type.
type MsgUpdateClassAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateClassAdminResponse) Reset() {
	*x = MsgUpdateClassAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassAdminResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassAdminResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassAdminResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{13}
}

// MsgUpdateClassIssuers is the Msg/UpdateClassIssuers request type.
type MsgUpdateClassIssuers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that is the admin of the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// class_id is the unique ID of the credit class.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// issuers are the updated account addresses of the approved issuers.
	Issuers []string `protobuf:"bytes,3,rep,name=issuers,proto3" json:"issuers,omitempty"`
}

func (x *MsgUpdateClassIssuers) Reset() {
	*x = MsgUpdateClassIssuers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassIssuers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassIssuers) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassIssuers.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassIssuers) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{14}
}

func (x *MsgUpdateClassIssuers) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateClassIssuers) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgUpdateClassIssuers) GetIssuers() []string {
	if x != nil {
		return x.Issuers
	}
	return nil
}

// MsgUpdateClassIssuersResponse is the MsgUpdateClassIssuers response type.
type MsgUpdateClassIssuersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateClassIssuersResponse) Reset() {
	*x = MsgUpdateClassIssuersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassIssuersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassIssuersResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassIssuersResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassIssuersResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{15}
}

// MsgUpdateClassMetadata is the Msg/UpdateClassMetadata request type.
type MsgUpdateClassMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that is the admin of the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// class_id is the unique ID of the credit class.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// metadata is the updated arbitrary metadata to be attached to the credit class.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgUpdateClassMetadata) Reset() {
	*x = MsgUpdateClassMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassMetadata) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassMetadata.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassMetadata) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{16}
}

func (x *MsgUpdateClassMetadata) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateClassMetadata) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgUpdateClassMetadata) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUpdateClassMetadataResponse is the MsgUpdateClassMetadata response type.
type MsgUpdateClassMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateClassMetadataResponse) Reset() {
	*x = MsgUpdateClassMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassMetadataResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassMetadataResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassMetadataResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{17}
}

// MsgSell is the Msg/Sell request type.
type MsgSell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// owner is the address of the owner of the credits being sold.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// orders are the sell orders being created.
	Orders []*MsgSell_Order `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *MsgSell) Reset() {
	*x = MsgSell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSell) ProtoMessage() {}

// Deprecated: Use MsgSell.ProtoReflect.Descriptor instead.
func (*MsgSell) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{18}
}

func (x *MsgSell) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *MsgSell) GetOrders() []*MsgSell_Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

// MsgSellResponse is the Msg/Sell response type.
type MsgSellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sell_order_ids are the sell order IDs of the newly created sell orders.
	SellOrderIds []uint64 `protobuf:"varint,1,rep,packed,name=sell_order_ids,json=sellOrderIds,proto3" json:"sell_order_ids,omitempty"`
}

func (x *MsgSellResponse) Reset() {
	*x = MsgSellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSellResponse) ProtoMessage() {}

// Deprecated: Use MsgSellResponse.ProtoReflect.Descriptor instead.
func (*MsgSellResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{19}
}

func (x *MsgSellResponse) GetSellOrderIds() []uint64 {
	if x != nil {
		return x.SellOrderIds
	}
	return nil
}

// MsgUpdateSellOrders is the Msg/UpdateSellOrders request type.
type MsgUpdateSellOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// owner is the owner of the sell orders.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// updates are updates to existing sell orders.
	Updates []*MsgUpdateSellOrders_Update `protobuf:"bytes,2,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (x *MsgUpdateSellOrders) Reset() {
	*x = MsgUpdateSellOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateSellOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateSellOrders) ProtoMessage() {}

// Deprecated: Use MsgUpdateSellOrders.ProtoReflect.Descriptor instead.
func (*MsgUpdateSellOrders) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{20}
}

func (x *MsgUpdateSellOrders) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *MsgUpdateSellOrders) GetUpdates() []*MsgUpdateSellOrders_Update {
	if x != nil {
		return x.Updates
	}
	return nil
}

// MsgUpdateSellOrdersResponse is the Msg/UpdateSellOrders response type.
type MsgUpdateSellOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateSellOrdersResponse) Reset() {
	*x = MsgUpdateSellOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateSellOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateSellOrdersResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateSellOrdersResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateSellOrdersResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{21}
}

// MsgBuy is the Msg/Buy request type.
type MsgBuy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buyer is the address of the credit buyer.
	Buyer string `protobuf:"bytes,1,opt,name=buyer,proto3" json:"buyer,omitempty"`
	// orders are the new buy orders.
	Orders []*MsgBuy_Order `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *MsgBuy) Reset() {
	*x = MsgBuy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBuy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBuy) ProtoMessage() {}

// Deprecated: Use MsgBuy.ProtoReflect.Descriptor instead.
func (*MsgBuy) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{22}
}

func (x *MsgBuy) GetBuyer() string {
	if x != nil {
		return x.Buyer
	}
	return ""
}

func (x *MsgBuy) GetOrders() []*MsgBuy_Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

// MsgBuyResponse is the Msg/Buy response type.
type MsgBuyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buy_order_ids are the buy order IDs of the newly created buy orders. Buy
	// orders may not settle instantaneously, but rather in batches at specified
	// batch epoch times.
	BuyOrderIds []uint64 `protobuf:"varint,1,rep,packed,name=buy_order_ids,json=buyOrderIds,proto3" json:"buy_order_ids,omitempty"`
}

func (x *MsgBuyResponse) Reset() {
	*x = MsgBuyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBuyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBuyResponse) ProtoMessage() {}

// Deprecated: Use MsgBuyResponse.ProtoReflect.Descriptor instead.
func (*MsgBuyResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{23}
}

func (x *MsgBuyResponse) GetBuyOrderIds() []uint64 {
	if x != nil {
		return x.BuyOrderIds
	}
	return nil
}

// MsgAllowAskDenom is the Msg/AllowAskDenom request type.
type MsgAllowAskDenom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// root_address is the address of the governance account which can authorize ask denoms
	RootAddress string `protobuf:"bytes,1,opt,name=root_address,json=rootAddress,proto3" json:"root_address,omitempty"`
	// denom is the denom to allow (ex. ibc/GLKHDSG423SGS)
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// display_denom is the denom to display to the user and is informational
	DisplayDenom string `protobuf:"bytes,3,opt,name=display_denom,json=displayDenom,proto3" json:"display_denom,omitempty"`
	// exponent is the exponent that relates the denom to the display_denom and is
	// informational
	Exponent uint32 `protobuf:"varint,4,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (x *MsgAllowAskDenom) Reset() {
	*x = MsgAllowAskDenom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAllowAskDenom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAllowAskDenom) ProtoMessage() {}

// Deprecated: Use MsgAllowAskDenom.ProtoReflect.Descriptor instead.
func (*MsgAllowAskDenom) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{24}
}

func (x *MsgAllowAskDenom) GetRootAddress() string {
	if x != nil {
		return x.RootAddress
	}
	return ""
}

func (x *MsgAllowAskDenom) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *MsgAllowAskDenom) GetDisplayDenom() string {
	if x != nil {
		return x.DisplayDenom
	}
	return ""
}

func (x *MsgAllowAskDenom) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

// MsgAllowAskDenomResponse is the Msg/AllowAskDenom response type.
type MsgAllowAskDenomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgAllowAskDenomResponse) Reset() {
	*x = MsgAllowAskDenomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAllowAskDenomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAllowAskDenomResponse) ProtoMessage() {}

// Deprecated: Use MsgAllowAskDenomResponse.ProtoReflect.Descriptor instead.
func (*MsgAllowAskDenomResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{25}
}

// MsgCreateBasket is the Msg/CreateBasket request type.
type MsgCreateBasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// curator is the address of the basket curator who is able to change certain
	// basket settings.
	Curator string `protobuf:"bytes,1,opt,name=curator,proto3" json:"curator,omitempty"`
	// name will be used to create a bank denom for this basket token of the form
	// ecocredit:{curator}:{name}.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// display_name will be used to create a bank Metadata display name for this
	// basket token of the form ecocredit:{curator}:{display_name}.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// exponent is the exponent that will be used for denom metadata. An exponent
	// of 6 will mean that 10^6 units of a basket token should be displayed
	// as one unit in user interfaces.
	Exponent uint32 `protobuf:"varint,4,opt,name=exponent,proto3" json:"exponent,omitempty"`
	// basket_criteria is the criteria by which credits can be added to the
	// basket. Basket criteria will be applied in order and the first criteria
	// which applies to a credit will determine its multiplier in the basket.
	BasketCriteria []*BasketCriteria `protobuf:"bytes,5,rep,name=basket_criteria,json=basketCriteria,proto3" json:"basket_criteria,omitempty"`
	// disable_auto_retire allows auto-retirement to be disabled.
	// The credits will be auto-retired if disable_auto_retire is
	// false unless the credits were previously put into the basket by the
	// address picking them from the basket, in which case they will remain
	// tradable.
	DisableAutoRetire bool `protobuf:"varint,6,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// allow_picking specifies whether an address which didn't deposit the credits
	// in the basket can pick those credits or not.
	AllowPicking bool `protobuf:"varint,7,opt,name=allow_picking,json=allowPicking,proto3" json:"allow_picking,omitempty"`
}

func (x *MsgCreateBasket) Reset() {
	*x = MsgCreateBasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBasket) ProtoMessage() {}

// Deprecated: Use MsgCreateBasket.ProtoReflect.Descriptor instead.
func (*MsgCreateBasket) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{26}
}

func (x *MsgCreateBasket) GetCurator() string {
	if x != nil {
		return x.Curator
	}
	return ""
}

func (x *MsgCreateBasket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MsgCreateBasket) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *MsgCreateBasket) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

func (x *MsgCreateBasket) GetBasketCriteria() []*BasketCriteria {
	if x != nil {
		return x.BasketCriteria
	}
	return nil
}

func (x *MsgCreateBasket) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *MsgCreateBasket) GetAllowPicking() bool {
	if x != nil {
		return x.AllowPicking
	}
	return false
}

// MsgCreateBasketResponse is the Msg/CreateBasket response type.
type MsgCreateBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// basket_denom is the unique denomination ID of the newly created basket.
	BasketDenom string `protobuf:"bytes,1,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
}

func (x *MsgCreateBasketResponse) Reset() {
	*x = MsgCreateBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBasketResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateBasketResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateBasketResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{27}
}

func (x *MsgCreateBasketResponse) GetBasketDenom() string {
	if x != nil {
		return x.BasketDenom
	}
	return ""
}

// MsgAddToBasket is the Msg/AddToBasket request type.
type MsgAddToBasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// owner is the owner of credits being added to the basket.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// basket_denom is the basket denom to add credits to.
	BasketDenom string `protobuf:"bytes,2,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// credits are credits to add to the basket. If they do not match the basket's
	// admission criteria the operation will fail.
	Credits []*BasketCredit `protobuf:"bytes,3,rep,name=credits,proto3" json:"credits,omitempty"`
}

func (x *MsgAddToBasket) Reset() {
	*x = MsgAddToBasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[28]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAddToBasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAddToBasket) ProtoMessage() {}

// Deprecated: Use MsgAddToBasket.ProtoReflect.Descriptor instead.
func (*MsgAddToBasket) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{28}
}

func (x *MsgAddToBasket) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *MsgAddToBasket) GetBasketDenom() string {
	if x != nil {
		return x.BasketDenom
	}
	return ""
}

func (x *MsgAddToBasket) GetCredits() []*BasketCredit {
	if x != nil {
		return x.Credits
	}
	return nil
}

// MsgAddToBasketResponse is the Msg/AddToBasket response type.
type MsgAddToBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// amount_received is the amount of basket tokens received.
	AmountReceived string `protobuf:"bytes,1,opt,name=amount_received,json=amountReceived,proto3" json:"amount_received,omitempty"`
}

func (x *MsgAddToBasketResponse) Reset() {
	*x = MsgAddToBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[29]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAddToBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAddToBasketResponse) ProtoMessage() {}

// Deprecated: Use MsgAddToBasketResponse.ProtoReflect.Descriptor instead.
func (*MsgAddToBasketResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{29}
}

func (x *MsgAddToBasketResponse) GetAmountReceived() string {
	if x != nil {
		return x.AmountReceived
	}
	return ""
}

// MsgTakeFromBasket is the Msg/TakeFromBasket request type.
type MsgTakeFromBasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// owner is the owner of the basket tokens.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// basket_denom is the basket denom to take credits from.
	BasketDenom string `protobuf:"bytes,2,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// amount is the number of credits to take from the basket.
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// retirement_location is the optional retirement location for the credits
	// which will be used only if retire_on_take is true for this basket.
	RetirementLocation string `protobuf:"bytes,4,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
}

func (x *MsgTakeFromBasket) Reset() {
	*x = MsgTakeFromBasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[30]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTakeFromBasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTakeFromBasket) ProtoMessage() {}

// Deprecated: Use MsgTakeFromBasket.ProtoReflect.Descriptor instead.
func (*MsgTakeFromBasket) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{30}
}

func (x *MsgTakeFromBasket) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *MsgTakeFromBasket) GetBasketDenom() string {
	if x != nil {
		return x.BasketDenom
	}
	return ""
}

func (x *MsgTakeFromBasket) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *MsgTakeFromBasket) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

// MsgTakeFromBasketResponse is the Msg/TakeFromBasket response type.
type MsgTakeFromBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// credits are the credits taken out of the basket.
	Credits []*BasketCredit `protobuf:"bytes,1,rep,name=credits,proto3" json:"credits,omitempty"`
}

func (x *MsgTakeFromBasketResponse) Reset() {
	*x = MsgTakeFromBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[31]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTakeFromBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTakeFromBasketResponse) ProtoMessage() {}

// Deprecated: Use MsgTakeFromBasketResponse.ProtoReflect.Descriptor instead.
func (*MsgTakeFromBasketResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{31}
}

func (x *MsgTakeFromBasketResponse) GetCredits() []*BasketCredit {
	if x != nil {
		return x.Credits
	}
	return nil
}

// MsgPickFromBasket is the Msg/PickFromBasket request type.
type MsgPickFromBasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// owner is the owner of the basket tokens.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// basket_denom is the basket denom to pick credits from.
	BasketDenom string `protobuf:"bytes,2,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// credits are the units of credits being picked from the basket
	Credits []*BasketCredit `protobuf:"bytes,3,rep,name=credits,proto3" json:"credits,omitempty"`
	// retirement_location is the optional retirement location for the credits
	// which will be used only if retire_on_take is true for this basket.
	RetirementLocation string `protobuf:"bytes,4,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
}

func (x *MsgPickFromBasket) Reset() {
	*x = MsgPickFromBasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[32]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPickFromBasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPickFromBasket) ProtoMessage() {}

// Deprecated: Use MsgPickFromBasket.ProtoReflect.Descriptor instead.
func (*MsgPickFromBasket) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{32}
}

func (x *MsgPickFromBasket) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *MsgPickFromBasket) GetBasketDenom() string {
	if x != nil {
		return x.BasketDenom
	}
	return ""
}

func (x *MsgPickFromBasket) GetCredits() []*BasketCredit {
	if x != nil {
		return x.Credits
	}
	return nil
}

func (x *MsgPickFromBasket) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

// MsgPickFromBasketResponse is the Msg/PickFromBasket response type.
type MsgPickFromBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgPickFromBasketResponse) Reset() {
	*x = MsgPickFromBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[33]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPickFromBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPickFromBasketResponse) ProtoMessage() {}

// Deprecated: Use MsgPickFromBasketResponse.ProtoReflect.Descriptor instead.
func (*MsgPickFromBasketResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{33}
}

// BatchIssuance represents the issuance of some credits in a batch to a
// single recipient.
type MsgCreateBatch_BatchIssuance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// recipient is the account of the recipient.
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// tradable_amount is the number of credits in this issuance that can be
	// traded by this recipient. Decimal values are acceptable.
	TradableAmount string `protobuf:"bytes,2,opt,name=tradable_amount,json=tradableAmount,proto3" json:"tradable_amount,omitempty"`
	// retired_amount is the number of credits in this issuance that are
	// effectively retired by the issuer on receipt. Decimal values are
	// acceptable.
	RetiredAmount string `protobuf:"bytes,3,opt,name=retired_amount,json=retiredAmount,proto3" json:"retired_amount,omitempty"`
	// retirement_location is the location of the beneficiary or buyer of the
	// retired credits. This must be provided if retired_amount is positive. It
	// is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters.
	RetirementLocation string `protobuf:"bytes,4,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
}

func (x *MsgCreateBatch_BatchIssuance) Reset() {
	*x = MsgCreateBatch_BatchIssuance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[34]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBatch_BatchIssuance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBatch_BatchIssuance) ProtoMessage() {}

// Deprecated: Use MsgCreateBatch_BatchIssuance.ProtoReflect.Descriptor instead.
func (*MsgCreateBatch_BatchIssuance) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{4, 0}
}

func (x *MsgCreateBatch_BatchIssuance) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *MsgCreateBatch_BatchIssuance) GetTradableAmount() string {
	if x != nil {
		return x.TradableAmount
	}
	return ""
}

func (x *MsgCreateBatch_BatchIssuance) GetRetiredAmount() string {
	if x != nil {
		return x.RetiredAmount
	}
	return ""
}

func (x *MsgCreateBatch_BatchIssuance) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

// SendCredits specifies a batch and the number of credits being transferred.
// This is split into tradable credits, which will remain tradable on receipt,
// and retired credits, which will be retired on receipt.
type MsgSend_SendCredits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// tradable_amount is the number of credits in this transfer that can be
	// traded by the recipient. Decimal values are acceptable within the
	// precision returned by Query/Precision.
	TradableAmount string `protobuf:"bytes,2,opt,name=tradable_amount,json=tradableAmount,proto3" json:"tradable_amount,omitempty"`
	// retired_amount is the number of credits in this transfer that are
	// effectively retired by the issuer on receipt. Decimal values are
	// acceptable within the precision returned by Query/Precision.
	RetiredAmount string `protobuf:"bytes,3,opt,name=retired_amount,json=retiredAmount,proto3" json:"retired_amount,omitempty"`
	// retirement_location is the location of the beneficiary or buyer of the
	// retired credits. This must be provided if retired_amount is positive. It
	// is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters.
	RetirementLocation string `protobuf:"bytes,4,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
}

func (x *MsgSend_SendCredits) Reset() {
	*x = MsgSend_SendCredits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[35]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSend_SendCredits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSend_SendCredits) ProtoMessage() {}

// Deprecated: Use MsgSend_SendCredits.ProtoReflect.Descriptor instead.
func (*MsgSend_SendCredits) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{6, 0}
}

func (x *MsgSend_SendCredits) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *MsgSend_SendCredits) GetTradableAmount() string {
	if x != nil {
		return x.TradableAmount
	}
	return ""
}

func (x *MsgSend_SendCredits) GetRetiredAmount() string {
	if x != nil {
		return x.RetiredAmount
	}
	return ""
}

func (x *MsgSend_SendCredits) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

// RetireCredits specifies a batch and the number of credits being retired.
type MsgRetire_RetireCredits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the number of credits being retired.
	// Decimal values are acceptable within the precision returned by
	// Query/Precision.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgRetire_RetireCredits) Reset() {
	*x = MsgRetire_RetireCredits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[36]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRetire_RetireCredits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRetire_RetireCredits) ProtoMessage() {}

// Deprecated: Use MsgRetire_RetireCredits.ProtoReflect.Descriptor instead.
func (*MsgRetire_RetireCredits) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{8, 0}
}

func (x *MsgRetire_RetireCredits) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *MsgRetire_RetireCredits) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

// CancelCredits specifies a batch and the number of credits being cancelled.
type MsgCancel_CancelCredits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the number of credits being cancelled.
	// Decimal values are acceptable within the precision returned by
	// Query/Precision.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgCancel_CancelCredits) Reset() {
	*x = MsgCancel_CancelCredits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[37]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCancel_CancelCredits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCancel_CancelCredits) ProtoMessage() {}

// Deprecated: Use MsgCancel_CancelCredits.ProtoReflect.Descriptor instead.
func (*MsgCancel_CancelCredits) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{10, 0}
}

func (x *MsgCancel_CancelCredits) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *MsgCancel_CancelCredits) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

// Order is the content of a new sell order.
type MsgSell_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the credit batch being sold.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// quantity is the quantity of credits being sold from this batch. If it is
	// less then the balance of credits the owner has available at the time this
	// sell order is matched, the quantity will be adjusted downwards to the
	// owner's balance. However, if the balance of credits is less than this
	// quantity at the time the sell order is created, the operation will fail.
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// ask_price is the price the seller is asking for each unit of the
	// batch_denom. Each credit unit of the batch will be sold for at least the
	// ask_price or more.
	AskPrice *v1beta1.Coin `protobuf:"bytes,3,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	// disable_auto_retire disables auto-retirement of credits which allows a
	// buyer to disable auto-retirement in their buy order enabling them to
	// resell the credits to another buyer.
	DisableAutoRetire bool `protobuf:"varint,4,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// expiration is an optional timestamp when the sell order expires. When the
	// expiration time is reached, the sell order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *MsgSell_Order) Reset() {
	*x = MsgSell_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[38]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSell_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSell_Order) ProtoMessage() {}

// Deprecated: Use MsgSell_Order.ProtoReflect.Descriptor instead.
func (*MsgSell_Order) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{18, 0}
}

func (x *MsgSell_Order) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *MsgSell_Order) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *MsgSell_Order) GetAskPrice() *v1beta1.Coin {
	if x != nil {
		return x.AskPrice
	}
	return nil
}

func (x *MsgSell_Order) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *MsgSell_Order) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// Update is an update to an existing sell order.
type MsgUpdateSellOrders_Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  sell_order_id is the ID of an existing sell order.
	SellOrderId uint64 `protobuf:"varint,1,opt,name=sell_order_id,json=sellOrderId,proto3" json:"sell_order_id,omitempty"`
	// new_quantity is the updated quantity of credits available to sell, if it
	// is set to zero then the order is cancelled.
	NewQuantity string `protobuf:"bytes,2,opt,name=new_quantity,json=newQuantity,proto3" json:"new_quantity,omitempty"`
	// new_ask_price is the new ask price for this sell order
	NewAskPrice *v1beta1.Coin `protobuf:"bytes,3,opt,name=new_ask_price,json=newAskPrice,proto3" json:"new_ask_price,omitempty"`
	// disable_auto_retire updates the disable_auto_retire field in the sell order.
	DisableAutoRetire bool `protobuf:"varint,4,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// new_expiration is an optional timestamp when the sell order expires. When the
	// expiration time is reached, the sell order is removed from state.
	NewExpiration *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=new_expiration,json=newExpiration,proto3" json:"new_expiration,omitempty"`
}

func (x *MsgUpdateSellOrders_Update) Reset() {
	*x = MsgUpdateSellOrders_Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[39]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateSellOrders_Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateSellOrders_Update) ProtoMessage() {}

// Deprecated: Use MsgUpdateSellOrders_Update.ProtoReflect.Descriptor instead.
func (*MsgUpdateSellOrders_Update) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{20, 0}
}

func (x *MsgUpdateSellOrders_Update) GetSellOrderId() uint64 {
	if x != nil {
		return x.SellOrderId
	}
	return 0
}

func (x *MsgUpdateSellOrders_Update) GetNewQuantity() string {
	if x != nil {
		return x.NewQuantity
	}
	return ""
}

func (x *MsgUpdateSellOrders_Update) GetNewAskPrice() *v1beta1.Coin {
	if x != nil {
		return x.NewAskPrice
	}
	return nil
}

func (x *MsgUpdateSellOrders_Update) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *MsgUpdateSellOrders_Update) GetNewExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.NewExpiration
	}
	return nil
}

// Order is a buy order.
type MsgBuy_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// selection is the buy order selection.
	Selection *MsgBuy_Order_Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	// quantity is the quantity of credits to buy. If the quantity of credits
	// available is less than this amount the order will be partially filled
	// unless disable_partial_fill is true.
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// bid price is the bid price for this buy order. A credit unit will be
	// settled at a purchase price that is no more than the bid price. The
	// buy order will fail if the buyer does not have enough funds available
	// to complete the purchase.
	BidPrice *v1beta1.Coin `protobuf:"bytes,3,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	// disable_auto_retire allows auto-retirement to be disabled. If it is set to true
	// the credits will not auto-retire and can be resold assuming that the
	// corresponding sell order has auto-retirement disabled. If the sell order
	// hasn't disabled auto-retirement and the buy order tries to disable it,
	// that buy order will fail.
	DisableAutoRetire bool `protobuf:"varint,4,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// disable_partial_fill disables the default behavior of partially filling
	// buy orders if the requested quantity is not available.
	DisablePartialFill bool `protobuf:"varint,5,opt,name=disable_partial_fill,json=disablePartialFill,proto3" json:"disable_partial_fill,omitempty"`
	// retirement_location is the optional retirement location for the credits
	// which will be used only if disable_auto_retire is false.
	RetirementLocation string `protobuf:"bytes,6,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
	// expiration is the optional timestamp when the buy order expires. When the
	// expiration time is reached, the buy order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *MsgBuy_Order) Reset() {
	*x = MsgBuy_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[40]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBuy_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBuy_Order) ProtoMessage() {}

// Deprecated: Use MsgBuy_Order.ProtoReflect.Descriptor instead.
func (*MsgBuy_Order) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{22, 0}
}

func (x *MsgBuy_Order) GetSelection() *MsgBuy_Order_Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

func (x *MsgBuy_Order) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *MsgBuy_Order) GetBidPrice() *v1beta1.Coin {
	if x != nil {
		return x.BidPrice
	}
	return nil
}

func (x *MsgBuy_Order) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *MsgBuy_Order) GetDisablePartialFill() bool {
	if x != nil {
		return x.DisablePartialFill
	}
	return false
}

func (x *MsgBuy_Order) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

func (x *MsgBuy_Order) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// Selection defines a buy order selection.
type MsgBuy_Order_Selection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum defines the type of selection.
	//
	// Types that are assignable to Sum:
	//	*MsgBuy_Order_Selection_SellOrderId
	Sum isMsgBuy_Order_Selection_Sum `protobuf_oneof:"sum"`
}

func (x *MsgBuy_Order_Selection) Reset() {
	*x = MsgBuy_Order_Selection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[41]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBuy_Order_Selection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBuy_Order_Selection) ProtoMessage() {}

// Deprecated: Use MsgBuy_Order_Selection.ProtoReflect.Descriptor instead.
func (*MsgBuy_Order_Selection) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP(), []int{22, 0, 0}
}

func (x *MsgBuy_Order_Selection) GetSum() isMsgBuy_Order_Selection_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *MsgBuy_Order_Selection) GetSellOrderId() uint64 {
	if x, ok := x.GetSum().(*MsgBuy_Order_Selection_SellOrderId); ok {
		return x.SellOrderId
	}
	return 0
}

type isMsgBuy_Order_Selection_Sum interface {
	isMsgBuy_Order_Selection_Sum()
}

type MsgBuy_Order_Selection_SellOrderId struct {
	// sell_order_id is the sell order ID against which the buyer is trying to buy.
	// When sell_order_id is set, this is known as a direct buy order because it
	// is placed directly against a specific sell order.
	SellOrderId uint64 `protobuf:"varint,1,opt,name=sell_order_id,json=sellOrderId,proto3,oneof"`
}

func (*MsgBuy_Order_Selection_SellOrderId) isMsgBuy_Order_Selection_Sum() {}

var File_regen_ecocredit_v1alpha2_tx_proto protoreflect.FileDescriptor

var file_regen_ecocredit_v1alpha2_tx_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x10, 0x4d, 0x73, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0xe6, 0x03, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x73, 0x73,
	0x75, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f,
	0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x1a, 0xae, 0x01, 0x0a, 0x0d, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x16, 0x4d, 0x73,
	0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0xba, 0x02, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73,
	0x1a, 0xaf, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x64,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x74, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74, 0x69, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x48, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x13,
	0x0a, 0x11, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x1a, 0x48, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x13, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x4d, 0x73,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x0a, 0x15, 0x4d, 0x73, 0x67,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x22, 0x1f, 0x0a,
	0x1d, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65,
	0x0a, 0x16, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x1e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd1, 0x02, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53,
	0x65, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6c, 0x6c, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0xee, 0x01, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0f, 0x4d,
	0x73, 0x67, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x22, 0x85, 0x03, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x4e, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d,
	0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x1a, 0x87, 0x02, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x41, 0x73, 0x6b, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0d, 0x6e,
	0x65, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x1b,
	0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9b, 0x04, 0x0a, 0x06,
	0x4d, 0x73, 0x67, 0x42, 0x75, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x79, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0xba, 0x03, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x69, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x69,
	0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x6c, 0x12, 0x2f, 0x0a, 0x13,
	0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x74, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90,
	0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x38, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x34, 0x0a, 0x0e, 0x4d, 0x73, 0x67,
	0x42, 0x75, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x62,
	0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x0b, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22,
	0x8c, 0x01, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x6b, 0x44,
	0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x22, 0x1a,
	0x0a, 0x18, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x6b, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa6, 0x02, 0x0a, 0x0f, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x0f, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x0e,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x2e,
	0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x69, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x22, 0x3c, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x40, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x22,
	0x41, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x42, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x54, 0x61, 0x6b, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x19, 0x4d, 0x73,
	0x67, 0x54, 0x61, 0x6b, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x11, 0x4d, 0x73,
	0x67, 0x50, 0x69, 0x63, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65,
	0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1b, 0x0a, 0x19, 0x4d,
	0x73, 0x67, 0x50, 0x69, 0x63, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xad, 0x0e, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x12, 0x69, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x28, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x1a, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2a, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x32, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x28, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x21, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65,
	0x6e, 0x64, 0x1a, 0x29, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73,
	0x67, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x06, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x1a, 0x2b, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74, 0x69, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d,
	0x73, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x1a, 0x2b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2d, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x35, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x7e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x1a, 0x37, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x81, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x38, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x21, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6c, 0x6c, 0x1a, 0x29,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x35, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x03, 0x42, 0x75, 0x79, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x79, 0x1a, 0x28, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x41,
	0x73, 0x6b, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x2a, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x6b, 0x44, 0x65,
	0x6e, 0x6f, 0x6d, 0x1a, 0x32, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d,
	0x73, 0x67, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x6b, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x1a, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73,
	0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x4d, 0x73, 0x67, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x1a, 0x30,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x72, 0x0a, 0x0e, 0x54, 0x61, 0x6b, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73,
	0x67, 0x54, 0x61, 0x6b, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x1a,
	0x33, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x61,
	0x6b, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x0e, 0x50, 0x69, 0x63, 0x6b, 0x46, 0x72, 0x6f, 0x6d,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x4d, 0x73, 0x67, 0x50, 0x69, 0x63, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d,
	0x73, 0x67, 0x50, 0x69, 0x63, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xff, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x07, 0x54, 0x78, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x58,
	0xaa, 0x02, 0x18, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x18, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xe2, 0x02, 0x24, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a,
	0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_regen_ecocredit_v1alpha2_tx_proto_rawDescOnce sync.Once
	file_regen_ecocredit_v1alpha2_tx_proto_rawDescData = file_regen_ecocredit_v1alpha2_tx_proto_rawDesc
)

func file_regen_ecocredit_v1alpha2_tx_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_v1alpha2_tx_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_v1alpha2_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_v1alpha2_tx_proto_rawDescData)
	})
	return file_regen_ecocredit_v1alpha2_tx_proto_rawDescData
}

var file_regen_ecocredit_v1alpha2_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 42)
var file_regen_ecocredit_v1alpha2_tx_proto_goTypes = []interface{}{
	(*MsgCreateClass)(nil),                 // 0: regen.ecocredit.v1alpha2.MsgCreateClass
	(*MsgCreateClassResponse)(nil),         // 1: regen.ecocredit.v1alpha2.MsgCreateClassResponse
	(*MsgCreateProject)(nil),               // 2: regen.ecocredit.v1alpha2.MsgCreateProject
	(*MsgCreateProjectResponse)(nil),       // 3: regen.ecocredit.v1alpha2.MsgCreateProjectResponse
	(*MsgCreateBatch)(nil),                 // 4: regen.ecocredit.v1alpha2.MsgCreateBatch
	(*MsgCreateBatchResponse)(nil),         // 5: regen.ecocredit.v1alpha2.MsgCreateBatchResponse
	(*MsgSend)(nil),                        // 6: regen.ecocredit.v1alpha2.MsgSend
	(*MsgSendResponse)(nil),                // 7: regen.ecocredit.v1alpha2.MsgSendResponse
	(*MsgRetire)(nil),                      // 8: regen.ecocredit.v1alpha2.MsgRetire
	(*MsgRetireResponse)(nil),              // 9: regen.ecocredit.v1alpha2.MsgRetireResponse
	(*MsgCancel)(nil),                      // 10: regen.ecocredit.v1alpha2.MsgCancel
	(*MsgCancelResponse)(nil),              // 11: regen.ecocredit.v1alpha2.MsgCancelResponse
	(*MsgUpdateClassAdmin)(nil),            // 12: regen.ecocredit.v1alpha2.MsgUpdateClassAdmin
	(*MsgUpdateClassAdminResponse)(nil),    // 13: regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse
	(*MsgUpdateClassIssuers)(nil),          // 14: regen.ecocredit.v1alpha2.MsgUpdateClassIssuers
	(*MsgUpdateClassIssuersResponse)(nil),  // 15: regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse
	(*MsgUpdateClassMetadata)(nil),         // 16: regen.ecocredit.v1alpha2.MsgUpdateClassMetadata
	(*MsgUpdateClassMetadataResponse)(nil), // 17: regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse
	(*MsgSell)(nil),                        // 18: regen.ecocredit.v1alpha2.MsgSell
	(*MsgSellResponse)(nil),                // 19: regen.ecocredit.v1alpha2.MsgSellResponse
	(*MsgUpdateSellOrders)(nil),            // 20: regen.ecocredit.v1alpha2.MsgUpdateSellOrders
	(*MsgUpdateSellOrdersResponse)(nil),    // 21: regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse
	(*MsgBuy)(nil),                         // 22: regen.ecocredit.v1alpha2.MsgBuy
	(*MsgBuyResponse)(nil),                 // 23: regen.ecocredit.v1alpha2.MsgBuyResponse
	(*MsgAllowAskDenom)(nil),               // 24: regen.ecocredit.v1alpha2.MsgAllowAskDenom
	(*MsgAllowAskDenomResponse)(nil),       // 25: regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse
	(*MsgCreateBasket)(nil),                // 26: regen.ecocredit.v1alpha2.MsgCreateBasket
	(*MsgCreateBasketResponse)(nil),        // 27: regen.ecocredit.v1alpha2.MsgCreateBasketResponse
	(*MsgAddToBasket)(nil),                 // 28: regen.ecocredit.v1alpha2.MsgAddToBasket
	(*MsgAddToBasketResponse)(nil),         // 29: regen.ecocredit.v1alpha2.MsgAddToBasketResponse
	(*MsgTakeFromBasket)(nil),              // 30: regen.ecocredit.v1alpha2.MsgTakeFromBasket
	(*MsgTakeFromBasketResponse)(nil),      // 31: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse
	(*MsgPickFromBasket)(nil),              // 32: regen.ecocredit.v1alpha2.MsgPickFromBasket
	(*MsgPickFromBasketResponse)(nil),      // 33: regen.ecocredit.v1alpha2.MsgPickFromBasketResponse
	(*MsgCreateBatch_BatchIssuance)(nil),   // 34: regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance
	(*MsgSend_SendCredits)(nil),            // 35: regen.ecocredit.v1alpha2.MsgSend.SendCredits
	(*MsgRetire_RetireCredits)(nil),        // 36: regen.ecocredit.v1alpha2.MsgRetire.RetireCredits
	(*MsgCancel_CancelCredits)(nil),        // 37: regen.ecocredit.v1alpha2.MsgCancel.CancelCredits
	(*MsgSell_Order)(nil),                  // 38: regen.ecocredit.v1alpha2.MsgSell.Order
	(*MsgUpdateSellOrders_Update)(nil),     // 39: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update
	(*MsgBuy_Order)(nil),                   // 40: regen.ecocredit.v1alpha2.MsgBuy.Order
	(*MsgBuy_Order_Selection)(nil),         // 41: regen.ecocredit.v1alpha2.MsgBuy.Order.Selection
	(*timestamppb.Timestamp)(nil),          // 42: google.protobuf.Timestamp
	(*BasketCriteria)(nil),                 // 43: regen.ecocredit.v1alpha2.BasketCriteria
	(*BasketCredit)(nil),                   // 44: regen.ecocredit.v1alpha2.BasketCredit
	(*v1beta1.Coin)(nil),                   // 45: cosmos.base.v1beta1.Coin
}
var file_regen_ecocredit_v1alpha2_tx_proto_depIdxs = []int32{
	34, // 0: regen.ecocredit.v1alpha2.MsgCreateBatch.issuance:type_name -> regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance
	42, // 1: regen.ecocredit.v1alpha2.MsgCreateBatch.start_date:type_name -> google.protobuf.Timestamp
	42, // 2: regen.ecocredit.v1alpha2.MsgCreateBatch.end_date:type_name -> google.protobuf.Timestamp
	35, // 3: regen.ecocredit.v1alpha2.MsgSend.credits:type_name -> regen.ecocredit.v1alpha2.MsgSend.SendCredits
	36, // 4: regen.ecocredit.v1alpha2.MsgRetire.credits:type_name -> regen.ecocredit.v1alpha2.MsgRetire.RetireCredits
	37, // 5: regen.ecocredit.v1alpha2.MsgCancel.credits:type_name -> regen.ecocredit.v1alpha2.MsgCancel.CancelCredits
	38, // 6: regen.ecocredit.v1alpha2.MsgSell.orders:type_name -> regen.ecocredit.v1alpha2.MsgSell.Order
	39, // 7: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.updates:type_name -> regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update
	40, // 8: regen.ecocredit.v1alpha2.MsgBuy.orders:type_name -> regen.ecocredit.v1alpha2.MsgBuy.Order
	43, // 9: regen.ecocredit.v1alpha2.MsgCreateBasket.basket_criteria:type_name -> regen.ecocredit.v1alpha2.BasketCriteria
	44, // 10: regen.ecocredit.v1alpha2.MsgAddToBasket.credits:type_name -> regen.ecocredit.v1alpha2.BasketCredit
	44, // 11: regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse.credits:type_name -> regen.ecocredit.v1alpha2.BasketCredit
	44, // 12: regen.ecocredit.v1alpha2.MsgPickFromBasket.credits:type_name -> regen.ecocredit.v1alpha2.BasketCredit
	45, // 13: regen.ecocredit.v1alpha2.MsgSell.Order.ask_price:type_name -> cosmos.base.v1beta1.Coin
	42, // 14: regen.ecocredit.v1alpha2.MsgSell.Order.expiration:type_name -> google.protobuf.Timestamp
	45, // 15: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_ask_price:type_name -> cosmos.base.v1beta1.Coin
	42, // 16: regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update.new_expiration:type_name -> google.protobuf.Timestamp
	41, // 17: regen.ecocredit.v1alpha2.MsgBuy.Order.selection:type_name -> regen.ecocredit.v1alpha2.MsgBuy.Order.Selection
	45, // 18: regen.ecocredit.v1alpha2.MsgBuy.Order.bid_price:type_name -> cosmos.base.v1beta1.Coin
	42, // 19: regen.ecocredit.v1alpha2.MsgBuy.Order.expiration:type_name -> google.protobuf.Timestamp
	0,  // 20: regen.ecocredit.v1alpha2.Msg.CreateClass:input_type -> regen.ecocredit.v1alpha2.MsgCreateClass
	2,  // 21: regen.ecocredit.v1alpha2.Msg.CreateProject:input_type -> regen.ecocredit.v1alpha2.MsgCreateProject
	4,  // 22: regen.ecocredit.v1alpha2.Msg.CreateBatch:input_type -> regen.ecocredit.v1alpha2.MsgCreateBatch
	6,  // 23: regen.ecocredit.v1alpha2.Msg.Send:input_type -> regen.ecocredit.v1alpha2.MsgSend
	8,  // 24: regen.ecocredit.v1alpha2.Msg.Retire:input_type -> regen.ecocredit.v1alpha2.MsgRetire
	10, // 25: regen.ecocredit.v1alpha2.Msg.Cancel:input_type -> regen.ecocredit.v1alpha2.MsgCancel
	12, // 26: regen.ecocredit.v1alpha2.Msg.UpdateClassAdmin:input_type -> regen.ecocredit.v1alpha2.MsgUpdateClassAdmin
	14, // 27: regen.ecocredit.v1alpha2.Msg.UpdateClassIssuers:input_type -> regen.ecocredit.v1alpha2.MsgUpdateClassIssuers
	16, // 28: regen.ecocredit.v1alpha2.Msg.UpdateClassMetadata:input_type -> regen.ecocredit.v1alpha2.MsgUpdateClassMetadata
	18, // 29: regen.ecocredit.v1alpha2.Msg.Sell:input_type -> regen.ecocredit.v1alpha2.MsgSell
	20, // 30: regen.ecocredit.v1alpha2.Msg.UpdateSellOrders:input_type -> regen.ecocredit.v1alpha2.MsgUpdateSellOrders
	22, // 31: regen.ecocredit.v1alpha2.Msg.Buy:input_type -> regen.ecocredit.v1alpha2.MsgBuy
	24, // 32: regen.ecocredit.v1alpha2.Msg.AllowAskDenom:input_type -> regen.ecocredit.v1alpha2.MsgAllowAskDenom
	26, // 33: regen.ecocredit.v1alpha2.Msg.CreateBasket:input_type -> regen.ecocredit.v1alpha2.MsgCreateBasket
	28, // 34: regen.ecocredit.v1alpha2.Msg.AddToBasket:input_type -> regen.ecocredit.v1alpha2.MsgAddToBasket
	30, // 35: regen.ecocredit.v1alpha2.Msg.TakeFromBasket:input_type -> regen.ecocredit.v1alpha2.MsgTakeFromBasket
	32, // 36: regen.ecocredit.v1alpha2.Msg.PickFromBasket:input_type -> regen.ecocredit.v1alpha2.MsgPickFromBasket
	1,  // 37: regen.ecocredit.v1alpha2.Msg.CreateClass:output_type -> regen.ecocredit.v1alpha2.MsgCreateClassResponse
	3,  // 38: regen.ecocredit.v1alpha2.Msg.CreateProject:output_type -> regen.ecocredit.v1alpha2.MsgCreateProjectResponse
	5,  // 39: regen.ecocredit.v1alpha2.Msg.CreateBatch:output_type -> regen.ecocredit.v1alpha2.MsgCreateBatchResponse
	7,  // 40: regen.ecocredit.v1alpha2.Msg.Send:output_type -> regen.ecocredit.v1alpha2.MsgSendResponse
	9,  // 41: regen.ecocredit.v1alpha2.Msg.Retire:output_type -> regen.ecocredit.v1alpha2.MsgRetireResponse
	11, // 42: regen.ecocredit.v1alpha2.Msg.Cancel:output_type -> regen.ecocredit.v1alpha2.MsgCancelResponse
	13, // 43: regen.ecocredit.v1alpha2.Msg.UpdateClassAdmin:output_type -> regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse
	15, // 44: regen.ecocredit.v1alpha2.Msg.UpdateClassIssuers:output_type -> regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse
	17, // 45: regen.ecocredit.v1alpha2.Msg.UpdateClassMetadata:output_type -> regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse
	19, // 46: regen.ecocredit.v1alpha2.Msg.Sell:output_type -> regen.ecocredit.v1alpha2.MsgSellResponse
	21, // 47: regen.ecocredit.v1alpha2.Msg.UpdateSellOrders:output_type -> regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse
	23, // 48: regen.ecocredit.v1alpha2.Msg.Buy:output_type -> regen.ecocredit.v1alpha2.MsgBuyResponse
	25, // 49: regen.ecocredit.v1alpha2.Msg.AllowAskDenom:output_type -> regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse
	27, // 50: regen.ecocredit.v1alpha2.Msg.CreateBasket:output_type -> regen.ecocredit.v1alpha2.MsgCreateBasketResponse
	29, // 51: regen.ecocredit.v1alpha2.Msg.AddToBasket:output_type -> regen.ecocredit.v1alpha2.MsgAddToBasketResponse
	31, // 52: regen.ecocredit.v1alpha2.Msg.TakeFromBasket:output_type -> regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse
	33, // 53: regen.ecocredit.v1alpha2.Msg.PickFromBasket:output_type -> regen.ecocredit.v1alpha2.MsgPickFromBasketResponse
	37, // [37:54] is the sub-list for method output_type
	20, // [20:37] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_v1alpha2_tx_proto_init() }
func file_regen_ecocredit_v1alpha2_tx_proto_init() {
	if File_regen_ecocredit_v1alpha2_tx_proto != nil {
		return
	}
	file_regen_ecocredit_v1alpha2_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateClass); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateClassResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateProject); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateProjectResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBatch); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBatchResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSend); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSendResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRetire); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRetireResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCancel); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCancelResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassAdmin); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassAdminResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassIssuers); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassIssuersResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassMetadata); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassMetadataResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSell); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSellResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateSellOrders); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateSellOrdersResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBuy); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBuyResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAllowAskDenom); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAllowAskDenomResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBasket); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[27].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBasketResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[28].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAddToBasket); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[29].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAddToBasketResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[30].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTakeFromBasket); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[31].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTakeFromBasketResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[32].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPickFromBasket); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[33].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPickFromBasketResponse); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[34].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBatch_BatchIssuance); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[35].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSend_SendCredits); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[36].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRetire_RetireCredits); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[37].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCancel_CancelCredits); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[38].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSell_Order); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[39].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateSellOrders_Update); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[40].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBuy_Order); i {
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
		file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[41].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBuy_Order_Selection); i {
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
	file_regen_ecocredit_v1alpha2_tx_proto_msgTypes[41].OneofWrappers = []interface{}{
		(*MsgBuy_Order_Selection_SellOrderId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_v1alpha2_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   42,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_regen_ecocredit_v1alpha2_tx_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_v1alpha2_tx_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_v1alpha2_tx_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_v1alpha2_tx_proto = out.File
	file_regen_ecocredit_v1alpha2_tx_proto_rawDesc = nil
	file_regen_ecocredit_v1alpha2_tx_proto_goTypes = nil
	file_regen_ecocredit_v1alpha2_tx_proto_depIdxs = nil
}
