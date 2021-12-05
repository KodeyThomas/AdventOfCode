package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createTensor(row int, column int, depth int) [][][]int {
	buffer := make([]int, row*column*depth)

	tensor := make([][][]int, row)
	for row := range tensor {
		tensor[row] = make([][]int, column)
		for column := range tensor[row] {
			tensor[row][column] = buffer[:depth:depth]
			buffer = buffer[depth:]
		}
	}

	return tensor
}

func extractDataFromFile(path string) (bingoNumbers []int, bingoTensor [][][]int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	bingoNumbers = make([]int, 0)

	for i := 0; i < 1; i++ {
		bingoNumberLineString := lines[i]

		bingoValues := strings.Split(bingoNumberLineString, ",")

		for _, value := range bingoValues {
			number, _ := strconv.Atoi(value)
			bingoNumbers = append(bingoNumbers, number)
		}
	}

	bingoTensor = createTensor(6, 6, 100)
	currentRow := -1
	currentDepth := -1

	for i := 1; i < len(lines); i++ {
		currentRow++

		currentLine := strings.Trim(lines[i], " ")
		currentLine = strings.ReplaceAll(currentLine, "  ", " ")

		if len(lines[i]) == 0 {
			currentDepth++
			currentRow = -1
			continue
		}

		currentLineArray := strings.Split(currentLine, " ")

		for column, value := range currentLineArray {
			valueInt, _ := strconv.Atoi(value)
			bingoTensor[currentRow][column][currentDepth] = valueInt
		}

	}

	return bingoNumbers, bingoTensor, nil
}

func calculateScoreTensor(bingoNumbers []int, bingoTensor [][][]int) (scoreTensor [][][]int, winningValue int, tensorDepth int) {
	scoreMatrix := make([][]int, 99)
	for i := 0; i < 99; i++ {
		scoreMatrix[i] = make([]int, 1)
	}

	scoreTensor = createTensor(6, 6, 100)

	row := 0
	column := 0
	i := 0

	for bingoIndex, value := range bingoNumbers {
		fmt.Println("Calling Number:", bingoIndex)

		for row = 0; row < 5; row++ {
			for column = 0; column < 5; column++ {
				for i = 0; i < 99; i++ {
					if bingoTensor[row][column][i] == value {
						scoreTensor[row][column][i] = 1
					}
					var score int

					for column := 0; column < 6; column++ {
						if scoreTensor[row][column][i] == 1 {
							score++
						}
					}

					if score == 5 {
						fmt.Print("\n")
						println("Bingo! \nTensor Depth:", i, "\nWinning Number:", value)
						fmt.Print("\n")
						return scoreTensor, value, i
					}
				}
			}

			if bingoTensor[row][column][i] == value {
				scoreTensor[row][column][i] = 1
			}
			var score int

			for column := 0; column < 6; column++ {
				if scoreTensor[row][column][i] == 1 {
					score++
				}
			}

			if score == 5 {
				fmt.Print("\n")
				println("Bingo! \nTensor Depth:", i, "\nWinning Number:", value)
				fmt.Print("\n")
				return scoreTensor, value, i
			}
		}
	}

	return nil, 0, 0
}

func main() {
	bingoNumbers, bingoTensor, err := extractDataFromFile("input.txt")

	if err != nil {
		panic(err)
	}

	scoreTensor, winningValue, tensorDepth := calculateScoreTensor(bingoNumbers, bingoTensor)

	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			fmt.Print(bingoTensor[row][column][tensorDepth], " ")
		}
		fmt.Println()
	}

	var notCalledSum int
	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if scoreTensor[row][column][tensorDepth] == 0 {
				notCalledSum += bingoTensor[row][column][tensorDepth]
			}
		}
	}

	fmt.Println("Not Called Sum:", notCalledSum)
	fmt.Println("Puzzle One Answer:", notCalledSum*winningValue)
}
