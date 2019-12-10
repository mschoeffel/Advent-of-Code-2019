package Day08

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay08Part1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\imagedata.txt")
	expected := 1088
	actual := Day8Part1(data[0])
	if actual != expected {
		t.Errorf("Test Day 08 Part 1 Test 1")
	}
}
