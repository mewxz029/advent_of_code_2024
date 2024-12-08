package main

import "testing"

func TestFindXmas2(t *testing.T) {
	t.Run("TestFindXmas2", func(t *testing.T) {
		input := []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		}
		want := 9
		got := FindXmas2(input)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
