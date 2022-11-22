package main

import (
	"math/rand"
)

// InitializeBoardCentral initialize a board with given size and pile (int object)
// in a way all the coins are in the central square
func InitializeBoardCentral(size, pile int) Board {
	var initialBoard Board

	initialBoard = make(Board, size)
	for i := range initialBoard {
		initialBoard[i] = make([]Square, size)
	}

	// find the middle square of the board and place the pile
	middle := size / 2
	initialBoard[middle][middle].coinNumber = pile

	// initialize the "inqueue" field of a square, "false" means that the square in not currently put into a queue waiting for topple
	for i := range initialBoard {
		for j := range initialBoard[i] {
			initialBoard[i][j].inqueue = false
		}
	}

	return initialBoard
}

// InitializeBoardRandom initialize a board with given size and pile (int object)
// in a way all the coins are nearly equally placed in randomly-chosen 100 squares
func InitializeBoardRandom(size, pile int) Board {
	var initialBoard Board
	var npile int

	initialBoard = make(Board, size)
	for i := range initialBoard {
		initialBoard[i] = make([]Square, size)
	}

	// npile is the approximately average number of coins placed on each square
	npile = pile / 100

	for i := 0; i < 100; i++ {
		// randomly pick squares
		x := rand.Intn(size)
		y := rand.Intn(size)
		if i < 99 {
			initialBoard[x][y].coinNumber += npile
		} else {
			// the last square takes the remaining coins
			initialBoard[x][y].coinNumber += pile - 99*npile
		}
	}

	// initialize the "inqueue" field of a square, "false" means that the square in not currently put into a queue waiting for topple
	for i := range initialBoard {
		for j := range initialBoard[i] {
			initialBoard[i][j].inqueue = false
		}
	}

	return initialBoard
}

// CopyBoard makes a deep copy of a board and return it
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

// SandpileSerial topple the coins serially
func SandpileSerial(initialBoard *Board) Board {
	var pair OrderPair

	// get all squares with more than 4 coins into a queue
	above4Queue := GenerateAbove4Queue(initialBoard)
	//topple until the queue is empty, which means no square has more than 4 coins
	for len(above4Queue) > 0 {
		pair = above4Queue[0]
		initialBoard.Topple(pair.x, pair.y, &above4Queue)
	}
	return *initialBoard
}

