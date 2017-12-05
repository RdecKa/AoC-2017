package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func puzzle1(input []int) int {
	count := 0
	index := 0
	for index >= 0 && index < len(input) {
		oldIndex := index
		index += input[index]
		input[oldIndex]++
		count++
	}
	return count
}

func puzzle2(input []int) int {
	count := 0
	index := 0
	for index >= 0 && index < len(input) {
		oldIndex := index
		index += input[index]
		if input[oldIndex] >= 3 {
			input[oldIndex]--
		} else {
			input[oldIndex]++
		}
		count++
	}
	return count
}

// Reads file in a string
func fileToString(fileName string) (inputStr string) {
	input, _ := ioutil.ReadFile(fileName)
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
	inputSplit := make([]int, len(input))
	for i, v := range input {
		inputSplit[i], _ = strconv.Atoi(v)
	}
	inputSplit2 := make([]int, len(inputSplit))
	copy(inputSplit2, inputSplit)
	fmt.Printf("Result:\nPart 1: %d\nPart 2: %d\n", puzzle1(inputSplit), puzzle2(inputSplit2))
}
