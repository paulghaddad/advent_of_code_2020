package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	input_file := os.Args[1]
	f, err := os.Open(input_file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	re := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-zA-Z]): ([a-zA-Z]+)")

	for scanner.Scan() {
		fmt.Printf("%q\n", re.FindAllStringSubmatch(scanner.Text(), -1))
	}
}
