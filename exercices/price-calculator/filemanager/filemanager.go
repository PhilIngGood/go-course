package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(InputFilePath, OutputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  InputFilePath,
		OutputFilePath: OutputFilePath,
	}
}

func (fm FileManager) ReadFiles() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	// File openning error handling
	if err != nil {
		return nil, errors.New("couldn't open a file")
	}

	scanner := bufio.NewScanner(file)
	fileContent := []string{}

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	// Scanner error handling
	if scanner.Err() != nil {
		file.Close()
		return nil, errors.New("couldn't scan a file")
	}

	file.Close()
	return fileContent, nil

}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("couldn't create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("couldn't encode to JSON")
	}

	file.Close()
	return nil

}
