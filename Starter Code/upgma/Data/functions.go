package main

type DistanceMatrix[][]

type Tree *Node

func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree{

}

func CountLeaves(v *Node) int{
	if v.Child1 == nil && Child2 == nil{
		return 1
	}
	if v.Child1 == nil{
		return CountLeaves(v.Child2)
	}
	if v.Child2 == nil{
		return CountLeaves(v.Child1)
	}
	else{
		return CountLeaves(v.Child1)+CountLeaves(v.Child2)
	}
}

Delete(){
	delete the larger one first
}