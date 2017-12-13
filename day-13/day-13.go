package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RdecKa/AoC-2017/useful"
)

var state = make([][]int, 0)

func puzzle1(input [][]string) int {
	// Init: state[i] = [currentPosition, range, movingUpOrDown(-1/1)]
	for _, line := range input {
		i, _ := strconv.Atoi(strings.Replace(line[0], ":", "", -1))
		if i >= len(state) {
			for n := len(state); n <= i; n++ {
				state = append(state, make([]int, 0))
				state[n] = append(state[n], 0)
			}
		}
		v, _ := strconv.Atoi(line[1])
		state[i] = append(state[i], v, 1)
	}

	return move()
}

// Returns score for going through scanners
func move() int {
	score := 0
	for pos := range state { // Steps
		if len(state[pos]) > 1 && state[pos][0] == 0 {
			score += pos * state[pos][1]
		}
		moveScanners()
	}
	return score
}

// Move scanners for one step
func moveScanners() {
	for i := range state {
		if len(state[i]) > 1 {
			if state[i][0] == 0 {
				state[i][2] = 1
			} else if state[i][0] == state[i][1]-1 {
				state[i][2] = -1
			}
			state[i][0] += state[i][2]
		}
	}
}

func puzzle2(input [][]string) int {
	pico := 0
	for {
		resetScanners()
		moveScannersDelay(pico)

		if !moveCheck() {
			pico++
		} else {
			return pico
		}
	}
}

// Check if we get caught when we go through the scanners
func moveCheck() bool {
	for pos := range state { // Steps
		if len(state[pos]) > 1 && state[pos][0] == 0 {
			return false
		}
		moveScanners()
	}
	return true
}

// Move scanners for 'delay' steps
func moveScannersDelay(delay int) {
	for i := range state {
		if len(state[i]) > 1 {
			state[i][0] = delay % ((state[i][1] - 1) * 2)
			if state[i][0] >= state[i][1] {
				if state[i][0] >= state[i][1]-1 {
					state[i][2] = -1
				}
				state[i][0] = 2*state[i][1] - state[i][0] - 2
			}
		}
	}
}

// Reset scanners to initial state
func resetScanners() {
	for i := range state {
		if len(state[i]) > 1 {
			state[i][0] = 0
			state[i][2] = 1
		}
	}
}

func main() {
	in := useful.StringTo2DArray(useful.FileToString("input.txt"))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in))
	fmt.Printf("Part2: %d\n", puzzle2(in))
}
