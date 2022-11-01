package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
)

// one lane with 1000 cells
// vi <- min{vi(t-1)+1, gsi(t-1), vmax}

func main() {

	var initiallane lane

	initiallane.length = 1000

	vehicles := make([]vehicle, 10)
	for i := 0; i < 10; i++ {
		vehicles[i].velocity = 1
		vehicles[i].position = rand.Intn(1000)
	}

	initiallane.vehicles = vehicles

	fmt.Println("Command line arguments read successfully.")
	fmt.Println("Simulating system.")

	timeStep := 1
	numGens := 20
	timePoints := SimulateLanes(initiallane, numGens, timeStep)

	fmt.Println("Boids has been simulated!")
	fmt.Println("Ready to draw images.")

	canvasWidth := 20
	imageFrequency := 1

	images := AnimateSystem(timePoints, canvasWidth, imageFrequency)

	fmt.Println("Images drawn!")
	fmt.Println("Making GIF.")

	gifhelper.ImagesToGIF(images, "Boids.2")

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")

}
