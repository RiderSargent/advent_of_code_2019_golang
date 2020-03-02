package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	program := getInput("day05/input")

	fmt.Println("Part 1:", run(program, []int{1}))
	fmt.Println("Part 2:", run(program, []int{5}))
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
		var p1 int
		var p2 int

		switch opcode {
		case 1: // ADD
			if mode1 == 1 {
				p1 = memory[i+1]
			} else {
				p1 = memory[memory[i+1]]
			}

			if mode2 == 1 {
				p2 = memory[i+2]
			} else {
				p2 = memory[memory[i+2]]
			}

			// fmt.Println("   writing", p1+p2, "(", p1, "+", p2, ") to memory[", memory[i+3], "]")
			memory[memory[i+3]] = p1 + p2
			i += 4
		case 2: // MULTIPLY
			if mode1 == 1 {
				p1 = memory[i+1]
			} else {
				p1 = memory[memory[i+1]]
			}

			if mode2 == 1 {
				p2 = memory[i+2]
			} else {
				p2 = memory[memory[i+2]]
			}

			// fmt.Println("   writing", p1*p2, "(", p1, "*", p2, ") to memory[", memory[i+3], "]")
			memory[memory[i+3]] = p1 * p2
			i += 4
		case 3: // INPUT
			// fmt.Println("   writing", input[0], "to memory[", memory[i+1], "]")
			memory[memory[i+1]] = input[0]
			input = input[1:]
			i += 2
		case 4: // OUTPUT
			if mode1 == 1 {
				p1 = memory[i+1]
			} else {
				p1 = memory[memory[i+1]]
			}
			// fmt.Println("   outputing", p1)
			output = append(output, p1)
			i += 2
		case 5: // JUMP-IF-TRUE
			if mode1 == 1 {
				p1 = memory[i+1]
			} else {
				p1 = memory[memory[i+1]]
			}

			if mode2 == 1 {
				p2 = memory[i+2]
			} else {
				p2 = memory[memory[i+2]]
			}

			if p1 > 0 {
				i = p2
			} else {
				i += 3
			}
			// fmt.Println("   set i to", i)
		case 6: // JUMP-IF-FALSE
			if mode1 == 1 {
				p1 = memory[i+1]
			} else {
				p1 = memory[memory[i+1]]
			}

			if mode2 == 1 {
				p2 = memory[i+2]
			} else {
				p2 = memory[memory[i+2]]
			}

			if p1 == 0 {
				i = p2
			} else {
				i += 3
			}
			// fmt.Println("   set i to", i)
		case 7: // LESS THAN
			if mode1 == 1 {
				p1 = memory[i+1]
			} else {
				p1 = memory[memory[i+1]]
			}

			if mode2 == 1 {
				p2 = memory[i+2]
			} else {
				p2 = memory[memory[i+2]]
			}

			if p1 < p2 {
				memory[memory[i+3]] = 1
			} else {
				memory[memory[i+3]] = 0
			}
			// fmt.Println("   set memory[", memory[i+3], "] to", memory[memory[i+3]])
			i += 4
		case 8: // EQUALS
			if mode1 == 1 {
				p1 = memory[i+1]
			} else {
				p1 = memory[memory[i+1]]
			}

			if mode2 == 1 {
				p2 = memory[i+2]
			} else {
				p2 = memory[memory[i+2]]
			}

			if p1 == p2 {
				memory[memory[i+3]] = 1
			} else {
				memory[memory[i+3]] = 0
			}
			// fmt.Println("   set memory[", memory[i+3], "] to", memory[memory[i+3]])
			i += 4
		case 99:
			return output
		default:
			fmt.Println(i, memory)
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
