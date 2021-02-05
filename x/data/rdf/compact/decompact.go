package compact

import "github.com/regen-network/regen-ledger/x/data/rdf"

func Decompact(resolver InternalIDResolver, dataset CompactDataset) ([]*rdf.Quad, error) {
	var quads []*rdf.Quad

	for _, node := range dataset.Nodes {
		for _, properties := range node.Properties {
			for _, objectGraphs := range properties.Objects {

			}
		}
	}
}
