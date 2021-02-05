package compact

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type InternalIDResolver interface {
	ResolveID(iri rdf.IRI) []byte
}

type compactCtx struct {
	resolver       InternalIDResolver
	resolvedIRIs   map[rdf.IRIOrBNode]resolvedIRIOrBlankNode
	dataset        *data.CompactDataset
	curNode        *data.CompactDataset_Node
	curProperties  *data.CompactDataset_Properties
	curObjectGraph *data.CompactDataset_ObjectGraph
	curSubject     rdf.IRIOrBNode
	curPredicate   rdf.IRIOrBNode
	curObject      rdf.Term
	curGraph       rdf.IRIOrBNode
}

type resolvedIRIOrBlankNode struct {
	internalId []byte
	localRef   int32
}

func (ctx *compactCtx) compactQuad(quad *rdf.Quad) error {
	if !quad.Subject.Equal(ctx.curSubject) {
		err := ctx.compactSubject(quad.Subject)
		if err != nil {
			return err
		}
	}

	if !quad.Predicate.Equal(ctx.curPredicate) {
		err := ctx.compactPredicate(quad.Predicate)
		if err != nil {
			return err
		}
	}

	if !quad.Object.Equal(ctx.curObject) {
		err := ctx.compactObject(quad.Object)
		if err != nil {
			return err
		}
	}

	if !quad.Graph.Equal(ctx.curGraph) {
		err := ctx.compactGraph(quad.Graph)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctx *compactCtx) compactSubject(subject rdf.IRIOrBNode) error {
	resolved, err := ctx.resolveIRIOrBNode(subject)
	if err != nil {
		return err
	}

	ctx.curSubject = subject
	ctx.curPredicate = nil
	ctx.curObject = nil
	ctx.curGraph = nil
	ctx.curNode = &data.CompactDataset_Node{}
	ctx.dataset.Nodes = append(ctx.dataset.Nodes, ctx.curNode)

	if len(resolved.internalId) != 0 {
		ctx.curNode.Subject = &data.CompactDataset_Node_InternalId{InternalId: resolved.internalId}
	} else {
		ctx.curNode.Subject = &data.CompactDataset_Node_LocalRef{LocalRef: resolved.localRef}
	}

	return nil
}

func (ctx *compactCtx) compactPredicate(predicate rdf.IRIOrBNode) error {
	resolved, err := ctx.resolveIRIOrBNode(predicate)
	if err != nil {
		return err
	}

	ctx.curPredicate = predicate
	ctx.curProperties = &data.CompactDataset_Properties{}

	if len(resolved.internalId) != 0 {
		ctx.curProperties.Predicate = &data.CompactDataset_Properties_InternalId{InternalId: resolved.internalId}
	} else {
		ctx.curProperties.Predicate = &data.CompactDataset_Properties_LocalRef{LocalRef: resolved.localRef}
	}

	ctx.curNode.Properties = append(ctx.curNode.Properties, ctx.curProperties)

	return nil
}

func (ctx *compactCtx) compactObject(object rdf.Term) error {
	ctx.curObject = object
	ctx.curObjectGraph = &data.CompactDataset_ObjectGraph{}
	ctx.curProperties.Objects = append(ctx.curProperties.Objects, ctx.curObjectGraph)

	switch object := object.(type) {
	case rdf.IRIOrBNode:
		resolved, err := ctx.resolveIRIOrBNode(object)
		if err != nil {
			return err
		}

		if len(resolved.internalId) != 0 {
			ctx.curProperties.Predicate = &data.CompactDataset_Properties_InternalId{InternalId: resolved.internalId}
		} else {
			ctx.curProperties.Predicate = &data.CompactDataset_Properties_LocalRef{LocalRef: resolved.localRef}
		}

		return nil
	case rdf.Literal:
		return fmt.Errorf("not implemented")
	default:
		return fmt.Errorf("unexpected case %T", object)
	}
}

func (ctx *compactCtx) compactGraph(graph rdf.IRIOrBNode) error {
	graphID := &data.CompactDataset_ObjectGraph_GraphID{}

	// if not default graph
	if graph != nil {
		// named graph
		resolved, err := ctx.resolveIRIOrBNode(graph)
		if err != nil {
			return err
		}

		if len(resolved.internalId) != 0 {
			graphID.Graph = &data.CompactDataset_ObjectGraph_GraphID_InternalId{InternalId: resolved.internalId}
		} else {
			graphID.Graph = &data.CompactDataset_ObjectGraph_GraphID_LocalRef{LocalRef: resolved.localRef}
		}
	}

	ctx.curObjectGraph.Graphs = append(ctx.curObjectGraph.Graphs, graphID)

	return nil
}

func (ctx *compactCtx) resolveIRIOrBNode(node rdf.IRIOrBNode) (resolvedIRIOrBlankNode, error) {
	if resolved, ok := ctx.resolvedIRIs[node]; ok {
		return resolved, nil
	}

	switch node := node.(type) {
	case rdf.IRI:
		var resolved resolvedIRIOrBlankNode
		id := ctx.resolver.ResolveID(node)
		if len(id) != 0 {
			resolved = resolvedIRIOrBlankNode{
				internalId: id,
			}
		} else {
			i := len(ctx.dataset.NewIris)
			ctx.dataset.NewIris = append(ctx.dataset.NewIris, string(node))
			resolved = resolvedIRIOrBlankNode{
				localRef: -int32(i),
			}
		}
		ctx.resolvedIRIs[node] = resolved
		return resolved, nil
	case rdf.BNode:
		bnode := string(node)
		if !strings.HasPrefix("c14n", bnode) {
			return resolvedIRIOrBlankNode{}, fmt.Errorf("blank node %s is not canonicalized", bnode)
		}
		bnode = strings.TrimPrefix(bnode, "c14n")
		i, err := strconv.Atoi(bnode)
		if err != nil {
			return resolvedIRIOrBlankNode{}, err
		}

		return resolvedIRIOrBlankNode{localRef: int32(i)}, nil
	default:
		return resolvedIRIOrBlankNode{}, fmt.Errorf("unexpected case %T", node)
	}
}
