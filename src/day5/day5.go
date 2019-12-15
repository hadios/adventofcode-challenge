package main

import (
	"fileutils"
	"log"
	"strconv"
	"strings"
)

func getParamsMode(instruc int64) [3]int64 {
	instruc = instruc / 100
	paramsModes := [3]int64{}

	for i := 0; i < 2; i++ {
		digit := instruc % 10
		paramsModes[i] = digit

		instruc = instruc / 10
	}

	return paramsModes
}

func getOpcode(instruc int64) int64 {
	// Get last two digits
	var opCode int64 = 0

	for i := 0; i < 2; i++ {
		digit := instruc % 10

		if i == 0 {
			opCode += digit
		} else {
			opCode += digit * 10
		}

		instruc = instruc / 10
	}

	return opCode
}

func executeOpsCommand(arr []int64, input []int64) ([]int64, []int64) {
	indexPtr := 0

	for indexPtr < len(arr) {
		instruc := arr[indexPtr]
		opCode := getOpcode(instruc)
		paramModes := getParamsMode(instruc)

		// log.Printf("%d %d %d", arr, indexPtr, opCode)

		switch opCode {
		case 1:
			dest := arr[indexPtr+3]

			src1 := arr[indexPtr+1]
			if paramModes[0] == 1 {
				src1 = int64(indexPtr + 1)
			}

			src2 := arr[indexPtr+2]
			if paramModes[1] == 1 {
				src2 = int64(indexPtr + 2)
			}

			arr[dest] = arr[src1] + arr[src2]
			indexPtr += 4
			break

		case 2:
			// Multiplication
			dest := arr[indexPtr+3]

			src1 := arr[indexPtr+1]
			if paramModes[0] == 1 {
				src1 = int64(indexPtr + 1)
			}

			src2 := arr[indexPtr+2]
			if paramModes[1] == 1 {
				src2 = int64(indexPtr + 2)
			}

			arr[dest] = arr[src1] * arr[src2]
			indexPtr += 4
			break

		case 3:
			// Takes input and store in position
			dest := arr[indexPtr+1]
			if paramModes[0] == 1 {
				dest = int64(indexPtr + 1)
			}
			arr[dest] = input[0]

			input = input[1:len(input)]

			indexPtr += 2
			break

		case 4:
			// Store the output at param position
			dest := arr[indexPtr+1]
			if paramModes[0] == 1 {
				dest = int64(indexPtr + 1)
			}

			input = append(input, arr[dest])

			indexPtr += 2
			break

		case 5:
			src1 := arr[indexPtr+1]
			if paramModes[0] == 1 {
				src1 = int64(indexPtr + 1)
			}

			src2 := arr[indexPtr+2]
			if paramModes[1] == 1 {
				src2 = int64(indexPtr + 2)
			}

			if arr[src1] != 0 {
				indexPtr = int(arr[src2])
			} else {
				indexPtr += 3
			}
			break

		case 6:
			src1 := arr[indexPtr+1]
			if paramModes[0] == 1 {
				src1 = int64(indexPtr + 1)
			}

			src2 := arr[indexPtr+2]
			if paramModes[1] == 1 {
				src2 = int64(indexPtr + 2)
			}

			if arr[src1] == 0 {
				indexPtr = int(arr[src2])
			} else {
				indexPtr += 3
			}
			break

		case 7:
			dest := arr[indexPtr+3]

			src1 := arr[indexPtr+1]
			if paramModes[0] == 1 {
				src1 = int64(indexPtr + 1)
			}

			src2 := arr[indexPtr+2]
			if paramModes[1] == 1 {
				src2 = int64(indexPtr + 2)
			}

			result := 0
			if arr[src1] < arr[src2] {
				result = 1
			}

			arr[dest] = int64(result)
			indexPtr += 4
			break

		case 8:
			dest := arr[indexPtr+3]

			src1 := arr[indexPtr+1]
			if paramModes[0] == 1 {
				src1 = int64(indexPtr + 1)
			}

			src2 := arr[indexPtr+2]
			if paramModes[1] == 1 {
				src2 = int64(indexPtr + 2)
			}

			result := 0
			if arr[src1] == arr[src2] {
				result = 1
			}

			arr[dest] = int64(result)
			indexPtr += 4
			break

		case 99:
			indexPtr = len(arr)
			continue

		default:
			log.Printf("Invalid opscode: %d", opCode)
			// Invalid opscode
			indexPtr = len(arr)
			break
		}
	}

	return arr, input
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
	dataArr := fileutils.ReadFile("./input.txt")
	arr := parseData(dataArr)

	// Part 1
	// _, output := executeOpsCommand(arr, int64(1))
	// log.Print(output)

	// Part 2
	_, output := executeOpsCommand(arr, []int64{5})
	log.Print(output)
}
