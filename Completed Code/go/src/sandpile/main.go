package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println("SandPile simulator")

	// read parameters from stdin
	size, _ := strconv.Atoi(os.Args[1])
	pile, _ := strconv.Atoi(os.Args[2])
	placement := os.Args[3]

	// initialize board according to the placement input
	var initialBoard Board
	if placement == "central" {
		initialBoard = InitializeBoardCentral(size, pile)
	} else if placement == "random" {
		initialBoard = InitializeBoardRandom(size, pile)
	} else {
		panic("wrong input")
	}

	// make a copy of the initialized board for two separate tests
	initialBoardCopy := CopyBoard(initialBoard)

	// run serial program and track the time
	fmt.Println("Running simulation.")
	start := time.Now()
	stableBoardSerial := SandpileSerial(&initialBoard)
	elapsed := time.Since(start)
	log.Printf("Simulating sandpile in serial took %s", elapsed)

	// run parallel program and trak the time
	start = time.Now()
	stableBoardParallel := SandpileMultiprocs(&initialBoardCopy, 4)
	elapsed = time.Since(start)
	log.Printf("Simulating sandpile in parallel took %s", elapsed)

	fmt.Println("Simulation run. drawing.")

	// draw the stable board to PNG
	canvasWidth := 20
	outFileNameSerial := "serial.png"
	outFileNameParallel := "parallel.png"
	DrawPNG(stableBoardSerial, canvasWidth, outFileNameSerial)
	DrawPNG(stableBoardParallel, canvasWidth, outFileNameParallel)

	fmt.Println("PNG Generated")
}
