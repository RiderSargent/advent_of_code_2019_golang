package main

import (
	"fmt"
)

func main() {
	p1Candidates := 0
	p2Candidates := 0

	for n := 134792; n <= 675810; n++ {
		digits := toDigits(n)
		if hasConsecutive(digits) && nonDecreasing(digits) {
			p1Candidates++
			if hasIsolatedConsecutive(digits) {
				p2Candidates++
			}
		}
	}

	fmt.Println("Part 1:", p1Candidates)
	fmt.Println("Part 2:", p2Candidates)
}

func hasConsecutive(n []int) bool {
	for i := 0; i < 5; i++ {
		if n[i] == n[i+1] {
			return true
		}
	}
	return false
}

func nonDecreasing(n []int) bool {
	for i := 0; i < 5; i++ {
		if n[i] > n[i+1] {
			return false
		}
	}
	return true
}

func hasIsolatedConsecutive(n []int) bool {
	if n[0] == n[1] && n[1] != n[2] {
		return true
	}
	if n[1] == n[2] && n[1] != n[0] && n[2] != n[3] {
		return true
	}
	if n[2] == n[3] && n[2] != n[1] && n[3] != n[4] {
		return true
	}
	if n[3] == n[4] && n[3] != n[2] && n[4] != n[5] {
		return true
	}
	if n[4] == n[5] && n[4] != n[3] {
		return true
	}
	return false
}

func toDigits(n int) []int {
	digits := []int{
		(n / 100000) % 10,
		(n / 10000) % 10,
		(n / 1000) % 10,
		(n / 100) % 10,
		(n / 10) % 10,
		n % 10,
	}

	return digits
}
