package main

import (
	"fmt"
	"os"
)

var garbageCount = 0

func collectGarbage(file *os.File, in []byte) {
	file.Read(in)
	for string(in) != ">" {
		if string(in) == "!" {
			file.Read(in)
		} else {
			garbageCount++
		}
		file.Read(in)
	}
	file.Read(in)
}

func puzzle1(file *os.File, in []byte, level int) int {
	score := 0
	for string(in) == "{" || string(in) == "," || string(in) == "<" {
		if string(in) == "," {
			file.Read(in)
			continue
		} else if string(in) == "<" {
			collectGarbage(file, in)
			continue
		}
		file.Read(in)
		score += puzzle1(file, in, level+1)
		if string(in) == "}" {
			file.Read(in)
		} else {
			fmt.Println("ERROR: Unclosed group")
		}
	}
	return level + score
}

func main() {
	f, _ := os.Open("input.txt")
	in := make([]byte, 1)
	f.Read(in)

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(f, in, 0))
	fmt.Printf("Part2: %d\n", garbageCount)
}
