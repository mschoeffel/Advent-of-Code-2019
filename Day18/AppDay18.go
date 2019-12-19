package Day18

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"math"
	"strconv"
)

func Main() {
	maze := Utils.ReadFileByLinesString("//Day18//maze.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day18Part1(maze)))

	maze = Utils.ReadFileByLinesString("//Day18//maze2.txt")
	fmt.Println("Result part two: " + strconv.Itoa(Day18Part2(maze)))
}

func Day18Part1(mazeLines []string) int {
	var startPosition [2]int
	var doorPositions [26][2]int
	var keyPositions [26][2]int
	numberOfKeysInMaze := 0

	for lineNumber, line := range mazeLines {
		for charNumber := 0; charNumber < len(line); charNumber++ {
			if line[charNumber] == '@' {
				startPosition = [2]int{lineNumber, charNumber}
			} else if line[charNumber] >= 'a' && line[charNumber] <= 'z' {
				keyPositions[line[charNumber]-'a'] = [2]int{lineNumber, charNumber}
				numberOfKeysInMaze++
			} else if line[charNumber] >= 'A' && line[charNumber] <= 'Z' {
				doorPositions[line[charNumber]-'A'] = [2]int{lineNumber, charNumber}
			}
		}
	}

	//Stores to each position(x,y), foundKeys and targetKey the distance
	positionMemory := map[[4]int]int{}

	//Calculates the min distance between position and targetKey
	findMinPath := func(currentPosition [2]int, foundKeys int, targetKeyNumber int) int {
		currentPositionData := [4]int{currentPosition[0], currentPosition[1], foundKeys, targetKeyNumber}
		if v, ok := positionMemory[currentPositionData]; ok {
			return v
		}

		//[0]posx, [1]posy, [2]moves made to get to this currentPosition
		nextPossibleMoves := [][3]int{{currentPosition[0], currentPosition[1], 0}}
		visitedPositions := map[[2]int]bool{}
		targetPosition := keyPositions[targetKeyNumber]
		directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		for len(nextPossibleMoves) > 0 {
			nextMove := nextPossibleMoves[0]
			nextPossibleMoves = nextPossibleMoves[1:]
			positionNextMove := [2]int{nextMove[0], nextMove[1]}
			if positionNextMove == targetPosition {
				positionMemory[currentPositionData] = nextMove[2]
				return nextMove[2]
			}
			if visitedPositions[positionNextMove] {
				continue
			}
			visitedPositions[positionNextMove] = true
			for _, direction := range directions {
				nextPossiblePosition := [3]int{nextMove[0] + direction[0], nextMove[1] + direction[1], nextMove[2] + 1}
				if nextPossiblePosition[0] < 0 || nextPossiblePosition[0] >= len(mazeLines) || nextPossiblePosition[1] < 0 || nextPossiblePosition[1] >= len(mazeLines[0]) {
					continue
				}
				c := mazeLines[nextPossiblePosition[0]][nextPossiblePosition[1]]
				if c == '#' {
					continue
				}
				b := 1 << uint(c-'a')
				if c >= 'a' && c <= 'z' && int(c-'a') != targetKeyNumber && (foundKeys&b) == 0 {
					continue
				}
				b = 1 << uint(c-'A')
				if c >= 'A' && c <= 'Z' && (foundKeys&b) == 0 {
					continue
				}
				nextPossibleMoves = append(nextPossibleMoves, nextPossiblePosition)
			}
		}
		positionMemory[currentPositionData] = -1
		return -1
	}

	min := math.MaxInt64
	mins := map[[3]int]int{}

	var findShortestPathCollectingAllKeys func([2]int, int, int)
	findShortestPathCollectingAllKeys = func(currentPosition [2]int, alreadyFoundKeys int, distanceGone int) {
		if alreadyFoundKeys == (1<<numberOfKeysInMaze)-1 {
			if distanceGone < min {
				min = distanceGone
			}
			return
		}
		if distanceGone >= min {
			return
		}

		key := [3]int{currentPosition[0], currentPosition[1], alreadyFoundKeys}
		if v, ok := mins[key]; ok && v <= distanceGone {
			return
		}
		mins[key] = distanceGone

		for nextKey := 0; nextKey < 26; nextKey++ {
			bit := 1 << uint(nextKey)
			if alreadyFoundKeys&bit != 0 {
				continue
			}
			distanceToNextKey := findMinPath(currentPosition, alreadyFoundKeys, nextKey)
			if distanceToNextKey == -1 {
				continue
			}
			findShortestPathCollectingAllKeys(keyPositions[nextKey], alreadyFoundKeys|bit, distanceGone+distanceToNextKey)
		}
	}
	findShortestPathCollectingAllKeys(startPosition, 0, 0)
	return min
}

func Day18Part2(mazeLines []string) int {

	numberOfPlayersInMaze := 0
	numberOfKeysInMaze := 0
	var startPositionPlayers [4][2]int
	var doorPositions [26][2]int
	var keyPositions [26][2]int

	for lineNumber, line := range mazeLines {
		for charNumber := 0; charNumber < len(line); charNumber++ {
			if line[charNumber] == '@' {
				startPositionPlayers[numberOfPlayersInMaze] = [2]int{lineNumber, charNumber}
				numberOfPlayersInMaze++
			} else if line[charNumber] >= 'a' && line[charNumber] <= 'z' {
				keyPositions[line[charNumber]-'a'] = [2]int{lineNumber, charNumber}
				numberOfKeysInMaze++
			} else if line[charNumber] >= 'A' && line[charNumber] <= 'Z' {
				doorPositions[line[charNumber]-'A'] = [2]int{lineNumber, charNumber}
			}
		}
	}

	positionMemory := map[[4]int]int{}

	findMinPath := func(currentPosition [2]int, foundKeys int, targetKeyNumber int) int {
		currentPositionData := [4]int{currentPosition[0], currentPosition[1], foundKeys, targetKeyNumber}
		if v, ok := positionMemory[currentPositionData]; ok {
			return v
		}

		nextPossibleMoves := [][3]int{{currentPosition[0], currentPosition[1], 0}}
		visitedPositions := map[[2]int]bool{}
		targetPosition := keyPositions[targetKeyNumber]
		directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		for len(nextPossibleMoves) > 0 {
			nextMove := nextPossibleMoves[0]
			nextPossibleMoves = nextPossibleMoves[1:]
			positionNextMove := [2]int{nextMove[0], nextMove[1]}
			if positionNextMove == targetPosition {
				positionMemory[currentPositionData] = nextMove[2]
				return nextMove[2]
			}
			if visitedPositions[positionNextMove] {
				continue
			}
			visitedPositions[positionNextMove] = true
			for _, dir := range directions {
				nextPossiblePosition := [3]int{nextMove[0] + dir[0], nextMove[1] + dir[1], nextMove[2] + 1}
				if nextPossiblePosition[0] < 0 || nextPossiblePosition[0] >= len(mazeLines) || nextPossiblePosition[1] < 0 || nextPossiblePosition[1] >= len(mazeLines[0]) {
					continue
				}
				c := mazeLines[nextPossiblePosition[0]][nextPossiblePosition[1]]
				if c == '#' {
					continue
				}
				b := 1 << uint(c-'a')
				if c >= 'a' && c <= 'z' && int(c-'a') != targetKeyNumber && (foundKeys&b) == 0 {
					continue
				}
				b = 1 << uint(c-'A')
				if c >= 'A' && c <= 'Z' && (foundKeys&b) == 0 {
					continue
				}
				nextPossibleMoves = append(nextPossibleMoves, nextPossiblePosition)
			}
		}
		positionMemory[currentPositionData] = -1
		return -1
	}

	keysSegments := [4]int{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 4; j++ {
			if findMinPath(startPositionPlayers[j], (1<<numberOfKeysInMaze)-1, i) != -1 {
				keysSegments[j] |= 1 << uint(i)
			}
		}
	}

	min := math.MaxInt64

	mins := map[[9]int]int{}

	var findShortestPathCollectingAllKeys func([4][2]int, int, int)
	findShortestPathCollectingAllKeys = func(currentPosition [4][2]int, alreadyFoundKeys int, distanceGone int) {
		if alreadyFoundKeys == (1<<numberOfKeysInMaze)-1 {
			if distanceGone < min {
				min = distanceGone
			}
			return
		}
		if distanceGone >= min {
			return
		}

		key := [9]int{currentPosition[0][0], currentPosition[0][1], currentPosition[1][0], currentPosition[1][1], currentPosition[2][0], currentPosition[2][1], currentPosition[3][0], currentPosition[3][1], alreadyFoundKeys}
		if v, ok := mins[key]; ok && v <= distanceGone {
			return
		}
		mins[key] = distanceGone

		for nextKey := 0; nextKey < 26; nextKey++ {
			bit := 1 << uint(nextKey)
			if alreadyFoundKeys&bit != 0 {
				continue
			}
			var bi int
			for i := 0; i < 4; i++ {
				if keysSegments[i]&bit != 0 {
					bi = i
				}
			}
			npos := currentPosition
			npos[bi] = keyPositions[nextKey]
			distanceToNextKey := findMinPath(currentPosition[bi], alreadyFoundKeys, nextKey)
			if distanceToNextKey == -1 {
				continue
			}
			findShortestPathCollectingAllKeys(npos, alreadyFoundKeys|bit, distanceGone+distanceToNextKey)
		}
	}
	findShortestPathCollectingAllKeys(startPositionPlayers, 0, 0)
	return min
}
