package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	arr := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		arr = strings.Split(text, "-")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}

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
	limits := readFile("./input.txt")

	min, _ := strconv.ParseInt(limits[0], 10, 64)
	max, _ := strconv.ParseInt(limits[1], 10, 64)

	count := 0
	for i := min; i <= max; i++ {
		// if isValidPassword_Part2(i) {
		if isIncreasingDigits(i) && hasDoubleDigits(i) && hasNoLargerDoubleDigits(i) {
			log.Print(i)
			count++
		}
	}

	log.Print(count)
}
