package fileoperations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var Hello string = "Adele"

func UpdateBalance(value int, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetBalance(fileName string) (int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("file not found")
	}
	valueText := string(data)
	val, err := strconv.Atoi(valueText)

	if err != nil {
		return 0, errors.New("unable to convert data to int datatype")
	}

	return val, nil
}
