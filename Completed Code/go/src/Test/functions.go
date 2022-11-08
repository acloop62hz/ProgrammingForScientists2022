package main

import (
	"fmt"
	"math"
)

func (galaxy *Galaxy) Push(x, y float64) {
	for i := range *galaxy {
		(*galaxy)[i].velocity.x = x
		(*galaxy)[i].velocity.y = y
	}
}

// BarnesHut is our highest level function.
// Input: initial Universe object, a number of generations, and a time interval.
// Output: collection of Universe objects corresponding to updating the system
// over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)
	timePoints[0] = initialUniverse

	//now range over the number of generations and update the universe each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = timePoints[i-1].UpdateUniverse(time, theta)
		fmt.Println(i)
	}

	return timePoints
}

// UpdateUniverse updates a given universe over a specified time interval (in seconds).
// Input: A Universe object, a float time and Barnes-Hut parameter theta
// Output: A Universe object corresponding to simulating gravity over time seconds
func (currentUniverse *Universe) UpdateUniverse(time, theta float64) *Universe {
	newUniverse := currentUniverse.CopyUniverse()

	var quadTree QuadTree
	quadTree = currentUniverse.BuildQuadtree()
	//range over all stars in the universe and update their acceleration,
	//velocity, and position
	for i := range newUniverse.stars {
		newUniverse.stars[i].UpdateAcceleration(currentUniverse, &quadTree, theta)
		newUniverse.stars[i].UpdateVelocity(time)
		newUniverse.stars[i].UpdatePosition(time)
	}

	return &newUniverse
}

func (star *Star) UpdateAcceleration(currentUniverse *Universe, quadTree *QuadTree, theta float64) {

	force := ComputeNetForce(quadTree.root, star, theta)

	star.acceleration.x = force.x / star.mass
	star.acceleration.y = force.y / star.mass

}

func (star *Star) UpdateVelocity(time float64) {
	//new velocity is current velocity + acceleration * time
	star.velocity.x = star.velocity.x + star.acceleration.x*time
	star.velocity.y = star.velocity.y + star.acceleration.y*time
}

// UpdatePosition
// Input: a Body b and a time step (float64).
// Output: The OrderedPair corresponding to the updated position of the body after a single time step, using the body's current acceleration and velocity.
func (star *Star) UpdatePosition(time float64) {
	star.position.x = 0.5*star.acceleration.x*time*time + star.velocity.x*time + star.position.x
	star.position.y = 0.5*star.acceleration.y*time*time + star.velocity.y*time + star.position.y
}

func ComputeNetForce(root *Node, star *Star, theta float64) OrderedPair {
	var force OrderedPair
	force.x = 0
	force.y = 0
	if root == nil {
		return force
	} else if root.children == nil {
		if root.star == nil || *(root.star) == *(star) {
			return force
		} else {
			force = ComputeSingleStarForce(root.star, star)
		}
	} else {
		dist := Distance(root.star.position, star.position)
		if root.sector.width/dist > theta {
			for i := 0; i < 4; i++ {
				force.x += ComputeNetForce(root.children[i], star, theta).x
				force.y += ComputeNetForce(root.children[i], star, theta).y
			}
		} else {
			force.x += ComputeSingleStarForce(root.star, star).x
			force.y += ComputeSingleStarForce(root.star, star).y
		}
	}
	return force
}

func ComputeSingleStarForce(caster, receiver *Star) OrderedPair {
	var force OrderedPair

	dist := Distance(caster.position, receiver.position)
	dx := caster.position.x - receiver.position.x
	dy := caster.position.y - receiver.position.y

	F := G * caster.mass * receiver.mass / (dist * dist)
	force.x = F * dx / dist
	force.y = F * dy / dist

	return force

}

func Distance(a, b OrderedPair) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	d := math.Sqrt(dx*dx + dy*dy)
	return d
}

