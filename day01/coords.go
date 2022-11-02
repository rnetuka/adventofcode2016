package day01

import "aoc2016/utils"

type Coords struct {
	x, y int
}

func (coords *Coords) add(vector Vector) {
	coords.x += vector.dx
	coords.y += vector.dy
}

func (coords Coords) distanceFrom(x, y int) int {
	dx := utils.Abs(coords.x - x)
	dy := utils.Abs(coords.y - y)
	return dx + dy
}
