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

func TestIsValidPassportStrict(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926", false},
		{"iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946", false},
		{"hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277", false},
		{"hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007", false},
		{"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f", true},
		{"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm", true},
		{"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022", true},
		{"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.input)

		t.Run(testname, func(t *testing.T) {
			got := IsValidPassportStrict(tt.input)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestValidBirthyear(t *testing.T) {
	var tests = []struct {
		year int
		want bool
	}{
		{2002, true},
		{2003, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %d\n", tt.year)

		t.Run(testname, func(t *testing.T) {
			got := ValidBirthyear(tt.year)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestValidIssueyear(t *testing.T) {
	var tests = []struct {
		year int
		want bool
	}{
		{2014, true},
		{2009, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %d\n", tt.year)

		t.Run(testname, func(t *testing.T) {
			got := ValidIssueyear(tt.year)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestValidExpirationyear(t *testing.T) {
	var tests = []struct {
		year int
		want bool
	}{
		{2021, true},
		{2019, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %d\n", tt.year)

		t.Run(testname, func(t *testing.T) {
			got := ValidExpirationyear(tt.year)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestValidHeight(t *testing.T) {
	var tests = []struct {
		height string
		want   bool
	}{
		{"60in", true},
		{"190cm", true},
		{"190in", false},
		{"190", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.height)

		t.Run(testname, func(t *testing.T) {
			got := ValidHeight(tt.height)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestValidHaircolor(t *testing.T) {
	var tests = []struct {
		haircolor string
		want      bool
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.haircolor)

		t.Run(testname, func(t *testing.T) {
			got := ValidHaircolor(tt.haircolor)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestValidEyecolor(t *testing.T) {
	var tests = []struct {
		eyecolor string
		want     bool
	}{
		{"brn", true},
		{"wat", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.eyecolor)

		t.Run(testname, func(t *testing.T) {
			got := ValidEyecolor(tt.eyecolor)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestValidPassportID(t *testing.T) {
	var tests = []struct {
		pid  string
		want bool
	}{
		{"000000001", true},
		{"0123456789", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %s\n", tt.pid)

		t.Run(testname, func(t *testing.T) {
			got := ValidPassportID(tt.pid)

			if got != tt.want {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
