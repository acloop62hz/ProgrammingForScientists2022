package main

import (
	"reflect"
	"testing"
)

func TestCountLeaves(t *testing.T) {
	bst := ReadBSTFromText("bst.txt")

	leaves := bst[0].CountLeaves()

	if leaves != 3 {
		t.Errorf("Expected %d but got %d", 3, leaves)
	}
}

func TestSearchRecursive(t *testing.T) {
	bst := ReadBSTFromText("bst.txt")

	if SearchRecursive(5, bst[0]) != nil {
		t.Error("Expected", nil, "but got", SearchRecursive(5, bst[0]))
	}

	var tester Node
	tester.label = 1

	tester_pointer := &tester

	if SearchRecursive(1, bst[0]).label != tester_pointer.label {
		t.Error("Expected", 1, "but got", SearchRecursive(1, bst[0]))
	}
}

func TestSearchIterative(t *testing.T) {
	bst := ReadBSTFromText("bst.txt")

	if SearchIterative(5, bst[0]) != nil {
		t.Error("Expected", nil, "but got", SearchIterative(5, bst[0]))
	}

	var tester Node
	tester.label = 1

	tester_pointer := &tester

	if SearchIterative(1, bst[0]).label != tester_pointer.label {
		t.Error("Expected", 1, "but got", SearchIterative(1, bst[0]))
	}
}

func TestInsertRecursive(t *testing.T) {
	bst := ReadBSTFromText("bst.txt")

	var correct_bst Tree

	correct_bst = make([]*Node, len(bst))

	for index, node := range bst {
		var vx Node
		vx.label = node.label
		vx.child1 = node.child1
		vx.child2 = node.child2
		correct_bst[index] = &vx
	}

	for _, node := range correct_bst {
		if node.label == 6 {
			var vx Node
			vx.label = 5
			node.child1 = &vx
			correct_bst = append(correct_bst, &vx)
		}
	}

	var empty *Node

	bst = append(bst, InsertRecursive(5, bst[0], empty))

	if !reflect.DeepEqual(bst, correct_bst) {
		t.Error("Should have gotten", correct_bst, "but got", bst)
	}
}
