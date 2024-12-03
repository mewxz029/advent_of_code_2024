package pkg

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseFileToIntSlices(filePath string) ([][]int, error) {
	currentDir, _ := os.Getwd()
	realFilePath := currentDir + "/" + filePath
	file, err := os.Open(realFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var numbers []int
		if len(parts) == 0 {
			break
		}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
		result = append(result, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
