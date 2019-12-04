package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	testDay1Part1()
	testDay1Part2()
	fmt.Println("-----")
	testDay2Part1()
	testDay2Part2()
	fmt.Println("-----")
	testDay3Part1()
	testDay3Part2()
	fmt.Println("-----")
	testDay4Part1()
	testDay4Part2()
}

func writeSuccessMesage(day int, part int, test int) {
	fmt.Println("Test Successful | Day " + strconv.Itoa(day) + " Part " + strconv.Itoa(part) + " Test " + strconv.Itoa(test))
}

func writeErrorMessage(day int, part int, test int, expected string, actual string) {
	log.Println("Test Failed | Day " + strconv.Itoa(day) + " Part " + strconv.Itoa(part) + " Test " + strconv.Itoa(test) + " Expected " + expected + " Actual " + actual)
}

func writeErrorMessageInt(day int, part int, test int, expected int, actual int) {
	writeErrorMessage(day, part, test, strconv.Itoa(expected), strconv.Itoa(actual))
}
