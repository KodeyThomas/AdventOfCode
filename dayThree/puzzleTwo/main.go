package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./inputs.txt")
	defer file.Close()

	array0 := make([]string, 0)
	array1 := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		if input[0] == '0' {
			array0 = append(array0, input)
		} else {
			array1 = append(array1, input)
		}
	}

	o2, co2 := "", ""

	if len(array1) >= len(array0) {
		o2 = calO2(array1, 1)
		co2 = calCO2(array0, 1)
	} else {
		o2 = calO2(array0, 1)
		co2 = calCO2(array1, 1)
	}

	ratingO2, _ := strconv.ParseInt(o2, 2, 64)
	ratingCO2, _ := strconv.ParseInt(co2, 2, 64)

	println(ratingO2 * ratingCO2)
}

func calO2(arr []string, index int) string {

	if len(arr) == 1 {
		return arr[0]
	}

	array0, array1 := breakdown(arr, index)

	if len(array1) >= len(array0) {
		return calO2(array1, index+1)
	} else {
		return calO2(array0, index+1)
	}
}

func calCO2(arr []string, index int) string {
	if len(arr) == 1 {
		return arr[0]
	}

	array0, array1 := breakdown(arr, index)

	if len(array0) <= len(array1) {
		return calCO2(array0, index+1)
	} else {
		return calCO2(array1, index+1)
	}
}

func breakdown(arr []string, index int) ([]string, []string) {
	array0 := make([]string, 0)
	array1 := make([]string, 0)

	for _, val := range arr {
		if val[index] == '0' {
			array0 = append(array0, val)
		} else {
			array1 = append(array1, val)
		}
	}
	return array0, array1
}
