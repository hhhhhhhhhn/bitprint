package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var quadrants = []string{
	// Bits in left to right, top to bottom order
	" ", // 0b0000
	"▘", // 0b0001
	"▝", // 0b0010
	"▀", // 0b0011
	"▖", // 0b0100
	"▌", // 0b0101
	"▞", // 0b0110
	"▛", // 0b0111
	"▗", // 0b1000
	"▚", // 0b1001
	"▐", // 0b1010
	"▜", // 0b1011
	"▄", // 0b1100
	"▙", // 0b1101
	"▟", // 0b1110
	"█", // 0b1111
}

type Font map[rune][][]int

func PrintGrid(grid [][]int) {
	for y := 0; y < len(grid); y += 2 {
		for x := 0; x < len(grid[0]); x += 2 {
			var quadrant uint = 0
			if  grid[y][x] != 0 {
				quadrant = quadrant | 0b0001
			}
			if len(grid[0]) > x + 1 && grid[y][x + 1] != 0 {
				quadrant = quadrant | 0b0010
			}
			if len(grid) > y + 1 && grid[y + 1][x] != 0 {
				quadrant = quadrant | 0b0100
			}
			if len(grid[0]) > x + 1 && len(grid) > y + 1 && grid[y + 1][x + 1] != 0 {
				quadrant = quadrant | 0b1000
			}
			fmt.Print(quadrants[quadrant])
		}
		fmt.Print("\n")
	}
}

// TODO: Flag for this
func WidenGrid(original [][]int) (new [][]int) {
	for _, row := range original {
		newRow := []int{}
		for _, value := range row {
			newRow = append(newRow, value, value)
		}
		new = append(new, newRow)
	}
	return new
}

func TextToGrid(text string, font Font) [][]int {
	height := len(font[' '])

	grid := make([][]int, height)

	for _, chr := range text {
		chrGrid := font[chr]
		for i, row := range chrGrid {
			grid[i] = append(grid[i], row...)
		}
	}
	return grid
}

func main() {
	in, _ := ioutil.ReadAll(os.Stdin)
	for _, line := range strings.Split(string(in), "\n") {
		grid := TextToGrid(strings.ReplaceAll(line, "\t", "    "), TomThumb)
		PrintGrid(grid)
	}
}
