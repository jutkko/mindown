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

type Graph struct {
	nodes []*Node
}

// TODO: implement this
func (g *Graph) Export() error {
	return nil
}
