package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := getInput("day08/input")
	var layers [][]int
	for r := 0; r < 100; r++ {
		layers = append(layers, input[r*150:r*150+150])
	}

	min := math.MaxInt32
	minIndex := 0
	for i, l := range layers {
		zeros := count(0, l)
		if zeros < min {
			min = zeros
			minIndex = i
		}
	}

	answer := count(1, layers[minIndex]) * count(2, layers[minIndex])

	fmt.Println("answer:", answer)
}

func count(n int, layer []int) int {
	result := 0
	for _, x := range layer {
		if x == n {
			result++
		}
	}
	return result
}

func getInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	panicIf(err)

	input := strings.TrimSpace(string(bytes))

	var result []int

	for _, s := range strings.Split(input, "") {
		result = append(result, toInt(s))
	}
	return result
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
