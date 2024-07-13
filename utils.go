package main

import (
	"bufio"
	"io"
	"os"
	"unicode/utf8"
)


func getFileByteSize(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	byteSize := 0
	for scanner.Scan() {
		line := scanner.Text()
		//added two for newline
		byteSize += len(line) + 2
	}

	return byteSize, nil
}

func countWordsInFile (filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
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
		return 0, scanError 
	}

	return words, nil
}

func countLinesInFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if (scanner.Err() != nil) {
		return 0, nil
	}

	return lineCount, nil
}

func countCharactersInFile (filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	char := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return 0, err
		}

		char += utf8.RuneCountInString(line)

		if err == io.EOF {
			break
		}
	}

	return char, nil
}