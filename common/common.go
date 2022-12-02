package common

import (
	"io"
	"os"
)

func ReadInput(path string) (input string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return
	}
	return string(bytes), nil
}
