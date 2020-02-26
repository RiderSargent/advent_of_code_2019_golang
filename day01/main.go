package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := getInput()
	fmt.Println("Part 1: ", partOne(input))
	fmt.Println("Part 2: ", partTwo(input))
}

func partOne(modules []int) int {
	fuel := 0

	for _, mass := range modules {
		fuel += fuelFor(mass)
	}

	return fuel
}

func partTwo(modules []int) int {
	fuel := 0

	for _, mass := range modules {
		fuel += recursiveFuelFor(mass)
	}

	return fuel
}

func fuelFor(mass int) int {
	return mass/3 - 2
}

func recursiveFuelFor(mass int) int {
	fuel := mass/3 - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + recursiveFuelFor(fuel)
}

func getInput() []int {
	file, err := os.Open("input")
	panicIf(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		result, err := strconv.Atoi(scanner.Text())
		panicIf(err)
		numbers = append(numbers, result)
	}

	return numbers
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
