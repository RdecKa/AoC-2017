package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RdecKa/AoC-2017/useful"
)

/*
	connectTo: index of node at the other side of the edge
	used: true if this edge has already been used in current path, false otherwise
*/
type edge struct {
	connectTo int
	used      bool
}

/*
	Input: list of connections (one connection = [2]int{vertex1, vertex2})
	Output: slice of maps of connections for each vertex
*/
func makeGraph(input [][2]int) []map[int]*edge {
	graph := make([]map[int]*edge, 51)
	for i := range graph {
		graph[i] = make(map[int]*edge)
	}
	for _, e := range input {
		// Undirected graph
		a0, a1 := e[0], e[1]
		graph[a0][a1] = &edge{a1, false}
		graph[a1][a0] = &edge{a0, false}
	}
	return graph
}

/*
	Recursive function to find the strongest bridge.
	Input: graph with marked used edges, current node
	Output: maximal strength of a bridge
*/
func findStrongest(graph []map[int]*edge, node int) int {
	max := 0
	for _, edge := range graph[node] {
		cur := 0
		if !edge.used {
			end := edge.connectTo
			graph[node][end].used = true
			graph[end][node].used = true

			cur += node + end + findStrongest(graph, end)

			graph[node][end].used = false
			graph[end][node].used = false
		}

		if cur > max {
			max = cur
		}
	}
	return max
}

/*
	Recursive function to find the longest bridge (if more, find the strongest among them).
	Input: graph with marked used edges, current node
	Output: maximal length and strength of a bridge
*/
func findLongest(graph []map[int]*edge, node int) (int, int) {
	maxLen := 0
	maxStr := 0
	for _, edge := range graph[node] {
		curLen, curStr := 0, 0
		if !edge.used {
			end := edge.connectTo
			graph[node][end].used = true
			graph[end][node].used = true

			l, s := findLongest(graph, end)
			curLen += l + 1
			curStr += node + end + s

			graph[node][end].used = false
			graph[end][node].used = false
		}

		if curLen > maxLen {
			maxLen = curLen
			maxStr = curStr
		} else if curLen == maxLen && curStr > maxStr {
			maxStr = curStr
		}
	}
	return maxLen, maxStr
}

/*
	Builds a slice of edges: [2]int{vertex1, vertex2}
*/
func readInput(file string) [][2]int {
	inp := useful.StringToLines(useful.FileToString(file))
	in := make([][2]int, len(inp))
	for i, line := range inp {
		t := strings.Split(line, "/")
		a0, _ := strconv.Atoi(t[0])
		a1, _ := strconv.Atoi(t[1])
		in[i] = [2]int{a0, a1}
	}
	return in
}

func puzzle1(input [][2]int) int {
	graph := makeGraph(input)
	return findStrongest(graph, 0)
}

func puzzle2(input [][2]int) int {
	graph := makeGraph(input)
	_, s := findLongest(graph, 0)
	return s
}

func main() {
	in := readInput("input.txt")

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in))
	fmt.Printf("Part2: %d\n", puzzle2(in))
}
