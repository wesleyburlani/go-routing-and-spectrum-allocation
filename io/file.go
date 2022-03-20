package io

import (
	"fmt"
	"os"
)

func OpenFile(path string) *os.File {
	data, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file ", path, err)
		os.Exit(1)
	}
	return data
}
