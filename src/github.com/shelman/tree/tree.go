package tree

type Tree struct {
	Root *Node
	nodes []*Node
}

type Node struct {
	name string
	children []*Node
}

func (n *Node) AddChild(name string) {
	if len(n.children) > 1 {
		panic("Can only have 2 children")
	}
	n.children = append(n.children, &Node{name, nil})
}

func (n *Node) Children() []*Node {
	return n.children
}

func (t *Tree) AddNode(name string, parent string) {
	for _, node := range t.nodes {
		if node.name == parent {
			node.AddChild(name)
			t.nodes = append(t.nodes, node.Children()[len(node.Children())-1])
			return
		}
	}
	panic("Couldn't find parent")
}

func (t *Tree) AddRoot(name string) {
	if t.Root != nil {
		panic("Cannot double up on root")
	}
	t.Root = &Node{name, nil}
	t.nodes = []*Node{t.Root}
}