package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func generateArrayFromFile(path string) (startingString string, bytestreamArray []string, err error, elapsed time.Duration) {
	start := time.Now()
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	byteStreamArray := make([]string, 0)

	for _, line := range lines {
		byteStreamArray = append(byteStreamArray, line)
	}

	for i, value := range byteStreamArray {
		if value == "" {
			byteStreamArray = remove(byteStreamArray, i)
		}
	}

	startingString, byteStreamArray = byteStreamArray[0], byteStreamArray[1:]

	return startingString, byteStreamArray, nil, time.Since(start)
}

func main() {
	startingString, byteStreamArray, err, _ := generateArrayFromFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	currentString := startingString
	matches := 0

	fmt.Println("Starting String:", startingString)

	for steps := 0; steps < 10; steps++ {
		var tmpString = currentString
		for _, value := range byteStreamArray {
			for i := 0; i < len(currentString)-1; i++ {
				currentPolymer := strings.Split(value, " -> ")
				// fmt.Println(currentPolymer[0], currentString[i:i+2])

				if currentPolymer[0] == currentString[i:i+2] {
					matches++
					fmt.Print("\033[H\033[2J")
					fmt.Println("String Length:", len(currentString))
					fmt.Println("Step:", steps+1)
					fmt.Println("Match Number:", matches)
					//fmt.Println(currentString[i:i+1] + currentPolymer[1] + currentString[i+1:i+2])
					tmpString = strings.Replace(
						tmpString,
						currentString[i:i+2],
						currentString[i:i+1]+currentPolymer[1]+currentString[i+1:i+2],
						1)
				}
			}
		}
		currentString = tmpString
	}
	fmt.Print("\n")
	fmt.Println("Final String Length:", len(currentString))
	fmt.Print("\n")

	charArray := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	lengthsArray := make([]int, 0, len(charArray))
	for _, char := range charArray {
		lengthsArray = append(lengthsArray, len(strings.Split(currentString, char)))
	}

	var highest int
	var highestI int
	var lowest int
	var lowestI int
	for index, value := range lengthsArray {
		if value > highest {
			highest = value
			highestI = index
		}
		if value != 1 {
			if value < lowest || lowest == 0 {
				lowest = value
				lowestI = index
			}
		}
	}

	fmt.Println("Most Common Letter:", charArray[highestI], "with a length of", highest)
	fmt.Println("Least Common Letter:", charArray[lowestI], "with a length of", lowest)
	fmt.Print("\n")
	fmt.Println("Puzzle One Answer:", highest-lowest)
}
