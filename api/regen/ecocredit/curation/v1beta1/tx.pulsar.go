package curationv1beta1

import (
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_MsgDefineTag protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgDefineTag = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgDefineTag")
}

var _ protoreflect.Message = (*fastReflection_MsgDefineTag)(nil)

type fastReflection_MsgDefineTag MsgDefineTag

func (x *MsgDefineTag) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgDefineTag)(x)
}

func (x *MsgDefineTag) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgDefineTag_messageType fastReflection_MsgDefineTag_messageType
var _ protoreflect.MessageType = fastReflection_MsgDefineTag_messageType{}

type fastReflection_MsgDefineTag_messageType struct{}

func (x fastReflection_MsgDefineTag_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgDefineTag)(nil)
}
func (x fastReflection_MsgDefineTag_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgDefineTag)
}
func (x fastReflection_MsgDefineTag_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineTag
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgDefineTag) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineTag
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgDefineTag) Type() protoreflect.MessageType {
	return _fastReflection_MsgDefineTag_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgDefineTag) New() protoreflect.Message {
	return new(fastReflection_MsgDefineTag)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgDefineTag) Interface() protoreflect.ProtoMessage {
	return (*MsgDefineTag)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgDefineTag) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgDefineTag) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTag does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineTag) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTag does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgDefineTag) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTag does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgDefineTag) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTag does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgDefineTag) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTag does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgDefineTag) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTag does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgDefineTag) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgDefineTag", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgDefineTag) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineTag) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgDefineTag) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgDefineTag) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgDefineTag)
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
		x := input.Message.Interface().(*MsgDefineTag)
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
		x := input.Message.Interface().(*MsgDefineTag)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineTag: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineTag: illegal tag %d (wire type %d)", fieldNum, wire)
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
	md_MsgDefineTagResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgDefineTagResponse = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgDefineTagResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgDefineTagResponse)(nil)

type fastReflection_MsgDefineTagResponse MsgDefineTagResponse

func (x *MsgDefineTagResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgDefineTagResponse)(x)
}

func (x *MsgDefineTagResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgDefineTagResponse_messageType fastReflection_MsgDefineTagResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgDefineTagResponse_messageType{}

type fastReflection_MsgDefineTagResponse_messageType struct{}

func (x fastReflection_MsgDefineTagResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgDefineTagResponse)(nil)
}
func (x fastReflection_MsgDefineTagResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgDefineTagResponse)
}
func (x fastReflection_MsgDefineTagResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineTagResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgDefineTagResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineTagResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgDefineTagResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgDefineTagResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgDefineTagResponse) New() protoreflect.Message {
	return new(fastReflection_MsgDefineTagResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgDefineTagResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgDefineTagResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgDefineTagResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgDefineTagResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTagResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineTagResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTagResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgDefineTagResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTagResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgDefineTagResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTagResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgDefineTagResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTagResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgDefineTagResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineTagResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgDefineTagResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgDefineTagResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgDefineTagResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineTagResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgDefineTagResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgDefineTagResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgDefineTagResponse)
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
		x := input.Message.Interface().(*MsgDefineTagResponse)
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
		x := input.Message.Interface().(*MsgDefineTagResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineTagResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineTagResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
	md_MsgDefineNumericAttr protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgDefineNumericAttr = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgDefineNumericAttr")
}

var _ protoreflect.Message = (*fastReflection_MsgDefineNumericAttr)(nil)

type fastReflection_MsgDefineNumericAttr MsgDefineNumericAttr

func (x *MsgDefineNumericAttr) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgDefineNumericAttr)(x)
}

func (x *MsgDefineNumericAttr) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgDefineNumericAttr_messageType fastReflection_MsgDefineNumericAttr_messageType
var _ protoreflect.MessageType = fastReflection_MsgDefineNumericAttr_messageType{}

type fastReflection_MsgDefineNumericAttr_messageType struct{}

func (x fastReflection_MsgDefineNumericAttr_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgDefineNumericAttr)(nil)
}
func (x fastReflection_MsgDefineNumericAttr_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgDefineNumericAttr)
}
func (x fastReflection_MsgDefineNumericAttr_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineNumericAttr
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgDefineNumericAttr) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineNumericAttr
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgDefineNumericAttr) Type() protoreflect.MessageType {
	return _fastReflection_MsgDefineNumericAttr_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgDefineNumericAttr) New() protoreflect.Message {
	return new(fastReflection_MsgDefineNumericAttr)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgDefineNumericAttr) Interface() protoreflect.ProtoMessage {
	return (*MsgDefineNumericAttr)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgDefineNumericAttr) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgDefineNumericAttr) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineNumericAttr) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgDefineNumericAttr) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgDefineNumericAttr) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgDefineNumericAttr) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgDefineNumericAttr) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgDefineNumericAttr) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgDefineNumericAttr) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineNumericAttr) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgDefineNumericAttr) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgDefineNumericAttr) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgDefineNumericAttr)
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
		x := input.Message.Interface().(*MsgDefineNumericAttr)
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
		x := input.Message.Interface().(*MsgDefineNumericAttr)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineNumericAttr: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineNumericAttr: illegal tag %d (wire type %d)", fieldNum, wire)
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
	md_MsgDefineNumericAttrResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgDefineNumericAttrResponse = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgDefineNumericAttrResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgDefineNumericAttrResponse)(nil)

