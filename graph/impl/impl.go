package impl

import (
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/types"
	"sort"
)

type graphImpl struct {
	rootNode  graph.Node
	nodeNames []types.HasURI
	nodes     map[string]graph.Node
}

// NewGraph creates a new Graph with no nodes
func NewGraph() graph.Graph {
	return &graphImpl{nodeNames: []types.HasURI{}, nodes: make(map[string]graph.Node)}
}

// NewNode creates a new Node with the provided ID
func NewNode(id types.HasURI) graph.Node {
	return &node{id: id, propertyNames: []graph.Property{}, properties: make(map[string]interface{})}
}

type node struct {
	id types.HasURI
	// TODO classes    []string
	propertyNames []graph.Property
	properties    map[string]interface{}
}

func (n *node) Classes() []graph.Class {
	panic("classes aren't supported yet")
}

type sortNodeNames []types.HasURI

func (s sortNodeNames) Len() int {
	return len(s)
}

func (s sortNodeNames) Less(i, j int) bool {
	return s[i].String() < s[j].String()
}

func (s sortNodeNames) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (g *graphImpl) WithNode(node graph.Node) {
	if node.ID() == nil {
		g.rootNode = node
	} else {
		key := node.ID().String()
		_, found := g.nodes[key]
		if !found {
			g.nodeNames = append(g.nodeNames, node.ID())
			sort.Sort(sortNodeNames(g.nodeNames))
		}
		g.nodes[key] = node
	}
}

func (g *graphImpl) WithoutNode(id types.HasURI) {
	if id == nil {
		g.rootNode = nil
	} else {
		key := id.String()
		_, found := g.nodes[key]
		if !found {
			return
		}
		delete(g.nodes, key)
		nNodes := len(g.nodeNames)
		i := sort.Search(nNodes, func(i int) bool {
			return g.nodeNames[i].String() >= key
		})
		if i < nNodes && g.nodeNames[i].String() == key {
			g.nodeNames = append(g.nodeNames[:i], g.nodeNames[i+1:]...)
		}
	}
}

func (n *node) SetID(id types.HasURI) {
	n.id = id
}

type sortProperties []graph.Property

func (s sortProperties) Len() int {
	return len(s)
}

func (s sortProperties) Less(i, j int) bool {
	return s[i].URI().String() < s[j].URI().String()
}

func (s sortProperties) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (n *node) SetProperty(property graph.Property, value interface{}) {
	// TODO validate value
	key := property.URI().String()
	_, found := n.properties[key]
	if !found {
		n.propertyNames = append(n.propertyNames, property)
		sort.Sort(sortProperties(n.propertyNames))
	}
	n.properties[key] = value
}

func (n *node) DeleteProperty(property graph.Property) {
	panic("implement me")
}

func (n *node) String() string {
	s, err := graph.CanonicalNodeString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func (g *graphImpl) RootNode() graph.Node {
	return g.rootNode
}

func (g *graphImpl) Nodes() []types.HasURI {
	return g.nodeNames
}

func (g *graphImpl) GetNode(url types.HasURI) graph.Node {
	return g.nodes[url.String()]
}

func (n *node) ID() types.HasURI {
	return n.id
}

func (n *node) Properties() []graph.Property {
	return n.propertyNames
}

func (n *node) GetProperty(property graph.Property) interface{} {
	return n.properties[property.URI().String()]
}

func (g *graphImpl) String() string {
	s, err := graph.CanonicalString(g)
	if err != nil {
		panic(err)
	}
	return s
}
