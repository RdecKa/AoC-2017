package main

import (
	"fmt"

	"github.com/RdecKa/AoC-2017/useful"
)

func puzzle1(input []string, numRepeat int) int {
	/*
		Store grid in a map:
			Key: coordinates
			Value:
				true: clean
				false: infected
	*/
	clean := make(map[[2]int]bool)
	center := int(len(input) / 2)
	for y := -center; y <= center; y++ {
		for x := -center; x <= center; x++ {
			if c := input[y+center][x+center]; c == '.' {
				clean[[2]int{x, y}] = true
			} else {
				clean[[2]int{x, y}] = false
			}
		}
	}

	/*
		Direction:
			0: right
			1: down
			2: left
			3: up
	*/
	dir := 3            // Current direction
	pos := [2]int{0, 0} // Current position of virus carrier
	count := 0          // Number of infections

	for r := 0; r < numRepeat; r++ {
		_, inMap := clean[pos]
		if !inMap {
			// Add to map
			clean[pos] = true
		}

		if clean[pos] == true {
			// Turn left
			dir = (dir + 3) % 4
		} else {
			// Turn right
			dir = (dir + 1) % 4
		}

		if clean[pos] {
			count++
		}
		clean[pos] = !clean[pos]

		// Move forward
		switch dir {
		case 0:
			pos[0]++
		case 1:
			pos[1]++
		case 2:
			pos[0]--
		case 3:
			pos[1]--
		}

		//draw(clean, 8)
	}

	return count
}

// Draw a part of the map (puzzle 1): [-size, size] in both directions
func draw(clean map[[2]int]bool, size int) {
	for y := -size; y <= size; y++ {
		for x := -size; x <= size; x++ {
			el, inMap := clean[[2]int{x, y}]
			if !inMap || el {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

// Draw a part of the map (puzzle 2): [-size, size] in both directions
func draw2(clean map[[2]int]int, size int) {
	for y := -size; y <= size; y++ {
		for x := -size; x <= size; x++ {
			el, inMap := clean[[2]int{x, y}]
			if !inMap || el == 0 {
				fmt.Print(".")
			} else if el == 1 {
				fmt.Print("W")
			} else if el == 2 {
				fmt.Print("#")
			} else {
				fmt.Print("F")
			}
		}
		fmt.Println()
	}
}

func puzzle2(input []string, numRepeat int) int {
	/*
		Key: coordinates
		Value:
			0: clean
			1: weakened
			2: infected
			3: flagged
	*/
	state := make(map[[2]int]int)
	center := int(len(input) / 2)
	for y := -center; y <= center; y++ {
		for x := -center; x <= center; x++ {
			if c := input[y+center][x+center]; c == '.' {
				state[[2]int{x, y}] = 0 // Clean
			} else {
				state[[2]int{x, y}] = 2 // Infected
			}
		}
	}

	dir := 3
	pos := [2]int{0, 0}
	count := 0

	for r := 0; r < numRepeat; r++ {
		_, inMap := state[pos]
		if !inMap {
			state[pos] = 0
		}

		if state[pos] == 0 {
			// Turn left
			dir = (dir + 3) % 4
		} else if state[pos] == 2 {
			// Turn right
			dir = (dir + 1) % 4
		} else if state[pos] == 3 {
			// Reverse
			dir = (dir + 2) % 4
		}

		if state[pos] == 1 {
			count++
		}
		state[pos] = (state[pos] + 1) % 4

		switch dir {
		case 0:
			pos[0]++
		case 1:
			pos[1]++
		case 2:
			pos[0]--
		case 3:
			pos[1]--
		}

		//draw2(state, 8)
	}

	return count
}

func main() {
	in := useful.StringToLines(useful.FileToString("input.txt"))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in, 10000))
	fmt.Printf("Part2: %d\n", puzzle2(in, 10000000))
}
