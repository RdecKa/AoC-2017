package main

import (
	"fmt"
	"strings"

	"github.com/RdecKa/AoC-2017/useful"
)

type node struct {
	index   string
	neigh   []*node
	inGroup bool
}

var gMap = make(map[string]*node)

func puzzle1(input [][]string) int {
	for _, v := range input { // each line of input
		ind := v[0] // current node
		_, ok := gMap[ind]
		if !ok {
			// vertex not yet in the map, add it
			gMap[ind] = &node{ind, make([]*node, 0), false}
		}
		for _, n := range v[2:] {
			// add conections to each neighbour
			c := strings.Replace(n, ",", "", -1) // get rid of commas
			addr, inMap := gMap[c]
			if !inMap {
				// neighbour not yet in the map, add it
				gMap[c] = &node{c, make([]*node, 0), false}
				addr = gMap[c]
			}
			gMap[ind].neigh = append(gMap[ind].neigh, addr)
		}
	}

	m := markGroup(gMap["0"])

	return m
}

// Return number of elements in group
func markGroup(n *node) int {
	if n.inGroup {
		return 0
	}
	count := 1 // itself
	n.inGroup = true
	for _, c := range n.neigh {
		count += markGroup(c)
	}
	return count
}

func puzzle2(input [][]string) int {
	numGr := 1 // group 0 already marked
	for _, v := range gMap {
		if !v.inGroup {
			markGroup(v)
			numGr++
		}
	}
	return numGr
}

func main() {
	in := useful.StringTo2DArray(useful.FileToString("input.txt"))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in))
	fmt.Printf("Part2: %d\n", puzzle2(in))
}
