package main

import (
	"fileutils"
	"log"
	"math"
	"sort"
)

type Position struct {
	x int
	y int
}

func getAsteriods(data []string) []Position {
	asteroidList := []Position{}

	for i, s := range data {

		for j := 0; j < len(s); j++ {
			point := s[j : j+1]

			if point == "#" {
				asteroidList = append(asteroidList, Position{j, i})
			}
		}
	}

	return asteroidList
}
func getHighestVisibleAsteroidCount2(asteroidList []Position) int {
	highestVisibleAsteriod := 0
	asteroidCount := make([][]float64, len(asteroidList))

	for i, s := range asteroidList {

		asteroidCount[i] = []float64{}
		visibleAsteroidList := make(map[float64][]Position)
		for j := 0; j < len(asteroidList); j++ {
			height := float64(asteroidList[j].y - s.y)
			width := float64(asteroidList[j].x - s.x)
			angle := math.Atan(height / width)

			radAngle := angle
			if height > 0 && width == 0 {
				radAngle = math.Pi / 2
			} else if height < 0 && width == 0 {
				radAngle = 3 * math.Pi / 2
			} else if height >= 0 && width < 0 {
				radAngle += 3.14
			} else if height < 0 && width < 0 {
				radAngle += 3.14
			} else if height <= 0 && width > 0 {
				radAngle += 3.14 * 2
			}

			// Round to 2 decimal places
			roundedAngle := math.Round(radAngle*180/math.Pi*100-math.Pi/2) / 100

			if !math.IsNaN(roundedAngle) {
				newList := append(visibleAsteroidList[roundedAngle], asteroidList[j])

				// Sort it by manhanttan distance
				sort.SliceStable(newList, func(i, j int) bool {
					manhatDistX1 := math.Abs(float64(newList[i].x - s.x))
					manhatDistY1 := math.Abs(float64(newList[i].y - s.y))

					manhatDistX2 := math.Abs(float64(newList[j].x - s.x))
					manhatDistY2 := math.Abs(float64(newList[j].y - s.y))

					return ((manhatDistX1 + manhatDistY1) - (manhatDistX2 + manhatDistY2)) > 0
				})

				visibleAsteroidList[roundedAngle] = newList
			}

			asteroidCount[i] = append(asteroidCount[i], roundedAngle)
		}

		log.Print(visibleAsteroidList)

		if highestVisibleAsteriod < len(visibleAsteroidList) {
			highestVisibleAsteriod = len(visibleAsteroidList)
		}
	}

	return highestVisibleAsteriod
}

func getHighestVisibleAsteroidCount(asteroidList []Position) int {
	highestVisibleAsteriod := 0
	asteroidCount := make([][]float64, len(asteroidList))

	for i, s := range asteroidList {

		asteroidCount[i] = []float64{}
		visibleAsteroidList := make(map[float64]int)
		for j := 0; j < len(asteroidList); j++ {
			height := float64(asteroidList[j].y - s.y)
			width := float64(asteroidList[j].x - s.x)
			angle := math.Atan(height / width)

			radAngle := angle
			if height > 0 && width == 0 {
				radAngle = math.Pi / 2
			} else if height < 0 && width == 0 {
				radAngle = 3 * math.Pi / 2
			} else if height >= 0 && width < 0 {
				radAngle += 3.14
			} else if height < 0 && width < 0 {
				radAngle += 3.14
			} else if height <= 0 && width > 0 {
				radAngle += 3.14 * 2
			}

			// Round to 2 decimal places
			roundedAngle := math.Round(radAngle*180/math.Pi*100-math.Pi/2) / 100

			if !math.IsNaN(roundedAngle) {
				visibleAsteroidList[roundedAngle]++
			}

			asteroidCount[i] = append(asteroidCount[i], roundedAngle)
		}

		if highestVisibleAsteriod < len(visibleAsteroidList) {
			highestVisibleAsteriod = len(visibleAsteroidList)
		}
	}

	return highestVisibleAsteriod
}

func main() {
	data := fileutils.ReadFile("./input.txt")

	// Parse data
	asteroidList := getAsteriods(data)

	// Part 1
	visibleAsteriod := getHighestVisibleAsteroidCount(asteroidList)
	log.Print(visibleAsteriod)

	// Part 2
}
