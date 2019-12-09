package Day02

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay02Part1(t *testing.T) {
	data := Utils.ReadFileBySeparatorInt("\\codes.txt", ",")
	expected := 3267740
	actual := Day2Part1(data, 12, 2)
	if actual != expected {
		t.Errorf("Test Day 02 Part 1 Test 1")
	}
}

func TestDay02Part2(t *testing.T) {
	data := 19690720
	expected := 7870
	actual := Day2Part2("\\codes.txt", data)
	if actual != expected {
		t.Errorf("Test Day 02 Part 2 Test 1")
	}
}
