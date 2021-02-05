package compact

import "github.com/regen-network/regen-ledger/x/data/rdf"

type InternalIDRegistrar interface {
	InternalIDResolver

	RegisterIRI(iri rdf.IRI) []byte
}
