package main

import (
	"AoC/useful"
	"fmt"
	"strconv"
	"strings"
)

type program struct {
	name   string
	weight int
	parent string
}

func returnMap(input []string) map[string]program {
	progs := make(map[string]program)
	for _, v := range input {
		in := useful.SplitOnWhitespace(v)
		weight, _ := strconv.Atoi(in[1][1 : len(in[1])-1])
		if _, inMap := progs[in[0]]; inMap {
			progs[in[0]] = program{in[0], weight, progs[in[0]].parent}
		} else {
			progs[in[0]] = program{in[0], weight, ""}
		}
		if len(in) > 3 {
			for _, v := range in[3:] {
				name := strings.Replace(v, ",", "", -1)
				if _, inMap := progs[name]; inMap {
					progs[name] = program{name, progs[name].weight, in[0]}
				} else {
					progs[name] = program{name, -1, in[0]}
				}
			}
		}
	}
	return progs
}

func puzzle1(input []string) string {
	progs := returnMap(input)
	for _, v := range progs {
		if v.parent == "" {
			return v.name
		}
	}
	return ""
}

type program2 struct {
	name     string
	weight   int
	children []*program2
}

func constructTree(name string, progs map[string]program) *program2 {
	c := 0
	for _, cl := range progs {
		if cl.parent == name {
			c++
		}
	}
	root := program2{name, progs[name].weight, make([]*program2, c)}
	i := 0
	for _, cl := range progs {
		if cl.parent == name {
			root.children[i] = constructTree(cl.name, progs)
			i++
		}
	}
	return &root
}

func puzzle2(input []string, root string) int {
	progs := returnMap(input)

	tree := constructTree(root, progs)

	printTree(tree, 0)

	a, _ := weight(tree)

	return a
}

func weight(node *program2) (int, bool) {
	if node.children == nil || len(node.children) == 0 {
		return node.weight, false
	}

	ws := make([]int, len(node.children))
	s := node.weight
	unbalanced := false

	for i := range node.children {
		ws[i], unbalanced = weight(node.children[i])
		if unbalanced {
			return ws[i], true
		}
		s += ws[i]
	}

	if len(ws) == 1 {
		return s, false
	}

	prev1 := ws[0]
	prev2 := ws[1]
	wrong := -1
	diff := 0
	for i, v := range ws[2:] {
		if prev1 != prev2 {
			if v == prev1 {
				wrong = 1
				diff = prev2 - prev1
			} else if v == prev2 {
				wrong = 0
				diff = prev1 - prev2
			}
		}
		if v != prev1 {
			wrong = i + 2
			diff = v - prev1
		}
	}

	if wrong >= 0 {
		ww := node.children[wrong].weight
		return ww - diff, true
	}

	return s, false
}

func printTree(root *program2, ind int) {
	for i := 0; i < ind; i++ {
		fmt.Print("    ")
	}
	fmt.Println(root.name, root.weight)
	for _, c := range root.children {
		printTree(c, ind+1)
	}
}

func main() {
	input := useful.StringToLines(useful.FileToString("input.txt"))

	root := puzzle1(input)
	fmt.Printf("Results:\nPart1: %s\nPart2: %d\n", root, puzzle2(input, root))
}
