package impl

type MinHeap struct {
	Root *Node
}

type Node struct {
	Value int
	Parent *Node
	LeftChild *Node
	RightChild *Node
}

func (n *Node) SmallerChild() *Node {
	if n.LeftChild == nil && n.RightChild == nil {
		return nil
	}
	if n.LeftChild == nil && n.RightChild != nil {
		return n.RightChild
	}
	if n.LeftChild != nil && n.RightChild == nil {
		return n.LeftChild
	}
	if n.LeftChild.Value < n.RightChild.Value {
		return n.LeftChild
	}
	return n.RightChild
}

func NewMinHeap(rootValue int) *MinHeap {
	return &MinHeap{&Node{rootValue, nil, nil, nil}}
}

func (h *MinHeap) Add(value int) {
	leaf := h.FindLeaf()
	leaf.LeftChild = &Node{value, leaf, nil, nil}
	h.BubbleUp(leaf.LeftChild)
}

func (h *MinHeap) BubbleUp(leaf *Node) {
	for leaf.Parent != nil && leaf.Parent.Value > leaf.Value {
		h.Swap(leaf, leaf.Parent)
		leaf = leaf.Parent
	}
}

func (h *MinHeap) Empty() bool {
	return h.Root == nil
}

func (h *MinHeap) FindLeaf() *Node {
	node := h.Root
	for true {
		if node.LeftChild != nil {
			node = node.LeftChild
		} else if node.RightChild != nil {
			node = node.RightChild
		} else {
			break
		}
	}
	return node
}

func (h *MinHeap) RemoveLeaf(leaf *Node) {
	if h.Root == leaf {
		h.Root = nil
	} else if leaf.Parent.LeftChild == leaf {
		leaf.Parent.LeftChild = nil
	} else {
		leaf.Parent.RightChild = nil
	}
}

func (h *MinHeap) RemoveMin() int {
	min := h.Root.Value
	leaf := h.FindLeaf()
	h.Swap(h.Root, leaf)
	h.RemoveLeaf(leaf)
	h.SiftDown()
	return min
}

func (h *MinHeap) SiftDown() {
	if h.Empty() {
		return
	}

	node := h.Root
	for true {
		smallerChild := node.SmallerChild()
		if smallerChild != nil && smallerChild.Value < node.Value {
			h.Swap(node, smallerChild)
			node = smallerChild
		} else {
			break
		}
	}
}

func (h *MinHeap) Swap(nodeOne *Node, nodeTwo *Node) {
	temp := nodeOne.Value
	nodeOne.Value = nodeTwo.Value
	nodeTwo.Value = temp
}
