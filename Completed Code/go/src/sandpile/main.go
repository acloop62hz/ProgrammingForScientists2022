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

	size, _ := strconv.Atoi(os.Args[1])

	pile, _ := strconv.Atoi(os.Args[2])

	placement := os.Args[3]

	var initialBoard Board

	if placement == "central" {
		initialBoard = InitializeBoardCentral(size, pile)
	} else if placement == "random" {
		initialBoard = InitializeBoardRandom(size, pile)
	} else {
		panic("wrong input")
	}

	// initialBoardCopy := CopyBoard(initialBoard)

	fmt.Println("Running simulation.")
	start := time.Now()
	stableBoard := SandpileSerial(initialBoard)
	//stableBoard := initialBoard
	elapsed := time.Since(start)
	log.Printf("Simulating sandpile in serial took %s", elapsed)

	fmt.Println("Simulation run. drawing.")
	canvasWidth := 50
	outFileName := "sandpile.png"

	DrawPNG(stableBoard, canvasWidth, outFileName)

	fmt.Println("PNG Generated")
}
