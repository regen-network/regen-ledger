package curationv1beta1

import (
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_TagDefinition               protoreflect.MessageDescriptor
	fd_TagDefinition_id            protoreflect.FieldDescriptor
	fd_TagDefinition_owner_address protoreflect.FieldDescriptor
	fd_TagDefinition_metadata      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_state_proto_init()
	md_TagDefinition = File_regen_ecocredit_curation_v1beta1_state_proto.Messages().ByName("TagDefinition")
	fd_TagDefinition_id = md_TagDefinition.Fields().ByName("id")
	fd_TagDefinition_owner_address = md_TagDefinition.Fields().ByName("owner_address")
	fd_TagDefinition_metadata = md_TagDefinition.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_TagDefinition)(nil)

type fastReflection_TagDefinition TagDefinition

func (x *TagDefinition) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TagDefinition)(x)
}

func (x *TagDefinition) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TagDefinition_messageType fastReflection_TagDefinition_messageType
var _ protoreflect.MessageType = fastReflection_TagDefinition_messageType{}

type fastReflection_TagDefinition_messageType struct{}

func (x fastReflection_TagDefinition_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TagDefinition)(nil)
}
func (x fastReflection_TagDefinition_messageType) New() protoreflect.Message {
	return new(fastReflection_TagDefinition)
}
func (x fastReflection_TagDefinition_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TagDefinition
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TagDefinition) Descriptor() protoreflect.MessageDescriptor {
	return md_TagDefinition
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TagDefinition) Type() protoreflect.MessageType {
	return _fastReflection_TagDefinition_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TagDefinition) New() protoreflect.Message {
	return new(fastReflection_TagDefinition)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TagDefinition) Interface() protoreflect.ProtoMessage {
	return (*TagDefinition)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TagDefinition) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_TagDefinition_id, value) {
			return
		}
	}
	if len(x.OwnerAddress) != 0 {
		value := protoreflect.ValueOfBytes(x.OwnerAddress)
		if !f(fd_TagDefinition_owner_address, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_TagDefinition_metadata, value) {
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
func (x *fastReflection_TagDefinition) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagDefinition.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.curation.v1beta1.TagDefinition.owner_address":
		return len(x.OwnerAddress) != 0
	case "regen.ecocredit.curation.v1beta1.TagDefinition.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagDefinition does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TagDefinition) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagDefinition.id":
		x.Id = uint64(0)
	case "regen.ecocredit.curation.v1beta1.TagDefinition.owner_address":
		x.OwnerAddress = nil
	case "regen.ecocredit.curation.v1beta1.TagDefinition.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagDefinition does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TagDefinition) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagDefinition.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.TagDefinition.owner_address":
		value := x.OwnerAddress
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.curation.v1beta1.TagDefinition.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagDefinition does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TagDefinition) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagDefinition.id":
		x.Id = value.Uint()
	case "regen.ecocredit.curation.v1beta1.TagDefinition.owner_address":
		x.OwnerAddress = value.Bytes()
	case "regen.ecocredit.curation.v1beta1.TagDefinition.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagDefinition does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TagDefinition) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagDefinition.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.curation.v1beta1.TagDefinition is not mutable"))
	case "regen.ecocredit.curation.v1beta1.TagDefinition.owner_address":
		panic(fmt.Errorf("field owner_address of message regen.ecocredit.curation.v1beta1.TagDefinition is not mutable"))
	case "regen.ecocredit.curation.v1beta1.TagDefinition.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.curation.v1beta1.TagDefinition is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagDefinition does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TagDefinition) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagDefinition.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.TagDefinition.owner_address":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.curation.v1beta1.TagDefinition.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagDefinition does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TagDefinition) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.TagDefinition", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TagDefinition) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TagDefinition) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TagDefinition) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TagDefinition) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TagDefinition)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.OwnerAddress)
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
		x := input.Message.Interface().(*TagDefinition)
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
		if len(x.OwnerAddress) > 0 {
			i -= len(x.OwnerAddress)
			copy(dAtA[i:], x.OwnerAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.OwnerAddress)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*TagDefinition)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TagDefinition: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TagDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
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
				x.OwnerAddress = append(x.OwnerAddress[:0], dAtA[iNdEx:postIndex]...)
				if x.OwnerAddress == nil {
					x.OwnerAddress = []byte{}
				}
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
	md_NumericAttributeDefinition                    protoreflect.MessageDescriptor
	fd_NumericAttributeDefinition_id                 protoreflect.FieldDescriptor
	fd_NumericAttributeDefinition_owner_address      protoreflect.FieldDescriptor
	fd_NumericAttributeDefinition_metadata           protoreflect.FieldDescriptor
	fd_NumericAttributeDefinition_max_decimal_places protoreflect.FieldDescriptor
	fd_NumericAttributeDefinition_min                protoreflect.FieldDescriptor
	fd_NumericAttributeDefinition_max                protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_state_proto_init()
	md_NumericAttributeDefinition = File_regen_ecocredit_curation_v1beta1_state_proto.Messages().ByName("NumericAttributeDefinition")
	fd_NumericAttributeDefinition_id = md_NumericAttributeDefinition.Fields().ByName("id")
	fd_NumericAttributeDefinition_owner_address = md_NumericAttributeDefinition.Fields().ByName("owner_address")
	fd_NumericAttributeDefinition_metadata = md_NumericAttributeDefinition.Fields().ByName("metadata")
	fd_NumericAttributeDefinition_max_decimal_places = md_NumericAttributeDefinition.Fields().ByName("max_decimal_places")
	fd_NumericAttributeDefinition_min = md_NumericAttributeDefinition.Fields().ByName("min")
	fd_NumericAttributeDefinition_max = md_NumericAttributeDefinition.Fields().ByName("max")
}

