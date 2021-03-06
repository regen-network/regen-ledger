// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v1alpha2/query.proto

package data

import (
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// QueryByContentHashRequest is the Query/ByContentHash request type.
type QueryByHashRequest struct {
	// hash is the hash-based identifier for the anchored content.
	Hash *ContentHash `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *QueryByHashRequest) Reset()         { *m = QueryByHashRequest{} }
func (m *QueryByHashRequest) String() string { return proto.CompactTextString(m) }
func (*QueryByHashRequest) ProtoMessage()    {}
func (*QueryByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7739eaec65300f, []int{0}
}
func (m *QueryByHashRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryByHashRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryByHashRequest.Merge(m, src)
}
func (m *QueryByHashRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryByHashRequest proto.InternalMessageInfo

func (m *QueryByHashRequest) GetHash() *ContentHash {
	if m != nil {
		return m.Hash
	}
	return nil
}

// QueryByContentHashResponse is the Query/ByContentHash response type.
type QueryByHashResponse struct {
	// entry is the ContentEntry
	Entry *ContentEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (m *QueryByHashResponse) Reset()         { *m = QueryByHashResponse{} }
func (m *QueryByHashResponse) String() string { return proto.CompactTextString(m) }
func (*QueryByHashResponse) ProtoMessage()    {}
func (*QueryByHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7739eaec65300f, []int{1}
}
func (m *QueryByHashResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryByHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryByHashResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryByHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryByHashResponse.Merge(m, src)
}
func (m *QueryByHashResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryByHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryByHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryByHashResponse proto.InternalMessageInfo

func (m *QueryByHashResponse) GetEntry() *ContentEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// QueryBySignerRequest is the Query/BySigner request type.
type QueryBySignerRequest struct {
	// signer is the address of the signer to query by.
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// pagination is the PageRequest to use for pagination.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryBySignerRequest) Reset()         { *m = QueryBySignerRequest{} }
func (m *QueryBySignerRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBySignerRequest) ProtoMessage()    {}
func (*QueryBySignerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7739eaec65300f, []int{2}
}
func (m *QueryBySignerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBySignerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBySignerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBySignerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBySignerRequest.Merge(m, src)
}
func (m *QueryBySignerRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBySignerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBySignerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBySignerRequest proto.InternalMessageInfo

func (m *QueryBySignerRequest) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *QueryBySignerRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryBySignerResponse is the Query/BySigner response type.
type QueryBySignerResponse struct {
	// entries is the ContentEntry's signed by the queried signer
	Entries []*ContentEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	// pagination is the pagination PageResponse.
	Pagination *query.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryBySignerResponse) Reset()         { *m = QueryBySignerResponse{} }
func (m *QueryBySignerResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBySignerResponse) ProtoMessage()    {}
func (*QueryBySignerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7739eaec65300f, []int{3}
}
func (m *QueryBySignerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBySignerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBySignerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBySignerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBySignerResponse.Merge(m, src)
}
func (m *QueryBySignerResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBySignerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBySignerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBySignerResponse proto.InternalMessageInfo

func (m *QueryBySignerResponse) GetEntries() []*ContentEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *QueryBySignerResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// ContentEntry describes data referenced and possibly stored on chain
type ContentEntry struct {
	// hash is the content hash
	Hash *ContentHash `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// iri is the content IRI
	Iri string `protobuf:"bytes,2,opt,name=iri,proto3" json:"iri,omitempty"`
	// timestamp is the anchor Timestamp
	Timestamp *types.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// signers are the signers, if any
	Signers []*SignerEntry `protobuf:"bytes,4,rep,name=signers,proto3" json:"signers,omitempty"`
	// content is the actual content if stored on-chain
	Content *Content `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *ContentEntry) Reset()         { *m = ContentEntry{} }
func (m *ContentEntry) String() string { return proto.CompactTextString(m) }
func (*ContentEntry) ProtoMessage()    {}
func (*ContentEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7739eaec65300f, []int{4}
}
func (m *ContentEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentEntry.Merge(m, src)
}
func (m *ContentEntry) XXX_Size() int {
	return m.Size()
}
func (m *ContentEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ContentEntry proto.InternalMessageInfo

func (m *ContentEntry) GetHash() *ContentHash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ContentEntry) GetIri() string {
	if m != nil {
		return m.Iri
	}
	return ""
}

func (m *ContentEntry) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ContentEntry) GetSigners() []*SignerEntry {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *ContentEntry) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryByHashRequest)(nil), "regen.data.v1alpha2.QueryByHashRequest")
	proto.RegisterType((*QueryByHashResponse)(nil), "regen.data.v1alpha2.QueryByHashResponse")
	proto.RegisterType((*QueryBySignerRequest)(nil), "regen.data.v1alpha2.QueryBySignerRequest")
	proto.RegisterType((*QueryBySignerResponse)(nil), "regen.data.v1alpha2.QueryBySignerResponse")
	proto.RegisterType((*ContentEntry)(nil), "regen.data.v1alpha2.ContentEntry")
}

func init() { proto.RegisterFile("regen/data/v1alpha2/query.proto", fileDescriptor_bf7739eaec65300f) }

var fileDescriptor_bf7739eaec65300f = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe3, 0xa6, 0x49, 0xc8, 0xc0, 0x01, 0x6d, 0x01, 0x59, 0x16, 0x72, 0x43, 0x0e, 0xb4,
	0x54, 0xb0, 0xab, 0x04, 0x04, 0x08, 0x6e, 0x45, 0x14, 0xc4, 0x01, 0x81, 0xe1, 0x04, 0xa7, 0x75,
	0x3a, 0xd8, 0x16, 0xc9, 0xae, 0xeb, 0xdd, 0x04, 0x72, 0xe7, 0x01, 0x78, 0x01, 0x9e, 0x07, 0x8e,
	0x3d, 0x72, 0x44, 0xc9, 0x8b, 0x20, 0xef, 0xae, 0x69, 0x03, 0x56, 0x5a, 0x71, 0xcb, 0x46, 0xdf,
	0xfc, 0xf3, 0xcf, 0x3f, 0x63, 0xd8, 0x2e, 0x30, 0x41, 0xc1, 0x0e, 0xb9, 0xe6, 0x6c, 0x36, 0xe0,
	0xe3, 0x3c, 0xe5, 0x43, 0x76, 0x34, 0xc5, 0x62, 0x4e, 0xf3, 0x42, 0x6a, 0x49, 0xb6, 0x0c, 0x40,
	0x4b, 0x80, 0x56, 0x40, 0xb0, 0x9d, 0x48, 0x99, 0x8c, 0x91, 0x19, 0x24, 0x9e, 0x7e, 0x60, 0x3a,
	0x9b, 0xa0, 0xd2, 0x7c, 0x92, 0xdb, 0xaa, 0x60, 0x6f, 0x24, 0xd5, 0x44, 0x2a, 0x16, 0x73, 0x85,
	0x56, 0x8e, 0xcd, 0x06, 0x31, 0x6a, 0x3e, 0x60, 0x39, 0x4f, 0x32, 0xc1, 0x75, 0x26, 0x85, 0x63,
	0x6b, 0x2d, 0xe8, 0x79, 0x8e, 0xca, 0x02, 0xfd, 0x17, 0x40, 0x5e, 0x97, 0x12, 0xfb, 0xf3, 0xe7,
	0x5c, 0xa5, 0x11, 0x1e, 0x4d, 0x51, 0x69, 0x72, 0x0f, 0x36, 0x53, 0xae, 0x52, 0xdf, 0xeb, 0x79,
	0xbb, 0x17, 0x87, 0x3d, 0x5a, 0xe3, 0x93, 0x3e, 0x91, 0x42, 0xa3, 0xd0, 0xa6, 0xcc, 0xd0, 0xfd,
	0x97, 0xb0, 0xb5, 0xa2, 0xa5, 0x72, 0x29, 0x14, 0x92, 0x07, 0xd0, 0x42, 0xa1, 0x8b, 0xb9, 0x53,
	0xbb, 0xb1, 0x4e, 0xed, 0x69, 0x09, 0x46, 0x96, 0xef, 0xcf, 0xe0, 0x8a, 0xd3, 0x7b, 0x93, 0x25,
	0x02, 0x8b, 0xca, 0xdd, 0x35, 0x68, 0x2b, 0xf3, 0x87, 0x51, 0xec, 0x46, 0xee, 0x45, 0x0e, 0x00,
	0x4e, 0x02, 0xf0, 0x37, 0x4c, 0xb7, 0x9b, 0xd4, 0xa6, 0x45, 0xcb, 0xb4, 0xa8, 0x0d, 0xdf, 0xa5,
	0x45, 0x5f, 0xf1, 0x04, 0x9d, 0x66, 0x74, 0xaa, 0xb2, 0xff, 0xcd, 0x83, 0xab, 0x7f, 0x35, 0x76,
	0xa3, 0x3c, 0x86, 0x4e, 0x69, 0x2d, 0x43, 0xe5, 0x7b, 0xbd, 0xe6, 0xf9, 0x86, 0xa9, 0x2a, 0xc8,
	0xb3, 0x15, 0x7b, 0x4d, 0x63, 0x6f, 0xe7, 0x4c, 0x7b, 0xb6, 0xf3, 0x8a, 0xbf, 0x2f, 0x1b, 0x70,
	0xe9, 0x74, 0x8b, 0xff, 0x5b, 0x17, 0xb9, 0x0c, 0xcd, 0xac, 0xc8, 0x4c, 0x4e, 0xdd, 0xa8, 0xfc,
	0x49, 0x1e, 0x42, 0xf7, 0xcf, 0xb1, 0x39, 0x83, 0x01, 0xb5, 0xe7, 0x48, 0xab, 0x73, 0xa4, 0x6f,
	0x2b, 0x22, 0x3a, 0x81, 0xc9, 0x23, 0xe8, 0xd8, 0x25, 0x28, 0x7f, 0xd3, 0x04, 0x53, 0x6f, 0xc2,
	0xc6, 0xe9, 0x72, 0x71, 0x05, 0xe4, 0x3e, 0x74, 0x46, 0xd6, 0x9c, 0xdf, 0x32, 0x3d, 0xaf, 0xaf,
	0x1b, 0x20, 0xaa, 0xe0, 0xe1, 0x77, 0x0f, 0x5a, 0x66, 0x4d, 0xe4, 0x3d, 0xb4, 0xed, 0xcd, 0x91,
	0x9d, 0xda, 0xd2, 0x7f, 0x2f, 0x3c, 0xd8, 0x3d, 0x1b, 0x74, 0x3b, 0xe7, 0x70, 0xa1, 0xba, 0x03,
	0x72, 0x6b, 0x5d, 0xd5, 0xca, 0x91, 0x06, 0x7b, 0xe7, 0x41, 0x6d, 0x8b, 0xfd, 0x83, 0x1f, 0x8b,
	0xd0, 0x3b, 0x5e, 0x84, 0xde, 0xaf, 0x45, 0xe8, 0x7d, 0x5d, 0x86, 0x8d, 0xe3, 0x65, 0xd8, 0xf8,
	0xb9, 0x0c, 0x1b, 0xef, 0x6e, 0x27, 0x99, 0x4e, 0xa7, 0x31, 0x1d, 0xc9, 0x09, 0x33, 0x7a, 0x77,
	0x04, 0xea, 0x4f, 0xb2, 0xf8, 0xe8, 0x5e, 0x63, 0x3c, 0x4c, 0xb0, 0x60, 0x9f, 0xcd, 0x17, 0x1e,
	0xb7, 0xcd, 0x92, 0xee, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x54, 0x77, 0x5e, 0x40, 0x79, 0x04,
	0x00, 0x00,
}

func (m *QueryByHashRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryByHashRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryByHashRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Hash != nil {
		{
			size, err := m.Hash.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryByHashResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryByHashResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryByHashResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Entry != nil {
		{
			size, err := m.Entry.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBySignerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBySignerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBySignerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBySignerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBySignerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBySignerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ContentEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Iri) > 0 {
		i -= len(m.Iri)
		copy(dAtA[i:], m.Iri)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Iri)))
		i--
		dAtA[i] = 0x12
	}
	if m.Hash != nil {
		{
			size, err := m.Hash.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryByHashRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryByHashResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Entry != nil {
		l = m.Entry.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBySignerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBySignerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ContentEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Iri)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, e := range m.Signers {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryByHashRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryByHashRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryByHashRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hash == nil {
				m.Hash = &ContentHash{}
			}
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryByHashResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryByHashResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryByHashResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entry == nil {
				m.Entry = &ContentEntry{}
			}
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBySignerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBySignerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBySignerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBySignerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBySignerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBySignerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &ContentEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ContentEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ContentEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hash == nil {
				m.Hash = &ContentHash{}
			}
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Iri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Iri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, &SignerEntry{})
			if err := m.Signers[len(m.Signers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &Content{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
