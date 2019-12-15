package main

import (
	"fileutils"
	"log"
	"strconv"
)

func countDigits(num string) [10]int {
	digitBucket := [10]int{}

	for i := 0; i < len(num); i++ {
		remain, _ := strconv.ParseInt(num[i:i+1], 10, 64)
		digitBucket[remain]++
	}

	return digitBucket
}

func parseImageData(data string, width int, height int) []string {
	layerSize := width * height
	totalLayers := len(data) / int(layerSize)

	layerArr := []string{}
	for i := 0; i < totalLayers; i++ {
		layerString := data[0:layerSize]

		if len(data) > layerSize {
			data = data[layerSize:len(data)]
		}

		layerArr = append(layerArr, layerString)
	}

	return layerArr
}

func part1(layerData []string) {
	digitLayers := [][]int{}
	lowestCount := 0

	for i, s := range layerData {
		digits := countDigits(s)

		currentDigits := digits[:len(digits)]
		digitLayers = append(digitLayers, currentDigits)

		if digits[0] < digitLayers[lowestCount][0] {
			lowestCount = i
		}
	}

	ans := digitLayers[lowestCount][1] * digitLayers[lowestCount][2]
	log.Print(ans)
}

func part2(layerData []string, layerSize int) {
	combinedLayer := []int64{}
	layerCount := len(layerData)

	for i := 0; i < layerSize; i++ {
		for j := 0; j < layerCount; j++ {
			num := layerData[j][i : i+1]
			currentPixel, _ := strconv.ParseInt(num, 10, 64)

			if currentPixel == 2 {
				continue
			}

			combinedLayer = append(combinedLayer, currentPixel)
			break
		}
	}

	log.Print(combinedLayer)
}

func main() {
	inputData := fileutils.ReadFile("./input.txt")

	layerWidth := 25
	layerHeight := 6
	layerData := parseImageData(inputData[0], layerWidth, layerHeight)

	// part1(layerData)
	part2(layerData, layerWidth*layerHeight)
}
