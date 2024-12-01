package main

import (
	"testing"

	"github.com/mewxz029/advent_of_code_2024/day1/pkg/util"
)

func TestConvertSliceToMap(t *testing.T) {
	t.Run("convert slice to map", func(t *testing.T) {
		slice := []int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5,
		}
		expected := map[int]int{
			1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0,
		}
		result := convertSliceToMap(slice)

		if len(result) != len(expected) {
			t.Errorf("got %v want %v", len(result), len(expected))
		}

		for k, v := range result {
			util.AssertEqual(t, v, expected[k])
		}
	})
}

func TestFindSimilarityScoreMap(t *testing.T) {
	t.Run("finds the similarity score", func(t *testing.T) {
		a := []int{1, 2, 3, 1, 2}
		b := []int{2, 3, 4, 2, 3}
		expected := map[int]int{1: 0, 2: 2, 3: 2}
		result := findSimilarityScoreMap(a, b)

		if len(result) != len(expected) {
			t.Errorf("got %v want %v", len(result), len(expected))
		}

		for k, v := range result {
			util.AssertEqual(t, v, expected[k])
		}
	})
}

func TestSummarizeSimilarityScore(t *testing.T) {
	t.Run("summarizes the similarity score", func(t *testing.T) {
		score := map[int]int{1: 0, 2: 2, 3: 2}
		expected := 10
		result := summarizeSimilarityScore(score)

		util.AssertEqual(t, result, expected)
	})
}
