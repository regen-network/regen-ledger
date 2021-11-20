package list

import "google.golang.org/protobuf/proto"

type Options struct {
	Reverse   bool
	IndexHint string
	Start     proto.Message
	End       proto.Message
}
