package Day01

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"strconv"
)

func Main() {
	masses := Utils.ReadFileByLinesInt("\\Day01\\masses.txt")

	fmt.Println("Day 1: Result part one: " + strconv.Itoa(Day1Part1(masses)))
	fmt.Println("Day 1: Result part two: " + strconv.Itoa(Day1Part2(masses)))
}

func Day1Part1(masses []int) int {
	result := 0
	for i := 0; i < len(masses); i++ {
		result += calculateFuel(masses[i])
	}
	return result
}

func Day1Part2(masses []int) int {
	result := 0
	for i := 0; i < len(masses); i++ {
		result += calculateFuelRecursive(masses[i])
	}
	return result
}

//Calculates the fuel (Question part 1)
func calculateFuel(mass int) int {
	fuel := mass/3 - 2
	return fuel
}

// Calculates the fuel recursively (Question part 2)
func calculateFuelRecursive(mass int) int {
	fuel := mass/3 - 2
	if fuel > 0 {
		return fuel + calculateFuelRecursive(fuel)
	} else {
		return 0
	}
}
