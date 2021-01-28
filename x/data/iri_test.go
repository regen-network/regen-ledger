package data

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
)

func TestAccAddressToDID(t *testing.T) {
	type args struct {
		address         types.AccAddress
		bech32AccPrefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccAddressToDID(tt.args.address, tt.args.bech32AccPrefix); got != tt.want {
				t.Errorf("AccAddressToDID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentHash_Graph_ToIRI(t *testing.T) {
	type fields struct {
		Hash                      []byte
		DigestAlgorithm           DigestAlgorithm
		CanonicalizationAlgorithm GraphCanonicalizationAlgorithm
		MerkleTree                GraphMerkleTree
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chg := ContentHash_Graph{
				Hash:                      tt.fields.Hash,
				DigestAlgorithm:           tt.fields.DigestAlgorithm,
				CanonicalizationAlgorithm: tt.fields.CanonicalizationAlgorithm,
				MerkleTree:                tt.fields.MerkleTree,
			}
			got, err := chg.ToIRI()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIRI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIRI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentHash_Raw_ToIRI(t *testing.T) {
	type fields struct {
		Hash            []byte
		DigestAlgorithm DigestAlgorithm
		MediaType       MediaType
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chr := ContentHash_Raw{
				Hash:            tt.fields.Hash,
				DigestAlgorithm: tt.fields.DigestAlgorithm,
				MediaType:       tt.fields.MediaType,
			}
			got, err := chr.ToIRI()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIRI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIRI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentHash_ToIRI(t *testing.T) {
	type fields struct {
		Sum isContentHash_Sum
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := ContentHash{
				Sum: tt.fields.Sum,
			}
			got, err := ch.ToIRI()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIRI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIRI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMediaType_ToExtension(t *testing.T) {
	tests := []struct {
		name    string
		mt      MediaType
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.mt.ToExtension()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToExtension() got = %v, want %v", got, tt.want)
			}
		})
	}
}
