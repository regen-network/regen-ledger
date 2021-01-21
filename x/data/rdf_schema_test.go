package data

type iri string

type indexedGraph struct {
	// nodeMap maps node IRIs to nodeProps
	nodeMap map[iri]nodeProps

	// propMap maps properties to propTargets for quickly initiating sh:targetSubectsOf and sh:targetObjectsOf
	propMap map[iri][]propTarget

	// classMap indexes nodes satisfying each rdfs:class for quickly initiating sh:targetClass
	classMap map[iri][]iri
}

type nodeProps struct {
	// rdfs:class values
	classes []iri
	// other properties mapped to their values
	props map[iri]map[interface{}]bool
}

type propTarget struct {
	sub iri
	// obj is the target IRI of the prop and empty if the property targets a literal
	obj iri
}

type shapesGraph struct {
	nodeShapes []nodeShape
}

type nodeShape struct {
	targetClass      []iri
	targetSubjectsOf []iri
	targetObjectsOf  []iri
	properties       []propertyShape
}

type propertyShape struct {
}

type validationContext struct {
	shapesGraph shapesGraph
	dataGraph   indexedGraph
}

func (ctx *validationContext) validate() {
	// can be run in parallel with a WaitGroup
	for _, ns := range ctx.shapesGraph.nodeShapes {
		// can be run in parallel with a WaitGroup
		for _, tc := range ns.targetClass {
			// can be run in parallel with a WaitGroup
			for _, targetIRI := range ctx.dataGraph.classMap[tc] {
				ctx.evalNodeShape(ns, targetIRI)
			}
		}
	}
}

func (ctx *validationContext) evalNodeShape(ns nodeShape, targetIRI iri) {
	np := ctx.dataGraph.nodeMap[targetIRI]
	// can be run in parallel with a WaitGroup
	for _, prop := range ns.properties {
		ctx.evalPropertyShape(prop, np)
	}
}

func (ctx *validationContext) evalPropertyShape(ps propertyShape, np nodeProps) {

}

type IRI string

type IndexedGraph interface {
	PredicateObjectAccessor(subject IRI) PredicateObjectAccessor
	SubectObjectAccessor(predicate IRI) SubjectObectAccessor
}

type PredicateObjectAccessor interface {
	ObjectAccessor(predicate IRI) ObjectAccessor
	Iterator() PredicateObjectIterator
}

type PredicateObjectIterator interface {
	Predicate() IRI
	Object() ObjectAccessor
	Next() bool
}

type ObjectAccessor interface {
	HasValue(interface{}) bool
	Iterator() ValueIterator
}

type ValueIterator interface {
	Current() interface{}
	Next() bool
}

type SubjectObectAccessor interface {
	ObjectAccesor(subject IRI) ObjectAccessor
	Iterator() SubjectObectIterator
}

type SubjectObectIterator interface {
	Subject() IRI
	Object() interface{}
	Next() bool
}
