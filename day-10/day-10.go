package main

import (
	"fmt"
	"strings"

	"github.com/RdecKa/AoC-2017/useful"
)

const listLen = 256

func puzzle1(input []int) int {
	list := make([]int, listLen)
	for i := range list {
		list[i] = i
	}
	shifted := 0
	skipSize := 0
	for _, len := range input {
		useful.Reverse(&list, 0, len)
		useful.CircularShift(&list, (len + skipSize))
		shifted += len + skipSize
		skipSize++
	}
	shifted %= len(list)
	shiftBack := len(list) - shifted
	useful.CircularShift(&list, shiftBack)

	return list[0] * list[1]
}

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

func puzzle2(inputStr string) string {
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

func main() {
	str := useful.FileToString("input.txt")
	input1 := useful.StringsToIntsArr1D(strings.Split(str, ","))
	input2 := strings.TrimSuffix(str, "\n")

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(input1))
	fmt.Printf("Part2: %s\n", puzzle2(input2))
}
