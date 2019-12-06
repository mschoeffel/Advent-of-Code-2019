package Day5

import (
	"Advent-of-Code-2019/Utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		case 3:
			//Save input to position given in parameter
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			} else {
				text = strings.Replace(text, "\n", "", -1)
				value, err := strconv.ParseInt(text, 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				code[code[position+1]] = int(value)
			}
			position += 2
			break
		case 4:
			//Output value of parameter
			fmt.Println(code[code[position+1]])
			position += 2
			break
		case 99:
			running = false
			break
		}
	}
	return code
}
