package main

import (
	"fmt"
	"strconv"

	"github.com/RdecKa/AoC-2017/useful"
)

////////////// Solution for Day 10 ///////////////////
const listLen = 256

// returns list of ASCII values for characters in input, adds suffix to the end
func stringToASCII(input string) []int {
	out := make([]int, len(input))
	for i, v := range input {
		out[i] = int(v)
	}
	out = append(out, 17, 31, 73, 47, 23)
	return out
}

// 256 integers ---(XOR)---> 16 integers
func compress(input []int) []int {
	out := make([]int, len(input)/16)
	for i := 0; i < len(out); i++ {
		x := 0
		for j := 0; j < 16; j++ {
			x = x ^ (input[i*16+j])
		}
		out[i] = x
	}
	return out
}

// returns heximal representation of number given as a list of integers (hex digits)
func toHex(input []int) string {
	out := ""
	for _, v := range input {
		h := fmt.Sprintf("%x", v)
		if len(h) == 1 {
			out += "0"
		}
		out += h
	}
	return out
}

func getHash(inputStr string) string {
	input := stringToASCII(inputStr)
	skipSize := 0
	shifted := 0
	list := make([]int, listLen)
	for i := range list {
		list[i] = i
	}

	for i := 0; i < 64; i++ {
		for _, len := range input {
			useful.Reverse(&list, 0, len)
			useful.CircularShift(&list, (len + skipSize))
			shifted += len + skipSize
			skipSize++
		}
		shifted %= len(list)
	}

	shiftBack := len(list) - shifted
	useful.CircularShift(&list, shiftBack)

	return toHex(compress(list))
}

////////////// End of solution for Day 10 ///////////////////

func hexToBinary(s string) string {
	res := ""
	for _, c := range s {
		switch c {
		case '0':
			res += "0000"
		case '1':
			res += "0001"
		case '2':
			res += "0010"
		case '3':
			res += "0011"
		case '4':
			res += "0100"
		case '5':
			res += "0101"
		case '6':
			res += "0110"
		case '7':
			res += "0111"
		case '8':
			res += "1000"
		case '9':
			res += "1001"
		case 'a':
			res += "1010"
		case 'b':
			res += "1011"
		case 'c':
			res += "1100"
		case 'd':
			res += "1101"
		case 'e':
			res += "1110"
		case 'f':
			res += "1111"
		}
	}
	return res
}

func puzzle1(input string) int {
	count := 0

	for i := 0; i < 128; i++ {
		in := input + "-" + strconv.Itoa(i)
		hash := getHash(in)
		bin := hexToBinary(hash)

		for _, bit := range bin {
			if bit == '1' {
				count++
			}
		}
	}

	return count
}

// Recursively mark all used squares adjacent to (i, j)
func markGroup(grid *[][]int, i, j int) {
	if (*grid)[i][j] == 0 {
		// Free square
		return
	}
	(*grid)[i][j] = 0 // Mark as free, so it will not be counted again

	// Check all neighbours
	if i > 0 {
		markGroup(grid, i-1, j)
	}
	if i < 127 {
		markGroup(grid, i+1, j)
	}
	if j > 0 {
		markGroup(grid, i, j-1)
	}
	if j < 127 {
		markGroup(grid, i, j+1)
	}
}

// Returns coordinates of first used square
func checkGrid(grid [][]int) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return i, j
			}
		}
	}
	return -1, -1
}

func puzzle2(input string) int {
	// Make a 2D grid filled with zeros (0)
	t := make([][]int, 128)
	for tt := range t {
		t[tt] = make([]int, 128)
	}

	// Add ones (1) to correct places in the grid - mark as used
	for i := 0; i < 128; i++ {
		in := input + "-" + strconv.Itoa(i)

		hash := getHash(in)

		bin := hexToBinary(hash)
		for j, bit := range bin {
			if bit == '1' {
				t[i][j] = 1
			}
		}
	}

	count := 0

	for {
		// While there is an used square, mark it's group as unused
		i, j := checkGrid(t)
		if i != -1 && j != -1 {
			markGroup(&t, i, j)
			count++
		} else {
			// No more used squares
			break
		}
	}

	return count
}

func main() {
	in := useful.FileToString("input.txt")

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in))
	fmt.Printf("Part2: %d\n", puzzle2(in))
}
