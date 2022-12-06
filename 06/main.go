package main

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

func solve1(input string) int {
	for i := 0; i <= len(input)-4; i++ {
		substring := input[i : i+4]
		isMarker := true
		for _, c := range substring {
			if strings.Count(substring, string(c)) > 1 {
				isMarker = false
				break
			}
		}
		if isMarker {
			return i + 4
		}
	}
	return 0
}

func solve2(input string) int {
	for i := 0; i <= len(input)-14; i++ {
		substring := input[i : i+14]
		isMarker := true
		for _, c := range substring {
			if strings.Count(substring, string(c)) > 1 {
				isMarker = false
				break
			}
		}
		if isMarker {
			return i + 14
		}
	}
	return 0
}

func main() {
	input := shared.ReadInput("06/data.txt")

	fmt.Printf("Solved 1: %d", solve1(input[0]))
	fmt.Println()
	fmt.Printf("Solved 2: %d", solve2(input[0]))
}
