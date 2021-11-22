package kvlayout

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

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

func (p PrimaryEntry) String() string {
	msgBz, err := protojson.Marshal(p.Value)
	msgStr := string(msgBz)
	if err != nil {
		msgStr = fmt.Sprintf("%+v", p.Value)
	}
	return fmt.Sprintf("PK(%s:%s)", fmtValues(p.Key), msgStr)
}

func fmtValues(values []protoreflect.Value) string {
	var xs []interface{}
	for _, v := range values {
		xs = append(xs, v.Interface())
	}
	list, err := structpb.NewList(xs)
	if err != nil {
		return fmt.Sprintf("%+v", values)
	}
	bz, err := protojson.Marshal(list)
	if err != nil {
		return fmt.Sprintf("%+v", values)
	}
	return string(bz)
}

func (p PrimaryEntry) isEntry() {}

type IndexEntry struct {
	IndexFields string
	Key         []protoreflect.Value
	PrimaryKey  []protoreflect.Value
}

func (i IndexEntry) isEntry() {}

func (i IndexEntry) String() string {
	return fmt.Sprintf("IDX(%s:%s:%s)", i.IndexFields, fmtValues(i.Key), fmtValues(i.PrimaryKey))
}

type SeqEntry struct {
	TableName string
	Value     uint64
}

func (s SeqEntry) isEntry() {}

func (s SeqEntry) String() string {
	return fmt.Sprintf("SEQ(%s:%d)", s.TableName, s.Value)
}

type SchemaEntry struct {
	Id             uint32
	FileDescriptor *descriptorpb.FileDescriptorProto
}

func (s SchemaEntry) isEntry() {}

var _, _, _, _ Entry = PrimaryEntry{}, IndexEntry{}, SeqEntry{}, SchemaEntry{}
