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
		timepoints[i] = UpdateUniverse(timepoints[i-1], time)
	}
	return timepoints
}

func UpdateUniverse(currentUniverse Universe, time float64) Universe {
	newUniverse := currentUniverse
	for i, b := range currentUniverse {
		b.acceleration = UpdateAcceleration()
	}
}

func UpdateVelocity(b Body, time float64) OrderedPair {
	var v OrderedPair
	v.x = b.acceleration.x*time + b.velocity.x
	v.y = b.acceleration.y*time + b.velocity.y
	return v
}

func UpdatePosition(b Body, time float64) OrderedPair {
	var p OrderedPair
	p.x = 1.0/2.0*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	p.y = 1.0/2.0*b.acceleration.y*time*time + b.velocity.y*time + b.position.y
	return p
}
