package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		errMsg := fmt.Sprintf("failed to open: %v", fm.InputFilePath)
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
		errMsg := fmt.Sprintf("failed to read content: %v", fm.InputFilePath)
		return nil, errors.New(errMsg)
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to write file")
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return errors.New("failed to convert data")
	}

	file.Close()

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
