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
// pattern: regen:{base58check(concat( byte(0x0), byte(digest_algorithm), hash))}.{media_type extension}
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

	ext, err := chr.MediaType.ToExtension()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("regen:%s.%s", hashStr, ext), nil
}

// ToIRI converts the ContentHash_Graph to an IRI (internationalized URI) based on the following
// pattern: regen:{base58check(concat(byte(0x1), byte(canonicalization_algorithm),
// byte(merkle_tree), byte(digest_algorithm), hash))}.rdf
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

// ToExtension converts the media type to a file extension based on the mediaTypeExtensions map.
func (rmt RawMediaType) ToExtension() (string, error) {
	ext, ok := mediaExtensionTypeToString[rmt]
	if !ok {
		return "", sdkerrors.ErrInvalidRequest.Wrapf("missing extension for %T %s", rmt, rmt)
	}

	return ext, nil
}

var mediaExtensionTypeToString = map[RawMediaType]string{
	RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED: "bin",
	RawMediaType_RAW_MEDIA_TYPE_TEXT_PLAIN:  "txt",
	RawMediaType_RAW_MEDIA_TYPE_CSV:         "csv",
	RawMediaType_RAW_MEDIA_TYPE_JSON:        "json",
	RawMediaType_RAW_MEDIA_TYPE_XML:         "xml",
	RawMediaType_RAW_MEDIA_TYPE_PDF:         "pdf",
	RawMediaType_RAW_MEDIA_TYPE_TIFF:        "tiff",
	RawMediaType_RAW_MEDIA_TYPE_JPG:         "jpg",
	RawMediaType_RAW_MEDIA_TYPE_PNG:         "png",
	RawMediaType_RAW_MEDIA_TYPE_SVG:         "svg",
	RawMediaType_RAW_MEDIA_TYPE_WEBP:        "webp",
	RawMediaType_RAW_MEDIA_TYPE_AVIF:        "avif",
	RawMediaType_RAW_MEDIA_TYPE_GIF:         "gif",
	RawMediaType_RAW_MEDIA_TYPE_APNG:        "apng",
	RawMediaType_RAW_MEDIA_TYPE_MPEG:        "mpeg",
	RawMediaType_RAW_MEDIA_TYPE_MP4:         "mp4",
	RawMediaType_RAW_MEDIA_TYPE_WEBM:        "webm",
	RawMediaType_RAW_MEDIA_TYPE_OGG:         "ogg",
}

var stringToMediaExtensionType = map[string]RawMediaType{}

func init() {
	for mt, ext := range mediaExtensionTypeToString {
		stringToMediaExtensionType[ext] = mt
	}
}

// ParseIRI parses an IRI string representation of a ContentHash into a ContentHash struct
// Currently IRIs must have a "regen:" prefix, and only ContentHash_Graph and ContentHash_Raw
// are supported.
func ParseIRI(iri string) (*ContentHash, error) {
	const regenPrefix = "regen:"

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

	if version != iriVersion0 {
		return nil, ErrInvalidIRI.Wrapf("failed to parse IRI %s: invalid version", iri)
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

		// look up extension as media type
		mediaType, ok := stringToMediaExtensionType[ext]
		if !ok {
			return nil, ErrInvalidMediaExtension.Wrapf("failed to resolve media type for extension %s, expected %s", ext, mediaExtensionTypeToString[mediaType])
		}

		// interpret next byte as digest algorithm
		digestAlg := DigestAlgorithm(b0)
		hash := rdr.Bytes()
		err = digestAlg.Validate(hash)
		if err != nil {
			return nil, err
		}

		return &ContentHash{Raw: &ContentHash_Raw{
			Hash:            hash,
			DigestAlgorithm: digestAlg,
			MediaType:       mediaType,
		}}, nil

	case IriPrefixGraph:
		// rdf extension is expected for graph data
		if ext != "rdf" {
			return nil, ErrInvalidIRI.Wrapf("invalid extension .%s for graph data, expected .rdf", ext)
		}

		// read next byte
		b0, err := rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// interpret next byte as canonicalization algorithm
		c14Alg := GraphCanonicalizationAlgorithm(b0)
		err = c14Alg.Validate()
		if err != nil {
			return nil, err
		}

		// read next byte
		b0, err = rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// interpret next byte as merklization algorithm
		mtAlg := GraphMerkleTree(b0)
		err = mtAlg.Validate()
		if err != nil {
			return nil, err
		}

		// read next byte
		b0, err = rdr.ReadByte()
		if err != nil {
			return nil, err
		}

		// interpret next byte as digest algorithm
		digestAlg := DigestAlgorithm(b0)
		hash := rdr.Bytes()
		err = digestAlg.Validate(hash)
		if err != nil {
			return nil, err
		}

		return &ContentHash{Graph: &ContentHash_Graph{
			Hash:                      hash,
			DigestAlgorithm:           digestAlg,
			CanonicalizationAlgorithm: c14Alg,
			MerkleTree:                mtAlg,
		}}, nil
	}

	return nil, ErrInvalidIRI.Wrapf("unable to parse IRI %s", iri)
}
