package main

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
)

func main() {

	data := Utils.ReadFileBySeparatorInt("\\Day2\\codes.txt", ",")

	running := true
	position := 0

	data[1] = 12
	data[2] = 2

	for running {
		switch data[position] {
		case 1:
			data[data[position+3]] = data[data[position+1]] + data[data[position+2]]
			position += 4
			break
		case 2:
			data[data[position+3]] = data[data[position+1]] * data[data[position+2]]
			position += 4
			break
		case 99:
			running = false
			break
		}
	}
	fmt.Println(data[0])
}
