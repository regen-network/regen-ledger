package compact

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type InternalIDResolver interface {
	GetIDForIRI(iri rdf.IRI) []byte
	GetIRIForID(id []byte) rdf.IRI
}

type compactCtx struct {
	resolver       InternalIDResolver
	resolvedIRIs   map[rdf.IRIOrBNode]iriOrBlankNodeRef
	dataset        *CompactDataset
	curNode        *Node
	curProperties  *Properties
	curObjectGraph *ObjectGraph
	curSubject     rdf.IRIOrBNode
	curPredicate   rdf.IRIOrBNode
	curObject      rdf.Term
	curGraph       rdf.IRIOrBNode
}

type iriOrBlankNodeRef struct {
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
	ctx.curNode = &Node{}
	ctx.dataset.Nodes = append(ctx.dataset.Nodes, ctx.curNode)

	if len(resolved.internalId) != 0 {
		ctx.curNode.Subject = &Node_InternalId{InternalId: resolved.internalId}
	} else {
		ctx.curNode.Subject = &Node_LocalRef{LocalRef: resolved.localRef}
	}

	return nil
}

func (ctx *compactCtx) compactPredicate(predicate rdf.IRIOrBNode) error {
	resolved, err := ctx.resolveIRIOrBNode(predicate)
	if err != nil {
		return err
	}

	ctx.curPredicate = predicate
	ctx.curProperties = &Properties{}

	if len(resolved.internalId) != 0 {
		ctx.curProperties.Predicate = &Properties_InternalId{InternalId: resolved.internalId}
	} else {
		ctx.curProperties.Predicate = &Properties_LocalRef{LocalRef: resolved.localRef}
	}

	ctx.curNode.Properties = append(ctx.curNode.Properties, ctx.curProperties)

	return nil
}

func (ctx *compactCtx) compactObject(object rdf.Term) error {
	ctx.curObject = object
	ctx.curObjectGraph = &ObjectGraph{}
	ctx.curProperties.Objects = append(ctx.curProperties.Objects, ctx.curObjectGraph)

	switch object := object.(type) {
	case rdf.IRIOrBNode:
		resolved, err := ctx.resolveIRIOrBNode(object)
		if err != nil {
			return err
		}

		if len(resolved.internalId) != 0 {
			ctx.curObjectGraph.Sum = &ObjectGraph_ObjectInternalId{ObjectInternalId: resolved.internalId}
		} else {
			ctx.curObjectGraph.Sum = &ObjectGraph_ObjectLocalRef{ObjectLocalRef: resolved.localRef}
		}

		return nil
	case rdf.Literal:
		datatypeIRI := object.Datatype()
		resolved, err := ctx.resolveIRIOrBNode(datatypeIRI)
		if err != nil {
			return err
		}

		if len(resolved.internalId) != 0 {
			ctx.curObjectGraph.Sum = &ObjectGraph_DataTypeInternalId{DataTypeInternalId: resolved.internalId}
		} else {
			ctx.curObjectGraph.Sum = &ObjectGraph_DataTypeLocalRef{DataTypeLocalRef: resolved.localRef}
		}

		value := object.LexicalForm()
		ctx.curObjectGraph.LiteralValue = &ObjectGraph_StrValue{StrValue: value}

		// TODO: language tag
		// TODO: well known data types + canonical lexical form (maybe Literal.LexicalForm() always returns canonical and this was dealt with in parsing??)
		return nil
	default:
		return fmt.Errorf("unexpected case %T", object)
	}
}

func (ctx *compactCtx) compactGraph(graph rdf.IRIOrBNode) error {
	graphID := &GraphID{}

	// if not default graph
	if graph != nil {
		// named graph
		resolved, err := ctx.resolveIRIOrBNode(graph)
		if err != nil {
			return err
		}

		if len(resolved.internalId) != 0 {
			graphID.Graph = &GraphID_InternalId{InternalId: resolved.internalId}
		} else {
			graphID.Graph = &GraphID_LocalRef{LocalRef: resolved.localRef}
		}
	}

	ctx.curObjectGraph.Graphs = append(ctx.curObjectGraph.Graphs, graphID)

	return nil
}

func (ctx *compactCtx) resolveIRIOrBNode(node rdf.IRIOrBNode) (iriOrBlankNodeRef, error) {
	if resolved, ok := ctx.resolvedIRIs[node]; ok {
		return resolved, nil
	}

	switch node := node.(type) {
	case rdf.IRI:
		var resolved iriOrBlankNodeRef
		id := ctx.resolver.GetIDForIRI(node)
		if len(id) != 0 {
			resolved = iriOrBlankNodeRef{
				internalId: id,
			}
		} else {
			i := len(ctx.dataset.NewIris)
			ctx.dataset.NewIris = append(ctx.dataset.NewIris, string(node))
			resolved = iriOrBlankNodeRef{
				localRef: -int32(i) - 1,
			}
		}
		ctx.resolvedIRIs[node] = resolved
		return resolved, nil
	case rdf.BNode:
		bnode := string(node)
		if !strings.HasPrefix("c14n", bnode) {
			return iriOrBlankNodeRef{}, fmt.Errorf("blank node %s is not canonicalized", bnode)
		}
		bnode = strings.TrimPrefix(bnode, "c14n")
		i, err := strconv.Atoi(bnode)
		if err != nil {
			return iriOrBlankNodeRef{}, err
		}

		return iriOrBlankNodeRef{localRef: int32(i) + 1}, nil
	default:
		return iriOrBlankNodeRef{}, fmt.Errorf("unexpected case %T", node)
	}
}
