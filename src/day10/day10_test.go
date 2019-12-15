package main

import (
	"testing"
)

type testData struct {
	input  []string
	result int
}

type testData2 struct {
	input  []int64
	input2 []int64
	result int64
}

func Equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Test(t *testing.T) {
	/*
	** Part 1
	 */
	// testDataArr := []testData{
	// 	testData{
	// 		input: []string{
	// 			".#..#",
	// 			".....",
	// 			"#####",
	// 			"....#",
	// 			"...##",
	// 		},
	// 		result: 8,
	// 	},
	// 	testData{
	// 		input: []string{
	// 			"......#.#.",
	// 			"#..#.#....",
	// 			"..#######.",
	// 			".#.#.###..",
	// 			".#..#.....",
	// 			"..#....#.#",
	// 			"#..#....#.",
	// 			".##.#..###",
	// 			"##...#..#.",
	// 			".#....####",
	// 		},
	// 		result: 33,
	// 	},
	// 	testData{
	// 		input: []string{
	// 			"#.#...#.#.",
	// 			".###....#.",
	// 			".#....#...",
	// 			"##.#.#.#.#",
	// 			"....#.#.#.",
	// 			".##..###.#",
	// 			"..#...##..",
	// 			"..##....##",
	// 			"......#...",
	// 			".####.###.",
	// 		},
	// 		result: 35,
	// 	},
	// 	testData{
	// 		input: []string{
	// 			".#..#..###",
	// 			"####.###.#",
	// 			"....###.#.",
	// 			"..###.##.#",
	// 			"##.##.#.#.",
	// 			"....###..#",
	// 			"..#.#..#.#",
	// 			"#..#.#.###",
	// 			".##...##.#",
	// 			".....#.#..",
	// 		},
	// 		result: 41,
	// 	},
	// 	testData{
	// 		input: []string{
	// 			".#..##.###...#######",
	// 			"##.############..##.",
	// 			".#.######.########.#",
	// 			".###.#######.####.#.",
	// 			"#####.##.#.##.###.##",
	// 			"..#####..#.#########",
	// 			"####################",
	// 			"#.####....###.#.#.##",
	// 			"##.#################",
	// 			"#####.##.###..####..",
	// 			"..######..##.#######",
	// 			"####.##.####...##..#",
	// 			".#####..#.######.###",
	// 			"##...#.##########...",
	// 			"#.##########.#######",
	// 			".####.#.###.###.#.##",
	// 			"....##.##.###..#####",
	// 			".#.#.###########.###",
	// 			"#.#.#.#####.####.###",
	// 			"###.##.####.##.#..##",
	// 		},
	// 		result: 210,
	// 	},
	// }

	// for _, s := range testDataArr {
	// 	asteroidList := getAsteriods(s.input)
	// 	result := getHighestVisibleAsteroidCount(asteroidList)

	// 	if result != s.result {
	// 		t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
	// 	}
	// }

	/*
	** Part 2
	 */
	testDataArr2 := []testData{
		testData{
			input: []string{
				".#....#####...#..",
				"##...##.#####..##",
				"##...#...#.#####.",
				"..#.....X...###..",
				"..#.#.....#....##",
			},
			result: 8,
		},
	}

	for _, s := range testDataArr2 {
		asteroidList := getAsteriods(s.input)
		result := getHighestVisibleAsteroidCount2(asteroidList)

		if result != s.result {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
		}
	}
}
