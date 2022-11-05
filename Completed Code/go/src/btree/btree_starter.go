package main

import "fmt"

type Tree []*Node

// we also think of a cluster as a *Node

type Node struct {
	label          int
	child1, child2 *Node
}

func main() {
	bst := ReadBSTFromText("bst.txt")

	//print tree
	for _, node := range bst {
		fmt.Println("label", node.label, "child1", node.child1, "child2", node.child2)
	}

	//Counting leaves
	//fmt.Println("Number of Leaves", bst[0].CountLeaves())

	//Searching recursive
	//fmt.Println(SearchRecursive(1, bst[0]))

	//Searching iteratively
	//fmt.Println(SearchIterative(7, bst[0]))

	//Inserting a new node
	var empty *Node
	bst = append(bst, InsertRecursive(5, bst[0], empty))
	for _, node := range bst {
		fmt.Println("label", node.label, "child1", node.child1, "child2", node.child2)
	}

	//Traversals
	// InOrderTraversal(bst[0])
	// PreOrderTraversal(bst[0])
	// PostOrderTraversal(bst[0])

}

//CountLeaves is a Node method that counts the number of leaves in the tree rooted at the node. It returns 1 at a leaf.
func (vx *Node) CountLeaves() int {
	if vx.child1 == nil || vx.child2 == nil {
		if vx.child1 == nil && vx.child2 == nil {
			return 1 // could have weird case of only one child being nil but the tree we produce won't have this.
		} else if vx.child1 == nil {
			return vx.child2.CountLeaves()
		} else {
			vx.child1.CountLeaves()
		}
	}
	// if we make it here, we have two non-nil kids
	return vx.child1.CountLeaves() + vx.child2.CountLeaves()
}

//SearchRecursive takes in a key and searches through a binary search tree node by node to find the key.
//Returns the node if found. Else, returns nil.
func SearchRecursive(key int, node *Node) *Node {
	if node == nil || node.label == key {
		return node
	} else if node.label > key {
		return SearchRecursive(key, node.child1)
	} else {
		return SearchRecursive(key, node.child2)
	}
	// key > node.label
}

//SearchIterative is an iterative version of SearchRecursive
func SearchIterative(key int, node *Node) *Node {
	for {
		if node == nil || node.label == key {
			return node
		}

		if key < node.label {
			node = node.child1
		} else {
			node = node.child2
		}
	}
}

//InsertRecursive takes in the key you want to add the tree, the current node you are looking at, and the parent node
//of the current node. It will find the appropriate spot to place the key, make a new node for the key, and point the new
//node to a child of the parent node.
//It returns the pointer node that was created so it can be added to the tree.
func InsertRecursive(key int, currNode, prevNode *Node) *Node {
	//if node is empty, place the key at that node
	if currNode == nil {
		var vx Node
		vx.label = key
		if key < prevNode.label {
			prevNode.child1 = &vx
		} else if key > prevNode.label {
			prevNode.child2 = &vx
		}
		return &vx
	}

	//otherwise, head down tree and find a suitable spot
	if key < currNode.label {
		return InsertRecursive(key, currNode.child1, currNode)
	} else {
		return InsertRecursive(key, currNode.child2, currNode)
	}
	//if node is already in tree, just return the node
}

//InOrderTraversal takes a BST and prints out the elements of the tree in increasing order
func InOrderTraversal(currNode *Node) {
	if currNode != nil {
		InOrderTraversal(currNode.child1)
		fmt.Print(currNode.label, " ")
		InOrderTraversal(currNode.child2)
	}

}

//PreOrderTraversal takes a BST and prints out current node, then left child, then right child
func PreOrderTraversal(currNode *Node) {
	if currNode != nil {
		fmt.Print(currNode.label, " ")
		PreOrderTraversal(currNode.child1)
		PostOrderTraversal(currNode.child2)

	}

}

//PostOrderTraversal takes a BST and prints out left child, then right child, then current node
func PostOrderTraversal(currNode *Node) {
	if currNode != nil {
		PreOrderTraversal(currNode.child1)
		PostOrderTraversal(currNode.child2)
		fmt.Print(currNode.label, " ")

	}

}
