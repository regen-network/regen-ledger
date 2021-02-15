package data

import (
	"bytes"
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/btcsuite/btcutil/base58"
)

// AccAddressToDID converts an account address to a DID using a chain-specific method prefix,
// which should generally be equivalent to the chain's bech32 account prefix.
func AccAddressToDID(address sdk.AccAddress, methodPrefix string) string {
	b58str := base58.CheckEncode(address, didVersion0)
	return fmt.Sprintf("did:%s:%s", methodPrefix, b58str)
}

// ToIRI converts the ContentHash to an IRI (internationalized URI) using the regen IRI scheme.
// A ContentHash IRI will look something like regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
// which is some base58check encoded data followed by a file extension or pseudo-extension.
// See ContentHash_Raw.ToIRI and ContentHash_Graph.ToIRI for more details on specific formatting.
func (ch ContentHash) ToIRI() (string, error) {
	switch hash := ch.Sum.(type) {
	case *ContentHash_Raw_:
		return hash.Raw.ToIRI()
	case *ContentHash_Graph_:
		return hash.Graph.ToIRI()
	default:
		return "", fmt.Errorf("invalid %T type %T", ch, hash)
	}
}

const (
	iriVersion0 byte = 0

	IriPrefixRaw   byte = 0
	IriPrefixGraph byte = 1

	didVersion0 byte = 0
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
func (mt MediaType) ToExtension() (string, error) {
	ext, ok := mediaTypeExtensions[mt]
	if !ok {
		return "", fmt.Errorf("missing extension for %T %s", mt, mt)
	}

	return ext, nil
}

var mediaTypeExtensions = map[MediaType]string{
	MediaType_MEDIA_TYPE_UNSPECIFIED: "bin",
	MediaType_MEDIA_TYPE_TEXT_PLAIN:  "txt",
	MediaType_MEDIA_TYPE_CSV:         "csv",
	MediaType_MEDIA_TYPE_JSON:        "json",
	MediaType_MEDIA_TYPE_XML:         "xml",
	MediaType_MEDIA_TYPE_PDF:         "pdf",
	MediaType_MEDIA_TYPE_TIFF:        "tiff",
	MediaType_MEDIA_TYPE_JPG:         "jpg",
	MediaType_MEDIA_TYPE_PNG:         "png",
	MediaType_MEDIA_TYPE_SVG:         "svg",
	MediaType_MEDIA_TYPE_WEBP:        "webp",
	MediaType_MEDIA_TYPE_AVIF:        "avif",
	MediaType_MEDIA_TYPE_GIF:         "gif",
	MediaType_MEDIA_TYPE_APNG:        "apng",
	MediaType_MEDIA_TYPE_MPEG:        "mpeg",
	MediaType_MEDIA_TYPE_MP4:         "mp4",
	MediaType_MEDIA_TYPE_WEBM:        "webm",
	MediaType_MEDIA_TYPE_OGG:         "ogg",
}

var mediaTypeExtensionsReverse = map[string]MediaType{}

func init() {
	for mt, ext := range mediaTypeExtensions {
		mediaTypeExtensionsReverse[ext] = mt
	}
}

func ParseIRI(methodPrefix string, iri string) (*ContentHash, sdk.AccAddress, error) {
	const regenPrefix = "regen:"
	didPrefix := "did:" + methodPrefix + ":"

	if strings.HasPrefix(iri, regenPrefix) {
		hashExtPart := iri[len(regenPrefix):]
		parts := strings.Split(hashExtPart, ".")
		if len(parts) != 2 {
			return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("error parsing IRI %s", iri))
		}

		hashPart := parts[0]
		ext := parts[1]

		res, version, err := base58.CheckDecode(hashPart)
		if err != nil {
			return nil, nil, err
		}

		if version != iriVersion0 {
			return nil, nil, fmt.Errorf("invalid version found when parsing IRI %s", iri)
		}

		rdr := bytes.NewBuffer(res)

		typ, err := rdr.ReadByte()
		if err != nil {
			return nil, nil, err
		}
		switch typ {
		case IriPrefixRaw:
			b0, err := rdr.ReadByte()
			if err != nil {
				return nil, nil, err
			}

			mediaType, ok := mediaTypeExtensionsReverse[ext]
			if !ok {
				return nil, nil, fmt.Errorf("can't resolve MediaType for extension %s", ext)
			}

			digestAlg := DigestAlgorithm(b0)
			hash := rdr.Bytes()
			err = digestAlg.Validate(hash)
			if err != nil {
				return nil, nil, err
			}

			return &ContentHash{Sum: &ContentHash_Raw_{Raw: &ContentHash_Raw{
				Hash:            hash,
				DigestAlgorithm: digestAlg,
				MediaType:       mediaType,
			}}}, nil, nil

		case IriPrefixGraph:
			if ext != "rdf" {
				return nil, nil, fmt.Errorf("expected extension .rdf for graph data, got .%s", ext)
			}

			b0, err := rdr.ReadByte()
			if err != nil {
				return nil, nil, err
			}

			c14Alg := GraphCanonicalizationAlgorithm(b0)
			err = c14Alg.Validate()
			if err != nil {
				return nil, nil, err
			}

			b0, err = rdr.ReadByte()
			if err != nil {
				return nil, nil, err
			}

			mtAlg := GraphMerkleTree(b0)
			err = mtAlg.Validate()
			if err != nil {
				return nil, nil, err
			}

			b0, err = rdr.ReadByte()
			if err != nil {
				return nil, nil, err
			}

			digestAlg := DigestAlgorithm(b0)
			hash := rdr.Bytes()
			err = digestAlg.Validate(hash)
			if err != nil {
				return nil, nil, err
			}

			return &ContentHash{Sum: &ContentHash_Graph_{Graph: &ContentHash_Graph{
				Hash:                      hash,
				DigestAlgorithm:           digestAlg,
				CanonicalizationAlgorithm: c14Alg,
				MerkleTree:                mtAlg,
			}}}, nil, nil
		}
	} else if strings.HasPrefix(iri, didPrefix) {
		addrPart := iri[len(didPrefix):]
		res, version, err := base58.CheckDecode(addrPart)
		if err != nil {
			return nil, nil, err
		}

		if version != didVersion0 {
			return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid version found when parsing IRI %s", iri))
		}

		return nil, res, nil
	}

	return nil, nil, fmt.Errorf("unable to parse IRI %s", iri)
}
