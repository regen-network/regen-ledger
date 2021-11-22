package ormdecode

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/structpb"
)

type Entry interface {
	isEntry()
	fmt.Stringer
}

type PrimaryKeyEntry struct {
	Key   []protoreflect.Value
	Value proto.Message
}

func (p PrimaryKeyEntry) String() string {
	msg := p.Value
	name := msg.ProtoReflect().Descriptor().FullName()
	msgBz, err := protojson.Marshal(msg)
	msgStr := string(msgBz)
	if err != nil {
		msgStr = fmt.Sprintf("%s:%+v", name, msg)
	}
	return fmt.Sprintf("PK:%s:%s:%s", name, fmtValues(p.Key), msgStr)
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

func (p PrimaryKeyEntry) isEntry() {}

type IndexKeyEntry struct {
	TableName       protoreflect.FullName
	IndexFieldNames string
	IndexKey        []protoreflect.Value
	PrimaryKeyRest  []protoreflect.Value
	PrimaryKey      []protoreflect.Value
}

func (i IndexKeyEntry) isEntry() {}

func (i IndexKeyEntry) string() string {
	return fmt.Sprintf("%s%s:%s:%s", i.TableName, i.IndexFieldNames, fmtValues(i.IndexKey), fmtValues(i.PrimaryKeyRest))
}

func (i IndexKeyEntry) String() string {
	return fmt.Sprintf("IDX:%s", i.string())
}

type UniqueKeyEntry struct {
	IndexKeyEntry
}

func (u UniqueKeyEntry) String() string {
	return fmt.Sprintf("UNIQ:%s", u.string())
}

type SeqEntry struct {
	TableName string
	Value     uint64
}

func (s SeqEntry) isEntry() {}

func (s SeqEntry) String() string {
	return fmt.Sprintf("SEQ:%s:%d", s.TableName, s.Value)
}

type SchemaEntry struct {
	Id             uint32
	FileDescriptor *descriptorpb.FileDescriptorProto
}

func (s SchemaEntry) String() string {
	return fmt.Sprintf("FILEDESC:%v", s.FileDescriptor.Name)
}

func (s SchemaEntry) isEntry() {}

var _, _, _, _ Entry = PrimaryKeyEntry{}, IndexKeyEntry{}, SeqEntry{}, SchemaEntry{}
