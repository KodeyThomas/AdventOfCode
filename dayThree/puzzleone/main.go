package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	array, _ := generateArrayFromFile("./inputs.txt")

	var gammaRate string
	var epsilonRate string

	var gammaInt int
	var epsilonInt int

	diagnosticMap := make(map[int]string)

	for _, diagnosticString := range array {
		diagnosticArray := strings.Split(diagnosticString, "")

		for index, diagnosticValue := range diagnosticArray {
			diagnosticMap[index] += diagnosticValue
		}
	}

	keys := make([]int, 0, len(diagnosticMap))

	for k := range diagnosticMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range keys {
		var countZeroBit int
		var countOneBit int

		diagnosticMapValue := diagnosticMap[key]

		diagnosticArray := strings.Split(diagnosticMapValue, "")

		for _, value := range diagnosticArray {
			if value == "0" {
				countZeroBit++
				continue
			} else if value == "1" {
				countOneBit++
				continue
			}
		}

		if countZeroBit < countOneBit {
			gammaRate += "1"
			epsilonRate += "0"
		} else if countZeroBit > countOneBit {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			println("WOAH HOLD UP!!!")
		}
	}

	println("Gamma Rate:", gammaRate)
	println("Epsilon Rate:", epsilonRate)

	if i, err := strconv.ParseInt(gammaRate, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The Gamma Rate as a Decimal is", i)
		gammaInt = int(i)
	}

	if i, err := strconv.ParseInt(epsilonRate, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The Epsilon Rate as a Decimal is", i)
		epsilonInt = int(i)
	}

	finalRate := gammaInt * epsilonInt

	fmt.Println("The final rate is", finalRate)
}
