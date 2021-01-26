package memory

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

func NewGraph() rdf.GraphBuilder {
	return &memoryGraph{
		bnodeId:    0,
		nodeMap:    map[rdf.IRIOrBNode]*nodeProps{},
		predSubMap: map[rdf.IRIOrBNode]map[rdf.IRIOrBNode]*nodeProps{},
		predObjMap: map[rdf.IRIOrBNode]map[rdf.IRIOrBNode]bool{},
	}
}

type memoryGraph struct {
	bnodeId uint64

	// nodeMap maps node IRIs to nodeProps
	nodeMap map[rdf.IRIOrBNode]*nodeProps

	predSubMap map[rdf.IRIOrBNode]map[rdf.IRIOrBNode]*nodeProps

	predObjMap map[rdf.IRIOrBNode]map[rdf.IRIOrBNode]bool
}

type nodeProps struct {
	// other properties mapped to their values
	props map[rdf.IRIOrBNode]map[rdf.Term]bool
}

//type propTarget struct {
//	sub rdf.IRIOrBNode
//	// obj is the target IRI of the prop and empty if the property targets a literal
//	obj rdf.IRIOrBNode
//}

var _ rdf.IndexedGraph = memoryGraph{}

func (m memoryGraph) HasTriple(triple rdf.Triple) bool {
	props, ok := m.nodeMap[triple.Subject]
	if !ok {
		return false
	}

	objs, ok := props.props[triple.Predicate]
	if !ok {
		return false
	}

	return objs[triple.Object]
}

func (m memoryGraph) Iterator() rdf.GraphIterator {
	panic("implement me")
}

var _ rdf.GraphBuilder = &memoryGraph{}

func (m *memoryGraph) AddTriple(triple rdf.Triple) {
	props, ok := m.nodeMap[triple.Subject]
	if !ok {
		props = &nodeProps{
			props: map[rdf.IRIOrBNode]map[rdf.Term]bool{},
		}
		m.nodeMap[triple.Subject] = props
	}

	objs, ok := props.props[triple.Predicate]
	if !ok {
		objs = map[rdf.Term]bool{}
		props.props[triple.Predicate] = objs
	}

	objs[triple.Object] = true

	// predSubMap index
	if m.predSubMap[triple.Predicate] == nil {
		m.predSubMap[triple.Predicate] = map[rdf.IRIOrBNode]*nodeProps{}
	}
	m.predSubMap[triple.Predicate][triple.Subject] = props

	// predObjMap index
	if iriOrBNode, ok := triple.Object.(rdf.IRIOrBNode); ok {
		if m.predObjMap[triple.Predicate] == nil {
			m.predObjMap[triple.Predicate] = map[rdf.IRIOrBNode]bool{}
		}
		m.predObjMap[triple.Predicate][iriOrBNode] = true
	}
}

func (m *memoryGraph) RemoveTriple(triple rdf.Triple) {
	props, ok := m.nodeMap[triple.Subject]
	if !ok {
		return
	}

	objs, ok := props.props[triple.Predicate]
	if !ok {
		return
	}

	delete(objs, triple.Object)

	// predSubMap index
	delete(m.predSubMap[triple.Predicate], triple.Subject)

	// predObjMap index
	if iriOrBNode, ok := triple.Object.(rdf.IRIOrBNode); ok {
		delete(m.predObjMap[triple.Predicate], iriOrBNode)
	}
}

func (m *memoryGraph) Merge(graph rdf.IndexedGraph) {
	it := graph.Iterator()
	for it.Next() {
		propIt := it.Properties().Iterator()

		for propIt.Next() {
			objIt := propIt.Object().Iterator()

			for objIt.Next() {
				m.AddTriple(rdf.Triple{
					Subject:   it.Subject(),
					Predicate: propIt.Predicate(),
					Object:    objIt.Object(),
				})
			}
		}
	}
}

func (m *memoryGraph) NewBNode() rdf.BNode {
	m.bnodeId = m.bnodeId + 1
	return rdf.BNode(fmt.Sprintf("%d", m.bnodeId))
}

func (m memoryGraph) BySubject(subject rdf.Node) rdf.PredicateObjectAccessor {
	return predObjAcc{nodeProps: m.nodeMap[subject]}
}

type predObjAcc struct {
	*nodeProps
}

func (p predObjAcc) ByPredicate(predicate rdf.Node) rdf.ObjectAccessor {
	if p.props == nil {
		return objAcc{}
	}

	return objAcc{objs: p.props[predicate]}
}

type objAcc struct {
	objs map[rdf.Term]bool
}

func (o objAcc) HasValue(term rdf.Term) bool {
	if o.objs == nil {
		return false
	}

	return o.objs[term]
}

func (o objAcc) Iterator() rdf.ObjectIterator {
	if o.objs == nil {
		return &objIterator{}
	}

	ch := make(chan rdf.Term)
	go func() {
		for k, _ := range o.objs {
			ch <- k
		}
		close(ch)
	}()
	return &objIterator{ch: ch}
}

type objIterator struct {
	ch  chan rdf.Term
	obj rdf.Term
}

func (o *objIterator) Next() bool {
	if o.ch == nil {
		return false
	}

	obj, ok := <-o.ch
	if !ok {
		return false
	}

	o.obj = obj
	return true
}

func (o *objIterator) Object() rdf.Term {
	return o.obj
}

func (p predObjAcc) Iterator() rdf.PredicateObjectIterator {
	if p.props == nil {
		return &predObjIterator{}
	}

	ch := make(chan *predObjPair)
	go func() {
		for k, v := range p.props {
			ch <- &predObjPair{
				pred: k,
				objs: v,
			}
		}
		close(ch)
	}()
	return &predObjIterator{ch: ch}
}

type predObjIterator struct {
	ch   chan *predObjPair
	pair *predObjPair
}

type predObjPair struct {
	pred rdf.IRIOrBNode
	objs map[rdf.Term]bool
}

func (p predObjIterator) Next() bool {
	if p.ch == nil {
		return false
	}

	pair, ok := <-p.ch
	if !ok {
		return false
	}

	p.pair = pair
	return true
}

func (p predObjIterator) Predicate() rdf.Node {
	if p.pair == nil {
		return nil
	}

	return p.pair.pred
}

func (p predObjIterator) Object() rdf.ObjectAccessor {
	if p.pair == nil {
		return nil
	}

	return objAcc{p.pair.objs}
}

//func (m memoryGraph) ByPredicate(predicate rdf.Node) rdf.SubjectObectAccessor {
//	return subObjAcc{subObjs: m.propMap[predicate]}
//}
//
//type subObjAcc struct {
//	subObjs []*propTarget
//}
//
//func (s subObjAcc) Iterator() rdf.SubjectObectIterator {
//	if s.subObjs == nil {
//		return &subObjIterator{}
//	}
//}
//
//type subObjIterator struct {
//}
//
//func (s subObjIterator) Next() bool {
//	panic("implement me")
//}
//
//func (s subObjIterator) Subject() rdf.Node {
//	panic("implement me")
//}
//
//func (s subObjIterator) Object() rdf.Term {
//	panic("implement me")
//}
