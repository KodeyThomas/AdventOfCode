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
	var aim int

	for _, value := range array {
		commandArray := strings.Fields(value)
		x, _ := strconv.Atoi(commandArray[1])

		switch commandArray[0] {

		case "forward":
			horizontalPosition += x
			depth += aim * x
			println(value, "Current Horizontal Position:", horizontalPosition, "Current Depth:", depth, "Current Aim:", aim)
			break

		case "down":
			aim += x
			println(value, "Current Horizontal Position:", horizontalPosition, "Current Depth:", depth, "Current Aim:", aim)
			break

		case "up":
			aim -= x
			println(value, "Current Horizontal Position:", horizontalPosition, "Current Depth:", depth, "Current Aim:", aim)
			break
		}
	}

	println("Final Horizontal Position:", horizontalPosition, "Final Depth:", depth, "Final Aim:", aim)
	println("Multiplying Horizontal Position and Depth:", horizontalPosition*depth)
}
