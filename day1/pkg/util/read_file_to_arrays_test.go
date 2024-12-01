package util

import (
	"testing"
)

func TestReadInputFileToArrays(t *testing.T) {
	t.Run("reads input file to arrays", func(t *testing.T) {
		expectedA := []int{27636, 92436, 68957}
		expectedB := []int{67663, 51410, 77912}
		a, b, _ := ReadInputFileToArrays("../../input/input_test.txt")

		AssertEqualArrays(t, a, expectedA)
		AssertEqualArrays(t, b, expectedB)
	})
}
