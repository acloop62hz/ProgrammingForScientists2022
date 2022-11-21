package main

import (
	"fmt"
	"math/rand"
)

func InitializeBoardCentral(size, pile int) Board {
	var initialBoard Board
	initialBoard = make(Board, size)
	for i := range initialBoard {
		initialBoard[i] = make([]Square, size)
	}

	middle := size / 2
	initialBoard[middle][middle].coinNumber = pile

	for i := range initialBoard {
		for j := range initialBoard[i] {
			initialBoard[i][j].inqueue = false
		}
	}

	return initialBoard
}

func InitializeBoardRandom(size, pile int) Board {
	var initialBoard Board
	var npile int
	initialBoard = make(Board, size)
	for i := range initialBoard {
		initialBoard[i] = make([]Square, size)
	}

	npile = pile / 100

	for i := 0; i < 100; i++ {
		x := rand.Intn(size)
		y := rand.Intn(size)
		fmt.Println(x, y)
		if i < 99 {
			initialBoard[x][y].coinNumber += npile
		} else {
			initialBoard[x][y].coinNumber += pile - 99*npile
		}

	}

	for i := range initialBoard {
		for j := range initialBoard[i] {
			initialBoard[i][j].inqueue = false
		}
	}

	return initialBoard
}

func CopyBoard(currentBoard Board) Board {
	var newBoard Board
	size := len(currentBoard)
	newBoard = make(Board, size)
	for i := range currentBoard {
		newBoard[i] = make([]Square, size)
	}

	for i := range newBoard {
		for j := range newBoard[i] {
			newBoard[i][j] = currentBoard[i][j]
		}
	}

	return newBoard
}

func SandpileSerial(initialBoard Board) Board {
	var pair OrderPair

	above4Queue := GenerateAbove4Queue(initialBoard)
	for len(above4Queue) > 0 {
		pair = above4Queue.Dequeue()
		initialBoard[pair.x][pair.y].inqueue = false
		initialBoard.Topple(pair.x, pair.y, &above4Queue)
	}

	return initialBoard
}

func (board *Board) Topple(x, y int, queue *Queue) {
	size := len(*board)
	sides := make([]int, 2)

	var newPair OrderPair
	if x == 0 {
		sides[0] = 1
	} else {
		(*board)[x-1][y].coinNumber += 1
		if (*board)[x-1][y].coinNumber >= 4 && (*board)[x-1][y].inqueue == false {
			//fmt.Println("!")
			newPair.x = x - 1
			newPair.y = y
			queue.Enqueue(newPair)
			(*board)[x-1][y].inqueue = true
		}
	}

	if x == size-1 {
		sides[1] = 1
	} else {
		(*board)[x+1][y].coinNumber += 1
		if (*board)[x+1][y].coinNumber >= 4 && (*board)[x+1][y].inqueue == false {
			newPair.x = x + 1
			newPair.y = y
			queue.Enqueue(newPair)
			(*board)[x+1][y].inqueue = true
		}
	}

	if y == 0 {
	} else {
		(*board)[x][y-1].coinNumber += 1
		if (*board)[x][y-1].coinNumber >= 4 && (*board)[x][y-1].inqueue == false {
			newPair.x = x
			newPair.y = y - 1
			queue.Enqueue(newPair)
			(*board)[x][y-1].inqueue = true
		}
	}

	if y == size-1 {
	} else {
		(*board)[x][y+1].coinNumber += 1
		if (*board)[x][y+1].coinNumber >= 4 && (*board)[x][y+1].inqueue == false {
			newPair.x = x
			newPair.y = y + 1
			//fmt.Println("!")
			queue.Enqueue(newPair)
			(*board)[x][y+1].inqueue = true
		}
	}

	(*board)[x][y].coinNumber -= 4
	if (*board)[x][y].coinNumber >= 4 {
		newPair.x = x
		newPair.y = y
		queue.Enqueue(newPair)
		(*board)[x][y].inqueue = true
	}
}

func (Q *Queue) Enqueue(item OrderPair) {
	*Q = append(*Q, item)
}

func (Q *Queue) Dequeue() OrderPair {

	if len(*Q) < 1 {
		panic("EMPTY QUEUE, CANNOT DEQUEUE")
	}

	retval := (*Q)[0]

	*Q = (*Q)[1:len(*Q)]

	return retval

}

func GenerateAbove4Queue(initialBoard Board) Queue {
	queue := make(Queue, 0)
	var newPair OrderPair

	for i := range initialBoard {
		for j := range initialBoard[i] {
			if initialBoard[i][j].coinNumber >= 4 {
				newPair.x = i
				newPair.y = j
				queue.Enqueue(newPair)
				initialBoard[i][j].inqueue = true
			}
		}
	}
	return queue
}
