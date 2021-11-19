package list

import "google.golang.org/protobuf/proto"

type ErrIterator struct {
	Err error
}

func (e ErrIterator) isIterator() {}

func (e ErrIterator) Next(proto.Message) (bool, error) { return false, e.Err }
