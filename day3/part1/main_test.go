package main

import "testing"

func TestMatchPattern(t *testing.T) {
	t.Run("Test matchPattern", func(t *testing.T) {
		text := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
		pattern := `mul[(](\d+),(\d+)[)]`
		matches := matchPattern(text, pattern)

		if len(matches) != 4 {
			t.Errorf("Expected 4 matches, got %d", len(matches))
		}

		expected := [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}

		for i, match := range matches {
			if match[0] != expected[i][0] || match[1] != expected[i][1] {
				t.Errorf("Expected %v, got %v", expected[i], match)
			}
		}
	})
}

func TestSummaryMultiplication(t *testing.T) {
	t.Run("Test summaryMultiplication", func(t *testing.T) {
		input := [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}
		sum := summaryMultiplication(input)

		if sum != 161 {
			t.Errorf("Expected 161, got %d", sum)
		}
	})
}
