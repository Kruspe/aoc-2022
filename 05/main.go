package main

import (
	"aoc-2022/shared"
	"fmt"
	"strconv"
	"strings"
)

func solve1(input []string) string {
	separator, err := shared.GetBlankLineIndex(input)
	if err != nil {
		panic(err)
	}

	initial := input[:separator]
	stacks := createStack(initial)

	instructions := input[separator+1:]
	for _, instruction := range instructions {
		split := strings.Split(instruction, "move ")[1]
		amount, _ := strconv.Atoi(strings.Split(split, " ")[0])
		from := strings.Split(strings.Split(split, " from ")[1], " to ")[0]
		to := strings.Split(split, " to ")[1]

		for i := 0; i < amount; i++ {
			movingObject := stacks[from][0]
			stacks[from] = stacks[from][1:]
			stacks[to] = append([]string{movingObject}, stacks[to]...)
		}
	}

	result := ""
	for i := 1; i <= len(stacks); i++ {
		result += stacks[strconv.Itoa(i)][0]
	}
	return result
}

func solve2(input []string) string {
	separator, err := shared.GetBlankLineIndex(input)
	if err != nil {
		panic(err)
	}

	initial := input[:separator]
	instructions := input[separator+1:]
	stacks := createStack(initial)

	for _, instruction := range instructions {
		split := strings.Split(instruction, "move ")[1]
		amount, _ := strconv.Atoi(strings.Split(split, " ")[0])
		from := strings.Split(strings.Split(split, " from ")[1], " to ")[0]
		to := strings.Split(split, " to ")[1]

		var movingObjects []string
		for i := 0; i < amount; i++ {
			movingObjects = append(movingObjects, stacks[from][0])
			stacks[from] = stacks[from][1:]
		}
		stacks[to] = append(movingObjects, stacks[to]...)
	}

	result := ""
	for i := 1; i <= len(stacks); i++ {
		result += stacks[strconv.Itoa(i)][0]
	}
	return result

}

func createStack(initial []string) map[string][]string {
	stacks := make(map[string][]string)
	stackNumbers := initial[len(initial)-1]
	for _, s := range initial[:len(initial)-1] {
		for i := 1; i <= len(stackNumbers); i += 4 {
			if string(s[i]) != " " {
				stacks[string(stackNumbers[i])] = append(stacks[string(stackNumbers[i])], string(s[i]))
			}
		}
	}
	return stacks
}

func main() {
	input := shared.ReadInput("05/data.txt")

	fmt.Printf("Solved 1: %s", solve1(input))
	fmt.Println()
	fmt.Printf("Solved 2: %s", solve2(input))
}
