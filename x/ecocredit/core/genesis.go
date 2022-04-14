package core

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	gogoproto "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// MergeParamsIntoTarget merges params message into the ormjson.WriteTarget.
func MergeParamsIntoTarget(cdc codec.JSONCodec, message gogoproto.Message, target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(message)))
	if err != nil {
		return err
	}

	bz, err := cdc.MarshalJSON(message)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}
