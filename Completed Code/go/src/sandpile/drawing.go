package main

import "fmt"

func DrawPNG(board Board, cellWidth int, filename string) {

	width := len(board) * cellWidth
	c := CreateNewCanvas(width, width)

	// declare colors
	zero := MakeColor(0, 0, 0)
	one := MakeColor(85, 85, 85)
	two := MakeColor(170, 170, 170)
	three := MakeColor(255, 255, 255)
	//four := MakeColor(255, 0, 0)

	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j].coinNumber == 0 {
				c.SetFillColor(zero)
			} else if board[i][j].coinNumber == 1 {
				c.SetFillColor(one)
			} else if board[i][j].coinNumber == 2 {
				c.SetFillColor(two)
			} else if board[i][j].coinNumber == 3 {
				c.SetFillColor(three)
			} else {
				//c.SetFillColor(four)
				fmt.Println(board[i][j])
				//panic("Error: Out of range value " + string(board[i][j].coinNumber) + " in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	c.SaveToPNG(filename)
}
