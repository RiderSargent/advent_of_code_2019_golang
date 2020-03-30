package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	program := getInput("day02/input")

	fmt.Println("Part 1:", run(program, 12, 2))
	fmt.Println("Part 2:", findInputs(program, 19690720))
}

func findInputs(program []int, target int) int {
	memory := program
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if run(memory, noun, verb) == target {
				return 100*noun + verb
			}
		}
	}
	panic("inputs not found")
}

func run(program []int, noun, verb int) int {
	memory := make([]int, len(program))
	copy(memory, program)

	memory[1] = noun
	memory[2] = verb
	i := 0
	for {
		opcode := memory[i]
		switch opcode {
		case 1:
			a, b, c := memory[i+1], memory[i+2], memory[i+3]
			memory[c] = memory[a] + memory[b]
			i += 4
		case 2:
			a, b, c := memory[i+1], memory[i+2], memory[i+3]
			memory[c] = memory[a] * memory[b]
			i += 4
		case 99:
			return memory[0]
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
