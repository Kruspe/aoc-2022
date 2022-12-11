package main

import (
	"aoc-2022/shared"
	"fmt"
	"strconv"
	"strings"
)

type exercise struct {
	cycle  int
	x      int
	result int
}

func newExercise() *exercise {
	return &exercise{
		cycle:  0,
		x:      1,
		result: 0,
	}
}

func (e *exercise) solve1(input []string) int {
	for _, i := range input {
		e.resolveInstruction(i)
	}
	return e.result
}

func (e *exercise) resolveInstruction(instruction string) {
	e.cycle += 1
	if e.cycle == 20 || e.cycle == 60 || e.cycle == 100 || e.cycle == 140 || e.cycle == 180 || e.cycle == 220 {
		e.result += e.cycle * e.x
	}
	if instruction != "noop" {
		y, _ := strconv.Atoi(instruction[5:])
		e.cycle += 1
		if e.cycle == 20 || e.cycle == 60 || e.cycle == 100 || e.cycle == 140 || e.cycle == 180 || e.cycle == 220 {
			e.result += e.cycle * e.x
		}
		e.x += y
	}
}

type drawing struct {
	image map[int][]string
	cycle int
	x     int
}

func newDrawing() *drawing {
	image := make(map[int][]string)
	for i := 0; i < 6; i++ {
		image[i] = strings.Split(strings.Repeat(".", 40), "")
	}
	return &drawing{
		image: image,
		cycle: 0,
		x:     1,
	}
}

func (d *drawing) draw() {
	if d.x-1 <= d.cycle%40 && d.x+1 >= d.cycle%40 {
		d.image[(d.cycle / 40)][d.cycle%40] = "#"
	}
}

func (d *drawing) execute(instruction string) {
	d.draw()
	d.cycle += 1
	if instruction != "noop" {
		d.draw()
		d.cycle += 1
		y, _ := strconv.Atoi(instruction[5:])
		d.x += y
	}
}

func (d *drawing) solve2(input []string) {
	for _, i := range input {
		d.execute(i)
	}
	for i := 0; i < 6; i++ {
		fmt.Println(d.image[i])
	}
}

func main() {
	input := shared.ReadInput("10/data.txt")
	e := newExercise()
	d := newDrawing()

	fmt.Printf("Solved 1: %d", e.solve1(input))
	fmt.Println()
	fmt.Println("Solved 2:")
	d.solve2(input)
}
