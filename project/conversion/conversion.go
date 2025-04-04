package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(lines []string) ([]float64, error) {
	floats := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New("error parsing float")
		}
		floats[lineIndex] = floatPrice
	}
	return floats, nil
}
