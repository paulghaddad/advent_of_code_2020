package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var requiredPassportKeys = []string{
	"ecl",
	"pid",
	"eyr",
	"hcl",
	"byr",
	"iyr",
	"hgt",
}

// IsValidPassport determines whether a passport string is valid
func IsValidPassport(input string) bool {
	passport := make(map[string]string)

	for _, kv := range strings.Split(strings.TrimSpace(input), " ") {
		pair := strings.Split(kv, ":")
		passport[pair[0]] = pair[1]
	}

	for _, reqKey := range requiredPassportKeys {
		_, pres := passport[reqKey]
		if !pres {
			return false
		}
	}

	return true
}

// IsValidPassportStrict validates the individual fields of each passport, in
// addition to the criteria in IsValidPassport
func IsValidPassportStrict(input string) bool {
	if !IsValidPassport(input) {
		return false
	}

	passport := make(map[string]string)

	for _, kv := range strings.Split(strings.TrimSpace(input), " ") {
		pair := strings.Split(kv, ":")
		passport[pair[0]] = pair[1]
	}

	parsedByr, _ := strconv.Atoi(passport["byr"])
	if !ValidBirthyear(parsedByr) {
		return false
	}

	parsedIyr, _ := strconv.Atoi(passport["iyr"])
	if !ValidIssueyear(parsedIyr) {
		return false
	}

	parsedEyr, _ := strconv.Atoi(passport["eyr"])
	if !ValidExpirationyear(parsedEyr) {
		return false
	}

	if !ValidHeight(passport["hgt"]) {
		return false
	}

	if !ValidHaircolor(passport["hcl"]) {
		return false
	}

	if !ValidEyecolor(passport["ecl"]) {
		return false
	}

	if !ValidPassportID(passport["pid"]) {
		return false
	}

	return true
}

// ValidBirthyear returns true if: four digits; at least 1920 and at most 2002.
func ValidBirthyear(year int) bool {
	return year >= 1920 && year <= 2002
}

// ValidIssueyear returns true if: four digits; at least 2010 and at most 2020.
func ValidIssueyear(year int) bool {
	return year >= 2010 && year <= 2020
}

// ValidExpirationyear returns true if: four digits; at least 2020 and at most 2030.
func ValidExpirationyear(year int) bool {
	return year >= 2020 && year <= 2030
}

// ValidHeight returns whether a given height in in or cm is within a valid
// range
func ValidHeight(height string) bool {
	r, _ := regexp.Compile("([1-9][0-9]{1,2})(in|cm)")
	match := r.MatchString(height)
	if !match {
		return false
	}

	groups := r.FindStringSubmatch(height)
	num, _ := strconv.Atoi(groups[1])

	if groups[2] == "cm" && num >= 150 && num <= 193 {
		return true
	}

	if groups[2] == "in" && num >= 59 && num <= 76 {
		return true
	}

	return false
}

// ValidHaircolor returns whether the haircolor code is valid
func ValidHaircolor(color string) bool {
	match, _ := regexp.MatchString("#[0-9a-f]{6}", color)
	return match
}

// ValidEyecolor returns whether the eye color is valid
func ValidEyecolor(color string) bool {
	validEyecolors := [7]string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	for _, validColor := range validEyecolors {
		if color == validColor {
			return true
		}
	}

	return false
}

// ValidPassportID returns whether a PID is valid
func ValidPassportID(pid string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", pid)
	return match
}

func main() {
	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))

	var buf bytes.Buffer
	validPassports := 0
	validPassportsStrict := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Compare(line, "") == 0 {
			bufString := buf.String()
			if IsValidPassport(bufString) {
				validPassports++
			}

			if IsValidPassportStrict(bufString) {
				validPassportsStrict++

			}
			buf.Truncate(0)
		}

		buf.WriteString(line)
		buf.WriteByte(' ')
	}

	fmt.Printf("There are %d valid passports and %d valid string passports.\n", validPassports, validPassportsStrict)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input")
	}
}
