package Day10

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"math"
	"strconv"
)

type Asteroit struct {
	x        int
	y        int
	gradient float64
	distance float64
}

func Main() {
	data := Utils.ReadFileByLinesString("\\Day10\\map.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day10Part1(data)))
}

func Day10Part1(data []string) int {
	spacemap := generateMap(data)
	asteroits := getAsteroits(spacemap)
	return calculateBestBase(asteroits)
}

func generateMap(data []string) [][]string {
	var spacemap [][]string
	for i := range data {
		runes := []rune(data[i])
		var spacemapTemp []string
		for j := range runes {
			spacemapTemp = append(spacemapTemp, string(runes[j]))
		}
		spacemap = append(spacemap, spacemapTemp)
	}
	return spacemap
}

func getAsteroits(spacemap [][]string) []Asteroit {
	var asteroits []Asteroit
	for y := range spacemap {
		for x := range spacemap[y] {
			if spacemap[y][x] == "#" {
				newAsteroit := Asteroit{x: x, y: y}
				asteroits = append(asteroits, newAsteroit)
			}
		}
	}
	return asteroits
}

func calculateBestBase(asteroits []Asteroit) int {
	maxCount := 0
	for i := range asteroits {
		sourceAsteroit := asteroits[i]
		calculatedMap := calculateMap(asteroits, sourceAsteroit)
		count := countView(calculatedMap, sourceAsteroit)
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func calculateMap(asteroits []Asteroit, source Asteroit) []Asteroit {
	for i := range asteroits {
		if asteroits[i] != source {
			diffx := float64(asteroits[i].x - source.x)
			diffy := float64(asteroits[i].y - source.y)
			y1 := float64(asteroits[i].y)
			y2 := float64(source.y)
			x1 := float64(asteroits[i].x)
			x2 := float64(source.x)
			top := math.Max(y1, y2) - math.Min(y1, y2)
			bot := math.Sqrt((x1-x2)*(x1-x2) + (top * top))
			angle := math.Acos(top/bot) * (180 / math.Pi)
			if diffx > 0 && diffy < 0 {

			} else if diffx >= 0 && diffy > 0 {
				angle = 180 - angle
			} else if diffx < 0 && diffy >= 0 {
				angle = 180 + angle
			} else if diffx < 0 && diffy < 0 {
				angle = 360 - angle
			}
			asteroits[i].gradient = math.Round(angle*1000) / 1000
			asteroits[i].distance = math.Sqrt(diffy*diffy + diffx*diffx)
		}
	}
	return asteroits
}

func countView(asteroits []Asteroit, source Asteroit) int {
	var gradients []float64
	count := 0
	for i := range asteroits {
		checkAsteroit := asteroits[i]
		if checkAsteroit != source {
			if !contains(gradients, checkAsteroit.gradient) {
				gradients = append(gradients, checkAsteroit.gradient)
				count++
			}
		}
	}
	return count
}

func contains(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
