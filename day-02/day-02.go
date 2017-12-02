package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func puzzle1(input string) int {
	inputSplit := strings.Split(input, "\n")
	sum := 0

	for _, list := range inputSplit {
		splitList := strings.Fields(list)
		max, min := 0, math.MaxInt64
		for _, v := range splitList {
			currentValue, _ := strconv.Atoi(v)
			if currentValue > max {
				max = currentValue
			}
			if currentValue < min {
				min = currentValue
			}
		}

		sum += max - min
	}

	return sum
}

func puzzle2(input string) int {
	inputSplit := strings.Split(input, "\n")
	sum := 0

	for _, list := range inputSplit {
		splitList := strings.Fields(list)
		for i, v1 := range splitList {
			cVal1, _ := strconv.Atoi(v1)
			for _, v2 := range splitList[i+1:] {
				cVal2, _ := strconv.Atoi(v2)
				if cVal1%cVal2 == 0 {
					sum += cVal1 / cVal2
					break
				} else if cVal2%cVal1 == 0 {
					sum += cVal2 / cVal1
					break
				}
			}
		}
	}

	return sum
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inputStr := strings.TrimSuffix(string(input), "\n")
	fmt.Printf("%d\n%d\n", puzzle1(inputStr), puzzle2(inputStr))
}
