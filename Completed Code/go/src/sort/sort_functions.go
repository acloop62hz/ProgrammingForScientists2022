package main

func min(inlist []int) (int, int) {

}

func SelectionSort(inlist []int) []int {
	outList := make([]int, 0)
	for i {
		idx, mval := min(inlist)
		outList = append(outList, mval)
		inlist = append(inlist[:idx], inlist[idx+1:]...)
	}
}

func InsertionSort(inlist []int) []int {

}

func BubbleSort(inlist []int) []int {
	inlist[i], inlist[i+1] = inlist[i+1], inlist[i] // assign in parallel
}
