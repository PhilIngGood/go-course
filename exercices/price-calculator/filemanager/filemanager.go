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
	// Go will not execute this close() now, it will wait for the Readfile function to end, then close the file
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileContent := []string{}

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	// Scanner error handling
	if scanner.Err() != nil {
		return nil, errors.New("couldn't scan a file")
	}

	return fileContent, nil

}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("couldn't create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("couldn't encode to JSON")
	}

	return nil

}
