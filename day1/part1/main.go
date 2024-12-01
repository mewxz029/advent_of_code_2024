package main

import (
	"fmt"

	"github.com/mewxz029/advent_of_code_2024/day1/pkg/util"
)

func insertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		j := i
		for j > 0 && numbers[j] < numbers[j-1] {
			numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			j--
		}
	}
	return numbers
}

func findDiffBetweenTwoSortedArrays(a, b []int) []int {
	result := make([]int, len(a))

	for i := range a {
		result[i] = abs(a[i] - b[i])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func summarizeNumbers(numbers []int) int {
	var result int

	for i := range numbers {
		result += numbers[i]
	}

	return result
}

func main() {
	a, b, err := util.ReadInputFileToArrays("input/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	a = insertionSort(a)
	b = insertionSort(b)
	diff := findDiffBetweenTwoSortedArrays(a, b)
	result := summarizeNumbers(diff)
	fmt.Println(result)
}
