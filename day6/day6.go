package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CalcUniqueResponses determines the number of unique letters in each response
func CalcUniqueResponses(responses []string) int {
	letters := createLetterBitmap(responses)

	return uniqLetterCount(letters)
}

// CalcCommonResponses returns the number of questions everyone answered
func CalcCommonResponses(responses []string) int {
	commonResponses := createLetterBitmap(responses[0:1])

	for _, response := range responses[1:] {
		for i, pres := range commonResponses {
			if pres == 1 && !strings.ContainsRune(response, rune(i+int('a'))) {
				commonResponses[i] = 0
			}
		}
	}

	return uniqLetterCount(commonResponses)
}

func createLetterBitmap(lines []string) [26]int {
	letters := [26]int{}

	for _, line := range lines {
		for _, letter := range line {
			letters[int(letter)-int('a')] = 1
		}
	}

	return letters
}

func uniqLetterCount(letterBitmap [26]int) int {
	numUniq := 0
	for _, pres := range letterBitmap {
		if pres == 1 {
			numUniq++
		}
	}

	return numUniq
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	totalUniqueAnswers := 0
	totalCommonAnswers := 0
	group := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Compare(line, "") == 0 {
			totalUniqueAnswers += CalcUniqueResponses(group)
			totalCommonAnswers += CalcCommonResponses(group)
			group = []string{}
		} else {
			group = append(group, line)
		}
	}

	totalUniqueAnswers += CalcUniqueResponses(group)
	totalCommonAnswers += CalcCommonResponses(group)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard error", err)
	}

	// Part 1
	fmt.Printf("The sum of all unique responses is: %d\n", totalUniqueAnswers)

	// Part 2
	fmt.Printf("The sum of all common responses is: %d\n", totalCommonAnswers)
}
