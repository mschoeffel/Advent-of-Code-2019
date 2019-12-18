package Day04

import "testing"

func TestDay04Part1Test1(t *testing.T) {
	min := 246515
	max := 739105
	expected := 1048
	actual := Day4Part1(min, max)
	if actual != expected {
		t.Errorf("Test Day 04 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay04Part2Test1(t *testing.T) {
	min := 246515
	max := 739105
	expected := 677
	actual := Day4Part2(min, max)
	if actual != expected {
		t.Errorf("Test Day 04 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
