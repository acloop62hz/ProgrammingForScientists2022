package main

//this file contains functions shared by the serial and parallel versions of our code.

import (
	"math"
	"math/rand"
	"runtime"
	"time"
)

// CopyBoard is a Board method that makes a deep copy of a board and returns
// a pointer to it.
func (b *Board) CopyBoard() *Board {
	var newBoard Board

	newBoard.width = b.width
	newBoard.height = b.height
	newBoard.particles = make([]*Particle, len(b.particles))

	for i, p := range b.particles {
		newBoard.particles[i] = p.CopyParticle()
	}

	return &newBoard
}

// CopyParticle is a Particle method that makes a deep copy of a Particle
// and returns a pointer to the new Particle.
func (p *Particle) CopyParticle() *Particle {
	var p2 Particle

	p2 = *p // shallow copy ok because all fields are elementary

	//but if you changed the data structure representing a Particle, this could be horrible!

	return &p2
}

// UpdateBoards takes a pointer to an initial Board object and a number of steps parameter.
// It returns a slice of pointers to Board objects corresponding to simulating diffusion
// over the number of steps given.
func UpdateBoards(initialBoard *Board, numSteps int) []*Board {
	boards := make([]*Board, numSteps+1)
	boards[0] = initialBoard

	for i := 1; i <= numSteps; i++ {
		boards[i] = boards[i-1].UpdateBoard()
	}

	return boards
}

// UpdateBoard is a Board method that returns a pointer to a new Board object
// corresponding to a single time step update of the Board.
func (b *Board) UpdateBoard() *Board {
	newBoard := b.CopyBoard()

	newBoard.Diffuse()

	return newBoard
}

// Diffuse is a Board method that diffuses each Particle in the Board over a single
// time step.
func (b *Board) Diffuse(isparallel bool) {
	if isparallel {
		numProcs := runtime.NumCPU()
		b.DiffuseParallel(numProcs)
	} else {
		for _, p := range b.particles {
			source := rand.NewSource(time.Now().UnixNano())
			generator := rand.New(source)
			p.RandStep(generator)
		}
	}

}

func (b *Board) DiffuseParallel(numProcs int) {
	n := len(b.particles)

	finished := make(chan bool, numProcs)

	for i := 0; i < numProcs; i++ {
		start := i * (n / numProcs)
		end := (i + 1) * (n / numProcs)
		if i < numProcs-1 {
			source := rand.NewSource(time.Now().UnixNano())
			generator := rand.New(source)
			go DiffuseOneCore(b.particles[start:end], generator, finished)
		} else {
			source := rand.NewSource(time.Now().UnixNano())
			generator := rand.New(source)
			go DiffuseOneCore(b.particles[start:], generator, finished)

		}
	}

	for i := 0; i < numProcs; i++ {
		<-finished
	}
}

func DiffuseOneCore(particles []*Particle, generator *(rand.Rand), finished chan bool) {

	for _, p := range particles {
		p.RandStep(generator)
	}

	finished <- true
}

// RandStep is a Particle method that moves the Particle by the Particle's diffusion rate
// parameter in a randomly chosen direction.

func (p *Particle) RandStep(generator *(rand.Rand)) {
	stepLength := p.diffusionRate
	angle := generator.Float64() * 2 * math.Pi
	p.position.x += stepLength * math.Cos(angle)
	p.position.y += stepLength * math.Sin(angle)
}

// InitializeBoard takes board parameters and initializes a Board with these parameters
// for a collection of randomly placed particles in the Board.
func InitializeBoard(boardWidth, boardHeight float64, numParticles int, particleRadius float64, diffusionRate float64, random bool) *Board {
	var b Board

	b.width = boardWidth
	b.height = boardHeight

	b.particles = make([]*Particle, numParticles)

	for i := range b.particles {
		var p Particle
		if random {
			p.position.x = rand.Float64() * boardWidth
			p.position.y = rand.Float64() * boardHeight
		} else {
			// default: non-random: assign all to center of board
			p.position.x = boardWidth / 2
			p.position.y = boardHeight / 2
		}
		p.radius = particleRadius
		p.diffusionRate = diffusionRate
		p.red, p.green, p.blue = 255, 255, 255
		b.particles[i] = &p
	}

	return &b
}
