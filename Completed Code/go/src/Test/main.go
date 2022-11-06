package main

import "fmt"

func main() {
	var currentuniverseg Universe
	currentuniverse := &currentuniverseg
	currentuniverse.stars = make([]*Star, 6)
	for i := 0; i < 6; i++ {
		var s Star
		currentuniverse.stars[i] = &s
	}

	currentuniverse.stars[0].position.x = 1
	currentuniverse.stars[0].position.y = 3
	currentuniverse.stars[0].mass = 0
	currentuniverse.stars[1].position.x = 2
	currentuniverse.stars[1].position.y = 1
	currentuniverse.stars[1].mass = 1
	currentuniverse.stars[2].position.x = 4
	currentuniverse.stars[2].position.y = 7
	currentuniverse.stars[2].mass = 2
	currentuniverse.stars[3].position.x = 8
	currentuniverse.stars[3].position.y = 9
	currentuniverse.stars[3].mass = 3
	currentuniverse.stars[4].position.x = 7
	currentuniverse.stars[4].position.y = 8
	currentuniverse.stars[4].mass = 4
	currentuniverse.stars[5].position.x = 1
	currentuniverse.stars[5].position.y = 1
	currentuniverse.stars[5].mass = 5
	currentuniverse.width = 10

	fmt.Println("check1")

	c := currentuniverse.BuildQuadtree()

	printtree(c.root)
	printtree(c.root.children[0])
	printtree(c.root.children[1])
	printtree(c.root.children[2])
	printtree(c.root.children[3])
	printtree(c.root.children[2].children[0])
	printtree(c.root.children[2].children[1])
	printtree(c.root.children[2].children[2])
	printtree(c.root.children[2].children[3])
	printtree(c.root.children[1].children[0])
	printtree(c.root.children[1].children[1])
	printtree(c.root.children[1].children[2])
	printtree(c.root.children[1].children[3])
	printtree(c.root.children[0].children[0])
	printtree(c.root.children[0].children[1])
	printtree(c.root.children[0].children[2])
	printtree(c.root.children[0].children[3])
	printtree(c.root.children[3].children[0])
	printtree(c.root.children[3].children[1])
	printtree(c.root.children[3].children[2])
	printtree(c.root.children[3].children[3])

}

func printtree(n *Node) {
	if n == nil {
		fmt.Println("nil")
	} else {
		if n.star == nil {
			fmt.Println("Starnil")
		} else {
			fmt.Println(n.star.mass)
		}
	}
}
