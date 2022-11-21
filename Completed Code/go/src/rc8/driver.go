package main

func main() {
	adjList := make(map[int][]int)
	adjList[0] = []int{1, 2, 5}
	adjList[1] = []int{0, 2}
	adjList[2] = []int{3, 4}
	adjList[4] = []int{3, 5}
	adjList[5] = []int{3, 6}
	graph := MakeGraph(adjList)

	println("---- DFS ----")
	printPath(graph[0].dfs(6))
	println("---- BFS ----")
	printPath(graph[0].bfs(6))
	println("---- Iterative Deepening DFS ----")
	printPath(graph[0].iterativeDeepeningDFS(6, 2))

	tree := parseTree("((a,(b,c)),(d,e));")
	/*
		      /-2
		   /-1
		  |  |   /-4
		  |   \-3
		--0      \-5
		  |
		  |   /-7
		   \-6
		      \-8
	*/
	println("---- TRAVERSALS ----")
	println("---- PRE ----")
	tree.preOrderDFS()
	println()
	println("---- IN ----")
	tree.inOrderDFS()
	println()
	println("---- POST ----")
	tree.postOrderDFS()
	println()
}
