package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//ReadBSTFromText takes in a text file, reads the unsorted elements in the file, and returns a BST.
func ReadBSTFromText(fileName string) Tree {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: couldn't open the file")
		os.Exit(1)
	}

	var lines []int = make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		element, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, element)
	}
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	file.Close()

	//initialize tree
	var t Tree
	numLeaves := len(lines)
	t = make([]*Node, numLeaves)

	for index, element := range lines {
		var vx Node

		//assign first element as root node
		if index == 0 {
			vx.label = element
		} else {
			vx.label = element
			//find spot in bst to place node
			found := false
			node := t[0]
			//traverse tree until an empty spot is found
			for found == false {
				if vx.label < node.label {
					//if there is an open spot, add node
					if node.child1 == nil {
						node.child1 = &vx
						found = true
					} else {
						//continue to the left
						node = node.child1
					}
				}
				if vx.label > node.label {
					//if there is an open spot, add node
					if node.child2 == nil {
						node.child2 = &vx
						found = true
					} else {
						//continue to the right
						node = node.child2
					}
				}
			}
		}
		//point the pointer at the node
		t[index] = &vx
	}
	return t
}
