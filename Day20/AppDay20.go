package Day20

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"strconv"
)

func Main() {
	maze := Utils.ReadFileByLinesString("//Day20//maze.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day20Part1(maze)))

	maze = Utils.ReadFileByLinesString("//Day20//maze.txt")
	fmt.Println("Result part two: " + strconv.Itoa(Day20Part2(maze)))
}

func Day20Part1(mazeLines []string) int {
	portals := map[[2]int][2]int{}
	var startPosition [2]int
	var endPosition [2]int

	possibleDirections := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	forewardDirections := []bool{true, true, false, false}

	valid := func(position [2]int) bool {
		return position[0] >= 0 && position[0] < len(mazeLines) && position[1] >= 0 && position[1] < len(mazeLines[position[0]])
	}

	tempPortals := map[[2]byte][2]int{}

	for i, line := range mazeLines {
		for j := 0; j < len(line); j++ {
			if mazeLines[i][j] != '.' {
				continue
			}
			position := [2]int{i, j}
			for k, dir := range possibleDirections {
				fwd := forewardDirections[k]

				p1 := [2]int{i + dir[0], j + dir[1]}
				p2 := [2]int{i + dir[0]*2, j + dir[1]*2}
				if !valid(p1) || !valid(p2) {
					continue
				}

				c1 := mazeLines[p1[0]][p1[1]]
				c2 := mazeLines[p2[0]][p2[1]]

				if c1 < 'A' || c1 > 'Z' || c2 < 'A' || c2 > 'Z' {
					continue
				}

				var label [2]byte
				if fwd {
					label = [2]byte{c1, c2}
				} else {
					label = [2]byte{c2, c1}
				}

				if label == [2]byte{'A', 'A'} {
					startPosition = position
					continue
				}
				if label == [2]byte{'Z', 'Z'} {
					endPosition = position
					continue
				}

				if ex, ok := tempPortals[label]; ok {
					portals[ex] = position
					portals[position] = ex
					delete(tempPortals, label)
				} else {
					tempPortals[label] = position
				}
			}
		}
	}

	visited := map[[2]int]int{}
	work := [][3]int{{startPosition[0], startPosition[1], 0}}

	for len(work) > 0 {
		currentPos := work[0]
		work = work[1:]

		currentPosition := [2]int{currentPos[0], currentPos[1]}

		if currentPosition == endPosition {
			return currentPos[2]
		}

		if dist, ok := visited[currentPosition]; ok && dist <= currentPos[2] {
			continue
		}
		visited[currentPosition] = currentPos[2]

		for _, v := range possibleDirections {
			nextPos := [3]int{currentPos[0] + v[0], currentPos[1] + v[1], currentPos[2] + 1}
			nextPosition := [2]int{nextPos[0], nextPos[1]}

			if !valid(nextPosition) {
				continue
			}
			if mazeLines[nextPosition[0]][nextPosition[1]] != '.' {
				continue
			}
			work = append(work, nextPos)
		}
		if jump, ok := portals[currentPosition]; ok {
			work = append(work, [3]int{jump[0], jump[1], currentPos[2] + 1})
		}
	}
	return -1
}

func Day20Part2(mazeLines []string) int {
	portals := map[[2]int][2]int{}
	var startPosition [2]int
	var endPosition [2]int

	possibleDirections := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	forewardDirections := []bool{true, true, false, false}

	valid := func(position [2]int) bool {
		return position[0] >= 0 && position[0] < len(mazeLines) && position[1] >= 0 && position[1] < len(mazeLines[position[0]])
	}

	isInner := func(p [2]int) bool {
		return p[0] > 6 && p[0] < len(mazeLines)-6 && p[1] > 6 && p[1] < len(mazeLines[4])-6
	}

	tempPortals := map[[2]byte][2]int{}

	for i, line := range mazeLines {
		for j := 0; j < len(line); j++ {
			if mazeLines[i][j] != '.' {
				continue
			}
			position := [2]int{i, j}
			for k, dir := range possibleDirections {
				fwd := forewardDirections[k]

				p1 := [2]int{i + dir[0], j + dir[1]}
				p2 := [2]int{i + dir[0]*2, j + dir[1]*2}
				if !valid(p1) || !valid(p2) {
					continue
				}

				c1 := mazeLines[p1[0]][p1[1]]
				c2 := mazeLines[p2[0]][p2[1]]

				if c1 < 'A' || c1 > 'Z' || c2 < 'A' || c2 > 'Z' {
					continue
				}

				var label [2]byte
				if fwd {
					label = [2]byte{c1, c2}
				} else {
					label = [2]byte{c2, c1}
				}

				if label == [2]byte{'A', 'A'} {
					startPosition = position
					continue
				}
				if label == [2]byte{'Z', 'Z'} {
					endPosition = position
					continue
				}

				if ex, ok := tempPortals[label]; ok {
					portals[ex] = position
					portals[position] = ex
					delete(tempPortals, label)
				} else {
					tempPortals[label] = position
				}
			}
		}
	}

	endPos := [3]int{endPosition[0], endPosition[1], 0}
	visited := map[[3]int]int{}
	work := [][4]int{{startPosition[0], startPosition[1], 0, 0}}

	for len(work) > 0 {
		currentPos := work[0]
		work = work[1:]

		currentPosition := [3]int{currentPos[0], currentPos[1], currentPos[2]}

		if currentPosition == endPos {
			return currentPos[3]
		}

		if dist, ok := visited[currentPosition]; ok && dist <= currentPos[3] {
			continue
		}
		visited[currentPosition] = currentPos[3]

		for _, v := range possibleDirections {
			nextPos := [4]int{currentPos[0] + v[0], currentPos[1] + v[1], currentPos[2], currentPos[3] + 1}
			nextPosition := [2]int{nextPos[0], nextPos[1]}

			if !valid(nextPosition) {
				continue
			}
			if mazeLines[nextPosition[0]][nextPosition[1]] != '.' {
				continue
			}
			work = append(work, nextPos)
		}

		currentPosTemp := [2]int{currentPos[0], currentPos[1]}
		if jump, ok := portals[currentPosTemp]; ok && currentPos[2] > 0 || isInner(currentPosTemp) {
			recursionDifference := -1
			if isInner(currentPosTemp) {
				recursionDifference = 1
			}
			work = append(work, [4]int{jump[0], jump[1], currentPos[2] + recursionDifference, currentPos[3] + 1})
		}
	}
	return -1
}
