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
		terrain        [][][]byte
		xSlope, ySlope int
		want           int
	}{
		{testTerrain, 3, 1, 7},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("xSlope %d, ySlope: %d\n", tt.xSlope, tt.ySlope)

		t.Run(testname, func(t *testing.T) {
			got := CountTrees(tt.terrain, tt.xSlope, tt.ySlope)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
