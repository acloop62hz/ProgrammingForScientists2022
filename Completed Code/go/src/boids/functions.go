package main

import "math"

//SimulateBoids simulates skies over a series of snap shots separated by equal unit time.
//Input: an initial sky, a number of generations, and a time parameter (in seconds).
//Output: a slice of sky objects corresponding to simulating the boid force over the number of generations time points.
func SimulateBoids(initialSky Sky, numGens int, timeStep float64) []Sky {
	timePoints := make([]Sky, numGens+1)
	timePoints[0] = initialSky

	//range over the number of generations and update the sky each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateSky(timePoints[i-1], timeStep)
	}

	return timePoints
}

//CopySky
//Input: a Sky object
//Output: a new Sky object (Deep copy.)
func CopySky(currentSky Sky) Sky {
	var newSky Sky

	newSky.width = currentSky.width
	newSky.alignmentFactor = currentSky.alignmentFactor
	newSky.cohesionFactor = currentSky.cohesionFactor
	newSky.separationFactor = currentSky.separationFactor
	newSky.proximity = currentSky.proximity
	newSky.maxBoidSpeed = currentSky.maxBoidSpeed

	//make the new Sky's slice of Boid objects
	numboids := len(currentSky.boids)
	newSky.boids = make([]Boid, numboids)

	//copy all of the boids' fields into our new boids
	for i := range currentSky.boids {
		newSky.boids[i].position.x = currentSky.boids[i].position.x
		newSky.boids[i].position.y = currentSky.boids[i].position.y
		newSky.boids[i].velocity.x = currentSky.boids[i].velocity.x
		newSky.boids[i].velocity.y = currentSky.boids[i].velocity.y
		newSky.boids[i].acceleration.x = currentSky.boids[i].acceleration.x
		newSky.boids[i].acceleration.y = currentSky.boids[i].acceleration.y
	}

	return newSky
}

