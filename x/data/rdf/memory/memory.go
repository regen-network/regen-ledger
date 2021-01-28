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
	}
}

type memoryGraph struct {
	ctx rdf.Context

	bnodeId uint64

	// nodeMap maps node IRIs to nodeProps
	nodeMap map[rdf.IRIOrBNode]*nodeProps

	predSubMap map[rdf.IRIOrBNode]map[rdf.IRIOrBNode]*nodeProps
}

type nodeProps struct {
	// other properties mapped to their values
	props map[rdf.IRIOrBNode]map[rdf.Term]bool
}

func (m memoryGraph) HasTriple(sub rdf.IRIOrBNode, pred rdf.IRIOrBNode, obj rdf.Term) bool {
	props, ok := m.nodeMap[sub]
	if !ok {
		return false
	}

	objs, ok := props.props[pred]
	if !ok {
		return false
	}

	return objs[obj]
}

func (m memoryGraph) FindBySubject(sub rdf.IRIOrBNode) rdf.TripleIterator {
	props, ok := m.nodeMap[sub]
	if !ok {
		return &chIterator{}
	}

	ch := make(chan rdf.Triple)
	go func() {
		for k, vs := range props.props {
			for v := range vs {
				ch <- rdf.Triple{
					Subject:   sub,
					Predicate: k,
					Object:    v,
				}
			}
		}
		close(ch)
	}()

	return &chIterator{ch: ch}
}

func (m memoryGraph) FindByPredicate(pred rdf.IRIOrBNode) rdf.TripleIterator {
	subMap, ok := m.predSubMap[pred]
	if !ok {
		return &chIterator{}
	}

	ch := make(chan rdf.Triple)
	go func() {
		for k, vs := range subMap {
			for _, vs := range vs.props {
				for v := range vs {
					ch <- rdf.Triple{
						Subject:   k,
						Predicate: pred,
						Object:    v,
					}
				}
			}
		}
		close(ch)
	}()

	return &chIterator{ch: ch}
}

func (m memoryGraph) FindByObject(obj rdf.Term) rdf.TripleIterator {
	return rdf.Filter(m.FindAll(), func(_ rdf.IRIOrBNode, _ rdf.IRIOrBNode, o rdf.Term) bool {
		return obj.Equals(m.ctx, o)
	})
}

func (m memoryGraph) FindBySubjectPredicate(sub rdf.IRIOrBNode, pred rdf.IRIOrBNode) rdf.TripleIterator {
	props, ok := m.nodeMap[sub]
	if !ok {
		return &chIterator{}
	}

	vs := props.props[pred]
	if vs == nil {
		return &chIterator{}
	}

	ch := make(chan rdf.Triple)
	go func() {
		for v := range vs {
			ch <- rdf.Triple{
				Subject:   sub,
				Predicate: pred,
				Object:    v,
			}
		}
		close(ch)
	}()

	return &chIterator{ch: ch}
}

func (m memoryGraph) FindBySubjectObject(sub rdf.IRIOrBNode, obj rdf.Term) rdf.TripleIterator {
	return rdf.Filter(m.FindBySubject(sub), func(_ rdf.IRIOrBNode, _ rdf.IRIOrBNode, o rdf.Term) bool {
		return obj.Equals(m.ctx, o)
	})
}

func (m memoryGraph) FindByPredicateObject(pred rdf.IRIOrBNode, obj rdf.Term) rdf.TripleIterator {
	return rdf.Filter(m.FindByPredicate(pred), func(_ rdf.IRIOrBNode, _ rdf.IRIOrBNode, o rdf.Term) bool {
		return obj.Equals(m.ctx, o)
	})
}

func (m memoryGraph) FindAll() rdf.TripleIterator {
	triple := &rdf.Triple{}

	ch := make(chan rdf.Triple)
	go func() {
		for s, po := range m.nodeMap {
			triple.Subject = s
			for p, os := range po.props {
				triple.Predicate = p
				for o := range os {
					triple.Object = o
					ch <- rdf.Triple{
						Subject:   s,
						Predicate: p,
						Object:    o,
					}
				}
			}
		}
		close(ch)
	}()

	return &chIterator{ch: ch}
}

type chIterator struct {
	ch     chan rdf.Triple
	triple rdf.Triple
}

func (it chIterator) Count() int {
	panic("implement me")
}

func (it chIterator) CountGTE(i int) bool {
	panic("implement me")
}

func (it chIterator) CountLTE(i int) bool {
	panic("implement me")
}

func (it *chIterator) Next() bool {
	if it.ch == nil {
		return false
	}

	var ok bool
	it.triple, ok = <-it.ch

	return ok
}

func (it chIterator) Subject() rdf.IRIOrBNode {
	return it.triple.Subject
}

func (it chIterator) Predicate() rdf.IRIOrBNode {
	return it.triple.Predicate
}

func (it chIterator) Object() rdf.Term {
	return it.triple.Object
}

func (it chIterator) Close() {}

var _ rdf.Graph = memoryGraph{}

var _ rdf.GraphBuilder = &memoryGraph{}

func (m *memoryGraph) AddTriple(sub rdf.IRIOrBNode, pred rdf.IRIOrBNode, obj rdf.Term) {
	props, ok := m.nodeMap[sub]
	if !ok {
		props = &nodeProps{
			props: map[rdf.IRIOrBNode]map[rdf.Term]bool{},
		}
		m.nodeMap[sub] = props
	}

	objs, ok := props.props[pred]
	if !ok {
		objs = map[rdf.Term]bool{}
		props.props[pred] = objs
	}

	objs[obj] = true

	// predSubMap index
	if m.predSubMap[pred] == nil {
		m.predSubMap[pred] = map[rdf.IRIOrBNode]*nodeProps{}
	}
	m.predSubMap[pred][sub] = props
}

func (m *memoryGraph) RemoveTriple(sub rdf.IRIOrBNode, pred rdf.IRIOrBNode, obj rdf.Term) {
	props, ok := m.nodeMap[sub]
	if !ok {
		return
	}

	objs, ok := props.props[pred]
	if !ok {
		return
	}

	delete(objs, obj)

	// predSubMap index
	delete(m.predSubMap[pred], sub)
}

func (m *memoryGraph) NewBNode() rdf.BNode {
	m.bnodeId = m.bnodeId + 1
	return rdf.BNode(fmt.Sprintf("%d", m.bnodeId))
}
