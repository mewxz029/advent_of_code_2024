package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func matchPattern(text string, pattern string) [][]int {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(text, -1)
	var results [][]int
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		results = append(results, []int{left, right})
	}
	return results
}

func summaryMultiplication(input [][]int) int {
	sum := 0
	for _, x := range input {
		sum += x[0] * x[1]
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	text := strings.Fields(line)[0]
	fmt.Println(text)

	matches := matchPattern(text, `mul[(](\d+),(\d+)[)]`)
	sum := summaryMultiplication(matches)

	fmt.Printf("Sum of all products: %d\n", sum)
}
