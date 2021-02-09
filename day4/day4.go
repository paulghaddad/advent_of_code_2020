package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))

	var buf bytes.Buffer
	validPassports := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Compare(line, "") == 0 {
			if IsValidPassport(buf.String()) {
				validPassports++
			}
			buf.Truncate(0)
		}

		buf.WriteString(line)
		buf.WriteByte(' ')
	}

	fmt.Printf("There are %d valid passports.\n", validPassports)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input")
	}
}