var _ protoreflect.Message = (*fastReflection_NumericAttributeDefinition)(nil)

type fastReflection_NumericAttributeDefinition NumericAttributeDefinition

func (x *NumericAttributeDefinition) ProtoReflect() protoreflect.Message {
	return (*fastReflection_NumericAttributeDefinition)(x)
}

func (x *NumericAttributeDefinition) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_NumericAttributeDefinition_messageType fastReflection_NumericAttributeDefinition_messageType
var _ protoreflect.MessageType = fastReflection_NumericAttributeDefinition_messageType{}

type fastReflection_NumericAttributeDefinition_messageType struct{}

func (x fastReflection_NumericAttributeDefinition_messageType) Zero() protoreflect.Message {
	return (*fastReflection_NumericAttributeDefinition)(nil)
}
func (x fastReflection_NumericAttributeDefinition_messageType) New() protoreflect.Message {
	return new(fastReflection_NumericAttributeDefinition)
}
func (x fastReflection_NumericAttributeDefinition_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_NumericAttributeDefinition
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_NumericAttributeDefinition) Descriptor() protoreflect.MessageDescriptor {
	return md_NumericAttributeDefinition
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_NumericAttributeDefinition) Type() protoreflect.MessageType {
	return _fastReflection_NumericAttributeDefinition_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_NumericAttributeDefinition) New() protoreflect.Message {
	return new(fastReflection_NumericAttributeDefinition)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_NumericAttributeDefinition) Interface() protoreflect.ProtoMessage {
	return (*NumericAttributeDefinition)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_NumericAttributeDefinition) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_NumericAttributeDefinition_id, value) {
			return
		}
	}
	if len(x.OwnerAddress) != 0 {
		value := protoreflect.ValueOfBytes(x.OwnerAddress)
		if !f(fd_NumericAttributeDefinition_owner_address, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_NumericAttributeDefinition_metadata, value) {
			return
		}
	}
	if x.MaxDecimalPlaces != uint32(0) {
		value := protoreflect.ValueOfUint32(x.MaxDecimalPlaces)
		if !f(fd_NumericAttributeDefinition_max_decimal_places, value) {
			return
		}
	}
	if x.XMin != nil {
		switch o := x.XMin.(type) {
		case *NumericAttributeDefinition_Min:
			v := o.Min
			value := protoreflect.ValueOfInt32(v)
			if !f(fd_NumericAttributeDefinition_min, value) {
				return
			}
		}
	}
	if x.XMax != nil {
		switch o := x.XMax.(type) {
		case *NumericAttributeDefinition_Max:
			v := o.Max
			value := protoreflect.ValueOfInt32(v)
			if !f(fd_NumericAttributeDefinition_max, value) {
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
func (x *fastReflection_NumericAttributeDefinition) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.owner_address":
		return len(x.OwnerAddress) != 0
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max_decimal_places":
		return x.MaxDecimalPlaces != uint32(0)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.min":
		if x.XMin == nil {
			return false
		} else if _, ok := x.XMin.(*NumericAttributeDefinition_Min); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max":
		if x.XMax == nil {
			return false
		} else if _, ok := x.XMax.(*NumericAttributeDefinition_Max); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NumericAttributeDefinition) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.id":
		x.Id = uint64(0)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.owner_address":
		x.OwnerAddress = nil
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.metadata":
		x.Metadata = nil
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max_decimal_places":
		x.MaxDecimalPlaces = uint32(0)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.min":
		x.XMin = nil
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max":
		x.XMax = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_NumericAttributeDefinition) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.owner_address":
		value := x.OwnerAddress
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max_decimal_places":
		value := x.MaxDecimalPlaces
		return protoreflect.ValueOfUint32(value)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.min":
		if x.XMin == nil {
			return protoreflect.ValueOfInt32(int32(0))
		} else if v, ok := x.XMin.(*NumericAttributeDefinition_Min); ok {
			return protoreflect.ValueOfInt32(v.Min)
		} else {
			return protoreflect.ValueOfInt32(int32(0))
		}
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max":
		if x.XMax == nil {
			return protoreflect.ValueOfInt32(int32(0))
		} else if v, ok := x.XMax.(*NumericAttributeDefinition_Max); ok {
			return protoreflect.ValueOfInt32(v.Max)
		} else {
			return protoreflect.ValueOfInt32(int32(0))
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_NumericAttributeDefinition) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.id":
		x.Id = value.Uint()
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.owner_address":
		x.OwnerAddress = value.Bytes()
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max_decimal_places":
		x.MaxDecimalPlaces = uint32(value.Uint())
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.min":
		cv := int32(value.Int())
		x.XMin = &NumericAttributeDefinition_Min{Min: cv}
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max":
		cv := int32(value.Int())
		x.XMax = &NumericAttributeDefinition_Max{Max: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition does not contain field %s", fd.FullName()))
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
func (x *fastReflection_NumericAttributeDefinition) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.owner_address":
		panic(fmt.Errorf("field owner_address of message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max_decimal_places":
		panic(fmt.Errorf("field max_decimal_places of message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.min":
		panic(fmt.Errorf("field min of message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max":
		panic(fmt.Errorf("field max of message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_NumericAttributeDefinition) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.owner_address":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max_decimal_places":
		return protoreflect.ValueOfUint32(uint32(0))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.min":
		return protoreflect.ValueOfInt32(int32(0))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition.max":
		return protoreflect.ValueOfInt32(int32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeDefinition"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeDefinition does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_NumericAttributeDefinition) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition._min":
		if x.XMin == nil {
			return nil
		}
		switch x.XMin.(type) {
		case *NumericAttributeDefinition_Min:
			return x.Descriptor().Fields().ByName("min")
		}
	case "regen.ecocredit.curation.v1beta1.NumericAttributeDefinition._max":
		if x.XMax == nil {
			return nil
		}
		switch x.XMax.(type) {
		case *NumericAttributeDefinition_Max:
			return x.Descriptor().Fields().ByName("max")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.NumericAttributeDefinition", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_NumericAttributeDefinition) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NumericAttributeDefinition) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_NumericAttributeDefinition) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_NumericAttributeDefinition) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*NumericAttributeDefinition)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.OwnerAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxDecimalPlaces != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxDecimalPlaces))
		}
		if x.Min != nil {
			n += 1 + runtime.Sov(uint64(*x.Min))
		}
		if x.Max != nil {
			n += 1 + runtime.Sov(uint64(*x.Max))
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
		x := input.Message.Interface().(*NumericAttributeDefinition)
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
		switch x := x.XMax.(type) {
		case *NumericAttributeDefinition_Max:
			i = runtime.EncodeVarint(dAtA, i, uint64(*x.Max))
			i--
			dAtA[i] = 0x30
		}
		switch x := x.XMin.(type) {
		case *NumericAttributeDefinition_Min:
			i = runtime.EncodeVarint(dAtA, i, uint64(*x.Min))
			i--
			dAtA[i] = 0x28
		}
		if x.Max != nil {
			i = runtime.EncodeVarint(dAtA, i, uint64(*x.Max))
			i--
			dAtA[i] = 0x30
		}
		if x.Min != nil {
			i = runtime.EncodeVarint(dAtA, i, uint64(*x.Min))
			i--
			dAtA[i] = 0x28
		}
		if x.MaxDecimalPlaces != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxDecimalPlaces))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.OwnerAddress) > 0 {
			i -= len(x.OwnerAddress)
			copy(dAtA[i:], x.OwnerAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.OwnerAddress)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*NumericAttributeDefinition)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NumericAttributeDefinition: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NumericAttributeDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
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
				x.OwnerAddress = append(x.OwnerAddress[:0], dAtA[iNdEx:postIndex]...)
				if x.OwnerAddress == nil {
					x.OwnerAddress = []byte{}
				}
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
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxDecimalPlaces", wireType)
				}
				x.MaxDecimalPlaces = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxDecimalPlaces |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
				}
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Min = &v
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
				}
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Max = &v
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
	md_ClassTagEntry          protoreflect.MessageDescriptor
	fd_ClassTagEntry_class_id protoreflect.FieldDescriptor
	fd_ClassTagEntry_tag_id   protoreflect.FieldDescriptor
	fd_ClassTagEntry_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_state_proto_init()
	md_ClassTagEntry = File_regen_ecocredit_curation_v1beta1_state_proto.Messages().ByName("ClassTagEntry")
	fd_ClassTagEntry_class_id = md_ClassTagEntry.Fields().ByName("class_id")
	fd_ClassTagEntry_tag_id = md_ClassTagEntry.Fields().ByName("tag_id")
	fd_ClassTagEntry_metadata = md_ClassTagEntry.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_ClassTagEntry)(nil)

type fastReflection_ClassTagEntry ClassTagEntry

func (x *ClassTagEntry) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ClassTagEntry)(x)
}

func (x *ClassTagEntry) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ClassTagEntry_messageType fastReflection_ClassTagEntry_messageType
var _ protoreflect.MessageType = fastReflection_ClassTagEntry_messageType{}

type fastReflection_ClassTagEntry_messageType struct{}

func (x fastReflection_ClassTagEntry_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ClassTagEntry)(nil)
}
func (x fastReflection_ClassTagEntry_messageType) New() protoreflect.Message {
	return new(fastReflection_ClassTagEntry)
}
func (x fastReflection_ClassTagEntry_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassTagEntry
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ClassTagEntry) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassTagEntry
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ClassTagEntry) Type() protoreflect.MessageType {
	return _fastReflection_ClassTagEntry_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ClassTagEntry) New() protoreflect.Message {
	return new(fastReflection_ClassTagEntry)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ClassTagEntry) Interface() protoreflect.ProtoMessage {
	return (*ClassTagEntry)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ClassTagEntry) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ClassId)
		if !f(fd_ClassTagEntry_class_id, value) {
			return
		}
	}
	if x.TagId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TagId)
		if !f(fd_ClassTagEntry_tag_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_ClassTagEntry_metadata, value) {
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
func (x *fastReflection_ClassTagEntry) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.class_id":
		return x.ClassId != uint64(0)
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.tag_id":
		return x.TagId != uint64(0)
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.ClassTagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.ClassTagEntry does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassTagEntry) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.class_id":
		x.ClassId = uint64(0)
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.tag_id":
		x.TagId = uint64(0)
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.ClassTagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.ClassTagEntry does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ClassTagEntry) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.class_id":
		value := x.ClassId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.tag_id":
		value := x.TagId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.ClassTagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.ClassTagEntry does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ClassTagEntry) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.class_id":
		x.ClassId = value.Uint()
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.tag_id":
		x.TagId = value.Uint()
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.ClassTagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.ClassTagEntry does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ClassTagEntry) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.curation.v1beta1.ClassTagEntry is not mutable"))
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.tag_id":
		panic(fmt.Errorf("field tag_id of message regen.ecocredit.curation.v1beta1.ClassTagEntry is not mutable"))
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.curation.v1beta1.ClassTagEntry is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.ClassTagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.ClassTagEntry does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ClassTagEntry) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.class_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.tag_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.ClassTagEntry.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.ClassTagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.ClassTagEntry does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ClassTagEntry) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.ClassTagEntry", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ClassTagEntry) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassTagEntry) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ClassTagEntry) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ClassTagEntry) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ClassTagEntry)
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
		if x.ClassId != 0 {
			n += 1 + runtime.Sov(uint64(x.ClassId))
		}
		if x.TagId != 0 {
			n += 1 + runtime.Sov(uint64(x.TagId))
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
		x := input.Message.Interface().(*ClassTagEntry)
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
		if x.TagId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TagId))
			i--
			dAtA[i] = 0x10
		}
		if x.ClassId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ClassId))
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
		x := input.Message.Interface().(*ClassTagEntry)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassTagEntry: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassTagEntry: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				x.ClassId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ClassId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
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
	md_TagEntry          protoreflect.MessageDescriptor
	fd_TagEntry_target   protoreflect.FieldDescriptor
	fd_TagEntry_tag_id   protoreflect.FieldDescriptor
	fd_TagEntry_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_state_proto_init()
	md_TagEntry = File_regen_ecocredit_curation_v1beta1_state_proto.Messages().ByName("TagEntry")
	fd_TagEntry_target = md_TagEntry.Fields().ByName("target")
	fd_TagEntry_tag_id = md_TagEntry.Fields().ByName("tag_id")
	fd_TagEntry_metadata = md_TagEntry.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_TagEntry)(nil)

