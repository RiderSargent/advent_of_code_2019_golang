package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	program := getInput("day05/input")
	input := []int{1}

	fmt.Println("Part 1:", run(program, input))
}

func run(program []int, input []int) []int {
	memory := make([]int, len(program))
	copy(memory, program)
	output := []int{}

	i := 0
	for {
		mode2 := (memory[i] / 1000) % 10
		mode1 := (memory[i] / 100) % 10
		opcode := memory[i] % 100
		var a int
		var b int

		switch opcode {
		case 1: // ADD
			if mode1 == 1 {
				a = memory[i+1]
			} else {
				a = memory[memory[i+1]]
			}

			if mode2 == 1 {
				b = memory[i+2]
			} else {
				b = memory[memory[i+2]]
			}

			c := memory[i+3]
			memory[c] = a + b
			i += 4
		case 2: // MULTIPLY
			if mode1 == 1 {
				a = memory[i+1]
			} else {
				a = memory[memory[i+1]]
			}

			if mode2 == 1 {
				b = memory[i+2]
			} else {
				b = memory[memory[i+2]]
			}

			c := memory[i+3]
			memory[c] = a * b
			i += 4
		case 3: // INPUT
			a := memory[i+1]
			memory[a] = input[0]
			input = input[1:]
			i += 2
		case 4: // OUTPUT
			if mode1 == 1 {
				a = memory[i+1]
			} else {
				a = memory[memory[i+1]]
			}
			output = append(output, a)
			i += 2
		case 99:
			return output
		default:
			panic("illegal operation:")
		}
	}
}

func getInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	catch(err)

	input := strings.TrimSpace(string(bytes))

	var memory []int

	for _, s := range strings.Split(input, ",") {
		memory = append(memory, toInt(s))
	}
	return memory
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	catch(err)
	return n
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
