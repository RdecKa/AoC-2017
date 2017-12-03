package main

import (
	"fmt"
	"math"
)

func puzzle1(input int) int {
	start := 1
	len := 1
	count := 1
	for {
		l := make([]int, len)
		for i := 0; i <= len/2; i++ {
			l[i] = start - i
			l[len-1-i] = start - i
		}
		for i := 0; i < 4; i++ {
			for j := 0; j <= len; j++ {
				count++
				if count == input {
					if j == len {
						return start + 1
					}
					return l[j]
				}
			}
		}
		len += 2
		start += 2
	}
}

func puzzle2(input int) int {
	size := int(math.Ceil(math.Sqrt(float64(input)))) // Wrong. But good enough.
	center := size / 2
	t := make([][]int, size)
	for i := 0; i < size; i++ {
		t[i] = make([]int, size)
	}
	t[center][center] = 1
	direction := 0 // 0: right, 1: up, 2: left, 3: down
	x, y := center, center
	for t[y][x] <= input {
		switch direction {
		case 0:
			x++
			if t[y-1][x] == 0 {
				direction = 1
			}
		case 1:
			y--
			if t[y][x-1] == 0 {
				direction = 2
			}
		case 2:
			x--
			if t[y+1][x] == 0 {
				direction = 3
			}
		case 3:
			y++
			if t[y][x+1] == 0 {
				direction = 0
			}
		}
		t[y][x] = sumOfNeighbours(t, x, y)
	}

	return t[y][x]
}

func sumOfNeighbours(t [][]int, x, y int) int {
	sum := 0
	for i := int(math.Max(0, float64(y-1))); i <= int(math.Min(float64(len(t)), float64(y+1))); i++ {
		for j := int(math.Max(0, float64(x-1))); j <= int(math.Min(float64(len(t[i])), float64(x+1))); j++ {
			if i == y && j == x {
				continue
			}
			sum += t[i][j]
		}
	}
	return sum
}

func main() {
	input := 265149
	fmt.Printf("Result:\n%d\n%d\n", puzzle1(input), puzzle2(input))
}