type fastReflection_TagEntry TagEntry

func (x *TagEntry) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TagEntry)(x)
}

func (x *TagEntry) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TagEntry_messageType fastReflection_TagEntry_messageType
var _ protoreflect.MessageType = fastReflection_TagEntry_messageType{}

type fastReflection_TagEntry_messageType struct{}

func (x fastReflection_TagEntry_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TagEntry)(nil)
}
func (x fastReflection_TagEntry_messageType) New() protoreflect.Message {
	return new(fastReflection_TagEntry)
}
func (x fastReflection_TagEntry_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TagEntry
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TagEntry) Descriptor() protoreflect.MessageDescriptor {
	return md_TagEntry
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TagEntry) Type() protoreflect.MessageType {
	return _fastReflection_TagEntry_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TagEntry) New() protoreflect.Message {
	return new(fastReflection_TagEntry)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TagEntry) Interface() protoreflect.ProtoMessage {
	return (*TagEntry)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TagEntry) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Target != "" {
		value := protoreflect.ValueOfString(x.Target)
		if !f(fd_TagEntry_target, value) {
			return
		}
	}
	if x.TagId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TagId)
		if !f(fd_TagEntry_tag_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_TagEntry_metadata, value) {
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
func (x *fastReflection_TagEntry) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagEntry.target":
		return x.Target != ""
	case "regen.ecocredit.curation.v1beta1.TagEntry.tag_id":
		return x.TagId != uint64(0)
	case "regen.ecocredit.curation.v1beta1.TagEntry.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagEntry does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TagEntry) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagEntry.target":
		x.Target = ""
	case "regen.ecocredit.curation.v1beta1.TagEntry.tag_id":
		x.TagId = uint64(0)
	case "regen.ecocredit.curation.v1beta1.TagEntry.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagEntry does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TagEntry) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagEntry.target":
		value := x.Target
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.curation.v1beta1.TagEntry.tag_id":
		value := x.TagId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.TagEntry.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagEntry does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TagEntry) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagEntry.target":
		x.Target = value.Interface().(string)
	case "regen.ecocredit.curation.v1beta1.TagEntry.tag_id":
		x.TagId = value.Uint()
	case "regen.ecocredit.curation.v1beta1.TagEntry.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagEntry does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TagEntry) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagEntry.target":
		panic(fmt.Errorf("field target of message regen.ecocredit.curation.v1beta1.TagEntry is not mutable"))
	case "regen.ecocredit.curation.v1beta1.TagEntry.tag_id":
		panic(fmt.Errorf("field tag_id of message regen.ecocredit.curation.v1beta1.TagEntry is not mutable"))
	case "regen.ecocredit.curation.v1beta1.TagEntry.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.curation.v1beta1.TagEntry is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagEntry does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TagEntry) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagEntry.target":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.curation.v1beta1.TagEntry.tag_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.TagEntry.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagEntry does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TagEntry) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.TagEntry", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TagEntry) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TagEntry) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TagEntry) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TagEntry) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TagEntry)
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
		l = len(x.Target)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TagId != 0 {
			n += 1 + runtime.Sov(uint64(x.TagId))
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
		x := input.Message.Interface().(*TagEntry)
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
		if x.TagId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TagId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Target) > 0 {
			i -= len(x.Target)
			copy(dAtA[i:], x.Target)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Target)))
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
		x := input.Message.Interface().(*TagEntry)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TagEntry: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TagEntry: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
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
				x.Target = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
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
	md_NumericAttributeEntry          protoreflect.MessageDescriptor
	fd_NumericAttributeEntry_target   protoreflect.FieldDescriptor
	fd_NumericAttributeEntry_attr_id  protoreflect.FieldDescriptor
	fd_NumericAttributeEntry_metadata protoreflect.FieldDescriptor
	fd_NumericAttributeEntry_value    protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_state_proto_init()
	md_NumericAttributeEntry = File_regen_ecocredit_curation_v1beta1_state_proto.Messages().ByName("NumericAttributeEntry")
	fd_NumericAttributeEntry_target = md_NumericAttributeEntry.Fields().ByName("target")
	fd_NumericAttributeEntry_attr_id = md_NumericAttributeEntry.Fields().ByName("attr_id")
	fd_NumericAttributeEntry_metadata = md_NumericAttributeEntry.Fields().ByName("metadata")
	fd_NumericAttributeEntry_value = md_NumericAttributeEntry.Fields().ByName("value")
}