// CopyUniverse
// Input: a Universe object
// Output: a new Universe object, all of whose fields are copied over
// into the new Universe's fields. (Deep copy.)
func (currentUniverse *Universe) CopyUniverse() Universe {
	var newUniverse Universe

	newUniverse.width = currentUniverse.width

	//let's make the new universe's slice of Star objects
	numStars := len(currentUniverse.stars)
	stars := make([]Star, numStars)
	newUniverse.stars = make([]*Star, numStars)

	for i := range newUniverse.stars {
		newUniverse.stars[i] = &stars[i]
	}

	//now, copy all of the stars' fields into our new stars
	for i := range currentUniverse.stars {
		newUniverse.stars[i].mass = currentUniverse.stars[i].mass
		newUniverse.stars[i].radius = currentUniverse.stars[i].radius
		newUniverse.stars[i].position.x = currentUniverse.stars[i].position.x
		newUniverse.stars[i].position.y = currentUniverse.stars[i].position.y
		newUniverse.stars[i].velocity.x = currentUniverse.stars[i].velocity.x
		newUniverse.stars[i].velocity.y = currentUniverse.stars[i].velocity.y
		newUniverse.stars[i].acceleration.x = currentUniverse.stars[i].acceleration.x
		newUniverse.stars[i].acceleration.y = currentUniverse.stars[i].acceleration.y
		newUniverse.stars[i].red = currentUniverse.stars[i].red
		newUniverse.stars[i].green = currentUniverse.stars[i].green
		newUniverse.stars[i].blue = currentUniverse.stars[i].blue
	}

	return newUniverse
}

func (currentUniverse *Universe) BuildQuadtree() QuadTree {
	//initialize root node
	var root Node
	root.sector.width = currentUniverse.width
	root.sector.x = currentUniverse.width / 2.0
	root.sector.y = currentUniverse.width / 2.0

	//initialize the quadtree
	var quadTree QuadTree
	quadTree.root = &root

	//insert all stars in the universe into the tree
	for i := range currentUniverse.stars {
		if currentUniverse.stars[i].StarOutOfBound(currentUniverse.width) == false {
			continue
		}
		quadTree.root.InsertStar(currentUniverse.stars[i])
	}

	return quadTree

}

func (star *Star) StarOutOfBound(width float64) bool {
	if star.position.x > width || star.position.x < 0 || star.position.y > width || star.position.y < 0 {
		return false
	} else {
		return true
	}
}

func (proot *Node) InsertStar(star *Star) {
	if proot.children != nil {
		index := proot.FindInsertIndex(star)
		proot.children[index].InsertStar(star)
		proot.UpdateNode(proot.children[index])
	} else if proot.children == nil && proot.star == nil {
		proot.star = star
		return
	} else {
		prootStar := proot.star
		proot.MakeDummyNode()
		proot.InsertStar(prootStar)
		proot.InsertStar(star)
	}
}

func (proot *Node) UpdateNode(child *Node) {
	proot.star.position.x = (proot.star.position.x*proot.star.mass + child.star.position.x*child.star.mass) / (proot.star.mass + child.star.mass)
	proot.star.position.y = (proot.star.position.y*proot.star.mass + child.star.position.y*child.star.mass) / (proot.star.mass + child.star.mass)
	proot.star.mass = proot.star.mass + child.star.mass
}

func (proot *Node) MakeDummyNode() {
	proot.children = make([]*Node, 4)
	var dummyStar Star
	proot.star = &dummyStar
	for i := 0; i < 4; i++ {
		var n Node
		if i == 1 || i == 3 {
			n.sector.x = proot.sector.x + proot.sector.width/4.0
		} else {
			n.sector.x = proot.sector.x - proot.sector.width/4.0
		}
		if i == 0 || i == 1 {
			n.sector.y = proot.sector.y + proot.sector.width/4.0
		} else {
			n.sector.y = proot.sector.y - proot.sector.width/4.0
		}
		n.sector.width = proot.sector.width / 2.0
		proot.children[i] = &n
	}
}

func (proot *Node) FindInsertIndex(star *Star) int {

	if star.position.x < proot.sector.x && star.position.y < proot.sector.y {
		return 2
	} else if star.position.x < proot.sector.x && star.position.y > proot.sector.y {
		return 0
	} else if star.position.x >= proot.sector.x && star.position.y >= proot.sector.y {
		return 1
	} else {
		return 3
	}

}
