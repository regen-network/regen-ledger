package module

import (
	"strings"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	datav1beta1 "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
)

func (am Module) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              datav1beta1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Anchor",
					Use:       "anchor [iri] --from sender",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "content_hash"},
					},
				},
				{
					RpcMethod: "Attest",
					Use:       "attest [iri] --from sender",
					Short:     `Attest to the veracity of anchored data.`,
					Long: `Attest to the veracity of anchored data. The data MUST be of graph type (rdf file extension).
Attest to the veracity of more than one entry using a comma-separated (no spaces) list of IRIs.`,
					Example: "regen tx data attest regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "content_hashes"},
					},
				},
				{
					RpcMethod: "DefineResolver",
					Use:       "define-resolver [resolver_url]",
					Short:     `Defines a resolver URL and assigns it a new integer ID that can be used in calls to RegisterResolver.`,
					Long: `DefineResolver defines a resolver URL and assigns it a new integer ID that can be used in calls to RegisterResolver.
Parameters:
  resolver_url: resolver_url is a resolver URL which should refer to an HTTP service which will respond to 
			  a GET request with the IRI of a ContentHash and return the content if it exists or a 404.
Flags:
  --from: from flag is the address of the resolver manager
		`,
					Example: "regen tx data define-resolver https://foo.bar --from manager",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "resolver_url"},
					},
				},
				{
					RpcMethod: "RegisterResolver",
					Use:       "register-resolver [resolver_id] [content_hashes_json]",
					Short:     `Registers data content hashes.`,
					Long: `RegisterResolver registers data content hashes.
Parameters:
    resolver_id: resolver id is the ID of a resolver
	content_hashes_json: contains list of content hashes which the resolver claims to serve
Flags:
	--from: manager is the address of the resolver manager
		`,
					Example: `
			regen tx data register-resolver 1 content.json

			where content.json contains
			{
				"content_hashes": [
					{
						"graph": {
							"hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
							"digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
							"canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
							"merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
						}
					}
				]
			}
			`,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "resolver_id"},
						{ProtoField: "content_hashes"},
					},
				},
			},
		},
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: datav1beta1.Query_ServiceDesc.ServiceName,
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
					RpcMethod: "AnchorByHash",
					Use:       "anchor-by-hash [hash-json]",
					Short:     "Query a data anchor by the ContentHash of the data",
					Long:      "Query a data anchor by the ContentHash of the data.",
					Example: formatExample(`
  regen q data anchor-by-hash hash.json

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "content_hash"},
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
					RpcMethod: "AttestationsByHash",
					Use:       "attestations-by-hash [hash-json]",
					Short:     "Query data attestations by the ContentHash of the data",
					Long:      "Query data attestations by the ContentHash of the data with optional pagination flags.",
					Example: formatExample(`
  regen q data attestations-by-hash hash.json
  regen q data attestations-by-hash hash.json --limit 10 --count-total

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "content_hash"},
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
					RpcMethod: "ResolversByHash",
					Use:       "resolvers-by-hash [hash-json]",
					Short:     "Query resolvers with registered data by the ContentHash of the data",
					Long:      "Query resolvers with registered data by the ContentHash of the data with optional pagination flags.",
					Example: formatExample(`
  regen q data resolvers-by-hash hash.json
  regen q data resolvers-by-hash hash.json --limit 10 --count-total

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "content_hash"},
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
				{
					RpcMethod: "ConvertHashToIRI",
					Use:       "convert-hash-to-iri [hash-json]",
					Short:     "Convert a ContentHash to an IRI",
					Long:      "Convert a ContentHash to an IRI.",
					Example: formatExample(`
  regen q data convert-hash-to-iri hash.json

  where hash.json contains:
  {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
      "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
    }
  }
		`),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "content_hash"},
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
