package main

import (
	// "golang.org/x/exp/errors/fmt"
	"fmt"
	"regexp"
)

func (node *TreeNode) preOrderDFS() {
	if node == nil {
		return
	}
	_, _ = fmt.Print(node.label, " ")
	if node.left != nil {
		node.left.preOrderDFS()
	}
	if node.right != nil {
		node.right.preOrderDFS()
	}
}

func (node *TreeNode) inOrderDFS() {
	if node == nil {
		return
	}
	if node.left != nil {
		node.left.inOrderDFS()
	}
	_, _ = fmt.Print(node.label, " ")
	if node.right != nil {
		node.right.inOrderDFS()
	}
}

func (node *TreeNode) postOrderDFS() {
	if node == nil {
		return
	}
	if node.left != nil {
		node.left.postOrderDFS()
	}
	if node.right != nil {
		node.right.postOrderDFS()
	}
	_, _ = fmt.Print(node.label, " ")
}

type TreeNode struct {
	label       int
	left, right *TreeNode
}

// -------- IGNORE BELOW HERE --------
var tokens [][]string

func parse(id int) (*TreeNode, string, int) {
	nodeId, token := id, tokens[0]
	tokens = tokens[1:]
	node, children := &TreeNode{label: -1}, []*TreeNode{nil, nil}
	child, delim, ch := 0, token[2], token[3]
	if ch == "(" {
		for ch == "(" || ch == "," {
			node, ch, id = parse(id + 1)
			children[child] = node
			child += 1
		}
		token, tokens = tokens[0], tokens[1:]
		delim, ch = token[2], token[3]
	}
	return &TreeNode{label: nodeId, left: children[0], right: children[1]}, delim, id
}

func parseTree(newick string) (tree *TreeNode) {
	if newick == "();" {
		return &TreeNode{label: 0}
	}
	r := regexp.MustCompile(`(?P<name>[^:;,()\s]*)(?P<delim>[,);])|(?P<ch>\S)`)
	tokens = r.FindAllStringSubmatch(newick, -1)
	tree, _, _ = parse(0)
	return tree
}