// Topple distribution the coins of (x,y) square to its neighbors
func (board *Board) Topple(x, y int, queue *Queue) []int {
	length := len(*board)
	size := len((*board)[0])
	// track the coins that fall off the edge
	sides := make([]int, 2)

	var newPair OrderPair
	// track the coins that fall off at bottom
	if x == 0 {
		sides[0] = 1
	} else {
		(*board)[x-1][y].coinNumber += 1
		// enqueue the square with >= 4 coins if the square is not in the queue
		if (*board)[x-1][y].coinNumber >= 4 && (*board)[x-1][y].inqueue == false {
			newPair.x = x - 1
			newPair.y = y
			queue.Enqueue(newPair)
			(*board)[x-1][y].inqueue = true
		}
	}

	// track the coins that fall off at top
	if x == length-1 {
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

	// track the coins that fall at left
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

	// track the coins that fall at right
	if y == size-1 {
	} else {
		(*board)[x][y+1].coinNumber += 1
		if (*board)[x][y+1].coinNumber >= 4 && (*board)[x][y+1].inqueue == false {
			newPair.x = x
			newPair.y = y + 1
			queue.Enqueue(newPair)
			(*board)[x][y+1].inqueue = true
		}
	}

	// remove 4 coins from the given square.
	(*board)[x][y].coinNumber -= 4
	// if the current square has less than 4 coins, remove it from the queue
	if (*board)[x][y].coinNumber < 4 {
		queue.Dequeue()
		(*board)[x][y].inqueue = false
	}
	return sides
}

// Enqueue put a new item into the queue
func (Q *Queue) Enqueue(item OrderPair) {
	*Q = append(*Q, item)
}

// Dequeue remove the first element from queue
func (Q *Queue) Dequeue() OrderPair {

	if len(*Q) < 1 {
		panic("EMPTY QUEUE, CANNOT DEQUEUE")
	}
	retval := (*Q)[0]
	*Q = (*Q)[1:len(*Q)]

	return retval

}

// GenerateAbove4Queue generate a queue containing squares with more or equal than 4 coins based on the given board
func GenerateAbove4Queue(initialBoard *Board) Queue {
	queue := make(Queue, 0)
	var newPair OrderPair

	for i := range *initialBoard {
		for j := range (*initialBoard)[i] {
			if (*initialBoard)[i][j].coinNumber >= 4 {
				newPair.x = i
				newPair.y = j
				queue.Enqueue(newPair)
				(*initialBoard)[i][j].inqueue = true
			}
		}
	}
	return queue
}

// SandpileMultiprocs runs topple in parallel
func SandpileMultiprocs(initialBoard *Board, numProcs int) Board {
	var start int
	var end int
	var finishCount int
	size := len(*initialBoard)
	var sliceMessage message

	// make a buffer channel
	messageChan := make(chan message, numProcs)

	// count the number of subslices that have been stable
	finishCount = 0

	// finish running topple when all subslices are stable
	for finishCount != numProcs {

		for i := 0; i < numProcs; i++ {
			// divide the board into numProcs parts
			start = i * size / numProcs
			end = (i + 1) * size / numProcs
			// run subroutines and put the running outcome into the buffer channel
			if i < numProcs-1 {
				subslice := (*initialBoard)[start:end]
				go SandpileSingleproc(&(subslice), messageChan, start, end-1)
			} else {
				subslice := (*initialBoard)[start:]
				go SandpileSingleproc(&(subslice), messageChan, start, size-1)
			}
		}

		finishCount = 0

		// get subroutine outcome from the buffer channel
		for i := 0; i < numProcs; i++ {
			sliceMessage = <-messageChan
			if sliceMessage.done {
				// count the parts that has finished
				finishCount = finishCount + 1
			} else {
				// if the subslice has not finished yet, add the coins fallen from edges to its neighbors
				if sliceMessage.upperIndex < size {
					for k := 0; k < size; k++ {
						(*initialBoard)[sliceMessage.upperIndex][k].coinNumber += sliceMessage.upperFall[k]
					}
				}
				if sliceMessage.bottomIndex >= 0 {
					for k := 0; k < size; k++ {
						(*initialBoard)[sliceMessage.bottomIndex][k].coinNumber += sliceMessage.bottomFall[k]
					}
				}
			}
		}

	}
	return *initialBoard

}

// SandpileSingleprocs runs topple for each subslice
func SandpileSingleproc(initialBoard *Board, c chan message, start, end int) {
	var pair OrderPair
	var chanMessage message
	size := len((*initialBoard)[0])
	length := len(*initialBoard)
	// track the coins that fall from the bottom
	chanMessage.bottomFall = make([]int, size)
	// track the coins that fall from the top
	chanMessage.upperFall = make([]int, size)
	// the bottomIndex indicate the row in the whole board to add the fallen coins from the bottom
	chanMessage.bottomIndex = start - 1
	// the upperIndex indicate the row in the whole board to add the fallen coins from the top
	chanMessage.upperIndex = end + 1
	// track if the subslice is already stable
	chanMessage.done = false

	// generate a queue for topple
	above4Queue := GenerateAbove4Queue(initialBoard)
	if len(above4Queue) == 0 {
		chanMessage.done = true
	}
	for len(above4Queue) > 0 {
		pair = above4Queue[0]
		add := initialBoard.Topple(pair.x, pair.y, &above4Queue)
		// count the coins fall from the bottom
		if pair.x == 0 {
			chanMessage.bottomFall[pair.y] += add[0]
		}
		// count the coins fall from the top
		if pair.x == length-1 {
			chanMessage.upperFall[pair.y] += add[1]
		}
	}

	// store the running outcome for each sublice into the buffer channel
	c <- chanMessage
}
