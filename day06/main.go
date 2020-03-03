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

	fmt.Println("Part 1:", countOrbits(orbits))

	fmt.Println("Part 2:", countOrbitalTransfers("YOU", "SAN", orbits))
}

func countOrbitalTransfers(o1, o2 string, orbits map[string]string) int {
	p1 := getParents(o1, orbits)
	p2 := getParents(o2, orbits)
	unique1 := p1
	unique2 := p2
	for i := 0; i < len(p1) && i < len(p2); i++ {
		if p1[i] == p2[i] {
			unique1 = p1[i+1:]
			unique2 = p2[i+1:]
		}
	}
	return len(unique1) + len(unique2)
}

func getParents(currentObject string, orbits map[string]string) []string {
	parent := orbits[currentObject]
	switch parent {
	case "COM":
		return []string{"COM"}
	default:
		return append(getParents(parent, orbits), parent)
	}
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
