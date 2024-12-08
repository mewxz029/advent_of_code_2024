package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindXmas2(input []string) int {
	matrix := [][]string{}
	for _, line := range input {
		matrix = append(matrix, strings.Split(line, ""))
	}

	aPosition := [][]int{}

	for i, line := range matrix {
		if i == 0 || i == len(matrix)-1 {
			continue
		}
		for j, char := range line {
			if j == 0 || j == len(line)-1 {
				continue
			}
			if char == "A" {
				aPosition = append(aPosition, []int{i, j})
			}
		}
	}

	count := 0

	for _, pos := range aPosition {
		leftTop := matrix[pos[0]-1][pos[1]-1]
		rightBottom := matrix[pos[0]+1][pos[1]+1]
		correctLtRb := (leftTop == "M" && rightBottom == "S") ||
			(leftTop == "S" && rightBottom == "M")
		rightTop := matrix[pos[0]-1][pos[1]+1]
		leftBottom := matrix[pos[0]+1][pos[1]-1]
		correctRtLb := (rightTop == "M" && leftBottom == "S") ||
			(rightTop == "S" && leftBottom == "M")

		if correctLtRb && correctRtLb {
			count++
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println(FindXmas2(input))
}
