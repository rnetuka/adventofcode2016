package day10

import (
	"strconv"
	"strings"
)

type Output struct {
	microchips 	 []int
}

func (output *Output) Add(microchip int) {
	output.microchips = append(output.microchips, microchip)
}

func (output *Output) String() string {
	var tokens []string
	for _, microchip := range output.microchips {
		tokens = append(tokens, strconv.Itoa(microchip))
	}
	return "[" + strings.Join(tokens, ", ") + "]"
}