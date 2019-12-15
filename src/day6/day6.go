package main

import (
	"fileutils"
	"log"
	"strings"
)

type ListOfPlanets struct {
	planets []string
}

func travel(
	planetList map[string]*ListOfPlanets,
	visitedPlanets map[string]int,
	currentPlanet string,
	destPlanet string,
	orbitDist int,
) int {
	if visitedPlanets[currentPlanet] == 1 {
		return 999999
	}

	visitedPlanets[currentPlanet] = 1

	if currentPlanet == destPlanet {
		return orbitDist
	}
	if planetList[currentPlanet] == nil {
		return 999999
	}

	totalCount := 999999
	for _, s := range planetList[currentPlanet].planets {
		currentJumpCount := travel(planetList, visitedPlanets, s, destPlanet, orbitDist+1)

		if currentJumpCount < totalCount {
			totalCount = currentJumpCount
		}
	}

	return totalCount
}

func explore(
	planetList map[string]*ListOfPlanets,
	visitedPlanets map[string]int,
	currentPlanet string,
	orbitDist int,
) int {
	if planetList[currentPlanet] == nil {
		return orbitDist
	}

	if visitedPlanets[currentPlanet] == 1 {
		return 0
	}

	visitedPlanets[currentPlanet] = 1

	totalCount := orbitDist
	for _, s := range planetList[currentPlanet].planets {
		totalCount += explore(planetList, visitedPlanets, s, orbitDist+1)
	}

	return totalCount
}

func parsePlanetOrbitArr(list []string) (map[string]*ListOfPlanets, string) {
	planetList := make(map[string]*ListOfPlanets)
	planetOrbitCount := make(map[string]int)

	for _, s := range list {
		planetArr := strings.Split(s, ")")

		// Update key value
		if planetList[planetArr[0]] == nil {
			planetList[planetArr[0]] = &ListOfPlanets{}
		}
		planetList[planetArr[0]].planets = append(planetList[planetArr[0]].planets, planetArr[1])

		if planetList[planetArr[1]] == nil {
			planetList[planetArr[1]] = &ListOfPlanets{}
		}
		planetList[planetArr[1]].planets = append(planetList[planetArr[1]].planets, planetArr[0])

		planetOrbitCount[planetArr[0]]++
		planetOrbitCount[planetArr[0]]--
		planetOrbitCount[planetArr[1]]++
	}

	firstPlanet := ""
	for i, m := range planetOrbitCount {
		if m == 0 {
			firstPlanet = i
			break
		}
	}

	return planetList, firstPlanet
}

func main() {
	arr := fileutils.ReadFile("./input.txt")

	planetList, _ := parsePlanetOrbitArr(arr)
	visitedPlanet := make(map[string]int)

	// Part 1
	// orbitCount := explore(planetList, visitedPlanet, firstPlanet, 0)

	// Part 2
	orbitCount := travel(
		planetList,
		visitedPlanet,
		"YOU",
		"SAN",
		0,
	) - 2

	log.Print(orbitCount)
}
