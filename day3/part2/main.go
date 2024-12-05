package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func matchPattern(text string, pattern string) [][]string {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(text, -1)
	return matches
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	do := true
	sum := 0
	fmt.Println("text", len(lines))
	for _, t := range lines {

		matches := matchPattern(t, `mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don\'t\(\))`)
		// fmt.Println(matches)
		for _, match := range matches {
			if match[1] != "" {
				left, _ := strconv.Atoi(match[1])
				right, _ := strconv.Atoi(match[2])
				if do {
					sum += left * right
				}
			} else if match[3] != "" {
				do = true
			} else if match[4] != "" {
				do = false
			}
		}
	}

	fmt.Printf("Sum of all products: %d\n", sum)
}
