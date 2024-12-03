package main

import "testing"

func TestIsSafeReport(t *testing.T) {
	type testCase struct {
		name  string
		input []int
		want  bool
	}

	testCases := []testCase{
		{"safe report cause all decreasing by 1 or 2", []int{7, 6, 4, 2, 1}, true},
		{"unsafe report cause increase of 5", []int{1, 2, 7, 8, 9}, false},
		{"unsafe report cause decrease of 4", []int{9, 7, 6, 2, 1}, false},
		{"unsafe report cause both increasing and decreasing", []int{1, 3, 2, 4, 5}, false},
		{"unsafe report cause neither increasing nor decreasing", []int{8, 6, 4, 4, 1}, false},
		{"safe report cause levels are all increasing by 1, 2 or 3", []int{1, 3, 6, 7, 9}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSafeReport(tc.input)
			assertGotWantBooleanValue(t, got, tc.want)
		})
	}
}

func TestCountReportsAreSafe(t *testing.T) {
	type testCase struct {
		name  string
		input [][]int
		want  int
	}

	testCases := []testCase{
		{"all reports are safe", [][]int{
			{1, 2, 3, 4, 5},
			{1, 3, 5, 7, 9},
			{1, 2, 5, 6, 9},
			{8, 6, 4, 3, 1},
			{9, 7, 6, 3, 1},
		}, 5},
		{
			"some reports are unsafe",
			[][]int{
				{1, 2, 3, 4, 5},
				{1, 3, 5, 7, 9},
				{1, 2, 5, 6, 9},
				{8, 6, 4, 4, 1},
				{9, 7, 6, 2, 1},
			},
			3,
		},
		{
			"all reports are unsafe",
			[][]int{
				{9, 2, 3, 4, 5},
				{1, 6, 5, 7, 9},
				{1, 2, 3, 7, 9},
				{8, 6, 4, 4, 1},
				{9, 7, 6, 2, 1},
			},
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := countReportsAreSafe(tc.input)
			assertGotWantIntValue(t, got, tc.want)
		})
	}
}

func assertGotWantBooleanValue(t testing.TB, got bool, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func assertGotWantIntValue(t testing.TB, got int, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