type fastReflection_MsgDefineNumericAttrResponse MsgDefineNumericAttrResponse

func (x *MsgDefineNumericAttrResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgDefineNumericAttrResponse)(x)
}

func (x *MsgDefineNumericAttrResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgDefineNumericAttrResponse_messageType fastReflection_MsgDefineNumericAttrResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgDefineNumericAttrResponse_messageType{}

type fastReflection_MsgDefineNumericAttrResponse_messageType struct{}

func (x fastReflection_MsgDefineNumericAttrResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgDefineNumericAttrResponse)(nil)
}
func (x fastReflection_MsgDefineNumericAttrResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgDefineNumericAttrResponse)
}
func (x fastReflection_MsgDefineNumericAttrResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineNumericAttrResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgDefineNumericAttrResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDefineNumericAttrResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgDefineNumericAttrResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgDefineNumericAttrResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgDefineNumericAttrResponse) New() protoreflect.Message {
	return new(fastReflection_MsgDefineNumericAttrResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgDefineNumericAttrResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgDefineNumericAttrResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgDefineNumericAttrResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgDefineNumericAttrResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineNumericAttrResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgDefineNumericAttrResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgDefineNumericAttrResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgDefineNumericAttrResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgDefineNumericAttrResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgDefineNumericAttrResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgDefineNumericAttrResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDefineNumericAttrResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgDefineNumericAttrResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgDefineNumericAttrResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgDefineNumericAttrResponse)
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
		x := input.Message.Interface().(*MsgDefineNumericAttrResponse)
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
		x := input.Message.Interface().(*MsgDefineNumericAttrResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineNumericAttrResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDefineNumericAttrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_MsgTag_2_list)(nil)

type _MsgTag_2_list struct {
	list *[]*MsgTag_Tagging
}

func (x *_MsgTag_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgTag_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgTag_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgTag_Tagging)
	(*x.list)[i] = concreteValue
}

func (x *_MsgTag_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgTag_Tagging)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgTag_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgTag_Tagging)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgTag_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgTag_2_list) NewElement() protoreflect.Value {
	v := new(MsgTag_Tagging)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgTag_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgTag          protoreflect.MessageDescriptor
	fd_MsgTag_curator  protoreflect.FieldDescriptor
	fd_MsgTag_taggings protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgTag = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgTag")
	fd_MsgTag_curator = md_MsgTag.Fields().ByName("curator")
	fd_MsgTag_taggings = md_MsgTag.Fields().ByName("taggings")
}

var _ protoreflect.Message = (*fastReflection_MsgTag)(nil)

type fastReflection_MsgTag MsgTag

func (x *MsgTag) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgTag)(x)
}

func (x *MsgTag) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgTag_messageType fastReflection_MsgTag_messageType
var _ protoreflect.MessageType = fastReflection_MsgTag_messageType{}

type fastReflection_MsgTag_messageType struct{}

