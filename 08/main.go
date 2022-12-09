package main

import (
	"aoc-2022/shared"
	"fmt"
	"strconv"
)

type exercise struct {
	input []string
}

func newExercise(input []string) *exercise {
	return &exercise{
		input: input,
	}
}

func (e *exercise) solve1() int {
	visibleTrees := 0
	for i := 1; i < len(e.input)-1; i++ {
		for j := 1; j < len(e.input[0])-1; j++ {
			treeHeight, _ := strconv.Atoi(string(e.input[i][j]))
			if e.isVisible(j, i, treeHeight) {
				visibleTrees += 1
			}
		}
	}
	return 2*(len(e.input)+len(e.input[0])) - 4 + visibleTrees
}

func (e *exercise) isVisible(x, y, treeHeight int) bool {
	visibleTop, visibleRight, visibleBottom, visibleLeft := true, true, true, true
	for c := x - 1; c >= 0; c-- {
		h, _ := strconv.Atoi(string(e.input[y][c]))
		if h >= treeHeight {
			visibleTop = false
			break
		}
	}
	for c := x + 1; c < len(e.input[0]); c++ {
		h, _ := strconv.Atoi(string(e.input[y][c]))
		if h >= treeHeight {
			visibleBottom = false
			break
		}
	}
	for c := y - 1; c >= 0; c-- {
		h, _ := strconv.Atoi(string(e.input[c][x]))
		if h >= treeHeight {
			visibleRight = false
			break
		}
	}
	for c := y + 1; c < len(e.input); c++ {
		h, _ := strconv.Atoi(string(e.input[c][x]))
		if h >= treeHeight {
			visibleLeft = false
			break
		}
	}
	return visibleTop || visibleLeft || visibleBottom || visibleRight
}

func (e *exercise) solve2() int {
	scenicScore := 0
	for i := 1; i < len(e.input)-1; i++ {
		for j := 1; j < len(e.input[0])-1; j++ {
			treeHeight, _ := strconv.Atoi(string(e.input[i][j]))
			newScore := e.getScenicScore(j, i, treeHeight)
			if newScore > scenicScore {
				scenicScore = newScore
			}
		}
	}
	return scenicScore
}

func (e *exercise) getScenicScore(x, y, treeHeight int) int {
	visibleTop, visibleRight, visibleBottom, visibleLeft := 0, 0, 0, 0
	for c := x - 1; c >= 0; c-- {
		h, _ := strconv.Atoi(string(e.input[y][c]))
		visibleTop += 1
		if h >= treeHeight {
			break
		}
	}
	for c := x + 1; c < len(e.input[0]); c++ {
		h, _ := strconv.Atoi(string(e.input[y][c]))
		visibleBottom += 1
		if h >= treeHeight {
			break
		}
	}
	for c := y - 1; c >= 0; c-- {
		h, _ := strconv.Atoi(string(e.input[c][x]))
		visibleRight += 1
		if h >= treeHeight {
			break
		}
	}
	for c := y + 1; c < len(e.input); c++ {
		h, _ := strconv.Atoi(string(e.input[c][x]))
		visibleLeft += 1
		if h >= treeHeight {
			break
		}
	}
	return visibleTop * visibleLeft * visibleBottom * visibleRight
}

func main() {
	input := shared.ReadInput("08/data.txt")
	e := newExercise(input)

	fmt.Printf("Solved 1: %d", e.solve1())
	fmt.Println()
	fmt.Printf("Solved 2: %d", e.solve2())
}
