package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func expenseReportSum(items []int, goal int) []int {
	visitedItems := make(map[int]bool)

	for _, item := range items {
		_, ok := visitedItems[goal-item]
		if ok {
			return []int{item, goal - item}
		}
		visitedItems[item] = true
	}

	return []int{}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	var expenseItems []int

	for scanner.Scan() {
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		expenseItems = append(expenseItems, item)
	}

	fmt.Printf("The values that add up to %d are %v\n", 2020, expenseReportSum(expenseItems, 2020))
}
