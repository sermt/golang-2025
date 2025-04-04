package file_handler

import (
	"fmt"
	"os"
)

func WriteValueToFile(value float64, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%.2f", value))
	return err
}

func GetValueFromFile(fileName string) (float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0.0, err
	}
	defer file.Close()

	var value float64
	_, err = fmt.Fscanf(file, "%f", &value)
	return value, err
}
