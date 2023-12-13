package data

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ToIRI converts the ContentHash to an IRI (internationalized URI) using the regen IRI scheme.
// A ContentHash IRI will look something like regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
// which is some base58check encoded data followed by a file extension or pseudo-extension.
// See ContentHash_Raw.ToIRI and ContentHash_Graph.ToIRI for more details on specific formatting.
func (ch ContentHash) ToIRI() (string, error) {
	if chr := ch.GetRaw(); chr != nil {
		return chr.ToIRI()
	} else if chg := ch.GetGraph(); chg != nil {
		return chg.ToIRI()
	}
	return "", sdkerrors.ErrInvalidType.Wrapf("invalid %T", ch)
}

const (
	iriVersion0 byte = 0

	IriPrefixRaw   byte = 0
	IriPrefixGraph byte = 1
)

// ToIRI converts the ContentHash_Raw to an IRI (internationalized URI) based on the following
// pattern: regen:{base58check(concat( byte(0x0), byte(digest_algorithm), hash), 0)}.{file_extension}
func (chr ContentHash_Raw) ToIRI() (string, error) {
	err := chr.Validate()
	if err != nil {
		return "", err
	}

	bz := make([]byte, len(chr.Hash)+2)
	bz[0] = IriPrefixRaw
	bz[1] = byte(chr.DigestAlgorithm)
	copy(bz[2:], chr.Hash)
	hashStr := base58.CheckEncode(bz, iriVersion0)
	ext := chr.FileExtension

	return fmt.Sprintf("regen:%s.%s", hashStr, ext), nil
}

// ToIRI converts the ContentHash_Graph to an IRI (internationalized URI) based on the following
// pattern: regen:{base58check(concat(byte(0x1), byte(canonicalization_algorithm),
// byte(merkle_tree), byte(digest_algorithm), hash), 0)}.rdf
func (chg ContentHash_Graph) ToIRI() (string, error) {
	err := chg.Validate()
	if err != nil {
		return "", err
	}

	bz := make([]byte, len(chg.Hash)+4)
	bz[0] = IriPrefixGraph
	bz[1] = byte(chg.CanonicalizationAlgorithm)
	bz[2] = byte(chg.MerkleTree)
	bz[3] = byte(chg.DigestAlgorithm)
	copy(bz[4:], chg.Hash)
	hashStr := base58.CheckEncode(bz, iriVersion0)

	return fmt.Sprintf("regen:%s.rdf", hashStr), nil
}

// ParseIRI parses an IRI string representation of a ContentHash into a ContentHash struct
// Currently IRIs must have a "regen:" prefix, and only ContentHash_Graph and ContentHash_Raw
// are supported.
func ParseIRI(iri string) (*ContentHash, error) {
	const regenPrefix = "regen:"

	if iri == "" {
		return nil, ErrInvalidIRI.Wrap("failed to parse IRI: empty string is not allowed")
	}

	if !strings.HasPrefix(iri, regenPrefix) {
		return nil, ErrInvalidIRI.Wrapf("failed to parse IRI %s: %s prefix required", iri, regenPrefix)
	}

	hashExtPart := iri[len(regenPrefix):]
	parts := strings.Split(hashExtPart, ".")
	if len(parts) != 2 {
		return nil, ErrInvalidIRI.Wrapf("failed to parse IRI %s: extension required", iri)
	}

	hashPart, ext := parts[0], parts[1]

	res, version, err := base58.CheckDecode(hashPart)
	if err != nil {
		return nil, ErrInvalidIRI.Wrapf("failed to parse IRI %s: %s", iri, err)
	}

	rdr := bytes.NewBuffer(res)

	// read first byte
	typ, err := rdr.ReadByte()
	if err != nil {
		return nil, err
	}

	// switch on first byte which represents the type prefix
	switch typ {
	case IriPrefixRaw:
		// read next byte
		b0, err := rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		hash := rdr.Bytes()

		if version == iriVersion0 {
			return &ContentHash{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: uint32(b0),
				FileExtension:   ext,
			}}, nil
		} else {
			return nil, ErrInvalidIRI.Wrapf("failed to parse IRI %s: invalid version %d", iri, version)
		}

	case IriPrefixGraph:
		// rdf extension is expected for graph data
		if ext != "rdf" {
			return nil, ErrInvalidIRI.Wrapf("invalid extension .%s for graph data, expected .rdf", ext)
		}

		// read canonicalization algorithm
		bC14NAlg, err := rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// read merkle tree algorithm
		bMtAlg, err := rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// read digest algorithm
		bDigestAlg, err := rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// read hash
		hash := rdr.Bytes()

		if version == iriVersion0 {
			return &ContentHash{Graph: &ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           uint32(bDigestAlg),
				CanonicalizationAlgorithm: uint32(bC14NAlg),
				MerkleTree:                uint32(bMtAlg),
			}}, nil
		} else {
			return nil, ErrInvalidIRI.Wrapf("failed to parse IRI %s: invalid version %d", iri, version)
		}
	}

	return nil, ErrInvalidIRI.Wrapf("unable to parse IRI %s", iri)
}
