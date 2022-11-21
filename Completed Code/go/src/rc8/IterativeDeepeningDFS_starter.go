package main

import "fmt"

// iterativeDeepeningDFS is sort of like a combo of DFS with BFS
func (node *Node) iterativeDeepeningDFS(searchFor int, depth int) (path []*Node) {
	enqueue(node)
	for len(queue) != 0 {
		printQueue()
		curr := dequeue()
		fmt.Printf("Current head is %d. ", curr.id)
		if visited[curr] {
			fmt.Printf("Already visited %d!\n", curr.id)
			continue // next iteration!
		}

		//TODO: Add limited search

		visit(curr) // don't come here again

		if path == nil {
			fmt.Printf("Children to enqueue are %v\n", curr.children)
			for _, child := range curr.children {
				enqueue(child)
			}

		} else {
			return path
		}
	}
	return nil
}

// subroutine
func (node *Node) depthLimitedSearch(searchFor int, depth int) (path []*Node) {
	fmt.Printf("called DFS limited on %d. ", node.id)

	if visited[node] {
		fmt.Println("already visited!")
		return nil
	}

	push(node)  // add current node to our path
	visit(node) // don't come here again

	fmt.Printf("visiting %d, with kids %v\n", node.id, node.children)

	// check if we have found our node
	if node.id == searchFor {
		return getPathFromStack()

		//TODO: check depth
	} else {
		// see if our children have a path
		for _, child := range node.children {

			//TODO: call limited search
			if p := nil; p != nil {
				return p // return the path if so
			}
		}
	}

	pop() // didn't find anything, remove this node from our path
	return nil
}
