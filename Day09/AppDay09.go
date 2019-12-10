package Day09

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
	data := Utils.ReadFileBySeparatorInt("\\Day09\\codes.txt", ",")
	fmt.Println("Result part one: " + strconv.Itoa(Day9Part1(data, "0", false)))

	data = Utils.ReadFileBySeparatorInt("\\Day09\\codes.txt", ",")
	fmt.Println("Result part two: " + strconv.Itoa(Day9Part2(data, "0", false)))
}

func Day9Part1(data []int, input string, blockOutput bool) int {
	//Give 1 input
	return runCodePart(data, input, blockOutput)
}

func Day9Part2(data []int, input string, blockOutput bool) int {
	//Give 2 input
	return runCodePart(data, input, blockOutput)
}

func runCodePart(code []int, input string, blockoutput bool) int {
	running := true
	position := 0
	lastoutput := 0
	relativebase := 0

	for running {
		number := code[position]
		modefirst := 0
		modesecond := 0
		modethird := 0
		opcode := 0
		slice := strconv.Itoa(number)
		if len(slice) == 5 {
			temp, err := strconv.ParseInt(string(slice[len(slice)-5]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			modethird = int(temp)

			temp, err = strconv.ParseInt(string(slice[len(slice)-4]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			modesecond = int(temp)

			temp, err = strconv.ParseInt(string(slice[len(slice)-3]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			modefirst = int(temp)

			temp, err = strconv.ParseInt(slice[len(slice)-2:], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			opcode = int(temp)
		} else if len(slice) == 4 {
			temp, err := strconv.ParseInt(string(slice[len(slice)-4]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			modesecond = int(temp)

			temp, err = strconv.ParseInt(string(slice[len(slice)-3]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			modefirst = int(temp)

			temp, err = strconv.ParseInt(slice[len(slice)-2:], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			opcode = int(temp)
		} else if len(slice) == 3 {
			temp, err := strconv.ParseInt(string(slice[len(slice)-3]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			modefirst = int(temp)

			temp, err = strconv.ParseInt(slice[len(slice)-2:], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			opcode = int(temp)
		} else if len(slice) == 2 {
			temp, err := strconv.ParseInt(slice[len(slice)-2:], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			opcode = int(temp)
		} else {
			temp, err := strconv.ParseInt(string(slice[len(slice)-1]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			opcode = int(temp)
		}

		switch opcode {
		case 1:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				positionTemp := code[position+1]
				if positionTemp >= len(code) {
					param1 = 0
				} else {
					param1 = code[positionTemp]
				}
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else if modesecond == 2 {
				param2 = code[relativebase+code[position+2]]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				positionTemp := code[position+3]
				if positionTemp >= len(code) {
					temp := make([]int, positionTemp-len(code)+1)
					code = append(code, temp...)
				}
				code[positionTemp] = param1 + param2
			} else if modethird == 2 {
				positionTemp := code[position+3] + relativebase
				if positionTemp >= len(code) {
					temp := make([]int, positionTemp-len(code)+1)
					code = append(code, temp...)
				}
				code[positionTemp] = param1 + param2
			} else {
				positionTemp := code[position+3]
				if positionTemp >= len(code) {
					temp := make([]int, positionTemp-len(code)+1)
					code = append(code, temp...)
				}
				code[positionTemp] = param1 + param2
			}
			position += 4
			break
		case 2:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else if modesecond == 2 {
				param2 = code[relativebase+code[position+2]]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				positionTemp := position + 3
				if positionTemp >= len(code) {
					temp := make([]int, positionTemp-len(code)+1)
					code = append(code, temp...)
				}
				code[positionTemp] = param1 * param2
			} else if modethird == 2 {
				positionTemp := code[position+3] + relativebase
				if positionTemp >= len(code) {
					temp := make([]int, positionTemp-len(code)+1)
					code = append(code, temp...)
				}
				code[positionTemp] = param1 * param2
			} else {
				positionTemp := code[position+3]
				if positionTemp >= len(code) {
					temp := make([]int, positionTemp-len(code)+1)
					code = append(code, temp...)
				}
				code[positionTemp] = param1 * param2
			}
			position += 4
			break
		case 3:
			//Save input to position given in parameter
			if input != "0" {
				value, err := strconv.ParseInt(input, 10, 32)
				if err != nil {
					log.Fatal(err)
				} else {
					if modefirst == 1 {
						code[position+3] = int(value)
					} else if modethird == 2 {
						code[code[position+3]+relativebase] = int(value)
					} else {
						code[code[position+3]] = int(value)
					}
				}
			} else {
				fmt.Println("Please provide input:")

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
					if modefirst == 1 {
						code[position+3] = int(value)
					} else if modethird == 2 {
						code[code[position+3]+relativebase] = int(value)
					} else {
						code[code[position+3]] = int(value)
					}
				}
			}
			position += 2
			break
		case 4:
			param1 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				param1 = code[code[position+1]]
			}
			if input != "0" {
				lastoutput = param1
				position += 2
				break
			}
			//Output value of parameter
			if !blockoutput {
				fmt.Println(param1)
			}
			lastoutput = param1
			position += 2
			break
		case 5:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else if modesecond == 2 {
				param2 = code[relativebase+code[position+2]]
			} else {
				param2 = code[code[position+2]]
			}
			if param1 != 0 {
				position = param2
			} else {
				position += 3
			}
			break
		case 6:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else if modesecond == 2 {
				param2 = code[relativebase+code[position+2]]
			} else {
				param2 = code[code[position+2]]
			}
			if param1 == 0 {
				position = param2
			} else {
				position += 3
			}
			break
		case 7:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else if modesecond == 2 {
				param2 = code[relativebase+code[position+2]]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				value := 0
				if param1 < param2 {
					value = 1
				}
				code[position+3] = value
			} else if modethird == 2 {
				value := 0
				if param1 < param2 {
					value = 1
				}
				code[code[position+3]+relativebase] = value
			} else {
				value := 0
				if param1 < param2 {
					value = 1
				}
				code[code[position+3]] = value
			}
			position += 4
			break
		case 8:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else if modesecond == 2 {
				param2 = code[relativebase+code[position+2]]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				value := 0
				if param1 == param2 {
					value = 1
				}
				code[position+3] = value
			} else if modethird == 2 {
				value := 0
				if param1 == param2 {
					value = 1
				}
				code[code[position+3]+relativebase] = value
			} else {
				value := 0
				if param1 == param2 {
					value = 1
				}
				positionTemp := code[position+3]
				if positionTemp >= len(code) {
					temp := make([]int, positionTemp-len(code)+1)
					code = append(code, temp...)
				}
				code[positionTemp] = value
			}
			position += 4
			break
		case 9:
			param1 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else if modefirst == 2 {
				param1 = code[relativebase+code[position+1]]
			} else {
				param1 = code[code[position+1]]
			}
			relativebase = relativebase + param1
			position += 2
			break
		case 99:
			running = false
			break
		}
	}
	return lastoutput
}
