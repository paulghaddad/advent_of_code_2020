package main

import (
	"fmt"
	"testing"
)

func TestCalcRow(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"FBFBBFFRLR", 44},
		{"BFFFBBFRRR", 70},
		{"FFFBBBFRRR", 14},
		{"BBFFBBFRLL", 102},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.input)

		t.Run(testname, func(t *testing.T) {
			got, _ := CalcRow(tt.input)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestCalcColumn(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"FBFBBFFRLR", 5},
		{"BFFFBBFRRR", 7},
		{"FFFBBBFRRR", 7},
		{"BBFFBBFRLL", 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.input)

		t.Run(testname, func(t *testing.T) {
			got, _ := CalcColumn(tt.input)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestCalcSeatID(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.input)

		t.Run(testname, func(t *testing.T) {
			got := CalcSeatID(tt.input)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
