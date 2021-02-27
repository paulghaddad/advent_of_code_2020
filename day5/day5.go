package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Subproblem: determine column from pass
// Calculate seat id and track highest

// CalcRow calculates the row number from the boarding pass
func CalcRow(boardingPass string) (int, error) {
	low, high := 0, 127

	for _, marker := range boardingPass[0:7] {
		if string(marker) == "F" {
			high = (high-low)/2 + low
		} else if string(marker) == "B" {
			low = (high-low)/2 + low + 1
		} else {
			return -1, errors.New("Invalid Boarding Pass Character; must be B or F")
		}
	}

	return low, nil
}

// CalcColumn calculates the column number from the boarding pass
func CalcColumn(boardingPass string) (int, error) {
	low, high := 0, 7

	for _, marker := range boardingPass[7:] {
		if string(marker) == "L" {
			high = (high-low)/2 + low
		} else if string(marker) == "R" {
			low = (high-low)/2 + low + 1
		} else {
			return -1, errors.New("Invalid Boarding Pass Character; must be L or R")
		}
	}

	return low, nil
}

// CalcSeatID calculates a boarding pass's seat id
func CalcSeatID(boardingPass string) int {
	row, err := CalcRow(boardingPass)
	if err != nil {
		return -1
	}

	column, err := CalcColumn(boardingPass)
	if err != nil {
		return -1
	}

	return 8*row + column
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	highestSeatID := 0
	var idBitmap [1024]bool

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		seatID := CalcSeatID(scanner.Text())
		idBitmap[seatID] = true

		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading line")
	}

	// Part 1
	fmt.Printf("The highest seat ID is %d\n", highestSeatID)

	// Part 2
	for i := 9; i <= 1015; i++ {
		if idBitmap[i] == false && idBitmap[i-1] == true && idBitmap[i+1] == true {
			fmt.Printf("The missing seat ID is %v %d\n", idBitmap[i], i)
			break
		}
	}
}
