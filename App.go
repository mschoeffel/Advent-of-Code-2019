package main

import (
	"Advent-of-Code-2019/Day01"
	"Advent-of-Code-2019/Day02"
	"Advent-of-Code-2019/Day03"
	"Advent-of-Code-2019/Day04"
	"Advent-of-Code-2019/Day05"
	"Advent-of-Code-2019/Day06"
	"Advent-of-Code-2019/Day07"
	"Advent-of-Code-2019/Day08"
	"Advent-of-Code-2019/Day09"
	"Advent-of-Code-2019/Day10"
	"Advent-of-Code-2019/Day11"
	"Advent-of-Code-2019/Day12"
	"Advent-of-Code-2019/Day13"
	"Advent-of-Code-2019/Day14"
	"Advent-of-Code-2019/Day15"
	"Advent-of-Code-2019/Day16"
	"Advent-of-Code-2019/Day17"
	"Advent-of-Code-2019/Day18"
	"Advent-of-Code-2019/Day19"
	"Advent-of-Code-2019/Day20"
	"Advent-of-Code-2019/Day21"
	"Advent-of-Code-2019/Day22"
	"Advent-of-Code-2019/Day23"
	"Advent-of-Code-2019/Day24"
	"Advent-of-Code-2019/Day25"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of day: ")
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")
	number, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		fmt.Println("Invalid input.")
		log.Fatal(err)
	} else {
		switch number {
		case 1:
			Day01.Main()
			break
		case 2:
			Day02.Main()
			break
		case 3:
			Day03.Main()
			break
		case 4:
			Day04.Main()
			break
		case 5:
			Day05.Main()
			break
		case 6:
			Day06.Main()
			break
		case 7:
			Day07.Main()
			break
		case 8:
			Day08.Main()
			break
		case 9:
			Day09.Main()
			break
		case 10:
			Day10.Main()
			break
		case 11:
			Day11.Main()
			break
		case 12:
			Day12.Main()
			break
		case 13:
			Day13.Main()
			break
		case 14:
			Day14.Main()
			break
		case 15:
			Day15.Main()
			break
		case 16:
			Day16.Main()
			break
		case 17:
			Day17.Main()
			break
		case 18:
			Day18.Main()
			break
		case 19:
			Day19.Main()
			break
		case 20:
			Day20.Main()
			break
		case 21:
			Day21.Main()
			break
		case 22:
			Day22.Main()
			break
		case 23:
			Day23.Main()
			break
		case 24:
			Day24.Main()
			break
		case 25:
			Day25.Main()
			break
		}
	}
}
