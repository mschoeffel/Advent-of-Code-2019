package main

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
)

func main() {
	masses := Utils.ReadFileByLinesInt("\\Day1\\masses.txt")

	resultPartOne := 0
	resultPartTwo := 0

	for i := 0; i < len(masses); i++ {
		resultPartOne += calculateFuel(masses[i])
		resultPartTwo += calculateFuelRecursive(masses[i])
	}

	fmt.Println(resultPartOne)
	fmt.Println(resultPartTwo)
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
