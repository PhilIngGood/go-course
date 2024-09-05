package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueString := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueString), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	dataByte, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}
	valueText := string(dataByte)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse stored values")
	}
	return value, nil
}
