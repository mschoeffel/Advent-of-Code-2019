package Day07

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"log"
	"strconv"
)

type Amp struct {
	register []int
	pointer  int
}

func Main() {
	data := Utils.ReadFileBySeparatorInt("\\Day07\\codes.txt", ",")
	fmt.Println("Result part one: " + strconv.Itoa(Day7Part1(data, false)))

	data = Utils.ReadFileBySeparatorInt("\\Day07\\codes.txt", ",")
	fmt.Println("Result part two: " + strconv.Itoa(Day7Part2(data, false)))
}

func Day7Part1(data []int, blockOutput bool) int {
	dataTemp := data
	maxResult := 0
	result := 0
	for a := 0; a <= 4; a++ {
		for b := 0; b <= 4; b++ {
			if b != a {
				for c := 0; c <= 4; c++ {
					if c != a && c != b {
						for d := 0; d <= 4; d++ {
							if d != c && d != b && d != a {
								for e := 0; e <= 4; e++ {
									if e != d && e != c && e != b && e != a {
										phases := []int{a, b, c, d, e}
										result = 0
										for i := range phases {
											copy(dataTemp, data)
											result = runComputer(dataTemp, result, phases[i], blockOutput)
										}
										if result > maxResult {
											maxResult = result
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return maxResult
}

func Day7Part2(data []int, blockOutput bool) int {
	dataTemp := data
	maxResult := 0
	result := 0
	for a := 5; a <= 9; a++ {
		for b := 5; b <= 9; b++ {
			if b != a {
				for c := 5; c <= 9; c++ {
					if c != a && c != b {
						for d := 5; d <= 9; d++ {
							if d != c && d != b && d != a {
								for e := 5; e <= 9; e++ {
									if e != d && e != c && e != b && e != a {
										breaked := false
										phases := []int{a, b, c, d, e}
										result = 0
										maxResultSave := 0
										i := 0
										phasesdone := false
										copy(dataTemp, data)
										ampA := Amp{register: dataTemp, pointer: 0}
										copy(dataTemp, data)
										ampB := Amp{register: dataTemp, pointer: 0}
										copy(dataTemp, data)
										ampC := Amp{register: dataTemp, pointer: 0}
										copy(dataTemp, data)
										ampD := Amp{register: dataTemp, pointer: 0}
										copy(dataTemp, data)
										ampE := Amp{register: dataTemp, pointer: 0}
										amps := []Amp{ampA, ampB, ampC, ampD, ampE}
										for !breaked {
											if phasesdone {
												phases[i] = 0
											}
											result = runComputer2(&amps[i], result, phases[i], blockOutput)
											if result == -1 {
												breaked = true
												break
											}
											if i == len(phases)-1 {
												maxResultSave = result
												i = 0
												phasesdone = true
											} else {
												i++
											}
										}
										if maxResultSave > maxResult {
											maxResult = maxResultSave
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return maxResult
}

func runComputer(code []int, input int, phase int, blockOutput bool) int {
	inputTemp := []int{phase, input}
	return runCode(code, inputTemp, blockOutput)
}

func runCode(code []int, input []int, blockOutput bool) int {
	running := true
	position := 0
	lastoutput := 0
	currentInput := 0

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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				code[position+3] = param1 + param2
			} else {
				code[code[position+3]] = param1 + param2
			}
			position += 4
			break
		case 2:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				code[position+3] = param1 * param2
			} else {
				code[code[position+3]] = param1 * param2
			}
			position += 4
			break
		case 3:
			code[code[position+1]] = input[currentInput]
			currentInput += 1
			position += 2
			break
		case 4:
			param1 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else {
				param1 = code[code[position+1]]
			}
			//Output value of parameter
			if !blockOutput {
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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				value := 0
				if param1 < param2 {
					value = 1
				}
				code[position+3] = value
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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				value := 0
				if param1 == param2 {
					value = 1
				}
				code[position+3] = value
			} else {
				value := 0
				if param1 == param2 {
					value = 1
				}
				code[code[position+3]] = value
			}
			position += 4
			break
		case 99:
			running = false
			break
		}
	}
	return lastoutput
}

func runComputer2(amp *Amp, input int, phase int, blockOutput bool) int {
	var inputTemp []int
	if phase != 0 {
		inputTemp = []int{phase, input}
	} else {
		inputTemp = []int{input}
	}
	return runCode2(amp, inputTemp, blockOutput)
}

//TODO: Not perfect (negatives fail)
func runCode2(amp *Amp, input []int, blockOutput bool) int {
	code := amp.register
	running := true
	position := amp.pointer
	lastoutput := 0
	currentInput := 0

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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				code[position+3] = param1 + param2
			} else {
				code[code[position+3]] = param1 + param2
			}
			position += 4
			break
		case 2:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				code[position+3] = param1 * param2
			} else {
				code[code[position+3]] = param1 * param2
			}
			position += 4
			break
		case 3:
			code[code[position+1]] = input[currentInput]
			currentInput += 1
			position += 2
			break
		case 4:
			param1 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else {
				param1 = code[code[position+1]]
			}
			//Output value of parameter
			if !blockOutput {
				fmt.Println(param1)
			}
			position += 2
			copy(amp.register, code)
			amp.pointer = position
			return param1
		case 5:
			param1 := 0
			param2 := 0
			if modefirst == 1 {
				param1 = code[position+1]
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				value := 0
				if param1 < param2 {
					value = 1
				}
				code[position+3] = value
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
			} else {
				param1 = code[code[position+1]]
			}
			if modesecond == 1 {
				param2 = code[position+2]
			} else {
				param2 = code[code[position+2]]
			}
			if modethird == 1 {
				value := 0
				if param1 == param2 {
					value = 1
				}
				code[position+3] = value
			} else {
				value := 0
				if param1 == param2 {
					value = 1
				}
				code[code[position+3]] = value
			}
			position += 4
			break
		case 99:
			running = false
			return -1
		}
	}
	return lastoutput
}
