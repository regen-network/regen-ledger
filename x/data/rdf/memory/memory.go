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

	triple := &rdf.Triple{
		Subject:   sub,
		Predicate: nil,
		Object:    nil,
	}

	ch := make(chan bool)
	go func() {
		for k, vs := range props.props {
			triple.Predicate = k
			for v := range vs {
				triple.Object = v
				ch <- true
			}
		}
		close(ch)
	}()

	return &chIterator{ch: ch, triple: triple}
}

func (m memoryGraph) FindByPredicate(pred rdf.IRIOrBNode) rdf.TripleIterator {
	subMap, ok := m.predSubMap[pred]
	if !ok {
		return &chIterator{}
	}

	triple := &rdf.Triple{
		Subject:   nil,
		Predicate: pred,
		Object:    nil,
	}

	ch := make(chan bool)
	go func() {
		for k, vs := range subMap {
			triple.Subject = k
			for _, vs := range vs.props {
				for v := range vs {
					triple.Object = v
					ch <- true
				}
			}
		}
		close(ch)
	}()

	return &chIterator{ch: ch, triple: triple}
}

func (m memoryGraph) FindByObject(obj rdf.Term) rdf.TripleIterator {
	panic("implement me")
}

func (m memoryGraph) FindBySubjectPredicate(sub rdf.IRIOrBNode, pred rdf.IRIOrBNode) rdf.TripleIterator {
	panic("implement me")
}

func (m memoryGraph) FindBySubjectObject(sub rdf.IRIOrBNode, obj rdf.Term) rdf.TripleIterator {
	panic("implement me")
}

func (m memoryGraph) FindByPredicateObject(pred rdf.IRIOrBNode, obj rdf.Term) rdf.TripleIterator {
	panic("implement me")
}

func (m memoryGraph) FindAll() rdf.TripleIterator {
	panic("implement me")
}

//type propTarget struct {
//	sub rdf.IRIOrBNode
//	// obj is the target IRI of the prop and empty if the property targets a literal
//	obj rdf.IRIOrBNode
//}

type chIterator struct {
	ch     chan bool
	triple *rdf.Triple
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

	_, ok := <-it.ch

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

func (m *memoryGraph) Merge(graph rdf.Graph) {
	it := graph.FindAll()
	defer it.Close()

	for it.Next() {
		m.AddTriple(rdf.Triple{
			Subject:   it.Subject(),
			Predicate: it.Predicate(),
			Object:    it.Object(),
		})
	}
}

func (m *memoryGraph) NewBNode() rdf.BNode {
	m.bnodeId = m.bnodeId + 1
	return rdf.BNode(fmt.Sprintf("%d", m.bnodeId))
}

//
