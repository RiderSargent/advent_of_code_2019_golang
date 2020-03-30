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

	points1 := toPoints(routes[0])
	points2 := toPoints(routes[1])
	intersections := intersections(points1, points2)

	fmt.Println("Part 1:", closestIntersection(intersections))
	fmt.Println("Part 2:", shortestPath(points1, points2, intersections))
}

func intersections(p1, p2 []point) []point {
	p1 = unique(p1)
	p2 = unique(p2)
	points := append(p1, p2...)
	var intersections []point

	pointMap := make(map[point]int)
	for _, p := range points {
		pointMap[p]++
		if pointMap[p] > 1 {
			intersections = append(intersections, p)
		}
	}

	return intersections[1:] // don't inlude (0, 0)
}

func closestIntersection(points []point) int {
	min := math.MaxInt32
	for _, p := range points {
		md := manhattanDistance(p)
		if md < min {
			min = md
		}
	}

	return min
}

func shortestPath(p1, p2, isects []point) int {
	min := math.MaxInt32
	for _, intersection := range isects {
		totalSteps := indexOf(intersection, p1) + indexOf(intersection, p2)
		if totalSteps < min {
			min = totalSteps
		}
	}
	return min
}

func indexOf(p point, list []point) int {
	for i, v := range list {
		if v == p {
			return i
		}
	}
	return -1
}

func pointMap(points []point) map[point]int {
	m := make(map[point]int)
	for _, p := range points {
		m[p]++
	}
	delete(m, point{0, 0})

	return m
}

func toPoints(route string) []point {
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
	return points
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
		panic(err)
	}
}
