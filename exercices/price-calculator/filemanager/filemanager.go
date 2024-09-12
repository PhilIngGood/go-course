package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadFiles(filename string) ([]string, error) {
	file, err := os.Open(filename)

	// File openning error handling
	if err != nil {
		return nil, errors.New("an error occured openning file")
	}

	scanner := bufio.NewScanner(file)
	fileContent := []string{}

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	// Scanner error handling
	if scanner.Err() != nil {
		file.Close()
		return nil, errors.New("an error occured scanning file")
	}

	file.Close()
	return fileContent, nil

}

func WriteJSON(data any, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return errors.New("an error occured creating file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("an error occured encoding to JSON")
	}

	file.Close()
	return nil

}
