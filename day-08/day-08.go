package main

import (
	"AoC/useful"
	"fmt"
	"strconv"
)

func puzzle(input []string, problem int) int {
	num := 0
	maxValue := 0
	indices := make(map[string]int)
	registers := make([]int, 0, 1)

	for _, v := range input {
		line := useful.SplitOnWhitespace(v)
		if _, inMap := indices[line[0]]; !inMap {
			indices[line[0]] = num
			num++
			registers = append(registers, 0)
		}
		diff, _ := strconv.Atoi(line[2])
		if line[1] == "dec" {
			diff = -diff
		}
		ind, exists := indices[line[4]]
		if !exists {
			indices[line[4]] = num
			num++
			registers = append(registers, 0)
			ind = indices[line[4]]
		}
		check := registers[ind]
		comp, _ := strconv.Atoi(line[6])
		switch line[5] {
		case ">":
			if check > comp {
				registers[indices[line[0]]] += diff
			}
		case "<":
			if check < comp {
				registers[indices[line[0]]] += diff
			}
		case ">=":
			if check >= comp {
				registers[indices[line[0]]] += diff
			}
		case "<=":
			if check <= comp {
				registers[indices[line[0]]] += diff
			}
		case "==":
			if check == comp {
				registers[indices[line[0]]] += diff
			}
		case "!=":
			if check != comp {
				registers[indices[line[0]]] += diff
			}
		}

		if registers[indices[line[0]]] > maxValue {
			maxValue = registers[indices[line[0]]]
		}
	}

	if problem == 1 {
		max := registers[0]
		for _, v := range registers[1:] {
			if v > max {
				max = v
			}
		}

		return max
	}
	return maxValue
}

func main() {
	input := useful.StringToLines(useful.FileToString("input.txt"))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle(input, 1))
	fmt.Printf("Part2: %d\n", puzzle(input, 2))
}
