package types

import "net/url"

type HasURI interface {
	URI() *url.URL
}
