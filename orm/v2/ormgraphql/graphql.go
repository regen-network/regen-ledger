package ormgraphql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/regen-network/regen-ledger/orm/v2/ormpb"

	"github.com/graphql-go/graphql/language/ast"

	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Builder struct {
	objects map[string]*graphql.Object
}

func (b Builder) buildTable(tableDesc *ormpb.TableDescriptor, desc protoreflect.MessageDescriptor) (*graphql.Field, error) {
	name := messageName(desc)
	objType, err := b.buildObject(desc)
	if err != nil {
		return nil, err
	}

	return &graphql.Field{
		Name:              name,
		Type:              objType,
		Args:              nil,
		Resolve:           nil,
		Subscribe:         nil,
		DeprecationReason: "",
		Description:       "",
	}, nil
}

func messageName(descriptor protoreflect.MessageDescriptor) string {
	return strings.ReplaceAll(string(descriptor.FullName()), ".", "_")
}

func (b Builder) buildObject(descriptor protoreflect.MessageDescriptor) (*graphql.Object, error) {
	name := messageName(descriptor)

	if obj, ok := b.objects[name]; ok {
		return obj, nil
	}

	fieldDescriptors := descriptor.Fields()
	n := fieldDescriptors.Len()
	fields := graphql.Fields{}

	for i := 0; i < n; i++ {
		f, err := b.buildField(fieldDescriptors.Get(i))
		// skip field errors - we can't deal with everything
		if err != nil {
			continue
		}
		fields[f.Name] = f
	}

	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})
	b.objects[name] = obj
	return obj, nil
}

func (b Builder) buildField(descriptor protoreflect.FieldDescriptor) (*graphql.Field, error) {
	typ, err := b.fieldType(descriptor)
	if err != nil {
		return nil, err
	}

	return &graphql.Field{
		Name: string(descriptor.Name()),
		Type: typ,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			mref := p.Source.(protoreflect.Message)
			return mref.Get(descriptor), nil
		},
	}, nil
}

func (b Builder) fieldType(descriptor protoreflect.FieldDescriptor) (graphql.Type, error) {
	simpleType, err := b.simpleFieldType(descriptor)
	if err != nil {
		return nil, err
	}

	if descriptor.IsList() {
		return graphql.NewList(simpleType), nil
	} else if descriptor.IsMap() {
		return nil, fmt.Errorf("map field %s not supported", descriptor)

	} else {
		return simpleType, nil
	}
}

func (b Builder) simpleFieldType(descriptor protoreflect.FieldDescriptor) (graphql.Type, error) {
	switch descriptor.Kind() {
	case protoreflect.BoolKind:
		return graphql.Boolean, nil
	case protoreflect.Int32Kind:
		return graphql.Int, nil
	case protoreflect.Uint32Kind:
		return b.uint32Scalar(), nil
	case protoreflect.MessageKind:
		return b.buildObject(descriptor.Message())
	default:
		return nil, fmt.Errorf("don't know how to convert field %s", descriptor)
	}
}

func (b Builder) uint32Scalar() *graphql.Scalar {
	return graphql.NewScalar(graphql.ScalarConfig{
		Name: "UInt32",
		Serialize: func(value interface{}) interface{} {
			return strconv.FormatUint(uint64(value.(uint32)), 10)

		},
		ParseValue: func(value interface{}) interface{} {
			x, err := strconv.ParseUint(value.(string), 10, 32)
			if err != nil {
				return uint32(0)
			}
			return uint32(x)
		},
		ParseLiteral: func(valueAST ast.Value) interface{} {
			panic("TODO")
		},
	})
}
