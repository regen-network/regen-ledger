package list

import "google.golang.org/protobuf/proto"

type ErrIterator struct {
	Err error
}

func (e ErrIterator) Next(proto.Message) (bool, error) { return false, e.Err }

func (e ErrIterator) Close() {}
