package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
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

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

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

	defer file.Close() //deferred methods are executed LIFO

	time.Sleep(3 * time.Second) //simulate slow process

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return errors.New("failed to convert data")
	}

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
