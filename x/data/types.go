package data

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (ch ContentHash) Validate() error {
	switch hash := ch.Sum.(type) {
	case *ContentHash_Raw_:
		return hash.Raw.Validate()
	case *ContentHash_Graph_:
		return hash.Graph.Validate()
	default:
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid %T type %T", ch, hash)
	}
}

func (chr ContentHash_Raw) Validate() error {
	err := chr.MediaType.Validate()
	if err != nil {
		return err
	}

	return chr.DigestAlgorithm.Validate(chr.Hash)
}

func (chg ContentHash_Graph) Validate() error {
	err := chg.CanonicalizationAlgorithm.Validate()
	if err != nil {
		return err
	}

	err = chg.MerkleTree.Validate()
	if err != nil {
		return err
	}

	return chg.DigestAlgorithm.Validate(chg.Hash)
}

func (x MediaType) Validate() error {
	if _, ok := MediaType_name[int32(x)]; !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("unknown %T %d", x, x)
	}

	return nil
}

func (x DigestAlgorithm) Validate(hash []byte) error {
	nBits, ok := DigestalgorithmLength[x]
	if !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid or unknown %T %s", x, x)
	}

	nBytes := nBits / 8
	if len(hash) != nBytes {
		return sdkerrors.ErrInvalidRequest.Wrapf("expected %d bytes for %s, got %d", nBytes, x, len(hash))
	}

	return nil
}

var DigestalgorithmLength = map[DigestAlgorithm]int{
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256: 256,
}

func (x GraphCanonicalizationAlgorithm) Validate() error {
	if _, ok := GraphCanonicalizationAlgorithm_name[int32(x)]; !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("unknown %T %d", x, x)
	}

	if x == GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid %T %s", x, x)
	}

	return nil
}

func (x GraphMerkleTree) Validate() error {
	if _, ok := GraphMerkleTree_name[int32(x)]; !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("unknown %T %d", x, x)
	}

	return nil
}
