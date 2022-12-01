package main

import (
	"aoc-2022/shared"
	"fmt"
	"sort"
	"strconv"
)

func solve02(caloriesPerElf []int) int {
	return caloriesPerElf[len(caloriesPerElf)-1] + caloriesPerElf[len(caloriesPerElf)-2] + caloriesPerElf[len(caloriesPerElf)-3]

}

func main() {
	input := shared.ReadInput("01/data.txt")
	caloriesList := getFoodPerElf(input)

	fmt.Printf("Solved 1: %d", caloriesList[len(caloriesList)-1])
	fmt.Println()
	fmt.Printf("Solved 2: %d", solve02(caloriesList))
}

func getFoodPerElf(foodList []string) []int {
	calories := 0
	var caloriesList []int
	for _, c := range foodList {
		if c == "" {
			caloriesList = append(caloriesList, calories)
			calories = 0
			continue
		}

		i, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		calories += i
	}
	caloriesList = append(caloriesList, calories)

	sort.Ints(caloriesList)
	return caloriesList
}
