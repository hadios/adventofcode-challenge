package main

import (
	"testing"
)

type testData struct {
	input  []string
	result int
}

func Test(t *testing.T) {
	/*
	** Part 1
	 */
	testDataArr := []testData{
		testData{
			input:  []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
			result: 42,
		},
	}

	for _, s := range testDataArr {
		planetList, firstPlanet := parsePlanetOrbitArr(s.input)
		visitedPlanet := make(map[string]int)
		result := explore(planetList, visitedPlanet, firstPlanet, 0)

		if result != s.result {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
		}
	}

	/*
	** Part 2
	 */
	testDataArr2 := []testData{
		testData{
			input: []string{
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
				"K)YOU",
				"I)SAN",
			},
			result: 4,
		},
	}

	for _, s := range testDataArr2 {
		planetList, _ := parsePlanetOrbitArr(s.input)
		visitedPlanet := make(map[string]int)
		result := travel(
			planetList,
			visitedPlanet,
			"YOU",
			"SAN",
			0,
		) - 2

		if result != s.result {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
		}
	}

}
