package utils

import (
	"fmt"
	"reflect"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/proto"
)

// MatchEvent matches the values in a proto message struct to the attributes in a sdk.Event.
func MatchEvent(event any, emitted sdk.Event) error {
	tag := "json"
	valOfEvent := reflect.ValueOf(event)
	typeOfEvent := valOfEvent.Type()
	if typeOfEvent.Kind() != reflect.Struct {
		return fmt.Errorf("expected event to be struct, got %T", event)
	}
	attrMap := mapAttributes(emitted)

	numExportedFields := 0
	for i := 0; i < typeOfEvent.NumField(); i++ {
		underlying := valOfEvent.Field(i)
		descriptor := typeOfEvent.Field(i)
		if !descriptor.IsExported() {
			continue
		}
		numExportedFields++
		key := strings.Split(descriptor.Tag.Get(tag), ",")[0]
		val, ok := attrMap[key]
		if !ok {
			return fmt.Errorf("event has no attribute '%s'", key)
		}
		if underlyingValue := fmt.Sprintf("%v", underlying.Interface()); underlyingValue != val {
			return fmt.Errorf("expected %s, got %s for field %s", underlyingValue, val, descriptor.Name)
		}
	}
	if numAttrs := len(emitted.Attributes); numExportedFields != numAttrs {
		return fmt.Errorf("emitted event has %d attributes, expected %d", numAttrs, numExportedFields)
	}
	return nil
}

func GetEvent(msg proto.Message, events []sdk.Event) (e sdk.Event, found bool) {
	for _, e := range events {
		if string(proto.MessageName(msg)) == e.Type {
			return e, true
		}
	}
	return e, false
}

// mapAttributes converts the sdk.Event attribute slice to a map.
func mapAttributes(event sdk.Event) map[string]string {
	m := make(map[string]string, len(event.Attributes))
	for _, attr := range event.Attributes {
		m[string(attr.Key)] = strings.Trim(string(attr.Value), `"`)
	}
	return m
}
