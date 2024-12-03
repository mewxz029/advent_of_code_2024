package main

import (
	"fmt"

	"github.com/mewxz029/advent_of_code_2024/day2/pkg"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafeReport(input []int) bool {
	var isIncreasing bool
	for i := 0; i < len(input)-1; i++ {
		diff := input[i] - input[i+1]
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if i == 0 {
			isIncreasing = diff > 0
			continue
		}
		if isIncreasing != (diff > 0) {
			return false
		}
	}
	return true
}

func countReportsAreSafe(reports [][]int) int {
	var count int
	for _, report := range reports {
		if isSafeReport(report) {
			count++
		}
	}
	return count
}

func main() {
	reports, err := pkg.ParseFileToIntSlices("input/input.txt")
	if err != nil {
		panic(err)
	}
	count := countReportsAreSafe(reports)
	fmt.Println(count)
}
