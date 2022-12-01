package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input_1")

	inventories_string := string(b)

	var inventories = list_inventories(inventories_string)
	var highest_calories = get_highest_calories(inventories)
	var three_highest_calories = get_three_highest_calories(inventories)

	fmt.Printf("Highest calorie count is %v\nThree highest calorie counts add to %v\n", highest_calories, three_highest_calories)
}

func list_inventories(input string) []string {
	inventories := strings.Split(input, "\n\n")
	return inventories
}

func get_highest_calories(inventories []string) int {
	var highest_calories = 0

	for _, inventory := range inventories {
		items := strings.Split(inventory, "\n")
		var calorie_total = 0
		for _, item_calories_string := range items {
			item_calories_int, _ := strconv.Atoi(item_calories_string)
			calorie_total = calorie_total + item_calories_int
		}
		if calorie_total > highest_calories {
			highest_calories = calorie_total
		}
	}

	return highest_calories
}

func get_three_highest_calories(inventories []string) int {
	var highest_calories = 0
	var second_highest_calories = 0
	var third_highest_calories = 0

	for _, inventory := range inventories {
		items := strings.Split(inventory, "\n")
		var calorie_total = 0
		for _, item_calories_string := range items {
			item_calories_int, _ := strconv.Atoi(item_calories_string)
			calorie_total = calorie_total + item_calories_int
		}
		if calorie_total >= highest_calories {
			third_highest_calories = second_highest_calories
			second_highest_calories = highest_calories
			highest_calories = calorie_total
		} else if calorie_total >= second_highest_calories {
			third_highest_calories = second_highest_calories
			second_highest_calories = calorie_total
		} else if calorie_total > third_highest_calories {
			third_highest_calories = calorie_total
		}
	}

	return highest_calories + second_highest_calories + third_highest_calories
}
