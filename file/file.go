package file

import (
	"io/ioutil"
	"os"
)

// ReadFile reads a file until an error or EOF
func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
