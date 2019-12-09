package Day03

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Wire struct {
	steps []Step
}

type Step struct {
	x int
	y int
}

func Main() {
	wiresRaw := Utils.ReadFileByLinesString("\\Day03\\wires.txt")
	fmt.Println("Day 3: Result part one: " + strconv.Itoa(Day3Part1(wiresRaw)))
	fmt.Println("Day 3: Result part one: " + strconv.Itoa(Day3Part2(wiresRaw)))
}

func Day3Part1(wiresRaw []string) int {
	var wires []Wire
	sourceStep := Step{0, 0}
	currentStep := Step{0, 0}

	// Each wire
	for i := 0; i < len(wiresRaw); i++ {
		pathsRaw := strings.Split(wiresRaw[i], ",")
		var steps []Step

		//Each path
		for j := 0; j < len(pathsRaw); j++ {
			runes := []rune(pathsRaw[j])

			direction := runes[0]
			value, err := strconv.ParseInt(string(runes[1:]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}

			dx := 0
			dy := 0

			switch direction {
			case 'U':
				dy = 1
				break
			case 'D':
				dy = -1
				break
			case 'L':
				dx = -1
				break
			case 'R':
				dx = 1
				break
			}

			// Each Step
			for k := 1; k <= int(value); k++ {
				steps = append(steps, Step{currentStep.x + k*dx, currentStep.y + k*dy})
			}
			currentStep.x = currentStep.x + int(value)*dx
			currentStep.y = currentStep.y + int(value)*dy
		}

		currentStep = Step{0, 0}

		wires = append(wires, Wire{steps: steps})
	}

	var intersections []Step

	for i := range wires[0].steps {
		for j := range wires[1].steps {
			if wires[0].steps[i] == wires[1].steps[j] {
				intersections = append(intersections, wires[1].steps[j])
			}
		}
	}

	minDist := 0
	for i := range intersections {
		dist := stepsDist(sourceStep, intersections[i])
		if minDist == 0 || dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

func Day3Part2(wiresRaw []string) int {
	var wires []Wire
	sourceStep := Step{0, 0}
	currentStep := Step{0, 0}

	// Each wire
	for i := 0; i < len(wiresRaw); i++ {
		pathsRaw := strings.Split(wiresRaw[i], ",")
		var steps []Step

		//Each path
		for j := 0; j < len(pathsRaw); j++ {
			runes := []rune(pathsRaw[j])

			direction := runes[0]
			value, err := strconv.ParseInt(string(runes[1:]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}

			dx := 0
			dy := 0

			switch direction {
			case 'U':
				dy = 1
				break
			case 'D':
				dy = -1
				break
			case 'L':
				dx = -1
				break
			case 'R':
				dx = 1
				break
			}

			// Each Step
			for k := 1; k <= int(value); k++ {
				steps = append(steps, Step{currentStep.x + k*dx, currentStep.y + k*dy})
			}
			currentStep.x = currentStep.x + int(value)*dx
			currentStep.y = currentStep.y + int(value)*dy
		}

		currentStep = Step{0, 0}

		wires = append(wires, Wire{steps: steps})
	}

	var intersections []Step

	for i := range wires[0].steps {
		for j := range wires[1].steps {
			if wires[0].steps[i] == wires[1].steps[j] {
				intersections = append(intersections, wires[1].steps[j])
			}
		}
	}

	minDist := 0
	for i := range intersections {
		dist := stepsDist(sourceStep, intersections[i])
		if minDist == 0 || dist < minDist {
			minDist = dist
		}
	}

	minPath := 0

	for j := 0; j < len(intersections); j++ {
		count := 0
		for i := range wires[0].steps {
			if wires[0].steps[i] == intersections[j] {
				count++
				break
			}
			count++
		}

		for i := range wires[1].steps {
			if wires[1].steps[i] == intersections[j] {
				count++
				break
			}
			count++
		}

		if minPath > count || minPath == 0 {
			minPath = count
		}

	}

	return minPath
}

func stepsDist(stepSrc Step, stepDest Step) int {
	dist := 0
	if stepDest.x > stepSrc.x {
		dist += stepDest.x - stepSrc.x
	} else {
		dist += stepSrc.x - stepDest.x
	}
	if stepDest.y > stepSrc.y {
		dist += stepDest.y - stepSrc.y
	} else {
		dist += stepSrc.y - stepDest.y
	}
	return dist
}
