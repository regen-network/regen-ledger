package data

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

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
	IRI_VERSION_0 byte = 0

	IRI_PREFIX_RAW   byte = 0
	IRI_PREFIX_GRAPH byte = 1

	DID_VERSION_0 byte = 0
)

func (chr ContentHash_Raw) ToIRI() (string, error) {
	err := chr.Validate()
	if err != nil {
		return "", err
	}

	bz := make([]byte, len(chr.Hash)+2)
	bz[0] = IRI_PREFIX_RAW
	bz[1] = byte(chr.DigestAlgorithm)
	copy(bz[2:], chr.Hash)
	hashStr := base58.CheckEncode(bz, IRI_VERSION_0)

	ext, err := chr.MediaType.ToExtension()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("regen:%s.%s", hashStr, ext), nil
}

func (chg ContentHash_Graph) ToIRI() (string, error) {
	err := chg.Validate()
	if err != nil {
		return "", err
	}

	bz := make([]byte, len(chg.Hash)+4)
	bz[0] = IRI_PREFIX_GRAPH
	bz[1] = byte(chg.CanonicalizationAlgorithm)
	bz[2] = byte(chg.MerkleTree)
	bz[3] = byte(chg.DigestAlgorithm)
	copy(bz[4:], chg.Hash)
	hashStr := base58.CheckEncode(bz, IRI_VERSION_0)

	return fmt.Sprintf("regen:%s.rdf", hashStr), nil
}

//func AccAddressToDID(address sdk.AccAddress, bech32AccPrefix string) string {
//	hash := base58.CheckEncode(address, DID_VERSION_0)
//	return fmt.Sprintf("did:%s:%s", bech32AccPrefix, hash)
//}
//
//func ParseIRI(iri string, bech32AccPrefix string) (IRIDescriptor, error) {
//	if strings.HasPrefix(iri, "regen:") {
//
//	} else if strings.HasPrefix(iri, "did:") {
//		subStr := iri[4:]
//		if !strings.HasPrefix(subStr, bech32AccPrefix + ":") {
//
//		}
//	}
//	return IRIDescriptor{Other: iri}, nil
//}
//
//type IRIDescriptor struct {
//	ContentHash *ContentHash
//	AccAddress  sdk.AccAddress
//	Other       string
//}

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
