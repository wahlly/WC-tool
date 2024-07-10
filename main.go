package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main () {
	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	opts := strings.Split(val, " ")

	if len(opts) != 3 {
		fmt.Println("Invalid command!")
		return
	}

	if opts[0] != "ccwc" {
		fmt.Println("Invalid command!")
		return
	}

	if opts[1] == "-c" {
		getFileByteSize(opts[2])
		return
	} else if opts[1] == "-l" {
		countFileLines(opts[2])
	}
}