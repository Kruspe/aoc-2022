package main

import (
	"aoc-2022/shared"
	"fmt"
	"strconv"
	"strings"
)

func solve1(input []string) int {
	result := 0
	for _, l := range input {
		elves := strings.Split(l, ",")
		startFirst, _ := strconv.Atoi(strings.Split(elves[0], "-")[0])
		endFirst, _ := strconv.Atoi(strings.Split(elves[0], "-")[1])
		startSecond, _ := strconv.Atoi(strings.Split(elves[1], "-")[0])
		endSecond, _ := strconv.Atoi(strings.Split(elves[1], "-")[1])
		if startFirst >= startSecond && endFirst <= endSecond {
			result += 1
			continue
		}
		if startSecond >= startFirst && endSecond <= endFirst {
			result += 1

		}
	}
	return result
}

func solve2(input []string) int {
	result := 0
	for _, l := range input {
		elves := strings.Split(l, ",")
		startFirst, _ := strconv.Atoi(strings.Split(elves[0], "-")[0])
		endFirst, _ := strconv.Atoi(strings.Split(elves[0], "-")[1])
		startSecond, _ := strconv.Atoi(strings.Split(elves[1], "-")[0])
		endSecond, _ := strconv.Atoi(strings.Split(elves[1], "-")[1])
		if startFirst >= startSecond && startFirst <= endSecond {
			result += 1
			continue
		}
		if endFirst >= startSecond && endFirst <= endSecond {
			result += 1
			continue
		}
		if startSecond >= startFirst && startSecond <= endFirst {
			result += 1
			continue
		}
		if endSecond >= startFirst && endSecond <= endFirst {
			result += 1
		}
	}
	return result
}

func main() {
	input := shared.ReadInput("04/data.txt")

	fmt.Printf("Solved 1: %d", solve1(input))
	fmt.Println()
	fmt.Printf("Solved 2: %d", solve2(input))
}
