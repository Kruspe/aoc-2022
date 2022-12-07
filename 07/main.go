package main

import (
	"aoc-2022/shared"
	"fmt"
	"strconv"
	"strings"
)

type directory struct {
	name     string
	children []string
	parent   string
	size     int
}

type exercise struct {
	directories map[string]*directory
}

func newExercise(input []string) *exercise {
	directories := map[string]*directory{
		"/": {
			name:     "/",
			children: nil,
			parent:   "",
			size:     0,
		},
	}
	currentDirectoryName := "/"
	for _, c := range input {
		if string(c[0]) == "$" {
			if string(c[2]) == "c" {
				if c[5:] == "/" {
					currentDirectoryName = "/"
					continue
				}
				if c[5:] == ".." {
					currentDirectoryName = directories[directories[currentDirectoryName].parent].name
					continue
				}

				newDirectoryName := fmt.Sprintf("%s/%s", currentDirectoryName, c[5:])
				directories[newDirectoryName] = &directory{
					name:     newDirectoryName,
					children: nil,
					parent:   currentDirectoryName,
					size:     0,
				}
				directories[currentDirectoryName].children = append(directories[currentDirectoryName].children, newDirectoryName)
				currentDirectoryName = newDirectoryName
			}
		} else {
			if string(c[0]) != "d" {
				size, err := strconv.Atoi(strings.Split(c, " ")[0])
				if err != nil {
					panic(err)
				}
				directories[currentDirectoryName].size += size
			}
		}
	}
	exercise := &exercise{
		directories: directories,
	}
	exercise.sumUpSize("/")
	return exercise
}

func (e *exercise) sumUpSize(name string) {
	directory := e.directories[name]
	for _, c := range directory.children {
		childDirectory := e.directories[c]
		if childDirectory.children == nil {
			directory.size += childDirectory.size
			continue
		} else {
			e.sumUpSize(c)
		}
	}
	if directory.parent != "" {
		e.directories[directory.parent].size += directory.size
	}
}

func solve1(e *exercise) int {
	result := 0
	for _, d := range e.directories {
		if d.size <= 100000 {
			result += d.size
		}
	}
	return result
}

func solve2(e *exercise) int {
	neededSpace := 30000000 - (70000000 - e.directories["/"].size)
	result := e.directories["/"].size
	for _, d := range e.directories {
		if d.size >= neededSpace && d.size < result {
			result = d.size
		}
	}
	return result
}

func main() {
	input := shared.ReadInput("07/data.txt")
	e := newExercise(input)

	fmt.Printf("Solved 1: %d", solve1(e))
	fmt.Println()
	fmt.Printf("Solved 2: %d", solve2(e))
}
