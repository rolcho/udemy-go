package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		errMsg := fmt.Sprintf("failed to open: %v", fileName)
		return nil, errors.New(errMsg)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	err = scanner.Err()

	if err != nil {
		errMsg := fmt.Sprintf("failed to read content: %v", fileName)
		return nil, errors.New(errMsg)
	}

	return lines, nil
}
