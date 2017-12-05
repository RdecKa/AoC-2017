package main

import (
	"AoC/useful"
	"fmt"
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

func main() {
	input := useful.StringToLines(useful.FileToString("input.txt"))
	inputSplit := useful.StringsToIntsArr1D(input)
	inputSplit2 := make([]int, len(inputSplit))
	copy(inputSplit2, inputSplit)
	fmt.Printf("Result:\nPart 1: %d\nPart 2: %d\n", puzzle1(inputSplit), puzzle2(inputSplit2))
}
