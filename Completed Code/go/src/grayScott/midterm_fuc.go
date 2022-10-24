package main

import "math"

//OrderedPair contains two float64 fields corresponding to
//the x and y coordinates of a point or vector in two-dimensional space.
type OrderedPair struct {
	x, y float64
}

//Boid represents our "bird" object. It contains two
//OrderedPair fields corresponding to its position, velocity, and acceleration.
type Boid struct {
	position, velocity, acceleration OrderedPair
}

//Sky represents a single time point of the simulation.
//It corresponds to a slice of Boid objects
type Sky []*Boid

//Insert your CountFlocks() function here, along with any subroutines and type declarations that you need.
func (s Sky) CountFlocks(flockDistance float64) int {
	n := len(s)
	count := n
	fmatrix := make([][]int, n)
	for i := range fmatrix {
		fmatrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if Distance(s[i], s[j]) <= flockDistance {
				if fmatrix[i][j] == 0 {
					fmatrix[i][j] = 1
					count -= 1
				}
			}
		}
	}
	return count

}

func Distance(a, b *Boid) float64 {
	deltaX := a.position.x - b.position.x
	deltaY := a.position.y - b.position.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
