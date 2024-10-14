package main

import "fmt"

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	totalOpponetEnergy := 0
	for _, v := range energy {
		totalOpponetEnergy += v
	}
	EnergyNeed := 0
	if initialEnergy <= totalOpponetEnergy {
		EnergyNeed = totalOpponetEnergy - initialEnergy + 1
	}
	ExpirenceNeed := 0
	currentExpirence := initialExperience
	for _, v := range experience {
		if currentExpirence <= v {
			ExpirenceNeed += v - currentExpirence + 1
			currentExpirence = v + 1
		}
		currentExpirence += v
	}
	return EnergyNeed + ExpirenceNeed
}

func main() {
	initialEnergy := 5
	initialExperience := 3
	energy := []int{1, 4, 3, 2}
	experience := []int{2, 6, 3, 1}
	fmt.Println(minNumberOfHours(initialEnergy, initialExperience, energy, experience))
}