var _ protoreflect.Message = (*fastReflection_NumericAttributeEntry)(nil)

type fastReflection_NumericAttributeEntry NumericAttributeEntry

func (x *NumericAttributeEntry) ProtoReflect() protoreflect.Message {
	return (*fastReflection_NumericAttributeEntry)(x)
}

func (x *NumericAttributeEntry) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_NumericAttributeEntry_messageType fastReflection_NumericAttributeEntry_messageType
var _ protoreflect.MessageType = fastReflection_NumericAttributeEntry_messageType{}

type fastReflection_NumericAttributeEntry_messageType struct{}

func (x fastReflection_NumericAttributeEntry_messageType) Zero() protoreflect.Message {
	return (*fastReflection_NumericAttributeEntry)(nil)
}
func (x fastReflection_NumericAttributeEntry_messageType) New() protoreflect.Message {
	return new(fastReflection_NumericAttributeEntry)
}
func (x fastReflection_NumericAttributeEntry_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_NumericAttributeEntry
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_NumericAttributeEntry) Descriptor() protoreflect.MessageDescriptor {
	return md_NumericAttributeEntry
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_NumericAttributeEntry) Type() protoreflect.MessageType {
	return _fastReflection_NumericAttributeEntry_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_NumericAttributeEntry) New() protoreflect.Message {
	return new(fastReflection_NumericAttributeEntry)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_NumericAttributeEntry) Interface() protoreflect.ProtoMessage {
	return (*NumericAttributeEntry)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_NumericAttributeEntry) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Target != "" {
		value := protoreflect.ValueOfString(x.Target)
		if !f(fd_NumericAttributeEntry_target, value) {
			return
		}
	}
	if x.AttrId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AttrId)
		if !f(fd_NumericAttributeEntry_attr_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_NumericAttributeEntry_metadata, value) {
			return
		}
	}
	if x.Value != int32(0) {
		value := protoreflect.ValueOfInt32(x.Value)
		if !f(fd_NumericAttributeEntry_value, value) {
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
func (x *fastReflection_NumericAttributeEntry) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.target":
		return x.Target != ""
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.attr_id":
		return x.AttrId != uint64(0)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.value":
		return x.Value != int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeEntry does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NumericAttributeEntry) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.target":
		x.Target = ""
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.attr_id":
		x.AttrId = uint64(0)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.metadata":
		x.Metadata = nil
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.value":
		x.Value = int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeEntry does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_NumericAttributeEntry) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.target":
		value := x.Target
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.attr_id":
		value := x.AttrId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.value":
		value := x.Value
		return protoreflect.ValueOfInt32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeEntry does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_NumericAttributeEntry) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.target":
		x.Target = value.Interface().(string)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.attr_id":
		x.AttrId = value.Uint()
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.value":
		x.Value = int32(value.Int())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeEntry does not contain field %s", fd.FullName()))
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
func (x *fastReflection_NumericAttributeEntry) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.target":
		panic(fmt.Errorf("field target of message regen.ecocredit.curation.v1beta1.NumericAttributeEntry is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.attr_id":
		panic(fmt.Errorf("field attr_id of message regen.ecocredit.curation.v1beta1.NumericAttributeEntry is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.curation.v1beta1.NumericAttributeEntry is not mutable"))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.value":
		panic(fmt.Errorf("field value of message regen.ecocredit.curation.v1beta1.NumericAttributeEntry is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeEntry does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_NumericAttributeEntry) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.target":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.attr_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.curation.v1beta1.NumericAttributeEntry.value":
		return protoreflect.ValueOfInt32(int32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.NumericAttributeEntry"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.NumericAttributeEntry does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_NumericAttributeEntry) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.NumericAttributeEntry", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_NumericAttributeEntry) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NumericAttributeEntry) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_NumericAttributeEntry) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_NumericAttributeEntry) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*NumericAttributeEntry)
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
		l = len(x.Target)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AttrId != 0 {
			n += 1 + runtime.Sov(uint64(x.AttrId))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Value != 0 {
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
		x := input.Message.Interface().(*NumericAttributeEntry)
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
		if x.Value != 0 {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(x.Value))
			i--
			dAtA[i] = 0x25
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if x.AttrId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AttrId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Target) > 0 {
			i -= len(x.Target)
			copy(dAtA[i:], x.Target)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Target)))
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
		x := input.Message.Interface().(*NumericAttributeEntry)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NumericAttributeEntry: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NumericAttributeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
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
				x.Target = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
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
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
				}
				x.Value = 0
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Value = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
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
// source: regen/ecocredit/curation/v1beta1/state.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TagDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Metadata     []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *TagDefinition) Reset() {
	*x = TagDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagDefinition) ProtoMessage() {}

