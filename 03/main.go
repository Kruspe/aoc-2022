package main

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
	"unicode"
)

func solve1(input []string) int {
	score := 0
	for _, r := range input {
		firstHalf := r[:len(r)/2]
		secondHalf := r[len(r)/2:]

		used := ""
		for _, c := range firstHalf {
			if strings.Contains(secondHalf, string(c)) {
				if !strings.Contains(used, string(c)) {
					used += string(c)
					if unicode.IsUpper(c) {
						score = score + int(c) - 38
					} else {
						score = score + int(c) - 96
					}
				}
			}
		}
	}
	return score
}

func solve2(input []string) int {
	score := 0

	for i := 1; i <= len(input)/3; i++ {
		team := input[3*(i-1) : 3*i]
		for _, c := range team[0] {
			if strings.Contains(team[1], string(c)) && strings.Contains(team[2], string(c)) {
				if unicode.IsUpper(c) {
					score = score + int(c) - 38
				} else {
					score = score + int(c) - 96
				}
				break
			}
		}
	}
	return score
}

func main() {
	input := shared.ReadInput("03/data.txt")

	fmt.Printf("Solved 1: %d", solve1(input))
	fmt.Println()
	fmt.Printf("Solved 2: %d", solve2(input))
}
