package main

import (
	"log"
	"fileutils"
	"strconv"
	"strings"
)

func executeOpsCommand(arr []int64) []int64 {
	indexPtr := 0

	for indexPtr < len(arr) {
		opCode := arr[indexPtr]

		switch opCode {
		case 1:
			dest := arr[indexPtr+3]
			src1 := arr[indexPtr+1]
			src2 := arr[indexPtr+2]
			arr[dest] = arr[src1] + arr[src2]
			indexPtr += 4

			break

		case 2:
			// Multiplication
			dest := arr[indexPtr+3]
			src1 := arr[indexPtr+1]
			src2 := arr[indexPtr+2]
			arr[dest] = arr[src1] * arr[src2]
			indexPtr += 4
			break

		case 99:
			indexPtr = len(arr)
			continue

		default:
			// Invalid opscode
			indexPtr = len(arr)
			break
		}
	}

	return arr
}

func runInitialState(x int64, y int64, arr []int64) []int64 {
	arr[1] = x
	arr[2] = y

	return arr
}

func parseData(dataArr []string) []int64 {
	arr := []int64{}

	for _, line := range dataArr {
		arrStr := strings.Split(line, ",")

		for _, s := range arrStr {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			arr = append(arr, i)
		}
	}

	return arr
}

func main() {
	var numToSearch int64 = 19690720

	dataArr := fileutils.ReadFile("./input.txt")

	// Brute force search
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			arr := parseData(dataArr)

			modifiedArr := runInitialState(int64(i), int64(j), arr)
			finalArr := executeOpsCommand(modifiedArr)

			if finalArr[0] == numToSearch {
				log.Print("found", i, j)
				log.Print("Answer: ", (100*i + j))
			}
		}
	}

}
