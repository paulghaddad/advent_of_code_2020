package main

import (
	"fmt"
	"testing"
)

func TestIsValidPassport(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"yr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", false},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", true},
		{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.input)

		t.Run(testname, func(t *testing.T) {
			got := IsValidPassport(tt.input)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
