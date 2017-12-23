package main

import (
	"fmt"
	"strconv"

	"github.com/RdecKa/AoC-2017/useful"
)

// Returns value of argument 'reg'
func getValueOfRegister(reg string, m map[string]int64) int64 {
	val, err := strconv.Atoi(reg)
	if err == nil {
		// 'reg' is a constant
		return int64(val)
	}

	// 'reg' is a register name
	el, inMap := m[reg]
	if inMap {
		return el
	}
	fmt.Println("Unknown register!")
	return 0
}

func puzzle1(input [][]string) int {
	reg := map[string]int64{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
	}

	count := 0
	PC := int64(0)
	for PC >= 0 && int(PC) < len(input) {
		switch instr := input[PC]; instr[0] {
		case "set":
			reg[instr[1]] = getValueOfRegister(instr[2], reg)
		case "sub":
			reg[instr[1]] -= getValueOfRegister(instr[2], reg)
		case "mul":
			reg[instr[1]] *= getValueOfRegister(instr[2], reg)
			count++
		case "jnz":
			if getValueOfRegister(instr[1], reg) != 0 {
				// Substract 1 because PC will be incremented at the end of the loop
				PC += getValueOfRegister(instr[2], reg) - 1
			}
		}
		PC++
	}

	return count
}

func puzzle2() int {
	var b, c, d, f, h int64
	b, c = 81, b
	b = b*100 + 100000
	c = b + 17000

	for {
		f, d = 1, 2
		for {
			// Too slow
			/*for e := int64(2); e != b; e++ {
				if d*e == b {
					f = 0
					break
				}
			}*/
			// End of Too slow

			// Optimised
			if b%d == 0 {
				f = 0
			}
			// End of Optimised

			d++
			if d != b {
				continue
			}

			if f == 0 {
				h++
			}

			if b != c {
				b += 17
				break
			} else {
				return int(h)
			}
		}
	}
}

func main() {
	in := useful.StringTo2DArray(useful.FileToString("input.txt"))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in))
	fmt.Printf("Part2: %d\n", puzzle2())
}
