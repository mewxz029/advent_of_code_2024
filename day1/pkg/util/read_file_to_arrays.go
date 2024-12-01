package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputFileToArrays(filename string) ([]int, []int, error) {
	var a, b []int

	currentPath, _ := os.Getwd()
	path := currentPath + "/" + filename
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("error converting numbers: %v, %v", err1, err2)
		}

		a = append(a, num1)
		b = append(b, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return a, b, nil
}
