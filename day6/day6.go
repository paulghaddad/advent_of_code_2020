package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CalcUniqueResponses determines the number of unique letters in each response
func CalcUniqueResponses(responses []string) int {
	letters := [26]int{}

	for _, response := range responses {
		for _, letter := range response {
			letters[int(letter)-int('a')] = 1
		}
	}

	numUniq := 0
	for _, pres := range letters {
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

	totalAnswers := 0
	group := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Compare(line, "") == 0 {
			totalAnswers += CalcUniqueResponses(group)
			group = []string{}
		}

		group = append(group, line)
	}

	totalAnswers += CalcUniqueResponses(group)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard error", err)
	}

	// Part 1
	fmt.Printf("The sum of all unique responses is: %d\n", totalAnswers)
}
