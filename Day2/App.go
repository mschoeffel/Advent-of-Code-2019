package Day2

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"strconv"
)

func main() {
	//Question part one
	data := Utils.ReadFileBySeparatorInt("\\Day2\\codes.txt", ",")
	fmt.Println("Result part one: " + strconv.Itoa(Day2Part1(data, 12, 2)))

	// Question part two
	expectedResult := 19690720
	fmt.Println("Result part two: " + strconv.Itoa(Day2Part2(expectedResult)))
}

func Day2Part1(code []int, noun int, verb int) int {
	runCode(code, noun, verb)
	return code[0]
}

func Day2Part2(expected int) int {
	noun, verb := bruteCode(expected)
	return 100*noun + verb
}

// Runs the given register code
func runCode(code []int, noun int, verb int) []int {
	running := true
	position := 0

	code[1] = noun
	code[2] = verb

	for running {
		switch code[position] {
		case 1:
			code[code[position+3]] = code[code[position+1]] + code[code[position+2]]
			position += 4
			break
		case 2:
			code[code[position+3]] = code[code[position+1]] * code[code[position+2]]
			position += 4
			break
		case 99:
			running = false
			break
		}
	}
	return code
}

// Brute forces the logic and returns the noun and verb to get a specific result (Question part two)
func bruteCode(result int) (int, int) {
	noun := 0
	verb := 0
	code := Utils.ReadFileBySeparatorInt("\\Day2\\codes.txt", ",")
	temp := Utils.ReadFileBySeparatorInt("\\Day2\\codes.txt", ",")

	for runCode(temp, noun, verb)[0] != result {
		// To Debug or print the progress:
		//fmt.Println("Noun: " + strconv.Itoa(noun) + " Verb: " + strconv.Itoa(verb) + " Result failed")
		verb++
		if verb > 99 {
			verb = 0
			noun++
		}
		if noun > 99 {
			break
		}
		copy(temp, code)
	}
	return noun, verb
}
