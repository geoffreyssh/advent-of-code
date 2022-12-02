package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./input")

	rounds := string(b)
	rounds_list := strings.Split(rounds, "\n")

	var total_points_first_interpretation = 0
	var total_points_second_interpretation = 0

	for _, single_round := range rounds_list {
		if single_round != "" {
			total_points_first_interpretation = total_points_first_interpretation + points_from_single_round_first_interpretation(single_round)
			total_points_second_interpretation = total_points_second_interpretation + points_from_single_round_second_interpretation(single_round)
		}
	}

	fmt.Printf("Total points from the first interpretation is %v\n", total_points_first_interpretation)
	fmt.Printf("Total points from the second interpretation is %v\n", total_points_second_interpretation)
}

func points_from_single_round_first_interpretation(strategy string) int {
	moves := strings.Split(strategy, " ")
	var opponent = moves[0]
	var myself = moves[1]
	var points int

	switch myself {
	case "X":
		points = 1
		switch opponent {
		case "A":
			points = points + 3
		case "B":
		case "C":
			points = points + 6
		}
	case "Y":
		points = 2
		switch opponent {
		case "A":
			points = points + 6
		case "B":
			points = points + 3
		case "C":
		}
	case "Z":
		points = 3
		switch opponent {
		case "A":
		case "B":
			points = points + 6
		case "C":
			points = points + 3
		}
	}

	return points
}

func points_from_single_round_second_interpretation(strategy string) int {
	moves := strings.Split(strategy, " ")
	var opponent = moves[0]
	var myself = moves[1]
	var points int

	switch myself {
	case "X":
		points = 0
		switch opponent {
		case "A":
			points = points + 3
		case "B":
			points = points + 1
		case "C":
			points = points + 2
		}
	case "Y":
		points = 3
		switch opponent {
		case "A":
			points = points + 1
		case "B":
			points = points + 2
		case "C":
			points = points + 3
		}
	case "Z":
		points = 6
		switch opponent {
		case "A":
			points = points + 2
		case "B":
			points = points + 3
		case "C":
			points = points + 1
		}
	}

	return points
}
