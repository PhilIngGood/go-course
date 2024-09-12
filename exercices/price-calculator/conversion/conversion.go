package conversion

import (
	"errors"
	"strconv"
)

func SliceStringToSliceFloat64(strSlice []string) ([]float64, error) {
	floatSlice := make([]float64, len(strSlice))
	for i, val := range strSlice {
		floatData, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return nil, errors.New("failed to parse string to float")
		}

		floatSlice[i] = floatData
	}
	return floatSlice, nil

}
