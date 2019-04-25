package types

import (
	"fmt"
	"github.com/regen-network/regen-ledger/util"
	"github.com/tendermint/tendermint/libs/bech32"
	"net/url"
)

type HasURI interface {
	fmt.Stringer
	// URI returns the URI representation of the underlying type
	URI() *url.URL
}

const (
	Bech32GeoAddressPrefix       = "xrn:geo/"
	Bech32GraphDataAddressPrefix = "xrn:g/"
	Bech32RawDataAddressPrefix   = "xrn:d/"
)

type GeoAddress []byte

type DataAddress []byte

const (
	DataAddressPrefixGraph byte = iota
	DataAddressPrefixRawData
)

// String returns the string URI representation mof the GeoAddress
func (addr GeoAddress) String() string {
	return util.MustEncodeBech32(Bech32GeoAddressPrefix, addr)
}

// URI returns the URI representation mof the GeoAddress
func (addr GeoAddress) URI() *url.URL {
	uri, err := url.Parse(addr.String())
	if err != nil {
		panic(err)
	}
	return uri
}

func (addr DataAddress) String() string {
	switch addr[0] {
	case DataAddressPrefixGraph:
		return util.MustEncodeBech32(Bech32GraphDataAddressPrefix, addr[1:])
	case DataAddressPrefixRawData:
		return util.MustEncodeBech32(Bech32RawDataAddressPrefix, addr[1:])
	default:
		panic(fmt.Errorf("unknown address prefix %d", addr[0]))

	}
}

func (addr DataAddress) URI() *url.URL {
	uri, err := url.Parse(addr.String())
	if err != nil {
		panic(err)
	}
	return uri
}

func GetDataAddressGraph(hash []byte) DataAddress {
	return append([]byte{DataAddressPrefixGraph}, hash...)
}

func GetDataAddressRawData(sha256hash []byte) DataAddress {
	return append([]byte{DataAddressPrefixRawData}, sha256hash...)
}

func MustDecodeBech32DataAddress(url string) DataAddress {
	addr, err := DecodeBech32DataAddress(url)
	if err != nil {
		panic(err)
	}
	return addr
}

func DecodeBech32DataAddress(url string) (DataAddress, error) {
	hrp, bz, err := bech32.DecodeAndConvert(url)
	if err != nil {
		return nil, err
	}
	if hrp == Bech32GraphDataAddressPrefix {
		return GetDataAddressGraph(bz), nil
	}
	if hrp == Bech32RawDataAddressPrefix {
		return GetDataAddressRawData(bz), nil
	}
	return nil, fmt.Errorf("can't decode data URL")
}

// IsGraphDataAddress indicates whether the provided DataAddress points to graph
// data - which has a well-known structure conformant with the schema module -
// as opposed to "raw" data which can have any format
func IsGraphDataAddress(addr DataAddress) bool {
	switch addr[0] {
	case DataAddressPrefixOnChainGraph:
		return true
	default:
		return false
	}
}

// IsRawDataAddress indicates whether the provided DataAddress points to raw
// data - i.e. data in any format - as opposed to well-structured graph data
func IsRawDataAddress(addr DataAddress) bool {
	switch addr[0] {
	case DataAddressPrefixOnChainGraph:
		return false
	default:
		return false
	}
}
