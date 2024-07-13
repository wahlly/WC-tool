package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)


func getFileByteSize(filePath string) (int64, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("File does not exist!", err)
		return 0, err
	}

	// fmt.Println(file.Size(), filePath)
	return file.Size(), nil
}

func countWordsInFile (filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		// fmt.Println("Unable to access file: ", err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	words := 0

	for scanner.Scan() {
		words++
	}

	scanError := scanner.Err()
	if scanError != nil {
		// fmt.Println("Unable to scan file! ", scanner.Err())
		return 0, scanError 
	}

	// fmt.Println(words, filePath)
	return words, nil
}

func countLinesInFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		// fmt.Println("Unable to access file! ", err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if (scanner.Err() != nil) {
		// fmt.Println("Unable to scan file! ", scanner.Err())
		return 0, nil
	}

	// fmt.Println(lineCount, filePath)
	return lineCount, nil
}

func countCharactersInFile (filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		// fmt.Println("Unable to access file! ", err)
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	char := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			// fmt.Println("Unable to read file! ", err)
			return 0, err
		}

		char += utf8.RuneCountInString(line)

		if err == io.EOF {
			break
		}
	}

	// fmt.Println(char, filePath)
	return char, nil
}