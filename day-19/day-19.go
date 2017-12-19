package main

import (
	"fmt"

	"github.com/RdecKa/AoC-2017/useful"
)

var network []string

// Bum around the network
// dir: 0-right, 1-down, 2-left, 3-up
func travel(start int) (string, int) {
	x, y, dir := start, 0, 1
	count := 0
	str := ""
	for dir >= 0 {
		switch dir {
		case 0:
			if x < len(network[0])-1 && network[y][x+1] != ' ' {
				x++
			} else if y > 0 && network[y-1][x] != ' ' {
				y--
				dir = 3
			} else if y < len(network)-1 && network[y+1][x] != ' ' {
				y++
				dir = 1
			} else {
				dir = -1
			}
		case 1:
			if y < len(network)-1 && network[y+1][x] != ' ' {
				y++
			} else if x > 0 && network[y][x-1] != ' ' {
				x--
				dir = 2
			} else if x < len(network[0])-1 && network[y][x+1] != ' ' {
				x++
				dir = 0
			} else {
				dir = -1
			}
		case 2:
			if x > 0 && network[y][x-1] != ' ' {
				x--
			} else if y > 0 && network[y-1][x] != ' ' {
				y--
				dir = 3
			} else if y < len(network)-1 && network[y+1][x] != ' ' {
				y++
				dir = 1
			} else {
				dir = -1
			}
		case 3:
			if y > 0 && network[y-1][x] != ' ' {
				y--
			} else if x > 0 && network[y][x-1] != ' ' {
				x--
				dir = 2
			} else if x < len(network[0])-1 && network[y][x+1] != ' ' {
				x++
				dir = 0
			} else {
				dir = -1
			}
		}
		count++
		if dir >= 0 && network[y][x] != '-' && network[y][x] != '+' && network[y][x] != '|' {
			str += string(network[y][x])
		}
	}
	return str, count
}

// Find begining of the path
func findStart() int {
	return useful.SliceIndex(0, len(network[0]), func(i int) bool { return network[0][i] != ' ' })
}

// Debugging: Print the whole network.
func printNetwork() {
	for row := range network {
		for col := range network[row] {
			fmt.Print(string(network[row][col]))
		}
		fmt.Println()
	}
}

func main() {
	// Read lines, save them as strings
	network = useful.StringToLines(useful.FileToString("input.txt"))
	maxLen := -1

	// Find maximal length and extend all lines to this length
	for _, line := range network {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	for i, line := range network {
		for j := len(line); j < maxLen; j++ {
			network[i] = network[i] + " "
		}
	}

	fmt.Println("Results:")
	p1, p2 := travel(findStart())
	fmt.Printf("Part1: %s\n", p1)
	fmt.Printf("Part2: %d\n", p2)
}
