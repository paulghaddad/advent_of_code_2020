package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountTrees(t *testing.T) {
	testTerrain := [][][]byte{}
	testTerrain = append(testTerrain, bytes.Split([]byte("..##......."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("#...#...#.."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#....#..#."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("..#.#...#.#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#...##..#."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("..#.##....."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#.#.#....#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#........#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("#.##...#..."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("#...##....#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#..#...#.#"), []byte("")))

	var tests = []struct {
		terrain [][][]byte
		slope   Slope
		want    int
	}{
		{testTerrain, Slope{X: 3, Y: 1}, 7},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Slope: %v\n", tt.slope)

		t.Run(testname, func(t *testing.T) {
			got := CountTrees(tt.terrain, tt.slope)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestMultipleRoutes(t *testing.T) {
	testTerrain := [][][]byte{}
	testTerrain = append(testTerrain, bytes.Split([]byte("..##......."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("#...#...#.."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#....#..#."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("..#.#...#.#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#...##..#."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("..#.##....."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#.#.#....#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#........#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("#.##...#..."), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte("#...##....#"), []byte("")))
	testTerrain = append(testTerrain, bytes.Split([]byte(".#..#...#.#"), []byte("")))

	var tests = []struct {
		terrain [][][]byte
		slopes  []Slope
		want    int
	}{
		{
			testTerrain,
			[]Slope{
				{X: 1, Y: 1},
				{X: 3, Y: 1},
				{X: 5, Y: 1},
				{X: 7, Y: 1},
				{X: 1, Y: 2},
			},
			336,
		}}

	for _, tt := range tests {
		testname := fmt.Sprintf("Slope %v\n", tt.slopes)

		t.Run(testname, func(t *testing.T) {
			got := MultipleRoutes(tt.terrain, tt.slopes)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
