package main

import (
	"fmt"
)

var seedA, seedB int64 = 16807, 48271
var startA, startB int64 = 783, 325
var rem, lowestBits int64 = 2147483647, 65536 - 1
var numRounds1, numRounds2 int = 40000000, 5000000

func generator(start, factor, _ int64) func() int64 {
	n := start
	return func() int64 {
		n = (n * factor) % rem
		return n
	}
}

func generatorPicky(start, factor, divide int64) func() int64 {
	n := start
	return func() int64 {
		n = (n * factor) % rem
		for n&(divide-1) != 0 {
			n = (n * factor) % rem
		}
		return n
	}
}

func puzzle(numRounds int, gen func(int64, int64, int64) func() int64) int {
	genA := gen(startA, seedA, 4)
	genB := gen(startB, seedB, 8)
	count := 0
	for r := 0; r < numRounds; r++ {
		newA := genA()
		newB := genB()
		if (newA & lowestBits) == (newB & lowestBits) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle(numRounds1, generator))
	fmt.Printf("Part2: %d\n", puzzle(numRounds2, generatorPicky))
}
