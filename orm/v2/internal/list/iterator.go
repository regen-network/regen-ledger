package list

import "google.golang.org/protobuf/proto"

type Iterator interface {
	Next(proto.Message) (bool, error)
	Close()
}