func (x fastReflection_MsgTag_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgTag)(nil)
}
func (x fastReflection_MsgTag_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgTag)
}
func (x fastReflection_MsgTag_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTag
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgTag) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTag
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgTag) Type() protoreflect.MessageType {
	return _fastReflection_MsgTag_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgTag) New() protoreflect.Message {
	return new(fastReflection_MsgTag)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgTag) Interface() protoreflect.ProtoMessage {
	return (*MsgTag)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgTag) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Curator != "" {
		value := protoreflect.ValueOfString(x.Curator)
		if !f(fd_MsgTag_curator, value) {
			return
		}
	}
	if len(x.Taggings) != 0 {
		value := protoreflect.ValueOfList(&_MsgTag_2_list{list: &x.Taggings})
		if !f(fd_MsgTag_taggings, value) {
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
func (x *fastReflection_MsgTag) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.curator":
		return x.Curator != ""
	case "regen.ecocredit.curation.v1beta1.MsgTag.taggings":
		return len(x.Taggings) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTag) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.curator":
		x.Curator = ""
	case "regen.ecocredit.curation.v1beta1.MsgTag.taggings":
		x.Taggings = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgTag) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.curator":
		value := x.Curator
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.curation.v1beta1.MsgTag.taggings":
		if len(x.Taggings) == 0 {
			return protoreflect.ValueOfList(&_MsgTag_2_list{})
		}
		listValue := &_MsgTag_2_list{list: &x.Taggings}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgTag) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.curator":
		x.Curator = value.Interface().(string)
	case "regen.ecocredit.curation.v1beta1.MsgTag.taggings":
		lv := value.List()
		clv := lv.(*_MsgTag_2_list)
		x.Taggings = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgTag) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.taggings":
		if x.Taggings == nil {
			x.Taggings = []*MsgTag_Tagging{}
		}
		value := &_MsgTag_2_list{list: &x.Taggings}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.curation.v1beta1.MsgTag.curator":
		panic(fmt.Errorf("field curator of message regen.ecocredit.curation.v1beta1.MsgTag is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgTag) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.curator":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.curation.v1beta1.MsgTag.taggings":
		list := []*MsgTag_Tagging{}
		return protoreflect.ValueOfList(&_MsgTag_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgTag) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgTag", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgTag) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTag) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgTag) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgTag) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgTag)
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
		if len(x.Taggings) > 0 {
			for _, e := range x.Taggings {
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
		x := input.Message.Interface().(*MsgTag)
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
		if len(x.Taggings) > 0 {
			for iNdEx := len(x.Taggings) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Taggings[iNdEx])
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
		x := input.Message.Interface().(*MsgTag)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTag: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTag: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Taggings", wireType)
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
				x.Taggings = append(x.Taggings, &MsgTag_Tagging{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Taggings[len(x.Taggings)-1]); err != nil {
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

var _ protoreflect.List = (*_MsgTag_Tagging_2_list)(nil)

type _MsgTag_Tagging_2_list struct {
	list *[]*TagTarget
}

func (x *_MsgTag_Tagging_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgTag_Tagging_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgTag_Tagging_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*TagTarget)
	(*x.list)[i] = concreteValue
}

func (x *_MsgTag_Tagging_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*TagTarget)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgTag_Tagging_2_list) AppendMutable() protoreflect.Value {
	v := new(TagTarget)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgTag_Tagging_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgTag_Tagging_2_list) NewElement() protoreflect.Value {
	v := new(TagTarget)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgTag_Tagging_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgTag_Tagging         protoreflect.MessageDescriptor
	fd_MsgTag_Tagging_tag_id  protoreflect.FieldDescriptor
	fd_MsgTag_Tagging_targets protoreflect.FieldDescriptor
	fd_MsgTag_Tagging_untag   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgTag_Tagging = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgTag").Messages().ByName("Tagging")
	fd_MsgTag_Tagging_tag_id = md_MsgTag_Tagging.Fields().ByName("tag_id")
	fd_MsgTag_Tagging_targets = md_MsgTag_Tagging.Fields().ByName("targets")
	fd_MsgTag_Tagging_untag = md_MsgTag_Tagging.Fields().ByName("untag")
}

var _ protoreflect.Message = (*fastReflection_MsgTag_Tagging)(nil)

type fastReflection_MsgTag_Tagging MsgTag_Tagging

func (x *MsgTag_Tagging) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgTag_Tagging)(x)
}

func (x *MsgTag_Tagging) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgTag_Tagging_messageType fastReflection_MsgTag_Tagging_messageType
var _ protoreflect.MessageType = fastReflection_MsgTag_Tagging_messageType{}

type fastReflection_MsgTag_Tagging_messageType struct{}

func (x fastReflection_MsgTag_Tagging_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgTag_Tagging)(nil)
}
func (x fastReflection_MsgTag_Tagging_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgTag_Tagging)
}
func (x fastReflection_MsgTag_Tagging_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTag_Tagging
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgTag_Tagging) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTag_Tagging
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgTag_Tagging) Type() protoreflect.MessageType {
	return _fastReflection_MsgTag_Tagging_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgTag_Tagging) New() protoreflect.Message {
	return new(fastReflection_MsgTag_Tagging)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgTag_Tagging) Interface() protoreflect.ProtoMessage {
	return (*MsgTag_Tagging)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgTag_Tagging) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.TagId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TagId)
		if !f(fd_MsgTag_Tagging_tag_id, value) {
			return
		}
	}
	if len(x.Targets) != 0 {
		value := protoreflect.ValueOfList(&_MsgTag_Tagging_2_list{list: &x.Targets})
		if !f(fd_MsgTag_Tagging_targets, value) {
			return
		}
	}
	if x.Untag != false {
		value := protoreflect.ValueOfBool(x.Untag)
		if !f(fd_MsgTag_Tagging_untag, value) {
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
func (x *fastReflection_MsgTag_Tagging) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.tag_id":
		return x.TagId != uint64(0)
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.targets":
		return len(x.Targets) != 0
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.untag":
		return x.Untag != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag.Tagging"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag.Tagging does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTag_Tagging) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.tag_id":
		x.TagId = uint64(0)
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.targets":
		x.Targets = nil
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.untag":
		x.Untag = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag.Tagging"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag.Tagging does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgTag_Tagging) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.tag_id":
		value := x.TagId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.targets":
		if len(x.Targets) == 0 {
			return protoreflect.ValueOfList(&_MsgTag_Tagging_2_list{})
		}
		listValue := &_MsgTag_Tagging_2_list{list: &x.Targets}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.untag":
		value := x.Untag
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag.Tagging"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag.Tagging does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgTag_Tagging) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.tag_id":
		x.TagId = value.Uint()
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.targets":
		lv := value.List()
		clv := lv.(*_MsgTag_Tagging_2_list)
		x.Targets = *clv.list
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.untag":
		x.Untag = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag.Tagging"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag.Tagging does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgTag_Tagging) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.targets":
		if x.Targets == nil {
			x.Targets = []*TagTarget{}
		}
		value := &_MsgTag_Tagging_2_list{list: &x.Targets}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.tag_id":
		panic(fmt.Errorf("field tag_id of message regen.ecocredit.curation.v1beta1.MsgTag.Tagging is not mutable"))
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.untag":
		panic(fmt.Errorf("field untag of message regen.ecocredit.curation.v1beta1.MsgTag.Tagging is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag.Tagging"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag.Tagging does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgTag_Tagging) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.tag_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.targets":
		list := []*TagTarget{}
		return protoreflect.ValueOfList(&_MsgTag_Tagging_2_list{list: &list})
	case "regen.ecocredit.curation.v1beta1.MsgTag.Tagging.untag":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTag.Tagging"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTag.Tagging does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgTag_Tagging) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgTag.Tagging", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgTag_Tagging) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTag_Tagging) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgTag_Tagging) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgTag_Tagging) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgTag_Tagging)
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
		if x.TagId != 0 {
			n += 1 + runtime.Sov(uint64(x.TagId))
		}
		if len(x.Targets) > 0 {
			for _, e := range x.Targets {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Untag {
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
		x := input.Message.Interface().(*MsgTag_Tagging)
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
		if x.Untag {
			i--
			if x.Untag {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x18
		}
		if len(x.Targets) > 0 {
			for iNdEx := len(x.Targets) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Targets[iNdEx])
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
		if x.TagId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TagId))
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
		x := input.Message.Interface().(*MsgTag_Tagging)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTag_Tagging: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTag_Tagging: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TagId", wireType)
				}
				x.TagId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TagId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Targets", wireType)
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
				x.Targets = append(x.Targets, &TagTarget{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Targets[len(x.Targets)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Untag", wireType)
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
				x.Untag = bool(v != 0)
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
	md_MsgTagResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgTagResponse = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgTagResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgTagResponse)(nil)

type fastReflection_MsgTagResponse MsgTagResponse

func (x *MsgTagResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgTagResponse)(x)
}

func (x *MsgTagResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgTagResponse_messageType fastReflection_MsgTagResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgTagResponse_messageType{}

type fastReflection_MsgTagResponse_messageType struct{}

func (x fastReflection_MsgTagResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgTagResponse)(nil)
}
func (x fastReflection_MsgTagResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgTagResponse)
}
func (x fastReflection_MsgTagResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTagResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgTagResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgTagResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgTagResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgTagResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgTagResponse) New() protoreflect.Message {
	return new(fastReflection_MsgTagResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgTagResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgTagResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgTagResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgTagResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTagResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTagResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTagResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgTagResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTagResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgTagResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTagResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgTagResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTagResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgTagResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgTagResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgTagResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgTagResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgTagResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgTagResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgTagResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgTagResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgTagResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgTagResponse)
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
		x := input.Message.Interface().(*MsgTagResponse)
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
		x := input.Message.Interface().(*MsgTagResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTagResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgTagResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_MsgSetNumericAttr_2_list)(nil)

type _MsgSetNumericAttr_2_list struct {
	list *[]*MsgSetNumericAttr_SetAttr
}

func (x *_MsgSetNumericAttr_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgSetNumericAttr_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgSetNumericAttr_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSetNumericAttr_SetAttr)
	(*x.list)[i] = concreteValue
}

func (x *_MsgSetNumericAttr_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSetNumericAttr_SetAttr)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgSetNumericAttr_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgSetNumericAttr_SetAttr)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSetNumericAttr_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgSetNumericAttr_2_list) NewElement() protoreflect.Value {
	v := new(MsgSetNumericAttr_SetAttr)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSetNumericAttr_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgSetNumericAttr          protoreflect.MessageDescriptor
	fd_MsgSetNumericAttr_curator  protoreflect.FieldDescriptor
	fd_MsgSetNumericAttr_set_attr protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgSetNumericAttr = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgSetNumericAttr")
	fd_MsgSetNumericAttr_curator = md_MsgSetNumericAttr.Fields().ByName("curator")
	fd_MsgSetNumericAttr_set_attr = md_MsgSetNumericAttr.Fields().ByName("set_attr")
}

