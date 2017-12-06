package main

import (
	"AoC/useful"
	"fmt"
)

const numBanks = 16 // Change to 4 for small_input.txt

func puzzle1(input [numBanks]int) int {
	count := 0
	seen := make(map[[numBanks]int]bool)
	for {
		count++
		// Find maximum
		maxInd := 0
		max := input[0]
		for i, v := range input {
			if v > max {
				maxInd = i
				max = v
			}
		}
		// Reallocate
		input[maxInd] = 0
		change := (maxInd + 1) % numBanks
		for max > 0 {
			input[change]++
			max--
			change = (change + 1) % numBanks
		}
		// Keys in map can be arrays, but not slices
		cpy := make([]int, numBanks)
		copy(cpy, input[:])
		if seen[input] {
			return count
		}
		seen[input] = true
	}
}

func puzzle2(input [numBanks]int) int {
	count := 0
	seen := make(map[[numBanks]int]int)
	for {
		count++
		maxInd := 0
		max := input[0]
		for i, v := range input {
			if v > max {
				maxInd = i
				max = v
			}
		}
		input[maxInd] = 0
		change := (maxInd + 1) % numBanks
		for max > 0 {
			input[change]++
			max--
			change = (change + 1) % numBanks
		}
		cpy := make([]int, numBanks)
		copy(cpy, input[:])
		_, inMap := seen[input]
		if inMap {
			return count - seen[input]
		}
		seen[input] = count
	}
}

func main() {
	input := useful.SplitOnWhitespace(useful.FileToString("input.txt"))
	inputSplit := useful.StringsToIntsArr1D(input)
	var in1 [numBanks]int
	var in2 [numBanks]int
	for i := 0; i < numBanks; i++ {
		in1[i] = inputSplit[i]
		in2[i] = inputSplit[i]
	}

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in1))
	fmt.Printf("Part2: %d\n", puzzle2(in2))
}
