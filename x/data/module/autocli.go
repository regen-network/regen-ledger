package module

import (
	"strings"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	datav1beta1 "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
)

func (am Module) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:              datav1beta1.Query_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "AnchorByIRI",
					Use:       "anchor-by-iri [iri]",
					Short:     "Query a data anchor by the IRI of the data",
					Long:      "Query a data anchor by the IRI of the data.",
					Example: formatExample(`
  regen q data anchor-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "iri"},
					},
				},
				{
					RpcMethod: "AttestationsByAttestor",
					Use:       "attestations-by-attestor [attestor]",
					Short:     "Query data attestations by an attestor",
					Long:      "Query data attestations by an attestor with optional pagination flags.",
					Example: formatExample(`
  regen q data attestations-by-attestor regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza
  regen q data attestations-by-attestor regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza --limit 10 --count-total
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "attestor"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "AttestationsByIRI",
					Use:       "attestations-by-iri [iri]",
					Short:     "Query data attestations by the IRI of the data",
					Long:      "Query data attestations by the IRI of the data with optional pagination flags.",
					Example: formatExample(`
  regen q data attestations-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
  regen q data attestations-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf --limit 10 --count-total
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "iri"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "Resolver",
					Use:       "resolver [id]",
					Short:     "Query a resolver by its unique identifier",
					Long:      "Query a resolver by its unique identifier.",
					Example: formatExample(`
  regen q data resolver 1
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
				{
					RpcMethod: "ResolversByIRI",
					Use:       "resolvers-by-iri [iri]",
					Short:     "Query resolvers with registered data by the IRI of the data",
					Long:      "Query resolvers with registered data by the IRI of the data with optional pagination flags.",
					Example: formatExample(`
  regen q data resolvers-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
  regen q data resolvers-by-iri regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf --limit 10 --count-total
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "iri"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "ResolversByURL",
					Use:       "resolvers-by-url [url]",
					Short:     "Query resolvers by URL",
					Long:      "Query resolvers by URL with optional pagination flags.",
					Example: formatExample(`
  regen q data resolvers-by-url https://foo.bar
  regen q data resolvers-by-url https://foo.bar --limit 10 --count-total
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "url"},
						{ProtoField: "pagination", Optional: true},
					},
				},
				{
					RpcMethod: "ConvertIRIToHash",
					Use:       "convert-iri-to-hash [iri]",
					Short:     "Convert an IRI to a ContentHash",
					Long:      "Convert an IRI to a ContentHash.",
					Example: formatExample(`
  regen q data convert-iri-to-hash regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "iri"},
					},
				},
			},
		},
	}
}

func formatExample(str string) string {
	str = strings.TrimPrefix(str, "\n")
	str = strings.TrimRight(str, "\t")
	return strings.TrimSuffix(str, "\n")
}