var _ protoreflect.Message = (*fastReflection_MsgSetNumericAttr)(nil)

type fastReflection_MsgSetNumericAttr MsgSetNumericAttr

func (x *MsgSetNumericAttr) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSetNumericAttr)(x)
}

func (x *MsgSetNumericAttr) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSetNumericAttr_messageType fastReflection_MsgSetNumericAttr_messageType
var _ protoreflect.MessageType = fastReflection_MsgSetNumericAttr_messageType{}

type fastReflection_MsgSetNumericAttr_messageType struct{}

func (x fastReflection_MsgSetNumericAttr_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSetNumericAttr)(nil)
}
func (x fastReflection_MsgSetNumericAttr_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSetNumericAttr)
}
func (x fastReflection_MsgSetNumericAttr_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSetNumericAttr
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSetNumericAttr) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSetNumericAttr
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSetNumericAttr) Type() protoreflect.MessageType {
	return _fastReflection_MsgSetNumericAttr_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSetNumericAttr) New() protoreflect.Message {
	return new(fastReflection_MsgSetNumericAttr)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSetNumericAttr) Interface() protoreflect.ProtoMessage {
	return (*MsgSetNumericAttr)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSetNumericAttr) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Curator != "" {
		value := protoreflect.ValueOfString(x.Curator)
		if !f(fd_MsgSetNumericAttr_curator, value) {
			return
		}
	}
	if len(x.SetAttr) != 0 {
		value := protoreflect.ValueOfList(&_MsgSetNumericAttr_2_list{list: &x.SetAttr})
		if !f(fd_MsgSetNumericAttr_set_attr, value) {
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
func (x *fastReflection_MsgSetNumericAttr) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.curator":
		return x.Curator != ""
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.set_attr":
		return len(x.SetAttr) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSetNumericAttr) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.curator":
		x.Curator = ""
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.set_attr":
		x.SetAttr = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSetNumericAttr) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.curator":
		value := x.Curator
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.set_attr":
		if len(x.SetAttr) == 0 {
			return protoreflect.ValueOfList(&_MsgSetNumericAttr_2_list{})
		}
		listValue := &_MsgSetNumericAttr_2_list{list: &x.SetAttr}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSetNumericAttr) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.curator":
		x.Curator = value.Interface().(string)
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.set_attr":
		lv := value.List()
		clv := lv.(*_MsgSetNumericAttr_2_list)
		x.SetAttr = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSetNumericAttr) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.set_attr":
		if x.SetAttr == nil {
			x.SetAttr = []*MsgSetNumericAttr_SetAttr{}
		}
		value := &_MsgSetNumericAttr_2_list{list: &x.SetAttr}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.curator":
		panic(fmt.Errorf("field curator of message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSetNumericAttr) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.curator":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.set_attr":
		list := []*MsgSetNumericAttr_SetAttr{}
		return protoreflect.ValueOfList(&_MsgSetNumericAttr_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSetNumericAttr) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgSetNumericAttr", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSetNumericAttr) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSetNumericAttr) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSetNumericAttr) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSetNumericAttr) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSetNumericAttr)
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
		if len(x.SetAttr) > 0 {
			for _, e := range x.SetAttr {
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
		x := input.Message.Interface().(*MsgSetNumericAttr)
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
		if len(x.SetAttr) > 0 {
			for iNdEx := len(x.SetAttr) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.SetAttr[iNdEx])
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
		x := input.Message.Interface().(*MsgSetNumericAttr)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSetNumericAttr: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSetNumericAttr: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SetAttr", wireType)
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
				x.SetAttr = append(x.SetAttr, &MsgSetNumericAttr_SetAttr{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SetAttr[len(x.SetAttr)-1]); err != nil {
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
	md_MsgSetNumericAttr_SetAttr         protoreflect.MessageDescriptor
	fd_MsgSetNumericAttr_SetAttr_attr_id protoreflect.FieldDescriptor
	fd_MsgSetNumericAttr_SetAttr_targets protoreflect.FieldDescriptor
	fd_MsgSetNumericAttr_SetAttr_value   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgSetNumericAttr_SetAttr = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgSetNumericAttr").Messages().ByName("SetAttr")
	fd_MsgSetNumericAttr_SetAttr_attr_id = md_MsgSetNumericAttr_SetAttr.Fields().ByName("attr_id")
	fd_MsgSetNumericAttr_SetAttr_targets = md_MsgSetNumericAttr_SetAttr.Fields().ByName("targets")
	fd_MsgSetNumericAttr_SetAttr_value = md_MsgSetNumericAttr_SetAttr.Fields().ByName("value")
}

var _ protoreflect.Message = (*fastReflection_MsgSetNumericAttr_SetAttr)(nil)

type fastReflection_MsgSetNumericAttr_SetAttr MsgSetNumericAttr_SetAttr

func (x *MsgSetNumericAttr_SetAttr) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSetNumericAttr_SetAttr)(x)
}

func (x *MsgSetNumericAttr_SetAttr) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSetNumericAttr_SetAttr_messageType fastReflection_MsgSetNumericAttr_SetAttr_messageType
var _ protoreflect.MessageType = fastReflection_MsgSetNumericAttr_SetAttr_messageType{}

type fastReflection_MsgSetNumericAttr_SetAttr_messageType struct{}

func (x fastReflection_MsgSetNumericAttr_SetAttr_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSetNumericAttr_SetAttr)(nil)
}
func (x fastReflection_MsgSetNumericAttr_SetAttr_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSetNumericAttr_SetAttr)
}
func (x fastReflection_MsgSetNumericAttr_SetAttr_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSetNumericAttr_SetAttr
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSetNumericAttr_SetAttr
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Type() protoreflect.MessageType {
	return _fastReflection_MsgSetNumericAttr_SetAttr_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) New() protoreflect.Message {
	return new(fastReflection_MsgSetNumericAttr_SetAttr)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Interface() protoreflect.ProtoMessage {
	return (*MsgSetNumericAttr_SetAttr)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.AttrId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AttrId)
		if !f(fd_MsgSetNumericAttr_SetAttr_attr_id, value) {
			return
		}
	}
	if x.Targets != nil {
		value := protoreflect.ValueOfMessage(x.Targets.ProtoReflect())
		if !f(fd_MsgSetNumericAttr_SetAttr_targets, value) {
			return
		}
	}
	if x.XValue != nil {
		switch o := x.XValue.(type) {
		case *MsgSetNumericAttr_SetAttr_Value:
			v := o.Value
			value := protoreflect.ValueOfUint32(v)
			if !f(fd_MsgSetNumericAttr_SetAttr_value, value) {
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
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.attr_id":
		return x.AttrId != uint64(0)
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.targets":
		return x.Targets != nil
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.value":
		if x.XValue == nil {
			return false
		} else if _, ok := x.XValue.(*MsgSetNumericAttr_SetAttr_Value); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.attr_id":
		x.AttrId = uint64(0)
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.targets":
		x.Targets = nil
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.value":
		x.XValue = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.attr_id":
		value := x.AttrId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.targets":
		value := x.Targets
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.value":
		if x.XValue == nil {
			return protoreflect.ValueOfUint32(uint32(0))
		} else if v, ok := x.XValue.(*MsgSetNumericAttr_SetAttr_Value); ok {
			return protoreflect.ValueOfUint32(v.Value)
		} else {
			return protoreflect.ValueOfUint32(uint32(0))
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.attr_id":
		x.AttrId = value.Uint()
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.targets":
		x.Targets = value.Message().Interface().(*TagTarget)
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.value":
		cv := uint32(value.Uint())
		x.XValue = &MsgSetNumericAttr_SetAttr_Value{Value: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSetNumericAttr_SetAttr) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.targets":
		if x.Targets == nil {
			x.Targets = new(TagTarget)
		}
		return protoreflect.ValueOfMessage(x.Targets.ProtoReflect())
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.attr_id":
		panic(fmt.Errorf("field attr_id of message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr is not mutable"))
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.value":
		panic(fmt.Errorf("field value of message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.attr_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.targets":
		m := new(TagTarget)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.value":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr._value":
		if x.XValue == nil {
			return nil
		}
		switch x.XValue.(type) {
		case *MsgSetNumericAttr_SetAttr_Value:
			return x.Descriptor().Fields().ByName("value")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSetNumericAttr_SetAttr) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSetNumericAttr_SetAttr) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSetNumericAttr_SetAttr)
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
		if x.AttrId != 0 {
			n += 1 + runtime.Sov(uint64(x.AttrId))
		}
		if x.Targets != nil {
			l = options.Size(x.Targets)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Value != nil {
			n += 5
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
		x := input.Message.Interface().(*MsgSetNumericAttr_SetAttr)
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
		switch x := x.XValue.(type) {
		case *MsgSetNumericAttr_SetAttr_Value:
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(*x.Value))
			i--
			dAtA[i] = 0x1d
		}
		if x.Value != nil {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(*x.Value))
			i--
			dAtA[i] = 0x1d
		}
		if x.Targets != nil {
			encoded, err := options.Marshal(x.Targets)
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
		if x.AttrId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AttrId))
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
		x := input.Message.Interface().(*MsgSetNumericAttr_SetAttr)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSetNumericAttr_SetAttr: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSetNumericAttr_SetAttr: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AttrId", wireType)
				}
				x.AttrId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AttrId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Targets", wireType)
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
				if x.Targets == nil {
					x.Targets = &TagTarget{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Targets); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
				}
				var v uint32
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
				x.Value = &v
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
	md_MsgSetNumericAttrResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_tx_proto_init()
	md_MsgSetNumericAttrResponse = File_regen_ecocredit_curation_v1beta1_tx_proto.Messages().ByName("MsgSetNumericAttrResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgSetNumericAttrResponse)(nil)

type fastReflection_MsgSetNumericAttrResponse MsgSetNumericAttrResponse

func (x *MsgSetNumericAttrResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSetNumericAttrResponse)(x)
}

func (x *MsgSetNumericAttrResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSetNumericAttrResponse_messageType fastReflection_MsgSetNumericAttrResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgSetNumericAttrResponse_messageType{}

type fastReflection_MsgSetNumericAttrResponse_messageType struct{}

func (x fastReflection_MsgSetNumericAttrResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSetNumericAttrResponse)(nil)
}
func (x fastReflection_MsgSetNumericAttrResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSetNumericAttrResponse)
}
func (x fastReflection_MsgSetNumericAttrResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSetNumericAttrResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSetNumericAttrResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSetNumericAttrResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSetNumericAttrResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgSetNumericAttrResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSetNumericAttrResponse) New() protoreflect.Message {
	return new(fastReflection_MsgSetNumericAttrResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSetNumericAttrResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgSetNumericAttrResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSetNumericAttrResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgSetNumericAttrResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSetNumericAttrResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSetNumericAttrResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSetNumericAttrResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSetNumericAttrResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSetNumericAttrResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSetNumericAttrResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSetNumericAttrResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSetNumericAttrResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSetNumericAttrResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSetNumericAttrResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSetNumericAttrResponse)
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
		x := input.Message.Interface().(*MsgSetNumericAttrResponse)
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
		x := input.Message.Interface().(*MsgSetNumericAttrResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSetNumericAttrResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSetNumericAttrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
// source: regen/ecocredit/curation/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgDefineTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDefineTag) Reset() {
	*x = MsgDefineTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDefineTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDefineTag) ProtoMessage() {}

// Deprecated: Use MsgDefineTag.ProtoReflect.Descriptor instead.
func (*MsgDefineTag) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

type MsgDefineTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDefineTagResponse) Reset() {
	*x = MsgDefineTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDefineTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDefineTagResponse) ProtoMessage() {}

// Deprecated: Use MsgDefineTagResponse.ProtoReflect.Descriptor instead.
func (*MsgDefineTagResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

type MsgDefineNumericAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDefineNumericAttr) Reset() {
	*x = MsgDefineNumericAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDefineNumericAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDefineNumericAttr) ProtoMessage() {}

// Deprecated: Use MsgDefineNumericAttr.ProtoReflect.Descriptor instead.
func (*MsgDefineNumericAttr) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

type MsgDefineNumericAttrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDefineNumericAttrResponse) Reset() {
	*x = MsgDefineNumericAttrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDefineNumericAttrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDefineNumericAttrResponse) ProtoMessage() {}

// Deprecated: Use MsgDefineNumericAttrResponse.ProtoReflect.Descriptor instead.
func (*MsgDefineNumericAttrResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

// MsgTag is the Msg/Tag request type.
type MsgTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// curator
	Curator string `protobuf:"bytes,1,opt,name=curator,proto3" json:"curator,omitempty"`
	// taggings are the tag operations.
	Taggings []*MsgTag_Tagging `protobuf:"bytes,2,rep,name=taggings,proto3" json:"taggings,omitempty"`
}

func (x *MsgTag) Reset() {
	*x = MsgTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTag) ProtoMessage() {}

// Deprecated: Use MsgTag.ProtoReflect.Descriptor instead.
func (*MsgTag) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgTag) GetCurator() string {
	if x != nil {
		return x.Curator
	}
	return ""
}

func (x *MsgTag) GetTaggings() []*MsgTag_Tagging {
	if x != nil {
		return x.Taggings
	}
	return nil
}

type MsgTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgTagResponse) Reset() {
	*x = MsgTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTagResponse) ProtoMessage() {}

