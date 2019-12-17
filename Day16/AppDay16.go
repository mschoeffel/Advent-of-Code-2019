package Day16

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"strconv"
)

func Main() {
	inputs := Utils.ReadFileByCharInt("//Day16//inputs.txt")
	fmt.Println("Result part one: " + Day16Part1(inputs))
}

func Day16Part1(inputs []int) string {
	pattern := []int{0, 1, 0, -1}
	pointerPatternMax := len(pattern)
	maxPhase := 100
	characters := len(inputs)

	for phase := 1; phase <= maxPhase; phase++ {

		rounds := make([]int, characters)

		for round := range rounds {
			result := 0
			counter := 0
			pointerPattern := 1
			for character := 0; character < characters; character++ {
				if character >= round {
					counter++
					currentPatternNumber := pattern[pointerPattern]
					temp := inputs[character] * currentPatternNumber
					result += temp
					if pointerPattern >= pointerPatternMax-1 && counter%(round+1) == 0 {
						pointerPattern = 0
					} else if counter%(round+1) == 0 {
						pointerPattern++
						counter = 0
					}
				}
			}
			if result < 0 {
				result = result * -1
			}
			rounds[round] = result % 10
		}
		inputs = rounds
	}

	return strconv.Itoa(inputs[0]) + strconv.Itoa(inputs[1]) + strconv.Itoa(inputs[2]) + strconv.Itoa(inputs[3]) +
		strconv.Itoa(inputs[4]) + strconv.Itoa(inputs[5]) + strconv.Itoa(inputs[6]) + strconv.Itoa(inputs[7])
}
