package main

import (
	"canvas"
	"image"
)

func BoardsToImages(boards []GameBoard, cellWidth int) []image.Image {
	imageList := make([]image.Image, len(boards))
	for i := range boards {
		imageList[i] = boards[i].BoardToImage(cellWidth)
	}
	return imageList
}

//BoardToImage converts a GameBoard to an image, in which
//each cell has a cell width given by a parameter
func (g GameBoard) BoardToImage(cellWidth int) image.Image {
	rows := len(g)
	cols := len(g[0])

	c := canvas.CreateNewCanvas(cellWidth*rows, cellWidth*cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g[i][j].strategy == "C" {
				c.SetFillColor(canvas.MakeColor(0, 0, 255))
			} else if g[i][j].strategy == "D" {
				c.SetFillColor(canvas.MakeColor(255, 0, 0))
			}

			x1, y1 := cellWidth*j, cellWidth*i
			x2, y2 := cellWidth*(j+1), cellWidth*(i+1)

			c.ClearRect(x1, y1, x2, y2)

			c.Fill()

		}
	}
	return c.GetImage()

}
