package main

import (
	"canvas"
	"image"
)

//AnimateSystem takes a slice of Sky objects along with a canvas width parameter and a frequency parameter.
//It generates a slice of images corresponding to drawing each Sky whose index is divisible by the frequency parameter on a canvasWidth x canvasWidth canvas
func AnimateSystem(timePoints []Sky, canvasWidth, drawingFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%drawingFrequency == 0 { //only draw if current index of universe
			//is divisible by some parameter frequency
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Sky object's boids at one time point
//on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(u Sky, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a white background
	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the boids and draw them.
	for _, b := range u.boids {
		c.SetFillColor(canvas.MakeColor(0, 0, 0))
		cx := (b.position.x / u.width) * float64(canvasWidth)
		cy := (b.position.y / u.width) * float64(canvasWidth)
		r := 5.0
		c.Circle(cx, cy, r)
		c.Fill()
	}
	// return an image
	return c.GetImage()
}
