package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)


func main () {
	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	opts := strings.Split(val, " ")

	var inputLen int = len(opts)

	if opts[0] != "ccwc" && opts[0] != "cat" {
		fmt.Println("Invalid command!")
		return
	}

	var filePath string = ""
	if inputLen == 3 {
		filePath = opts[2]
	} else {
		filePath = opts[1]
	}

	//decide where to read from, either file or command terminal
	var fileInput io.Reader
	if opts[0] == "cat" {
		if inputLen != 4 && inputLen != 5 {
			fmt.Println("Invalid command execution! not recognised")
			return
		}
		inputCmd := opts[0] + " " + opts[2] + " " + opts[3]
		if inputCmd != "cat | ccwc" {
			fmt.Println("Invalid command execution! Invalid use of pipes with stdin")
			return
		}
		/**execute a bash command in the terminal and capture its output*/
		cmd := exec.Command("cat", filePath)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("An error occured while executing command: ", err)
		}

		//use the commnd output as input
		fileInput = strings.NewReader(string(output))
	} else{
		if inputLen != 2 && inputLen != 3 {
			fmt.Println("Invalid command! not recognized")
			return
		}
		//read from a given filepath
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Unable to access file! ", err)
			// return 0, err
		}
		fileInput = file
		defer file.Close()
	}

	options := ""
	if inputLen == 3 {
		options = opts[1]
	} else if inputLen == 5 {
		options = opts[4]
	}
	var result string = ""
	if options == "-c" || options == "" {
		res, err := getFileByteSize(fileInput)
		if err != nil {
			fmt.Println("Error while getting file bytes: ", err)
			return
		} else{
			result += " " + strconv.Itoa(res)
		}
	}

	if options == "-l" || options == "" {
		res, err := countLinesInFile(filePath)
		if err != nil {
			fmt.Println("Error while counting lines in file: ", err)
			return
		} else{
			result += " " + strconv.Itoa(res)
		}
	}

	if options == "-w" || options == "" {
		res, err := countWordsInFile(filePath)
		if err != nil {
			fmt.Println("Error while counting words in file: ", err)
			return
		}

		result += " " + strconv.Itoa(res)
	}

	if options == "-m" {
		res, err := countCharactersInFile(filePath)
		if err != nil {
			fmt.Println("Error while counting characters in file: ", err)
			return
		}

		result += " " + strconv.Itoa(res)
	}

	fmt.Println(result, " ", filePath)
}