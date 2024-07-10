package main

import (
	"bufio"
	"fmt"
	"os"
)


func countFileLines(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to access file! ", err)
		return
	}
	// defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if (scanner.Err() != nil) {
		fmt.Println("Unable to scan file! ", scanner)
	}

	fmt.Println(lineCount, filePath)
}