// Deprecated: Use MsgTagResponse.ProtoReflect.Descriptor instead.
func (*MsgTagResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{5}
}

type MsgSetNumericAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Curator string                       `protobuf:"bytes,1,opt,name=curator,proto3" json:"curator,omitempty"`
	SetAttr []*MsgSetNumericAttr_SetAttr `protobuf:"bytes,2,rep,name=set_attr,json=setAttr,proto3" json:"set_attr,omitempty"`
}

func (x *MsgSetNumericAttr) Reset() {
	*x = MsgSetNumericAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetNumericAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetNumericAttr) ProtoMessage() {}

// Deprecated: Use MsgSetNumericAttr.ProtoReflect.Descriptor instead.
func (*MsgSetNumericAttr) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{6}
}

func (x *MsgSetNumericAttr) GetCurator() string {
	if x != nil {
		return x.Curator
	}
	return ""
}

func (x *MsgSetNumericAttr) GetSetAttr() []*MsgSetNumericAttr_SetAttr {
	if x != nil {
		return x.SetAttr
	}
	return nil
}

type MsgSetNumericAttrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSetNumericAttrResponse) Reset() {
	*x = MsgSetNumericAttrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetNumericAttrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetNumericAttrResponse) ProtoMessage() {}

// Deprecated: Use MsgSetNumericAttrResponse.ProtoReflect.Descriptor instead.
func (*MsgSetNumericAttrResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{7}
}

