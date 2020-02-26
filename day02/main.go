package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	program := getInput("day02/input")
	result, err := run(program)
	panicIf(err)
	fmt.Println("Part 1:", result)
}

func run(program []int) (int, error) {
	memory := program
	memory[1] = 12
	memory[2] = 2
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
			return memory[0], nil
		default:
			return 0, errors.New("illegal operation")
		}
	}
}

func getInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	panicIf(err)

	input := strings.TrimSpace(string(bytes))

	var memory []int

	for _, s := range strings.Split(input, ",") {
		memory = append(memory, toInt(s))
	}
	return memory
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	panicIf(err)
	return n
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
