package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindXmasAndCount(input []string, prevNextPair map[string]string, startChar string) int {
	horizontal := 0
	for _, line := range input {
		stack := []string{}
		for _, char := range strings.Split(line, "") {
			if char == startChar {
				stack = []string{}
				stack = append(stack, char)
			} else if len(stack) > 0 {
				if char == prevNextPair[stack[len(stack)-1]] {
					stack = append(stack, char)
					if len(stack) == 4 {
						horizontal++
						stack = []string{}
					}
				} else {
					stack = []string{}
				}
			}
		}
	}
	return horizontal
}

func transposeVertical(input [][]string) []string {
	columns := len(input[0])
	var transposed []string
	for i := 0; i < columns; i++ {
		s := ""
		for j := 0; j < len(input); j++ {
			s += input[j][i]
		}
		transposed = append(transposed, s)
		s = ""
	}
	return transposed
}

func transposeDiagonal(input [][]string) []string {
	columns := len(input[0])
	rows := len(input)
	transposed := []string{}

	for i := 0; i < columns*rows; i++ {
		s := ""
		for j := 0; j < rows; j++ {
			for k := 0; k < columns; k++ {
				if j+k == i {
					s += input[j][k]
				}
			}
		}
		if len(s) >= 4 {
			transposed = append(transposed, s)
		}
	}

	return transposed
}

func MapInputMatrix(input []string) [][]string {
	inputMatrix := [][]string{}
	for _, line := range input {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
	}
	return inputMatrix
}

func MapInputMatrixReverse(input []string) [][]string {
	inputMatrix := [][]string{}
	for _, line := range input {
		splitted := strings.Split(line, "")
		reversed := []string{}
		for i := len(splitted) - 1; i >= 0; i-- {
			reversed = append(reversed, splitted[i])
		}
		inputMatrix = append(inputMatrix, reversed)
	}
	return inputMatrix
}

func FindXMAS(input []string) int {
	xmasMap := map[string]string{
		"X": "M",
		"M": "A",
		"A": "S",
	}
	xmasBackwardMap := map[string]string{
		"S": "A",
		"A": "M",
		"M": "X",
	}
	horizontal := FindXmasAndCount(input, xmasMap, "X")
	backwardHorizontal := FindXmasAndCount(input, xmasBackwardMap, "S")

	inputMatrix := MapInputMatrix(input)
	verticalTransposed := transposeVertical(inputMatrix)
	vertical := FindXmasAndCount(verticalTransposed, xmasMap, "X")
	verticalBackward := FindXmasAndCount(verticalTransposed, xmasBackwardMap, "S")
	leftBottomToRightTopDiagonalTransposed := transposeDiagonal(inputMatrix)
	leftBottomToRightTopDiagonal := FindXmasAndCount(
		leftBottomToRightTopDiagonalTransposed,
		xmasMap,
		"X",
	)
	leftBottomToRightTopDiagonalBackward := FindXmasAndCount(
		leftBottomToRightTopDiagonalTransposed,
		xmasBackwardMap,
		"S",
	)
	reverseInputMatrix := MapInputMatrixReverse(input)
	rightBottomToLeftTopDiagonalTransposed := transposeDiagonal(reverseInputMatrix)
	rightBottomToLeftTopDiagonal := FindXmasAndCount(
		rightBottomToLeftTopDiagonalTransposed,
		xmasMap,
		"X",
	)
	rightBottomToLeftTopDiagonalBackward := FindXmasAndCount(
		rightBottomToLeftTopDiagonalTransposed,
		xmasBackwardMap,
		"S",
	)

	return horizontal + backwardHorizontal + vertical + verticalBackward + leftBottomToRightTopDiagonal + leftBottomToRightTopDiagonalBackward + rightBottomToLeftTopDiagonal + rightBottomToLeftTopDiagonalBackward
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println(FindXMAS(input))
}
