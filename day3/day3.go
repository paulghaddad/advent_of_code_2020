package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func CountTrees(terrain [][][]byte, xSlope, ySlope int) int {
	treeCount := 0
	xPos := 0
	yPos := 0
	terrainWidth := len(terrain[0])

	for yPos < len(terrain) {
		pos := terrain[yPos][xPos%terrainWidth]

		if bytes.Equal(pos, []byte("#")) {
			treeCount++
		}

		xPos += xSlope
		yPos++
	}

	return treeCount
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

	numTrees := CountTrees(terrain, 3, 1)
	fmt.Printf("The number of trees encountered was %d\n", numTrees)
}
