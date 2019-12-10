package Day10

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Asteroit struct {
	x        int
	y        int
	gradient float64
	distance float64
}

type ByDegree []Asteroit

func (a ByDegree) Len() int           { return len(a) }
func (a ByDegree) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDegree) Less(i, j int) bool { return a[i].gradient < a[j].gradient }

func Main() {
	data := Utils.ReadFileByLinesString("\\Day10\\map.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day10Part1(data)))

	data = Utils.ReadFileByLinesString("\\Day10\\map.txt")
	fmt.Println("Result part two: " + strconv.Itoa(Day10Part2(data)))
}

func Day10Part1(data []string) int {
	spacemap := generateMap(data)
	asteroits := getAsteroits(spacemap)
	return calculateBestBase(asteroits)
}

func Day10Part2(data []string) int {
	spacemap := generateMap(data)
	asteroits := getAsteroits(spacemap)
	base := getBestBase(asteroits)
	asteroitsAlive := calculateMap(asteroits, base)
	count := 0
	for len(asteroitsAlive)-1 > 0 {
		asteroitsView := getView(asteroitsAlive, base)
		for i := range asteroitsView {
			count++
			if count == 200 {
				return asteroitsView[i].x*100 + asteroitsView[i].y
			}
			//fmt.Println(strconv.Itoa(count) + " Removed: (" + strconv.Itoa(asteroitsView[i].x) + "," + strconv.Itoa(asteroitsView[i].y) + ") Gradient: " + fmt.Sprintf("%f", asteroitsView[i].gradient))
		}
		asteroitsAlive = removeAsteroits(asteroitsAlive, asteroitsView)
	}
	return 000
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
		if diffx == 0 && diffy == 0 {
			asteroits[i].gradient = 0
		} else {
			asteroits[i].gradient = math.Round(angle*1000) / 1000
		}
		asteroits[i].distance = math.Sqrt(diffy*diffy + diffx*diffx)
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

func getView(asteroits []Asteroit, source Asteroit) []Asteroit {
	checkedView := make(map[float64]Asteroit)
	var asteroitsView []Asteroit
	for i := range asteroits {
		checkAsteroit := asteroits[i]
		if !(checkAsteroit.x == source.x && checkAsteroit.y == source.y) {
			i, ok := checkedView[checkAsteroit.gradient]
			if ok {
				if i.distance > checkAsteroit.distance {
					checkedView[checkAsteroit.gradient] = checkAsteroit
				}
			} else {
				checkedView[checkAsteroit.gradient] = checkAsteroit
			}
		}
	}
	for _, value := range checkedView {
		asteroitsView = append(asteroitsView, value)
	}
	sort.Sort(ByDegree(asteroitsView))
	return asteroitsView
}

func getBestBase(asteroits []Asteroit) Asteroit {
	maxCount := 0
	bestBase := Asteroit{}
	for i := range asteroits {
		sourceAsteroit := asteroits[i]
		calculatedMap := calculateMap(asteroits, sourceAsteroit)
		count := countView(calculatedMap, sourceAsteroit)
		if count > maxCount {
			maxCount = count
			bestBase = sourceAsteroit
		}
	}
	return bestBase
}

func removeAsteroits(asteroitsdest []Asteroit, asteroitssrc []Asteroit) []Asteroit {
	for i := range asteroitssrc {
		x := asteroitssrc[i].x
		y := asteroitssrc[i].y
		for j := range asteroitsdest {
			if asteroitsdest[j].x == x && asteroitsdest[j].y == y {
				asteroitsdest = append(asteroitsdest[:j], asteroitsdest[j+1:]...)
				break
			}
		}
	}
	return asteroitsdest
}

func contains(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
