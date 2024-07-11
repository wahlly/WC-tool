package main

import (
	"bufio"
	"fmt"
	"os"
)



func countWordsInFile (filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to access file: ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	words := 0

	for scanner.Scan() {
		words++
	}

	if scanner.Err() != nil {
		fmt.Println("Unable to scan file! ", scanner.Err())
		return
	}

	fmt.Println(words, filePath)
}