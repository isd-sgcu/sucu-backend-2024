package utils

import (
	"bytes"
	"fmt"
	"io"
)

func ToBytesReader(file io.Reader) (*bytes.Reader, error) {
	// read all data in io.Reader
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file, %v", err)
	}

	// convert to *bytes.Reader
	buffer := bytes.NewReader(data)

	return buffer, nil
}
