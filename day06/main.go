package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readLines("day06/input")

	orbits := parseOrbits(lines)
	count := countOrbits(orbits)

	fmt.Println("Part 1:", count)
}

func countOrbits(orbits map[string]string) int {
	count := 0
	for _, o := range orbits {
		count += countOrbitsFor(o, orbits)
	}
	return count
}

func countOrbitsFor(currentObject string, orbits map[string]string) int {
	switch currentObject {
	case "COM":
		return 1
	default:
		return 1 + countOrbitsFor(orbits[currentObject], orbits)
	}
}

func parseOrbits(lines []string) map[string]string {
	orbits := make(map[string]string)
	for _, line := range lines {
		objects := strings.Split(line, ")")
		orbits[objects[1]] = objects[0]
	}
	return orbits
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
		log.Fatal(err)
	}
}
