package main

import (
	"math/rand"
)

func SimulateLanes(initiallane lane, numGens int, timeStep int) []lane {
	timePoints := make([]lane, numGens+1)
	timePoints[0] = initiallane

	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateLane(timePoints[i-1], timeStep)
	}

	return timePoints
}

func CopyLane(currentlane lane) lane {
	var newLane lane
	newLane.length = currentlane.length

	numVehicles := len(currentlane.vehicles)
	newVehicles := make([]vehicle, numVehicles)

	for i := 0; i < numVehicles; i++ {
		newVehicles[i].position = currentlane.vehicles[i].position
		newVehicles[i].velocity = currentlane.vehicles[i].velocity
	}

	newLane.vehicles = newVehicles

	return newLane

}

func UpdateLane(currentlane lane, timStep int) lane {

	newLane := CopyLane(currentlane)

	numVehicles := len(newLane.vehicles)
	for i := 0; i < numVehicles; i++ {
		newLane.vehicles[i] = UpdateVehicle(&currentlane, i)
	}

	return newLane
}

func UpdateVehicle(currentlane *lane, k int) vehicle {
	var newVehiclei vehicle
	newVehiclei.velocity = UpdateVelocity(currentlane, k)
	newVehiclei.position = UpdatePosition(currentlane, k)
	return newVehiclei
}

func UpdateVelocity(currentlane *lane, k int) int {

	var newVelocity int
	vm := 50
	frontIdex := findFrontVehicle(currentlane, k)
	NormRn := rand.NormFloat64()
	p := 1.65
	if NormRn > p {
		newVelocity = Randomisation(currentlane.vehicles[k])
	} else {
		newVelocity = NormalUpdateVelocity(currentlane.vehicles[frontIdex], currentlane.vehicles[k], vm)
	}
	return newVelocity

}

func UpdatePosition(currentlane *lane, k int) int {
	p := currentlane.vehicles[k].position + currentlane.vehicles[k].velocity
	if p > 1000 {
		p = p - 1000
	}
	return p
}

func findFrontVehicle(currentlane *lane, k int) int {
	numVehicles := len(currentlane.vehicles)

	minSpace := 1000
	frontIdex := k

	for i := 0; i < numVehicles; i++ {
		if i == k {
			continue
		}
		space := Distance(currentlane.vehicles[i], currentlane.vehicles[k])
		if space > 0 && space < minSpace {
			minSpace = space
			frontIdex = i
		}
	}

	return frontIdex

}

func Distance(vi, vk vehicle) int {
	return vi.position - vk.position
}

func Randomisation(vk vehicle) int {
	if vk.velocity-1 > 0 {
		return vk.velocity - 1
	} else {
		return 0
	}
}

func NormalUpdateVelocity(vf, vk vehicle, vm int) int {
	v1 := vk.velocity + 1
	v2 := vf.position - vk.position
	if v1 > v2 {
		if vm > v1 {
			return vm
		} else {
			return v1
		}
	} else {
		if vm > v2 {
			return vm
		} else {
			return v2
		}
	}
}
