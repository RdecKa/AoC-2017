package main

import (
	"fmt"
)

func puzzle1(numRepeat int) int {
	state, pos := 'A', 0
	tape := make([]bool, 1)

	for r := 0; r < numRepeat; r++ {
		switch state {
		case 'A':
			if !tape[pos] {
				tape[pos] = true
				pos++
				state = 'B'
			} else {
				tape[pos] = true
				pos--
				state = 'E'
			}
		case 'B':
			if !tape[pos] {
				tape[pos] = true
				pos++
				state = 'C'
			} else {
				tape[pos] = true
				pos++
				state = 'F'
			}
		case 'C':
			if !tape[pos] {
				tape[pos] = true
				pos--
				state = 'D'
			} else {
				tape[pos] = false
				pos++
				state = 'B'
			}
		case 'D':
			if !tape[pos] {
				tape[pos] = true
				pos++
				state = 'E'
			} else {
				tape[pos] = false
				pos--
				state = 'C'
			}
		case 'E':
			if !tape[pos] {
				tape[pos] = true
				pos--
				state = 'A'
			} else {
				tape[pos] = false
				pos++
				state = 'D'
			}
		case 'F':
			if !tape[pos] {
				tape[pos] = true
				pos++
				state = 'A'
			} else {
				tape[pos] = true
				pos++
				state = 'C'
			}
		default:
			fmt.Println("Unknown state", state)
			return -1
		}

		if pos >= len(tape) {
			tape = append(tape, false)
		} else if pos < 0 {
			tape = append([]bool{false}, tape...)
			pos++
		}
	}

	s := 0
	for _, v := range tape {
		if v {
			s++
		}
	}

	return s
}

func main() {
	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(12459852))
	fmt.Println("DONE!")
}
