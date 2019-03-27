package xrb

type Graph interface {
	Nodes() []string
	RootNode() Node
	GetNode(id string) Node
}

type Node interface {
	ID() string
	Properties() []string
	GetClasses() []string
	GetProperty(url string) interface{}
}
