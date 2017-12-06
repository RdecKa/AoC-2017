package useful

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Min : returns minimum of integers a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max : returns maximum of integers a and b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FileToString : reads file into a string
func FileToString(fileName string) (inputStr string) {
	input, _ := ioutil.ReadFile(fileName)
	inputStr = strings.TrimSuffix(string(input), "\n")
	return
}

// StringToLines : Split string on \n, returns a list of strings
func StringToLines(input string) (inputSplit []string) {
	inputSplit = strings.Split(input, "\n")
	return
}

// SplitOnWhitespace : Converts string to list of words
func SplitOnWhitespace(input string) (inputSplit []string) {
	inputSplit = strings.Fields(input)
	return
}

// StringTo2DArray : creates a list of lists of strings from a string (splits on whitespaces)
func StringTo2DArray(input string) (array2D [][]string) {
	lines := StringToLines(input)
	array2D = make([][]string, len(lines))
	for i, l := range lines {
		array2D[i] = SplitOnWhitespace(l)
	}
	return
}

// StringsToIntsArr1D : accepts list of strings and returns list of integers
func StringsToIntsArr1D(input []string) (output []int) {
	output = make([]int, len(input))
	for i, v := range input {
		output[i], _ = strconv.Atoi(v)
	}
	return
}

// StringsToIntsArr2D : accepts 2D array of strings and returns 2D array of integers
func StringsToIntsArr2D(input [][]string) (output [][]int) {
	output = make([][]int, len(input))
	for i, l := range input {
		output[i] = make([]int, len(l))
		for j, v := range l {
			output[i][j], _ = strconv.Atoi(v)
		}
	}
	return
}
