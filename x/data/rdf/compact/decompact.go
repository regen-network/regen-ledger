package compact

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/data/rdf/internal"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type decompactCtx struct {
	resolver InternalIDResolver
	dataset  *CompactDataset
}

func Decompact(resolver InternalIDResolver, dataset *CompactDataset) rdf.QuadIterator {
	ch := make(chan internal.QuadOrErr)

	go func() {
		ctx := decompactCtx{
			resolver: resolver,
			dataset:  dataset,
		}

		for _, node := range dataset.Nodes {
			subject, err := ctx.decompactSubject(node)
			if err != nil {
				ch <- internal.QuadOrErr{Err: err}
			}

			for _, properties := range node.Properties {
				predicate, err := ctx.decompactPredicate(properties)
				if err != nil {
					ch <- internal.QuadOrErr{Err: err}
				}

				for _, objectGraphs := range properties.Objects {

					for _, graphID := range objectGraphs.Graphs {
						graph, err := ctx.decompactGraphID(graphID)
						if err != nil {
							ch <- internal.QuadOrErr{Err: err}
						}

						quad := rdf.NewQuad(subject, predicate, nil, graph)
						ch <- internal.QuadOrErr{Quad: quad}
					}
				}
			}
		}

	}()

	return internal.ChanIterator{Chan: ch}
}

func (ctx decompactCtx) decompactSubject(subject *Node) (rdf.IRIOrBNode, error) {
	ref, err := subjectRef(subject)
	if err != nil {
		return nil, err
	}

	return ctx.resolveRef(ref)
}

func (ctx decompactCtx) decompactPredicate(properties *Properties) (rdf.IRIOrBNode, error) {
	ref, err := predicateRef(properties)
	if err != nil {
		return nil, err
	}

	return ctx.resolveRef(ref)
}

func (ctx decompactCtx) decompactGraphID(id *GraphID) (rdf.IRIOrBNode, error) {
	if id == nil || id.Graph == nil {
		return nil, nil
	}

	ref, err := graphIDRef(id)
	if err != nil {
		return nil, err
	}

	return ctx.resolveRef(ref)
}

func (ctx decompactCtx) resolveRef(ref iriOrBlankNodeRef) (rdf.IRIOrBNode, error) {
	if len(ref.internalId) != 0 {
		return ctx.resolver.GetIRIForID(ref.internalId), nil
	} else if ref.localRef > 0 {
		bnodeId := ref.localRef - 1
		return rdf.BNode(fmt.Sprintf("c14n%d", bnodeId)), nil
	} else if ref.localRef < 0 {
		localIdx := int(-(ref.localRef + 1))
		n := len(ctx.dataset.NewIris)
		if localIdx >= n {
			return nil, fmt.Errorf("local index %d is out of bounds, only %d new IRIs in dataset", localIdx, n)
		}
		return rdf.IRI(ctx.dataset.NewIris[localIdx]), nil
	} else {
		return nil, nil
	}
}

func subjectRef(node *Node) (iriOrBlankNodeRef, error) {
	switch subject := node.Subject.(type) {
	case *Node_InternalId:
		return iriOrBlankNodeRef{internalId: subject.InternalId}, nil
	case *Node_LocalRef:
		return iriOrBlankNodeRef{localRef: subject.LocalRef}, nil
	default:
		return iriOrBlankNodeRef{}, fmt.Errorf("unexpected case %T", subject)
	}
}

func predicateRef(properties *Properties) (iriOrBlankNodeRef, error) {
	switch predicate := properties.Predicate.(type) {
	case *Properties_InternalId:
		return iriOrBlankNodeRef{internalId: predicate.InternalId}, nil
	case *Properties_LocalRef:
		return iriOrBlankNodeRef{localRef: predicate.LocalRef}, nil
	default:
		return iriOrBlankNodeRef{}, fmt.Errorf("unexpected case %T", predicate)
	}
}

func graphIDRef(id *GraphID) (iriOrBlankNodeRef, error) {
	switch graph := id.Graph.(type) {
	case *GraphID_InternalId:
		return iriOrBlankNodeRef{internalId: graph.InternalId}, nil
	case *GraphID_LocalRef:
		return iriOrBlankNodeRef{localRef: graph.LocalRef}, nil
	default:
		return iriOrBlankNodeRef{}, fmt.Errorf("unexpected case %T", graph)
	}
}
