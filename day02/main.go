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
	program[1] = 12
	program[2] = 2
	i := 0
	for {
		opcode := program[i]
		switch opcode {
		case 1:
			a, b, c := program[i+1], program[i+2], program[i+3]
			program[c] = program[a] + program[b]
			i += 4
		case 2:
			a, b, c := program[i+1], program[i+2], program[i+3]
			program[c] = program[a] * program[b]
			i += 4
		case 99:
			return program[0], nil
		default:
			return 0, errors.New("illegal operation")
		}
	}
}

func getInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	panicIf(err)

	input := strings.TrimSpace(string(bytes))

	var program []int

	for _, s := range strings.Split(input, ",") {
		program = append(program, toInt(s))
	}
	return program
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