// Deprecated: Use TagDefinition.ProtoReflect.Descriptor instead.
func (*TagDefinition) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_state_proto_rawDescGZIP(), []int{0}
}

func (x *TagDefinition) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TagDefinition) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *TagDefinition) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// NumericAttributeDefinition allows curators to define numeric attributes
// which can be assigned to credit classes
type NumericAttributeDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Metadata     []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// max_decimal_places
	MaxDecimalPlaces uint32 `protobuf:"varint,4,opt,name=max_decimal_places,json=maxDecimalPlaces,proto3" json:"max_decimal_places,omitempty"`
	Min              *int32 `protobuf:"varint,5,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max              *int32 `protobuf:"varint,6,opt,name=max,proto3,oneof" json:"max,omitempty"`
}

func (x *NumericAttributeDefinition) Reset() {
	*x = NumericAttributeDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumericAttributeDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumericAttributeDefinition) ProtoMessage() {}

// Deprecated: Use NumericAttributeDefinition.ProtoReflect.Descriptor instead.
func (*NumericAttributeDefinition) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_state_proto_rawDescGZIP(), []int{1}
}

func (x *NumericAttributeDefinition) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NumericAttributeDefinition) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *NumericAttributeDefinition) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *NumericAttributeDefinition) GetMaxDecimalPlaces() uint32 {
	if x != nil {
		return x.MaxDecimalPlaces
	}
	return 0
}

func (x *NumericAttributeDefinition) GetMin() int32 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *NumericAttributeDefinition) GetMax() int32 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

type ClassTagEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId  uint64 `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	TagId    uint64 `protobuf:"varint,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ClassTagEntry) Reset() {
	*x = ClassTagEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassTagEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassTagEntry) ProtoMessage() {}

// Deprecated: Use ClassTagEntry.ProtoReflect.Descriptor instead.
func (*ClassTagEntry) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_state_proto_rawDescGZIP(), []int{2}
}

func (x *ClassTagEntry) GetClassId() uint64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *ClassTagEntry) GetTagId() uint64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *ClassTagEntry) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type TagEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	TagId    uint64 `protobuf:"varint,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *TagEntry) Reset() {
	*x = TagEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagEntry) ProtoMessage() {}

