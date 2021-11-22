package kvlayout

import (
	"google.golang.org/protobuf/types/descriptorpb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Entry interface {
	isEntry()
}

type PrimaryEntry struct {
	Key   []protoreflect.Value
	Value proto.Message
}

func (p PrimaryEntry) isEntry() {}

type IndexEntry struct {
	IndexFields string
	Key         []protoreflect.Value
	PrimaryKey  []protoreflect.Value
}

func (i IndexEntry) isEntry() {}

type SeqEntry struct {
	TableName string
	Value     uint64
}

func (s SeqEntry) isEntry() {}

type SchemaEntry struct {
	Id             uint32
	FileDescriptor *descriptorpb.FileDescriptorProto
}

func (s SchemaEntry) isEntry() {}

var _, _, _, _ Entry = PrimaryEntry{}, IndexEntry{}, SeqEntry{}, SchemaEntry{}
