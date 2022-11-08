package main

import (
	"fmt"
	"gifhelper"
	"math"
	"os"
)

func main() {

	// a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	universeType := os.Args[1]
	if universeType == "jupiter" {
		var jupiter, io, europa, ganymede, callisto Star

		jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
		io.red, io.green, io.blue = 249, 249, 165
		europa.red, europa.green, europa.blue = 132, 83, 52
		ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
		callisto.red, callisto.green, callisto.blue = 0, 153, 76

		jupiter.mass = 1.898 * math.Pow(10, 27)
		io.mass = 8.9319 * math.Pow(10, 22)
		europa.mass = 4.7998 * math.Pow(10, 22)
		ganymede.mass = 1.4819 * math.Pow(10, 23)
		callisto.mass = 1.0759 * math.Pow(10, 23)

		jupiter.radius = 71000000
		io.radius = 1821000
		europa.radius = 1569000
		ganymede.radius = 2631000
		callisto.radius = 2410000

		jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
		io.position.x, io.position.y = 2000000000-421600000, 2000000000
		europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
		ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
		callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

		jupiter.velocity.x, jupiter.velocity.y = 0, 0
		io.velocity.x, io.velocity.y = 0, -17320
		europa.velocity.x, europa.velocity.y = -13740, 0
		ganymede.velocity.x, ganymede.velocity.y = 0, 10870
		callisto.velocity.x, callisto.velocity.y = 8200, 0

		// declaring universe and setting its fields.
		var jupiterSystem Universe
		jupiterSystem.width = 4000000000
		jupiterSystem.stars = []*Star{&jupiter, &io, &europa, &ganymede, &callisto}

		// now evolve the universe: feel free to adjust the following parameters.
		numGens := 10000
		time := 50.0
		theta := 0.5
		canvasWidth := 1000
		frequency := 300
		scalingFactor := 1e1

		fmt.Println("checkPoint1")

		timePoints := BarnesHut(&jupiterSystem, numGens, time, theta)

		fmt.Println("Simulation run. Now drawing images.")

		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "jupiter")
		fmt.Println("GIF drawn.")

	} else if universeType == "galaxy" {
		g0 := InitializeGalaxy(50, 4e21, 7e22, 2e22)
		width := 1.0e23
		galaxies := []Galaxy{g0}

		initialUniverse := InitializeUniverse(galaxies, width)

		// now evolve the universe: feel free to adjust the following parameters.
		numGens := 10000
		time := 2e15
		theta := 0.5

		timePoints := BarnesHut(initialUniverse, numGens, time, theta)

		fmt.Println("Simulation run. Now drawing images.")
		canvasWidth := 1000
		frequency := 300
		scalingFactor := 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "galaxy")
		fmt.Println("GIF drawn.")

	} else if universeType == "collision" {
		g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
		g1 := InitializeGalaxy(500, 4e21, 3e22, 7e22)

		// you probably want to apply a "push" function at this point to these galaxies to move
		// them toward each other to collide.
		// be careful: if you push them too fast, they'll just fly through each other.
		// too slow and the black holes at the center collide and hilarity ensues.
		g0.Push(-4e3-8e2, 5e3+8e2)
		g1.Push(4e3+8e2, -5e3-8e2)

		width := 1.0e23
		galaxies := []Galaxy{g0, g1}

		initialUniverse := InitializeUniverse(galaxies, width)

		// now evolve the universe: feel free to adjust the following parameters.
		numGens := 20000
		time := 4e14
		theta := 0.5

		timePoints := BarnesHut(initialUniverse, numGens, time, theta)

		fmt.Println("Simulation run. Now drawing images.")
		canvasWidth := 1000
		frequency := 250
		scalingFactor := 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "collision2")
		fmt.Println("GIF drawn.")

	}

	// // the following sample parameters may be helpful for the "collide" command
	// // all units are in SI (meters, kg, etc.)
	// // but feel free to change the positions of the galaxies.

	// g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
	// g1 := InitializeGalaxy(500, 4e21, 3e22, 7e22)

	// // you probably want to apply a "push" function at this point to these galaxies to move
	// // them toward each other to collide.
	// // be careful: if you push them too fast, they'll just fly through each other.
	// // too slow and the black holes at the center collide and hilarity ensues.

	// width := 1.0e23
	// galaxies := []Galaxy{g0, g1}

	// initialUniverse := InitializeUniverse(galaxies, width)

	// // now evolve the universe: feel free to adjust the following parameters.
	// numGens := 500000
	// time := 2e14
	// theta := 0.5

	// timePoints := BarnesHut(initialUniverse, numGens, time, theta)

	// fmt.Println("Simulation run. Now drawing images.")
	// canvasWidth := 1000
	// frequency := 1000
	// scalingFactor := 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	// imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	// fmt.Println("Images drawn. Now generating GIF.")
	// gifhelper.ImagesToGIF(imageList, "galaxy")
	// fmt.Println("GIF drawn.")
}
