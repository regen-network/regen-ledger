package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/proto"
)

// MatchEvent matches the values in a proto message struct to the attributes in a sdk.Event.
func MatchEvent(event any, emitted sdk.Event) error {
	tag := "json"
	valOfEvent := reflect.ValueOf(event)
	// if the value is a pointer, get the underlying element
	if valOfEvent.Kind() == reflect.Ptr {
		valOfEvent = valOfEvent.Elem()
	}
	typeOfEvent := valOfEvent.Type()
	if typeOfEvent.Kind() != reflect.Struct {
		return fmt.Errorf("expected event to be struct, got %T", event)
	}

	// convert key/value of event attributes to map
	attrMap := mapAttributes(emitted)

	numExportedFields := 0
	// iterate over the struct fields
	for i := 0; i < typeOfEvent.NumField(); i++ {
		underlying := valOfEvent.Field(i)
		descriptor := typeOfEvent.Field(i)
		if !descriptor.IsExported() {
			continue
		}
		numExportedFields++

		// extract the json tag, which is the key of the event attribute
		key := strings.Split(descriptor.Tag.Get(tag), ",")[0]
		val, ok := attrMap[key]
		if !ok {
			return fmt.Errorf("event has no attribute '%s'", key)
		}

		// convert the underlying struct field value to a string
		underlyingValue := fmt.Sprintf("%v", underlying.Interface())

		// handle special case for null values
		if underlyingValue == "<nil>" {
			underlyingValue = "null"
		} else if len(val) > 0 && val[0] == '{' { // handle nested struct case
			// nested event structs end up as a json marshalled string, so we must unmarshal it
			sdkEvent, err := jsonToEvent(val)
			if err != nil {
				return err
			}
			typ := underlying.Type()
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
				dst := reflect.New(typ).Elem()
				err := json.Unmarshal([]byte(val), dst.Addr().Interface())
				if err != nil {
					return err
				}
				if err := MatchEvent(dst.Addr().Interface(), sdkEvent); err != nil {
					return err
				}
			} else {
				dst := reflect.New(typ).Elem()
				err := json.Unmarshal([]byte(val), dst.Addr().Interface())
				if err != nil {
					return err
				}
				if err := MatchEvent(dst.Interface(), sdkEvent); err != nil {
					return err
				}
			}
		} else {
			// it's not a nested struct or nil, so we can just simply compare
			if underlyingValue != val {
				return fmt.Errorf("expected %s, got %s for field %s", underlyingValue, val, descriptor.Name)
			}
		}
	}
	// check that no fields were missed. amount of exported struct fields should be equal to amount of attributes
	if numAttrs := len(emitted.Attributes); numExportedFields != numAttrs {
		return fmt.Errorf("emitted event has %d attributes, expected %d", numAttrs, numExportedFields)
	}
	return nil
}

func jsonToEvent(jsn string) (sdk.Event, error) {
	event := sdk.Event{}
	m := make(map[string]string, 0)
	err := json.Unmarshal([]byte(jsn), &m)
	if err != nil {
		return sdk.Event{}, err
	}
	for k, v := range m {
		event.Attributes = append(event.Attributes, types.EventAttribute{
			Key:   []byte(k),
			Value: []byte(v),
		})
	}
	return event, nil
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
