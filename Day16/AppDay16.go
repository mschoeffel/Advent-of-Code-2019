package Day16

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"log"
	"strconv"
)

func Main() {
	inputs := Utils.ReadFileByCharInt("//Day16//inputs.txt")
	fmt.Println("Result part one: " + Day16Part1(inputs))

	inputs = Utils.ReadFileByCharInt("//Day16//inputs.txt")
	fmt.Println("Result part two: " + Day16Part2(inputs))

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

func Day16Part2(signal []int) string {
	repeats := 10000
	maxPhases := 100

	position := strconv.Itoa(signal[0]) + strconv.Itoa(signal[1]) + strconv.Itoa(signal[2]) + strconv.Itoa(signal[3]) +
		strconv.Itoa(signal[4]) + strconv.Itoa(signal[5]) + strconv.Itoa(signal[6])

	positionInt, err := strconv.ParseInt(position, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	offset := int(positionInt)

	msg := make([]int, len(signal)*repeats-offset)
	for i := range msg {
		msg[i] = signal[(offset+i)%len(signal)]
	}

	n := len(msg)

	if offset < len(signal)*repeats/2 {
		log.Fatal("offset too great:", offset)
	}

	for ; maxPhases > 0; maxPhases-- {
		sum := 0
		for i := n - 1; i >= 0; i-- {
			sum += msg[i]
			if sum < 0 {
				sum = -sum
			}
			msg[i] = sum % 10
		}
	}

	return strconv.Itoa(msg[0]) + strconv.Itoa(msg[1]) + strconv.Itoa(msg[2]) + strconv.Itoa(msg[3]) +
		strconv.Itoa(msg[4]) + strconv.Itoa(msg[5]) + strconv.Itoa(msg[6]) + strconv.Itoa(msg[7])
}
