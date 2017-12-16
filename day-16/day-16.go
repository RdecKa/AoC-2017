package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RdecKa/AoC-2017/useful"
)

var order *[]rune
var numRounds = 1000000000
var m map[string]int

func puzzle1(input []string) string {
	for _, move := range input {
		sp := strings.Split(move, "/")
		if len(sp) == 1 {
			// SPIN (s14)
			val, _ := strconv.Atoi(string(sp[0][1:]))
			useful.CircularShiftRunes(order, len(*order)-val)
		} else {
			var firstAd *rune
			var secondAd *rune
			if sp[0][0] == 'x' {
				// EXCHANGE (x4/11)
				f, _ := strconv.Atoi(sp[0][1:])
				s, _ := strconv.Atoi(sp[1])
				firstAd = &(*order)[f]
				secondAd = &(*order)[s]
			} else if sp[0][0] == 'p' {
				// PARTNER (pb/c)
				firstAd = &(*order)[useful.SliceIndex(0, len(*order),
					func(i int) bool { return (*order)[i] == rune(sp[0][1]) })]
				secondAd = &(*order)[useful.SliceIndex(0, len(*order),
					func(i int) bool { return (*order)[i] == rune(sp[1][0]) })]
			}
			*firstAd, *secondAd = *secondAd, *firstAd
		}
	}
	return string(*order)
}

func puzzle2(input []string) string {
	m = make(map[string]int)
	var cycle int

	// Find cycle length
	for i := 1; i < numRounds; i++ {
		v, inMap := m[string(*order)]
		if inMap {
			cycle = i - v
			break
		}
		m[string(*order)] = i

		puzzle1(input)
	}

	for i := 1; i < numRounds%cycle; i++ {
		puzzle1(input)
	}

	return string(*order)
}

func main() {
	order = &([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'})
	in := strings.Split(useful.FileToString("input.txt"), ",")

	fmt.Println("Results:")
	fmt.Printf("Part1: %s\n", puzzle1(in))
	fmt.Printf("Part2: %s\n", puzzle2(in))
}
