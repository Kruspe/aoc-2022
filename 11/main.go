package main

import (
	"aoc-2022/shared"
	"fmt"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation []string
	divisible int
	d1        int
	d2        int
	inspected int
}

func newMonkeys(input []string) map[int]*monkey {
	monkeys := make(map[int]*monkey)
	for i := 0; i < len(input); i = i + 7 {
		m := &monkey{
			inspected: 0,
		}
		for _, item := range strings.Split(strings.Split(input[i+1], ": ")[1], ", ") {
			w, _ := strconv.Atoi(item)
			m.items = append(m.items, w)
		}

		m.operation = strings.Split(strings.Split(input[i+2], "new = old ")[1], " ")

		d, _ := strconv.Atoi(strings.Split(input[i+3], "divisible by ")[1])
		m.divisible = d

		d1, _ := strconv.Atoi(strings.Split(input[i+4], "to monkey ")[1])
		m.d1 = d1
		d2, _ := strconv.Atoi(strings.Split(input[i+5], "to monkey ")[1])
		m.d2 = d2

		monkeys[i/7] = m
	}

	return monkeys
}

func solve1(monkeys map[int]*monkey) int {
	for i := 0; i < 20; i++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]
			m.inspected += len(m.items)
			for _, item := range m.items {
				var operand int
				if m.operation[1] == "old" {
					operand = item
				} else {
					o, _ := strconv.Atoi(m.operation[1])
					operand = o
				}

				var newItem int
				if m.operation[0] == "+" {
					newItem = item + operand
				} else {
					newItem = item * operand
				}
				newItem = newItem / 3

				if newItem%m.divisible == 0 {
					monkeys[m.d1].items = append(monkeys[m.d1].items, newItem)
				} else {
					monkeys[m.d2].items = append(monkeys[m.d2].items, newItem)
				}
			}
			monkeys[i].items = []int{}
		}
	}

	first, second := 0, 0
	for _, m := range monkeys {
		if m.inspected > first {
			second = first
			first = m.inspected
		} else if m.inspected > second {
			second = m.inspected
		}
	}

	return first * second
}

func solve2(monkeys map[int]*monkey) int {
	var divisible []int
	for _, m := range monkeys {
		divisible = append(divisible, m.divisible)
	}
	lcm := lcm(divisible[0], divisible[1], divisible[2:]...)

	for i := 0; i < 10000; i++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]
			m.inspected += len(m.items)
			for _, item := range m.items {
				var operand int
				if m.operation[1] == "old" {
					operand = item
				} else {
					o, _ := strconv.Atoi(m.operation[1])
					operand = o
				}

				var newItem int
				if m.operation[0] == "+" {
					newItem = (item + operand) % lcm
				} else {
					newItem = item * operand % lcm
				}

				if newItem%m.divisible == 0 {
					monkeys[m.d1].items = append(monkeys[m.d1].items, newItem)
				} else {
					monkeys[m.d2].items = append(monkeys[m.d2].items, newItem)
				}
			}
			monkeys[i].items = []int{}
		}
	}

	first, second := 0, 0
	for _, m := range monkeys {
		if m.inspected > first {
			second = first
			first = m.inspected
		} else if m.inspected > second {
			second = m.inspected
		}
	}

	return first * second
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func main() {
	input := shared.ReadInput("11/data.txt")
	m1 := newMonkeys(input)
	m2 := newMonkeys(input)

	fmt.Printf("Solved 1: %d", solve1(m1))
	fmt.Println()
	fmt.Printf("Solved 2: %d", solve2(m2))
}
