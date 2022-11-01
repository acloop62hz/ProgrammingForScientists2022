package main

type vehicle struct {
	velocity int
	position int
}

type lane struct {
	vehicles []vehicle
	length   int
}
