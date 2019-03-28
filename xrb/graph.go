package xrb

type Graph interface {
	Nodes() []string
	RootNode() Node
	GetNode(url string) Node
}

type Node interface {
	ID() string
	Properties() []string
	GetProperty(url string) interface{}
	// TODO GetClasses() []string
}

type graph struct {
	rootNode  *node
	nodeNames []string
	nodes     map[string]*node
}

type node struct {
	id string
	// TODO classes    []string
	propertyNames []string
	properties    map[string]interface{}
}

func (g graph) Nodes() []string {
	return g.nodeNames
}

func (g graph) RootNode() Node {
	return g.rootNode
}

func (g graph) GetNode(url string) Node {
	return g.nodes[url]
}

func (n node) ID() string {
	return n.ID()
}

func (n node) Properties() []string {
	return n.propertyNames
}

func (n node) GetProperty(url string) interface{} {
	return n.properties[url]
}

//func (n node) GetClasses() []string {
//	panic("implement me")
//}
