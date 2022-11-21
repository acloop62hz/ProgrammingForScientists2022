package main

// LIFO!
var stack = make([]*Node, 0) // track our path

// Question: why can't we just use the stack to keep track of where we've visited?

func (node *Node) dfs(searchFor int) (path []*Node) {
	if visited[node] {
		print("already visited")
		return nil
	}
	push(node)
	visit(node)
	if node.id == searchFor {
		print("Find")
	} else {
		for _, child := range node.children {
			if p := child.dfs(searchFor); p != nil {
				return p
			}
		}
	}
	pop()
	return nil
}

// push onto the stack!
func push(node *Node) {
	stack = append(stack, node)
}

// pop off the stack!
func pop() {
	stack = stack[:len(stack)-1]
}

// getPath is a helper to clean up our global variables and create our return variable
func getPathFromStack() []*Node {
	var path = make([]*Node, len(stack))
	copy(path, stack)              // store solution
	stack = make([]*Node, 0)       // reset stack
	visited = make(map[*Node]bool) // reset visited
	return path
}
