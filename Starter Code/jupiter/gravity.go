package main

import (
	"math"
)

//let's place our gravity simulation functions here.

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func SimulateGravity(initialUniverse Universe, numGen int, time float64) []Universe {
	timepoints := make([]Universe, numGen+1)
	timepoints[0] = initialUniverse
	for i := 1; i <= numGen; i++ {
		timepoints[i] = UpdateUniverse(timepoints[i-1])
	}
}

func UpdateUniverse(currentUniverse Universe, time float64) Universe {

}
