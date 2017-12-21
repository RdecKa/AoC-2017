package main

import (
	"fmt"

	"github.com/RdecKa/AoC-2017/useful"
)

func puzzle(input [][]string, numRounds int) int {
	patterns2 := make([][]string, 0)
	patterns3 := make([][]string, 0)

	for _, rule := range input {
		if len(rule[0]) == 11 {
			patterns3 = append(patterns3, rule)
		} else {
			patterns2 = append(patterns2, rule)
		}
	}

	grid := make([][]string, 1)
	grid[0] = make([]string, 1)
	grid[0][0] = ".#./..#/###"

	for r := 0; r < numRounds; r++ {
		fmt.Println("... Round", r+1)
		if len(grid)*len(grid[0][0])%2 == 0 {
			// Patterns of size 2
			if len(grid[0][0]) == 11 {
				grid = transform3to2(grid)
			}
			for y := range grid {
				for x := range grid[y] {
					grid[y][x] = change(grid[y][x], patterns2, matchPattern2, rotate2, flip2)
				}
			}
		} else {
			// Patterns of size 3
			newGrid := make([][]string, len(grid)*2)
			for i := 0; i < len(newGrid); i++ {
				newGrid[i] = make([]string, len(grid[0])*2)
			}
			for y := range grid {
				for x := range grid[0] {
					newPatt := change(grid[y][x], patterns3, matchPattern3, rotate3, flip3)
					newGrid[y*2+0][x*2+0] = join(newPatt, []int{0, 1, 4, 5, 6})
					newGrid[y*2+0][x*2+1] = join(newPatt, []int{2, 3, 4, 7, 8})
					newGrid[y*2+1][x*2+0] = join(newPatt, []int{10, 11, 4, 15, 16})
					newGrid[y*2+1][x*2+1] = join(newPatt, []int{12, 13, 4, 17, 18})
				}
			}
			grid = newGrid
		}
	}
	return count(grid)
}

// Transforms grid of patterns 3×3 to grid of patterns 2×2
func transform3to2(grid [][]string) [][]string {
	tmpGrid := make([][]string, len(grid)*3)
	for i := range tmpGrid {
		tmpGrid[i] = make([]string, len(grid)*3)
	}

	for y := range grid {
		for x := range grid[y] {
			p := grid[y][x]
			tmpGrid[y*3+0][x*3+0] = string(p[0])
			tmpGrid[y*3+0][x*3+1] = string(p[1])
			tmpGrid[y*3+0][x*3+2] = string(p[2])
			tmpGrid[y*3+1][x*3+0] = string(p[4])
			tmpGrid[y*3+1][x*3+1] = string(p[5])
			tmpGrid[y*3+1][x*3+2] = string(p[6])
			tmpGrid[y*3+2][x*3+0] = string(p[8])
			tmpGrid[y*3+2][x*3+1] = string(p[9])
			tmpGrid[y*3+2][x*3+2] = string(p[10])
		}
	}

	newGrid := make([][]string, len(grid)*3/2)
	for i := range newGrid {
		newGrid[i] = make([]string, len(grid)*3/2)
	}

	for y := range newGrid {
		for x := range newGrid[y] {
			newGrid[y][x] = "" +
				tmpGrid[y*2+0][x*2+0] +
				tmpGrid[y*2+0][x*2+1] +
				"/" +
				tmpGrid[y*2+1][x*2+0] +
				tmpGrid[y*2+1][x*2+1]
		}
	}

	return newGrid
}

// Conut the number of set pixels
func count(grid [][]string) int {
	c := 0
	for y := range grid {
		for x := range grid {
			for _, ch := range grid[y][x] {
				if ch == '#' {
					c++
				}
			}
		}
	}
	return c
}

// Find a pattern that matches current square
func change(square string, rules [][]string, match func(string, string) bool, rotate func(string) string, flipX func(string) string) string {
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			for _, rule := range rules {
				if match(rule[0], square) {
					return rule[2]
				}
			}
			square = rotate(square)
		}
		square = flipX(square)
	}
	fmt.Println("Cannot find", square)
	return "ERROR"
}

// Returns a string composed of characters on indices 'indices' in pattern
func join(pattern string, indices []int) string {
	str := ""
	for _, i := range indices {
		str += string(pattern[i])
	}
	return str
}

// Check match
// example: matchOnePattern("../.#", []string{"..", ".#"}) -> true
func matchOnePattern(pattern string, parts []string) bool {
	lenPattern := len(parts[0])
	for i, p := range parts {
		if p != pattern[i*lenPattern+i:(i+1)*lenPattern+i] {
			return false
		}
	}
	return true
}

// Check if square and pattern match (size 3×3)
func matchPattern3(pattern, square string) bool {
	if matchOnePattern(pattern, []string{
		join(square, []int{0, 1, 2}),
		join(square, []int{4, 5, 6}),
		join(square, []int{8, 9, 10}),
	}) {
		return true
	}
	return false
}

// Check if square and pattern match (size 2×2)
func matchPattern2(pattern, square string) bool {
	if matchOnePattern(pattern, []string{
		join(square, []int{0, 1}),
		join(square, []int{3, 4}),
	}) {
		return true
	}
	return false
}

// Rotate a square 2×2 for 90 degrees clockwise
// Example: #./.. -> .#/..
func rotate2(square string) string {
	new := ""
	for _, el := range []int{3, 0, 2, 4, 1} {
		new += string(square[el])
	}
	return new
}

// Rotate a square 3×3 for 90 degrees clockwise
// Example: #../.../... -> ..#/.../...
func rotate3(square string) string {
	new := ""
	for _, el := range []int{8, 4, 0, 3, 9, 5, 1, 7, 10, 6, 2} {
		new += string(square[el])
	}
	return new
}

// Flips a square of size 2×2
func flip2(square string) string {
	return square[3:5] + "/" + square[0:2]
}

// Flips a square of size 3×3
func flip3(square string) string {
	return square[8:11] + "/" + square[4:7] + "/" + square[0:3]
}

func main() {
	in := useful.StringTo2DArray(useful.FileToString("input.txt"))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle(in, 5))
	fmt.Printf("Part2: %d\n", puzzle(in, 18))
}