// Tagging is a tag operations.
type MsgTag_Tagging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId uint64 `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	// targets are the tag operation targets.
	Targets []*TagTarget `protobuf:"bytes,2,rep,name=targets,proto3" json:"targets,omitempty"`
	// untag indicates that a tag should be removed if set to true.
	Untag bool `protobuf:"varint,3,opt,name=untag,proto3" json:"untag,omitempty"`
}

func (x *MsgTag_Tagging) Reset() {
	*x = MsgTag_Tagging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTag_Tagging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTag_Tagging) ProtoMessage() {}

// Deprecated: Use MsgTag_Tagging.ProtoReflect.Descriptor instead.
func (*MsgTag_Tagging) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{4, 0}
}

func (x *MsgTag_Tagging) GetTagId() uint64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *MsgTag_Tagging) GetTargets() []*TagTarget {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *MsgTag_Tagging) GetUntag() bool {
	if x != nil {
		return x.Untag
	}
	return false
}

type MsgSetNumericAttr_SetAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrId  uint64     `protobuf:"varint,1,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
	Targets *TagTarget `protobuf:"bytes,2,opt,name=targets,proto3" json:"targets,omitempty"`
	// value should be unset to delete
	Value *uint32 `protobuf:"fixed32,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *MsgSetNumericAttr_SetAttr) Reset() {
	*x = MsgSetNumericAttr_SetAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetNumericAttr_SetAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetNumericAttr_SetAttr) ProtoMessage() {}

// Deprecated: Use MsgSetNumericAttr_SetAttr.ProtoReflect.Descriptor instead.
func (*MsgSetNumericAttr_SetAttr) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP(), []int{6, 0}
}

func (x *MsgSetNumericAttr_SetAttr) GetAttrId() uint64 {
	if x != nil {
		return x.AttrId
	}
	return 0
}

func (x *MsgSetNumericAttr_SetAttr) GetTargets() *TagTarget {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *MsgSetNumericAttr_SetAttr) GetValue() uint32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

var File_regen_ecocredit_curation_v1beta1_tx_proto protoreflect.FileDescriptor

var file_regen_ecocredit_curation_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2c, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x63,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x67, 0x22, 0x16, 0x0a, 0x14, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x22, 0x1e, 0x0a, 0x1c, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xef, 0x01, 0x0a, 0x06,
	0x4d, 0x73, 0x67, 0x54, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x4c, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x7d,
	0x0a, 0x07, 0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64,
	0x12, 0x45, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x07,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x74, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x61, 0x67, 0x22, 0x10, 0x0a,
	0x0e, 0x4d, 0x73, 0x67, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x96, 0x02, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x41, 0x74, 0x74, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x56, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x07,
	0x73, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x1a, 0x8e, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x74, 0x74, 0x72, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x07,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x54, 0x61, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x07, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x4d, 0x73, 0x67, 0x53,
	0x65, 0x74, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf7, 0x03, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x75, 0x0a,
	0x09, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x67, 0x12, 0x2e, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x67, 0x1a, 0x36, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x41, 0x74, 0x74, 0x12, 0x36, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x41, 0x74, 0x74,
	0x72, 0x1a, 0x3e, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x54, 0x61, 0x67, 0x1a, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74,
	0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x12, 0x33, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d,
	0x73, 0x67, 0x53, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x41, 0x74, 0x74, 0x72,
	0x1a, 0x3b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0xae, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x07, 0x54, 0x78, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x63,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x45, 0x43, 0xaa, 0x02, 0x20, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x52, 0x65, 0x67, 0x65, 0x6e,
	0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x43, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2c, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x43, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x43,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescOnce sync.Once
	file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescData = file_regen_ecocredit_curation_v1beta1_tx_proto_rawDesc
)

func file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescData)
	})
	return file_regen_ecocredit_curation_v1beta1_tx_proto_rawDescData
}

var file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_regen_ecocredit_curation_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgDefineTag)(nil),                 // 0: regen.ecocredit.curation.v1beta1.MsgDefineTag
	(*MsgDefineTagResponse)(nil),         // 1: regen.ecocredit.curation.v1beta1.MsgDefineTagResponse
	(*MsgDefineNumericAttr)(nil),         // 2: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr
	(*MsgDefineNumericAttrResponse)(nil), // 3: regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse
	(*MsgTag)(nil),                       // 4: regen.ecocredit.curation.v1beta1.MsgTag
	(*MsgTagResponse)(nil),               // 5: regen.ecocredit.curation.v1beta1.MsgTagResponse
	(*MsgSetNumericAttr)(nil),            // 6: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr
	(*MsgSetNumericAttrResponse)(nil),    // 7: regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse
	(*MsgTag_Tagging)(nil),               // 8: regen.ecocredit.curation.v1beta1.MsgTag.Tagging
	(*MsgSetNumericAttr_SetAttr)(nil),    // 9: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr
	(*TagTarget)(nil),                    // 10: regen.ecocredit.curation.v1beta1.TagTarget
}
var file_regen_ecocredit_curation_v1beta1_tx_proto_depIdxs = []int32{
	8,  // 0: regen.ecocredit.curation.v1beta1.MsgTag.taggings:type_name -> regen.ecocredit.curation.v1beta1.MsgTag.Tagging
	9,  // 1: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.set_attr:type_name -> regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr
	10, // 2: regen.ecocredit.curation.v1beta1.MsgTag.Tagging.targets:type_name -> regen.ecocredit.curation.v1beta1.TagTarget
	10, // 3: regen.ecocredit.curation.v1beta1.MsgSetNumericAttr.SetAttr.targets:type_name -> regen.ecocredit.curation.v1beta1.TagTarget
	0,  // 4: regen.ecocredit.curation.v1beta1.Msg.DefineTag:input_type -> regen.ecocredit.curation.v1beta1.MsgDefineTag
	2,  // 5: regen.ecocredit.curation.v1beta1.Msg.DefineNumericAtt:input_type -> regen.ecocredit.curation.v1beta1.MsgDefineNumericAttr
	4,  // 6: regen.ecocredit.curation.v1beta1.Msg.Tag:input_type -> regen.ecocredit.curation.v1beta1.MsgTag
	6,  // 7: regen.ecocredit.curation.v1beta1.Msg.SetNumericAttr:input_type -> regen.ecocredit.curation.v1beta1.MsgSetNumericAttr
	1,  // 8: regen.ecocredit.curation.v1beta1.Msg.DefineTag:output_type -> regen.ecocredit.curation.v1beta1.MsgDefineTagResponse
	3,  // 9: regen.ecocredit.curation.v1beta1.Msg.DefineNumericAtt:output_type -> regen.ecocredit.curation.v1beta1.MsgDefineNumericAttrResponse
	5,  // 10: regen.ecocredit.curation.v1beta1.Msg.Tag:output_type -> regen.ecocredit.curation.v1beta1.MsgTagResponse
	7,  // 11: regen.ecocredit.curation.v1beta1.Msg.SetNumericAttr:output_type -> regen.ecocredit.curation.v1beta1.MsgSetNumericAttrResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_curation_v1beta1_tx_proto_init() }
func file_regen_ecocredit_curation_v1beta1_tx_proto_init() {
	if File_regen_ecocredit_curation_v1beta1_tx_proto != nil {
		return
	}
	file_regen_ecocredit_curation_v1beta1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDefineTag); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDefineTagResponse); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDefineNumericAttr); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDefineNumericAttrResponse); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTag); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTagResponse); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetNumericAttr); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetNumericAttrResponse); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTag_Tagging); i {
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
		file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetNumericAttr_SetAttr); i {
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
	file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_curation_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_regen_ecocredit_curation_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_curation_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_curation_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_curation_v1beta1_tx_proto = out.File
	file_regen_ecocredit_curation_v1beta1_tx_proto_rawDesc = nil
	file_regen_ecocredit_curation_v1beta1_tx_proto_goTypes = nil
	file_regen_ecocredit_curation_v1beta1_tx_proto_depIdxs = nil
}
