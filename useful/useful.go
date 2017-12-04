package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Reads file in a string
func fileToString(fileName string) (inputStr string) {
	input, _ := ioutil.ReadFile("input.txt")
	inputStr = strings.TrimSuffix(string(input), "\n")
	return
}

// Split string on \n
func stringToLines(input string) (inputSplit []string) {
	inputSplit = strings.Split(input, "\n")
	return
}

// Converts string to list
func splitOnWhitespace(input string) (inputSplit []string) {
	inputSplit = strings.Fields(input)
	return
}

func main() {
	input := stringToLines(fileToString("input.txt"))
	for _, v := range input {
		fmt.Println(splitOnWhitespace(v))
	}
	fmt.Printf("%s\n", input)
}
