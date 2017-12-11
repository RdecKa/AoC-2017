package main

import (
	"fmt"
	"strings"

	"github.com/RdecKa/AoC-2017/useful"
)

func puzzle(input []string, problem int) int {
	x, y, max := 0, 0, 0
	for _, step := range input {
		switch step {
		case "n":
			y--
		case "ne":
			x++
			y--
		case "se":
			x++
		case "s":
			y++
		case "sw":
			x--
			y++
		case "nw":
			x--
		}
		max = useful.Max(max, getDistance(x, y))
	}
	if problem == 1 {
		return getDistance(x, y)
	}
	return max
}

// get distance between (0, 0) and (x, y) in a hex grid
func getDistance(x, y int) int {
	if (x <= 0 && y <= 0) || (x >= 0 && y >= 0) {
		return useful.Abs(x) + useful.Abs(y)
	}
	xA, yA := useful.Abs(x), useful.Abs(y)
	return xA + yA - useful.Min(xA, yA)
}

func main() {
	input := strings.Split(useful.FileToString("input.txt"), ",")

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle(input, 1))
	fmt.Printf("Part2: %d\n", puzzle(input, 2))
}
