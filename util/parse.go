package util

type Node struct {
	children []*Node
	title    string
	value    string
}

func NewNode(title, value string) *Node {
	return &Node{
		title: title,
		value: value,
	}
}

func (n *Node) AddChild(childNode *Node) {
	n.children = append(n.children, childNode)
}

func (n *Node) GetChildren() []*Node {
	return n.children
}

func (n *Node) GetTitle() string {
	return n.title
}

type Graph struct {
	nodes []*Node
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) GetNodes() []*Node {
	return g.nodes
}
