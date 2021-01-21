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

func validate(shapesGraph shapesGraph, dataGraph indexedGraph) {
	for _, ns := range shapesGraph.nodeShapes {
		for _, tc := range ns.targetClass {

		}
	}
}
