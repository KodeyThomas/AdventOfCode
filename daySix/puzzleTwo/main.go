package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func getRawDataFromFile(path string) (lineIntArray []int, err error, elapsed time.Duration) {
	start := time.Now()

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err, time.Since(start)
	}

	str := string(bytes)
	data := strings.Split(str, ",")

	lineIntArray = make([]int, 0)
	for _, i := range data {
		intV, err := strconv.Atoi(i)
		if err != nil {
			return nil, err, time.Since(start)
		}
		lineIntArray = append(lineIntArray, intV)
	}

	return lineIntArray, nil, time.Since(start)
}

func calculateCurrentPopulation(data []int) int {
	var total int

	for _, value := range data {
		total += value
	}

	return total
}

func makeInitialPopulation(data []int, fishArray []int) []int {
	for _, value := range data {
		switch {
		case value == 0:
			fishArray[0]++
		case value == 1:
			fishArray[1]++
		case value == 2:
			fishArray[2]++
		case value == 3:
			fishArray[3]++
		case value == 4:
			fishArray[4]++
		case value == 5:
			fishArray[5]++
		case value == 6:
			fishArray[6]++
		case value == 7:
			fishArray[7]++
		case value == 8:
			fishArray[8]++
		}
	}

	return fishArray
}

func getPopulationModel(data []int, days int) (int, time.Duration) {
	start := time.Now()
	fishArray := make([]int, 9, 9)
	addToPopulation := 0

	fishArray = makeInitialPopulation(data, fishArray)

	for i := 1; i <= days; i++ {
		fmt.Println("Starting Day", i, "- Total Population", calculateCurrentPopulation(fishArray))
		addToPopulation = 0
		tempArray := make([]int, 9, 9)

		for index, value := range fishArray {
			if index == 0 {
				addToPopulation = value
			} else {
				tempArray[index-1] = value
			}
		}
		fishArray = tempArray

		fishArray[6] = fishArray[6] + addToPopulation
		fishArray[8] = fishArray[8] + addToPopulation
	}

	return calculateCurrentPopulation(fishArray), time.Since(start)
}

func main() {
	start := time.Now()

	data, err, fileTime := getRawDataFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalFish, populationTime := getPopulationModel(data, 256)

	fmt.Println("Total Population:", totalFish)
	fmt.Print("\n")
	fmt.Println("----- Program Timing -----")
	elapsed := time.Since(start)
	fmt.Printf("Creating Raw Data Array took %s\n", fileTime)
	fmt.Printf("Running Population Model took %s\n", populationTime)
	fmt.Print("\n")
	fmt.Printf("Program Execution took %s\n", elapsed)
	fmt.Print("\n")
}
