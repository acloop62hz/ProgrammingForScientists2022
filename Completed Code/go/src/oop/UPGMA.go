package main

import "fmt"

func main() {
	a := make([]int, 0, 20)
	//a = append(a, 1)
	fmt.Println(len(a))
}

type Node struct {
	name     rune
	Distance int
	Children []Node
	Parent   *Node
}

type Matrix struct {
	Nodes  []*Node
	Matrix [][]int
}
