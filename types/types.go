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
	Bech32GeoAddressPrefix  = "xrn:geo/"
	Bech32DataAddressPrefix = "xrn:g/"
)

type GeoAddress []byte

type DataAddress []byte

const (
	DataAddressPrefixOnChainGraph byte = iota
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
	case DataAddressPrefixOnChainGraph:
		return util.MustEncodeBech32(Bech32DataAddressPrefix, addr[1:])
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

func GetDataAddressOnChainGraph(hash []byte) DataAddress {
	return append([]byte{DataAddressPrefixOnChainGraph}, hash...)
}

func MustDecodeDataURL(url string) DataAddress {
	hrp, bz, err := bech32.DecodeAndConvert(url)
	if err != nil {
		panic(err)
	}
	if hrp == Bech32DataAddressPrefix {
		return GetDataAddressOnChainGraph(bz)
	}
	panic("can't decode data URL")
}
