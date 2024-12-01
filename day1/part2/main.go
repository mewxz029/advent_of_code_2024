package main

import (
	"fmt"

	"github.com/mewxz029/advent_of_code_2024/day1/pkg/util"
)

func convertSliceToMap(slice []int) map[int]int {
	m := make(map[int]int)

	for _, item := range slice {
		m[item] = 0
	}

	return m
}

func findSimilarityScoreMap(a, b []int) map[int]int {
	m := convertSliceToMap(a)

	for _, item := range b {
		_, ok := m[item]
		if ok {
			m[item]++
		}
	}

	return m
}

func summarizeSimilarityScore(score map[int]int) int {
	sum := 0

	for k, v := range score {
		sum += k * v
	}

	return sum
}

func main() {
	a, b, err := util.ReadInputFileToArrays("input/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	score := findSimilarityScoreMap(a, b)
	result := summarizeSimilarityScore(score)
	fmt.Println(result)
}
