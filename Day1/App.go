package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	masses := readMassInput()

	resultPartOne := 0
	resultPartTwo := 0

	for i := 0; i < len(masses); i++ {
		resultPartOne += calculateFuel(masses[i])
		resultPartTwo += calculateFuelRecursive(masses[i])
	}

	fmt.Println(resultPartOne)
	fmt.Println(resultPartTwo)
}

// Reads the input data from a txt file
func readMassInput() []int {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dir + "\\Day1\\masses.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var l []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		} else {
			l = append(l, int(i))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return l
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
