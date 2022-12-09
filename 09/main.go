package main

import (
	"aoc-2022/shared"
	"fmt"
	"strconv"
)

type coordinates struct {
	x int
	y int
}

type exercise struct {
	knotPositions map[int]coordinates
	visited       map[string]bool
}

func newExercise(knots int) *exercise {
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d", 0, 0)] = true
	knotPositions := make(map[int]coordinates)
	for i := 0; i < knots; i++ {
		knotPositions[i] = coordinates{
			x: 0,
			y: 0,
		}
	}
	return &exercise{
		knotPositions: knotPositions,
		visited:       visited,
	}
}

func (e *exercise) solve1(input []string) int {
	for _, i := range input {
		times, _ := strconv.Atoi(i[2:])
		direction := string(i[0])
		for i := 1; i <= times; i++ {
			e.moveHead(direction)
			e.moveTail(0)
			e.visited[fmt.Sprintf("%d,%d", e.knotPositions[1].x, e.knotPositions[1].y)] = true
		}
	}

	return len(e.visited)
}

func (e *exercise) moveTail(head int) {
	if e.knotPositions[head].x-2 == e.knotPositions[head+1].x && e.knotPositions[head].y == e.knotPositions[head+1].y {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x + 1,
			y: e.knotPositions[head+1].y,
		}
	} else if e.knotPositions[head].x+2 == e.knotPositions[head+1].x && e.knotPositions[head].y == e.knotPositions[head+1].y {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x - 1,
			y: e.knotPositions[head+1].y,
		}
	} else if e.knotPositions[head].y-2 == e.knotPositions[head+1].y && e.knotPositions[head].x == e.knotPositions[head+1].x {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x,
			y: e.knotPositions[head+1].y + 1,
		}
	} else if e.knotPositions[head].y+2 == e.knotPositions[head+1].y && e.knotPositions[head].x == e.knotPositions[head+1].x {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x,
			y: e.knotPositions[head+1].y - 1,
		}
	} else if e.knotPositions[head].x-2 == e.knotPositions[head+1].x && e.knotPositions[head].y+1 == e.knotPositions[head+1].y || e.knotPositions[head].y+2 == e.knotPositions[head+1].y && e.knotPositions[head].x-1 == e.knotPositions[head+1].x || e.knotPositions[head].y+2 == e.knotPositions[head+1].y && e.knotPositions[head].x-2 == e.knotPositions[head+1].x {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x + 1,
			y: e.knotPositions[head+1].y - 1,
		}
	} else if e.knotPositions[head].x+2 == e.knotPositions[head+1].x && e.knotPositions[head].y+1 == e.knotPositions[head+1].y || e.knotPositions[head].y+2 == e.knotPositions[head+1].y && e.knotPositions[head].x+1 == e.knotPositions[head+1].x || e.knotPositions[head].y+2 == e.knotPositions[head+1].y && e.knotPositions[head].x+2 == e.knotPositions[head+1].x {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x - 1,
			y: e.knotPositions[head+1].y - 1,
		}
	} else if e.knotPositions[head].x+2 == e.knotPositions[head+1].x && e.knotPositions[head].y-1 == e.knotPositions[head+1].y || e.knotPositions[head].y-2 == e.knotPositions[head+1].y && e.knotPositions[head].x+1 == e.knotPositions[head+1].x || e.knotPositions[head].y-2 == e.knotPositions[head+1].y && e.knotPositions[head].x+2 == e.knotPositions[head+1].x {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x - 1,
			y: e.knotPositions[head+1].y + 1,
		}
	} else if e.knotPositions[head].x-2 == e.knotPositions[head+1].x && e.knotPositions[head].y-1 == e.knotPositions[head+1].y || e.knotPositions[head].y-2 == e.knotPositions[head+1].y && e.knotPositions[head].x-1 == e.knotPositions[head+1].x || e.knotPositions[head].y-2 == e.knotPositions[head+1].y && e.knotPositions[head].x-2 == e.knotPositions[head+1].x {
		e.knotPositions[head+1] = coordinates{
			x: e.knotPositions[head+1].x + 1,
			y: e.knotPositions[head+1].y + 1,
		}
	}
}

func (e *exercise) moveHead(direction string) {
	switch direction {
	case "U":
		e.knotPositions[0] = coordinates{
			x: e.knotPositions[0].x,
			y: e.knotPositions[0].y + 1,
		}
	case "R":
		e.knotPositions[0] = coordinates{
			x: e.knotPositions[0].x + 1,
			y: e.knotPositions[0].y,
		}
	case "D":
		e.knotPositions[0] = coordinates{
			x: e.knotPositions[0].x,
			y: e.knotPositions[0].y - 1,
		}
	case "L":
		e.knotPositions[0] = coordinates{
			x: e.knotPositions[0].x - 1,
			y: e.knotPositions[0].y,
		}
	}
}

func (e *exercise) solve2(input []string) int {
	for _, i := range input {
		times, _ := strconv.Atoi(i[2:])
		direction := string(i[0])
		for i := 1; i <= times; i++ {
			e.moveHead(direction)
			for tail := 1; tail < 10; tail++ {
				lastPositionX := e.knotPositions[tail].x
				lastPositionY := e.knotPositions[tail].y
				e.moveTail(tail - 1)
				if lastPositionX == e.knotPositions[tail].x && lastPositionY == e.knotPositions[tail].y {
					break
				}
				if tail == 9 {
					e.visited[fmt.Sprintf("%d,%d", e.knotPositions[tail].x, e.knotPositions[tail].y)] = true
				}
			}
		}
	}
	return len(e.visited)
}

func main() {
	input := shared.ReadInput("09/data.txt")
	e1 := newExercise(2)
	e2 := newExercise(10)

	fmt.Printf("Solved 1: %d", e1.solve1(input))
	fmt.Println()
	fmt.Printf("Solved 2: %d", e2.solve2(input))
}
