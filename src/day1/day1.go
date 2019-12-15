package main

import (
	"log"
	"math"
	"fileutils"
	"strconv"
)

func calculateFuel(module int64) int64 {
	return int64(math.Floor(float64(module/3)) - 2)
}

func calculateFuelRepeat(mass int64) int64 {
	var totalFuel int64 = 0

	mass = calculateFuel(mass)
	for mass > 0 {
		totalFuel += mass
		mass = calculateFuel(mass)
	}

	return totalFuel
}

func parseData(dataArr []string) []int64 {
	arr := []int64{}

	for _, line := range dataArr {
		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		arr = append(arr, i)
	}

	return arr
}

func main() {
	dataStr := fileutils.ReadFile("./input.txt")
	arr := parseData(dataStr)

	var totalFuel int64 = 0
	for _, s := range arr {
		totalFuel += calculateFuelRepeat(s)
	}

	log.Print(totalFuel)
}
