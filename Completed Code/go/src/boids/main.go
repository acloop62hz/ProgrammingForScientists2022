package main

import (
	"fmt"
	"gifhelper"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	//os.Args[1] is going to be the number of boids in a sky
	numBoids, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	if numBoids < 0 {
		panic("Negative number of generations given.")
	}

	//os.Args[2] is going to be time step parameter
	skyWidth, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		panic(err2)
	}

	//os.Args[3] is the value for initial speed for all boids
	initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3 != nil {
		panic(err3)
	}

	//os.Args[4] is the upper bound for boids' speed
	maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
	if err4 != nil {
		panic(err4)
	}

	//os.Args[5] is the number of generations
	numGens, err5 := strconv.Atoi(os.Args[5])
	if err5 != nil {
		panic(err5)
	}

	//os.Args[6] is the largest distance for boids to interact with each other
	proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6 != nil {
		panic(err6)
	}

	//os.Args[7] is  the coefficient for calculating the separation force
	separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7 != nil {
		panic(err7)
	}

	//os.Args[8] is the coefficient for calculating the alignment force
	alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8 != nil {
		panic(err8)
	}

	//os.Args[9] is the coefficient for calculating the cohesion force
	cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
	if err9 != nil {
		panic(err9)
	}

	//os.Args[10] is the time step for each generation
	timeStep, err10 := strconv.ParseFloat(os.Args[10], 64)
	if err10 != nil {
		panic(err10)
	}

	//os.Args[11] is the canvas width
	canvasWidth, err11 := strconv.Atoi(os.Args[11])
	if err11 != nil {
		panic(err11)
	}

	//os.Args[12] is how often to make a canvas
	imageFrequency, err12 := strconv.Atoi(os.Args[12])
	if err12 != nil {
		panic(err12)
	}

	//create a initial Sky
	var initialSky Sky
	initialSky.width = skyWidth
	initialSky.maxBoidSpeed = maxBoidSpeed
	initialSky.proximity = proximity
	initialSky.separationFactor = separationFactor
	initialSky.alignmentFactor = alignmentFactor
	initialSky.cohesionFactor = cohesionFactor

	//generate boids with random positions and random velocity directions
	boids := make([]Boid, numBoids)
	for i := 0; i < numBoids; i++ {
		boids[i].position.x = 0 + rand.Float64()*skyWidth
		boids[i].position.y = 0 + rand.Float64()*skyWidth
		boids[i].velocity.x = initialSpeed * math.Cos(0+rand.Float64()*2*math.Pi)
		boids[i].velocity.y = initialSpeed * math.Sin(0+rand.Float64()*2*math.Pi)
	}
	initialSky.boids = boids

	// Simulate the system to get skies at different timepoints
	fmt.Println("Command line arguments read successfully.")

	fmt.Println("Simulating system.")

	timePoints := SimulateBoids(initialSky, numGens, timeStep)

	fmt.Println("Boids has been simulated!")
	fmt.Println("Ready to draw images.")

	// Get images files for all timepoints

	images := AnimateSystem(timePoints, canvasWidth, imageFrequency)
	fmt.Println("Images drawn!")

	fmt.Println("Making GIF.")

	// put images together to generate a gif

	gifhelper.ImagesToGIF(images, "Boids.2")

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")

}
