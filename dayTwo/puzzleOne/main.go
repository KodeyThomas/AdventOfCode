package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func generateArrayFromFile(path string) (byteStreamArray []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	byteStreamArray = make([]string, 0)

	for _, line := range lines {
		byteStreamArray = append(byteStreamArray, line)
	}

	return byteStreamArray, nil
}

func main() {
	array, _ := generateArrayFromFile("./input.txt")

	var horizontalPosition int
	var depth int

	for _, value := range array {
		commandArray := strings.Fields(value)
		commandValue, _ := strconv.Atoi(commandArray[1])

		switch commandArray[0] {

		case "forward":
			horizontalPosition += commandValue
			println(value, "Current Horizontal Position:", horizontalPosition, "Current Depth:", depth)
			break

		case "down":
			depth += commandValue
			println(value, "Current Horizontal Position:", horizontalPosition, "Current Depth:", depth)
			break

		case "up":
			depth -= commandValue
			println(value, "Current Horizontal Position:", horizontalPosition, "Current Depth:", depth)
			break
		}
	}

	println("Final Horizontal Position:", horizontalPosition, "Final Depth:", depth)
	println("Multiplying Horizontal Position and Depth:", horizontalPosition*depth)
}
