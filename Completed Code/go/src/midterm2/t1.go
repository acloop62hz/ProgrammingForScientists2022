package main

import "fmt"

//Tree is an object that contains a pointer to a Node
type Tree struct {
	root *Node
}

//Node contains a label and two pointers to children nodes
//(one or both may be nil)
type Node struct {
	label       string
	left, right *Node
}

//DO NOT EDIT ANYTHING ABOVE THIS LINE (except for adding package declarations)

//Insert your PathToNode() function here, along with any subroutines that you need.
func (t Tree) PathToNode(vx *Node) []*Node {
	p := t.root
	path := traversal(p, vx)
	return path

}

func traversal(p, vx *Node) []*Node {
	path := make([]*Node, 0)
	pathl := make([]*Node, 0)
	pathr := make([]*Node, 0)
	if p.label == vx.label {
		path = make([]*Node, 1)
		path[0] = p
		return path
	}
	if p.left == nil && p.right == nil {
	} else {
		if p.left != nil {
			pathl = traversal(p.left, vx)
		}
		if p.right != nil {
			pathr = traversal(p.right, vx)
		}
	}
	if len(pathl) > 0 {
		path = make([]*Node, 1)
		path[0] = p
		path = append(path, pathl...)
		fmt.Println(len(path))
	} else if len(pathr) > 0 {
		path = make([]*Node, 1)
		path[0] = p
		path = append(path, pathl...)
		fmt.Println(len(path))
	} else {
	}
	return path
}
