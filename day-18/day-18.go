package main

import (
	"fmt"
	"strconv"

	"github.com/RdecKa/AoC-2017/useful"
)

// Returns value of argument 'reg'
func getValueOfRegister(reg string, m map[string]int64) int64 {
	val, err := strconv.Atoi(reg)
	if err == nil {
		// 'reg' is a constant
		return int64(val)
	}

	// 'reg' is a register name
	el, inMap := m[reg]
	if inMap {
		return el
	}

	// New register, add it to map
	m[reg] = 0
	return 0
}

func puzzle1(input [][]string) int64 {
	m := make(map[string]int64)
	var PC, lastPlayed int64

	for PC >= 0 && PC < int64(len(input)) {
		switch instr := input[PC]; instr[0] {
		case "snd":
			lastPlayed = getValueOfRegister(instr[1], m)
		case "set":
			m[instr[1]] = getValueOfRegister(instr[2], m)
		case "add":
			m[instr[1]] = getValueOfRegister(instr[1], m) + getValueOfRegister(instr[2], m)
		case "mul":
			m[instr[1]] = getValueOfRegister(instr[1], m) * getValueOfRegister(instr[2], m)
		case "mod":
			m[instr[1]] = getValueOfRegister(instr[1], m) % getValueOfRegister(instr[2], m)
		case "rcv":
			if getValueOfRegister(instr[1], m) > 0 {
				m[instr[1]] = lastPlayed
				return lastPlayed
			}
		case "jgz":
			if getValueOfRegister(instr[1], m) > 0 {
				// Substract 1 because PC will be incremented at the end of the iteration
				PC += getValueOfRegister(instr[2], m) - 1
			}
		}

		PC++
	}

	return lastPlayed
}

// RETURN : canContinue, newPC, sendValue
func updateVal(m map[string]int64, PC int64, instrs [][]string) (bool, int64, int64) {
	if PC < 0 || PC >= int64(len(instrs)) {
		return false, PC, -3 // Terminate
	}

	// Set default return values
	canContinue := true
	newPC := PC + 1
	sendValue := int64(-1)

	switch instr := instrs[PC]; instr[0] {
	case "snd":
		canContinue = false
		sendValue = getValueOfRegister(instr[1], m)
	case "set":
		m[instr[1]] = getValueOfRegister(instr[2], m)
	case "add":
		m[instr[1]] = getValueOfRegister(instr[1], m) + getValueOfRegister(instr[2], m)
	case "mul":
		m[instr[1]] = getValueOfRegister(instr[1], m) * getValueOfRegister(instr[2], m)
	case "mod":
		m[instr[1]] = getValueOfRegister(instr[1], m) % getValueOfRegister(instr[2], m)
	case "rcv":
		canContinue = false
		newPC = PC
	case "jgz":
		if getValueOfRegister(instr[1], m) > 0 {
			newPC = PC + getValueOfRegister(instr[2], m)
		}
	default:
		// Terminate
		fmt.Println("Unknown instruction:", instr)
		canContinue = false
		newPC = PC
		sendValue = -3
	}
	return canContinue, newPC, sendValue
}

func puzzle2(input [][]string) int {
	reg0 := map[string]int64{"p": 0}
	reg1 := map[string]int64{"p": 1}

	// Queues to store sent but not yet read data
	qu0, qu1 := make([]int64, 0), make([]int64, 0)

	var term0, term1, wait0, wait1 bool
	var PC0, PC1, sendValue, newPC int64

	count := 0
	var canContinue = true

	for !term0 && !term1 { // While at least one still runs
		for !term0 && canContinue {
			canContinue, newPC, sendValue = updateVal(reg0, PC0, input)
			if sendValue == -3 {
				term0 = true // Process 0 terminated
				break
			}
			if !canContinue {
				if input[PC0][0] == "rcv" {
					if len(qu0) > 0 {
						// Read from queue
						reg0[input[PC0][1]] = qu0[0]
						qu0 = qu0[1:]
						canContinue = true
						newPC++
					} else {
						// No data to read, wait
						wait0 = true
						if wait1 {
							// Deadlock
							term0 = true
							term1 = true
						}
					}
				} else if input[PC0][0] == "snd" {
					qu1 = append(qu1, sendValue)
					canContinue = true
					// If process 1 was waiting, it doesn't have to wait anymore
					wait1 = false
				} else {
					// Process 0 terminated
					term0 = true
				}
			}
			PC0 = newPC
		}

		canContinue = true

		for !term1 && canContinue {
			canContinue, newPC, sendValue = updateVal(reg1, PC1, input)
			if sendValue == -3 {
				term1 = true
				break
			}
			if !canContinue {
				if input[PC1][0] == "rcv" {
					if len(qu1) > 0 {
						reg1[input[PC1][1]] = qu1[0]
						qu1 = qu1[1:]
						canContinue = true
						newPC++
					} else {
						wait1 = true
						if wait0 {
							term0 = true
							term1 = true
						}
					}
				} else if input[PC1][0] == "snd" {
					qu0 = append(qu0, sendValue)
					canContinue = true
					count++
					wait0 = false
				} else {
					term1 = true
				}
			}
			PC1 = newPC
		}

		canContinue = true
	}

	return count
}

func main() {
	in := useful.StringTo2DArray(useful.FileToString("input.txt"))

	fmt.Println("Results:")
	fmt.Printf("Part1: %d\n", puzzle1(in))
	fmt.Printf("Part2: %d\n", puzzle2(in))
}
