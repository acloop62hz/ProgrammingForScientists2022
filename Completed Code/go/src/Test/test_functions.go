package main

import "fmt"

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
		fmt.Println("startInset", i)
		quadTree.root.InsertStar(currentUniverse.stars[i])
		fmt.Println("EndInset", i)
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

func (proot *Node) MakeDummyNode() {
	proot.children = make([]*Node, 4)
	proot.star = nil
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
