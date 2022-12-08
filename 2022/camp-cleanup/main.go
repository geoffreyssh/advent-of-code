package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")

	defer file.Close()

	contained_pair_counter := 0
	overlapping_pair_counter := 0

	var elves_listed []string
	var first_elf_bounds []string
	var second_elf_bounds []string

	var first_elf_lower_bound int
	var first_elf_upper_bound int
	var second_elf_lower_bound int
	var second_elf_upper_bound int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elves_listed = strings.Split(scanner.Text(), ",")
		first_elf_bounds = strings.Split(elves_listed[0], "-")
		second_elf_bounds = strings.Split(elves_listed[1], "-")

		first_elf_lower_bound, _ = strconv.Atoi(first_elf_bounds[0])
		first_elf_upper_bound, _ = strconv.Atoi(first_elf_bounds[1])
		second_elf_lower_bound, _ = strconv.Atoi(second_elf_bounds[0])
		second_elf_upper_bound, _ = strconv.Atoi(second_elf_bounds[1])

		if calculate_if_contained_pair(first_elf_lower_bound, first_elf_upper_bound, second_elf_lower_bound, second_elf_upper_bound) == true {
			contained_pair_counter = contained_pair_counter + 1
		}

		if calculate_if_overlapping_pair(first_elf_lower_bound, first_elf_upper_bound, second_elf_lower_bound, second_elf_upper_bound) == true {
			overlapping_pair_counter = overlapping_pair_counter + 1
		}
	}

	fmt.Printf("Contained pair counter is %v\n", contained_pair_counter)
	fmt.Printf("Overlapping pair counter is %v\n", overlapping_pair_counter)
}

func calculate_if_contained_pair(first_elf_lower_bound int, first_elf_upper_bound int, second_elf_lower_bound int, second_elf_upper_bound int) bool {

	elf_tasks_contained := false

	if first_elf_lower_bound <= second_elf_lower_bound && first_elf_upper_bound >= second_elf_upper_bound {
		elf_tasks_contained = true
	} else if first_elf_lower_bound >= second_elf_lower_bound && first_elf_upper_bound <= second_elf_upper_bound {
		elf_tasks_contained = true
	}

	return elf_tasks_contained
}

func calculate_if_overlapping_pair(first_elf_lower_bound int, first_elf_upper_bound int, second_elf_lower_bound int, second_elf_upper_bound int) bool {

	elf_tasks_overlap := true

	if first_elf_upper_bound < second_elf_lower_bound {
		elf_tasks_overlap = false
	} else if first_elf_lower_bound > second_elf_upper_bound {
		elf_tasks_overlap = false
	}

	return elf_tasks_overlap
}
