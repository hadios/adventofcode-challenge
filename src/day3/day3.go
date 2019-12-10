package main

import (
	"bufio"
	"log"
	"math"
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
		arr = append(arr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}

func decomposeInstructionString(instr string) []string {
	arrStr := strings.Split(instr, ",")
	return arrStr
}

func drawWires(wireInputs []string) int {
	const size = 5000
	wireGrid := make(map[string]uint8)

	intersections := []string{}

	for index, s := range wireInputs {
		instrucs := decomposeInstructionString(s)

		xPos := size / 2
		yPos := size / 2

		key := strconv.Itoa(xPos) + "#" + strconv.Itoa(yPos)
		wireGrid[key] = 0000

		for _, t := range instrucs {
			dir := t[:1]
			steps, _ := strconv.ParseInt(t[1:len(t)], 10, 64)

			xDir := 0
			yDir := 0
			switch dir {
			case "U":
				yDir = 1
				break
			case "D":
				yDir = -1
				break
			case "L":
				xDir = -1
				break
			case "R":
				xDir = 1
				break
			}

			for i := 0; i < int(steps); i++ {
				xPos += xDir
				yPos += yDir

				key = strconv.Itoa(xPos) + "#" + strconv.Itoa(yPos)
				wireGrid[key] |= uint8(index + 1)

				if wireGrid[key] > 2 {
					intersections = append(intersections, key)
				}
			}
		}
	}

	min := size * size
	for _, s := range intersections {
		strArr := strings.Split(s, "#")

		x, _ := strconv.ParseInt(strArr[0], 10, 64)
		manhanValX := int(math.Abs(float64(size/2 - x)))

		y, _ := strconv.ParseInt(strArr[1], 10, 64)
		manhanValY := int(math.Abs(float64(size/2 - y)))

		newMin := manhanValX + manhanValY
		if min > newMin {
			min = newMin
		}
	}

	return min
}

func drawSignalWires(wireInputs []string) int {
	const size = 5000
	wireGrid1 := make(map[string]int)
	wireGrid2 := make(map[string]int)

	intersections := []string{}

	for index, s := range wireInputs {
		instrucs := decomposeInstructionString(s)

		xPos := size / 2
		yPos := size / 2

		currentWireGrid := wireGrid1
		if index == 1 {
			currentWireGrid = wireGrid2
		}

		totalStepsCount := 0
		for _, t := range instrucs {
			dir := t[:1]
			steps, _ := strconv.ParseInt(t[1:len(t)], 10, 64)

			xDir := 0
			yDir := 0
			switch dir {
			case "U":
				yDir = 1
				break
			case "D":
				yDir = -1
				break
			case "L":
				xDir = -1
				break
			case "R":
				xDir = 1
				break
			}

			for i := 0; i < int(steps); i++ {
				xPos += xDir
				yPos += yDir
				totalStepsCount++

				key := strconv.Itoa(xPos) + "#" + strconv.Itoa(yPos)

				if currentWireGrid[key] == 0 {
					currentWireGrid[key] = totalStepsCount
				}

				if wireGrid1[key] > 0 && wireGrid2[key] > 0 {
					intersections = append(intersections, key)
				}
			}
		}
	}

	min := size * size
	for _, s := range intersections {
		newMin := wireGrid1[s] + wireGrid2[s]

		if min > newMin {
			min = newMin
		}
	}

	return min
}

func main() {
	wireInputs := readFile("./input.txt")

	manDist := drawSignalWires(wireInputs)
	log.Print(manDist)
}
