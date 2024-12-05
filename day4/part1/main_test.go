package main

import "testing"

func TestFindXMAS(t *testing.T) {
	t.Run("TestFindXMAS", func(t *testing.T) {
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
		want := 18
		got := FindXMAS(input)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestTransposeDiagonal(t *testing.T) {
	t.Run("TestTransposedDiagonal", func(t *testing.T) {
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

		want := []string{
			"SAMM",
			"XMXSX",
			"XXSAMX",
			"MMXMAXS",
			"ASMASAMS",
			"SMASAMSAM",
			"MSAMMMMXAM",
			"AMSXXSAMX",
			"MMAXAMMM",
			"XMASAMX",
			"MMXSXA",
			"ASAMX",
			"SAMM",
		}
		got := transposeDiagonal(MapInputMatrix(input))

		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("got %s, want %s", got[i], want[i])
			}
		}
	})
}
