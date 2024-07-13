package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main () {
	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	opts := strings.Split(val, " ")

	var inputLen int = len(opts)
	if inputLen != 2 && inputLen != 3 {
		fmt.Println("Invalid command.!")
		return
	}

	if opts[0] != "ccwc" {
		fmt.Println("Invalid command!")
		return
	}

	var filePath string = ""
	if inputLen == 2 {
		filePath = opts[1]
	} else{
		filePath = opts[2]
	}

	var result string = ""
	if opts[1] == "-c" || inputLen == 2 {
		res, err := getFileByteSize(filePath)
		if err != nil {
			fmt.Println("Error while getting file bytes: ", err)
			return
		} else{
			result += " " + strconv.FormatInt(res, 10)
		}
	}

	if opts[1] == "-l" || inputLen == 2 {
		res, err := countLinesInFile(filePath)
		if err != nil {
			fmt.Println("Error while counting lines in file: ", err)
			return
		} else{
			result += " " + strconv.Itoa(res)
		}
	}

	if opts[1] == "-w" || inputLen == 2 {
		res, err := countWordsInFile(filePath)
		if err != nil {
			fmt.Println("Error while counting words in file: ", err)
			return
		}

		result += " " + strconv.Itoa(res)
	}

	if opts[1] == "-m" {
		res, err := countCharactersInFile(filePath)
		if err != nil {
			fmt.Println("Error while counting characters in file: ", err)
			return
		}

		result += " " + strconv.Itoa(res)
	}

	fmt.Println(result, " ", filePath)
}