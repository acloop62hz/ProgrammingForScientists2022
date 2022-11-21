package main

import "fmt"

// FIFO!
var queue = make([][]*Node, 0) // track the nodes we're gonna visit

func (node *Node) bfs(searchFor int) (path []*Node) {
	enqueue(node)
	for len(queue) != 0 {
		printQueue()
		curr := dequeue()
		if visited[curr] {
			print("already visited")
			continue
		}
		visit(curr)
		if curr.id == searchFor {
			print("Find it")
			return getPathFromQueue()
		} else {
			for _, child := range curr.children {
				enqueue(child)
			}
		}
	}
	return nil
}

var currPath = make([]*Node, 0)

// enqueue item onto the queue!
func enqueue(node *Node) {
	path := make([]*Node, len(currPath))
	copy(path, currPath)
	path = append(path, node)
	queue = append(queue, path)
}

// dequeue item from the queue!
func dequeue() *Node {
	currPath, queue = queue[0], queue[1:]
	node := currPath[len(currPath)-1]
	return node
}

// recovering a path for a bfs is non-trivial, don't worry about this
func getPathFromQueue() []*Node {
	queue = make([][]*Node, 0)     // reset stack
	visited = make(map[*Node]bool) // reset visited

	path := make([]*Node, len(currPath))
	copy(path, currPath)
	currPath = make([]*Node, 0)
	return path
}

func printQueue() {
	currQueue := make([]*Node, 0)
	for _, path := range queue {
		currQueue = append(currQueue, path[len(path)-1])
	}
	fmt.Printf("Queue: %v\n", currQueue)
}