// Deprecated: Use TagEntry.ProtoReflect.Descriptor instead.
func (*TagEntry) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_state_proto_rawDescGZIP(), []int{3}
}

func (x *TagEntry) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *TagEntry) GetTagId() uint64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *TagEntry) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type NumericAttributeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	AttrId   uint64 `protobuf:"varint,2,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// value represents a 32-bit decimal value with 2 decimal places, equivalent to value/100.
	Value int32 `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NumericAttributeEntry) Reset() {
	*x = NumericAttributeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumericAttributeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumericAttributeEntry) ProtoMessage() {}

// Deprecated: Use NumericAttributeEntry.ProtoReflect.Descriptor instead.
func (*NumericAttributeEntry) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_state_proto_rawDescGZIP(), []int{4}
}

func (x *NumericAttributeEntry) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *NumericAttributeEntry) GetAttrId() uint64 {
	if x != nil {
		return x.AttrId
	}
	return 0
}

func (x *NumericAttributeEntry) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *NumericAttributeEntry) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_regen_ecocredit_curation_v1beta1_state_proto protoreflect.FileDescriptor

var file_regen_ecocredit_curation_v1beta1_state_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x85, 0x01, 0x0a, 0x0d, 0x54, 0x61, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x3a, 0x23, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x1d, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x10, 0x01, 0x18, 0x01, 0x22, 0xfe, 0x01, 0x0a, 0x1a, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x88, 0x01, 0x01, 0x3a, 0x23, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x1d, 0x0a, 0x06, 0x0a, 0x02, 0x69,
	0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x18, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x22, 0x7a, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x54, 0x61, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x1b, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x15, 0x0a,
	0x11, 0x0a, 0x0f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x2c, 0x74, 0x61, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x22, 0x70, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x19, 0xf2, 0x9e, 0xd3,
	0x8e, 0x03, 0x13, 0x0a, 0x0f, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2c, 0x74, 0x61,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x22, 0xa7, 0x01, 0x0a, 0x15, 0x4e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x74, 0x74, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x2b, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x25, 0x0a, 0x0f, 0x0a, 0x0d, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2c, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x0c,
	0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x2c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x01, 0x18, 0x01,
	0x42, 0xb1, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x3b, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x43, 0xaa, 0x02, 0x20, 0x52, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x43,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2,
	0x02, 0x2c, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5c, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x23, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x3a, 0x3a, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_curation_v1beta1_state_proto_rawDescOnce sync.Once
	file_regen_ecocredit_curation_v1beta1_state_proto_rawDescData = file_regen_ecocredit_curation_v1beta1_state_proto_rawDesc
)

