package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func readFile(path string) []int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	arr := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		arr = append(arr, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}

func CalculateFuel(module int64) int64 {
	return int64(math.Floor(float64(module/3)) - 2)
}

func CalculateFuelRepeat(mass int64) int64 {
	var totalFuel int64 = 0

	mass = CalculateFuel(mass)
	for mass > 0 {
		totalFuel += mass
		mass = CalculateFuel(mass)
	}

	return totalFuel
}

func main() {
	arr := readFile("./input.txt")

	var totalFuel int64 = 0
	for _, s := range arr {
		totalFuel += CalculateFuelRepeat(s)
	}

	log.Print(totalFuel)
}
