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

	var previousNumber int
	var totalNumberOfIncreases int

	for i, value := range array {
		if i == 0 {
			println(value, "(N/A - no previous measurement)")
			previousNumber = value
		} else {
			if previousNumber < value {
				totalNumberOfIncreases++
				println(value, "(increased)")
				previousNumber = value
			} else if previousNumber > value {
				println(value, "(decreased)")
				previousNumber = value
			} else {
				println(value, "(value stayed the same)")
				previousNumber = value
			}
		}
	}

	println("Total Number of Increases:", totalNumberOfIncreases)
}
