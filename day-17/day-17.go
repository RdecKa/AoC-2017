package main

import (
	"fmt"

	"github.com/RdecKa/AoC-2017/useful"
)

const step = 344 // input
const numRepeat = 2017
const giantNumRepeat = 50000000

func puzzle1() int {
	buffer := make([]int, 1)
	for i := 1; i <= numRepeat; i++ {
		// Move current position to index 0
		// +1 to first update current position to last inserted value
		useful.CircularShift(&buffer, 1+step)

		// Insert new element to position 1
		buffer = append([]int{0}, buffer...)
		buffer[0] = buffer[1]
		buffer[1] = i
	}
	return buffer[2]
}

func puzzle2() int {
	curLength := 1 // Current length of buffer
	curPos := 0    // Current position
	curEl1 := -1   // Current element at position 1 (after 0)
	for i := 1; i <= giantNumRepeat; i++ {
		// Step forward
		curPos = (curPos + step) % curLength

		// Insert
		curLength++

		// Update current position to last inserted value
		curPos = (curPos + 1) % curLength

		// Check if the value was inserted directly after 0
		if curPos == 1 {
			curEl1 = i
		}
	}
	return curEl1
}

func main() {
	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1())
	fmt.Printf("Part2: %d\n", puzzle2())
}
