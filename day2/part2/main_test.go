package main

import "testing"

func TestValidREcord(t *testing.T) {
	type testcase struct {
		name string
		rec  []int
		want bool
	}
	tests := []testcase{
		{"Safe without removing any level.", []int{7, 6, 4, 2, 1}, true},
		{"Unsafe regardless of which level is removed.", []int{1, 2, 7, 8, 9}, false},
		{"Unsafe regardless of which level is removed.", []int{9, 7, 6, 2, 1}, false},
		{"Safe by removing the second level, 3.", []int{1, 3, 2, 4, 5}, true},
		{"Safe by removing the third level, 4.", []int{8, 6, 4, 4, 1}, true},
		{"Safe without removing any level.", []int{1, 3, 6, 7, 9}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recs := generateRemoveOne(tt.rec)
			got := false
			for _, rec := range recs {
				if validRecord(rec) {
					got = true
					break
				}
			}
			if got != tt.want {
				t.Errorf("validRecord(%v) = %v; want %v", tt.rec, got, tt.want)
			}
		})
	}
}

func TestGenerateRemoveOne(t *testing.T) {
	type testcase struct {
		name string
		rec  []int
		want [][]int
	}
	tests := []testcase{
		{"Empty slice", []int{}, [][]int{}},
		{"Three element slice", []int{1, 2, 3}, [][]int{{2, 3}, {1, 3}, {1, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRemoveOne(tt.rec)
			if len(got) != len(tt.want) {
				t.Errorf("generateRemoveOne(%v) = %v; want %v", tt.rec, got, tt.want)
			}
			for i, rec := range got {
				if len(rec) != len(tt.want[i]) {
					t.Errorf("generateRemoveOne(%v) = %v; want %v", tt.rec, got, tt.want)
				}
				for j, n := range rec {
					if n != tt.want[i][j] {
						t.Errorf("generateRemoveOne(%v) = %v; want %v", tt.rec, got, tt.want)
					}
				}
			}
		})
	}
}
