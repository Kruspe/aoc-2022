package main

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

func solve1(input []string) int {
	score := 0
	for _, i := range input {
		s := strings.Split(i, " ")

		if s[1] == "X" {
			score += 1
		}
		if s[1] == "Y" {
			score += 2
		}
		if s[1] == "Z" {
			score += 3
		}

		if s[0] == "A" {
			if s[1] == "Y" {
				score += 6
			}
			if s[1] == "X" {
				score += 3
			}
		}
		if s[0] == "B" {
			if s[1] == "Z" {
				score += 6
			}
			if s[1] == "Y" {
				score += 3
			}
		}
		if s[0] == "C" {
			if s[1] == "X" {
				score += 6
			}
			if s[1] == "Z" {
				score += 3
			}
		}
	}

	return score
}

func solve2(input []string) int {
	score := 0
	for _, i := range input {
		s := strings.Split(i, " ")

		if s[1] == "Y" {
			score += 3
		}
		if s[1] == "Z" {
			score += 6
		}

		if s[0] == "A" {
			if s[1] == "X" {
				score += 3
			}
			if s[1] == "Y" {
				score += 1
			}
			if s[1] == "Z" {
				score += 2
			}
		}
		if s[0] == "B" {
			if s[1] == "X" {
				score += 1
			}
			if s[1] == "Y" {
				score += 2
			}
			if s[1] == "Z" {
				score += 3
			}
		}
		if s[0] == "C" {
			if s[1] == "X" {
				score += 2
			}
			if s[1] == "Y" {
				score += 3
			}
			if s[1] == "Z" {
				score += 1
			}
		}
	}

	return score
}

func main() {
	input := shared.ReadInput("02/data.txt")

	fmt.Printf("Solved 1: %d", solve1(input))
	fmt.Println()
	fmt.Printf("Solved 2: %d", solve2(input))
}
