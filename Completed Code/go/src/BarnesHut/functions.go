package main

//BarnesHut is our highest level function.
//Input: initial Universe object, a number of generations, and a time interval.
//Output: collection of Universe objects corresponding to updating the system
//over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)
	timePoints[0] = initialUniverse

	//now range over the number of generations and update the universe each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = timePoints[i-1].UpdateUniverse(time, theta)
	}

	return timePoints
}

//UpdateUniverse updates a given universe over a specified time interval (in seconds).
//Input: A Universe object, a float time and Barnes-Hut parameter theta
//Output: A Universe object corresponding to simulating gravity over time seconds
func (currentUniverse *Universe) UpdateUniverse(time, theta float64) *Universe {
	newUniverse := currentUniverse.CopyUniverse()

	//range over all stars in the universe and update their acceleration,
	//velocity, and position
	for i := range newUniverse.stars {
		newUniverse.stars[i].UpdateAcceleration(currentUniverse)
		newUniverse.stars[i].UpdateVelocity(time)
		newUniverse.stars[i].UpdatePosition(time)
	}

	return &newUniverse
}

//CopyUniverse
//Input: a Universe object
//Output: a new Universe object, all of whose fields are copied over
//into the new Universe's fields. (Deep copy.)
func (currentUniverse *Universe) CopyUniverse() Universe {
	var newUniverse Universe

	newUniverse.width = currentUniverse.width

	//let's make the new universe's slice of Star objects
	numStars := len(currentUniverse.stars)
	newUniverse.stars = make([]*Star, numStars)

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
	root.sector.x = currentUniverse.width / 2
	root.sector.y = currentUniverse.width / 2

	//initialize the quadtree
	var quadTree QuadTree
	quadTree.root = &root
	pQuadTree := &quadTree

	//insert all stars in the universe into the tree
	for i := range currentUniverse.stars {
		if currentUniverse.stars[i].StarOutOfBound(currentUniverse.width) == true {
			continue
		}
		pQuadTree.root.InsertStar(currentUniverse.stars[i])
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
	} else {
		if proot.star == nil {
			proot.star = star
		} else {
			prootStar := proot.star
			*proot = proot.MakeDummyNode()
			proot.InsertStar(prootStar)
			proot.InsertStar(star)
		}
	}
}

func (proot *Node) MakeDummyNode() Node {
	var dummyNode Node
	dummyNode.children = make([]*Node, 4)
	dummyNode.sector.x = proot.sector.x
	dummyNode.sector.y = proot.sector.y
	dummyNode.sector.width = proot.sector.width
	for i := 0; i < 4; i++ {
		if i == 1 || i == 3 {
			dummyNode.children[i].sector.x = proot.sector.x + proot.sector.width/2
		} else {
			dummyNode.children[i].sector.x = proot.sector.x - proot.sector.width/2
		}
		if i == 0 || i == 1 {
			dummyNode.children[i].sector.y = proot.sector.y + proot.sector.width/2
		} else {
			dummyNode.children[i].sector.y = proot.sector.y - proot.sector.width/2
		}
		dummyNode.sector.width = proot.sector.width / 2
	}

	return dummyNode
}

func (proot *Node) FindInsertIndex(star *Star) int {

	if star.position.x < proot.sector.x && star.position.y < proot.sector.y {
		return 3
	} else if star.position.x < proot.sector.x && star.position.y > proot.sector.y {
		return 1
	} else if star.position.x >= proot.sector.x && star.position.y >= proot.sector.y {
		return 2
	} else {
		return 4
	}

}
