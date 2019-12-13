package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
		text := scanner.Text()
		arrStr := strings.Split(text, ",")

		for _, s := range arrStr {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			arr = append(arr, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}

func getParamsMode(instruc int64) [3]int64 {
	instruc = instruc / 100
	paramsModes := [3]int64{}

	for i := 0; i < 3; i++ {
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

func getParameterSrc(
	mode int64,
	paramIndex int,
	arr []int64,
	ptr int,
	relativeOffset int64,
) int64 {
	switch mode {
	case 0:
		return arr[ptr+paramIndex]

	case 1:
		return int64(ptr + paramIndex)

	case 2:
		return arr[ptr+paramIndex] + relativeOffset
	}

	log.Print("Invalid parameter mode")
	return arr[ptr+paramIndex]
}

func executeOpsCommand(arr []int64, input []int64) ([]int64, []int64) {
	indexPtr := 0
	var relativeOffset int64 = 0

	// Pad the memory with more 0's
	for i := 0; i < 1000; i++ {
		arr = append(arr, 0)
	}

	for indexPtr < len(arr) {
		instruc := arr[indexPtr]
		opCode := getOpcode(instruc)
		paramModes := getParamsMode(instruc)

		// log.Printf("%d %d %d %d", instruc, arr, indexPtr, opCode)

		switch opCode {
		case 1:
			log.Printf("%d %d %d %d", instruc, indexPtr, opCode, paramModes)
			src1 := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			src2 := getParameterSrc(paramModes[1], 2, arr, indexPtr, relativeOffset)
			dest := getParameterSrc(paramModes[2], 3, arr, indexPtr, relativeOffset)

			log.Printf("dest: %d, src1: %d, src2: %d", dest, src1, src2)
			arr[dest] = arr[src1] + arr[src2]
			indexPtr += 4
			break

		case 2:
			// Multiplication
			src1 := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			src2 := getParameterSrc(paramModes[1], 2, arr, indexPtr, relativeOffset)
			dest := getParameterSrc(paramModes[2], 3, arr, indexPtr, relativeOffset)

			arr[dest] = arr[src1] * arr[src2]
			indexPtr += 4
			break

		case 3:
			// Takes input and store in position
			src := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			arr[src] = input[0]

			input = input[1:len(input)]

			indexPtr += 2
			break

		case 4:
			// Store the output at param position
			src := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			input = append(input, arr[src])

			indexPtr += 2
			break

		case 5:
			src1 := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			src2 := getParameterSrc(paramModes[1], 2, arr, indexPtr, relativeOffset)

			if arr[src1] != 0 {
				indexPtr = int(arr[src2])
			} else {
				indexPtr += 3
			}
			break

		case 6:
			src1 := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			src2 := getParameterSrc(paramModes[1], 2, arr, indexPtr, relativeOffset)

			if arr[src1] == 0 {
				indexPtr = int(arr[src2])
			} else {
				indexPtr += 3
			}
			break

		case 7:
			src1 := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			src2 := getParameterSrc(paramModes[1], 2, arr, indexPtr, relativeOffset)
			dest := getParameterSrc(paramModes[2], 3, arr, indexPtr, relativeOffset)

			result := 0
			if arr[src1] < arr[src2] {
				result = 1
			}

			arr[dest] = int64(result)
			indexPtr += 4
			break

		case 8:
			src1 := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			src2 := getParameterSrc(paramModes[1], 2, arr, indexPtr, relativeOffset)
			dest := getParameterSrc(paramModes[2], 3, arr, indexPtr, relativeOffset)

			result := 0
			if arr[src1] == arr[src2] {
				result = 1
			}

			arr[dest] = int64(result)
			indexPtr += 4
			break

		case 9:
			src1 := getParameterSrc(paramModes[0], 1, arr, indexPtr, relativeOffset)
			relativeOffset += arr[src1]

			indexPtr += 2
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

func main() {
	arr := readFile("./input.txt")

	// Part 1
	// _, output := executeOpsCommand(arr, []int64{1})

	// Part 2
	_, output := executeOpsCommand(arr, []int64{2})
	log.Print(output)
}
