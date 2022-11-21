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
