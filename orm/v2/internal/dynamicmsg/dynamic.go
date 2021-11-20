package dynamicmsg

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

func Unmarshal(descriptor protoreflect.MessageDescriptor, bz []byte) (proto.Message, error) {
	msgType, err := protoregistry.GlobalTypes.FindMessageByName(descriptor.FullName())
	if err != nil {
		msgType = dynamicpb.NewMessageType(descriptor)
	}

	msg := msgType.New().Interface()
	err = proto.Unmarshal(bz, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
