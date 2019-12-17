package Day14

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Input struct {
	amount int
	name   string
}

type Process struct {
	name   string
	output int
	input  []Input
}

func Main() {
	processes := Utils.ReadFileByLinesString("//Day14//processes.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day14Part1(processes)))
}

func Day14Part1(processes []string) int {
	var processList []Process
	for i := range processes {
		temp := strings.Split(processes[i], " => ")
		output := strings.Split(temp[1], " ")
		outputAmount, err := strconv.ParseInt(output[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		process := Process{output: int(outputAmount), name: output[1]}
		inputs := strings.Split(temp[0], ", ")
		for j := range inputs {
			input := strings.Split(inputs[j], " ")
			inputAmount, err := strconv.ParseInt(input[0], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			inputObj := Input{amount: int(inputAmount), name: input[1]}
			process.input = append(process.input, inputObj)
		}
		processList = append(processList, process)
	}
	return calcNeeds("FUEL", 1, processList, map[string]int{})

}

func calcNeeds(name string, amount int, processes []Process, excess map[string]int) int {
	if name == "ORE" {
		return amount
	}

	if excess[name] >= amount {
		excess[name] -= amount
		return 0
	}

	if excess[name] > 0 {
		amount -= excess[name]
		excess[name] = 0
	}

	process := Process{}

	for i := range processes {
		if processes[i].name == name {
			process = processes[i]
			break
		}
	}

	number := int(math.Ceil(float64(amount) / float64(process.output)))

	ore := 0

	for i := range process.input {
		ore += calcNeeds(process.input[i].name, process.input[i].amount*number, processes, excess)
	}

	numProduced := number * process.output
	excess[name] += numProduced - amount

	return ore

}
