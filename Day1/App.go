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

	result := 0

	for i := 0; i < len(masses); i++ {
		result += calculateFuel(masses[i])
	}

	fmt.Println(result)

}

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

func calculateFuel(mass int) int {
	fuel := mass/3 - 2
	return fuel
}
