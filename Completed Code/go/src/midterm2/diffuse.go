package main

type Board [][]float64

func DiffuseMultiProc(initialBoard Board, convolution [3][3]float64, numSteps int, numProcs int) []Board {
	boards := make([]Board, numSteps+1)
	boards[0] = initialBoard

	for i := 1; i < numSteps; i++ {
		boards[i] = UpdateBoardParallel(boards[i-1], convolution, numProcs)
	}
	return boards
}

func UpdateBoardParallel(currentBoard Board, kernel [3][3]float64, numProcs int) Board {
	var start int
	var end int
	c := make(chan []Board, numProcs)
	n := len(currentBoard)

	for j := 0; j < numProcs; j++ {
		if j < numProcs-1 {
			start = j * (n / numProcs)
			end = (j + 1) * (n / numProcs)
		} else {
			start = j * (n / numProcs)
			end = n - 1
		}
		go DiffuseBoardOneParticleTorus(currentBoard[start:end+1], kernel, c)
	}
}

func DiffuseBoardOneParticleTorus(currentBoard Board, kernel [3][3]float64, c chan []Board) {
	newboard := InitializeBoard(CountRows(currentBoard), CountCols(currentBoard))
	for i := range currentBoard {
		for j := range currentBoard[i] {
			neighbors := GetNeighborBoard(i, j, currentBoard)
			newboard[i][j] = currentBoard[i][j] + Convolution(neighbors, kernel)
		}
	}
	c <- newboard
}

func InitializeBoard(rows, cols int) Board {
	mtx := make(Board, rows)
	for i := range mtx {
		mtx[i] = make([]float64, cols)
	}
	return mtx
}

func TorusIndex(mtx Board, row, col int) []int {
	rc := make([]int, 2)
	if row < 0 {
		rc[0] = CountRows(mtx) + row
	} else if row >= CountRows(mtx) {
		rc[0] = row - CountRows(mtx)
	} else {
		rc[0] = row
	}
	if col < 0 {
		rc[1] = CountRows(mtx) + col
	} else if col >= CountRows(mtx) {
		rc[1] = col - CountRows(mtx)
	} else {
		rc[1] = col
	}
	return rc
}

func GetNeighborBoard(irow, jcol int, currentboard Board) Board {
	neighbors := InitializeBoard(3, 3)
	for i := irow - 1; i <= irow+1; i++ {
		for j := jcol - 1; j <= jcol+1; j++ {
			newidex := TorusIndex(currentboard, i, j)
			neighbors[i+1-irow][j+1-jcol] = currentboard[newidex[0]][newidex[1]]
		}
	}
	return neighbors
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

func Convolution(neighbors Board, kernel [3][3]float64) float64 {
	sum := 0.0
	for i := range neighbors {
		for j := range neighbors[i] {
			sum += neighbors[i][j] * kernel[i][j]
		}
	}
	return sum
}
