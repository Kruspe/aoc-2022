package shared

import (
	"errors"
	"os"
	"strings"
)

func ReadInput(pathToFile string) []string {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func GetBlankLineIndex(input []string) (int, error) {
	for index, l := range input {
		if l == "" {
			return index, nil
		}
	}
	return 0, errors.New("no separator")
}
