package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func puzzle1(input string) int {
	inputSplit := strings.Split(input, "")
	sum := 0

	v1 := inputSplit[len(inputSplit)-1]

	for _, v := range inputSplit {
		if v1 == v {
			v2, _ := strconv.Atoi(v)
			sum += v2
		}
		v1 = v
	}

	return sum
}

func puzzle2(input string) int {
	inputSplit := strings.Split(input, "")
	sum := 0

	jump := len(inputSplit) / 2

	for i := range inputSplit {
		if inputSplit[i] == inputSplit[(i+jump)%len(inputSplit)] {
			v, _ := strconv.Atoi(inputSplit[i])
			sum += v
		}
	}

	return sum
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inputStr := strings.TrimSuffix(string(input), "\n")
	fmt.Printf("%d\n%d\n", puzzle1(inputStr), puzzle2(inputStr))
}
