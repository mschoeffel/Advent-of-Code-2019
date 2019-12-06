package Day5

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay5Part1(t *testing.T) {
	data := Utils.ReadFileBySeparatorInt("\\codes.txt", ",")
	expected := 4887191
	actual := Day5Part1(data, "1")
	if actual != expected {
		t.Errorf("Test Day 5 Part 1 Test 1")
	}
}
