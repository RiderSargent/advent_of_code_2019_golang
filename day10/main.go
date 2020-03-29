package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	lines := readLines("day10/input")
	asteroids := getAsteroids(lines)
	visibles := map[point]map[float64]point{}
	max := 0

	for _, a := range asteroids {
		visibles[a] = make(map[float64]point)
		for _, b := range asteroids {
			if a == b {
				continue
			}
			slope := getSlope(a, b)
			visibles[a][slope] = b
		}
		if len(visibles[a]) > max {
			max = len(visibles[a])
		}
	}

	fmt.Println("Part 1:", max)
}

type point struct {
	x, y int
}

func getSlope(a, b point) float64 {
	dy := float64(b.y - a.y)
	dx := float64(b.x - a.x)

	return math.Atan2(dy, dx)
}

func getAsteroids(lines []string) []point {
	height := len(lines)
	width := len(lines[0])

	var points []point

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if lines[row][col] == '#' {
				points = append(points, point{col, row})
			}
		}
	}
	return points
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	catch(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
