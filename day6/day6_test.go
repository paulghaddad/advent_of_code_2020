package main

import (
	"fmt"
	"testing"
)

func TestCalcUniqueResponses(t *testing.T) {
	var tests = []struct {
		input []string
		want  int
	}{
		{[]string{"abc"}, 3},
		{[]string{"a", "b", "c"}, 3},
		{[]string{"ab", "ac"}, 3},
		{[]string{"a", "a", "a", "a"}, 1},
		{[]string{"b"}, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.input)

		t.Run(testname, func(t *testing.T) {
			got := CalcUniqueResponses(tt.input)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestCalcCommonResponses(t *testing.T) {
	var tests = []struct {
		input []string
		want  int
	}{
		{[]string{"abc"}, 3},
		{[]string{"a", "b", "c"}, 0},
		{[]string{"ab", "ac"}, 1},
		{[]string{"a", "a", "a", "a"}, 1},
		{[]string{"b"}, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.input)

		t.Run(testname, func(t *testing.T) {
			got := CalcCommonResponses(tt.input)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
