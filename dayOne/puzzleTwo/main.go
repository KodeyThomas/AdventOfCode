package main

import (
	"bufio"
	"os"
	"strconv"
)

func generateArrayFromFile(path string) (byteStreamArray []int, err error) {
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

	byteStreamArray = make([]int, 0)

	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		byteStreamArray = append(byteStreamArray, number)
	}

	return byteStreamArray, nil
}

func main() {
	array, _ := generateArrayFromFile("./inputs.txt")

	var previousSum int
	var currentSum int
	var totalNumberOfIncreases int

	for i, value := range array {
		if i < 3 {
			previousSum += value
			currentSum += value
			continue
		} else {
			if i == 3 {
				println(currentSum, "(N/A - no previous sum)")
			} else {
				currentSum -= array[i-3]
				currentSum += value

				if currentSum > previousSum {
					totalNumberOfIncreases++
					previousSum = currentSum
					println(currentSum, "(increased)")
				} else if currentSum < previousSum {
					previousSum = currentSum
					println(currentSum, "(decreased)")
				} else {
					previousSum = currentSum
					println(currentSum, "(no change)")
				}
			}
		}
	}

	println("Total number of increases:", totalNumberOfIncreases)
}
