package types

import (
	"fmt"
	"github.com/regen-network/regen-ledger/util"
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
	return util.MustEncodeBech32(Bech32DataAddressPrefix, addr)
}

func (addr DataAddress) URI() *url.URL {
	uri, err := url.Parse(addr.String())
	if err != nil {
		panic(err)
	}
	return uri
}
