package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Time: O(n); Space: O(n)
func expenseReportTwoSum(items []int, goal int) []int {
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

// Time: O(n^2); Space: O(n)
func expenseReportThreeSum(items []int, goal int) []int {
	nums := make(map[int]int)
	for _, num := range items {
		nums[num] = goal - num
	}

	for num, sumNeeded := range nums {
		for _, comp1 := range items {
			for _, comp2 := range items {
				if num == comp1 || num == comp2 || comp1 == comp2 {
					break
				}

				if comp1+comp2 == sumNeeded {
					return []int{num, comp1, comp2}
				}
			}
		}
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

	fmt.Printf("The two values that add up to %d are %v\n", 2020, expenseReportTwoSum(expenseItems, 2020))
	fmt.Printf("The three values that add up to %d are %v\n", 2020, expenseReportThreeSum(expenseItems, 2020))

}
