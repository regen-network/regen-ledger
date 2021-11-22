package ormsql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gorm.io/gorm"
)

func testgorm() {
	_ = gorm.DB{}
}

type tableCommiter struct {
	msgType     protoreflect.MessageType
	tableStruct reflect.Type
}

func buildTableCommiter(messageType protoreflect.MessageType, tableDesc *ormpb.TableDescriptor) (*tableCommiter, error) {
	if tableDesc.PrimaryKey == nil {
		return nil, fmt.Errorf("missing primary key")
	}

	pk := tableDesc.PrimaryKey
	pkFields := strings.Split(pk.Fields, ",")
	if len(pkFields) == 0 {
		return nil, fmt.Errorf("missing primary key fields")
	}
	pkFieldMap := map[string]bool{}
	for _, k := range pkFields {
		pkFieldMap[k] = true
	}

	desc := messageType.Descriptor()
	fieldDescriptors := desc.Fields()
	n := fieldDescriptors.Len()
	var structFields []reflect.StructField
	for i := 0; i < n; i++ {
		field := fieldDescriptors.Get(i)
		structField, err := buildStructField(field, pkFieldMap[string(field.Name())])
		if err != nil {
			return nil, err
		}
		structFields = append(structFields, structField)
	}
	return &tableCommiter{
		msgType:     messageType,
		tableStruct: reflect.StructOf(structFields),
	}, nil
}

func buildStructField(field protoreflect.FieldDescriptor, isPrimaryKey bool) (reflect.StructField, error) {
	tag := fmt.Sprintf(`gorm:"column:%s`, field.Name())
	if isPrimaryKey {
		tag = tag + fmt.Sprintf(`;primaryKey;autoIncrement:false`)
	}
	var fieldName = strings.ToTitle(string(field.Name()))
	typ, err := structFieldType(field)
	return reflect.StructField{
		Name: fieldName,
		Type: typ,
		Tag:  reflect.StructTag(tag + `"`),
	}, err
}

func structFieldType(descriptor protoreflect.FieldDescriptor) (reflect.Type, error) {
	return reflect.TypeOf(descriptor.Default().Interface()), nil
}
