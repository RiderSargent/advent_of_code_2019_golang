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
	layers := getLayers(input)
	fmt.Println("Part 1:", part1(layers))
	fmt.Println("Part 2:")
	part2(layers)
}

func part2(layers [][]int) {
	decoded := make([]int, 150)
	numLayers := len(layers)

	for i := 0; i < 150; i++ {
	layerLoop:
		for l := 0; l < numLayers; l++ {
			if layers[l][i] != 2 {
				decoded[i] = layers[l][i]
				break layerLoop
			}
		}
	}

	i := 0
	for row := 0; row < 6; row++ {
		for col := 0; col < 25; col++ {
			if decoded[i] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("█")
			}
			i++
		}
		fmt.Println()
	}
}

func part1(layers [][]int) int {
	min := math.MaxInt32
	minIndex := 0
	for i, l := range layers {
		zeros := count(0, l)
		if zeros < min {
			min = zeros
			minIndex = i
		}
	}
	l := layers[minIndex]
	return count(1, l) * count(2, l)
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

func getLayers(input []int) [][]int {
	var layers [][]int
	for r := 0; r < 100; r++ {
		layers = append(layers, input[r*150:r*150+150])
	}
	return layers
}

func getInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	catch(err)

	input := strings.TrimSpace(string(bytes))

	var result []int

	for _, s := range strings.Split(input, "") {
		result = append(result, toInt(s))
	}
	return result
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
