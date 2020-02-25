package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part 1: 3303995
// Part 2: 4953118
func main() {
	modules := getInput()
	totalFuel := 0

	for _, mass := range modules {
		totalFuel += mass/3 - 2
	}

	fmt.Print("Part 1: ")
	fmt.Println(totalFuel)
}

func getInput() []int {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		result, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, result)
	}

	return numbers
}
