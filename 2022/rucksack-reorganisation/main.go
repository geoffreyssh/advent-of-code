package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")

	defer file.Close()

	total_priority_misplaced_items := 0
	var rucksack_list []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total_priority_misplaced_items = total_priority_misplaced_items + assign_points_to_letter(get_duplicated_item_within_single_line(scanner.Text()))
		rucksack_list = append(rucksack_list, scanner.Text())
	}

	fmt.Printf("Total priority of misplaced items is %v", total_priority_misplaced_items)

	fmt.Println(rucksack_list)

	total_priority_badges := 0

	for i := 0; i < 100; i++ {
		total_priority_badges = total_priority_badges + assign_points_to_letter(get_duplicated_item_across_three_lines(rucksack_list[3*i:(3*i)+3]))
	}

	fmt.Printf("Total priority of badges is %v", total_priority_badges)
}

func get_duplicated_item_within_single_line(rucksack string) string {
	first_half, second_half := get_two_halves_of_rucksack(rucksack)

	common_item := ""

	for i := range first_half {
		for j := range second_half {
			if string(first_half[i]) == string(second_half[j]) {
				common_item = string(first_half[i])
				break
			}
		}
	}

	return common_item
}

func get_duplicated_item_across_three_lines(rucksack []string) string {
	first_rucksack := []rune(rucksack[0])
	second_rucksack := []rune(rucksack[1])
	third_rucksack := []rune(rucksack[2])

	common_item := ""

	for i := range first_rucksack {
		for j := range second_rucksack {
			for k := range third_rucksack {
				if string(first_rucksack[i]) == string(second_rucksack[j]) && string(second_rucksack[j]) == string(third_rucksack[k]) {
					common_item = string(first_rucksack[i])
					break
				}
			}
		}
	}

	return common_item
}

func get_two_halves_of_rucksack(rucksack string) ([]rune, []rune) {
	rucksack_as_runes := []rune(rucksack)

	first_half := rucksack_as_runes[:len(rucksack_as_runes)/2]
	second_half := rucksack_as_runes[(len(rucksack_as_runes) / 2):]

	return first_half, second_half
}

func assign_points_to_letter(letter string) int {
	priorities := make(map[string]int)
	priorities["a"] = 1
	priorities["b"] = 2
	priorities["c"] = 3
	priorities["d"] = 4
	priorities["e"] = 5
	priorities["f"] = 6
	priorities["g"] = 7
	priorities["h"] = 8
	priorities["i"] = 9
	priorities["j"] = 10
	priorities["k"] = 11
	priorities["l"] = 12
	priorities["m"] = 13
	priorities["n"] = 14
	priorities["o"] = 15
	priorities["p"] = 16
	priorities["q"] = 17
	priorities["r"] = 18
	priorities["s"] = 19
	priorities["t"] = 20
	priorities["u"] = 21
	priorities["v"] = 22
	priorities["w"] = 23
	priorities["x"] = 24
	priorities["y"] = 25
	priorities["z"] = 26
	priorities["A"] = 27
	priorities["B"] = 28
	priorities["C"] = 29
	priorities["D"] = 30
	priorities["E"] = 31
	priorities["F"] = 32
	priorities["G"] = 33
	priorities["H"] = 34
	priorities["I"] = 35
	priorities["J"] = 36
	priorities["K"] = 37
	priorities["L"] = 38
	priorities["M"] = 39
	priorities["N"] = 40
	priorities["O"] = 41
	priorities["P"] = 42
	priorities["Q"] = 43
	priorities["R"] = 44
	priorities["S"] = 45
	priorities["T"] = 46
	priorities["U"] = 47
	priorities["V"] = 48
	priorities["W"] = 49
	priorities["X"] = 50
	priorities["Y"] = 51
	priorities["Z"] = 52

	points := priorities[letter]

	return points
}
