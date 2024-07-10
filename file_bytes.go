package main

import (
	"fmt"
	"os"
)



func getFileByteSize(filePath string) {
	file, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("File does not exist!", err)
		return
	}

	fmt.Println(file.Size(), filePath)
}