package data

import (
	"reflect"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var DigestAlgorithmLength = map[DigestAlgorithm]int{
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256: 256,
}

func (ch ContentHash) Validate() error {
	hashRaw := ch.GetRaw()
	hashGraph := ch.GetGraph()

	switch {
	case hashRaw != nil && hashGraph != nil:
		return sdkerrors.ErrInvalidRequest.Wrapf("content hash must be one of raw type or graph type")
	case hashRaw != nil:
		return hashRaw.Validate()
	case hashGraph != nil:
		return hashGraph.Validate()
	}

	return sdkerrors.ErrInvalidRequest.Wrapf("content hash must be one of raw type or graph type")
}

func (chr ContentHash_Raw) Validate() error {
	err := chr.DigestAlgorithm.Validate(chr.Hash)
	if err != nil {
		return err
	}

	return chr.MediaType.Validate()
}

func (chg ContentHash_Graph) Validate() error {
	err := chg.DigestAlgorithm.Validate(chg.Hash)
	if err != nil {
		return err
	}

	err = chg.CanonicalizationAlgorithm.Validate()
	if err != nil {
		return err
	}

	return chg.MerkleTree.Validate()
}

func (da DigestAlgorithm) Validate(hash []byte) error {
	if reflect.DeepEqual(hash, []byte(nil)) {
		return sdkerrors.ErrInvalidRequest.Wrapf("hash cannot be empty")
	}

	if da == DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid %T %s", da, da)
	}

	nBits, ok := DigestAlgorithmLength[da]
	if !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("unknown %T %s", da, da)
	}

	nBytes := nBits / 8
	if len(hash) != nBytes {
		return sdkerrors.ErrInvalidRequest.Wrapf("expected %d bytes for %s, got %d", nBytes, da, len(hash))
	}

	return nil
}

func (rmt RawMediaType) Validate() error {
	if _, ok := RawMediaType_name[int32(rmt)]; !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("unknown %T %d", rmt, rmt)
	}

	return nil
}

func (gca GraphCanonicalizationAlgorithm) Validate() error {
	if _, ok := GraphCanonicalizationAlgorithm_name[int32(gca)]; !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("unknown %T %d", gca, gca)
	}

	if gca == GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid %T %s", gca, gca)
	}

	return nil
}

func (gmt GraphMerkleTree) Validate() error {
	if _, ok := GraphMerkleTree_name[int32(gmt)]; !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("unknown %T %d", gmt, gmt)
	}

	return nil
}
