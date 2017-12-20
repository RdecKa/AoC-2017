package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RdecKa/AoC-2017/useful"
)

var numRepeat = 1000

func puzzle1(input [][]int) int {
	for i := 0; i < numRepeat; i++ {
		for l := 0; l < len(input); l++ {
			input[l][3] += input[l][6]
			input[l][4] += input[l][7]
			input[l][5] += input[l][8]
			input[l][0] += input[l][3]
			input[l][1] += input[l][4]
			input[l][2] += input[l][5]
		}
	}

	minD := useful.Abs(input[0][0]) + useful.Abs(input[0][1]) + useful.Abs(input[0][2])
	minI := 0
	for i, data := range input {
		d := useful.Abs(data[0]) + useful.Abs(data[1]) + useful.Abs(data[2])
		if d < minD {
			minD = d
			minI = i
		}
	}

	return minI
}

func puzzle2(input [][]int) int {
	for i := 0; i < len(input); i++ {
		input[i] = append(input[i], 1)
	}

	for i := 0; i < numRepeat; i++ {
		toRemove := make([]int, 0)

		for j := 0; j < len(input); j++ {
			if input[j][9] == 1 { // Still alive
				for k := j + 1; k < len(input); k++ {
					if input[k][9] == 1 {
						if input[j][0] == input[k][0] &&
							input[j][1] == input[k][1] &&
							input[j][2] == input[k][2] {
							toRemove = append(toRemove, j, k)
						}
					}
				}
			}
		}

		for _, e := range toRemove {
			input[e][9] = 0
		}

		for l := 0; l < len(input); l++ {
			input[l][3] += input[l][6]
			input[l][4] += input[l][7]
			input[l][5] += input[l][8]
			input[l][0] += input[l][3]
			input[l][1] += input[l][4]
			input[l][2] += input[l][5]
		}
	}

	count := 0
	for _, p := range input {
		if p[9] == 1 {
			count++
		}
	}

	return count
}

func parseInput(input string) [][]int {
	lines := useful.StringToLines(input)
	out := make([][]int, len(lines))
	for i, line := range lines {
		out[i] = make([]int, 9)
		l := strings.Fields(line)

		pref := []string{"p=<", "v=<", "a=<"}
		suff := []string{">,", ">,", ">"}

		for j, coord := range l {
			p := strings.TrimPrefix(coord, pref[j])
			p = strings.TrimSuffix(p, suff[j])
			for k, v := range strings.Split(p, ",") {
				out[i][j*3+k], _ = strconv.Atoi(v)
			}
		}
	}
	return out
}

func main() {
	file := "input.txt"
	in := parseInput(useful.FileToString(file))
	in2 := parseInput(useful.FileToString(file))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in))
	fmt.Printf("Part2: %d\n", puzzle2(in2))
}
