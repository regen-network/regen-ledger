package basket

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
)

// ToApi converts to pulsar based data structure
func (d *DateCriteria) ToApi() *basketv1.DateCriteria {
	if d == nil {
		return nil
	}
	if x := d.GetMinStartDate(); x != nil {
		return &basketv1.DateCriteria{Sum: &basketv1.DateCriteria_MinStartDate{
			&timestamppb.Timestamp{Seconds: x.Seconds, Nanos: x.Nanos}}}
	} else if x := d.GetStartDateWindow(); x != nil {
		return &basketv1.DateCriteria{Sum: &basketv1.DateCriteria_StartDateWindow{
			&durationpb.Duration{Seconds: x.Seconds, Nanos: x.Nanos}}}
	}
	return nil
}

var _ codec.AminoMarshaler = &DateCriteria{}

func (d DateCriteria) MarshalAminoJSON() ([]byte, error) {
	return d.Marshal()
}

func (d *DateCriteria) UnmarshalAminoJSON(bytes []byte) error {
	return d.Unmarshal(bytes)
}

func (d DateCriteria) MarshalAmino() ([]byte, error) {
	return d.Marshal()
}

func (d *DateCriteria) UnmarshalAmino(bytes []byte) error {
	return d.Unmarshal(bytes)
}
