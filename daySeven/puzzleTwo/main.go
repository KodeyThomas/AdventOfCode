package main

import (
	"fmt"
	"io/ioutil"
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

func minMax(lineIntArray []int) (min, max int, timeTaken time.Duration) {
	start := time.Now()
	for _, value := range lineIntArray {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	timeTaken = time.Since(start)

	return min, max, timeTaken
}

func calculateCheapestFuel(lineIntArray []int, min, max int) (int, time.Duration) {
	start := time.Now()
	var cheapestFuel int
	for minMaxValue := min; minMaxValue <= max; minMaxValue++ {
		var tempFuelCost int
		for _, value := range lineIntArray {
			if value < minMaxValue {
				diff := minMaxValue - value
				for x := 0; x < diff; x++ {
					tempFuelCost += x + 1
				}
			} else if value > minMaxValue {
				diff := value - minMaxValue
				for x := 0; x < diff; x++ {
					tempFuelCost += x + 1
				}
			} else if value == minMaxValue {
				tempFuelCost += 0
			}
		}

		if tempFuelCost < cheapestFuel || cheapestFuel == 0 {
			fmt.Println("Found New Cheapest Fuel: ", tempFuelCost)
			fmt.Println("Horizontal Position: ", minMaxValue)
			fmt.Print("\n")
			cheapestFuel = tempFuelCost
		}
	}

	return cheapestFuel, time.Since(start)
}

func main() {
	start := time.Now()

	lineIntArray, err, fileTime := getRawDataFromFile("./inputs.txt")
	if err != nil {
		panic(err)
	}

	min, max, minMaxTime := minMax(lineIntArray)
	cheapestFuelCost, fuelCalTime := calculateCheapestFuel(lineIntArray, min, max)

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Print("\n")

	fmt.Printf("Cheapest Fuel Cost: %d\n", cheapestFuelCost)
	fmt.Print("\n")

	fmt.Println("----- Program Timing -----")
	elapsed := time.Since(start)
	fmt.Printf("Creating Raw Data Array took %s\n", fileTime)
	fmt.Printf("Calculating Min and Max took %s\n", minMaxTime)
	fmt.Printf("Calculating Cheapest Fuel took %s\n", fuelCalTime)
	fmt.Print("\n")
	fmt.Printf("Program Execution took %s\n", elapsed)
	fmt.Print("\n")
}
