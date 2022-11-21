package main

import "fmt"

// Node a single node that composes the tree
type Node struct {
	id       int
	children []*Node
}

// String is essentially overriding the toString method.
// It tells Go how to print a Node object
func (node *Node) String() string {
	return fmt.Sprintf("%d", node.id)
}

// use this for both DFS and BFS
var visited = make(map[*Node]bool) // prevent cycles!

// visit a node!
func visit(node *Node) {
	visited[node] = true
}

// MakeGraph takes an adjacency list and turns it into an actual graph
// returns a map from id to node pointer
func MakeGraph(adjacencyList map[int][]int) map[int]*Node {
	nodes := make(map[int]*Node)
	for node, neighbors := range adjacencyList {
		// make sure parent exists
		if _, ok := nodes[node]; !ok {
			nodes[node] = &Node{id: node}
		}

		// go through and add children
		for _, neighbor := range neighbors {
			// make sure child exists
			if _, ok := nodes[neighbor]; !ok {
				nodes[neighbor] = &Node{id: neighbor}
			}
			nodes[node].children = append(nodes[node].children, nodes[neighbor])
		}
	}

	return nodes
}

func printPath(path []*Node) {
	if path == nil {
		println("could not find path!")
		return
	}

	if len(path) == 1 {
		fmt.Println(path[0].id)
	} else {
		first, path := path[0], path[1:]
		fmt.Printf("%d", first.id)
		for _, node := range path {
			fmt.Printf(" -> %d", node.id)
		}
		fmt.Println() // flush buffer
	}
	fmt.Println()
}
