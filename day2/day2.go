package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValidPasswordSledRental(password, letter string, min, max int) bool {
	count := strings.Count(password, letter)

	return count >= min && count <= max
}

func isValidPasswordToboggan(password, letter []byte, min, max int) bool {
	minPresent := password[min-1] == letter[0]
	maxPresent := password[max-1] == letter[0]

	return (minPresent || maxPresent) && !(minPresent && maxPresent)
}

func main() {
	input_file := os.Args[1]
	f, err := os.Open(input_file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	re := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-zA-Z]): ([a-zA-Z]+)")

	numValidPasswordsSledRental := 0
	numValidPasswordsToboggan := 0
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)[0]
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		letter := matches[3]
		password := matches[4]

		if isValidPasswordSledRental(password, letter, min, max) {
			numValidPasswordsSledRental++
		}

		if isValidPasswordToboggan([]byte(password), []byte(letter), min, max) {
			numValidPasswordsToboggan++
		}
	}

	fmt.Printf("%d\n", numValidPasswordsSledRental)
	fmt.Printf("%d\n", numValidPasswordsToboggan)
}
