package main

type Square struct {
	coinNumber int
	inqueue    bool
}

type Board [][]Square

type OrderPair struct {
	x int
	y int
}

type Queue []OrderPair

type message struct {
	bottomFall  []int
	upperFall   []int
	bottomIndex int
	upperIndex  int
	done        bool
}
