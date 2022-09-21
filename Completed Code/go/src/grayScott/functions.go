package main

//place your functions from the assignment here.

func SimulateGrayScott(initialBoard Board, numGens int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) []Board {
	boards := make([]Board, numGens+1)
	boards[0] = initialBoard
	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1], feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
	}
	return boards
}

func UpdateBoard(currentBoard Board, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Board {
	//no need to initailize a new board --> cause out of memory error
	for i := range currentBoard {
		for j := range currentBoard[i] {
			currentBoard[i][j] = UpdateCell(currentBoard, i, j, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
		}
	}
	return currentBoard
}

func UpdateCell(currentBoard Board, row, col int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {
	reaction_change := ChangeDueToReactions(currentBoard[row][col], feedRate, killRate)
	diffusion_change := ChangeDueToDiffusion(currentBoard, row, col, preyDiffusionRate, predatorDiffusionRate, kernel)
	return SumCells(currentBoard[row][col], reaction_change, diffusion_change)
}

func ChangeDueToReactions(currentCell Cell, feedRate, killRate float64) Cell {
	reaction_change := [2]float64{0.0, 0.0}
	reaction_change[0] = feedRate*(1-currentCell[0]) - currentCell[0]*currentCell[1]*currentCell[1]
	reaction_change[1] = -killRate*currentCell[1] + currentCell[0]*currentCell[1]*currentCell[1]
	return reaction_change
}

func ChangeDueToDiffusion(currentBoard Board, row, col int, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {
	add := Cell{Convolution(row, col, currentBoard, kernel)[0] * preyDiffusionRate, Convolution(row, col, currentBoard, kernel)[1] * predatorDiffusionRate}
	return add
}

func SumCells(cells ...Cell) Cell {
	var cellSum Cell
	cellSum[0] = 0.0
	cellSum[1] = 0.0
	for _, val := range cells {
		cellSum[0] += val[0]
		cellSum[1] += val[1]
	}
	return cellSum
}

func InitializeBoard(rows, cols int) Board {
	mtx := make(Board, rows)
	for i := range mtx {
		mtx[i] = make([]Cell, cols)
	}
	return mtx
}

func InField(mtx Board, row, col int) bool {
	if row < 0 || col < 0 {
		return false
	} else if row >= CountRows(mtx) || col >= CountCols(mtx) {
		return false
	} else {
		return true
	}
}

func CountRows(mtx Board) int {
	return len(mtx)
}

func CountCols(mtx Board) int {
	if CountRows(mtx) == 0 {
		panic("Error:empty board")
	}
	return len(mtx[0])
}

func Convolution(row, col int, currentBoard Board, kernel [3][3]float64) Cell {
	sum := Cell{0.0, 0.0}
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if InField(currentBoard, i, j) {
				add := Cell{currentBoard[i][j][0] * kernel[i+1-row][j+1-col], currentBoard[i][j][1] * kernel[i+1-row][j+1-col]}
				sum = SumCells(sum, add)
			}
		}
	}
	return sum
}

// This function is deleted to avoid repeated loops
// func GetNeighborBoard(irow, jcol int, currentboard Board) Board {
// 	neighbors := InitializeBoard(3, 3)
// 	for i := irow - 1; i <= irow+1; i++ {
// 		for j := jcol - 1; j <= jcol+1; j++ {
// 			if InField(currentboard, i, j) {
// 				neighbors[i+1-irow][j+1-jcol] = currentboard[i][j]
// 			} else {
// 				neighbors[i+1-irow][j+1-jcol] = Cell{0.0, 0.0}
// 			}
// 		}
// 	}
// 	return neighbors
// }
