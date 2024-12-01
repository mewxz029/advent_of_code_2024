package main

import (
	"testing"

	"github.com/mewxz029/advent_of_code_2024/day1/pkg/util"
)

func TestInsertionSort(t *testing.T) {
	t.Run("sorts a list of numbers", func(t *testing.T) {
		numbers := []int{3, 2, 1}
		expected := []int{1, 2, 3}
		result := insertionSort(numbers)

		util.AssertEqualArrays(t, result, expected)
	})
}

func TestFindDiffBetweenTwoSortedArrays(t *testing.T) {
	t.Run("finds the difference between two sorted arrays", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{2, 3, 4}
		expected := []int{1, 1, 1}
		result := findDiffBetweenTwoSortedArrays(a, b)

		util.AssertEqualArrays(t, result, expected)
	})
}

func TestSummarizeNumbers(t *testing.T) {
	t.Run("summarizes a list of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		expected := 6
		result := summarizeNumbers(numbers)

		if result != expected {
			t.Errorf("got %v want %v", result, expected)
		}
	})
}