//UpdateSky updates a given sky over a specified time interval (in seconds).
//Input: A Sky object and a float time.
//Output: A Sky object corresponding to simulating gravity over time seconds, assuming that acceleration is constant over this time.
func UpdateSky(currentSky Sky, time float64) Sky {
	//make a deep copu of current sky
	newSky := CopySky(currentSky)

	//range over all boids in the Sky and update their acceleration,velocity, and position
	for i := range newSky.boids {
		newSky.boids[i].acceleration = UpdateAcceleration(currentSky, newSky.boids[i])
		newSky.boids[i].velocity = UpdateVelocity(newSky.boids[i], time, newSky.maxBoidSpeed)
		newSky.boids[i].position = UpdatePosition(newSky.boids[i], time, newSky.width)
	}

	return newSky
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

//ComputeSeparationForce
//Input: Two body objects b1 and b2, a parameter separationFactor.
//Output: The force due to Separation rule acting on b1 subject to b2.
func ComputeSeparationForce(b1, b2 Boid, separationFactor, dist float64) OrderedPair {
	var force OrderedPair

	deltaX := b1.position.x - b2.position.x
	deltaY := b1.position.y - b2.position.y

	force.x = separationFactor * deltaX / (dist * dist)
	force.y = separationFactor * deltaY / (dist * dist)

	return force
}

//ComputeAlignmentForce
//Input: Two body objects b1 and b2, a parameter alignmentFactor.
//Output: The force due to Alignment rule acting on b1 subject to b2.
func ComputeAlignmentForce(b1, b2 Boid, alignmentFactor, dist float64) OrderedPair {
	var force OrderedPair

	force.x = alignmentFactor * b2.velocity.x / dist
	force.y = alignmentFactor * b2.velocity.y / dist

	return force
}

//ComputeAlignmentForce
//Input: Two body objects b1 and b2, a parameter cohesionFactor.
//Output: The force due to Cohesion rule acting on b1 subject to b2.
func ComputeCohesionForce(b1, b2 Boid, cohesionFactor, dist float64) OrderedPair {
	var force OrderedPair

	deltaX := b2.position.x - b1.position.x
	deltaY := b2.position.y - b1.position.y

	force.x = cohesionFactor * deltaX / dist
	force.y = cohesionFactor * deltaY / dist

	return force
}

//ComputeAverageForce
//Input: A slice of Boid objects and an individual body, parameters for separation, alignment, and cohesion, priximity
//Output: The average force vector (OrderedPair) acting on the given body
//due to separation force, alignment force and cohesion force from all other boids in the Sky
func ComputeAverageForce(boids []Boid, b Boid, separationFactor, alignmentFactor, cohesionFactor, proximity float64) OrderedPair {
	var netForce OrderedPair

	closeNum := 0
	netForce.x = 0.0
	netForce.y = 0.0

	for i := range boids {
		//only do a force computation if current boid is not the input Boid
		if boids[i] != b {
			dist := Distance(b.position, boids[i].position)
			// only calculate force if the current boid and the input boid are within the given distance
			if dist > proximity {
				netForce.x += 0.0
				netForce.y += 0.0
			} else {
				// count number the boids that interact with the input one
				closeNum += 1
				// calculate three force
				separationForce := ComputeSeparationForce(b, boids[i], separationFactor, dist)
				alignmentForce := ComputeAlignmentForce(b, boids[i], alignmentFactor, dist)
				cohesionForce := ComputeCohesionForce(b, boids[i], cohesionFactor, dist)
				//now add its components into net force components
				netForce.x += separationForce.x + alignmentForce.x + cohesionForce.x
				netForce.y += separationForce.y + alignmentForce.y + cohesionForce.y
			}
		}
	}
	// do not do division if there's no nearby boids
	if closeNum != 0 {
		netForce.x = netForce.x / float64(closeNum)
		netForce.y = netForce.y / float64(closeNum)
	}

	return netForce
}

//UpdateAcceleration
//Input: Sky object and a boid B
//Output: The net acceleration on B due to the average force calculated by every boid in the Sky
func UpdateAcceleration(currentSky Sky, b Boid) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on b
	force := ComputeAverageForce(currentSky.boids, b, currentSky.separationFactor, currentSky.alignmentFactor, currentSky.cohesionFactor, currentSky.proximity)

	//now, calculate acceleration (F = ma)
	accel.x = force.x
	accel.y = force.y

	return accel
}

//UpdateVelocity
//Input: a Body object and a time step (float64).
//Output: The orderedPair corresponding to the velocity of this object
//after a single time step, using the body's current acceleration.
func UpdateVelocity(b Boid, time, maxSpeed float64) OrderedPair {
	var vel OrderedPair

	//new velocity is current velocity + acceleration * time
	vel.x = b.velocity.x + b.acceleration.x*time
	vel.y = b.velocity.y + b.acceleration.y*time

	speed := math.Sqrt(vel.x*vel.x + vel.y*vel.y)

	maxSpeedX := maxSpeed * math.Abs(vel.x) / speed
	maxSpeedY := maxSpeed * math.Abs(vel.y) / speed

	//limit the speed to given maxSpeed

	if math.Abs(vel.x) > maxSpeedX {
		vel.x = vel.x * maxSpeedX / math.Abs(vel.x)
	}

	if math.Abs(vel.y) > maxSpeedY {
		vel.y = vel.y * maxSpeedY / math.Abs(vel.y)
	}

	return vel
}

//UpdatePosition
//Input: a Body b and a time step (float64).
//Output: The OrderedPair corresponding to the updated position of the body after a single time step, using the body's current acceleration and velocity.
func UpdatePosition(b Boid, time, width float64) OrderedPair {
	var pos OrderedPair

	pos.x = 0.5*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	pos.y = 0.5*b.acceleration.y*time*time + b.velocity.y*time + b.position.y

	// when boids fly out of the canvas
	for pos.x < 0 {
		pos.x = width + pos.x
	}
	for pos.x > width {
		pos.x = pos.x - width
	}
	for pos.y < 0 {
		pos.y = width + pos.y
	}
	for pos.y > width {
		pos.y = pos.y - width
	}
	return pos
}
