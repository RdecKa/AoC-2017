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

// Abs : returns absolute value of integer a
func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// Reverse : takes pointer to array of integers, reverse part between start and end
func Reverse(input *[]int, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		(*input)[i], (*input)[j] = (*input)[j], (*input)[i]
	}
}

// CircularShift : takes pointer to array of integers, performs circular shift for 'shift' positions to the left
func CircularShift(input *[]int, shift int) {
	shift = shift % len(*input)
	*input = append((*input)[shift:], (*input)[0:shift]...)
}

// CircularShiftRunes : takes pointer to array of runes, performs circular shift for 'shift' positions to the left
func CircularShiftRunes(input *[]rune, shift int) {
	shift = shift % len(*input)
	*input = append((*input)[shift:], (*input)[0:shift]...)
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

// SliceIndex : returns index of first element in a slice with property 'predicate' in range [start : end]
func SliceIndex(start, end int, predicate func(i int) bool) int {
	for i := start; i < end; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
