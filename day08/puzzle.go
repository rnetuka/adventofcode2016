package day08

import (
	"aoc2016/utils"
	"strconv"
	"strings"
)

type Screen struct {
	width  int
	height int
	pixels [6][50]bool
}

func NewScreen() *Screen {
	screen := new(Screen)
	screen.width = 50
	screen.height = 6
	return screen
}

func (screen *Screen) LightUp(width int, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			screen.pixels[i][j] = true
		}
	}
}

func (screen *Screen) RotateRow(i int, times int) {
	for t := 0; t < times; t++ {
		off := screen.pixels[i][screen.width - 1]
		for j := screen.width - 1; j > 0; j-- {
			screen.pixels[i][j] = screen.pixels[i][j - 1]
		}
		screen.pixels[i][0] = off
	}
}

func (screen *Screen) RotateColumn(j int, times int) {
	for t := 0; t < times; t++ {
		off := screen.pixels[screen.height - 1][j]
		for i := screen.height - 1; i > 0; i-- {
			screen.pixels[i][j] = screen.pixels[i - 1][j]
		}
		screen.pixels[0][j] = off
	}
}

func (screen *Screen) row(i int) string {
	row := strings.Builder{}
	for j := 0; j < screen.width; j++ {
		row.WriteByte(pixelToByte(screen.pixels[i][j]))
	}
	return row.String()
}

func (screen *Screen) ToString() string {
	result := strings.Builder{}
	for i := 0; i < screen.height; i++ {
		result.WriteString(screen.row(i))
		result.WriteString("\n")
	}
	return result.String()
}

func pixelToByte(pixel bool) byte {
	if pixel {
		return '@'
	} else {
		return ' '
	}
}

func (screen *Screen) Process(instructions []string) {
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "rect") {
			data := instruction[len("rect "):]
			parts := strings.Split(data, "x")
			width, _ := strconv.Atoi(parts[0])
			height, _ := strconv.Atoi(parts[1])
			screen.LightUp(width, height)
		} else if strings.HasPrefix(instruction, "rotate row") {
			data := instruction[len("rotate row "):]
			parts := strings.Split(data, " by ")
			i, _ := strconv.Atoi(parts[0][len("y="):])
			times, _ := strconv.Atoi(parts[1])
			screen.RotateRow(i, times)
		} else if strings.HasPrefix(instruction, "rotate column") {
			data := instruction[len("rotate column "):]
			parts := strings.Split(data, " by ")
			j, _ := strconv.Atoi(parts[0][len("y="):])
			times, _ := strconv.Atoi(parts[1])
			screen.RotateColumn(j, times)
		}
	}
}

func Solution1() int {
	instructions := utils.ReadLines("assets/input08.txt")
	screen := NewScreen()
	screen.Process(instructions)

	litPixels := 0
	for i := 0; i < screen.height; i++ {
		for j := 0; j < screen.width; j++ {
			if screen.pixels[i][j] {
				litPixels++
			}
		}
	}
	return litPixels
}

func Solution2() string {
	instructions := utils.ReadLines("assets/input08.txt")
	screen := NewScreen()
	screen.Process(instructions)
	return screen.ToString()
}