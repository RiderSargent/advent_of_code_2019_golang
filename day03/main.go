package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	routes := getInput("day03/input")

	fmt.Println("Part 1:", getClosestIntersection(routes))
}

func getClosestIntersection(routes []string) int {
	points1 := getUniquePoints(routes[0])
	points2 := getUniquePoints(routes[1])
	points := append(points1, points2...)

	pointMap := pointMap(points)

	min := math.MaxInt32
	for point, count := range pointMap {
		if count > 1 {
			d := manhattanDistance(point)
			if d < min {
				min = d
			}
		}
	}

	return min
}

func pointMap(points []point) map[point]int {
	m := make(map[point]int)
	for _, p := range points {
		m[p]++
	}
	delete(m, point{0, 0})

	return m
}

func getUniquePoints(route string) []point {
	points := make([]point, 1, 155000)
	for _, segment := range strings.Split(route, ",") {
		last := points[len(points)-1]
		direction, distance := parseSegment(segment)

		switch direction {
		case "U":
			for i := 1; i <= distance; i++ {
				points = append(points, point{last.x, last.y - i})
			}
		case "D":
			for i := 1; i <= distance; i++ {
				points = append(points, point{last.x, last.y + i})
			}
		case "L":
			for i := 1; i <= distance; i++ {
				points = append(points, point{last.x - i, last.y})
			}
		case "R":
			for i := 1; i <= distance; i++ {
				points = append(points, point{last.x + i, last.y})
			}
		default:
			panic("illegal direction")
		}
	}
	return unique(points)
}

func unique(points []point) []point {
	seen := map[point]bool{}
	result := []point{}

	for v := range points {
		if seen[points[v]] != true {
			seen[points[v]] = true
			result = append(result, points[v])
		}
	}
	return result
}

func parseSegment(s string) (string, int) {
	for i := range s {
		if i > 0 {
			distanceStr := s[i:]
			return s[:i], toInt(distanceStr)
		}
	}
	return "", 0
}

func getInput(filename string) []string {
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

type point struct {
	x, y int
}

func manhattanDistance(p point) int {
	return abs(p.x) + abs(p.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	catch(err)
	return n
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
