package Day08

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"log"
	"strconv"
)

type Layer struct {
	data [6][25]int
}

func Main() {
	data := Utils.ReadFileByLinesString("\\Day08\\imagedata.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day8Part1(data[0])))

	data = Utils.ReadFileByLinesString("\\Day08\\imagedata.txt")
	fmt.Println("Result part two: ")
	Day8Part2(data[0])
}

func Day8Part1(imagedata string) int {
	width := 25
	height := 6

	layer := createLayers(imagedata, width, height)

	layerFewestZero := returnLayerWithFewestZero(layer)

	number1 := countDigitOfLayer(layerFewestZero, 1)
	number2 := countDigitOfLayer(layerFewestZero, 2)

	return number1 * number2
}

func Day8Part2(imagedata string) {
	width := 25
	height := 6

	layer := createLayers(imagedata, width, height)
	printlayer(layer, width, height)
}

func printlayer(layers []Layer, width int, height int) {
	black := " "
	white := "#"
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			for i := range layers {
				pixel := layers[i].data[x][y]
				if pixel == 0 {
					fmt.Print(black)
					break
				} else if pixel == 1 {
					fmt.Print(white)
					break
				}
			}
		}
		fmt.Print("\n")
	}
}

func createLayers(data string, width int, height int) []Layer {
	layersize := width * height
	runes := []rune(data)
	layers := []string{""}
	index := 0
	for i := 0; i < len(runes); i++ {
		if i%(layersize) == 0 && i != 0 {
			index++
			layers = append(layers, "")
		}
		layers[index] = layers[index] + string(runes[i])
	}

	var layer []Layer
	for i := range layers {
		layer = append(layer, createLayer(layers[i], width, height))
	}
	return layer
}

func createLayer(data string, width int, height int) Layer {
	runes := []rune(data)
	test := [6][25]int{}
	layer := Layer{data: test}
	index := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			x, err := strconv.ParseInt(string(runes[index]), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			layer.data[h][w] = int(x)
			index++
		}
	}
	return layer
}

func returnLayerWithFewestZero(layers []Layer) Layer {
	minZeroCount := 0
	minLayer := Layer{}
	for i := range layers {
		layer := layers[i]
		layerZeroCount := countDigitOfLayer(layer, 0)
		if layerZeroCount < minZeroCount || minZeroCount == 0 {
			minZeroCount = layerZeroCount
			minLayer = layer
		}
	}
	return minLayer
}

func countDigitOfLayer(layer Layer, digit int) int {
	digitCount := 0
	for x := range layer.data {
		for y := range layer.data[x] {
			if layer.data[x][y] == digit {
				digitCount++
			}
		}
	}
	return digitCount
}
