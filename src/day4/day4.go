package main

import (
	"fileutils"
	"log"
	"strconv"
	"strings"
)

func isIncreasingDigits(num int64) bool {
	prevRem := num % 10
	num /= 10

	for num > 0 {
		curRem := num % 10

		if curRem > prevRem {
			return false
		}

		num /= 10
		prevRem = curRem
	}

	return true
}

func hasDoubleDigits(num int64) bool {
	prevRem := num % 10
	num /= 10

	for num > 0 {
		curRem := num % 10

		if curRem == prevRem {
			return true
		}

		num /= 10
		prevRem = curRem
	}

	return false
}

func hasNoLargerDoubleDigits(num int64) bool {
	box := [10]int{}

	for num > 0 {
		curRem := num % 10
		box[curRem]++

		num /= 10
	}

	lastRepeatIndex := 0
	doubleIndex := 0
	for i, s := range box {
		if s == 2 {
			doubleIndex = i
		}

		if s > 2 {
			lastRepeatIndex = i
		}
	}

	return doubleIndex != 0 && doubleIndex != lastRepeatIndex
}

func main() {
	dataStr := fileutils.ReadFile("./input.txt")

	limits := strings.Split(dataStr[0], "-")
	min, _ := strconv.ParseInt(limits[0], 10, 64)
	max, _ := strconv.ParseInt(limits[1], 10, 64)

	count := 0
	for i := min; i <= max; i++ {
		// Part 1
		// if isIncreasingDigits(i) && hasDoubleDigits(i) {

		// Part 2
		if isIncreasingDigits(i) && hasDoubleDigits(i) && hasNoLargerDoubleDigits(i) {
			count++
		}
	}

	log.Print(count)
}
