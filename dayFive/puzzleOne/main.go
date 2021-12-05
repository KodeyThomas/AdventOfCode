package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func createTensor(row int, column int, depth int) ([][][]int, time.Duration) {
	start := time.Now()
	buffer := make([]int, row*column*depth)

	tensor := make([][][]int, row)
	for row := range tensor {
		tensor[row] = make([][]int, column)
		for column := range tensor[row] {
			tensor[row][column] = buffer[:depth:depth]
			buffer = buffer[depth:]
		}
	}

	return tensor, time.Since(start)
}

func getRawDataFromFile(path string) (data []string, err error, elapsed time.Duration) {
	start := time.Now()
	file, err := os.Open(path)
	if err != nil {
		return nil, err, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil, time.Since(start)
}

func parseLine(line string) ([]int, error) {
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "->", ",")

	lineStringArray := strings.Split(line, ",")

	lineIntArray := make([]int, 0)
	for _, i := range lineStringArray {
		intV, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		lineIntArray = append(lineIntArray, intV)
	}

	return lineIntArray, nil
}

func minMax(data []string) (maxX int, maxY int, err error, duration time.Duration) {
	start := time.Now()
	for _, line := range data {
		parsedLine, err := parseLine(line)

		if err != nil {
			log.Fatal(err)
		}

		if parsedLine[0] > maxX {
			maxX = parsedLine[0]
		}

		if parsedLine[1] > maxY {
			maxY = parsedLine[1]
		}

		if parsedLine[2] > maxX {
			maxX = parsedLine[2]
		}

		if parsedLine[3] > maxY {
			maxY = parsedLine[3]
		}
	}

	return maxX, maxY, nil, time.Since(start)
}

func createHydrothermalTensor(data []string, maxX int, maxY int) ([][][]int, error, time.Duration, time.Duration) {
	start := time.Now()
	totalTransformations := len(data)
	fmt.Println("Tensor Depth:", totalTransformations)
	fmt.Println("Max X:", maxX, "Max Y:", maxY)

	tensor, tensorInitTime := createTensor(maxX+1, maxY+1, totalTransformations)

	tensor, err := addDataToTensor(tensor, data)
	if err != nil {
		return nil, err, tensorInitTime, time.Since(start)
	}
	return tensor, nil, tensorInitTime, time.Since(start)
}

func addDataToTensor(tensor [][][]int, data []string) ([][][]int, error) {
	for depth, line := range data {
		parsedLine, err := parseLine(line)

		if err != nil {
			return nil, err
		}

		if parsedLine[1] == parsedLine[3] { // X Transform
			x1 := parsedLine[0]
			x2 := parsedLine[2]

			if x1 < x2 {
				for i := 0; i <= x2-x1; i++ {
					tensor[x1+i][parsedLine[1]][depth] = 1
				}
			} else {
				for i := 0; i <= x1-x2; i++ {
					tensor[x1-i][parsedLine[1]][depth] = 1
				}
			}
		} else if parsedLine[0] == parsedLine[2] { // Y Transform
			y1 := parsedLine[1]
			y2 := parsedLine[3]

			if y1 < y2 {
				for i := 0; i <= y2-y1; i++ {
					tensor[parsedLine[0]][y1+i][depth] = 1
				}
			} else {
				for i := 0; i <= y1-y2; i++ {
					tensor[parsedLine[0]][y1-i][depth] = 1
				}
			}
		} else {
			// Ignoring At the moment... think i can tell what part two is
		}
	}

	return tensor, nil
}

func flattenTensor(tensor [][][]int, maxX int, maxY int) ([][]int, time.Duration) {
	start := time.Now()

	hydrothermalMatrix := make([][]int, maxX)
	for row := range hydrothermalMatrix {
		hydrothermalMatrix[row] = make([]int, maxY)
	}

	for row := 0; row < maxX; row++ {
		for column := 0; column < maxY; column++ {
			for depth := range tensor[row][column] {
				hydrothermalMatrix[row][column] += tensor[row][column][depth]
			}
		}
	}

	return hydrothermalMatrix, time.Since(start)
}

func countOverlaps(hydrothermalMatrix [][]int, maxX int, maxY int) int {
	overlaps := 0
	for row := 0; row < maxX; row++ {
		for column := 0; column < maxY; column++ {
			if hydrothermalMatrix[row][column] >= 2 {
				overlaps++
			}
		}
	}

	return overlaps
}

func printMatrix(hydrothermalMatrix [][]int, maxX int, maxY int, defintion int) {
	for row := 0; row < maxX; row += defintion {
		for column := 0; column < maxY; column += defintion {
			fmt.Print(hydrothermalMatrix[row][column], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	return
}

func main() {
	start := time.Now()

	data, err, fileTime := getRawDataFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	maxX, maxY, err, minMaxTime := minMax(data)
	if err != nil {
		log.Fatal(err)
	}

	hydrothermalTensor, err, initTensorTime, populatingTensorTime := createHydrothermalTensor(data, maxX, maxY)
	if err != nil {
		log.Fatal(err)
	}

	hydrothermalMatrix, flatteningTensorTime := flattenTensor(hydrothermalTensor, maxX, maxY)
	fmt.Print("\n")

	fmt.Println("Number of points where at least two lines overlap:", countOverlaps(hydrothermalMatrix, maxX, maxY))
	fmt.Print("\n")

	fmt.Println("----- Program Timing -----")
	elapsed := time.Since(start)
	fmt.Printf("Creating Raw Data Array took %s\n", fileTime)
	fmt.Printf("Min Max Calculation took %s\n", minMaxTime)
	fmt.Printf("Creating Tensor took %s\n", initTensorTime)
	fmt.Printf("Populating Tensor took %s\n", populatingTensorTime)
	fmt.Printf("Flattening Tensor took %s\n", flatteningTensorTime)
	fmt.Print("\n")
	fmt.Printf("Program Execution took %s\n", elapsed)
}
