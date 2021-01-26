package main

import (
	"fmt"
	"testing"
)

func TestIsValidPasswordSledRental(t *testing.T) {
	var tests = []struct {
		password, letter string
		min, max         int
		want             bool
	}{
		{"abcde", "a", 1, 3, true},
		{"cdefg", "b", 1, 3, false},
		{"ccccccccc", "c", 2, 9, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Password %s, letter %s, min: %d, max: %d\n", tt.password, tt.letter, tt.min, tt.max)

		t.Run(testname, func(t *testing.T) {
			got := isValidPasswordSledRental(tt.password, tt.letter, tt.min, tt.max)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestIsValidPasswordToboggan(t *testing.T) {
	var tests = []struct {
		password, letter []byte
		min, max         int
		want             bool
	}{
		{[]byte("abcde"), []byte("a"), 1, 3, true},
		{[]byte("cdefg"), []byte("b"), 1, 3, false},
		{[]byte("ccccccccc"), []byte("c"), 2, 9, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Password %s, letter %s, min: %d, max: %d\n", tt.password, tt.letter, tt.min, tt.max)

		t.Run(testname, func(t *testing.T) {
			got := isValidPasswordToboggan(tt.password, tt.letter, tt.min, tt.max)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
