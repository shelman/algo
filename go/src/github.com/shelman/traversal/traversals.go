package main

import (
	"fmt"

	"github.com/shelman/tree"
)

func main() {
	t := &tree.Tree{}

	t.AddRoot("0")
	t.AddNode("1", "0")
	t.AddNode("2", "1")
	t.AddNode("3", "1")
	t.AddNode("4", "0")
	t.AddNode("5", "4")
	t.AddNode("6", "4")

	preorderTraversal(t, printNode)
	inorderTraversal(t, printNode)
	postorderTraversal(t, printNode)
}

func _inorderTraversal(n *tree.Node, apply func(*tree.Node)) {
	children := n.Children()
	if len(children) > 0 {
		_inorderTraversal(children[0], apply)
	}
	apply(n)
	if len(children) > 1 {
		_inorderTraversal(children[1], apply)
	}
}

func inorderTraversal(t *tree.Tree, apply func(*tree.Node)) {
	_inorderTraversal(t.Root, apply)
}

func _preorderTraversal(n *tree.Node, apply func(*tree.Node)) {
	apply(n)
	for _, child := range n.Children() {
		_preorderTraversal(child, apply)
	}
}

func preorderTraversal(t *tree.Tree, apply func(*tree.Node)) {
	_preorderTraversal(t.Root, apply)
}

func _postorderTraversal(n *tree.Node, apply func(*tree.Node)) {
	children := n.Children()
	if len(children) > 0 {
		_postorderTraversal(children[0], apply)
	}
	if len(children) > 1 {
		_postorderTraversal(children[1], apply)
	}
	apply(n)
}

func postorderTraversal(t *tree.Tree, apply func(*tree.Node)) {
	_postorderTraversal(t.Root, apply)
}

func printNode(n *tree.Node) {
	fmt.Printf("%v\n", n)
}
