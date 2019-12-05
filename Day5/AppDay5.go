package Day5

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"strconv"
)

func Main() {
	//Question part one
	data := Utils.ReadFileBySeparatorInt("\\Day5\\codes.txt", ",")
	fmt.Println("Result part one: " + strconv.Itoa(Day5Part1(data, 12, 2)))

}

func Day5Part1(code []int, noun int, verb int) int {
	runCode(code, noun, verb)
	return code[0]
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
