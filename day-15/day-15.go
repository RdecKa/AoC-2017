package main

import (
	"fmt"
)

var aSeed, bSeed, rem, lowestBits int64 = 16807, 48271, 2147483647, 65536 - 1

func genPicky(old, seed, divide int64) int64 {
	n := (old * seed) % rem
	for n&(divide-1) != 0 {
		n = (n * seed) % rem
	}
	return n
}

func gen(old, seed, _ int64) int64 {
	return (old * seed) % rem
}

func puzzle(startA, startB, numRounds int, genFunc func(int64, int64, int64) int64) int {
	count := 0
	var a, b = int64(startA), int64(startB)
	for r := 0; r < numRounds; r++ {
		newA := genFunc(a, aSeed, 4)
		newB := genFunc(b, bSeed, 8)
		if (newA & lowestBits) == (newB & lowestBits) {
			count++
		}
		a, b = newA, newB
	}
	return count
}

func main() {
	startA, startB := 783, 325

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle(startA, startB, 40000000, gen))
	fmt.Printf("Part2: %d\n", puzzle(startA, startB, 5000000, genPicky))
}
