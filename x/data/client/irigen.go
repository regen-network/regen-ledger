package client

import (
	"crypto"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/piprate/json-gold/ld"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/data/v3"
)

func GenerateIRI() *cobra.Command {
	return &cobra.Command{
		Use:     "generate-iri [filename]",
		Short:   "Creates the content IRI for a file",
		Long:    "Creates the content IRI for a file. If the extension is .jsonld, a graph IRI will be created, otherwise a raw IRI will be created.",
		Example: formatExample(`regen q data generate-iri myfile.ext`),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			filename := args[0]
			ext := filepath.Ext(filename)
			contents, err := os.ReadFile(filename)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}

			ch := &data.ContentHash{}
			switch ext {
			case ".jsonld":
				proc := ld.NewJsonLdProcessor()
				opts := ld.NewJsonLdOptions("")
				opts.Format = "application/n-quads"
				opts.Algorithm = ld.AlgorithmURDNA2015

				var doc map[string]interface{}
				err = json.Unmarshal(contents, &doc)
				if err != nil {
					return fmt.Errorf("failed to unmarshal json: %w", err)
				}

				normalizedTriples, err := proc.Normalize(doc, opts)
				if err != nil {
					return fmt.Errorf("failed to normalize json: %w", err)
				}

				hash, err := blake2b256hash(fmt.Sprintf("%s", normalizedTriples))
				if err != nil {
					return err
				}

				ch.Graph = &data.ContentHash_Graph{
					Hash:                      hash,
					DigestAlgorithm:           uint32(data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
					CanonicalizationAlgorithm: uint32(data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0),
					MerkleTree:                uint32(data.GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED),
				}
			default:
				hash, err := blake2b256hash(string(contents))
				if err != nil {
					return err
				}

				ext = ext[1:] // take the . off the extension

				ch.Raw = &data.ContentHash_Raw{
					Hash:            hash,
					DigestAlgorithm: uint32(data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
					FileExtension:   ext,
				}
			}

			iri, err := ch.ToIRI()
			if err != nil {
				return fmt.Errorf("failed to convert content hash to IRI: %w", err)
			}

			return ctx.PrintString(iri)
		},
	}
}

func blake2b256hash(contents string) ([]byte, error) {
	hasher := crypto.BLAKE2b_256.New()
	_, err := hasher.Write([]byte(contents))
	if err != nil {
		return nil, fmt.Errorf("failed to hash normalized triples: %w", err)
	}
	return hasher.Sum(nil), nil
}
