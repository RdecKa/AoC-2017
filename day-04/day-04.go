package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func puzzle1(input [][]string) int {
	count := 0
	for _, v := range input {
		if !hasDuplicates(v) {
			count++
		}
	}
	return count
}

func puzzle2(input [][]string) int {
	count := 0
	for _, v := range input {
		if !hasAnagrams(v) {
			count++
		}
	}
	return count
}

func hasAnagrams(list []string) bool {
	m := make([]map[rune]int, len(list))
	for i := range m {
		m[i] = make(map[rune]int)
	}
	for i, word := range list {
		for _, char := range word {
			m[i][char]++
		}
		for j := 0; j < i; j++ {
			if isAnagram(m[i], m[j]) && isAnagram(m[j], m[i]) {
				return true
			}
		}
	}
	return false
}

func isAnagram(word1 map[rune]int, word2 map[rune]int) bool {
	for key := range word1 {
		v2, ok := word2[key]
		if !ok {
			return false
		}
		if word1[key] != v2 {
			return false
		}
	}
	return true
}

func hasDuplicates(list []string) bool {
	m := make(map[string]int)
	for _, word := range list {
		m[word]++
	}
	for _, word := range list {
		if m[word] > 1 {
			return true
		}
	}
	return false
}

// Reads file in a string
func fileToString(fileName string) (inputStr string) {
	input, _ := ioutil.ReadFile(fileName)
	inputStr = strings.TrimSuffix(string(input), "\n")
	return
}

// Split string on \n
func stringToLines(input string) (inputSplit []string) {
	inputSplit = strings.Split(input, "\n")
	return
}

// Converts string to list
func splitOnWhitespace(input string) (inputSplit []string) {
	inputSplit = strings.Fields(input)
	return
}

func main() {
	input := stringToLines(fileToString("input.txt"))
	inputSplit := make([][]string, len(input))
	for i, v := range input {
		inputSplit[i] = splitOnWhitespace(v)
	}
	fmt.Printf("Result:\nPart 1: %d\nPart 2: %d\n", puzzle1(inputSplit), puzzle2(inputSplit))
}
