package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// Slope defines the change in the x and y coordinates
type Slope struct {
	X int
	Y int
}

// CountTrees determines the number of trees encountered on its path from the
// top-left of the terrain to the bottom
func CountTrees(terrain [][][]byte, slope Slope) int {
	treeCount := 0
	xPos := 0
	yPos := 0
	terrainWidth := len(terrain[0])

	for yPos < len(terrain) {
		pos := terrain[yPos][xPos%terrainWidth]

		if bytes.Equal(pos, []byte("#")) {
			treeCount++
		}

		xPos += slope.X
		yPos += slope.Y
	}

	return treeCount
}

// MultipleRoutes returns the product of the number of trees encountered for
// for all routes taken
func MultipleRoutes(terrain [][][]byte, slopes []Slope) int {
	productOfTrees := 1

	for _, slope := range slopes {
		treesEncountered := CountTrees(terrain, slope)
		productOfTrees *= treesEncountered
	}

	return productOfTrees
}

func main() {
	input := os.Args[1]
	f, err := os.Open(input)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	terrain := [][][]byte{}

	for scanner.Scan() {
		line := []byte(scanner.Text())

		chars := bytes.Split(line, []byte(""))
		terrain = append(terrain, chars)
	}

	// Problem A

	numTrees := CountTrees(terrain, Slope{X: 3, Y: 1})
	fmt.Printf("Problem A: The number of trees encountered was %d\n", numTrees)

	// Problem B

	product := 1
	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product = MultipleRoutes(terrain, slopes)

	fmt.Printf("Problem B: The product of the trees encountered is: %d\n", product)
}
