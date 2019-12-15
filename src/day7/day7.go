package main

import (
	"bufio"
	"fileutils"
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

func executeOpsCommand(arr []int64, input []int64) ([]int64, []int64, bool) {
	indexPtr := 0
	hasHalted := false

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
			// log.Printf("%d %d %d", arr, indexPtr, opCode)
			// Takes input and store in position
			dest := arr[indexPtr+1]
			if paramModes[0] == 1 {
				dest = int64(indexPtr + 1)
			}
			arr[dest] = input[0]

			// if len(input) > 1 {
			input = input[1:len(input)]
			// }

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
			hasHalted = true
			log.Print("HALTED!")
			continue

		default:
			log.Printf("Invalid opscode: %d", opCode)
			// Invalid opscode
			indexPtr = len(arr)
			break
		}
	}

	return arr, input, hasHalted
}

func amplifyChain(arr []int64, sequence []int64) int64 {
	var nextInput int64 = 0

	for _, s := range sequence {
		currentInput := []int64{s, int64(nextInput)}

		newArr := make([]int64, len(arr))
		copy(newArr, arr)
		_, newInput, _ := executeOpsCommand(newArr, currentInput)
		nextInput = newInput[0]
	}

	return nextInput
}

func amplifyLoopChain(arr []int64, sequence []int64) int64 {
	nextInput := []int64{0}

	// Create a copy of the ops software state
	softStates := [5][]int64{}

	for i := 0; i < 5; i++ {
		softStates[i] = arr
	}

	// shouldStop := false

	// for shouldStop == false {
	for i, s := range sequence {
		log.Printf("Sequence %d", s)
		currentInput := []int64{s}

		currentInput = append(currentInput, nextInput...)
		// newArr := arr
		log.Print(currentInput)

		_, newInput, hasHalted := executeOpsCommand(softStates[i], currentInput)
		nextInput = newInput

		if hasHalted {
			log.Print("Should stop")
			// shouldStop = true
			// copy(newArr, arr)
			// log.Print(newArr)
		}

		log.Print(nextInput)
	}

	log.Print("RUnning")
	// }

	log.Print("EXIT")

	return nextInput[0]
}

func isNumberValid(numToTest int) bool {
	arrCount := [10]int{}

	num := numToTest
	for num > 0 {
		remain := num % 10
		arrCount[remain]++

		num /= 10
	}

	if numToTest < 10000 {
		arrCount[0]++
	}

	for i := 0; i < 5; i++ {
		if arrCount[i] != 1 {
			return false
		}
	}

	return true
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

	var highestCount int64 = 0
	for i := 1234; i <= 43210; i++ {
		if isNumberValid(i) == false {
			continue
		}

		order := [5]int64{}
		temp := i

		if temp < 10000 {
			order[0] = 0
		}

		count := 4
		for temp > 0 {
			order[count] = int64(temp % 10)
			temp /= 10
			count--
		}

		sequence := []int64{}

		for _, s := range order {
			sequence = append(sequence, s)
		}

		currentCount := amplifyChain(arr, sequence)

		if currentCount > highestCount {
			highestCount = currentCount
		}

		log.Printf("%d: %d, %d", i, currentCount, sequence)
	}

	log.Print(highestCount)
}
