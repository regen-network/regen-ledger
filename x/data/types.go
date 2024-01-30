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

func (chr *ContentHash_Raw) Validate() error {
	err := validateHash(chr.Hash, chr.DigestAlgorithm)
	if err != nil {
		return err
	}

	ext := chr.FileExtension
	extLen := len(ext)
	if extLen < 2 {
		return sdkerrors.ErrInvalidRequest.Wrapf("file extension cannot be shorter than 2 characters")
	}

	if extLen > 6 {
		return sdkerrors.ErrInvalidRequest.Wrapf("file extension cannot be longer than 6 characters")
	}

	// check that ext is all lowercase or numeric
	for _, c := range ext {
		if c < '0' || c > '9' && c < 'a' || c > 'z' {
			return sdkerrors.ErrInvalidRequest.Wrapf("file extension must be all lowercase or numeric")
		}
	}

	return nil
}

func (chg *ContentHash_Graph) Validate() error {
	err := validateHash(chg.Hash, chg.DigestAlgorithm)
	if err != nil {
		return err
	}

	if chg.CanonicalizationAlgorithm == 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("canonicalization algorithm cannot be empty")
	}

	return nil
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

func validateHash(hash []byte, digestAlgorithm uint32) error {
	hashLen := len(hash)
	if hashLen < 20 {
		return sdkerrors.ErrInvalidRequest.Wrapf("hash cannot be shorter than 20 bytes")
	}

	if hashLen > 64 {
		return sdkerrors.ErrInvalidRequest.Wrapf("hash cannot be longer than 64 bytes")
	}

	if digestAlgorithm == 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("digest algorithm cannot be empty")
	}

	return nil
}
