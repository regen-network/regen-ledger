package list

import "google.golang.org/protobuf/proto"

type Options struct {
	Reverse  bool
	UseIndex string
	Start    proto.Message
	End      proto.Message
}
