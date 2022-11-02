package main

import (
	"canvas"
	"image"
)

func AnimateSystem(timePoints []lane, canvasWidth, drawingFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%drawingFrequency == 0 { //only draw if current index of universe
			//is divisible by some parameter frequency
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

func DrawToCanvas(u lane, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a white background
	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the boids and draw them.
	for _, b := range u.vehicles {
		c.SetFillColor(canvas.MakeColor(0, 0, 0))
		cx := (float64(b.position) / float64(u.length)) * float64(canvasWidth)
		cy := canvasWidth / 2
		r := 2.0
		c.Circle(float64(cx), float64(cy), r)
		c.Fill()
	}
	// return an image
	return c.GetImage()
}