func file_regen_ecocredit_curation_v1beta1_state_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_curation_v1beta1_state_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_curation_v1beta1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_curation_v1beta1_state_proto_rawDescData)
	})
	return file_regen_ecocredit_curation_v1beta1_state_proto_rawDescData
}

var file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_regen_ecocredit_curation_v1beta1_state_proto_goTypes = []interface{}{
	(*TagDefinition)(nil),              // 0: regen.ecocredit.curation.v1beta1.TagDefinition
	(*NumericAttributeDefinition)(nil), // 1: regen.ecocredit.curation.v1beta1.NumericAttributeDefinition
	(*ClassTagEntry)(nil),              // 2: regen.ecocredit.curation.v1beta1.ClassTagEntry
	(*TagEntry)(nil),                   // 3: regen.ecocredit.curation.v1beta1.TagEntry
	(*NumericAttributeEntry)(nil),      // 4: regen.ecocredit.curation.v1beta1.NumericAttributeEntry
}
var file_regen_ecocredit_curation_v1beta1_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_curation_v1beta1_state_proto_init() }
func file_regen_ecocredit_curation_v1beta1_state_proto_init() {
	if File_regen_ecocredit_curation_v1beta1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagDefinition); i {
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
		file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumericAttributeDefinition); i {
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
		file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassTagEntry); i {
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
		file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagEntry); i {
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
		file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumericAttributeEntry); i {
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
	file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_curation_v1beta1_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_curation_v1beta1_state_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_curation_v1beta1_state_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_curation_v1beta1_state_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_curation_v1beta1_state_proto = out.File
	file_regen_ecocredit_curation_v1beta1_state_proto_rawDesc = nil
	file_regen_ecocredit_curation_v1beta1_state_proto_goTypes = nil
	file_regen_ecocredit_curation_v1beta1_state_proto_depIdxs = nil
}
