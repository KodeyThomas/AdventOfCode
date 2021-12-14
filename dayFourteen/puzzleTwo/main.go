package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func parseInput(path string) (string, polymerMap) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil
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

	startingString, byteStreamArray := byteStreamArray[0], byteStreamArray[1:]

	return startingString, parsePolymerMap(byteStreamArray)
}

type polymerData struct {
	polymerStart string
	polymerEnd   string
}

type polymerMap map[polymerData]string
type pairMap map[polymerData]int

func initializePairs(startingString string) pairMap {
	pairMapInstance := make(pairMap)
	reader := strings.NewReader(startingString)
	polymerStart, _, _ := reader.ReadRune()
	for {
		polymerEnd, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		pairMapInstance[polymerData{string(polymerStart), string(polymerEnd)}]++
		polymerStart = polymerEnd
	}
	return pairMapInstance
}

func parsePolymerMap(lines []string) polymerMap {
	polymerMapInstance := make(polymerMap)
	regex := regexp.MustCompile(`([A-Z])([A-Z]) -> ([A-Z])`)
	for i := 0; i < len(lines); i++ {
		matches := regex.FindStringSubmatch(lines[i])
		polymerMapInstance[polymerData{matches[1], matches[2]}] = matches[3]
	}
	return polymerMapInstance
}

func main() {
	startingString, polymerMapInstance := parseInput("./input.txt")
	pairMapInstance := generatePairMap(startingString, polymerMapInstance)

	totalPairsMap := countPairs(pairMapInstance, string(startingString[0]))
	maxValue := 0
	minValue := 0
	for _, pairValue := range totalPairsMap {
		if pairValue > maxValue {
			maxValue = pairValue
		}
		if pairValue < minValue || minValue == 0 {
			minValue = pairValue
		}
	}
	fmt.Println("Puzzle Two Answer:", maxValue-minValue)
}

func countPairs(pairMapInstance pairMap, first string) map[string]int {
	totalPairs := make(map[string]int)
	for currentPair, value := range pairMapInstance {
		totalPairs[currentPair.polymerEnd] += value
	}

	totalPairs[first]++
	return totalPairs
}

func generatePairMap(startingString string, polymerMap polymerMap) pairMap {
	pairInstance := initializePairs(startingString)
	for i := 0; i < 40; i++ {
		fmt.Println("Polymer Length: ", getLengthOfPolymer(pairInstance, startingString))
		newPairInstance := copyMapInstance(pairInstance)
		for currentPair, addedPolymer := range polymerMap {
			pairValue := pairInstance[currentPair]
			if pairValue > 0 {
				newPairInstance[currentPair] -= pairValue
				newPairInstance[polymerData{currentPair.polymerStart, addedPolymer}] += pairValue
				newPairInstance[polymerData{addedPolymer, currentPair.polymerEnd}] += pairValue
			}
		}

		pairInstance = newPairInstance
	}
	return pairInstance
}

func copyMapInstance(originalMap pairMap) pairMap {
	newMapInstance := make(pairMap)
	for key, value := range originalMap {
		newMapInstance[key] = value
	}
	return newMapInstance
}

func getLengthOfPolymer(pairInstance pairMap, startingString string) int {
	totalPairsMap := countPairs(pairInstance, string(startingString[0]))
	total := 0
	for _, pairValue := range totalPairsMap {
		total += pairValue
	}
	return total
}
